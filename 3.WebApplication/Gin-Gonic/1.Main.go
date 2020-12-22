package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)
func main() {
	fmt.Println("Running server.....")
	r := gin.Default()
	r.GET("/", homePage)
	r.POST("/", postPage )
	r.POST("/post", postBody )
	r.GET("/query", queryHand )
	r.GET("/path/:name/:id", pathParams)
	r.Run()
}
func homePage(c *gin.Context){
	c.JSON(200, gin.H{
		"msg" : "Hello Gopher",
	})
}
func postPage(c *gin.Context){

	c.JSON(200,  gin.H{
		"msg":"From Post",
	})
}
func postBody(c *gin.Context){
	 body := c.Request.Body
	 val , err := ioutil.ReadAll(body)
	 if err != nil{
		 panic(err.Error())
	 }
	c.JSON(200,  gin.H{
		"msg": string(val),
	})
}
func queryHand(c *gin.Context){
  name :=c.Query("name")
  id   :=c.Query("id")

  c.JSON(200, gin.H{
	  "name": name,
	  "id" : id,
  })
}
func pathParams(c *gin.Context){
	name := c.Param("name")
	id := c.Param("id")
	c.JSON(200, gin.H{
		"name": name,
		"id" : id,
	} )
}