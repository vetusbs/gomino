package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/vetusbs/gomino/controller"
)

// Create a new instance of the logger. You can have any number of instances.

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	mux := controller.Register()

	logrus.Infof("listening in port %v \n", port)

	panic(http.ListenAndServe(":"+port, mux))
}
