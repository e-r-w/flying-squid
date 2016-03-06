package main

import "github.com/e-r-w/flying-squid/lib"

func main() {

	var app flyingSquid.App
	app.
		Bootstrap().
		CreateServer().
		RunOnAddr(":8080")

}
