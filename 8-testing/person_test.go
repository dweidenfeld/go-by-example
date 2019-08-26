package person

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFindPerson(t *testing.T) {
	t.Run("should find a person", func(t *testing.T) {
		person, err := findPerson("dominik")

		if err != nil {
			t.Error("could not find dominik")
		}

		if person != nil && person.FirstName != "Dominik" {
			t.Error("wrong match")
		}
	})

	t.Run("should not find a person", func(t *testing.T) {
		person, err := findPerson("notfound")

		if err == nil {
			t.Error("falsely found a person")
		}

		if person != nil {
			t.Error("falsely returned a person")
		}
	})
}

func TestPersonEndpoint(t *testing.T) {
	r := gin.Default()
	r.GET("/person/:name", personEndpoint)

	doRequest := func(name string) (*httptest.ResponseRecorder, error) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/person/%v", name), nil)
		if err != nil {
			return nil, err
		}

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		return w, nil
	}

	t.Run("should find a person", func(t *testing.T) {
		w, err := doRequest("dominik")
		if err != nil {
			t.Error(err)
		}

		if w.Code != http.StatusOK {
			t.Error("expected response to be 200")
		}

		var person Person
		if err := json.Unmarshal(w.Body.Bytes(), &person); err != nil {
			t.Error(err)
		}

		if person.FirstName != "Dominik" {
			t.Error("wrong person returned")
		}
	})

	t.Run("should not find a person", func(t *testing.T) {
		w, err := doRequest("notfound")

		if err != nil {
			t.Error(err)
		}

		if w.Code != http.StatusNotFound {
			t.Error("expected response to be 404")
		}
	})
}
