package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/consts"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	common "github.com/lj1570693659/gfcq_protoc/common/v1"
	inspirit "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	product "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

var Product = productService{}

type productService struct{}

// GetList 项目清单
func (s *productService) GetList(ctx context.Context, in *model.ProductApiGetListReq) (res *response.GetListResponse, err error) {
	res = &response.GetListResponse{}
	resData := make([]model.GetProduct, 0)
	productList, productEntity, err := dao.Product.GetList(ctx, in)
	if err != nil {
		return res, err
	}

	modeList, err := boot.ModeServer.GetAll(ctx, &product.GetAllModeReq{})
	if err != nil {
		return res, err
	}

	if productList.TotalSize > 0 {
		for _, v := range productEntity {
			info := model.GetProduct{
				ProductInfo: v,
				ProductPm:   model.Employee{},
				ProductPml:  model.Employee{},
			}
			// 研发模式
			info.ProductMode, err = s.getModeInfo(gconv.Int32(v.ModeId), modeList.GetData())
			if err != nil {
				return res, err
			}
			// 项目所处阶段
			info.ProductStage, err = s.getStageInfo(ctx, v.ProTypeStageId)
			if err != nil {
				return res, err
			}
			//项目经理
			if v.PmId > 0 {
				pmInfo, err := boot.EmployeeServer.GetOne(ctx, &common.GetOneEmployeeReq{Id: gconv.Int32(v.PmId)})

				if err != nil {
					return res, err
				}
				gconv.Struct(pmInfo.GetEmployee(), &info.ProductPm)
			}

			//项目负责人
			if v.PmlId > 0 {
				pmlInfo, err := boot.EmployeeServer.GetOne(ctx, &common.GetOneEmployeeReq{Id: gconv.Int32(v.PmlId)})
				if err != nil {
					return res, err
				}
				gconv.Struct(pmlInfo.GetEmployee(), &info.ProductPml)
			}

			resData = append(resData, info)
		}
	}

	res.Size = productList.Size
	res.Page = productList.Page
	res.TotalSize = productList.TotalSize
	res.Data = resData
	return res, nil
}

// GetAll 项目筛选清单
func (s *productService) GetAll(ctx context.Context, in model.ProductWhere) (res []model.Product, err error) {
	res, err = dao.Product.GetAll(ctx, in)

	return res, err
}

func (s *productService) GetOne(ctx context.Context, in *model.ProductApiGetOneReq) (res model.Product, err error) {
	res, err = dao.Product.GetOne(ctx, in.Product)
	if err != nil {
		return res, err
	}

	// 项目所处阶段
	stageInfo, err := dao.ProductStageRule.GetOne(ctx, &model.ProductStageRule{Id: res.ProTypeStageId})
	if err != nil {
		return res, err
	}
	res.ProTypeStageId = stageInfo.ProStageId
	return res, nil
}

// GetDetail 项目详情
func (s *productService) GetDetail(ctx context.Context, in *model.ProductApiGetOneReq) (res model.GetProduct, err error) {
	res = model.GetProduct{
		ProductPm:  model.Employee{},
		ProductPml: model.Employee{},
	}
	res.ProductInfo, err = dao.Product.GetOne(ctx, in.Product)
	if err != nil {
		return res, err
	}

	// 研发模式
	modeInfo, err := boot.ModeServer.GetOne(ctx, &product.GetOneModeReq{Mode: &product.ModeInfo{Id: gconv.Int32(res.ProductInfo.ModeId)}})
	if err != nil {
		return res, err
	}
	res.ProductMode = model.Mode{Name: modeInfo.GetMode().GetName()}

	// 项目类型
	typeInfo, err := boot.TypeServer.GetOne(ctx, &product.GetOneTypeReq{Type: &product.TypeInfo{Id: gconv.Int32(res.ProductInfo.Tid)}})
	if err != nil {
		return res, err
	}
	res.ProductType = model.ProductType{
		Name: typeInfo.GetType().GetName(),
	}
	//项目经理
	if res.ProductInfo.PmId > 0 {
		pmInfo, err := boot.EmployeeServer.GetOne(ctx, &common.GetOneEmployeeReq{Id: gconv.Int32(res.ProductInfo.PmId)})

		if err != nil {
			return res, err
		}
		gconv.Struct(pmInfo.GetEmployee(), &res.ProductPm)
	}
	//项目负责人
	if res.ProductInfo.PmlId > 0 {
		pmlInfo, err := boot.EmployeeServer.GetOne(ctx, &common.GetOneEmployeeReq{Id: gconv.Int32(res.ProductInfo.PmlId)})
		if err != nil {
			return res, err
		}
		gconv.Struct(pmlInfo.GetEmployee(), &res.ProductPml)
	}

	// 项目所处阶段
	stageInfo, err := dao.ProductStageRule.GetOne(ctx, &model.ProductStageRule{Id: res.ProductInfo.ProTypeStageId})
	if err != nil {
		return res, err
	}
	res.ProductStage = &model.ModeStage{
		Name:       stageInfo.Name,
		QuotaRadio: stageInfo.QuotaRadio,
	}
	return res, nil
}

