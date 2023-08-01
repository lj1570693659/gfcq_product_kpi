// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/dao/internal"
)

// internalCrewDutyIndexDao is internal type for wrapping internal DAO implements.
type internalCrewDutyIndexDao = *internal.CrewDutyIndexDao

// crewDutyIndexDao is the data access object for table cqgf_crew_duty_index.
// You can define custom methods on it to extend its functionality as you wish.
type crewDutyIndexDao struct {
	internalCrewDutyIndexDao
}

var (
	// CrewDutyIndex is globally public accessible object for table cqgf_crew_duty_index operations.
	CrewDutyIndex = crewDutyIndexDao{
		internal.NewCrewDutyIndexDao(),
	}
)

// Fill with you ideas below.