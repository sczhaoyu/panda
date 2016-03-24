package model

import (
	"github.com/sczhaoyu/panda/develop/config"
	"strconv"
)

type Table struct {
	Name    string `json:"name"`    //表名称
	Comment string `json:"comment"` //表注释
}

//查询表名称
func FindTable(page, limit int) ([]Table, int64, error) {
	var ret []Table
	sql := "select table_name,TABLE_COMMENT from information_schema.tables where table_schema=? and table_type='base table' limit ?,?;"
	rs, err := DB.Query(sql, config.DB("db").String(), page*limit-limit, limit)
	if err != nil {
		return nil, 0, err
	}
	for i := 0; i < len(rs); i++ {
		var tmp Table
		tmp.Name = string(rs[i]["table_name"])
		tmp.Comment = string(rs[i]["TABLE_COMMENT"])
		ret = append(ret, tmp)
	}
	count, _ := FindTableCount()
	return ret, count, NoData(len(ret) > 0)
}

//查询表名称
func FindTableCount() (int64, error) {
	sql := "select  count(*) as count from information_schema.tables where table_schema=? and table_type='base table' ;"
	rs, err := DB.Query(sql, config.DB("db").String())
	if err != nil {
		return 0, err
	}
	if len(rs) == 0 {
		return 0, NoData(len(rs) > 0)
	}
	count, _ := strconv.ParseInt(string(rs[0]["count"]), 10, 64)
	return count, nil

}
func GetTable(name string) (*Table, error) {
	var ret Table
	sql := "select * from information_schema.tables where table_schema=? and table_type='base table' and table_name=?;"
	rs, err := DB.Query(sql, config.DB("db").String(), name)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 1; i++ {
		ret.Name = string(rs[0]["table_name"])
		ret.Comment = string(rs[0]["TABLE_COMMENT"])
	}

	return &ret, NoData(ret.Name != "")
}
