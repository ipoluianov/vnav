package main

import (
	"context"
	"fmt"
	"vnav/system"
)

// App struct
type App struct {
	ctx context.Context
	s   *system.System
}

// NewApp creates a new App application struct
func NewApp() *App {
	c := &App{}
	c.s = system.NewSystem()
	return c
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

func (a App) Run(cmd string) string {
	return "[" + cmd + "]"
}

func (a App) GetFilePanelContentAsJson(index int) string {
	return a.s.GetFilePanelContentAsJson(index)
}

func (c *App) SetCurrentItemIndex(panelIndex, index int) {
	c.s.SetCurrentItemIndex(panelIndex, index)
}

func (c *App) UpdateContent(panelIndex int) {
	c.s.UpdateContent(panelIndex)
}

func (c *App) MainAction(panelIndex int) {
	c.s.MainAction(panelIndex)
}

func (c *App) GoBack(panelIndex int) {
	c.s.GoBack(panelIndex)
}

func (c *App) CreateDirectory(name string, panelIndex int) string {
	err := c.s.CreateDirectory(name, panelIndex)
	//time.Sleep(1 * time.Second)
	//runtime.EventsEmit(c.ctx, "updateContent", panelIndex)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (c *App) Remove(panelIndex int) string {
	err := c.s.Remove(panelIndex)
	//time.Sleep(1 * time.Second)
	if err != nil {
		return err.Error()
	}
	return ""
}
