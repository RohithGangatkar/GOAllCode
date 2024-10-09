We have content on autotls, cookies, Controlling Log output coloring, Custom HTTP configuration,Custom log file, Custom Middleware


1]Encrypt HTTPS servers with autotls and we have custom autotls
Refer note and program. for more Gin official page.

2] cookies
cookie, err := c.Cookie("gin_cookie") - Get 
c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true) - Set

3] Render static file
r.Static("/static", "./static")
a]URL Path Prefix (/static): This is the URL path prefix that clients will use to access the static files.
b]File System Path (./static): This is the file system path where the static files are located.

Example:
Request URL: http://localhost:8080/static/css/styles.css
URL Path Prefix: /static
File System Path: ./static

Request: http://localhost:8080/public/css/styles.css
URL Path Prefix: /public
File System Path: ./assets
Result: Serves the file ./assets/css/styles.css

By accessing http://localhost:8080/public/css/styles.css, Gin will strip the /public prefix and map it to ./assets/css/styles.css. Similarly, other static files in the assets directory will be served under the /public URL path.

NOTE: you can choose any name for the URL path prefix when serving static files with Gin. The name you choose for the URL path prefix does not have to match the name of the directory on your file system.

4]Controlling Log output coloring
// Disable log's color
gin.DisableConsoleColor()

// Force log's color : Always colorize logs:
gin.ForceConsoleColor()

5] Custom HTTP configuration
 http.ListenAndServe(":8080", router)
 OR
 s := &http.Server{}
 s.ListenAndServe()

6]Custom log file
Need to have gin.New()
router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {}
OUTPUT: (for in program code)
::1 - [Fri, 07 Dec 2018 17:04:38 JST] "GET /ping HTTP/1.1 200 122.767µs "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.80 Safari/537.36" "


7] Custom Middleware
 Logger()

8]Goroutines inside a middleware
When starting new Goroutines inside a middleware or handler, you SHOULD NOT use the original context inside it, you have to use a read-only copy.

Original Context: The original gin.Context (c) is not used directly inside the Goroutine.
Context Copy: A shallow copy of the context is created and used inside the Goroutine to ensure thread safety.
Thread Safety: This approach prevents concurrency issues and race conditions that might occur if the original context is modified simultaneously by the main request handler and the Goroutine.

func myHandler(c *gin.Context) {
    // Create a read-only copy of the context
    ctxCopy := c.Copy()

    go func() {
        // Use the copied context in the new Goroutine
        handleRequestInGoroutine(ctxCopy)
    }()

    c.JSON(http.StatusOK, gin.H{"status": "processing"})
}

func handleRequestInGoroutine(c *gin.Context) {
    // Perform operations using the copied context
    // For example, you can read parameters, headers, etc.
    param := c.Para