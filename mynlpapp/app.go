package main

import (
	"context"
	"fmt"

	"github.com/ddecoen/nlp-week9/mynlpapp/backend"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CheckTermEmbed(term string) string {
	answer, exists, err := backend.ValidTermEmbed(term)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	if exists {
		return fmt.Sprintf("For term '%s', found answer: %s", term, answer)
	}
	return fmt.Sprintf("No answer found for term '%s', please select from the valid term list.", term)
}
