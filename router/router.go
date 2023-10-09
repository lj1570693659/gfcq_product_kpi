package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/api"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
)

func init() {
	s := g.Server()
	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,
			service.Middleware.SaveUserLog,
		)

		// 系统管理
		group.Group("/system", func(sg *ghttp.RouterGroup) {
			// 账号管理
			sg.Group("/account", func(sga *ghttp.RouterGroup) {
				// 个人资料、修改密码、日志 TODO
				// 登录账号相关
				sga.ALL("/user", api.User)

				// 操作日志管理
				sga.Group("/log", func(lsg *ghttp.RouterGroup) {
					// 清单
					lsg.GET("/lists", api.User.GetUserLogLists)

				})
			})
			// 组织管理
			sg.Group("/organize", func(sgo *ghttp.RouterGroup) {
				// 部门、职级、员工 无需鉴权部分
				sgo.Group("/employee", func(sgoe *ghttp.RouterGroup) {
					sgoe.GET("/all", api.Employee.GetAll)
					sgoe.GET("/lists", api.Employee.GetList)
				})
				sgo.Group("/level", func(lgo *ghttp.RouterGroup) {
					lgo.GET("/all", api.JobLevel.GetAll)
				})
				sgo.Group("/department", func(dgo *ghttp.RouterGroup) {
					dgo.GET("/lists", api.Department.GetList)
				})

				// 员工信息相关
				sgo.Group("/employee", func(sgoe *ghttp.RouterGroup) {
					sgoe.Middleware(service.Middleware.LoggedIn)
					sgoe.GET("/isSyncEmployee", api.Employee.IsSyncEmployee)
					sgoe.POST("/create", api.Employee.Create)
					sgoe.PUT("/modify", api.Employee.Modify)
					sgoe.GET("/info", api.Employee.GetOne)
				})
				// 部门信息
				sgo.Group("/department", func(dgo *ghttp.RouterGroup) {
					dgo.Middleware(service.Middleware.Role, service.Middleware.BusinessRole)
					dgo.GET("/info", api.Department.GetOne)
					dgo.POST("/create", api.Department.Create)
					dgo.PUT("/modify", api.Department.Modify)
					dgo.DELETE("/delete", api.Department.Delete)
				})
				// 职级信息
				sgo.Group("/level", func(lgo *ghttp.RouterGroup) {
					lgo.Middleware(service.Middleware.Role, service.Middleware.BusinessRole)
					lgo.GET("/lists", api.JobLevel.GetList)
					lgo.GET("/info", api.JobLevel.GetOne)
					lgo.POST("/create", api.JobLevel.Create)
					lgo.PUT("/modify", api.JobLevel.Modify)
					lgo.DELETE("/delete", api.JobLevel.Delete)
				})
				// 岗位信息
				sgo.Group("/job", func(jgo *ghttp.RouterGroup) {
					jgo.GET("/lists", api.Job.GetList)
					jgo.GET("/all", api.Job.GetAll)
					jgo.POST("/create", api.Job.Create)
					jgo.PUT("/modify", api.Job.Modify)
					jgo.DELETE("/delete", api.Job.Delete)
				})
			})
		})

		// 配置管理
		group.Group("/config", func(cg *ghttp.RouterGroup) {
			cg.Middleware(service.Middleware.Role, service.Middleware.BusinessRole)
			// 项目配置
			cg.Group("/product", func(cgp *ghttp.RouterGroup) {
				// 评级、优先级、研发模式、项目类型、阶段、项目角色
				// 项目评级标准
				cgp.Group("/assess", func(agp *ghttp.RouterGroup) {
					agp.GET("/lists", api.LevelAssess.GetList)
					agp.GET("/info", api.LevelAssess.GetOne)
					agp.POST("/create", api.LevelAssess.Create)
					agp.PUT("/modify", api.LevelAssess.Modify)
					agp.DELETE("/delete", api.LevelAssess.Delete)
				})
				// 项目优先级确认
				cgp.Group("/confirm", func(cgp *ghttp.RouterGroup) {
					cgp.GET("/lists", api.LevelConfirm.GetList)
					cgp.POST("/create", api.LevelConfirm.Create)
					cgp.PUT("/modify", api.LevelConfirm.Modify)
					cgp.DELETE("/delete", api.LevelConfirm.Delete)
				})
				// 研发模式
				cgp.Group("/mode", func(mgp *ghttp.RouterGroup) {
					mgp.GET("/all", api.Mode.GetAll)
					mgp.POST("/create", api.Mode.Create)
					mgp.PUT("/modify", api.Mode.Modify)
					mgp.DELETE("/delete", api.Mode.Delete)
				})
				// 项目类型
				cgp.Group("/type", func(tgp *ghttp.RouterGroup) {
					tgp.GET("/all", api.Type.GetAll)
					tgp.POST("/create", api.Type.Create)
					tgp.PUT("/modify", api.Type.Modify)
					tgp.DELETE("/delete", api.Type.Delete)
				})
				// 项目阶段
				cgp.Group("/stage", func(stgp *ghttp.RouterGroup) {
					stgp.GET("/all/{tid}", api.Type.GetStageAll)
					stgp.POST("/create", api.Type.CreateModeStage)
					stgp.PUT("/modify", api.Type.ModifyModeStage)
					stgp.DELETE("/delete", api.Type.DeleteModeStage)
				})
				// 项目角色
				cgp.Group("/roles", func(rgp *ghttp.RouterGroup) {
					rgp.GET("/lists", api.ProductRoles.GetList)
					rgp.POST("/create", api.ProductRoles.Create)
					rgp.PUT("/modify", api.ProductRoles.Modify)
					rgp.DELETE("/delete", api.ProductRoles.Delete)
				})
			})
			// 绩效配置
			cg.Group("/inspirit", func(cgi *ghttp.RouterGroup) {
				// 激励预算、激励应发、管理指数。。。
				// 激励预算
				cgi.Group("/budget", func(bgp *ghttp.RouterGroup) {
					bgp.GET("/lists", api.BudgetAccess.GetList)
					bgp.GET("/all", api.BudgetAccess.GetAll)
					bgp.POST("/create", api.BudgetAccess.Create)
					bgp.PUT("/modify", api.BudgetAccess.Modify)
					bgp.DELETE("/delete", api.BudgetAccess.Delete)
				})
				// 激励应发
				cgi.Group("/radio", func(rgp *ghttp.RouterGroup) {
					rgp.GET("/all", api.BudgetRadio.GetAll)
					rgp.POST("/create", api.BudgetRadio.Create)
					rgp.PUT("/modify", api.BudgetRadio.Modify)
					rgp.DELETE("/delete", api.BudgetRadio.Delete)
				})
				// 管理指数
				cgi.Group("/manage", func(mgp *ghttp.RouterGroup) {
					mgp.GET("/all", api.CrewManageIndex.GetAll)
					mgp.POST("/create", api.CrewManageIndex.Create)
					mgp.PUT("/modify", api.CrewManageIndex.Modify)
					mgp.DELETE("/delete", api.CrewManageIndex.Delete)
				})
				// 工时指数
				cgi.Group("/hours", func(hgp *ghttp.RouterGroup) {
					hgp.GET("/all", api.CrewHoursIndex.GetAll)
					hgp.POST("/create", api.CrewHoursIndex.Create)
					hgp.PUT("/modify", api.CrewHoursIndex.Modify)
					hgp.DELETE("/delete", api.CrewHoursIndex.Delete)
				})
				// 责任指数
				cgi.Group("/duty", func(dgp *ghttp.RouterGroup) {
					dgp.GET("/all", api.CrewDutyIndex.GetAll)
					dgp.POST("/create", api.CrewDutyIndex.Create)
					dgp.PUT("/modify", api.CrewDutyIndex.Modify)
					dgp.DELETE("/delete", api.CrewDutyIndex.Delete)
				})
				// 问题解决
				cgi.Group("/solve", func(sgp *ghttp.RouterGroup) {
					sgp.GET("/all", api.CrewSolveRule.GetAll)
					sgp.POST("/create", api.CrewSolveRule.Create)
					sgp.PUT("/modify", api.CrewSolveRule.Modify)
					sgp.DELETE("/delete", api.CrewSolveRule.Delete)
				})
				// 加班贡献
				cgi.Group("/overtime", func(ogp *ghttp.RouterGroup) {
					ogp.GET("/all", api.CrewOvertimeRule.GetAll)
					ogp.POST("/create", api.CrewOvertimeRule.Create)
					ogp.PUT("/modify", api.CrewOvertimeRule.Modify)
					ogp.DELETE("/delete", api.CrewOvertimeRule.Delete)
				})
				// 绩效等级
				cgi.Group("/kpiRule", func(kgp *ghttp.RouterGroup) {
					kgp.GET("/all", api.CrewKpiRule.GetAll)
					kgp.POST("/create", api.CrewKpiRule.Create)
					kgp.PUT("/modify", api.CrewKpiRule.Modify)
					kgp.DELETE("/delete", api.CrewKpiRule.Delete)
				})
			})
		})

		// 项目绩效
		group.Group("/achieve", func(ag *ghttp.RouterGroup) {
			ag.Middleware(service.Middleware.Role, service.Middleware.BusinessRole)
			// 阶段绩效
			ag.Group("/product", func(agp *ghttp.RouterGroup) {
				agp.POST("/create", api.ProductStageKpi.Create)
				agp.PUT("/modify", api.ProductStageKpi.Modify)
				agp.GET("/lists", api.ProductStageKpi.GetList)
				agp.GET("/info", api.ProductStageKpi.GetOne)
				// 阶段团队成员绩效
				agp.Group("/member", func(amgp *ghttp.RouterGroup) {
					amgp.POST("/export", api.ProductMemberKpi.Export)
					amgp.POST("/import", api.ProductMemberKpi.Import)
					amgp.POST("/create", api.ProductMemberKpi.Create)
					amgp.PUT("/modify", api.ProductMemberKpi.Modify)
					amgp.GET("/lists", api.ProductMemberKpi.GetList)
					//amgp.GET("/info", api.ProductMemberKpi.GetOne)
					//agp.DELETE("/delete", api.LevelAssess.Delete)
				})
				// 阶段团队成员关键事件
				agp.Group("/crucial", func(cgp *ghttp.RouterGroup) {
					cgp.GET("/lists", api.ProductMemberKey.GetList)
					cgp.POST("/create", api.ProductMemberKey.Create)
					cgp.PUT("/modify", api.ProductMemberKey.Modify)
					cgp.DELETE("/delete", api.ProductMemberKey.Delete)
				})
				// 阶段团队成员激励计算
				agp.Group("/prize", func(apgp *ghttp.RouterGroup) {
					apgp.POST("/compute", api.ProductMemberPrize.Compute)
					apgp.POST("/export", api.ProductMemberPrize.Export)
					apgp.GET("/lists", api.ProductMemberPrize.GetList)
				})
			})

		})

		// 项目管理
		group.Group("/product", func(pg *ghttp.RouterGroup) {
			pg.Middleware(service.Middleware.Role, service.Middleware.BusinessRole)
			// 清单、详情
			pg.Middleware(service.Middleware.LoggedIn, service.Middleware.Role, service.Middleware.BusinessRole)
			pg.GET("/lists", api.Product.GetList)
			pg.GET("/all", api.Product.GetAll)
			pg.GET("/info", api.Product.GetOne)
			pg.GET("/detail", api.Product.GetDetail)
			pg.POST("/create", api.Product.Create)
			pg.DELETE("/delete", api.Product.Delete)
			pg.PUT("/modify", api.Product.Modify)
			// 项目组成员信息
			pg.Group("/member", func(pmg *ghttp.RouterGroup) {
				pmg.POST("/import", api.ProductMember.Import)
				pmg.GET("/lists", api.ProductMember.GetList)
				pmg.GET("/info", api.ProductMember.GetOne)
				pmg.POST("/create", api.ProductMember.Create)
				pmg.PUT("/modify", api.ProductMember.Modify)
				pmg.DELETE("/delete", api.ProductMember.Delete)
			})
			// 项目阶段信息
			pg.Group("/stage", func(psg *ghttp.RouterGroup) {
				psg.GET("/all/{proId}", api.ProductStage.GetList)
			})
		})

		// 统计数据
		group.Group("/statistics", func(sg *ghttp.RouterGroup) {
			// 总量统计
			sg.Group("/summation", func(sumg *ghttp.RouterGroup) {
				sumg.GET("/inspire", api.StatisticsSummation.GetInspire)
				sumg.GET("/stage", api.StatisticsSummation.GetStage)
			})
			// 项目阶段统计
			sg.Group("/product", func(sumg *ghttp.RouterGroup) {
				sumg.GET("/stage", api.StatisticsSummation.GetProductStage)
				sumg.GET("/score", api.StatisticsSummation.GetProductStageScore)
				sumg.GET("/top", api.StatisticsSummation.GetProductStageTop)
			})
			// 绩效等级
			sg.Group("/level", func(sumg *ghttp.RouterGroup) {
				sumg.GET("/index", api.StatisticsSummation.GetProductMemberLevel)
			})
		})
	})
}
