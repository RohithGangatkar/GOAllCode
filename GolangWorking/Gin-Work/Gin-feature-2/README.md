Module all about basic auth, upload files and bind body into different structs


1] Using BasicAuth middleware
gin.Accounts is a shortcut for map[string]string
("/admin", gin.BasicAuth(gin.Accounts{"foo":    "bar"}))
// get user, it was set by the BasicAuth middleware
user := c.MustGet(gin.AuthUserKey).(string)

2] Upload files


3] bind body into different structs
PROBLEM: The normal methods for binding request body consumes c.Request.Body and they cannot be called multiple times.
SOLUTION : c.ShouldBindBodyWith stores body into the context before binding. This has a slight impact to performance, so you should not use this method if you are enough to call binding at once.

NOTE: This feature is only needed for some formats – JSON, XML, MsgPack, ProtoBuf. 
For other formats, Query, Form, FormPost, FormMultipart, can be called by c.ShouldBind() multiple times without any damage to performance