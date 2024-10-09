package main

import (
	// "log"

	// "github.com/gin-gonic/autotls"
	"fmt"
	"log"
	"time"

	// "net/http"
	// "time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {

	// Disable log's color
	gin.DisableConsoleColor()

	// Force log's color : Always colorize logs:
	gin.ForceConsoleColor()

	r := gin.Default()
	r.Use(Logger())

	// // LetsEncrypt HTTPS servers. with autotls and we have custom autotls
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.String(200, "pong")
	// })

	// log.Fatal(autotls.Run(r, "example1.com", "example2.com"))

	//--------------------------------------Cookies-----------------------------------
	r.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")
		fmt.Printf("Cookie value: %s   Err: %s \n", cookie, err.Error())

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		cookieNew, _ := c.Cookie("gin_cookie")
		fmt.Printf("Cookie value: %s \n", cookie)
		fmt.Printf("Cookie value: %s \n", cookieNew)
	})
	//--------------------------------------End of Cookies-----------------------------------

	//-----------------Render static file-------------------------------
	r.Static("/static", "./static")

	// Serve the main HTML file
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	//-----------------End of Render static file------------------------

	//-----------------Custom HTTP configuration-------------------------------
	// http.ListenAndServe(":8080", router)

	//----------------------OR--------------------

	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
	//-----------------End Custom HTTP configuration-------------------------

	//-----------------Custom log file LoggerWithFormatter-------------------------------
	// router := gin.New()
	// // LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// // By default gin.DefaultWriter = os.Stdout
	// router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 			param.ClientIP,
	// 			param.TimeStamp.Format(time.RFC1123),
	// 			param.Method,
	// 			param.Path,
	// 			param.Request.Proto,
	// 			param.StatusCode,
	// 			param.Latency,
	// 			param.Request.UserAgent(),
	// 			param.ErrorMessage,
	// 	)
	// }))
	// router.Use(gin.Recovery())
	//-----------------End Custom log file LoggerWithFormatter---------------------------

	//--------------------------------Custom logger--------------------------------------
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
	})
	//--------------------------------End Custom logger-----------------------------------

	//------------------------Goroutines inside a middlewar - Context-------------------------
	//When starting new Goroutines inside a middleware or handler in Gin,
	//you should create a read-only copy of the original context to avoid concurrency issues and
	//ensure thread-safe operations.
	r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
	})
	//------------------------Goroutines inside a middlewar - Context-------------------------

	r.Run(":8080")

}
