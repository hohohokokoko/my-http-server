package myhttp

// ListenAndServe start to listen and server http on addr with router
func ListenAndServe(addr string, router *Router) error {
	server := &Server{Addr: addr, Router: router}
	return server.ListenAndServe()
}
