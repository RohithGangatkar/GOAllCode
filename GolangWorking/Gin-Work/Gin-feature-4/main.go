package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func runService1() {
	r1 := gin.Default()
	r1.GET("/endpoint1", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Service 1"})
	})
	if err := r1.Run(":8081"); err != nil {
		log.Fatal("Service 1 failed to start: ", err)
	}
}

func runService2() {
	r2 := gin.Default()
	r2.GET("/endpoint2", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Service 2"})
	})
	if err := r2.Run(":8082"); err != nil {
		log.Fatal("Service 2 failed to start: ", err)
	}
}

func main() {

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//--------------------------HTTP2 server push--------------------------------------------
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// Push CSS files
			cssFiles := []string{
				"/static/css/styles1.css",
				"/static/css/styles2.css",
			}
			for _, css := range cssFiles {
				if err := pusher.Push(css, nil); err != nil {
					c.String(http.StatusInternalServerError, "Failed to push CSS: %v", err)
					return
				}
			}

			// Push JavaScript files
			jsFiles := []string{
				"/static/js/script1.js",
				"/static/js/script2.js",
			}
			for _, js := range jsFiles {
				if err := pusher.Push(js, nil); err != nil {
					c.String(http.StatusInternalServerError, "Failed to push JavaScript: %v", err)
					return
				}
			}
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "HTTP/2 Server Push Example",
		})
	})

	// Load HTML files
	// r.LoadHTMLGlob("templates/*")

	// Run the server with HTTP/2 enabled
	// server := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: r,
	// }

	// // Enable HTTP/2 by using the net/http package
	// if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
	// 	panic(err)
	// }
	//-----------------------------End HTTP2 server push--------------------------------------

	//-----------------------------JSONP (JSON with Padding)----------------------------------
	//Add callback to response body if the query parameter callback exists.
	//where x is function name ofr call back function
	r.GET("/JSONP?callback=x", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		//callback is x
		// Will output  :   x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})
	r.GET("/jsonp", func(c *gin.Context) {
		callback := c.DefaultQuery("callback", "")
		data := gin.H{
			"message": "Hello, JSONP!",
		}

		if callback != "" {
			c.JSONP(http.StatusOK, data)
		} else {
			c.JSON(http.StatusOK, data)
		}
	})
	//-------------------------END of JSONP (JSON with Padding)----------------------------

	//---------------------Map as querystring or postform parameters----------------------
	r.POST("/postQuery", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	//-------------------End of Map as querystring or postform parameters-------------------

	//---------------------------------wildcard parameters----------------------------------
	r.GET("/user/:name/*action", func(c *gin.Context) {
		// Retrieve the path parameters
		name := c.Param("name")
		action := c.Param("action")

		// Respond with a formatted string
		c.String(http.StatusOK, "%s is doing %s", name, action)
	})
	//------------------------------End of wildcard parameters------------------------------

	//--------------------------------------PureJSON----------------------------------------
	// Serves unicode entities
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	//-----------------------------------End PureJSON---------------------------------------

	//-----------------------------------Query and post form--------------------------------
	r.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
	})
	//------------------------------End of Query and post form-----------------------------

	//--------------------------------------Redirect---------------------------------------
	// Define a route that redirects
	r.GET("/redirect", func(c *gin.Context) {
		// Redirect to another URL
		c.Redirect(302, "https://gin-gonic.com/docs/examples/redirects/")
	})
	//-----------------------------------End of Redirect------------------------------------

	//--------------------------------Run multiple service----------------------------------
	go runService1()
	go runService2()
	r.Run(":8080")
	// Keep the main goroutine running
	select {}
	//------------------------------End of Run multiple service-----------------------------

}
