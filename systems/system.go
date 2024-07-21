package systems

import (
	"go-tetris/entities"
	"time"
)

type System interface {
	Init(*entities.Game)
	Tick(dt time.Duration)
	Close()
}
