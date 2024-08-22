package statics

import "github.com/charmbracelet/lipgloss"

var DashStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#DC143C")).
	Margin(0, 0, 0, 4)

var IdStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#DC143C"))

var ContentStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFD700"))

var InfoStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#32CD32"))

var ErrorStyle = lipgloss.NewStyle().
	Bold(true).
	Italic(true).
	Foreground(lipgloss.Color("#FF0000"))

var WarningStyle = lipgloss.NewStyle().
	Bold(true).
	Italic(true).
	Foreground(lipgloss.Color("#FA8072"))
