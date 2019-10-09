package pooling

import (
	"errors"

	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// Pool is a struct to work with many Servants.
type Pool struct {
	Servants   []interface{}
	CurrentIdx int
}

// AddToPool add a servant to the pool.
func (ePool *Pool) AddToPool(serv interface{}) {
	ePool.Servants = append(ePool.Servants, serv)
}

// GetFromPool get a servant from the pool.
func (ePool *Pool) GetFromPool() interface{} {
	if len(ePool.Servants) <= 0 {
		utils.PrintError(errors.New("empty pool"), "unable to get object because the pool is empty.")
		return nil
	}
	servHolder := ePool.Servants[ePool.CurrentIdx]
	ePool.CurrentIdx = (ePool.CurrentIdx + 1) % len(ePool.Servants)
	return servHolder
}

// EndPool end a pool.
func EndPool(cPool *Pool) {
	for i := 0; i < len(cPool.Servants); i++ {
		cPool.Servants[i] = nil
	}
	cPool = nil
}

// InitPool initialize a pool.
func InitPool(servs []interface{}) *Pool {
	echoP := Pool{Servants: servs, CurrentIdx: 0}
	return &echoP
}
