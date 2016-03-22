package auto_code

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"huijujiayuan.com/util"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	TableSufix                       = []string{"t_s_", "hj_"} //忽略表前缀
	PackageName                      = "model"                 //包名称
	Path                             = ""                      //存储路径
	DBSrc                            = "DocDB"                 //数据库变量名称
	DBUser                           = "root"                  //数据库账号
	DBUrl                            = "10.0.0.252:3306"       //数据库连接地址
	DBPwd                            = "root"                  //数据库密码
	DB                               = "doc"                   //操作数据库
	Import         map[string]string = make(map[string]string) //需要导入的包
	IsImportDBCode                   = true                    //是否导入数据源
)

type Colunm struct {
	Type         string //数据类型
	Name         string //列名
	Comment      string //列注释
	Key          string //主键
	GoType       string //数据类型
	GoFiled      string //go程序中的字段
	IsPriKey     bool   //是否是主键
	TableName    string //所属数据库表
	GoStructName string //所属结构体名字
	GoJsonName   string //json转换后的名称
}

func filert(name string) string {

	for i := 0; i < len(TableSufix); i++ {
		name = strings.Replace(name, TableSufix[i], "", -1)
	}
	return name
}

//获取当前字段的数据类型
func (c *Colunm) GetGoDataType() {
	if c.Key == "PRI" {
		c.IsPriKey = true
	}
	if c.Key == "PRI" && c.Type == "int" {
		c.GoType = GetGoDataType(c.Type + "-" + "PRI")
	} else {
		c.GoType = GetGoDataType(c.Type)
	}
	c.GoFiled = TF(c.Name)

}

//获取表名称sql
func GetDBTableNameSql(DBName string) string {
	sql := fmt.Sprintf("select table_name,TABLE_COMMENT from information_schema.tables where table_schema='%s' and table_type='base table';", DBName)
	return sql
}

//获取表注释sql
func GetDBTableZsSql(DBName, tableName string) string {
	sql := fmt.Sprintf("select TABLE_COMMENT from information_schema.tables where table_schema='%s' and table_type='base table' and table_name='%s';", DBName, tableName)
	return sql
}

