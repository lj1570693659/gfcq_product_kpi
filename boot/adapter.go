package boot

import (
	"context"
	casbin "github.com/dobyte/gf-casbin"
	"github.com/gogf/gf/v2/frame/g"
	"log"
)

var Enforcer *casbin.Enforcer
var err error

func init() {
	tableName, _ := g.Config("config.toml").Get(context.Background(), "user.casbin.table")
	linkUrl, _ := g.Config("config.toml").Get(context.Background(), "user.casbin.link")
	Enforcer, err = casbin.NewEnforcer(&casbin.Options{
		Model:    "./document/auth/model.conf",
		Debug:    false,
		Enable:   true,
		AutoLoad: true,
		Table:    tableName.String(),
		Link:     linkUrl.String(),
	})

	if err != nil {
		log.Fatalf("Casbin init failure:%s \n", err.Error())
	}

	Enforcer.AddFunction("my_func", KeyMatchFunc)
}

// KeyMatch 我们当然也可以定义自己的函数。先定义一个函数，返回 bool：
func KeyMatch(requestAct, ruleAct string) bool {
	g.Log("auth").Info(context.Background(), requestAct, ruleAct)
	if requestAct == ruleAct {
		return true
	}

	return false
}

// KeyMatchFunc 然后将这个函数用interface{}类型包装一层：
func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(KeyMatch(name1, name2)), nil
}
