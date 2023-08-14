package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/consts"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	common "github.com/lj1570693659/gfcq_protoc/common/v1"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

var ProductMemberPrize = productMemberPrizeService{}

var HoursIndexLists []*v1.CrewHoursIndexInfo
var ManageIndexLists []*v1.CrewManageIndexInfo
var DutyIndexLists []*v1.CrewDutyIndexInfo

// HoursIndexRadio 基准指数中工时指数占比
var HoursIndexRadio float64

// ManageIndexRadio 基准指数中管理指数占比
var ManageIndexRadio float64

// DutyIndexRadio 基准指数中责任指数占比
var DutyIndexRadio float64

type productMemberPrizeService struct{}

// 初始化数据
func init() {
	ctx := context.Background()
	res, err := boot.CrewHoursIndexServer.GetAll(ctx, &v1.GetAllCrewHoursIndexReq{CrewHoursIndex: &v1.CrewHoursIndexInfo{}})
	if err != nil {
		g.Log("config").Error(ctx, err)
	}
	if g.IsEmpty(res.GetData()) {
		panic("工时指数未配置，请先完善数据")
	}
	HoursIndexLists = res.GetData()
	hoursIndexRadio, err := dao.Config.GetKeyValueByKeyName(ctx, consts.HoursIndexRadio)
	if err != nil {
		g.Log("config").Error(ctx, err)
	}
	if g.IsEmpty(hoursIndexRadio) {
		panic("工时指数占比未配置，请先完善数据")
	}
	HoursIndexRadio = util.Decimal(gconv.Float64(hoursIndexRadio))

	resManage, err := boot.CrewManageIndexServer.GetAll(ctx, &v1.GetAllCrewManageIndexReq{CrewManageIndex: &v1.CrewManageIndexInfo{}})
	if err != nil {
		g.Log("config").Error(ctx, err)
	}
	if g.IsEmpty(resManage.GetData()) {
		panic("管理指数未配置，请先完善数据")
	}
	ManageIndexLists = resManage.GetData()
	manageIndexRadio, err := dao.Config.GetKeyValueByKeyName(ctx, consts.ManageIndexRadio)
	if err != nil {
		g.Log("config").Error(ctx, err)
	}
	if g.IsEmpty(manageIndexRadio) {
		panic("管理指数占比未配置，请先完善数据")
	}
	ManageIndexRadio = util.Decimal(gconv.Float64(manageIndexRadio))

	resDuty, err := boot.CrewDutyIndexServer.GetAll(ctx, &v1.GetAllCrewDutyIndexReq{CrewDutyIndex: &v1.CrewDutyIndexInfo{}})
	if err != nil {
		g.Log("config").Error(ctx, err)
	}
	if g.IsEmpty(resDuty.GetData()) {
		panic("责任指数未配置，请先完善数据")
	}
	DutyIndexLists = resDuty.GetData()
	dutyIndexRadio, err := dao.Config.GetKeyValueByKeyName(ctx, consts.DutyIndexRadio)
	if err != nil {
		g.Log("config").Error(ctx, err)
	}
	if g.IsEmpty(dutyIndexRadio) {
		panic("责任指数占比未配置，请先完善数据")
	}
	DutyIndexRadio = util.Decimal(gconv.Float64(dutyIndexRadio))
}

