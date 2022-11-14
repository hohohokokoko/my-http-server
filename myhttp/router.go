package myhttp

// routerKey key in router map
type routerKey struct {
	Method string
	URL    string
}

// Router using map to store the business layer handler with key
type Router struct {
	RouterMap map[routerKey]func(resp *Response, req *Request)
}

// NewRouter new a router
func NewRouter() *Router {
	r := Router{RouterMap: map[routerKey]func(resp *Response, req *Request){}}
	return &r
}

// MatchHandler finding business layer handler with method and url
func (r *Router) MatchHandler(method string, url string) (res func(resp *Response, req *Request)) {
	rk := routerKey{
		Method: method,
		URL:    url,
	}
	if handler, ok := r.RouterMap[rk]; ok {
		return handler
	}
	return NotFoundHandler
}

// NotFoundHandler return 404 if not match
func NotFoundHandler(resp *Response, req *Request) {
	resp.RawText(404, req.Method+" "+req.URL.String()+" not Found")
}

// GET add Get method handler to the router map
func (r *Router) GET(url string, handler func(resp *Response, req *Request)) {
	key := routerKey{
		"GET",
		url,
	}
	if _, ok := r.RouterMap[key]; ok {
		panic("Duplicated handler key")
	}
	r.RouterMap[key] = handler
}

//POST add Post method handler to the router map
func (r *Router) POST(url string, handler func(resp *Response, req *Request)) {
	key := routerKey{
		"POST",
		url,
	}
	if _, ok := r.RouterMap[key]; ok {
		panic("Duplicated handler key")
	}
	r.RouterMap[key] = handler
}
