package app

import (
	"context"
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

func (a *App) Cmd(command string) {
	command = strings.Join(strings.Split(command, "/")[:len(strings.Split(command, "/"))-1], "/")
	cmd := exec.Command("open", command)
	_, err := cmd.Output()
	if err != nil {
		panic(err)
	}
}
