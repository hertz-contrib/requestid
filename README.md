# RequestID
 
Request ID middleware for Hertz framework. 
- Adds an identifier to the response using the `X-Request-ID` header. 
- Passes the `X-Request-ID` value back to the caller if it's sent in the request headers.

## Install
```shell
go get https://github.com/hertz-contrib/requestid
```

## Usage


```go
func main() {
   h := server.Default()

   h.Use(
      // provide your own request id generator here 
      requestid.New(requestid.WithGenerator(func() string {
         return "cloudwego.io"
      })),
      // set custom header for request id
      requestid.WithCustomHeaderStrKey("your-customer-key"),
)

   // Example ping request.
   h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
      c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
   })

   h.Spin()
}
```

How to get the request identifier:

```go
// Example / request.
h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
    c.JSON(consts.StatusOK, utils.H{"ping": "pong", "request-id": requestid.Get(c)})
})
```