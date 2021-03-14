package api

import (
	"github.com/kazuhe/bookmarks/api/controllers"
)

func (server *Server) Router() {

	// Users
	server.mux.HandleFunc("/v1/users/", controllers.UsersHandler)

}
