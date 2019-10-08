package pooling

import (
	"errors"

	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// Pool is a struct to work with many Servants.
//
// Members:
//  Servants - list of Servant objects.
//  CurrentIdx - total number of used Servants.
//
type Pool struct {
	Servants   []interface{}
	CurrentIdx int
}

// AddToPool is a function to add a servant to the pool.
//
// Parameters:
//  serv - an object of the type of the pool servants.
//
// Returns:
//  none
//
func (cPool *Pool) AddToPool(serv interface{}) {
	cPool.Servants = append(cPool.Servants, serv)
}

// GetFromPool is a function to get a servant from the pool.
//
// Parameters:
//  none
//
// Returns:
//  the servant.
//
func (cPool *Pool) GetFromPool() interface{} {
	if len(cPool.Servants) <= 0 {
		utils.PrintError(errors.New("empty pool"), "unable to get object from empty pool.")
		return nil
	}
	servHolder := cPool.Servants[cPool.CurrentIdx]
	cPool.CurrentIdx = (cPool.CurrentIdx + 1) % len(cPool.Servants)
	return servHolder
}

// EndPool is a function to end a pool.
//
// Parameters:
//  cPool - the pool.
//
// Returns:
//  none
//
func EndPool(cPool *Pool) {
	for i := 0; i < len(cPool.Servants); i++ {
		cPool.Servants[i] = nil
	}
	cPool = nil
}

// InitPool is a function to initialize a pool.
//
// Parameters:
//  servs - the servants of the pool.
//
// Returns:
//  the pool.
//
func InitPool(servs []interface{}) *Pool {
	calcP := Pool{Servants: servs, CurrentIdx: 0}
	return &calcP
}
