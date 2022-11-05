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
| usage                                                                                                                           | description                                                                                       |
|---------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| [default](https://github.com/cloudwego/hertz-examples/blob/main/middleware/requestid/default/main.go)                           | This is using requestid by default                                                                |
| [custom key](https://github.com/cloudwego/hertz-examples/blob/main/middleware/requestid/custom_key/main.go)                     | How to use requestid for custom key                                                               |
| [custom generator](https://github.com/cloudwego/hertz-examples/blob/main/middleware/requestid/custom_generator/main.go)         | How to use requestid for custom generator                                                         |
| [custom handler](https://github.com/cloudwego/hertz-examples/blob/main/middleware/requestid/custom_handler/main.go)             | How to use requestid for custom handler                                                           |
| [get requestid](https://github.com/cloudwego/hertz-examples/blob/main/middleware/requestid/get_requestid/main.go)               | How to get requestid                                                                              |
| [log with hertzlogrus](https://github.com/cloudwego/hertz-examples/blob/main/middleware/requestid/log_with_hertzlogrus/main.go) | How to log requestid with [hertzlogrus](https://github.com/hertz-contrib/logger/tree/main/logrus) |

## License
This project is under the Apache License 2.0. See the LICENSE file for the full license text.