/*
@author: NanYan
*/
package repo

import (
	"time"
)

type Cache interface {
	Put(key string, value any, expire time.Duration) error
	Get(key string) (any, error)
}
