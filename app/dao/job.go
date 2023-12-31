// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/dao/internal"
)

// internalJobDao is internal type for wrapping internal DAO implements.
type internalJobDao = *internal.JobDao

// jobDao is the data access object for table cqgf_job.
// You can define custom methods on it to extend its functionality as you wish.
type jobDao struct {
	internalJobDao
}

var (
	// Job is globally public accessible object for table cqgf_job operations.
	Job = jobDao{
		internal.NewJobDao(),
	}
)

// Fill with you ideas below.
