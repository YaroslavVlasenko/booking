#Bookings and Reservations

This is repository for my bookings and reservations project.

- Built in Go version 1.16
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses [alex edwards SCS session management](https://github.com/alexedwards/scs)
- Uses [nosurf](https://github.com/justinas/nosurf)
- Uses [soda](https://gobuffalo.io/en/docs/db/toolbox/)
- Uses [Go Simple Mail](go get github.com/xhit/go-simple-mail/v2)

To remember:
####Run migrations: soda migrate
####Revert migration: soda migrate/down
####Migration generate table: soda generate fizz CreateNameTable
####Migration generate FK/Index: soda generate fizz CreateFKNameTable
####Seed generation: soda generate sql SeedNameTable

####Run the project on Windows: ./run.bat
####Run the project on Linux: ./run.sh

####Run all the tests: go test -v ./...
####Run test in folder with coverage: go test -coverprofile=coverage.out && go tool cover -html=coverage.out

- Migration generate table:
soda generate fizz *CreateNameTable*

- Migration generate FK/Index:
soda generate fizz *CreateFKNameTable*

- Seed generation:
soda generate sql *SeedNameTable*
  
- Run the project:
*./run.bat(Windows)*
*./run.sh(Linux)*
  
- Run all the tests
*go test -v ./...*
  
- Run test in folder with coverage
*go test -coverprofile=coverage.out && go tool cover -html=coverage.out*
  

Soda installation:
  ```
  $ go get github.com/gobuffalo/pop/...
  $ go install github.com/gobuffalo/pop/soda
  ```
  See also [full guide](https://gobuffalo.io/en/docs/db/toolbox/#from-source)
