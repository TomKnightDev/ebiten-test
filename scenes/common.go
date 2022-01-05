package scenes

import (
	"github.com/solarlune/resolv"
)

// Used to map sprite tiles in sheet
type SpriteLocation struct {
	x int
	y int
}

type Entity struct {
	sprites               map[string]SpriteLocation
	xPos                  int
	yPos                  int
	direction             string
	moveTurnTimer         int
	currentMoveTurnTime   int
	actionTurnTimer       int
	currentActionTurnTime int
	entityObj             *resolv.Object
}

func NewEntity(s *SceneManager, xPos int, yPos int, dir string, turnTime int, actionTime int, tag string) *Entity {
	c := new(Entity)
	c.xPos = xPos
	c.yPos = yPos
	c.direction = dir
	c.sprites = make(map[string]SpriteLocation)
	c.moveTurnTimer = turnTime
	c.actionTurnTimer = actionTime
	c.entityObj = resolv.NewObject(float64(xPos), float64(yPos), 8, 8, tag)
	s.Space.Add(c.entityObj)

	return c
}

func (c *Entity) MapSprites(dir string, spriteIndex int, rowCount int) {
	c.sprites[dir] = SpriteLocation{
		x: (spriteIndex % rowCount) * tileSize,
		y: (spriteIndex / rowCount) * tileSize,
	}
}
