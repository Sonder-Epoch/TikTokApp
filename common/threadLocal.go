package common

import (
	"github.com/timandy/routine"
)

var threadLocal = routine.NewInheritableThreadLocal()

func SetUser(userId int64) {
	threadLocal.Set(userId)
}

func GetUser() int64 {
	get := threadLocal.Get()
	switch get.(type) {
	case int64:
		return get.(int64)
	}
	return 0
}
