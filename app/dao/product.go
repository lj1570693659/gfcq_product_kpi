// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao/internal"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/do"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
)

// internalProductDao is internal type for wrapping internal DAO implements.
type internalProductDao = *internal.ProductDao

// productDao is the data access object for table cqgf_product.
// You can define custom methods on it to extend its functionality as you wish.
type productDao struct {
	internalProductDao
}

var (
	// Product is globally public accessible object for table cqgf_product operations.
	Product = productDao{
		internal.NewProductDao(),
	}
)

// Fill with you ideas below.

func (s *productDao) GetList(ctx context.Context, in model.ProductWhere, page, size int32) (res *response.GetListResponse, productEntity []model.Product, err error) {
	res = &response.GetListResponse{}
	productEntity = make([]model.Product, 0)
	query := s.Ctx(ctx)
	// 项目名称
	if len(in.Name) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", s.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.Name, "%")})
	}
	// 项目简称
	if len(in.SubName) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", s.Columns().SubName), g.Slice{fmt.Sprintf("%s%s", in.SubName, "%")})
	}
	// 项目优先级ID
	if len(in.LccId) > 0 {
		query = query.WhereIn(s.Columns().LccId, in.LccId)
	}
	// 项目类型ID
	if len(in.Tid) > 0 {
		query = query.WhereIn(s.Columns().Tid, in.Tid)
	}
	// 项目经理ID
	if len(in.PmId) > 0 {
		query = query.WhereIn(s.Columns().PmId, in.PmId)
	}
	// 项目责任人ID
	if len(in.PmlId) > 0 {
		query = query.WhereIn(s.Columns().PmlId, in.PmId)
	}
	// 项目经理投入程度
	if in.Attribute > 0 {
		query = query.Where(s.Columns().Attribute, in.Attribute)
	}
	// 项目当前状态
	if len(in.Status) > 0 {
		query = query.WhereIn(s.Columns().Status, in.PmId)
	}

	query, totalSize, page, size, err := util.GetListWithPage(query, page, size)
	if err != nil {
		return res, productEntity, err
	}

	if err = query.Scan(&productEntity); err != nil {
		return res, productEntity, err
	}

	res.Page = page
	res.Size = size
	res.TotalSize = totalSize
	res.Data = productEntity
	return res, productEntity, nil
}

func (s *productDao) GetOne(ctx context.Context, in model.Product) (res model.Product, err error) {
	productEntity := model.Product{}
	query := s.Ctx(ctx)
	// 项目名称
	if len(in.Name) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", s.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.Name, "%")})
	}
	// 项目简称
	if len(in.SubName) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", s.Columns().SubName), g.Slice{fmt.Sprintf("%s%s", in.SubName, "%")})
	}
	// 项目经理投入程度
	if in.Attribute > 0 {
		query = query.Where(s.Columns().Attribute, in.Attribute)
	}
	if in.Id > 0 {
		query = query.Where(s.Columns().Id, in.Id)
	}

	if err = query.Scan(&productEntity); err != nil {
		return res, err
	}

	return productEntity, nil
}

// Create 创建项目基础数据
func (s *productDao) Create(ctx context.Context, in *entity.Product) (*entity.Product, error) {
	data := do.Product{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return in, err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	lastInsertId, err := s.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return in, err
	}

	in.Id = gconv.Uint(lastInsertId)
	return in, nil
}

// Modify 编辑项目基础数据
func (s *productDao) Modify(ctx context.Context, in *entity.Product) (*entity.Product, error) {
	data := do.Product{}
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

func (s *productDao) Delete(ctx context.Context, id uint) (bool, error) {
	_, err := s.Ctx(ctx).Where(s.Columns().Id, id).Delete()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *productDao) GetOneByCondition(ctx context.Context, condition g.Map) (*entity.Product, error) {
	info := &entity.Product{}
	err := s.Ctx(ctx).Where(condition).Scan(info)
	return info, err
}

func (s *productDao) GetAll(ctx context.Context, in model.ProductWhere) (res []model.Product, err error) {
	res = make([]model.Product, 0)
	query := s.Ctx(ctx)
	// 项目名称
	if len(in.Name) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", s.Columns().Name), g.Slice{fmt.Sprintf("%s%s", in.Name, "%")})
	}
	// 项目简称
	if len(in.SubName) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", s.Columns().SubName), g.Slice{fmt.Sprintf("%s%s", in.SubName, "%")})
	}
	// 项目优先级ID
	if len(in.LccId) > 0 {
		query = query.WhereIn(s.Columns().LccId, in.LccId)
	}
	// 项目类型ID
	if len(in.Tid) > 0 {
		query = query.WhereIn(s.Columns().Tid, in.Tid)
	}
	// 项目经理ID
	if len(in.PmId) > 0 {
		query = query.WhereIn(s.Columns().PmId, in.PmId)
	}
	// 项目责任人ID
	if len(in.PmlId) > 0 {
		query = query.WhereIn(s.Columns().PmlId, in.PmId)
	}
	// 项目经理投入程度
	if in.Attribute > 0 {
		query = query.Where(s.Columns().Attribute, in.Attribute)
	}
	// 项目当前状态
	if len(in.Status) > 0 {
		query = query.WhereIn(s.Columns().Status, in.PmId)
	}

	err = query.Scan(&res)
	return res, err
}