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
	inspirit "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	product "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

var Product = productService{}

type productService struct{}

// GetList 项目清单
func (s *productService) GetList(ctx context.Context, in *model.ProductApiGetListReq) (res *response.GetListResponse, err error) {
	res, err = dao.Product.GetList(ctx, in.Product, in.Page, in.Size)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *productService) GetOne(ctx context.Context, in *model.ProductApiGetOneReq) (res *entity.Product, err error) {
	res, err = dao.Product.GetOne(ctx, in.Product)
	if err != nil {
		return res, err
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

	res, err = dao.Product.Modify(ctx, data)
	return res, err
}

func (s *productService) checkInputData(ctx context.Context, in *model.ProductApiChangeReq) (*model.ProductApiChangeReq, error) {
	conditionName := g.Map{
		fmt.Sprintf("%s = ?", dao.Product.Columns().Name): in.Name,
	}
	conditionSubName := g.Map{
		fmt.Sprintf("%s = ?", dao.Product.Columns().SubName): in.SubName,
	}
	if in.Id > 0 {
		conditionName["id != ?"] = in.Id
		conditionSubName["id != ?"] = in.Id
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

	// 根据优先级评分计算优先级
	if in.LccId, in.LccName, err = s.getLccByLcScore(ctx, in.LcScore); err != nil {
		return in, err
	}

	// 根据评分计算总激励预算
	in.IncentiveBudget, err = s.getIncentiveBudgetByLcScore(ctx, in.LcScore, in.FixType, in.NetProfit, in.FixBudget)

	return in, err
}

func (s *productService) getLccByLcScore(ctx context.Context, lcScore uint) (lccId uint, lccName string, err error) {
	// 查询项目优先级确认配置信息
	levelConfirmList, err := boot.LevelConfirmServer.GetList(ctx, &product.GetListLevelConfirmReq{
		Page: 1,
		Size: 100,
	})
	if err != nil {
		return 0, "", err
	}
	if g.IsEmpty(levelConfirmList.TotalSize) {
		return 0, "", errors.New("请先完善项目优先级相关配置信息")
	}

	for _, v := range levelConfirmList.GetData() {
		switch v.ScoreRange {
		case consts.ScoreRangeMin:
			// 左闭右开
			if v.ScoreMin <= gconv.Float32(lcScore) && gconv.Float32(lcScore) < v.ScoreMax {
				return gconv.Uint(v.Id), v.Name, nil
			}
		case consts.ScoreRangeMax:
			// 左开右闭
			if v.ScoreMin < gconv.Float32(lcScore) && gconv.Float32(lcScore) <= v.ScoreMax {
				return gconv.Uint(v.Id), v.Name, nil
			}
		case consts.ScoreRangeMinAndMax:
			// 左闭右闭
			if v.ScoreMin <= gconv.Float32(lcScore) && gconv.Float32(lcScore) <= v.ScoreMax {
				return gconv.Uint(v.Id), v.Name, nil
			}
		}
	}
	return
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

	checkLcScore := gconv.Uint32(lcScore)
	lcBudgetInfo := &inspirit.BudgetAssessInfo{}
	for _, v := range levelAssessList.GetData() {
		switch v.ScoreRange {
		case consts.ScoreRangeMin:
			// 左闭右开
			if v.ScoreMin <= checkLcScore && checkLcScore < v.ScoreMax {
				lcBudgetInfo = v
				break
			}
		case consts.ScoreRangeMax:
			// 左开右闭
			if v.ScoreMin < checkLcScore && checkLcScore <= v.ScoreMax {
				lcBudgetInfo = v
				break
			}
		case consts.ScoreRangeMinAndMax:
			// 左闭右闭
			if v.ScoreMin <= checkLcScore && checkLcScore <= v.ScoreMax {
				lcBudgetInfo = v
				break
			}
		}
	}
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
