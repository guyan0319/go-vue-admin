// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysJobLog is the golang structure for table sys_job_log.
type SysJobLog struct {
	JobLogId      int64       `json:"jobLogId"      description:"任务日志ID"`
	JobName       string      `json:"jobName"       description:"任务名称"`
	JobGroup      string      `json:"jobGroup"      description:"任务组名"`
	InvokeTarget  string      `json:"invokeTarget"  description:"调用目标字符串"`
	JobMessage    string      `json:"jobMessage"    description:"日志信息"`
	Status        string      `json:"status"        description:"执行状态（0正常 1失败）"`
	ExceptionInfo string      `json:"exceptionInfo" description:"异常信息"`
	CreateTime    *gtime.Time `json:"createTime"    description:"创建时间"`
}
