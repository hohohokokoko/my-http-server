package myhttp

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/url"
	"strconv"
	"strings"
)

// Server server struct
type Server struct {
	Addr   string
	Router *Router
}

// conn connection struct
type conn struct {
	server *Server
	rwc    net.Conn
}

// Request struct
type Request struct {
	Method        string
	URL           *url.URL
	Header        Header
	Body          io.Reader
	ContentLength int64
	Host          string
	RemoteAddr    string
}

// Header header struct
type Header map[string][]string

// ListenAndServe Establish TCP
func (srv *Server) ListenAndServe() error {
	addr := srv.Addr
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Printf("Listening and serving on %s\n", addr)
	return srv.Serve(ln)
}

// Serve start to serve HTTP
func (srv *Server) Serve(l net.Listener) error {
	for {
		rw, err := l.Accept()
		if err != nil {
			return err
		}
		c := &conn{
			server: srv,
			rwc:    rw,
		}
		go c.serve()
	}
}

// serve handle connection
func (c *conn) serve() {
	defer c.rwc.Close()
	req, err := c.parseRequest()
	if err != nil {
		log.Printf("parse requset failed, err:%v", err)
		return
	}
	// Implement Router and Call Handler Here
	resp := NewResponseWriter(req, c)
	router := c.server.Router
	handler := router.MatchHandler(req.Method, req.URL.String())
	handler(resp, req)
	return
}

// parseRequest read and parse request from incoming connection
func (c *conn) parseRequest() (*Request, error) {
	reader := bufio.NewReader(c.rwc)
	line, _, _ := reader.ReadLine()
	items := strings.SplitN(string(line), " ", 3)
	method := items[0]
	path := items[1]
	u, _ := url.Parse(path)
	header := make(Header)

	for len(line) > 0 {
		line, _, _ = reader.ReadLine()
		if len(line) == 0 {
			break
		}
		items = strings.SplitN(string(line), ": ", 2)
		key, value := items[0], items[1]
		header[key] = append(header[key], value)
	}

	contentLength := 0
	if val, ok := header["Content-Length"]; ok {
		contentLength, _ = strconv.Atoi(val[0])
	}
	host := ""
	if val, ok := header["Host"]; ok {
		host = val[0]
	}
	r := &Request{
		Method:        method,
		URL:           u,
		Header:        header,
		Body:          reader,
		ContentLength: int64(contentLength),
		Host:          host,
		RemoteAddr:    c.rwc.RemoteAddr().String(),
	}
	return r, nil
}
