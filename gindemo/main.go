package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func getDemo(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   1,
		"Name": "Sumathy",
	})
}

func postDemo(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	designation := c.Query("designation")
	salary := c.Query("salary")
	c.JSON(200, gin.H{
		"name":        name,
		"designation": designation,
		"salary":      salary,
	})
}

func pathParameters(c *gin.Context) {
	name := c.Param("name")
	designation := c.Param("designation")
	salary := c.Param("salary")
	c.JSON(200, gin.H{
		"name":        name,
		"designation": designation,
		"salary":      salary,
	})

}
func main() {
	fmt.Println("Gin Framework Dmeo")

	r := gin.Default()

	r.GET("/ginget", getDemo) //  func is replaced by getDemo that is defined above
	//r.GET("/gin", func(c *gin.Context) {})
	// 	c.JSON(200, gin.H{
	// 		"id":   1,
	// 		"Name": "Sumathy",
	// 	})
	// })
	r.POST("/ginpost", postDemo)
	//using Querystring where the values to the url are sent as : localhost:8080/query?name=sumathy&designation=developer&salary=50000
	r.GET("/query", QueryString) // query?name=sumathy&salary=50000
	//using pathaparameters where the values to the url are sent as: localhost:8080/path/Sumathy/Developer/50000
	r.GET("/path/:name/:designation/:salary", pathParameters)
	r.Run()
}
