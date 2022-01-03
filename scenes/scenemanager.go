package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	transitionFrom = ebiten.NewImage(384, 384)
	transitionTo   = ebiten.NewImage(384, 384)
)

type Scene interface {
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
	GetChildScenes() []Scene
	// AddChildScene(s *Scene)
}

const transitionMaxCount = 20

type GameState struct {
	SceneManager *SceneManager
	// Input        *Input
}

type SceneManager struct {
	current         Scene
	next            Scene
	transitionCount int
}

func (s *SceneManager) Update() error {
	if s.transitionCount == 0 {
		mainErr := s.current.Update(&GameState{
			SceneManager: s,
			// Input:        input,
		})

		return mainErr
	}

	s.transitionCount--
	if s.transitionCount > 0 {
		return nil
	}

	s.current = s.next
	s.next = nil
	return nil
}

func (s *SceneManager) Draw(r *ebiten.Image) {

	if s.transitionCount == 0 {
		s.current.Draw(r)
		for _, scene := range s.current.GetChildScenes() {
			scene.Draw(r)
		}
		return
	}

	transitionFrom.Clear()
	s.current.Draw(transitionFrom)

	transitionTo.Clear()
	s.next.Draw(transitionTo)

	r.DrawImage(transitionFrom, nil)

	alpha := 1 - float64(s.transitionCount)/float64(transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	r.DrawImage(transitionTo, op)
}

func (s *SceneManager) GoTo(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = transitionMaxCount
	}
}
