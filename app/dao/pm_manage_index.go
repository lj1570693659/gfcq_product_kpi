// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/dao/internal"
)

// internalPmManageIndexDao is internal type for wrapping internal DAO implements.
type internalPmManageIndexDao = *internal.PmManageIndexDao

// pmManageIndexDao is the data access object for table cqgf_pm_manage_index.
// You can define custom methods on it to extend its functionality as you wish.
type pmManageIndexDao struct {
	internalPmManageIndexDao
}

var (
	// PmManageIndex is globally public accessible object for table cqgf_pm_manage_index operations.
	PmManageIndex = pmManageIndexDao{
		internal.NewPmManageIndexDao(),
	}
)

// Fill with you ideas below.
