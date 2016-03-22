package controller

import (
	"math"
)

type PaginationHelper struct {
	Count       int `json:"count"`       ///总数据条数
	CurrentPage int `json:"currentPage"` //当前页面
	TotalPage   int `json:"totalPage"`   //总页数
	ShowNum     int `json:"showNum"`     //显示条数
	StartNum    int `json:"startNum"`    //开始记录
	Endnum      int `json:"endnum"`      //截止记录
}

type Pagination struct {
	Count       int         `json:"total"`       ///总数据条数
	CurrentPage int         `json:"currentPage"` //当前页面
	TotalPage   int         `json:"totalPage"`   //总页数
	ShowNum     int         `json:"showNum"`     //显示条数
	StartNum    int         `json:"startNum"`    //开始记录
	Endnum      int         `json:"endnum"`      //截止记录
	Ret         interface{} `json:"rows"`        //分页后的数据  数组或者对象
}

//分页计算 总数据条数，当前页，显示多少条
func GetPagination(c int64, currentPage, pageSize int, ret interface{}) *Pagination {
	count := int(c)
	var p Pagination
	p.Count = count //总数据条数
	if pageSize > count || currentPage < 0 {
		p.ShowNum = count //每页显示条数
	} else if pageSize <= 0 {
		p.ShowNum = count
	} else {
		p.ShowNum = pageSize //每页显示条数
	}
	///总页数  处理
	if p.ShowNum == 0 {
		p.TotalPage = 1
	} else {
		c := float64(p.Count)
		s := float64(p.ShowNum)
		p.TotalPage = int(math.Ceil(c / s))
	}

	if currentPage >= p.TotalPage {
		p.CurrentPage = p.TotalPage
	} else if p.CurrentPage < 0 {
		p.CurrentPage = 1
	} else {
		p.CurrentPage = currentPage
	}

	///处理结束位置
	if currentPage*pageSize >= count {
		p.Endnum = count
	} else {
		p.Endnum = p.CurrentPage * p.ShowNum
	}
	///处理开始位置

	p.StartNum = p.CurrentPage*p.ShowNum - p.ShowNum
	p.Ret = ret
	return &p
}

//加入对象格式化
func (this *Pagination) Put(ret interface{}) {
	this.Ret = ret
}
