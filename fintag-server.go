package main

import (
	"github.com/kunterbunt/fintag-server/controller"
	"github.com/kunterbunt/fintag-server/model"
	"fmt"
	"runtime"
	"path"
)

func main() {
	// Connect to mongodb.
	mongoDb := model.NewMongoDb("fintag")
	// Find own package directory.
	_, filename, _, ok := runtime.Caller(0)
  if !ok {
    panic("No caller information")
  }
	// Instatiate server and register REST API controllers.
	server := controller.NewServer(path.Dir(filename))
	server.RegisterController("/coordinates", controller.NewCoordinatesController(mongoDb))
	// We're done setting up, fire it up!
	fmt.Println("Fintag is ready to rumble!")
	server.StartListen()
}
