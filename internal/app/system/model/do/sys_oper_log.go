// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOperLog is the golang structure of table sys_oper_log for DAO operations like Where/Data.
type SysOperLog struct {
	g.Meta        `orm:"table:sys_oper_log, do:true"`
	OperId        interface{} // 日志主键
	Title         interface{} // 模块标题
	BusinessType  interface{} // 业务类型（0其它 1新增 2修改 3删除）
	Method        interface{} // 方法名称
	RequestMethod interface{} // 请求方式
	OperatorType  interface{} // 操作类别（0其它 1后台用户 2手机端用户）
	OperName      interface{} // 操作人员
	DeptName      interface{} // 部门名称
	OperUrl       interface{} // 请求URL
	OperIp        interface{} // 主机地址
	OperLocation  interface{} // 操作地点
	OperParam     interface{} // 请求参数
	JsonResult    interface{} // 返回参数
	Status        interface{} // 操作状态（0正常 1异常）
	ErrorMsg      interface{} // 错误消息
	OperTime      *gtime.Time // 操作时间
	CostTime      interface{} // 消耗时间
}
