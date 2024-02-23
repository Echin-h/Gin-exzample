package casbin

import (
	"LearningGo/internal/global/log"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

const (
	casModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
)

var Enforce *MyEnforce

func Init() {
	var err error
	Enforce, err = createEnforcer()
	if err != nil {
		log.SugarLogger.Error(err)
		return
	}
	err1 := Enforce.LoadPolicy()
	if err1 != nil {
		log.SugarLogger.Error(err)
		return
	}
	Enforce.EnableAutoSave(true)
	Enforce.AddBasedPolicies()
}

func createEnforcer() (*MyEnforce, error) {
	m, err := model.NewModelFromString(casModel)
	if err != nil {
		return nil, err
	}
	adapter, err := gormadapter.NewAdapter("mysql", "user:password@tcp(127.0.0.1:3307)/mysql", true)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}
	return &MyEnforce{e}, nil
}