func (s *productService) Create(ctx context.Context, in *model.ProductApiChangeReq) (*entity.Product, error) {
	res := &entity.Product{}
	in, err := s.checkInputData(ctx, in)
	if err != nil {
		return res, err
	}

	data := &entity.Product{}
	input, _ := json.Marshal(in)
	err = json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}
	err = dao.Product.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		res, err = dao.Product.Create(ctx, data)

		// 项目阶段占比个性化配置
		err = ProductStageRule.CreateDefault(ctx, in.Tid, res.Id)

		// 查询项目当前自由阶段
		stageInfo, err := dao.ProductStageRule.GetOne(ctx, &model.ProductStageRule{ProId: res.Id, ProStageId: in.ProTypeStageId})
		if err != nil {
			return err
		}
		// 更新项目个性化阶段定制信息
		data.ProTypeStageId = stageInfo.Id
		_, err = dao.Product.Modify(ctx, data)
		return err
	})

	return res, err
}

func (s *productService) Modify(ctx context.Context, in *model.ProductApiChangeReq) (*entity.Product, error) {
	res := &entity.Product{}
	if g.IsEmpty(in.Id) {
		return res, errors.New("缺少编辑对象")
	}

	in, err := s.checkInputData(ctx, in)
	if err != nil {
		return res, err
	}

	data := &entity.Product{}
	input, _ := json.Marshal(in)
	err = json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	// 查询项目当前自由阶段
	stageInfo, err := dao.ProductStageRule.GetOne(ctx, &model.ProductStageRule{ProId: in.Id, ProStageId: in.ProTypeStageId})
	if err != nil {
		return res, err
	}
	// 更新项目个性化阶段定制信息
	data.ProTypeStageId = stageInfo.Id

	res, err = dao.Product.Modify(ctx, data)
	return res, err
}

func (s *productService) Delete(ctx context.Context, in *model.Product) (model.Product, error) {
	res := model.Product{}
	if g.IsEmpty(in.Id) {
		return res, errors.New("缺少删除对象")
	}

	res, err := dao.Product.GetOne(ctx, model.Product{Id: in.Id})
	if err != nil {
		return res, err
	}

	productMember, err := dao.ProductMember.GetOne(ctx, model.ProductMember{ProId: in.Id})
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return res, err
	}
	if !g.IsEmpty(productMember.Id) {
		return res, errors.New("请先移除项目组成员信息")
	}

	err = dao.Product.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.Product.Delete(ctx, in.Id)
		if err != nil {
			return err
		}

		// 查询项目阶段信息
		_, err = dao.ProductStageRule.Delete(ctx, in.Id, 0)
		if err != nil {
			return err
		}
		return nil
	})

	return res, err
}

func (s *productService) checkInputData(ctx context.Context, in *model.ProductApiChangeReq) (*model.ProductApiChangeReq, error) {
	conditionName := g.Map{
		fmt.Sprintf("%s = ?", dao.Product.Columns().Name): in.Name,
	}
	conditionSubName := g.Map{
		fmt.Sprintf("%s = ?", dao.Product.Columns().SubName): in.SubName,
	}
	conditionProNumber := g.Map{
		fmt.Sprintf("%s = ?", dao.Product.Columns().ProNumber): in.ProNumber,
	}
	if in.Id > 0 {
		conditionName["id != ?"] = in.Id
		conditionSubName["id != ?"] = in.Id
		conditionProNumber["id != ?"] = in.Id
	}
	// 1: 项目名称唯一
	uniqueName, err := dao.Product.GetOneByCondition(ctx, conditionName)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}
	if !g.IsNil(uniqueName) && !g.IsEmpty(uniqueName.Id) {
		return in, errors.New("项目名称已存在，请确认输入信息是否正确")
	}

	// 2: 项目简称唯一
	uniqueSubName, err := dao.Product.GetOneByCondition(ctx, conditionSubName)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}
	if !g.IsNil(uniqueSubName) && !g.IsEmpty(uniqueSubName.Id) {
		return in, errors.New("项目简称已存在，请确认输入信息是否正确")
	}

	// 3: 项目编号唯一
	uniqueProNumber, err := dao.Product.GetOneByCondition(ctx, conditionProNumber)
	if err != nil && err != sql.ErrNoRows {
		return in, err
	}
	if !g.IsNil(uniqueProNumber) && !g.IsEmpty(uniqueProNumber.Id) {
		return in, errors.New("项目编号已存在，请确认输入信息是否正确")
	}

	// 根据优先级评分计算优先级
	if info, err := s.getLccByLcScore(ctx, in.LcScore); err != nil {
		in.LccId = gconv.Uint(info.Id)
		in.LccName = info.Name
		return in, err
	}

	// 根据评分计算总激励预算
	in.IncentiveBudget, err = s.getIncentiveBudgetByLcScore(ctx, in.LcScore, in.FixType, in.NetProfit, in.FixBudget)

	return in, err
}

