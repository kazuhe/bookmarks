package controllers

func (server *Server) Router() {

	// Users
	server.mux.HandleFunc("/v1/users/", UsersHandler)

}
