package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gobuffalo/packr/v2"
	"github.com/jinzhu/gorm"
	"github.com/meinto/cntr/counter"
	"github.com/rs/cors"
)

type Server struct {
	db      *gorm.DB
	counter *counter.Counter
}

func NewServer(db *gorm.DB, c *counter.Counter) *Server {
	return &Server{db, c}
}

func (s *Server) Start() {
	go func() {
		box := packr.New("App", "../app/build")
		const defaultPort = "5564"

		port := os.Getenv("PORT")
		if port == "" {
			port = defaultPort
		}

		mux := http.NewServeMux()
		mux.Handle("/", http.FileServer(box))
		mux.HandleFunc("/getKeys", func(w http.ResponseWriter, r *http.Request) {
			startYear, _ := strconv.ParseInt(r.URL.Query().Get("startYear"), 10, 64)
			startMonth, _ := strconv.ParseInt(r.URL.Query().Get("startMonth"), 10, 64)
			startDay, _ := strconv.ParseInt(r.URL.Query().Get("startDay"), 10, 64)
			startDate := time.Date(
				int(startYear),
				time.Month(startMonth),
				int(startDay),
				0, 0, 0, 0,
				time.Now().Location(),
			)

			// tmp
			startDate = time.Now().Add(-1 * 240 * time.Hour)

			type Tuple struct {
				Keys   int    `json:"keys,omitempty"`
				Clicks int    `json:"clicks,omitempty"`
				Date   string `json:"date,omitempty"`
			}
			var response []Tuple
			date := startDate
			for i := 0; i < 10; i++ {
				duration := 24 * time.Hour
				date = date.Add(duration)
				keys := s.counter.GetKeys(date.Year(), date.YearDay())
				clicks := s.counter.GetClicks(date.Year(), date.YearDay())
				response = append(response, Tuple{keys, clicks, date.Format("02.01.2006")})
			}
			json.NewEncoder(w).Encode(response)
		})

		log.Printf("connect to http://localhost:%s/", port)

		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5564"},
			AllowCredentials: true,
			// Enable Debugging for testing, consider disabling in production
			Debug: false,
		})
		handler := c.Handler(mux)
		log.Fatal(http.ListenAndServe(":"+port, handler))
	}()
}
