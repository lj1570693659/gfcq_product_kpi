package product

import (
	"context"
	"fmt"
	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

var url = "http://127.0.0.1:8199/product/member"
var ctx context.Context

func Test_Import(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		route := fmt.Sprintf("%s/%s/%d", url, "import", 1)
		client := g.Client()
		client.SetPrefix(route)

		data := `{
			"product_member_list":"E:\资料\项目绩效相关\项目成员.xlsx"
}`
		client.Post(ctx, route, data)
		//client.GetContent(ctx, url, data)
	})
}
