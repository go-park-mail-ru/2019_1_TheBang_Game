package room

import "fmt"

var (
	leftBorder  uint = 0
	rightBorder uint = width - 1
	upBorder    uint = height - 1
	downBorder  uint = 0
)

type GameInst struct {
	Map          GameMap
	PlayersPos   map[string]Position
	PlayersScore map[string]uint
	GemsCount    uint // захардкодить число гемов
	Room         *Room
}

func NewGame(r *Room) GameInst {
	score := make(map[string]uint)
	pos := make(map[string]Position)

	// заглушка для синглплеера, сделать генерацию плееров и их позиции на карте
	score["test"] = 0
	pos["test"] = Position{}

	return GameInst{
		Map:          NewMap(),
		PlayersPos:   pos,
		PlayersScore: score,
		GemsCount:    1, // захардкодить число гемов
		Room:         r,
	}
}

func (g *GameInst) Snap() GameSnap {
	return GameSnap{
		Map: g.Map,
		// PlayersScore: g.PlayersScore,
		GemsCount: 1,
	}
}

func (g *GameInst) Aggregation(actions ...Action) {
	for _, action := range actions {
		g.AcceptAction(action)
		fmt.Println()
	}
}

func (g *GameInst) AcceptAction(action Action) {
	var (
		pos Position
		ok  bool
	)

	if pos, ok = g.PlayersPos[action.Player]; !ok {
		return
	}

	newpos := pos

	switch {
	case action.Move == left:
		if newpos.X > leftBorder {
			newpos.X--
		}

	case action.Move == right:
		if newpos.X < rightBorder {
			newpos.X++
		}

		// up и down инвертированы
	case action.Move == down:
		if newpos.Y < upBorder {
			newpos.Y++
		}

	case action.Move == up:
		if newpos.Y > downBorder {
			newpos.Y--
		}
	}

	// if g.Map[pos.X][pos.Y] == gem {
	// 	g.PlayersScore[action.Player]++
	// 	g.GemsCount--
	// }

	g.PlayersPos[action.Player] = newpos
	g.Map[pos.X][pos.Y] = groung
	g.Map[newpos.X][newpos.Y] = player
}
