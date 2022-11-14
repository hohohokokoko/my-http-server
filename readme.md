# Simple HTTP server
In this project, we’re going to build a multi-threaded web server by Golang. Our goals are listed as below:

1. Built a concurrency model using Coroutine Pool and non-blocking socket.  
1. Using Goroutine to implement IO multiplexing.  
1. Used regular expression to parse HTTP request packets, support parsing GET and POST requests.  
1. Use WebBench to test the concurrency of our web server and concurrent connection data exchange (stress test).  

## TODOs:
1. ✅Handle http request concurrently with Goroutine
1. ✅ResponseWriter Buffer
1. ✅Implement Route Matching (Router)
1. ✅Read and parse request without using net/http package


## References:
1. https://zhuanlan.zhihu.com/p/101995755 (Chinese)
1. https://pkg.go.dev/net/http