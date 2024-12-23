package main

import (
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	csrf "github.com/utrack/gin-csrf"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store)) // this middleware attaches a session Object to every request. The session Object can store data of request.
	r.Use(csrf.Middleware(csrf.Options{ // this middleware add a token to each request
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) { // if the req do not have valid csrf token, the err func will be called
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))


	// First you need to make a GET request to /protected endpoint to get a CSRF token
	r.GET("/protected", func(c *gin.Context) {
		c.String(200, csrf.GetToken(c))
	})

	// Attach the CSRF token we received from prev. request to send request to /protected path as header:
	// header key: `X-CSRF-TOKEN` value: the value you get from the GET /protected request
	r.POST("/protected", func(c *gin.Context) {
		// If the token is valid, you will receive this:
		c.String(200, "CSRF token is valid")
		// else, you will receive the ErrFunc response:
		// aka c.String(400, "CSRF token mismatch")
	})

	r.Run(":8080")
}