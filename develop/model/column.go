package model

import (
	"github.com/sczhaoyu/panda/develop/config"
)

type Column struct {
	TableName string `json:"tableName"` //表名称
	DataType  string `json:"dataType"`  //数据类型
	Comment   string `json:"comment"`   //表注释
	Key       string `json:"key"`       //主键类型
	Name      string `json:"name"`      //列名
}

//获取列名称sql
func FindColumns(tableName string) ([]Column, error) {
	var ret []Column
	sql := "select TABLE_NAME,COLUMN_COMMENT,COLUMN_KEY,DATA_TYPE,COLUMN_NAME from information_schema.columns where table_schema=? and table_name=?;"
	rs, err := DB.Query(sql, config.DB("db").String(), tableName)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(rs); i++ {
		var tmp Column
		tmp.TableName = string(rs[i]["TABLE_NAME"])
		tmp.Comment = string(rs[i]["COLUMN_COMMENT"])
		tmp.Key = string(rs[i]["COLUMN_KEY"])
		tmp.DataType = string(rs[i]["DATA_TYPE"])
		tmp.Name = string(rs[i]["COLUMN_NAME"])
		ret = append(ret, tmp)
	}

	return ret, NoData(len(ret) > 0)
}
