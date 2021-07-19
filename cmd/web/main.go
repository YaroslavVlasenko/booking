package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/yaroslavvlasenko/bookings/internal/config"
	"github.com/yaroslavvlasenko/bookings/internal/driver"
	"github.com/yaroslavvlasenko/bookings/internal/handlers"
	"github.com/yaroslavvlasenko/bookings/internal/helpers"
	"github.com/yaroslavvlasenko/bookings/internal/models"
	"github.com/yaroslavvlasenko/bookings/internal/render"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

var dbConnection struct {
	Host     string
	Port     int
	DataBase string
	User     string
	Password string
}

var httpdConfig struct {
	Port int
}

// main is the entrypoint of application
func main() {
	/* Init database configuration */
	dbConfigFile, err := os.Open("config/db.json")
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}

	dbJsonParser := json.NewDecoder(dbConfigFile)
	if err = dbJsonParser.Decode(&dbConnection); err != nil {
		fmt.Println("parsing config file", err.Error())
	}

	/* Init web server configuration */
	httpdConfigFile, err := os.Open("config/httpd.json")
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}

	httpdJsonParser := json.NewDecoder(httpdConfigFile)
	if err = httpdJsonParser.Decode(&httpdConfig); err != nil {
		fmt.Println("parsing config file", err.Error())
	}

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("Starting application on port %s", httpdConfig.Port))

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(httpdConfig.Port),
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.RoomRestriction{})

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	//connect to database
	db, err := driver.ConnectSql("host=localhost port=" + strconv.Itoa(dbConnection.Port) +
		" dbname=" + dbConnection.DataBase +
		" user=" + dbConnection.User +
		" password=" + dbConnection.Password)
	if err != nil {
		log.Fatal("cannot connect to database! Dying...")
		return nil, err
	}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
