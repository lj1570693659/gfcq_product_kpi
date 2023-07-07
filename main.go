package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/frame/g"
	_ "github.com/lj1570693659/gfcq_product_kpi/boot"
	_ "github.com/lj1570693659/gfcq_product_kpi/router"
)

// @title       `gf-demo`示例服务API
// @version     1.0
// @description `GoFrame`基础开发框架示例服务API接口文档。
// @schemes     http
func main() {
	g.Server().Run()
}
