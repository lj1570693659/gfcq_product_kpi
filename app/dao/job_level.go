// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/dao/internal"
)

// internalJobLevelDao is internal type for wrapping internal DAO implements.
type internalJobLevelDao = *internal.JobLevelDao

// jobLevelDao is the data access object for table cqgf_job_level.
// You can define custom methods on it to extend its functionality as you wish.
type jobLevelDao struct {
	internalJobLevelDao
}

var (
	// JobLevel is globally public accessible object for table cqgf_job_level operations.
	JobLevel = jobLevelDao{
		internal.NewJobLevelDao(),
	}
)

// Fill with you ideas below.
