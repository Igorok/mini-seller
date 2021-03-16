package main

import (
	"fmt"
	"log"
	"mini-seller/infrastructure/mongohelper"
	"mini-seller/infrastructure/viperhelper"
	"mini-seller/web/customer/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	// default
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	vip := viperhelper.Viper{ConfigType: "", ConfigName: "", ConfigPath: "infrastructure/viperhelper"}
	vip.Read()

	fmt.Println("WEB_PORT", viper.GetString("WEB_PORT"))

	db, err := mongohelper.Connect()
	if err != nil {
		log.Fatal(err)
	}

	gqlHandler := handlers.GetHandler(db)

	r.GET("/graphql", gqlHandler)
	r.POST("/graphql", gqlHandler)

	http.ListenAndServe(":"+viper.GetString("WEB_PORT"), r)
}
