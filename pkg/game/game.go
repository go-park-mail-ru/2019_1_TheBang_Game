package game

import (
	"BangGame/config"
	"fmt"
)

const (
	width  uint = 10
	height uint = 10

	left  = "left"
	right = "right"
	up    = "up"
	down  = "down"
)

type cell string

const (
	gem      cell = "gem"
	player   cell = "player"
	groung   cell = "ground"
	box      cell = "box"
	teleport cell = "teleport"
)

type Action struct {
	Time   string `json:"time"`
	Player string `json:"player"`
	Move   string `json:"move" ` // left | right | up | down
}

type Position struct {
	X uint
	Y uint
}

type GameInst struct {
	Map          GameMap
	PlayersPos   map[string]Position
	PlayersScore map[string]uint
	GemsCount    uint // захардкодить число гемов
}

type GameSnap struct {
	Map          GameMap         `json:"map"`
	PlayersScore map[string]uint `json:"players_score"`
	GemsCount    uint            `json:"gems_count"`
}

type GameMap [height][width]cell

func newMap() GameMap {
	config.Logger.Infow("NewMap",
		"msg", fmt.Sprint("NewMap was generated"))

	return GameMap{
		{player, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{groung, groung, groung, groung, groung, groung, groung, groung, groung, groung},
		{gem, groung, groung, groung, groung, groung, groung, groung, groung, groung}, // захадкожены гемы
	}
}
