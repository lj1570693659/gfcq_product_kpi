// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/dao/internal"
)

// internalPmHoursIndexDao is internal type for wrapping internal DAO implements.
type internalPmHoursIndexDao = *internal.PmHoursIndexDao

// pmHoursIndexDao is the data access object for table cqgf_pm_hours_index.
// You can define custom methods on it to extend its functionality as you wish.
type pmHoursIndexDao struct {
	internalPmHoursIndexDao
}

var (
	// PmHoursIndex is globally public accessible object for table cqgf_pm_hours_index operations.
	PmHoursIndex = pmHoursIndexDao{
		internal.NewPmHoursIndexDao(),
	}
)

// Fill with you ideas below.
