package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

type FruitSelection struct {
	Fruits []string `form:"fruits[]"`
}
type QueryParams struct {
	Name  string `form:"name"`
	Age   int    `form:"age"`
	Email string `form:"email"`
}

func main() {
	r := gin.Default()

	//Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII characters.
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/someExample", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// r.GET("/someProtoBuf", func(c *gin.Context) {
	// reps := []int64{int64(1), int64(2)}
	// label := "test"
	// The specific definition of protobuf is written in the testdata/protoexample file.
	// data := &protoexample.Test{
	// 	Label: &label,
	// 	Reps:  reps,
	// }
	// Note that data becomes binary data in the response
	// Will output protoexample.Test protobuf serialized data
	// c.ProtoBuf(http.StatusOK, data)
	// })

	//Bind Uri to struct
	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})

	//Bind query string or post data
	r.GET("/user", func(c *gin.Context) {
		var queryParams QueryParams

		// Bind the query parameters to the struct
		if err := c.BindQuery(&queryParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"name":  queryParams.Name,
			"age":   queryParams.Age,
			"email": queryParams.Email,
		})
	})

	//-------------------------------Bind html checkboxes-----------------------------------
	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	//Example 1
	// Render the form
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})

	// Handle form submission
	r.POST("/submit", func(c *gin.Context) {
		var selection FruitSelection
		if err := c.ShouldBind(&selection); err != nil {
			c.String(http.StatusBadRequest, "Binding error: %v", err)
			return
		}
		c.String(http.StatusOK, "Selected fruits: %v", selection.Fruits)
	})

	//example 2
	r.GET("/render", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form2.html", nil)
	})

	// Handle form submission
	r.POST("/submit2", func(c *gin.Context) {
		var selection FruitSelection
		if err := c.ShouldBind(&selection); err != nil {
			c.String(http.StatusBadRequest, "Binding error: %v", err)
			return
		}
		c.HTML(http.StatusOK, "success.html", gin.H{
			"fruits": selection.Fruits,
		})
	})
	//-------------------------------Bind html checkboxes-----------------------------------

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
	// r.Run(":3000") for a hard coded port

}