//获取列名称sql
func GetDBColunmSql(TableName string) string {
	sql := fmt.Sprintf("select * from information_schema.columns where table_schema='%s' and table_name='%s';", DB, TableName)
	return sql
}
func FindTable() {
	ret := Query(GetDBTableNameSql(DB))
	if ret == nil {
		fmt.Println("not found table!")
		return
	}
	if IsImportDBCode {
		WriteFile(Path, "db.go", DBCode())
	}
	for i := 0; i < len(ret); i++ {

		//获取表名

		GetTableInfo(string(ret[i]["table_name"]))

	}

}
func GetTableInfo(tableName string) {
	//查询表的注释
	ret := Query(GetDBTableZsSql(DB, tableName))
	zs := ""
	if ret != nil {
		for i := 0; i < len(ret); i++ {
			zs = string(ret[i]["TABLE_COMMENT"])
		}
	}
	table := CreateTable(tableName, zs, FindColunm(tableName))
	WriteFile(Path, tableName+".go", table)
	Import = make(map[string]string) //需要导入的包

}
func WriteFile(path, fileName, body string) {
	_, err := os.Stat(path)
	if err != nil {
		//创建目录
		os.MkdirAll(path, 0777)
	}
	f, err := os.Create(path + fileName)
	if err != nil {
		return
	}
	io.WriteString(f, body)
	defer f.Close()
}
func CreateColunm(c Colunm) string {
	if c.Comment == "" {
		return fmt.Sprintf("	%s    %s   `json:\"%s\"`", c.GoFiled, c.GoType, c.GoJsonName)
	}
	return fmt.Sprintf("	%s    %s   `json:\"%s\"`    //%s", c.GoFiled, c.GoType, c.GoJsonName, c.Comment)
}
func CreateTable(name, zs string, c []Colunm) string {
	srcName := name
	name = filert(name) //过滤前缀
	f := ""             //需要生成的函数
	a := `type  ` + TF(name) + ` struct{` + "\n"
	if zs != "" {
		a = "//" + zs + "\n" + a
	}
	bl := false
	for i := 0; i < len(c); i++ {
		a = a + CreateColunm(c[i]) + "\n"
		if c[i].IsPriKey && bl == false {
			//获取删除函数
			f = f + DeleteFunc(c[i].GoStructName, c[i].Name, c[i].GoFiled) + "\n"
			//获取ID查询
			f = f + PrikeyFunc(c[i].GoStructName, c[i].GoJsonName, c[i].Name, c[i].GoType) + "\n"
			f = f + UpdateFunc(c[i].GoStructName, c[i].GoFiled, c[i].Name) + "\n"
			f = f + FindFunc(c[i].GoStructName, c[i].Name)
			f = f + FindFuncPages(c[i].GoStructName)
			bl = true
		}
	}

	b := ` } `
	//获取表名
	tn := a + b + "\n" + CrateTableFunc(TF(name), srcName)
	//获取创建函数
	tn = tn + "\n" + CreateFunc(TF(name))

	//导入包

	return "package " + PackageName + "\n" + GetImport() + tn + "\n" + f
}
func CrateTableFunc(name, srcName string) string {
	if name == TF(srcName) {
		return ""
	}
	s := "//" + srcName + "返回表名称\n"
	s += "func (t %s)  TableName() string {\n	return \"%s\"\n}"
	return fmt.Sprintf(s, name, srcName)
}
func FindColunm(TableName string) []Colunm {

	m := Query(GetDBColunmSql(TableName))
	if m == nil {
		fmt.Println("not found column!")
		return nil
	}

	TableName = filert(TableName) //过滤前缀

	var ret []Colunm = make([]Colunm, 0, 200)
	for i := 0; i < len(m); i++ {
		var tmp Colunm
		tmp.Type = string(m[i]["DATA_TYPE"])         //数据类型
		tmp.Name = string(m[i]["COLUMN_NAME"])       //列名
		tmp.Comment = string(m[i]["COLUMN_COMMENT"]) //列注释
		tmp.Key = string(m[i]["COLUMN_KEY"])         //主键
		tmp.TableName = TableName                    //表名
		tmp.GoStructName = TF(TableName)             //在结构中的名字
		tmp.GetGoDataType()                          //装配数据类型
		tmp.GoJsonName = PSK(tmp.Name)

		ret = append(ret, tmp)
	}
	return ret
}

func GetGoDataType(t string) string {
	tp := ""
	switch t {
	case "varchar", "char", "text", "longtext", "tinyblob", "longblob", "blob":
		tp = "string"
	case "int-PRI":
		tp = "int64"
	case "tinyint", "int", "smallint":
		tp = "int"
	case "date", "timestamp", "datetime":
		tp = "time.Time"
		Import["time"] = ""
	case "decimal", "float", "numeric", "double":
		tp = "float64"
	}
	return tp

}

//名字的驼峰命名
func TF(name string) string {
	s := strings.Split(name, "_")
	ret := ""
	for i := 0; i < len(s); i++ {

		ret = ret + strings.ToUpper(util.SubString(s[i], 0, 1)) + util.SubString(s[i], 1, len(s[i]))
	}
	return ret
}

//名字帕斯卡命名
func PSK(name string) string {

	ret := TF(name)
	if isOk, _ := regexp.MatchString("^[A-Z]+$", ret); isOk {
		return strings.ToLower(ret)
	}

	return strings.ToLower(util.SubString(ret, 0, 1)) + util.SubString(ret, 1, len(ret))
}

//创建函数
func CreateFunc(name string) string {
	fun := "//保存单条信息\n"
	fun += `func (%s *%s) Save() error {
	_, err := %s.Insert(%s)
	return err
	}
	`
	one := strings.ToLower(util.SubString(name, 0, 1))
	return fmt.Sprintf(fun, one, name, DBSrc, one)
}

//生成导入包的代码
func GetImport() string {
	ret := "import (\n%s\n)\n"
	tmp := ""
	for k, _ := range Import {
		tmp = tmp + "\t\"" + k + "\"" + "\n"
	}
	if tmp == "" {
		return ""
	}

	return fmt.Sprintf(ret, tmp)
}

