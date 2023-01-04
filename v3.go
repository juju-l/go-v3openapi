package main

import (
	//
	"github.com/gin-gonic/gin"
	"go-v3openapi/v3"
)

//

func main() {
	r := gin.Default()
	r.GET("/openapi/v3", func(c *gin.Context) { c.File("./openapiv3.yml") }, /**/)//
	r.GET("/swg/openapi/*any", func(c *gin.Context) { v3.OpenapiV3().ServeHTTP(c.Writer,c.Request) }, /**/)//
	r.GET("/", func(c *gin.Context) { c.Redirect(307, "/swg/openapi") }, /**/)//
	r.Run(":8077")
}

func init() {
	//
}