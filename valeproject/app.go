package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
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

func (a *App) AnalyzeText(text string) (string, error) {
	fmt.Println("Received text for analysis:", text)
	cmd := exec.Command("vale", "--output", "line", "-")

	//Set up command input and output pipes
	cmd.Stdin = strings.NewReader(text)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	//Run the command
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("Error running Vale: %v\nStderr: %s", err, stderr.String())
	}

	//return out.String(), nil
	return "Analysis result", nil
}

//Search application to be developed in Go
//this will rely on the knowledge base
func (a *App) Search(query string) string {
	return fmt.Sprintf("We provide an answer to the question: %s?", query)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
