// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/dao/internal"
)

// internalProductLevelConfirmDao is internal type for wrapping internal DAO implements.
type internalProductLevelConfirmDao = *internal.ProductLevelConfirmDao

// productLevelConfirmDao is the data access object for table cqgf_product_level_confirm.
// You can define custom methods on it to extend its functionality as you wish.
type productLevelConfirmDao struct {
	internalProductLevelConfirmDao
}

var (
	// ProductLevelConfirm is globally public accessible object for table cqgf_product_level_confirm operations.
	ProductLevelConfirm = productLevelConfirmDao{
		internal.NewProductLevelConfirmDao(),
	}
)

// Fill with you ideas below.
