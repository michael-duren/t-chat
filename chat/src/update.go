package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowDimensions.Width = msg.Width
		m.WindowDimensions.Height = msg.Height
	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "h":
			m.page = "home"
		case "a":
			m.page = "about"
		}
	}
	return m, nil
}
