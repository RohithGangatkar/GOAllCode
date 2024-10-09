1] AsciiJSON
Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII characters.

2] XML/JSON/YAML/ProtoBuf rendering

3] Using HTTP method
router.GET("/someGet", getting)
router.POST("/somePost", posting)
router.PUT("/somePut", putting)
router.DELETE("/someDelete", deleting)
router.PATCH("/somePatch", patching)
router.HEAD("/someHead", head)
router.OPTIONS("/someOptions", options)


4] Bind Uri param to struct
struct syntax (ID   string `uri:"id"`) form need to be there
"/:name/:id" 
c.ShouldBindUri(&person)

5]Bind query string or post data
struct syntax ( Name  string `form:"name"`) form need to be there
c.ShouldBind(&person) or c.BindQuery(&queryParams)
curl -X GET "localhost:8080/user?name=John&age=30&email=john@example.com"
GET http://localhost:8080/user?name=John&age=30&email=john@example.com
The form tag avoids issues related to case sensitivity. HTTP query parameters are case-sensitive, and using the form tag ensures that even if your struct fields use different casing, the correct query parameters will still be matched.

6] Load HTML templates and Bind html checkboxes
create one html page under templete folder and create struct a to map from data
a]r.LoadHTMLGlob("templates/*") - LoadHTMLGlob method allows Gin to parse and load HTML template files so they can be rendered in response to HTTP requests
b]Gin uses the Go html/template package to parse the HTML files. It reads all files matching the specified pattern and stores them in a template.Template object
c]c.HTML method in your handler functions to render a specific template.
For example, c.HTML(http.StatusOK, "form.html", nil) will render the form.html template.

2 Examples have been provide (one dispay output and other one renders output in html form)
"/" - endpoint will load form.html page and onec we select options, it will display respective output.
Form Submission:
The POST handler binds the form data to the FruitSelection struct and then renders the success.html template, passing the selected fruits as data.