//生成删除的字段
func DeleteFunc(name string, filed, goFiled string) string {
	fun := "//根据ID删除\n"
	fun += `func (%s *%s) Delete() error {
	_, err := %s.Where("%s=?",%s.%s).Delete(%s{})
	return err
	}`
	one := strings.ToLower(util.SubString(name, 0, 1))
	return fmt.Sprintf(fun, one, name, DBSrc, filed, one, goFiled, name)
}

//生成ID查询
func PrikeyFunc(name, PskgoFiled, DBFiled, dataType string) string {
	t := "fmt.Sprintf(\"" + name + " Not Found Value: %v\"," + PskgoFiled + ")"
	fun := "//根据ID查询\n"
	fun += `func Get` + name + `By%s(%s %s) (*` + name + ` , error) {
	var ret ` + name + ` 
	b, err := %s.Where("%s=?", ` + PskgoFiled + `).Get(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, NoDataMsg(b,%s)
	}`
	Import["fmt"] = ""
	return fmt.Sprintf(fun, TF(PskgoFiled), PskgoFiled, dataType, DBSrc, DBFiled, t)
}

//获取更新函数
func UpdateFunc(name, goFiled, dbFiled string) string {
	fun := "//根据ID更新数据\n"
	fun += `func (%s *%s) Update() error {
	_,err := %s.Where("%s=?",%s.%s).Update(%s)
	return err
	}
  `
	one := strings.ToLower(util.SubString(name, 0, 1))
	return fmt.Sprintf(fun, one, name, DBSrc, dbFiled, one, goFiled, one)
}
func FindFunc(name, dbFiled string) string {
	fun := "//" + name + "查询数据分页\n"
	fun += `func Find` + name + `(page, limit int) ([]` + name + `, error) {
	var ret []` + name + `
	err := ` + DBSrc + `.Desc("` + dbFiled + `").Limit(limit, page*limit-limit).Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret,NoDataMsg(len(ret)>0,"` + name + ` [] null")
}`
	return fun
}
func FindFuncPages(name string) string {
	fun := "\n//生成" + name + "的分页查询，WEB查询时使用\n"
	fun += `func Find` + name + `Pages(page, limit int) ([]` + name + `, int64, error) {
	var ret []` + name + `
	err := ` + DBSrc + `.Limit(limit, page*limit-limit).Find(&ret)
	if err != nil {
		return  nil,0, err
	}
	//取出数据总记录数
	count, err := DB.Count(&` + name + `{})
	if err != nil {
		return  nil,0, err
	}
	return ret, count, NoDataMsg(len(ret) > 0, "` + name + ` [] null")
	}`
	return fun
}

//生成数据库代码
func DBCode() string {
	d := `package ` + PackageName + `
import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"os"
)

const (
	SQL_NUM     = 150 //SQL批处理条数
	MAX_CLIENT  = 400 //最大链接个数
	INIT_CLIENT = 10  //初始化链接个数
)

var (
	` + DBSrc + ` *xorm.Engine //数据库
	 
)

func init() {
	//====================================================================

	url := "` + DBUser + `:` + DBPwd + `@tcp(` + DBUrl + `)/"
	 
	` + DBSrc + `, _ = xorm.NewEngine("mysql", url+"` + DB + `?charset=utf8")
	if os.Getenv("GO_DEV") == "1" {
		` + DBSrc + `.ShowSQL = true
	 
	}
	` + DBSrc + `.SetMaxIdleConns(INIT_CLIENT)
	` + DBSrc + `.SetMaxOpenConns(MAX_CLIENT)
	//====================================================================

}
func NoData(b bool) error {

	if b {
		return nil
	}
	return errors.New("empty")
}

//错误消息定义
func NoDataMsg(b bool, msg string) error {
	if b {
		return nil
	}
	return errors.New(msg)
}
`
	return d
}
func Query(sqL string) []map[string]string {
	var ret []map[string]string = make([]map[string]string, 0, 500)
	db, err := sql.Open("mysql", DBUser+":"+DBPwd+"@tcp("+DBUrl+")/"+DB+"?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	rows, err := db.Query(sqL)
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		ret = append(ret, record)
	}
	defer db.Close()
	return ret
}
