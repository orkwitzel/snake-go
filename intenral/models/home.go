package models

import (
	"time"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

// UI Palette and Styles
var (
	// Title block in vibrant neon green
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FFA3")).
			MarginBottom(1)

	// Tagline / Subtitle
	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true).
			MarginBottom(2)

	// Active button: bright text on Charm purple background
	activeButtonStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#7D56F4")).
				Padding(0, 3).
				Margin(0, 1)

	// Inactive button: muted text on dark tile background
	inactiveButtonStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#6C7086")).
				Background(lipgloss.Color("#1E1E2E")).
				Padding(0, 3).
				Margin(0, 1)

	// Footer helper text
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#585B70")).
			MarginTop(2)

	// Outer card container
	cardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#45475A")).
			Padding(1, 3).
			Align(lipgloss.Center)
)

const titleAscii = `
  ███████╗███╗   ██╗ █████╗ ██╗  ██╗███████╗
  ██╔════╝████╗  ██║██╔══██╗██║ ██╔╝██╔════╝
  ███████╗██╔██╗ ██║███████║█████╔╝ █████╗  
  ╚════██║██║╚██╗██║██╔══██║██╔═██╗ ██╔══╝  
  ███████║██║ ╚████║██║  ██║██║  ██╗███████╗
  ╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝`

type homeSelection string

const (
	homeSelectionPlay  homeSelection = "play"
	homeSelectionExist homeSelection = "exit"
)

type homeModel struct {
	selection homeSelection
	gameModel tea.Model
}

func NewHomeModel() homeModel {
	return homeModel{
		selection: homeSelectionPlay,
		gameModel: NewGameModel(),
	}
}

func (m homeModel) Init() tea.Cmd {
	return nil
}

func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "down", "left", "right", "tab":
			if m.selection == homeSelectionPlay {
				m.selection = homeSelectionExist
			} else {
				m.selection = homeSelectionPlay
			}

		case "enter":
			if m.selection == homeSelectionPlay {
				return m.gameModel, tickCmd(time.Second / 4)
			}
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m homeModel) View() tea.View {
	// 1. Render the title & subtitle
	title := titleStyle.Render(titleAscii)
	subtitle := subtitleStyle.Render("— CLASSIC TERMINAL EDITION —")

	// 2. Render buttons with active highlight state
	var playBtn, exitBtn string

	if m.selection == homeSelectionPlay {
		playBtn = activeButtonStyle.Render("▶ PLAY")
		exitBtn = inactiveButtonStyle.Render("  EXIT")
	} else {
		playBtn = inactiveButtonStyle.Render("  PLAY")
		exitBtn = activeButtonStyle.Render("▶ EXIT")
	}

	buttonRow := lipgloss.JoinHorizontal(lipgloss.Center, playBtn, exitBtn)

	// 3. Render helper controls
	controls := helpStyle.Render("arrows / tab: switch  •  enter: select  •  q: quit")

	// 4. Assemble vertically inside the rounded border frame
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		subtitle,
		buttonRow,
		controls,
	)

	return tea.NewView(cardStyle.Render(content))
}
