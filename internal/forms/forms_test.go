package forms

import (
	"net/http"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r, err := http.NewRequest("POST", "/whatever", nil)
	if err != nil {
		t.Error(err)
	}
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r, err := http.NewRequest("POST", "/whatever", nil)
	if err != nil {
		t.Error(err)
	}
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, err = http.NewRequest("POST", "/whatever", nil)
	if err != nil {
		t.Error(err)
	}

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("show does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	has := form.Has("whatever")

	if has {
		t.Error("form shows has field when it does")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")

	if !has {
		t.Error("shows form does not have field when it does")
	}

}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent fields")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("whatever", "whatever")
	form = New(postedData)
	form.MinLength("whatever", 100)
	if form.Valid() {
		t.Error("show min length of 100 met when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abs123")
	form = New(postedData)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows min length of 1 is not met when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non existent fields")
	}

	postedData = url.Values{}
	postedData.Add("email", "me@here.com")
	form = New(postedData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form shows not valid email for valid email")
	}

	postedData = url.Values{}
	postedData.Add("email", "x")
	form = New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form shows valid email for not valid email")
	}
}
