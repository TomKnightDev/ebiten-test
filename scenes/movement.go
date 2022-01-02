package scenes

func CanTraverse(tile Tile) bool {
	for _, t := range Obstacles {
		if t.x == tile.x && t.y == tile.y {
			return false
		}
	}

	return true
}
