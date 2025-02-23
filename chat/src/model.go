package main

import (
	"t-chat/src/views"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type WindowDimensions struct {
	Width  int
	Height int
}

type model struct {
	page             string
	WindowDimensions WindowDimensions
}

func InitialModel() model {
	return model{
		page: "home",
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) View() string {
	s := "What should we buy at the market?\n\n"

	switch m.page {
	case "home":
		s += views.Home()
	case "about":
		s += views.About()
	}

	s += "\nPress q to quit.\n"

	return lipgloss.Place(
		m.WindowDimensions.Width,
		m.WindowDimensions.Height,
		lipgloss.Center,
		lipgloss.Center,
		s,
	)
}
