package main

import "github.com/kazuhe/bookmarks/api"

var server = api.Server{}

func main() {

	server.Serve()

}
