package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"go-vue-admin/api/v1/common"
	"go-vue-admin/internal/app/system/model/entity"
)

type GetSysConfigListReq struct {
	g.Meta     `path:"/system/sys/config/list" method:"Get" tags:"Get Get sys_config" summary:"Get Get sys_config"`
	ConfigName string      `p:"configName"`
	ConfigKey  string      `p:"configKey"`
	ConfigType string      `p:"configType"`
	CreateTime *gtime.Time `p:"createTime"`
	common.PageReq
}

type GetSysConfigListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*entity.SysConfig `json:"rows"`
	Total  int                 `json:"total"`
}

type PostSysConfigReq struct {
	g.Meta      `path:"/system/sys/config" method:"Post" tags:"Post sys_config" summary:"Post sys_config"`
	ConfigName  string `p:"configName" `
	ConfigKey   string `p:"configKey" `
	ConfigValue string `p:"configValue" `
	ConfigType  string `p:"configType" `
	Remark      string `p:"remark" `
}

type PostSysConfigRes struct {
	g.Meta `mime:"application/json"`
}

type PutSysConfigReq struct {
	g.Meta      `path:"/system/sys/config" method:"Put" tags:"Put sys_config" summary:"Put sys_config"`
	ConfigId    int    `p:"configId"  v:"required" `
	ConfigName  string `p:"configName" `
	ConfigKey   string `p:"configKey" `
	ConfigValue string `p:"configValue" `
	ConfigType  string `p:"configType" `
	Remark      string `p:"remark" `
}

type PutSysConfigRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteSysConfigReq struct {
	g.Meta   `path:"/system/sys/config/{configId}" method:"Delete" tags:"Delete sys_config" summary:"Delete sys_config"`
	ConfigId string `p:"configId"  v:"required" `
}

type DeleteSysConfigRes struct {
	g.Meta `mime:"application/json"`
}

type GetSysConfigReq struct {
	g.Meta   `path:"/system/sys/config/{configId}" method:"Get" tags:"Get sys_config" summary:"Get sys_config"`
	ConfigId int `p:"configId"  v:"required" `
}

type GetSysConfigRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysConfig
}
