// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao/internal"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/do"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
)

// internalProductTaskQualityDao is internal type for wrapping internal DAO implements.
type internalProductTaskQualityDao = *internal.ProductTaskQualityDao

// productTaskQualityDao is the data access object for table cqgf_product_task_quality.
// You can define custom methods on it to extend its functionality as you wish.
type productTaskQualityDao struct {
	internalProductTaskQualityDao
}

var (
	// ProductTaskQuality is globally public accessible object for table cqgf_product_task_quality operations.
	ProductTaskQuality = productTaskQualityDao{
		internal.NewProductTaskQualityDao(),
	}
)

// Fill with you ideas below.
func (s *productTaskQualityDao) GetAll(ctx context.Context, in *entity.ProductTaskQuality) (res []*model.ProductTaskQualityField, err error) {
	res = make([]*model.ProductTaskQualityField, 0)
	query := s.Ctx(ctx)
	// 项目经理投入程度
	if in.Id > 0 {
		query = query.Where(s.Columns().Id, in.Id)
	}
	if in.ProductId > 0 {
		query = query.Where(s.Columns().ProductId, in.ProductId)
	}
	if in.TaskStatus > 0 {
		query = query.Where(s.Columns().TaskStatus, in.TaskStatus)
	}
	if !g.IsEmpty(in.PalnStartTime) {
		gt1 := gtime.New(in.PalnStartTime)
		query = query.WhereGTE(s.Columns().PalnStartTime, gt1.StartOfDay())
	}
	if !g.IsEmpty(in.PalnEndTime) {
		gt2 := gtime.New(in.PalnEndTime)
		query = query.WhereLTE(s.Columns().PalnEndTime, gt2.EndOfDay())
	}
	if !g.IsEmpty(in.UpgradeFirstTime) {
		query = query.WhereLTE(s.Columns().UpgradeFirstTime, in.UpgradeFirstTime)
	}
	if !g.IsEmpty(in.UpgradeTwoTime) {
		query = query.WhereLTE(s.Columns().UpgradeTwoTime, in.UpgradeTwoTime)
	}
	if !g.IsEmpty(in.UpgradeThreeTime) {
		query = query.WhereLTE(s.Columns().UpgradeThreeTime, in.UpgradeThreeTime)
	}

	err = query.InnerJoin(Product.Table(), fmt.Sprintf("%s.%s=%s.%s", s.Table(), s.Columns().ProductId, Product.Table(), Product.Columns().Id)).
		FieldsPrefix(s.Table(), s.Columns()).
		FieldsPrefix(Product.Table(), Product.Columns().Name).
		FieldsPrefix(Product.Table(), Product.Columns().SubName).
		Scan(&res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *productTaskQualityDao) GetUpgradeAll(ctx context.Context, in *entity.ProductTaskQuality) (res []*model.ProductTaskQualityField, err error) {
	res = make([]*model.ProductTaskQualityField, 0)
	query := s.Ctx(ctx).WhereLT(s.Columns().PalnEndTime, gtime.Now())

	if in.ProductId > 0 {
		query = query.Where(s.Columns().ProductId, in.ProductId)
	}
	if in.TaskStatus > 0 {
		query = query.Where(s.Columns().TaskStatus, in.TaskStatus)
	}
	if !g.IsEmpty(in.UpgradeFirstTime) {
		query = query.WhereLTE(s.Columns().UpgradeFirstTime, in.UpgradeFirstTime)
	}
	if !g.IsEmpty(in.UpgradeTwoTime) {
		query = query.WhereLTE(s.Columns().UpgradeTwoTime, in.UpgradeTwoTime)
	}
	if !g.IsEmpty(in.UpgradeThreeTime) {
		query = query.WhereLTE(s.Columns().UpgradeThreeTime, in.UpgradeThreeTime)
	}

	err = query.InnerJoin(Product.Table(), fmt.Sprintf("%s.%s=%s.%s", s.Table(), s.Columns().ProductId, Product.Table(), Product.Columns().Id)).
		FieldsPrefix(s.Table(), s.Columns()).
		FieldsPrefix(Product.Table(), Product.Columns().Name).
		FieldsPrefix(Product.Table(), Product.Columns().SubName).
		Scan(&res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *productTaskQualityDao) GetList(ctx context.Context, in *model.ProductTaskWhere) (res *response.GetListResponse, err error) {
	res = &response.GetListResponse{}
	resData := make([]*model.ProductTaskQualityField, 0)
	query := s.Ctx(ctx).InnerJoin(Product.Table(), fmt.Sprintf("%s.%s=%s.%s", s.Table(), s.Columns().ProductId, Product.Table(), Product.Columns().Id))
	// 项目经理投入程度
	if in.Id > 0 {
		query = query.Where(s.Columns().Id, in.Id)
	}
	if in.ProID > 0 {
		query = query.Where(s.Columns().ProductId, in.ProID)
	}
	if len(in.TaskName) > 0 {
		query = query.Where(s.Columns().TaskName, in.TaskName)
	}
	if len(in.Name) > 0 {
		query = query.Where(fmt.Sprintf("%s.%s", Product.Table(), Product.Columns().Name), in.Name)
	}
	if len(in.SubName) > 0 {
		query = query.Where(fmt.Sprintf("%s.%s", Product.Table(), Product.Columns().SubName), in.SubName)
	}

	if in.ProductId > 0 {
		query = query.Where(s.Columns().ProductId, in.ProductId)
	}
	if len(in.DutyWorkNumber) > 0 {
		query = query.Where(fmt.Sprintf("FIND_IN_SET('%s', replace(%s.%s,'|',',')) > 0 ", in.DutyWorkNumber, s.Table(), s.Columns().DutyWorkNumber))
	}
	if len(in.JoinWorkNumber) > 0 {
		query = query.Where(fmt.Sprintf("FIND_IN_SET('%s', replace(%s.%s,'|',',')) > 0 ", in.JoinWorkNumber, s.Table(), s.Columns().JoinWorkNumber))
	}
	if in.TaskStatus > 0 {
		query = query.Where(s.Columns().TaskStatus, in.TaskStatus)
	}
	if !g.IsEmpty(in.PalnStartTime) {
		gt1 := gtime.New(in.PalnStartTime)
		query = query.WhereGTE(s.Columns().PalnStartTime, gt1.StartOfDay())
		//query = query.WhereLTE(s.Columns().PalnStartTime, gt1.EndOfDay())
	}
	if !g.IsEmpty(in.PalnEndTime) {
		gt2 := gtime.New(in.PalnEndTime)
		//query = query.WhereGTE(s.Columns().PalnEndTime, gt2.StartOfDay())
		query = query.WhereLTE(s.Columns().PalnEndTime, gt2.EndOfDay())
	}
	if !g.IsEmpty(in.UpgradeFirstTime) {
		query = query.WhereLTE(s.Columns().UpgradeFirstTime, in.UpgradeFirstTime)
	}
	if !g.IsEmpty(in.UpgradeTwoTime) {
		query = query.WhereLTE(s.Columns().UpgradeTwoTime, in.UpgradeTwoTime)
	}
	if !g.IsEmpty(in.UpgradeThreeTime) {
		query = query.WhereLTE(s.Columns().UpgradeThreeTime, in.UpgradeThreeTime)
	}

	query, totalSize, page, size, err := util.GetListWithPage(query, in.Page, in.Size)

	err = query.FieldsPrefix(s.Table(), s.Columns()).
		FieldsPrefix(Product.Table(), Product.Columns().Name).
		FieldsPrefix(Product.Table(), Product.Columns().SubName).
		FieldsPrefix(Product.Table(), Product.Columns().ProNumber).Scan(&resData)

	if err != nil {
		return res, err
	}

	res.Page = page
	res.Size = size
	res.TotalSize = totalSize
	res.Data = resData
	return res, nil
}

func (s *productTaskQualityDao) GetOne(ctx context.Context, in *model.ProductTaskWhere) (res model.ProductTaskQualityField, err error) {
	productEntity := model.ProductTaskQualityField{}
	query := s.Ctx(ctx)
	// 项目名称
	if len(in.TaskName) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", s.Columns().TaskName), g.Slice{fmt.Sprintf("%s%s", in.TaskName, "%")})
	}
	if in.Id > 0 {
		query = query.Where(s.Columns().Id, in.Id)
	}

	query = query.InnerJoin(Product.Table(), fmt.Sprintf("%s.%s=%s.%s", s.Table(), s.Columns().ProductId, Product.Table(), Product.Columns().Id)).
		FieldsPrefix(s.Table(), s.Columns()).
		FieldsPrefix(Product.Table(), Product.Columns().Name).
		FieldsPrefix(Product.Table(), Product.Columns().SubName)

	if err = query.Scan(&productEntity); err != nil {
		return res, err
	}

	return productEntity, nil
}

// Modify 编辑项目基础数据
func (s *productTaskQualityDao) Modify(ctx context.Context, in model.ProductTaskQualityField) (model.ProductTaskQualityField, error) {
	data := do.ProductTaskQuality{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return in, err
	}

	data.UpdateTime = gtime.Now()
	_, err = s.Ctx(ctx).Where(s.Columns().Id, in.Id).Data(data).Update()
	if err != nil {
		return in, err
	}

	return in, nil
}
