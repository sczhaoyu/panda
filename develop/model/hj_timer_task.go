package model
import (
	"time"
	"fmt"

)
type  TimerTask struct{
	Id    int64   `json:"timerTaskId" form:"timerTaskId"`    //ID
	Name    string   `json:"name" form:"name"`    //任务名称
	Ctime    time.Time   `json:"ctime" form:"ctime"`    //开始时间
	Etime    time.Time   `json:"etime" form:"etime"`    //任务结束时间
	Millisecond    int   `json:"millisecond" form:"millisecond"`    //任务开销时间毫秒
	Success    int   `json:"success" form:"success"`    //任务成功或者失败（0成功，1失败）
	Remarks    string   `json:"remarks" form:"remarks"`    //任务备注信息
 } 
//hj_timer_task返回表名称
func (t TimerTask)  TableName() string {
	return "hj_timer_task"
}
//保存单条信息
func (t *TimerTask) Save() error {
	_, err := DB.Insert(t)
	return err
	}
	
//根据ID删除
func (t *TimerTask) Delete() error {
	_, err := DB.Where("id=?",t.Id).Delete(TimerTask{})
	return err
	}
//根据ID查询
func GetTimerTaskById(id int64) (*TimerTask , error) {
	var ret TimerTask 
	b, err := DB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, NoDataMsg(b,fmt.Sprintf("TimerTask Not Found Value: %v",id))
	}
//根据ID更新数据
func (t *TimerTask) Update() error {
	_,err := DB.Where("id=?",t.Id).Update(t)
	return err
	}
  
//TimerTask查询数据分页
func FindTimerTask(page, limit int) ([]TimerTask, error) {
	var ret []TimerTask
	err := DB.Desc("id").Limit(limit, page*limit-limit).Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret,NoDataMsg(len(ret)>0,"TimerTask [] null")
}
//生成TimerTask的分页查询，WEB查询时使用
func FindTimerTaskPages(page, limit int) ([]TimerTask, int64, error) {
	var ret []TimerTask
	err := DB.Limit(limit, page*limit-limit).Find(&ret)
	if err != nil {
		return  nil,0, err
	}
	//取出数据总记录数
	count, err := DB.Count(&TimerTask{})
	if err != nil {
		return  nil,0, err
	}
	return ret, count, NoDataMsg(len(ret) > 0, "TimerTask [] null")
	}