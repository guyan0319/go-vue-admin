// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysJobLog is the golang structure of table sys_job_log for DAO operations like Where/Data.
type SysJobLog struct {
	g.Meta        `orm:"table:sys_job_log, do:true"`
	JobLogId      interface{} // 任务日志ID
	JobName       interface{} // 任务名称
	JobGroup      interface{} // 任务组名
	InvokeTarget  interface{} // 调用目标字符串
	JobMessage    interface{} // 日志信息
	Status        interface{} // 执行状态（0正常 1失败）
	ExceptionInfo interface{} // 异常信息
	CreateTime    *gtime.Time // 创建时间
}