// MemberBaseIndexChange 更新项目成员基准指数
func (s *productMemberPrizeService) MemberBaseIndexChange(ctx context.Context, in *model.ProductMemberKpi) (err error) {
	prize := &model.ProductMemberPrize{
		ProEmpId:      in.ProEmpId,
		ProId:         in.ProId,
		ProStageId:    in.ProStageId,
		OvertimeRadio: in.OvertimeRadio,
		KpiLevel:      in.KpiLevel,
		KpiRadio:      in.KpiRadio,
		FloatRaio:     in.FloatRaio,
	}
	// 1：项目组成员
	// 基准指数、权重基准（自动）、权重基准（PMO）、发放基数、剩余额度、实发额度
	// 工时指数
	overtimeIndex, err := s.getHoursIndexByRadio(ctx, gconv.Float32(in.OvertimeRadio))
	if err != nil {
		return err
	}
	prize.OvertimeIndex = gconv.Uint(overtimeIndex)
	// 管理指数
	memberInfo, err := dao.ProductMember.GetOne(ctx, model.ProductMember{Id: in.ProEmpId})
	if err != nil {
		return err
	}
	// 责任指数
	prize.DutyIndex = gconv.Uint(memberInfo.DutyIndex)
	prize.ManageIndex = memberInfo.ManageIndex
	prize.PrId = memberInfo.PrId
	prize.PrName = memberInfo.PrName
	prize.JbId = memberInfo.JbId
	prize.JbName = memberInfo.JbName

	// 基准指数
	prize.BaseIndex = gconv.Float64(prize.DutyIndex)*DutyIndexRadio + gconv.Float64(prize.ManageIndex)*ManageIndexRadio + gconv.Float64(prize.OvertimeIndex)*HoursIndexRadio

	info, err := dao.ProductMemberPrize.GetOne(ctx, model.ProductMemberPrize{
		ProId:      in.ProId,
		ProEmpId:   in.ProEmpId,
		ProStageId: in.ProStageId,
	})

	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return err
	}
	prize.Id = in.Id
	prize.IsPm = consts.IsNotPm
	if g.IsEmpty(info.Id) {
		_, err = dao.ProductMemberPrize.Create(ctx, prize)
	} else {
		_, err = dao.ProductMemberPrize.Modify(ctx, prize)
	}
	return err
}

// PmBaseIndexChange 更新项目经理基准指数
func (s *productMemberPrizeService) PmBaseIndexChange(ctx context.Context, in *model.ProductMemberKpi) (err error) {
	// 项目经理奖金分配数据完善
	productStageKpi, err := dao.ProductStageKpi.GetOne(ctx, &model.ProductStageKpi{ProId: in.ProId, StageId: in.ProStageId})
	if err != nil {
		return err
	}

	// 项目组绩效数据完善
	pmMemberInfo, err := dao.ProductMember.GetOne(ctx, model.ProductMember{ProId: in.ProId, IsSpecial: consts.IsPm})
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return err
	}

	getPmPrize, err := dao.ProductMemberPrize.GetOne(ctx, model.ProductMemberPrize{
		ProId:      in.ProId,
		ProStageId: in.ProStageId,
		IsPm:       consts.IsPm,
	})

	data := &model.ProductMemberPrize{
		Id:             getPmPrize.Id,
		ProId:          in.ProId,
		IsPm:           consts.IsPm,
		ProEmpId:       pmMemberInfo.Id,
		ProStageId:     in.ProStageId,
		ManageIndex:    pmMemberInfo.ManageIndex,
		PrName:         pmMemberInfo.PrName,
		JbId:           pmMemberInfo.JbId,
		JbName:         pmMemberInfo.JbName,
		DutyIndex:      gconv.Uint(pmMemberInfo.DutyIndex),
		WeightPmoRadio: productStageKpi.PmRadio,
		SentBase:       productStageKpi.PmBase,
		RemaindQueto:   productStageKpi.CrewQuota - productStageKpi.PmBase,
		FloatRaio:      productStageKpi.PmFloatRadio,
		KpiLevelId:     productStageKpi.PmKpiLevelId,
		KpiLevel:       productStageKpi.PmKpiLevelName,
		KpiRadio:       productStageKpi.PmKpiLevelRadio,
		SentQueto:      productStageKpi.PmIncentiveQuota,
	}
	fmt.Println("pm.data----------------", data)
	fmt.Println("pm.getPmPrize----------------", getPmPrize)
	fmt.Println("pm.productStageKpi----------------", productStageKpi)
	if g.IsEmpty(getPmPrize.Id) {
		getPmPrize, err = dao.ProductMemberPrize.Create(ctx, data)
		if err != nil {
			return err
		}
	} else {
		getPmPrize, err = dao.ProductMemberPrize.Modify(ctx, data)
	}
	return err
}

// 根据工时占比获取工时指数
func (s *productMemberPrizeService) getHoursIndexByRadio(ctx context.Context, overtimeRadio float32) (uint32, error) {
	if g.IsEmpty(overtimeRadio) {
		return 0, nil
	}
	overtimeIndex := util.GetHoursIndexByScore(HoursIndexLists, overtimeRadio)
	return overtimeIndex, nil
}

