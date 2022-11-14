package myhttp

import (
	"bufio"
	"fmt"
	"time"
)

// Response struct
type Response struct {
	req        *Request
	w          *bufio.Writer
	statusCode int
	conn       *conn
	header     Header
}

// NewResponseWriter new a response writer
func NewResponseWriter(req *Request, conn *conn) *Response {
	res := Response{
		req:        req,
		w:          bufio.NewWriter(conn.rwc),
		statusCode: 200,
		conn:       conn,
		header:     Header{},
	}
	return &res
}

// statusMap HTTP status code
var statusMap = map[int]string{
	200: "OK",
	400: "Bad Request",
	404: "Not Found",
}

// Write write data to writer
func (r *Response) Write(p []byte) (n int, err error) {
	return r.w.Write(p)
}

// finish flush buffer
func (r *Response) finish() (err error) {
	return r.w.Flush()
}

// send data to client
func (r *Response) send(code int, content string) {
	fmt.Fprintf(r, "HTTP/1.1 %d %s\n", code, statusMap[code])
	r.header["Content-Type"] = []string{"text/plain; charset=utf-8"}
	r.header["Date"] = []string{time.Now().Format(time.RFC1123)}
	r.header["Content-Length"] = []string{fmt.Sprintf("%d", len(content)+1)}
	for key, values := range r.header {
		for _, value := range values {
			fmt.Fprintf(r, "%s: %s\n", key, value)
		}
	}
	fmt.Fprintf(r, "\n")
	fmt.Fprintf(r, "%s\n", content)
}

// RawText Response with raw text
func (r *Response) RawText(code int, text string) {
	r.send(code, text)
	r.finish()
}
