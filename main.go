package main

import (
	"ErrorHandling/Message"
	"ErrorHandling/middleware"
	"errors"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	r := gin.Default()
	// Ensure to set this middleware before others in order to catch any panics that occur
	r.Use(middleware.ErrorHandling())

	r.GET("/", func(c *gin.Context) {
		panic(Message.PanicMessage{MessageKey: 2})
	})

	// You can also pass the error in oder to log the error
	r.GET("/error", func(c *gin.Context) {
		err := errors.New("test error")
		panic(Message.PanicMessage{3, &err})
	})

	r.Run("0.0.0.0:8080")
}
