package main

import (
	"context"

	"github.com/b1rr0/PackMe/execution"
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
func (a *App) Greet(name string) int {
	return 5 * 5
}

func (a *App) Pack(pathsToDir []string, outputFileName string) {
	execution.Pack(pathsToDir, outputFileName)
}

func (a *App) Unpack(pathToArchive string) {
	execution.Unpack(pathToArchive)
}
