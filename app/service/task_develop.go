package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	wechat "github.com/lj1570693659/gfcq_protoc/wechat/v1"
	"strings"
)

var TaskDevelop = taskDevelopService{}

type taskDevelopService struct{}

// AutoRemindNotEnd 今天未结束
func (s *taskDevelopService) AutoRemindNotEnd(ctx context.Context) {
	// dutyMsgData 责任人
	dutyMsgData := make(map[string]map[string][]string, 0)
	// joinMsgData 关联责任人
	joinMsgData := make(map[string]map[string][]string, 0)
	resData, _ := dao.ProductTaskDevelop.GetAll(ctx, &entity.ProductTaskDevelop{
		TaskStatus:  2,
		PalnEndTime: gtime.Now(),
	})
	if len(resData) > 0 {
		for _, v := range resData {
			if !g.IsEmpty(v.DutyWorkNumber) {
				if g.IsNil(dutyMsgData[v.Name]) {
					dutyMsgData[v.Name] = make(map[string][]string)
				}
				for _, uv := range strings.Split(v.DutyWorkNumber, "|") {
					dutyMsgData[v.Name][uv] = append(dutyMsgData[v.Name][uv], fmt.Sprintf("任务：%s,需要在%s前完成, 任务描述：%s", v.TaskName, v.PalnEndTime.Format("Y-m-d"), v.TaskDesc))
				}
			}
			if !g.IsEmpty(v.JoinWorkNumber) {
				if g.IsNil(joinMsgData[v.Name]) {
					joinMsgData[v.Name] = make(map[string][]string)
				}
				for _, jv := range strings.Split(v.JoinWorkNumber, "|") {
					joinMsgData[v.Name][jv] = append(joinMsgData[v.Name][jv], fmt.Sprintf("任务：%s,需要在%s前完成, 任务描述：%s", v.TaskName, v.PalnEndTime.Format("Y-m-d"), v.TaskDesc))
				}
			}
		}
	}
	if len(dutyMsgData) > 0 {
		for dk, dv := range dutyMsgData {
			if len(dv) > 0 {
				for duk, duv := range dv {
					boot.WechatCheckIn.SendMsg(ctx, &wechat.SendTextMsgReq{
						Touser:  []string{duk},
						Msgtype: "text",
						Content: &wechat.TextContent{
							Content: fmt.Sprintf("%s-项目任务提醒！！ \n 研发版块 \n %s", dk, strings.Join(duv, "\n")),
						},
					})
				}
			}
		}
	}
	if len(joinMsgData) > 0 {
		for dk, dv := range joinMsgData {
			if len(dv) > 0 {
				for duk, duv := range dv {
					boot.WechatCheckIn.SendMsg(ctx, &wechat.SendTextMsgReq{
						Touser:  []string{duk},
						Msgtype: "text",
						Content: &wechat.TextContent{
							Content: fmt.Sprintf("%s-项目任务提醒！！\n 研发版块 %s \n请跟进任务进展！", dk, strings.Join(duv, "\n")),
						},
					})
				}
			}
		}
	}
}

// AutoRemindNotStart 今天未开启
func (s *taskDevelopService) AutoRemindNotStart() {
	fmt.Println("AutoRemindNotStart------------", gtime.Now().String())
	ctx := context.Background()
	// dutyMsgData 责任人
	dutyMsgData := make(map[string]map[string][]string, 0)
	// joinMsgData 关联责任人
	joinMsgData := make(map[string]map[string][]string, 0)

	resData, _ := dao.ProductTaskDevelop.GetAll(ctx, &entity.ProductTaskDevelop{
		TaskStatus:    1,
		PalnStartTime: gtime.Now(),
	})

	if len(resData) > 0 {
		for _, v := range resData {
			if !g.IsEmpty(v.DutyWorkNumber) {
				if g.IsNil(dutyMsgData[v.Name]) {
					dutyMsgData[v.Name] = make(map[string][]string)
				}
				for _, uv := range strings.Split(v.DutyWorkNumber, "|") {
					dutyMsgData[v.Name][uv] = append(dutyMsgData[v.Name][uv], fmt.Sprintf("任务：%s,需要在%s开启, 任务描述：%s", v.TaskName, v.PalnEndTime.Format("Y-m-d"), v.TaskDesc))
				}
			}
			if !g.IsEmpty(v.JoinWorkNumber) {
				if g.IsNil(joinMsgData[v.Name]) {
					joinMsgData[v.Name] = make(map[string][]string)
				}
				for _, jv := range strings.Split(v.JoinWorkNumber, "|") {
					joinMsgData[v.Name][jv] = append(joinMsgData[v.Name][jv], fmt.Sprintf("任务：%s,需要在%s开启, 任务描述：%s", v.TaskName, v.PalnEndTime.Format("Y-m-d"), v.TaskDesc))
				}
			}
		}
	}

	if len(dutyMsgData) > 0 {
		for dk, dv := range dutyMsgData {
			if len(dv) > 0 {
				for duk, duv := range dv {
					boot.WechatCheckIn.SendMsg(ctx, &wechat.SendTextMsgReq{
						Touser:  []string{duk},
						Msgtype: "text",
						Content: &wechat.TextContent{
							Content: fmt.Sprintf("%s-项目任务提醒！！ \n 研发版块 \n %s", dk, strings.Join(duv, "\n")),
						},
					})
				}
			}
		}
	}
	if len(joinMsgData) > 0 {
		for dk, dv := range joinMsgData {
			if len(dv) > 0 {
				for duk, duv := range dv {
					boot.WechatCheckIn.SendMsg(ctx, &wechat.SendTextMsgReq{
						Touser:  []string{duk},
						Msgtype: "text",
						Content: &wechat.TextContent{
							Content: fmt.Sprintf("%s-项目任务提醒！！ \n 研发版块 %s \n请跟进任务进展！", dk, strings.Join(duv, "\n")),
						},
					})
				}
			}
		}
	}
}

