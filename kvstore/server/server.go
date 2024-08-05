package server

import (
	"fmt"
	"kvstore/api/routes"
	"net/http"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /event", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /event")
		w.WriteHeader(200)
		w.Write([]byte("Success"))
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: routes.Routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Something went wrong", err)
		return err
	}
	fmt.Printf("Listening on %s\n", s.addr)
	return nil
}
