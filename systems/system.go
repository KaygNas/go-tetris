package systems

import "go-tetris/entities"

type System interface {
	Init(*entities.Game)
	Tick()
	Close()
}
