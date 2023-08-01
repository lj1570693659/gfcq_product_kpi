package util

import (
	"context"
	"fmt"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/consts"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	inspirit "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"strconv"
	"strings"
)

const (
	GSHA1 = "gsha1"
	MD5   = "md5"
)

func GetListWithPage(query *gdb.Model, page, size int32) (*gdb.Model, int32, int32, int32, error) {
	if g.IsEmpty(page) {
		page = 1
	}
	if g.IsEmpty(size) {
		size = 10
	}
	totalSize, err := query.Count()
	if err != nil {
		return query, gconv.Int32(totalSize), page, size, err
	}

	query = query.Limit(gconv.Int((page-1)*size), gconv.Int(size))
	return query, gconv.Int32(totalSize), page, size, nil
}

func Encrypt(str string) string {
	var encryptStr string
	types, _ := g.Config("config.toml").Get(context.Background(), "user.encrypt")
	switch types.String() {
	case GSHA1:
		encryptStr = gsha1.Encrypt(str)
	case MD5:
		encryptStr, _ = gmd5.Encrypt(str)
	}
	return encryptStr
}

func DeleteIntSlice(a []string) []string {
	ret := make([]string, 0, len(a))
	for _, val := range a {
		if !g.IsEmpty(val) {
			ret = append(ret, val)
		}
	}
	return ret
}

func GetEmploySex(name v1.SexEnum) string {
	attributeName := map[v1.SexEnum]string{
		v1.SexEnum_unknow: "未知",
		v1.SexEnum_man:    "男",
		v1.SexEnum_woman:  "女",
	}
	return attributeName[name]
}

// GetEmployStatus  '在职状态（1：在职 2：试用期 3：实习期 4：已离职）',
func GetEmployStatus(name v1.StatusEnum) string {
	attributeName := map[v1.StatusEnum]string{
		v1.StatusEnum_anything:   "未知",
		v1.StatusEnum_working:    "在职",
		v1.StatusEnum_tryout:     "试用期",
		v1.StatusEnum_interns:    "实习期",
		v1.StatusEnum_terminated: "已离职",
	}
	return attributeName[name]
}

func GetEmployAttribute(name string) uint {
	attributeName := map[string]uint{
		"兼职": consts.PartTime,
		"全职": consts.FullTime,
	}
	return attributeName[name]
}

func GetIsGuide(name uint) string {
	isGuideName := map[uint]string{
		0: "否",
		1: "是",
	}
	return isGuideName[name]
}

func GetFloatKeyType(name string) uint {
	devote := consts.ElseDevote
	//1：加班贡献 2：解决问题贡献 3：其他事件贡献
	if strings.Contains(name, "加班") {
		devote = consts.OverTimeDevote
	} else if strings.Contains(name, "解决问题") {
		devote = consts.SolveProblemDevote
	}
	return gconv.Uint(devote)
}

// GetFloatKeyProperty 1：正向激励 2：有待提高
func GetFloatKeyProperty(floatRaio float64) uint {
	var devote uint
	if floatRaio > 0 {
		devote = consts.ForwardDirection
	} else {
		devote = consts.ReverseDirection
	}
	return devote
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func DecimalLong(value float64, len int) float64 {
	formatStr := fmt.Sprintf("%%.%df", len)
	value, _ = strconv.ParseFloat(fmt.Sprintf(formatStr, value), 64)
	return value
}

// Letter 遍历a-z
func Letter(length int) []string {
	var str []string
	for i := 0; i < length; i++ {
		str = append(str, string(rune('A'+i)))
	}
	return str
}

// GetTreeNode 递归获取子节点
func GetTreeNode(ctx context.Context, perms []model.DepartmentApiGetList, GroupBy, GetFiledNameCount string) (context.Context, []model.DepartmentApiGetList, string, string) {
	//定义子节点
	for k, v := range perms {
		// 计算直属上级部门员工数量
		var childCountSum int32
		getCount, err := boot.EmployeeJobServer.GetCount(ctx, &v1.GetCountEmployeeJobReq{
			EmployeeJob: &v1.EmployeeJobInfo{
				DepartId: gconv.Int32(v.ID),
			},
			GroupBy:           GroupBy,
			GetFiledNameCount: GetFiledNameCount,
		})
		if err != nil {
			return ctx, perms, GroupBy, GetFiledNameCount
		}

		// 计算下级部门
		getChild, err := boot.DepertmentServer.GetListWithoutPage(ctx, &v1.GetListWithoutDepartmentReq{
			Department: &v1.DepartmentInfo{
				Pid: gconv.Int32(v.ID),
			},
		})
		if err != nil {
			return ctx, perms, GroupBy, GetFiledNameCount
		}
		info := make([]model.DepartmentApiGetList, 0)
		gconv.Scan(getChild.GetData(), &info)
		perms[k].ChildDepart = info

		if len(info) > 0 {
			for ik, iv := range info {
				getCount, err := boot.EmployeeJobServer.GetCount(ctx, &v1.GetCountEmployeeJobReq{
					EmployeeJob: &v1.EmployeeJobInfo{
						DepartId: gconv.Int32(iv.ID),
					},
					GroupBy:           GroupBy,
					GetFiledNameCount: GetFiledNameCount,
				})
				if err != nil {
					return ctx, perms, GroupBy, GetFiledNameCount
				}
				info[ik].EmployeeCount = getCount.GetCount()
				childCountSum += getCount.GetCount()
			}
		}

		perms[k].EmployeeCount = getCount.GetCount() + childCountSum
		GetTreeNode(ctx, info, GroupBy, GetFiledNameCount)
	}
	return ctx, perms, GroupBy, GetFiledNameCount
}

func GetHoursIndexByScore(lists []*inspirit.CrewHoursIndexInfo, score float32) uint32 {
	for _, v := range lists {
		switch v.ScoreRange {
		case consts.ScoreRangeMin:
			// 左闭右开
			if v.ScoreMin <= score && score < v.ScoreMax {
				return v.ScoreIndex
			}
		case consts.ScoreRangeMax:
			// 左开右闭
			if v.ScoreMin < score && score <= v.ScoreMax {
				return v.ScoreIndex
			}
		case consts.ScoreRangeMinAndMax:
			// 左闭右闭
			if v.ScoreMin <= score && score <= v.ScoreMax {
				return v.ScoreIndex
			}
		}
	}
	return 0
}
