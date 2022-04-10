package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"encoding/json"
	"io"
	"bytes"
)

var db *sql.DB

type Response struct {
	Name    string `json:"name"`
	Number  string `json:"phone_number"`
	City    string `json:"city"`
	State   string `json:"state"`
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	Zip     string `json:"zip_code"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	var err error
	db, err = sql.Open("mysql", "username:password@/cetec")
	fmt.Printf("Done %T", db)
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing

	defer db.Close()
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/person/:Id/info", getPersonById)
	router.POST("/person/create", createPerson)

	router.Run("localhost:8080")

}

func getPersonById(c *gin.Context) {
	id := c.Param("Id")

	results := db.QueryRow("SELECT name for person where Id= ?", id)
	Resp := Response{}
	err := results.Scan(&Resp.Name)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"}) // proper error handling instead of panic in your app
	} else {

		results = db.QueryRow("SELECT number for phone where person_id= ?", id)
		err := results.Scan(&Resp.Number)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Number not found"})
		} else {
			results = db.QueryRow("Select city,state,street1,street2,zip_code from address where id = ?", id)
			err := results.Scan(&Resp.City, &Resp.State, &Resp.Street1, &Resp.Street2, &Resp.Zip)
			if err != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Address not found"})
			} else {
				c.IndentedJSON(http.StatusOK, Resp)
			}

		}
	}
}

func createPerson(c *gin.Context) {

	var requestBody Response
	bytes := StreamToByte(c.Request.Body)
   if err := json.Unmarshal(bytes,&requestBody); err != nil {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
   }else{
	  	
   }

  fmt.Println(requestBody)
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	  buf.ReadFrom(stream)
	  return buf.Bytes()
  }
  
