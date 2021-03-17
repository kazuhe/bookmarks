package main

import "github.com/kazuhe/bookmarks/controllers"

var server = controllers.Server{}

func main() {

	server.Serve()

}
