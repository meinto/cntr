package server

import (
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db}
}

func (s *Server) Start() {
	go func() {
		const defaultPort = "5564"

		port := os.Getenv("PORT")
		if port == "" {
			port = defaultPort
		}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("fdsa"))
		})

		log.Printf("connect to http://localhost:%s/", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()
}
