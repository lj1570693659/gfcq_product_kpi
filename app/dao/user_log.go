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

// internalUserLogDao is internal type for wrapping internal DAO implements.
type internalUserLogDao = *internal.UserLogDao

// userLogDao is the data access object for table cqgf_user_log.
// You can define custom methods on it to extend its functionality as you wish.
type userLogDao struct {
	internalUserLogDao
}

var (
	// UserLog is globally public accessible object for table cqgf_user_log operations.
	UserLog = userLogDao{
		internal.NewUserLogDao(),
	}
)

// Fill with you ideas below.

func (s *userLogDao) Create(ctx context.Context, in entity.UserLog) error {
	data := do.UserLog{}
	input, _ := json.Marshal(in)
	err := json.Unmarshal(input, &data)
	if err != nil {
		return err
	}

	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	_, err = s.Ctx(ctx).OmitEmpty().Data(data).InsertAndGetId()
	return err
}

func (s *userLogDao) GetList(ctx context.Context, in *model.UserLogApiReq) (res *response.GetListResponse, productEntity []model.UserLog, err error) {
	res = &response.GetListResponse{}
	productEntity = make([]model.UserLog, 0)
	query := s.Ctx(ctx)
	// 项目名称
	if len(in.WorkNumber) > 0 {
		query = query.Where(fmt.Sprintf("%s like ?", s.Columns().WorkNumber), g.Slice{fmt.Sprintf("%s%s", in.WorkNumber, "%")})
	}

	query, totalSize, page, size, err := util.GetListWithPage(query, in.Page, in.Size)
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