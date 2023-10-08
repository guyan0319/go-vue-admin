package consts

const (
	ErrLoginFailMsg           = "用户名或密码错误"
	ErrLoginCodeFailMsg       = "验证码错误"
	ErrAuthFailMsg            = "没有授权或请求超时"
	Success                   = "操作成功"
	AdminAllPer               = "*:*:*"
	ProAdmin                  = "admin"
	ProAdminId          int64 = 1
	ProAdminRoleId      int64 = 1

	//ctx
	CtxAdminId   = "CtxAdminId"
	CtxAdminName = "CtxAdminName"

	//sys_menu
	SysMenuStatusOk = "0"
	SysMenuStatusNo = "1"

	CacheKeyPermsUrl = "redis.perms.url"

	//sys_role
	SysRoleStatusOk = "0"
	SysRoleStatusNo = "1"

	SysRoleDataScopeAll         = "1"
	SysRoleDataScopeCustom      = "2"
	SysRoleDataScopeCurrent     = "3"
	SysRoleDataScopeCurrents    = "4"
	SysRoleDataScopeCurrentUser = "5"

	//sys_post
	SysPostStatusOk = "0"
	SysPostStatusNo = "1"

	//sys_user
	SysUserStatusOk = "0"
	SysUserStatusNo = "1"

	//gen_table
	GenTableStatusOk = "0"
	GenTableStatusNo = "1"

	//sys_dept
	GetAllDeptListCache = "GetAllDeptListCache"
	SysDeptStatusOk     = "0"
	SysDeptStatusNo     = "1"

	//page
	PageSize = 10

	OpenRedis = true
	//必须跟hack/config.yaml配置一直，默认false，
	//当数据表字段类型为时间类型时，代码生成的属性类型使用标准库的time.Time而不是框架的*gtime.Time类型。
	StdTime = false
	//当数据表字段类型为JSON类型时，代码生成的属性类型使用*gjson.Json类型。
	GJsonSupport = false

	//gen code
	//api调用路径
	ApiUrlPath     = "/system/"
	ApiPackageName = "v1"
	//模板存放位置
	TemplateDir = "default/"

	//存放生成代码目录
	GenCodeDir = "resource/temp/down/"
	//生成代码压缩包
	GenCodeZipDir = "resource/temp/zip/"
	//预览代码
	GenCodeViewDir = "resource/temp/view/"
	//采用restfull
	ListMethod   = "Get"
	AddMethod    = "Post"
	UpdateMethod = "Put"
	DeleteMethod = "Delete"
)
