We have content on write log file, HTML rendering, HTTP2 server push

1] write log file
Disable Console Color, you don't need console color when writing the logs to file.
gin.DisableConsoleColor()
Use the following code if you need to write the logs to file and console at the same time.
gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

2]HTML rendering
Using LoadHTMLGlob() or LoadHTMLFiles()
router.LoadHTMLGlob("templates/*")
router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

3]HTTP2 server push

c.Writer.Pusher();
pusher.Push(css, nil);

HTTP/2 server push is a feature of the HTTP/2 protocol that allows a server to send resources to a client proactively, without the client having explicitly requested them. This can improve page load performance by reducing the number of round trips between the client and server.

How HTTP/2 Server Push Works
a]Initial Request: The client makes an initial request to the server for a specific resource, like an HTML page.
b]Push Promise: The server responds with the requested resource and includes a "push promise" frame, which indicates that it will also be sending additional resources that the client is likely to need (e.g., CSS, JavaScript, images).
c]Sending Resources: The server then pushes these additional resources to the client. The client can use these resources immediately or cache them for future use.


4] JSONP
JSONP (JSON with Padding) is a technique used to overcome the limitations of cross-domain requests in web browsers. Traditional AJAX calls are restricted by the same-origin policy, which means a web page can only make requests to the same domain from which it was loaded. JSONP allows you to request data from a server in a different domain by using <script> tags, which are not restricted by the same-origin policy.

r.GET("/JSONP?callback=x"
where x is function name ofr call back function
callback is x
Will output  :   x({\"foo\":\"bar\"})

Example output : the server will respond with

1]if you access the URL http://localhost:8080/jsonp?callback=handleResponse
If callback parameter is provided: handleResponse({"message":"Hello, JSONP!"})


2]http://localhost:8080/jsonp
If callback parameter is not provided:
{
  "message": "Hello, JSONP!"
}


5] Map as querystring or postform parameters
ids := c.QueryMap("ids")
names := c.PostFormMap("names")
POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1

Content-Type: application/x-www-form-urlencoded
names[first]=thinkerou&names[second]=tianou

6] Only bind query string
ShouldBindQuery function only binds the query params and not the post data.

7] Wildcard Parameter in path
Example Requests and Responses
Request: GET /user/john/walking
Response: john is doing /walking

Request: GET /user/jane/running/in/park
Response: jane is doing /running/in/park

Detailed Explanation
a]Static Segment (/user/):
The URL must start with /user/ for the route to match.
Named Parameter (:name):

b]Matches a single segment of the URL path.
The segment after /user/ is captured as name.
Wildcard Parameter (*action):

c]Matches any remaining part of the URL after the named parameter.
Captures everything following the name segment as action, including any subsequent slashes and segments.

8] PureJSON
Normally, JSON replaces special HTML characters with their unicode entities, e.g. < becomes \u003c. If you want to encode such characters literally, you can use PureJSON instead. This feature is unavailable in Go 1.6 and lower.

9]Query and post form
Query string parameters
id := c.Query("id")

page := c.DefaultQuery("page", "0")
name := c.PostForm("name")

POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded
name=manu&message=this_is_great

10] Redirects
This method is straightforward and allows you to send an HTTP redirect response to the client.
c.Redirect(302, "https://www.example.com") sends a 302 Found redirect to the client, which instructs the browser to navigate to https://www.example.com.

You can use different HTTP status codes for redirects, such as:
301 for "Moved Permanently"
302 for "Found" (temporary redirect)
307 for "Temporary Redirect"
308 for "Permanent Redirect"

11] Multiple Gin Instances
You can create multiple Gin instances and run them on different ports. This approach is more suitable if you want to run completely independent services.