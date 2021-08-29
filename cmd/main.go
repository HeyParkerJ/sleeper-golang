package main

import (
	"fmt"
	"github.com/MasEo9/sleeper-golang/Internal/http"
)

type App struct{}

func (app *App) Run() error {
	handler.GetRoster()
	handler.GetLeague()
	handler.GetUsers()

	return nil
}

func main() {
	fmt.Println("Running sleeper datafetch")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error running sleeper datafetch")
		fmt.Println(err)
	}
}
