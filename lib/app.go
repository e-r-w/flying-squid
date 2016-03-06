package flyingSquid

import (
	"fmt"
	"net/http"
	"os"

	"github.com/facebookgo/inject"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// App ...
type App struct {
	ImageRepository    ec2imageRepository    `inject:"imageRepository"`
	InstanceRepository ec2instanceRepository `inject:"instanceRepository"`
	StackRepository    stackRepository       `inject:"stackRepository"`
	HTTPTransport      http.RoundTripper     `inject:""`
}

// CreateServer ...
func (app App) CreateServer() *martini.ClassicMartini {

	repo := app.ImageRepository
	instanceRepo := app.InstanceRepository
	stackRepo := app.StackRepository

	server := martini.Classic()
	server.Use(render.Renderer())

	server.Get("/favicon.ico", func(renderer render.Render) {
		renderer.Text(404, "not found")
	})

	server.Get("/images", func(renderer render.Render) {
		renderer.JSON(200, repo.Fetch(""))
	})

	server.Get("/images/:imageName", func(renderer render.Render, params martini.Params) {
		renderer.JSON(200, repo.Fetch(params["imageName"]))
	})

	server.Get("/instances", func(renderer render.Render) {
		renderer.JSON(200, instanceRepo.Fetch(""))
	})

	server.Get("/instances/:instanceName", func(renderer render.Render, params martini.Params) {
		renderer.JSON(200, instanceRepo.Fetch(params["instanceName"]))
	})

	server.Get("/stacks", func(renderer render.Render) {
		renderer.JSON(200, stackRepo.Fetch(""))
	})

	server.Get("/stacks/:stackName", func(renderer render.Render, params martini.Params) {
		renderer.JSON(200, stackRepo.Fetch(params["stackName"])[0])
	})

	server.Get("/stacks/:stackName/resources", func(renderer render.Render, params martini.Params) {
		renderer.JSON(200, stackRepo.Resources(params["stackName"]))
	})

	return server

}

// Bootstrap ...
func (app App) Bootstrap() App {

	//_, _, _ = getCredentials()

	var graph inject.Graph
	var imageRepo ec2imageRepository // dependencies must be pointers to null
	var instanceRepo ec2instanceRepository
	var stackRepo stackRepository

	//dependency injection map
	if err := graph.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: imageRepo, Name: "imageRepository"},
		&inject.Object{Value: instanceRepo, Name: "instanceRepository"},
		&inject.Object{Value: stackRepo, Name: "stackRepository"},
		&inject.Object{Value: http.DefaultTransport},
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := graph.Populate(); err != nil {
		fmt.Println("foo")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return app

}