func (s *productService) getLccByLcScore(ctx context.Context, lcScore uint) (confirmInfo *product.LevelConfirmInfo, err error) {
	confirmInfo = &product.LevelConfirmInfo{}
	// 查询项目优先级确认配置信息
	levelConfirmList, err := boot.LevelConfirmServer.GetList(ctx, &product.GetListLevelConfirmReq{
		Page: 1,
		Size: 100,
	})
	if err != nil {
		return confirmInfo, err
	}
	if g.IsEmpty(levelConfirmList.TotalSize) {
		return confirmInfo, errors.New("请先完善项目优先级相关配置信息")
	}

	confirmInfo = util.GetLevelConfirmByScore(levelConfirmList.GetData(), lcScore)

	return confirmInfo, nil
}

func (s *productService) getIncentiveBudgetByLcScore(ctx context.Context, lcScore, fixType uint, netProfit, fixBudget float64) (incentiveBudget float64, err error) {
	var lcBudget float32
	// 查询项目优先级确认配置信息
	levelAssessList, err := boot.BudgetAssessServer.GetList(ctx, &inspirit.GetListBudgetAssessReq{
		Page: 1,
		Size: 100,
	})
	if err != nil {
		return 0, err
	}
	if g.IsEmpty(levelAssessList.TotalSize) {
		return 0, errors.New("请先完善项目预算相关配置信息")
	}

	// 查询配置信息product_budget_by_score_type（项目预算在预算区间中取值方式（1：取最小 2：取最大））
	keyValue, err := dao.Config.GetKeyValueByKeyName(ctx, consts.ProductBudgetByScoreType)
	if err != nil {
		return 0, err
	}
	if g.IsEmpty(keyValue) {
		return 0, errors.New("请先完善项目预算取值配置信息")
	}

	lcBudgetInfo := util.GetLevelAssessByScore(levelAssessList.GetData(), gconv.Uint32(lcScore))

	if gconv.Int(keyValue) == consts.ProductBudgetByMin {
		lcBudget = lcBudgetInfo.GetBudgetMin()
	} else {
		lcBudget = lcBudgetInfo.GetBudgetMax()
	}

	// 查询首年净利润比例
	npValue, err := dao.Config.GetKeyValueByKeyName(ctx, consts.BudgetNpRadio)
	if err != nil {
		return 0, err
	}
	if g.IsEmpty(npValue) {
		return 0, errors.New("请先完善项目首年利润比例配置信息")
	}

	// 首年利润10%和得分对应值向下取值
	budgetMin := 0.00
	if gconv.Float64(npValue)*netProfit < gconv.Float64(lcBudget) {
		budgetMin = gconv.Float64(npValue) * netProfit
	} else {
		budgetMin = gconv.Float64(lcBudget)
	}

	// 预算修正
	incentiveBudget = budgetMin
	if !g.IsEmpty(fixBudget) {
		switch fixType {
		case consts.BudgetFixAdd:
			incentiveBudget = budgetMin + fixBudget
		case consts.BudgetFixLess:
			incentiveBudget = budgetMin - fixBudget
		}
	}

	return incentiveBudget, nil
}

func (s *productService) getModeInfo(modeId int32, modeList []*product.ModeInfo) (res model.Mode, err error) {
	if len(modeList) == 0 {
		return
	}
	for _, v := range modeList {
		if modeId == v.Id {
			gconv.Structs(v, &res)
			return
		}
	}
	return
}

func (s *productService) getStageInfo(ctx context.Context, proTypeStageId uint) (res *model.ModeStage, err error) {
	res = &model.ModeStage{}
	if proTypeStageId == 0 {
		return
	}
	stageInfo, err := dao.ProductStageRule.GetOne(ctx, &model.ProductStageRule{Id: proTypeStageId})
	res.Name = stageInfo.Name
	res.QuotaRadio = stageInfo.QuotaRadio
	return res, err
}
