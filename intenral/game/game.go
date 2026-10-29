package game

import (
	"math/rand/v2"
	"reflect"
	"slices"
)

func NewGame() Game {
	width := 10
	height := 10

	return Game{
		direction:    DirectionRight,
		turn:         0,
		width:        width,
		height:       height,
		gameOver:     false,
		snake:        newSnake(width, height),
		foodPosition: Position{X: 0, Y: 0},
	}
}

type Game struct {
	direction    Direction
	turn         int
	snake        Snake
	foodPosition Position
	width        int
	height       int
	gameOver     bool
}

func (g *Game) Tick(direction Direction) {
	g.moveSnake(direction)
	head := g.snake[0]

	if reflect.DeepEqual(head, g.foodPosition) {
		g.growSnake()
		g.respawnFood()
	}

	if head.X > g.width || head.X < 0 || head.Y > g.height || head.Y < 0 {
		g.gameOver = true
	}

	g.turn++
}

func (g Game) GameOver() bool {
	return g.gameOver
}

func (g Game) Snake() Snake {
	return g.snake
}

func (g *Game) moveSnake(direction Direction) {
	newValue := g.snake[0]

	switch direction {
	case DirectionUp:
		newValue.Y++
	case DirectionDown:
		newValue.Y--
	case DirectionLeft:
		newValue.X--
	case DirectionRight:
		newValue.X++
	}

	for i := range g.snake {
		copy := g.snake[i]
		g.snake[i] = newValue
		newValue = copy
	}
}

func (g *Game) growSnake() {
	s := g.snake

	lastPosition := s[len(s)-1]
	s = append(s, lastPosition)

	g.snake = s
}

func (g Game) getEmptyCells() []Position {
	s := g.snake

	positions := []Position{}
	for y := range g.GetHeight() {
		for x := range g.GetWidth() {
			pos := Position{X: x, Y: y}
			if !slices.Contains(s, pos) {
				positions = append(positions, pos)
			}
		}
	}

	return positions
}

func (g *Game) respawnFood() {
	emeptyCells := g.getEmptyCells()
	randomIndex := rand.N(len(emeptyCells))
	g.foodPosition = emeptyCells[randomIndex]
}

func (g Game) GetWidth() int {
	return g.width
}

func (g Game) GetHeight() int {
	return g.height
}

type Snake []Position
