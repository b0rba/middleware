package pooling

import (
	"log"
)

type Pool struct {
	Servants   []interface{}
	CurrentIdx int
}

//adds a servant to the pool.
func (cPool *Pool) AddToPool(serv interface{}) {
	cPool.Servants = append(cPool.Servants, serv)
}

// get a servant from the pool.
func (cPool *Pool) GetFromPool() interface{} {
	if len(cPool.Servants) <= 0 {
		log.Fatalln("Empty pool.")
		return nil
	}
	servHolder := cPool.Servants[cPool.CurrentIdx]
	cPool.CurrentIdx = (cPool.CurrentIdx + 1) % len(cPool.Servants)
	return servHolder
}

func EndPool(cPool *Pool) {
	for i := 0; i < len(cPool.Servants); i++ {
		cPool.Servants[i] = nil
	}
	cPool = nil
}

func InitPool(servs []interface{}) *Pool {
	calcP := Pool{Servants: servs, CurrentIdx: 0}
	return &calcP
}
