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
func (ePool *Pool) AddToPool(serv interface{}) {
	ePool.Servants = append(ePool.Servants, serv)
}

// GetFromPool is a function to get a servant from the pool.
//
// Parameters:
//  none
//
// Returns:
//  the servant.
//
func (ePool *Pool) GetFromPool() interface{} {
	if len(ePool.Servants) <= 0 {
		utils.PrintError(errors.New("empty pool"), "unable to get object from empty pool.")
		return nil
	}
	servHolder := ePool.Servants[ePool.CurrentIdx]
	ePool.CurrentIdx = (ePool.CurrentIdx + 1) % len(ePool.Servants)
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

// InitPool initialize a pool.
func InitPool(servs []interface{}) *Pool {
	echoP := Pool{Servants: servs, CurrentIdx: 0}
	return &echoP
}