// AutoUpgrade 升级提醒
func (s *taskDevelopService) AutoUpgrade() {
	fmt.Println("AutoUpgrade------------", gtime.Now().String())
	ctx := context.Background()
	// threeMsgData 第三次升级时提醒人
	threeSendUser, _ := s.autoUpgradeSearch(ctx, &entity.ProductTaskDevelop{
		TaskStatus:       2,
		UpgradeThreeTime: gtime.Now(),
	}, "第三次", []string{})

	// twoMsgData 第二次升级时提醒人
	twoSendUser, _ := s.autoUpgradeSearch(ctx, &entity.ProductTaskDevelop{
		TaskStatus:     2,
		UpgradeTwoTime: gtime.Now(),
	}, "第二次", threeSendUser)

	// 第一次升级数据
	twoSendUser = append(twoSendUser, threeSendUser...)
	s.autoUpgradeSearch(ctx, &entity.ProductTaskDevelop{
		TaskStatus:       2,
		UpgradeFirstTime: gtime.Now(),
	}, "第一次", twoSendUser)
}

func (s *taskDevelopService) autoUpgradeSearch(ctx context.Context, where *entity.ProductTaskDevelop, keyStr string, notSendList []string) (sendUserList []string, err error) {
	// msgData 升级时提醒人
	msgData := make(map[string]map[string][]string, 0)
	sendUserList = make([]string, 0)
	// 未完成升级
	upData, err := dao.ProductTaskDevelop.GetUpgradeAll(ctx, where)

	if len(upData) > 0 {
		for _, v := range upData {
			if !g.IsEmpty(v.UpgradeFirst) {
				if g.IsNil(msgData[v.Name]) {
					msgData[v.Name] = make(map[string][]string)
				}
				for _, uv := range strings.Split(v.UpgradeFirst, "|") {
					msgData[v.Name][uv] = append(msgData[v.Name][uv], fmt.Sprintf("任务：%s,在计划结束时间%s内未完成, 任务描述：%s", v.TaskName, v.PalnEndTime.Format("Y-m-d"), v.TaskDesc))
					sendUserList = append(sendUserList, uv)
				}
			}
		}
	}

	if len(msgData) > 0 {
		for dk, dv := range msgData {
			if len(dv) > 0 {
				for duk, duv := range dv {
					if !util.CheckInStr(notSendList, duk) {
						boot.WechatCheckIn.SendMsg(ctx, &wechat.SendTextMsgReq{
							Touser:  []string{duk},
							Msgtype: "text",
							Content: &wechat.TextContent{
								Content: fmt.Sprintf("%s-任务延期%s升级提醒！！ \n 研发版块 \n %s \n请跟进任务进展！", dk, keyStr, strings.Join(duv, "\n")),
							},
						})
					}
				}
			}
		}
	}
	return sendUserList, nil
}

// GetDevelopList 获取员工信息列表
func (s *taskDevelopService) GetDevelopList(ctx context.Context, input *model.ProductTaskWhere) (apiRes *response.GetListResponse, err error) {
	apiRes, err = dao.ProductTaskDevelop.GetList(ctx, input)
	if err != nil {
		return apiRes, err
	}

	return apiRes, nil
}

// DevelopModify 任务状态变更
func (s *taskDevelopService) DevelopModify(ctx context.Context, input *model.ProductTaskWhere) (apiRes *response.GetListResponse, err error) {
	apiRes = &response.GetListResponse{}
	info, err := dao.ProductTaskDevelop.GetOne(ctx, input)
	if err != nil {
		return apiRes, err
	}
	oldStatus := info.TaskStatus
	info.TaskStatus = input.TaskStatus
	info.Remark = input.Remark
	if input.TaskStatus == 3 {
		info.RealEndTime = gtime.Now()
	}
	if input.TaskStatus == 2 {
		info.RealStartTime = gtime.Now()
	}

	apiRes.Data, err = dao.ProductTaskDevelop.Modify(ctx, info)
	touser := strings.Split(info.DutyWorkNumber, "|")
	touser = append(touser, strings.Split(info.JoinWorkNumber, "|")...)

	boot.WechatCheckIn.SendMsg(ctx, &wechat.SendTextMsgReq{
		Touser:  touser,
		Msgtype: "text",
		Content: &wechat.TextContent{
			Content: util.GetTaskStatusChangeMsg(info.SubName, info.TaskName, info.TaskDesc, oldStatus, input.TaskStatus),
		},
	})
	return apiRes, nil
}
