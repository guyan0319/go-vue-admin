// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-vue-admin/internal/app/system/dao/internal"
)

// internalSysLogininforDao is internal type for wrapping internal DAO implements.
type internalSysLogininforDao = *internal.SysLogininforDao

// sysLogininforDao is the data access object for table sys_logininfor.
// You can define custom methods on it to extend its functionality as you wish.
type sysLogininforDao struct {
	internalSysLogininforDao
}

var (
	// SysLogininfor is globally public accessible object for table sys_logininfor operations.
	SysLogininfor = sysLogininforDao{
		internal.NewSysLogininforDao(),
	}
)

// Fill with you ideas below.