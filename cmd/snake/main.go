package main

import (
	"fmt"
	"snake/intenral/models"

	tea "charm.land/bubbletea/v2"
)

func main() {
	p := tea.NewProgram(models.NewHomeModel())

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}
