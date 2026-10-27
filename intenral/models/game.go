package models

import (
	"fmt"
	"snake/intenral/game"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func NewGameModel() gameModel {
	return gameModel{
		currentGame:      game.NewGame(),
		currentDirection: game.DirectionRight,
		tickDuration:     time.Second / 4,
	}
}

type gameModel struct {
	currentGame      game.Game
	currentDirection game.Direction
	tickDuration     time.Duration
}

func (m gameModel) Init() tea.Cmd {
	return tickCmd(m.tickDuration)
}

func (m gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.currentDirection = game.DirectionUp
		case "down":
			m.currentDirection = game.DirectionDown
		case "left":
			m.currentDirection = game.DirectionLeft
		case "right":
			m.currentDirection = game.DirectionRight
		}
	case tickMsg:
		if m.currentGame.GameOver() {
			return m, nil
		}
		m.currentGame.Tick(m.currentDirection)
		return m, tickCmd(m.tickDuration)
	}

	return m, nil
}

func tickCmd(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type tickMsg time.Time

var (
	// Board Container with a rounded border
	boardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")). // Charm purple
			Padding(0)

	// Cell Styles
	// Even if emojis have fixed foregrounds, Lipgloss can enforce background color & cell width
	floorStyle = lipgloss.NewStyle().
			Width(2).
			Align(lipgloss.Center).
			Background(lipgloss.Color("#1E1E2E")) // Dark catppuccin-style grid background

	foodStyle  = floorStyle.Copy().SetString("🍎")
	headStyle  = floorStyle.Copy().SetString("🟢")
	bodyStyle  = floorStyle.Copy().SetString("🟩")
	deadStyle  = floorStyle.Copy().SetString("💀")
	emptyStyle = floorStyle.Copy().SetString("· ")

	// Score & Header Style
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#04B575")).
			MarginBottom(1)

	// Alert Banner
	gameOverBanner = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF5F87")).
			MarginTop(1)
)

func (m gameModel) View() tea.View {
	w := m.currentGame.GetWidth()
	h := m.currentGame.GetHeight()

	// 1. Render the grid rows
	var rows []string
	for y := h - 1; y >= 0; y-- {
		var row strings.Builder
		for x := 0; x < w; x++ {
			switch m.currentGame.CellAt(x, y) {
			case game.CellSnakeHead:
				if m.currentGame.GameOver() {
					row.WriteString(deadStyle.Render())
				} else {
					row.WriteString(headStyle.Render())
				}
			case game.CellSnakeBody:
				row.WriteString(bodyStyle.Render())
			case game.CellFood:
				row.WriteString(foodStyle.Render())
			default:
				row.WriteString(emptyStyle.Render())
			}
		}
		rows = append(rows, row.String())
	}

	// 2. Wrap the board inside Lipgloss's border
	boardContent := strings.Join(rows, "\n")
	boardView := boardStyle.Render(boardContent)

	// 3. Assemble Header and Footer
	header := headerStyle.Render(fmt.Sprintf("🐍 SNAKE | Length: %d", len(m.currentGame.Snake())))

	footer := "Controls: Arrow keys | 'q' to quit"
	if m.currentGame.GameOver() {
		footer = gameOverBanner.Render("💥 GAME OVER! Press 'q' to quit.")
	}

	// 4. Combine everything vertically
	output := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		boardView,
		footer,
	)

	return tea.NewView(output)
}
