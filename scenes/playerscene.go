package scenes

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	//go:embed resources/player.png
	playerPng     []byte
	playerSprite  *ebiten.Image
	player        Player
	playerSprites map[string]SpriteLocation
	playerXNum    = 8
)

type SpriteLocation struct {
	x int
	y int
}

type PlayerScene struct {
	count       int
	childScenes []Scene
}

type Player struct {
	xPos      int
	yPos      int
	direction string
}

func init() {
	// Get image here
	img, err := png.Decode(bytes.NewReader(playerPng))
	if err != nil {
		log.Fatal(err)
	}

	playerSprite = ebiten.NewImageFromImage(img)
	player = Player{
		xPos:      18 * tileSize,
		yPos:      15 * tileSize,
		direction: "down",
	}

	playerSprites = make(map[string]SpriteLocation)
	playerSprites["up"] = SpriteLocation{
		x: (9 % playerXNum) * tileSize,
		y: (9 / playerXNum) * tileSize,
	}
	playerSprites["down"] = SpriteLocation{

		x: (3 % playerXNum) * tileSize,
		y: (3 / playerXNum) * tileSize,
	}
	playerSprites["right"] = SpriteLocation{
		x: (0 % playerXNum) * tileSize,
		y: (0 / playerXNum) * tileSize,
	}
	playerSprites["left"] = SpriteLocation{

		x: (6 % playerXNum) * tileSize,
		y: (6 / playerXNum) * tileSize,
	}
}

func (s *PlayerScene) Draw(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(player.xPos), float64(player.yPos))

	x := playerSprites[player.direction].x
	y := playerSprites[player.direction].y

	r.DrawImage(playerSprite.SubImage(image.Rect(x, y, x+tileSize, y+tileSize)).(*ebiten.Image), op)
}

func (s *PlayerScene) Update(state *GameState) error {
	nextXPos := player.xPos
	nextYPos := player.yPos
	nextDirection := player.direction

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		nextYPos -= tileSize
		nextDirection = "up"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		nextYPos += tileSize
		nextDirection = "down"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		nextXPos -= tileSize
		nextDirection = "left"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		nextXPos += tileSize
		nextDirection = "right"
	}

	if CanTraverse(Tile{x: nextXPos, y: nextYPos}) {
		player.xPos = nextXPos
		player.yPos = nextYPos
		player.direction = nextDirection
	}

	return nil
}

func NewPlayerScene() *PlayerScene {
	return &PlayerScene{}
}

func (s *PlayerScene) GetChildScenes() []Scene {
	return s.childScenes
}
