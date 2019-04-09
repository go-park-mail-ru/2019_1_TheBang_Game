package game

var (
	leftBorder  uint = 0
	rightBorder uint = width - 1
	upBorder    uint = height - 1
	downBorder  uint = 0
)

func NewGame() GameInst {
	return GameInst{
		Map: newMap(),
	}
}

func (g *GameInst) Snap() GameSnap {
	return GameSnap{
		Map:          g.Map,
		PlayersScore: g.PlayersScore,
		GemsCount:    1,
	}
}

func (g *GameInst) Aggregation(actions []Action) {
	for _, action := range actions {
		g.AcceptAction(&action)
	}
}

func (g *GameInst) AcceptAction(action *Action) {
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

	case action.Move == up:
		if newpos.X < upBorder {
			newpos.Y++
		}

	case action.Move == down:
		if newpos.X > downBorder {
			newpos.Y--
		}
	}

	if g.Map[pos.X][pos.Y] == gem {
		g.PlayersScore[action.Player]++
		g.GemsCount--
	}

	g.PlayersPos[action.Player] = newpos
	g.Map[pos.X][pos.Y] = groung
	g.Map[newpos.X][newpos.Y] = player
}

func (g *GameInst) Run() {
}
