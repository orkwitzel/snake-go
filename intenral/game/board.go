package game

func newSnake(width int, height int) Snake {
	return []Position{
		{X: width / 2, Y: height / 2},
		{X: width/2 + 1, Y: height / 2},
	}
}

type CellType int

const (
	CellEmpty CellType = iota
	CellSnakeHead
	CellSnakeBody
	CellFood
)

func (g Game) CellAt(x, y int) CellType {
	if g.foodPosition.X == x && g.foodPosition.Y == y {
		return CellFood
	}
	for i, p := range g.snake {
		if p.X == x && p.Y == y {
			if i == 0 {
				return CellSnakeHead
			}
			return CellSnakeBody
		}
	}
	return CellEmpty
}
