package systems

import (
	"time"

	"github.com/KaygNas/go-tetris/entities"
)

type System interface {
	Init(*entities.Game)
	Tick(dt time.Duration) error
	Close()
}
