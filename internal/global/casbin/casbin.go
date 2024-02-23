package casbin

import (
	"github.com/casbin/casbin/v2"
)

type MyEnforce struct {
	*casbin.Enforcer
}

func GetEnforce() *MyEnforce {
	return Enforce
}
