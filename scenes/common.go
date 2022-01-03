package scenes

// Used to map sprite tiles in sheet
type SpriteLocation struct {
	x int
	y int
}

type Character struct {
	sprites         map[string]SpriteLocation
	xPos            int
	yPos            int
	direction       string
	turnTimer       int
	currentTurnTime int
}

func NewCharacter(xPos int, yPos int, dir string, turnTime int) *Character {
	c := new(Character)
	c.xPos = xPos
	c.yPos = yPos
	c.direction = dir
	c.sprites = make(map[string]SpriteLocation)
	c.turnTimer = turnTime

	return c
}

func (c *Character) MapSprites(dir string, spriteIndex int) {
	c.sprites[dir] = SpriteLocation{
		x: (spriteIndex % 8) * tileSize,
		y: (spriteIndex / 8) * tileSize,
	}
}
