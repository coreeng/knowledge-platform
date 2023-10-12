package router

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type InternalServer struct {
	mux *http.ServeMux
}

// NewInternalServer creates a http server that exposes prometheus
// metrics on an endpoint. This server is not accessible via the
// internet by default.
func NewInternalServer() *InternalServer {
	mux := http.NewServeMux()
	s := &InternalServer{mux: mux}
	s.routes()
	return s
}

func (s *InternalServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *InternalServer) routes() {
	s.mux.Handle("/metrics", promhttp.Handler())
	s.mux.Handle("/status", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
}

type Server struct {
	mux *http.ServeMux
}

// NewServer creates a http server that exposes a canaryTest endpoint
func NewServer() *Server {
	mux := http.NewServeMux()
	s := Server{mux: mux}
	s.routes()
	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) routes() {
	s.mux.Handle("/canaryTest", http.HandlerFunc(canarytestHandler))
}

func canarytestHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("status") != "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("canary test is passing."))
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("canary test is failing."))
}
