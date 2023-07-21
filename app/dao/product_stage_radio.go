// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/dao/internal"
)

// internalProductStageRadioDao is internal type for wrapping internal DAO implements.
type internalProductStageRadioDao = *internal.ProductStageRadioDao

// productStageRadioDao is the data access object for table cqgf_product_stage_radio.
// You can define custom methods on it to extend its functionality as you wish.
type productStageRadioDao struct {
	internalProductStageRadioDao
}

var (
	// ProductStageRadio is globally public accessible object for table cqgf_product_stage_radio operations.
	ProductStageRadio = productStageRadioDao{
		internal.NewProductStageRadioDao(),
	}
)

// Fill with you ideas below.
