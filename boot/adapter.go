package boot

import (
	"context"
	"fmt"
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

	// TODO
	Enforcer.AddFunction("my_func", KeyMatchFunc)
}

// KeyMatch 我们当然也可以定义自己的函数。先定义一个函数，返回 bool：
func KeyMatch(requestKey, ruleKey, requestAct, ruleAct string) bool {

	fmt.Println("eeeeeeeeeeeeeeeee------request----------", requestKey)
	fmt.Println("eeeeeeeeeeeeeeeee-------rule---------", ruleKey)
	fmt.Println("eeeeeeeeeeeeeeeee-------requestAct---------", requestAct)
	fmt.Println("eeeeeeeeeeeeeeeee-------ruleAct---------", ruleAct)

	return true
}

// KeyMatchFunc 然后将这个函数用interface{}类型包装一层：
func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	name3 := args[2].(string)
	name4 := args[3].(string)

	return (bool)(KeyMatch(name1, name2, name3, name4)), nil
}
