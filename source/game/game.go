package game

type Game struct {
	Board         [][]int // 0 = vide, 1 = joueur1, 2 = joueur2
	CurrentPlayer int
	Winner        int
}

func NewGame() *Game {
	board := make([][]int, 6)
	for i := range board {
		board[i] = make([]int, 7)
	}
	return &Game{
		Board:         board,
		CurrentPlayer: 1,
	}
}

func (g *Game) PlayMove(col int) bool {
	if col < 0 || col >= len(g.Board[0]) || g.Winner != 0 {
		return false
	}
	for row := len(g.Board) - 1; row >= 0; row-- {
		if g.Board[row][col] == 0 {
			g.Board[row][col] = g.CurrentPlayer
			if g.checkWin(row, col) {
				g.Winner = g.CurrentPlayer
			}
			g.switchPlayer()
			return true
		}
	}
	return false
}

func (g *Game) switchPlayer() {
	if g.CurrentPlayer == 1 {
		g.CurrentPlayer = 2
	} else {
		g.CurrentPlayer = 1
	}
}
func (g *Game) checkWin(r, c int) bool {
	player := g.Board[r][c]
	directions := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}
	for _, d := range directions {
		count := 1
		count += g.countDirection(r, c, d[0], d[1], player)
		count += g.countDirection(r, c, -d[0], -d[1], player)
		if count >= 4 {
			return true
		}
	}
	return false
}

func (g *Game) countDirection(r, c, dr, dc, player int) int {
	count := 0
	for {
		r += dr
		c += dc
		if r < 0 || r >= len(g.Board) || c < 0 || c >= len(g.Board[0]) {
			break
		}
		if g.Board[r][c] != player {
			break
		}
		count++
	}
	return count
}
g