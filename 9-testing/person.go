package person

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int8   `json:"age"`
}

var persons map[string]Person

func init() {
	persons = make(map[string]Person, 2)

	persons["dominik"] = Person{
		FirstName: "Dominik",
		LastName:  "Weidenfeld",
		Age:       29,
	}
	persons["holz"] = Person{
		FirstName: "Holz",
		LastName:  "Michel",
		Age:       33,
	}
}

func personEndpoint(c *gin.Context) {
	name := c.Param("name")

	person, err := findPerson(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, person)
}

func findPerson(name string) (*Person, error) {
	person, found := persons[name]
	if !found {
		return nil, errors.New("person not found")
	}
	return &person, nil
}
