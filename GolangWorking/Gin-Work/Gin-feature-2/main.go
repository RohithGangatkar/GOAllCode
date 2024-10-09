package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

type JsonInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// XML binding struct
type XmlInput struct {
	Name  string `xml:"name" binding:"required"`
	Email string `xml:"email" binding:"required,email"`
}

// Form binding struct
type FormInput struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required,email"`
}

func main() {

	r := gin.Default()

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",   // usernames and password
		"lena":   "hello2", // when we enter /admin/secerets endpoint asked to enter name and password
		"manu":   "4321",   // which will then print user details
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	//
	r.POST("/bind", func(c *gin.Context) {
		// Try binding as JSON
		var jsonInput JsonInput
		if err := c.ShouldBindBodyWith(&jsonInput, binding.JSON); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "The body is JSON",
				"data":    jsonInput,
			})
			return
		}

		// Try binding as XML
		var xmlInput XmlInput
		if err := c.ShouldBindBodyWith(&xmlInput, binding.XML); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "The body is XML",
				"data":    xmlInput,
			})
			return
		}

		// If none of the bindings were successful
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")

}
