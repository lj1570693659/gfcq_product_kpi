package response

const (
	// 用户信息相关（user）错误码1开头
	// 员工信息相关（Employee）错误码2开头
	// 部门信息相关（Department）错误码3开头
	Success            = 0
	NotSignedIn        = 101
	NotSyncEmployee    = 201
	FormatFailEmployee = 202
	CreateFailEmployee = 203
)
