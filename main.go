package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	pass "github.com/sethvargo/go-password/password"
)

// pwgen struct represents a password generation request
// attributes are defined as follows:
// MinLen: minimum length
// Schar: number of special characters in the password
// Nchar: number of numbers in the password
// Num: number of passwords that must be created.
type pwgen struct {
	MinLen int `json:"minLen"`
	Schar  int `json:"schar"`
	Nchar  int `json:"nchar"`
	Num    int `json:"num"`
}

// maxLen is an upper limit on password length; as the request
// only requires to set a minimum, this constant limits the resulting
// length, to create "reasonably long" passwords
const maxLen int = 65

// pwLength generates a random integer `n` to use as
// password length, where `minLen <= n < maxLen`
func (pg *pwgen) pwLength() int {
	rand := rand.Intn(maxLen)
	if pg.MinLen > rand {
		return pg.MinLen
	} else {
		return rand
	}
}

// generate performs the password generation
// by using the `go-password` library as many times
// as requested
func (pg *pwgen) generate() []string {
	var pwds []string

	max := pg.pwLength()

	for i := 0; i < pg.Num; i++ {
		res, err := pass.Generate(max, pg.Nchar, pg.Schar, false, false)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		pwds = append(pwds, res)
	}

	return pwds
}

// postPw implements the HTTP handler
// containing the password generation request;
// expects to receive a valid JSON matching the `pwgen` struct
func postPw(c *gin.Context) {
	var request pwgen

	// process body data and map JSON keys
	// to struct fields
	err := c.BindJSON(&request)
	if err != nil {
		log.Fatal(err)
		return
	}

	// calls the generator and returns an
	// array of passwords
	pwds := request.generate()
	if err != nil {
		log.Fatal(err)
		return
	}

	// process the array as valid JSON response
	c.JSON(http.StatusCreated, pwds)
}

func main() {
	// initialize HTTP router and perform the mapping
	// of URI + method + handler function
	router := gin.Default()
	router.POST("/pwgen", postPw)

	// run the HTTP server
	router.Run(":8080")
}
