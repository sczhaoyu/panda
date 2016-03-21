package model

import (
	"strconv"
)

type Table struct {
	Name    string `json:"name"`    //表名称
	Comment string `json:"comment"` //表注释
}

//查询表名称
func FindTable(DBName string, page, limit int) ([]Table, int64, error) {
	var ret []Table
	sql := "select table_name,TABLE_COMMENT from information_schema.tables where table_schema=? and table_type='base table' limit ?,?;"
	rs, err := DB.Query(sql, DBName, page*limit-limit, limit)
	if err != nil {
		return nil, 0, err
	}
	for i := 0; i < len(rs); i++ {
		var tmp Table
		tmp.Name = string(rs[i]["table_name"])
		tmp.Comment = string(rs[i]["TABLE_COMMENT"])
		ret = append(ret, tmp)
	}
	count, _ := FindTableCount(DBName)
	return ret, count, noData(len(ret) > 0)
}

//查询表名称
func FindTableCount(DBName string) (int64, error) {
	sql := "select count(*) as count from information_schema.tables where table_schema=? and table_type='base table';"
	rs, err := DB.Query(sql, DBName)
	if err != nil {
		return 0, err
	}
	if len(rs) == 0 {
		return 0, noData(len(rs) > 0)
	}
	count, _ := strconv.ParseInt(string(rs[0]["count"]), 10, 64)
	return count, nil

}