// Compute 计算成员奖金
func (s *productMemberPrizeService) Compute(ctx context.Context, in *model.ProductMemberPrizeComputeReq) (err error) {
	// 项目经理奖金分配数据查询
	productStageKpi, err := dao.ProductStageKpi.GetOne(ctx, &model.ProductStageKpi{ProId: in.ProId, StageId: in.StageId})
	if err != nil {
		return err
	}

	// 项目组奖金数据查询
	getPmPrize := &model.ProductMemberPrize{}
	getPmPrize, err = dao.ProductMemberPrize.GetOne(ctx, model.ProductMemberPrize{ProId: in.ProId, ProStageId: in.StageId, IsPm: consts.IsPm})
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return err
	}

	// 项目组成员全部数据
	getAllPrize, err := dao.ProductMemberPrize.GetAll(ctx, model.ProductMemberPrize{ProId: in.ProId, ProStageId: in.StageId, IsPm: consts.IsNotPm})
	if err != nil {
		return err
	}

	// 成员奖金基准指数和
	memberBaseIndexSum, err := dao.ProductMemberPrize.GetFieldSum(ctx, model.ProductMemberPrize{ProId: in.ProId, ProStageId: in.StageId, IsPm: consts.IsNotPm}, dao.ProductMemberPrize.Columns().BaseIndex)
	if err != nil {
		return err
	}

	memberBaseQuota := getPmPrize.SentBase
	for k, v := range getAllPrize {
		//权重基准（自动）、权重基准（PMO）、发放基数、剩余额度、实发额度
		// 权重基准（自动） TODO
		//getAllPrize[k].WeightAutoRadio = v.BaseIndex / (memberBaseIndexSum + productStageKpi.PmRadio)
		// 权重基准（PMO）
		getAllPrize[k].WeightPmoRadio = (1 - productStageKpi.PmRadio) * (v.BaseIndex / memberBaseIndexSum)
		// 发放基数 = 团队额度 * 成员比例
		getAllPrize[k].SentBase = getAllPrize[k].WeightPmoRadio * productStageKpi.CrewQuota
		memberBaseQuota = memberBaseQuota + getAllPrize[k].SentBase
		// 剩余额度
		if (productStageKpi.CrewQuota - memberBaseQuota) < 0 {
			getAllPrize[k].RemaindQueto = 0.00
		} else {
			getAllPrize[k].RemaindQueto = productStageKpi.CrewQuota - memberBaseQuota
		}
		// 实发额度
		getAllPrize[k].SentQueto = getAllPrize[k].SentBase * (v.FloatRaio + v.KpiRadio)
	}
	for _, v := range getAllPrize {
		err = dao.ProductMemberPrize.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			if _, err = dao.ProductMemberPrize.Modify(ctx, v); err != nil {
				return err
			}
			return nil
		})
	}

	return err
}

func (s *productMemberPrizeService) GetList(ctx context.Context, in model.ProductMemberPrizeApiGetListReq) (res *response.GetListResponse, err error) {
	resData := make([]model.ProductMemberPrizeApiGetListRes, 0)
	res, entity, err := dao.ProductMemberPrize.GetList(ctx, in)
	if err != nil {
		return res, err
	}

	// 部门清单
	departmentList, err := boot.DepertmentServer.GetListWithoutPage(ctx, &common.GetListWithoutDepartmentReq{})
	if err != nil {
		return res, err
	}

	if res.TotalSize > 0 {
		for _, v := range entity {
			info := model.ProductMemberPrizeApiGetListRes{
				ProductMemberPrize: v,
			}
			info.ProductMemberKpi, err = dao.ProductMemberKpi.GetOne(ctx, model.ProductMemberKpi{ProEmpId: v.ProEmpId})
			if err != nil {
				return res, err
			}

			// 项目组成员信息
			info.ProductMember, err = ProductMember.GetOne(ctx, &model.ProductMemberApiGetOneReq{model.ProductMember{Id: v.ProEmpId}})
			if err != nil {
				return res, err
			}

			empInfo, err := boot.EmployeeServer.GetOne(ctx, &common.GetOneEmployeeReq{Id: gconv.Int32(info.ProductMember.EmpId)})
			if err != nil {
				return res, err
			}
			info.UserName = empInfo.GetEmployee().GetUserName()
			info.DepartmentName = Department.GetDepartmentName(empInfo.GetEmployee().GetDepartId(), departmentList.GetData())
			resData = append(resData, info)
		}
	}

	res.Data = resData
	return res, nil
}
