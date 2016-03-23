package model
import (
	"fmt"

)
//短信验证码
type  AuthCode struct{
	Id    int64   `json:"authCodeId" form:"authCodeId"`    //记录ID
	Mobile    string   `json:"mobile" form:"mobile"`    //手机号
	AuthType    int   `json:"authType" form:"authType"`    //验证码类型
	AuthCode    string   `json:"authCode" form:"authCode"`    //验证码内容
	Ctime    int   `json:"ctime" form:"ctime"`    //创建时间戳
	Etime    int   `json:"etime" form:"etime"`    //失效时间戳
 } 
//hj_auth_code返回表名称
func (t AuthCode)  TableName() string {
	return "hj_auth_code"
}
//保存单条信息
func (a *AuthCode) Save() error {
	_, err := DB.Insert(a)
	return err
	}
	
//根据ID删除
func (a *AuthCode) Delete() error {
	_, err := DB.Where("id=?",a.Id).Delete(AuthCode{})
	return err
	}
//根据ID查询
func GetAuthCodeById(id int64) (*AuthCode , error) {
	var ret AuthCode 
	b, err := DB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, NoDataMsg(b,fmt.Sprintf("AuthCode Not Found Value: %v",id))
	}
//根据ID更新数据
func (a *AuthCode) Update() error {
	_,err := DB.Where("id=?",a.Id).Update(a)
	return err
	}
  
//AuthCode查询数据分页
func FindAuthCode(page, limit int) ([]AuthCode, error) {
	var ret []AuthCode
	err := DB.Desc("id").Limit(limit, page*limit-limit).Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret,NoDataMsg(len(ret)>0,"AuthCode [] null")
}
//生成AuthCode的分页查询，WEB查询时使用
func FindAuthCodePages(page, limit int) ([]AuthCode, int64, error) {
	var ret []AuthCode
	err := DB.Limit(limit, page*limit-limit).Find(&ret)
	if err != nil {
		return  nil,0, err
	}
	//取出数据总记录数
	count, err := DB.Count(&AuthCode{})
	if err != nil {
		return  nil,0, err
	}
	return ret, count, NoDataMsg(len(ret) > 0, "AuthCode [] null")
	}