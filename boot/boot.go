package boot

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	_ "github.com/lj1570693659/gfcq_product_kpi/packed"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

var (
	BaseServer        = grpcx.Client.MustNewGrpcClientConn("employee")
	DepertmentServer  = v1.NewDepartmentClient(BaseServer)
	JobServer         = v1.NewJobClient(BaseServer)
	JobLevelServer    = v1.NewJobLevelClient(BaseServer)
	EmployeeServer    = v1.NewEmployeeClient(BaseServer)
	EmployeeJobServer = v1.NewEmployeeJobClient(BaseServer)
)

// 用于应用初始化。
func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
}
