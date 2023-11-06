package model

// GetCheckIn 查询成员打卡日报
type GetCheckIn struct {
	StartTime  string   `json:"startTime"` // 部门ID
	EndTime    string   `json:"endTime"`   // 部门ID
	DepartId   int32    `json:"departId"`  // 部门ID
	UseridList []string `json:"useridList"`
	ProId      int32    `json:"proId"`
}
