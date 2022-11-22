package main

import (
	"fmt"
)

type App struct{}

func (app *App) Run() error {
	// handler.GetRoster()
	// handler.GetLeague()
	// handler.GetUsers()
	handler.GetMatchups()
	// handler.GetTransactions()
	// handler.GetTradedPicks()
	// handler.GetNflState()
	// handler.GetLeagueDraft()
	// handler.GetDraftPicks()
	// handler.GetTradedDraftPicks()
	// handler.GetPlayers()
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
