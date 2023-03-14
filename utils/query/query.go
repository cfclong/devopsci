package query

import (
	"fmt"
	"math"

	"github.com/astaxie/beego/orm"
)

type FilterQuery struct {
	PageIndex int    `json:"page_index"`
	PageSize  int    `json:"page_size"`
	FilterKey string `json:"filter_key"`
	FilterVal string `json:"filter_val"`
	IsLike    bool   `json:"-"`
}

type PageInfo struct {
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
	Page    int `json:"page"`
	Pages   int `json:"pages"`
}

// QueryResult ..
type QueryResult struct {
	PageInfo
	Item interface{} `json:"item"`
}

// FillPageInfo ..
func FillPageInfo(queryResult *QueryResult, pageIndex int, pageSize int, totalNum int) error {
	queryResult.PageInfo.Page = pageIndex
	queryResult.PageInfo.PerPage = pageSize
	queryResult.PageInfo.Total = totalNum

	queryResult.PageInfo.Pages = int(math.Ceil(float64(queryResult.PageInfo.Total) / float64(queryResult.PageInfo.PerPage)))

	return nil
}

// FilterCondition ..
func FilterCondition(filter *FilterQuery, filterKeys string) *orm.Condition {
	cond := orm.NewCondition()
	if filter.FilterVal != "" {
		cond = cond.Or(fmt.Sprintf("%s__icontains", filterKeys), filter.FilterVal)
	}
	if cond.IsEmpty() {
		return nil
	}
	return cond
}

func NewFilterQuery(isLike bool) *FilterQuery {
	return &FilterQuery{IsLike: isLike}
}
