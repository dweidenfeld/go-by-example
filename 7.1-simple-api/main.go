package main

import (
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

func main() {
	r := gin.Default()

	r.GET("/health", healthEndpoint)
	r.GET("/person/:name", personEndpoint)

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}

func personEndpoint(c *gin.Context) {
	name := c.Param("name")

	person, found := persons[name]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}

	c.JSON(http.StatusOK, person)
}

func healthEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
