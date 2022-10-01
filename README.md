# RequestID (This is a community driven project)
 
Request ID middleware for Hertz framework, inspired by [requestid](https://github.com/gin-contrib/requestid).
This project would not have been possible without the support from the CloudWeGo community and previous work done by the gin community.

- Adds an identifier to the response using the `X-Request-ID` header. 
- Passes the `X-Request-ID` value back to the caller if it's sent in the request headers.


## Install
```shell
go get github.com/hertz-contrib/requestid
```

## Usage
### Example

```go
func main() {
    h := server.Default()

    h.Use(
        // provide your own request id generator here
        requestid.New(
            requestid.WithGenerator(func() string {
                return "cloudwego.io"
            }),
            // set custom header for request id
            requestid.WithCustomHeaderStrKey("your-customised-key"),
        ),
    )
    
    // Example ping request.
    h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
        c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
    })
    
    h.Spin()
}
```

### Getting the request ID

`requestid.Get(c)` is a helper function to retrieve request id from request headers. It also works with customised header as defined with `WithCustomHeaderStrKey`. 
Note that you may get empty string if it's not present in the request.

```go
// Example / request.
h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
    c.JSON(consts.StatusOK, utils.H{"ping": "pong", "request-id": requestid.Get(c)})
})
```

## License
This project is under the Apache License 2.0. See the LICENSE file for the full license text.