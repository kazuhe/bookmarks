package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// .envから値をロード
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

type Server struct {
	mux *http.ServeMux
}

func (server *Server) Serve() {
	server.mux = http.NewServeMux()

	// starting up the server
	httpServer := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: server.mux,
	}

	server.Router()

	log.Printf("Listening to port :" + os.Getenv("PORT"))
	log.Fatal(httpServer.ListenAndServe())
}
