package main

import (
	"mini_project/routes"
)

func main() {

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
