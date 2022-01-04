package scenes

// Used to map sprite tiles in sheet
type SpriteLocation struct {
	x int
	y int
}

type Character struct {
	sprites               map[string]SpriteLocation
	xPos                  int
	yPos                  int
	direction             string
	moveTurnTimer         int
	currentMoveTurnTime   int
	actionTurnTimer       int
	currentActionTurnTime int
}

func NewCharacter(xPos int, yPos int, dir string, turnTime int, actionTime int) *Character {
	c := new(Character)
	c.xPos = xPos
	c.yPos = yPos
	c.direction = dir
	c.sprites = make(map[string]SpriteLocation)
	c.moveTurnTimer = turnTime
	c.actionTurnTimer = actionTime

	return c
}

func (c *Character) MapSprites(dir string, spriteIndex int, rowCount int) {
	c.sprites[dir] = SpriteLocation{
		x: (spriteIndex % rowCount) * tileSize,
		y: (spriteIndex / rowCount) * tileSize,
	}
}
