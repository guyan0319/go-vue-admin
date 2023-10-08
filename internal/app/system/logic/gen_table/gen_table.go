package gen_table

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gcompress"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/consts"
	"go-vue-admin/internal/app/system/dao"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/model/do"
	"go-vue-admin/internal/app/system/model/entity"
	"go-vue-admin/internal/app/system/service"
	"go-vue-admin/utility"
	"go-vue-admin/utility/lib"
	"strconv"
	"strings"
)

func init() {
	service.RegisterGenTable(New())
}

func New() *sGenTable {
	return &sGenTable{}
}

type sGenTable struct {
	IsPreview bool
}
type ApiFunc struct {
	FuncName string
	Req      string
	Res      string
	Url      string
}

var apiMap = map[string]ApiFunc{}

var apiJsMap = map[string]string{}
var fileCodeMap = map[string]string{}

func (s *sGenTable) GetTableList(ctx context.Context, req v1.ToolListCommon) (tableList v1.ToolListCommonRes, err error) {
	tableList = v1.ToolListCommonRes{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}

	err = g.Try(ctx, func(ctx context.Context) {
		var tables []*entity.GenTable
		m := dao.GenTable.Ctx(ctx)
		//状态
		if req.Status != "" {
			m = m.Where(dao.GenTable.Columns().Status, req.Status)
		}
		//时间
		if len(req.Params) > 0 {
			m = m.WhereBetween(dao.GenTable.Columns().CreateTime, req.Params["beginTime"], req.Params["endTime"])
		}

		//表名
		if req.TableName != "" {
			m = m.WhereLike(dao.GenTable.Columns().TableName, "%"+req.TableName+"%")
		}
		//表描述
		if req.TableComment != "" {
			m = m.WhereLike(dao.GenTable.Columns().TableComment, "%"+req.TableComment+"%")
		}

		tableList.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&tables)
		tableRows := make([]*model.GenTableList, len(tables))
		for k, value := range tables {
			ul := &model.GenTableList{}
			ul.GenTable = value
			tableRows[k] = ul
		}
		tableList.Rows = tableRows
	})
	return
}

func (s *sGenTable) GetGenTablesUpdate(ctx context.Context, req *v1.GetGenTablesUpdateReq) (tableList *v1.GetGenTablesUpdateRes, err error) {
	tableList = &v1.GetGenTablesUpdateRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		var tables *entity.GenTable
		err := dao.GenTable.Ctx(ctx).Where(dao.GenTable.Columns().TableId, req.TableId).Scan(&tables)
		utility.WriteErrLog(ctx, err, "获取生成表数据失败")
		info := &model.GenTableInfo{}
		info.GenTable = tables
		//获取生成表字段数据
		tableList.Rows, err = service.GenTableColumn().GetTableColumnByTableId(ctx, req.TableId)
		if err != nil {
			return
		}
		info.Columns = tableList.Rows
		tableList.Info = info
	})
	return
}

func (s *sGenTable) Delete(ctx context.Context, req *v1.DeleteGenTableReq) (tableList *v1.DeleteGenTableRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		tableIds := lib.ParamStrToSlice(req.TableId, ",")
		//删除生成表信息
		_, e := dao.GenTable.Ctx(ctx).WhereIn(dao.GenTable.Columns().TableId, tableIds).Update(do.GenTable{
			Status:     consts.GenTableStatusNo,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "删除生成表数据失败")
	})
	return
}

func (s *sGenTable) Add(ctx context.Context, req *v1.PostGenTableAddReq) (tableList *v1.PostGenTableAddRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		tables := strings.Split(req.Tables, ",")
		//添加生成表信息
		_, e := dao.GenTable.Ctx(ctx).WhereIn(dao.GenTable.Columns().TableName, tables).Update(do.GenTable{
			Status:     consts.GenTableStatusOk,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "添加生成表数据失败")
	})
	return
}
func (s *sGenTable) Update(ctx context.Context, req *v1.PutGenTableUpdateReq) (tableList *v1.PutGenTableUpdateRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//修改生成表信息
			_, e := dao.GenTable.Ctx(ctx).TX(tx).Where(dao.GenTable.Columns().TableId, req.TableId).Update(do.GenTable{
				BusinessName:   req.BusinessName,
				ClassName:      req.ClassName,
				FunctionName:   req.FunctionName,
				FunctionAuthor: req.FunctionAuthor,
				GenPath:        req.GenPath,
				GenType:        req.GenType,
				ModuleName:     req.ModuleName,
				Options:        req.Options,
				PackageName:    req.PackageName,
				Remark:         req.Remark,
				SubTableFkName: req.SubTableFkName,
				SubTableName:   req.SubTableName,
				TableComment:   req.TableComment,
				TableName:      req.TableName,
				TplCategory:    req.TplCategory,
				UpdateTime:     gtime.Now(),
				UpdateBy:       adminName,
			})
			utility.WriteErrLog(ctx, e, "修改生成表数据失败")
			for _, val := range req.Columns {
				_, e := dao.GenTableColumn.Ctx(ctx).TX(tx).Where(dao.GenTableColumn.Columns().ColumnId, val.ColumnId).Update(do.GenTableColumn{
					ColumnComment: val.ColumnComment,
					ColumnName:    val.ColumnName,
					ColumnType:    val.ColumnType,
					DictType:      val.DictType,
					GoField:       val.GoField,
					GoType:        val.GoType,
					HtmlType:      val.HtmlType,
					IsEdit:        val.IsEdit,
					IsInsert:      val.IsInsert,
					IsList:        val.IsList,
					IsIncrement:   val.IsIncrement,
					IsPk:          val.IsPk,
					IsQuery:       val.IsQuery,
					IsRequired:    val.IsRequired,
					JsonField:     val.JsonField,
					QueryType:     val.QueryType,
					Sort:          val.Sort,
					UpdateTime:    gtime.Now(),
					UpdateBy:      adminName,
				})
				utility.WriteErrLog(ctx, e, "修改生成表字段数据失败")
			}
		})
		return err
	})
	return
}

// 删除指定表名数据
func (s *sGenTable) DeleteTable(ctx context.Context, tableName string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var tableSilce *entity.GenTable
		err = dao.GenTable.Ctx(ctx).Where(dao.GenTable.Columns().TableName, tableName).Scan(&tableSilce)
		utility.WriteErrLog(ctx, err, "获取生成代码表名数据失败")
		//删除生成代码字段表数据
		_, err = dao.GenTableColumn.Ctx(ctx).Where(dao.GenTableColumn.Columns().TableId, tableSilce.TableId).Delete()
		utility.WriteErrLog(ctx, err, "删除生成代码字段表失败")
		//删除表名为tableName数据
		_, err = dao.GenTable.Ctx(ctx).Where(dao.GenTable.Columns().TableId, tableSilce.TableId).Delete()
	})
	return
}

// 生成表数据
func (s *sGenTable) InitTables(ctx context.Context, tableName string) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			db := g.DB()
			var tableArr []gdb.Record
			//获取所有表名
			if tableName == "" {
				tableArr, err = db.GetAll(ctx, "SELECT * FROM information_schema.Tables WHERE TABLE_SCHEMA='"+db.GetSchema()+"'")
			} else {
				tableArr, err = db.GetAll(ctx, "SELECT * FROM information_schema.Tables WHERE TABLE_SCHEMA='"+db.GetSchema()+"' AND TABLE_NAME='"+tableName+"'")
				if err == nil {
					//删除旧数据
					err = s.DeleteTable(ctx, tableName)
				}
			}
			utility.WriteErrLog(ctx, err, "获取数据库表名失败")
			adminName := gconv.String(ctx.Value(consts.CtxAdminName))
			for _, table := range tableArr {

				tableName := table["TABLE_NAME"].String()
				tableComment := table["TABLE_COMMENT"].String()

				exist, err := s.CheckTable(ctx, tableName)
				utility.WriteErrLog(ctx, err, "查询生成表失败")
				if exist { //生成过，跳过
					continue
				}
				moduleName := lib.GetPrefixName(tableName)
				tableId, err := dao.GenTable.Ctx(ctx).TX(tx).InsertAndGetId(do.GenTable{
					TableName:    tableName,
					TableComment: tableComment,
					ClassName:    lib.StrFirstToUpperS(tableName),
					//处理路径
					PackageName: "resource/gen/" + moduleName,
					//生成模块名
					ModuleName: moduleName,
					//生成业务名
					BusinessName: tableName,
					//生成功能名
					FunctionName: tableComment,
					//生成作者
					FunctionAuthor: adminName,
					CreateTime:     gtime.Now(),
					CreateBy:       adminName,
					UpdateTime:     gtime.Now(),
					UpdateBy:       adminName,
				})
				//fmt.Println(tableId)
				utility.WriteErrLog(ctx, err, "添加表数据失败")
				fieldList, err := db.Ctx(ctx).TableFields(ctx, tableName)
				//fmt.Println(tableList)
				for _, field := range fieldList {
					//特殊字段跳过 update_by update_time create_by
					if field.Name == "update_by" || field.Name == "update_time" || field.Name == "create_by" {
						continue
					}
					columnExist, err := s.CheckTableColumn(ctx, field.Name, tableId)
					utility.WriteErrLog(ctx, err, "查询生成表字段失败")
					if columnExist { //生成过，跳过
						continue
					}
					isList := "0"
					typeName, err := s.GetGolangTypeBySqlType(ctx, field)
					if err != nil {
						return
					}
					isInsert := "1"
					isEdit := "1"
					sort := 0
					isIncrement := "0"
					if field.Extra == "auto_increment" {
						isIncrement = "1"
						isInsert = "0"
						isEdit = "0"
					}
					isPk := "0"
					if field.Key == "PRI" {
						isPk = "1"
						isList = "1"
						isInsert = "0"
						isEdit = "0"
						sort = 11
					}
					isQuery := "0"
					if strings.Contains(field.Name, "name") {
						isQuery = "1"
						isList = "1"
						sort = 10
					}
					queryType := "EQ"
					htmlType := ""
					if field.Name == "create_time" {
						isQuery = "1"
						isList = "1"
						queryType = "BETWEEN"
						htmlType = "datetime"
						isInsert = "0"
						isEdit = "0"
					}
					dictType := ""
					if field.Name == "status" {
						isQuery = "1"
						isList = "1"
						dictType = "sys_normal_disable"
						htmlType = "select"
						sort = 9
					}
					isRequired := "0"
					if !field.Null {
						isRequired = "1"
					}
					//判断显示列表
					//处理对应的代码生成业务表字段
					_, e := dao.GenTableColumn.Ctx(ctx).TX(tx).Data(do.GenTableColumn{
						TableId:       tableId,
						ColumnName:    field.Name,
						ColumnComment: lib.SubStr(field.Comment, "（"),
						ColumnType:    field.Type,
						ColumnDef:     field.Default,
						GoField:       lib.StrFirstToUpperS(field.Name),
						GoType:        typeName,
						JsonField:     lib.StrFirstToUpper(field.Name),
						IsIncrement:   isIncrement,
						IsPk:          isPk,
						IsInsert:      isInsert,
						IsEdit:        isEdit,
						IsQuery:       isQuery,
						IsList:        isList,
						QueryType:     queryType,
						HtmlType:      htmlType,
						DictType:      dictType,
						IsRequired:    isRequired,
						Sort:          sort,
						UpdateTime:    gtime.Now(),
						UpdateBy:      adminName,
						CreateTime:    gtime.Now(),
						CreateBy:      adminName,
					}).Insert()
					utility.WriteErrLog(ctx, e, "添加代码生成业务表字段失败")
				}
			}
		})
		return err
	})

	return

}

func (s *sGenTable) GetGolangTypeBySqlType(ctx context.Context, field *gdb.TableField) (typeName string, err error) {
	db := g.DB()
	typeName, err = db.CheckLocalTypeForField(ctx, field.Type, nil)
	if err != nil {
		return
	}
	switch typeName {
	case gdb.LocalTypeDate, gdb.LocalTypeDatetime:
		if consts.StdTime {
			typeName = "time.Time"
		} else {
			typeName = "*gtime.Time"
		}

	case gdb.LocalTypeInt64Bytes:
		typeName = "int64"

	case gdb.LocalTypeUint64Bytes:
		typeName = "uint64"

	// Special type handle.
	case gdb.LocalTypeJson, gdb.LocalTypeJsonb:
		if consts.GJsonSupport {
			typeName = "*gjson.Json"
		} else {
			typeName = "string"
		}
	}
	return
}

func (s *sGenTable) CheckTable(ctx context.Context, tableName string) (bool, error) {
	count, err := dao.GenTable.Ctx(ctx).Where(do.GenTable{
		TableName: tableName,
	}).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *sGenTable) CheckTableColumn(ctx context.Context, columnName string, tableId int64) (bool, error) {
	count, err := dao.GenTableColumn.Ctx(ctx).Where(do.GenTableColumn{
		TableId:    tableId,
		ColumnName: columnName,
	}).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *sGenTable) GenPreview(ctx context.Context, req *v1.GetGenPreviewReq) (res *v1.GetGenPreviewRes, err error) {
	res = &v1.GetGenPreviewRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		table, err := s.GetGenTableSilceByTableId(ctx, req.TableId)
		if err != nil {
			return
		}
		//设置预览模式
		s.IsPreview = true
		ids := []int64{req.TableId}
		tableColumnMap, err := service.GenTableColumn().GetTableColumnByTableIds(ctx, ids)
		for _, columns := range tableColumnMap {
			//阅览文件
			err = s.GenApiFile("dir", table, columns)
			utility.WriteErrLog(ctx, err, "阅览文件错误")
		}
		res.List = fileCodeMap
	})
	return
}
func (s *sGenTable) BatchGenCode(ctx context.Context, req *v1.GetBatchGenCodeReq) (tableList *v1.GetBatchGenCodeRes, err error) {
	//adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	UserId := gconv.String(ctx.Value(consts.CtxAdminId))
	//创建临时存放代码目录
	dir := s.CreateDir(UserId)
	err = g.Try(ctx, func(ctx context.Context) {
		tables := strings.Split(req.Tables, ",")

		//获取gentable数据
		tableSilce, err := s.GetGenTableSilceByName(ctx, tables)
		if err != nil {
			return
		}
		var ids []int64
		for _, table := range tableSilce {
			ids = append(ids, table.TableId)
		}
		tableColumnMap, err := service.GenTableColumn().GetTableColumnByTableIds(ctx, ids)
		for _, table := range tableSilce {
			tableId := strconv.FormatInt(table.TableId, 10)
			//fmt.Println(tableColumnMap[tableId])
			//生成文件
			err = s.GenApiFile(dir, table, tableColumnMap[tableId])
			utility.WriteErrLog(ctx, err, "生成文件错误")

		}
		//文件打包
		downFilePath := consts.GenCodeZipDir + UserId
		if !gfile.IsDir(downFilePath) {
			gfile.Mkdir(downFilePath)
		}
		downFile := downFilePath + "/govueadmin.zip"
		err = gcompress.ZipPath(dir+"/", downFile)
		utility.WriteErrLog(ctx, err, "打包错误")

		//下载文件
		ghttp.RequestFromCtx(ctx).Response.ServeFileDownload(downFile)
		return
	})
	//删除生成代码目录
	if err == nil {
		gfile.Remove(dir)
	}
	return
}

// 生成controller文件
func (s *sGenTable) GenControllerFile(codeDir string, table *entity.GenTable, tplCtrlContent []model.TplCtrlContent) (err error) {
	params := g.Map{
		"packageName":    table.TableName,
		"apiPackageName": consts.ApiPackageName,
		"nameController": lib.StrFirstToUpper(table.TableName) + "Controller",
		"controllerName": table.ClassName,
		"funcSilce":      tplCtrlContent,
	}
	src := consts.TemplateDir + "internal/controller/controller.tpl"
	dst := codeDir + "/internal/controller/" + table.TableName + "/" + table.TableName + ".go"
	//写入controller接口文件
	err = s.WriteDownFile(src, dst, params, "controller")
	//格式化文件
	lib.FmtGoFile(dst)
	return
}

func (s *sGenTable) LogicTplContent(funcType string, table *entity.GenTable, tableColumnMap []*entity.GenTableColumn) (tlc model.TplLogicContent) {
	tlc = model.TplLogicContent{ApiPackageName: consts.ApiPackageName, ClassName: table.ClassName}
	build := strings.Builder{}
	switch funcType {
	case "list":
		tlc.FuncName = consts.ListMethod + table.ClassName + "List"
		tlc.Notes = "获取" + table.TableComment + "列表"
		tlc.ParamReq = consts.ListMethod + table.ClassName + "ListReq"
		tlc.ParamRes = consts.ListMethod + table.ClassName + "ListRes"
		build.WriteString("res = &" + tlc.ApiPackageName + "." + tlc.ParamRes + "{}\n")
		build.WriteString("\tif req.PageNum == 0 {\n\t\treq.PageNum = 1\n\t}\n\tif req.PageSize == 0 {\n\t\treq.PageSize = consts.PageSize\n\t}\n")
		build.WriteString("\terr = g.Try(ctx, func(ctx context.Context) {\n")
		build.WriteString("var list []*entity." + table.ClassName + "\n")
		build.WriteString("m := dao." + table.ClassName + ".Ctx(ctx)\n")
		for _, field := range tableColumnMap {
			//查询字段
			if field.IsQuery == "1" {
				if field.QueryType == "LIKE" {
					build.WriteString("if req." + field.GoField + " != \"\" {\n")
					build.WriteString("m = m.WhereLike(dao." + table.ClassName + ".Columns()." + field.GoField + ", \"%\"+req." + field.GoField + "+\"%\")\n")

				} else if field.QueryType == "BETWEEN" && field.ColumnType == "datetime" {
					build.WriteString("if len(req.Params) > 0 {\n")
					build.WriteString("m = m.WhereBetween(dao." + table.ClassName + ".Columns()." + field.GoField + ", req.Params[\"beginTime\"], req.Params[\"endTime\"])\n")
				} else {
					build.WriteString("if req." + field.GoField + " != \"\" {\n")
					build.WriteString("m = m.Where(dao." + table.ClassName + ".Columns()." + field.GoField + " , req." + field.GoField + ")\n")
				}
				build.WriteString("}\n")
			}
		}
		build.WriteString("\t\tres.Total, err = m.Count()\n\t\terr = m.Page(req.PageNum, req.PageSize).Scan(&list)\n\n")
		build.WriteString("\t\tutility.WriteErrLog(ctx, err, \"获取" + table.TableComment + "失败\")\n")
		build.WriteString("res.Rows = list\n")
		build.WriteString("})\n")
		//fmt.Println(tableColumnMap)
		tlc.Content = build.String()
	case "add":
		tlc.FuncName = "Add"
		tlc.Notes = "添加" + table.TableComment
		tlc.ParamReq = consts.AddMethod + table.ClassName + "Req"
		tlc.ParamRes = consts.AddMethod + table.ClassName + "Res"
		build.WriteString("adminName := gconv.String(ctx.Value(consts.CtxAdminName))\n")
		build.WriteString("err = g.Try(ctx, func(ctx context.Context) {\n")
		build.WriteString("_, e := dao." + table.ClassName + ".Ctx(ctx).Data(do." + table.ClassName + "{\n")
		for _, field := range tableColumnMap {
			if field.IsInsert == "1" && field.IsIncrement != "1" {
				build.WriteString(field.GoField + ":   req." + field.GoField + ",\n")
			}
		}
		build.WriteString("\t\t\tUpdateTime: gtime.Now(),\n\t\t\tUpdateBy:   adminName,\n\t\t\tCreateTime: gtime.Now(),\n\t\t\tCreateBy:   adminName,\n")

		build.WriteString("}).Insert()\n")
		build.WriteString("utility.WriteErrLog(ctx, e, \"添加" + table.TableComment + "数据失败\")\n")
		build.WriteString("\t})\n")
		tlc.Content = build.String()
	case "update":
		tlc.FuncName = "Update"
		tlc.Notes = "修改" + table.TableComment
		tlc.ParamReq = consts.UpdateMethod + table.ClassName + "Req"
		tlc.ParamRes = consts.UpdateMethod + table.ClassName + "Res"
		build.WriteString("adminName := gconv.String(ctx.Value(consts.CtxAdminName))\n")
		build.WriteString("err = g.Try(ctx, func(ctx context.Context) {\n")
		for _, field := range tableColumnMap {
			if field.IsPk == "1" && field.IsIncrement == "1" {
				build.WriteString("_, e := dao." + table.ClassName + ".Ctx(ctx).WherePri(&req." + field.GoField + ").Update(do." + table.ClassName + "{\n")
				break
			}
		}
		for _, field := range tableColumnMap {
			if field.IsEdit == "1" && field.IsIncrement != "1" {
				build.WriteString(field.GoField + ":   req." + field.GoField + ",\n")
			}
		}
		build.WriteString("\t\t\tUpdateTime: gtime.Now(),\n\t\t\tUpdateBy:   adminName,\n")
		build.WriteString("\t})\n")
		build.WriteString("utility.WriteErrLog(ctx, e, \"修改" + table.TableComment + "数据失败\")\n")
		build.WriteString("\t})\n")
		tlc.Content = build.String()
	case "delete":
		tlc.FuncName = "Delete"
		tlc.Notes = "删除" + table.TableComment
		tlc.ParamReq = consts.DeleteMethod + table.ClassName + "Req"
		tlc.ParamRes = consts.DeleteMethod + table.ClassName + "Res"
		build.WriteString("err = g.Try(ctx, func(ctx context.Context) {\n")
		statusExists := false
		for _, field := range tableColumnMap {
			if field.IsPk == "1" {
				build.WriteString("postIds := lib.ParamStrToSlice(req." + field.GoField + ", \",\")\n")
				build.WriteString("_, e := dao." + table.ClassName + ".Ctx(ctx).WhereIn(dao." + table.ClassName + ".Columns()." + field.GoField + ",postIds).")
				break
			}
			if field.ColumnName == "status" {
				statusExists = true
			}
		}
		if statusExists {
			build.WriteString("Update(do." + table.ClassName + "{\n")
			build.WriteString("Status:   consts." + table.ClassName + "StatusNo ,\n")
			build.WriteString("\t\t\tUpdateTime: gtime.Now(),\n\t\t\tUpdateBy:   gconv.String(ctx.Value(consts.CtxAdminName)),\n")
			build.WriteString("\t})\n")
		} else {
			build.WriteString("Delete()\n")
		}
		build.WriteString("utility.WriteErrLog(ctx, e, \"删除" + table.TableComment + "数据失败\")\n")

		build.WriteString("\t})\n")
		tlc.Content = build.String()
	case "get":
		tlc.FuncName = consts.ListMethod + table.ClassName
		tlc.Notes = "获取" + table.TableComment
		tlc.ParamReq = consts.ListMethod + table.ClassName + "Req"
		tlc.ParamRes = consts.ListMethod + table.ClassName + "Res"
		build.WriteString("res = &" + tlc.ApiPackageName + "." + tlc.ParamRes + "{}\n")
		build.WriteString("err = g.Try(ctx, func(ctx context.Context) {\n")
		build.WriteString("var table *entity." + table.ClassName + "\n")

		for _, field := range tableColumnMap {
			if field.IsPk == "1" {
				build.WriteString("err = dao." + table.ClassName + ".Ctx(ctx).Where(dao." + table.ClassName + ".Columns()." + field.GoField + ", req." + field.GoField + ").Scan(&table)\n")
				break
			}
		}
		build.WriteString(" res." + table.ClassName + "=table\n")
		build.WriteString("utility.WriteErrLog(ctx, err, \"删除" + table.TableComment + "数据失败\")\n")
		build.WriteString("\t})\n")
		tlc.Content = build.String()
	default:

	}
	return
}

func (s *sGenTable) ApiJsTplContent(funcType string, table *entity.GenTable) (tlc model.TplApiJsContent) {
	tlc = model.TplApiJsContent{}
	build := strings.Builder{}
	url := s.GetApiPath(table.TableName)
	switch funcType {
	case "list":
		tlc.FuncName = "list" + table.ClassName
		apiJsMap["list"] = tlc.FuncName
		tlc.Notes = "获取" + table.TableComment + "列表"
		build.WriteString("url: '" + url + "/list',\n    method: 'get',\n    params: query\n")
		tlc.Content = build.String()
	case "add":
		tlc.FuncName = "add" + table.ClassName
		apiJsMap["add"] = tlc.FuncName
		tlc.Notes = "添加" + table.TableComment
		build.WriteString("    url: '" + url + "',\n    method: 'post',\n    data: query\n")
		tlc.Content = build.String()
	case "update":
		tlc.FuncName = "update" + table.ClassName
		tlc.Notes = "修改" + table.TableComment
		apiJsMap["update"] = tlc.FuncName
		build.WriteString("    url: '" + url + "',\n    method: 'put',\n    data: query\n")
		tlc.Content = build.String()
	case "delete":
		tlc.FuncName = "delete" + table.ClassName
		tlc.Notes = "删除" + table.TableComment
		apiJsMap["delete"] = tlc.FuncName
		build.WriteString("    url: '" + url + "/' + query,\n    method: 'delete'\n")
		tlc.Content = build.String()
	case "get":
		tlc.FuncName = "get" + table.ClassName
		tlc.Notes = "获取" + table.TableComment
		apiJsMap["get"] = tlc.FuncName
		build.WriteString("    url: '" + url + "/' + query,\n    method: 'get'\n")
		tlc.Content = build.String()
	default:

	}
	return
}

// 生成logic文件
func (s *sGenTable) GenLogicFile(codeDir string, table *entity.GenTable, tplLogicContent []model.TplLogicContent) (err error) {
	params := g.Map{
		"packageName": table.TableName,
		"className":   table.ClassName,
		"logicSilce":  tplLogicContent,
	}
	//fmt.Println(tplLogicContent)
	src := consts.TemplateDir + "internal/logic/logic.tpl"
	dst := codeDir + "/internal/logic/" + table.TableName + "/" + table.TableName + ".go"
	//写入logic接口文件
	err = s.WriteDownFile(src, dst, params, "logic")
	//格式化文件
	if err == nil {
		lib.FmtGoFile(dst)
	}
	return
}
func (s *sGenTable) GetApiPath(tableName string) string {
	if tableName == "" {
		return ""
	}
	tableName = strings.Replace(tableName, "sys_", "system_", -1)
	return "/" + strings.Replace(tableName, "_", "/", -1)
}

// 生成api文件
func (s *sGenTable) GenApiFile(codeDir string, table *entity.GenTable, tableColumnMap []*entity.GenTableColumn) (err error) {
	var listReq, listRes, addReq, addRes, updateReq, updateRes, getUpdateReq, getUpdateRes, deleteReq, deleteRes model.TplContent
	var apiSilce []model.TplContent
	var listUrl, getUpdateUrl, deleteUrl, urlPath, pkId string
	urlPath = s.GetApiPath(table.TableName)
	listUrl = urlPath + "/list"
	listBuild := strings.Builder{}
	listBuild.WriteString(s.GetReqMeta(listUrl, consts.ListMethod, "Get "+table.TableName))
	listReq.TypeName = consts.ListMethod + table.ClassName + "ListReq"

	addReqBuild := strings.Builder{}
	addReqBuild.WriteString(s.GetReqMeta(urlPath, consts.AddMethod, table.TableName))
	updateReqBuild := strings.Builder{}
	updateReqBuild.WriteString(s.GetReqMeta(urlPath, consts.UpdateMethod, table.TableName))
	deleteReqBuild := strings.Builder{}
	getUpdateReqBuild := strings.Builder{}
	for _, column := range tableColumnMap {
		if column.IsQuery == "1" { //列表检索字段
			listBuild.WriteString(column.GoField)
			listBuild.WriteString(" ")
			listBuild.WriteString(column.GoType)
			listBuild.WriteString(" `p:\"")
			listBuild.WriteString(column.JsonField)
			listBuild.WriteString("\"`\n")
		}
		if column.IsInsert == "1" && column.IsIncrement != "1" { //插入
			addReqBuild.WriteString(column.GoField)
			addReqBuild.WriteString("  ")
			addReqBuild.WriteString(column.GoType)
			addReqBuild.WriteString(" `p:\"")
			addReqBuild.WriteString(column.JsonField)
			addReqBuild.WriteString("\"")
			if column.IsRequired == "1" { //是否必填
				addReqBuild.WriteString("  v:\"required\"")
			}
			addReqBuild.WriteString(" `\n")
		}
		if column.IsEdit == "1" || column.IsPk == "1" { //编辑
			updateReqBuild.WriteString(column.GoField)
			updateReqBuild.WriteString("  ")
			updateReqBuild.WriteString(column.GoType)
			updateReqBuild.WriteString(" `p:\"")
			updateReqBuild.WriteString(column.JsonField)
			updateReqBuild.WriteString("\"")
			if column.IsRequired == "1" { //是否必填
				updateReqBuild.WriteString("  v:\"required\"")
			}
			updateReqBuild.WriteString(" `\n")
		}
		if column.IsPk == "1" { //处理删除
			pkId = column.JsonField
			deleteReqBuild.WriteString(column.GoField)
			deleteReqBuild.WriteString("  ")
			deleteReqBuild.WriteString("string")
			deleteReqBuild.WriteString(" `p:\"")
			deleteReqBuild.WriteString(column.JsonField)
			deleteReqBuild.WriteString("\"  v:\"required\"")
			deleteReqBuild.WriteString(" `\n")
			//getUpdate
			getUpdateReqBuild.WriteString(column.GoField)
			getUpdateReqBuild.WriteString("  ")
			getUpdateReqBuild.WriteString(column.GoType)
			getUpdateReqBuild.WriteString(" `p:\"")
			getUpdateReqBuild.WriteString(column.JsonField)
			getUpdateReqBuild.WriteString("\"  v:\"required\"")
			getUpdateReqBuild.WriteString(" `\n")

		}
	}
	listBuild.WriteString("	common.PageReq \n")
	listReq.TypeContent = listBuild.String()
	apiSilce = append(apiSilce, listReq)

	listRes.TypeName = consts.ListMethod + table.ClassName + "ListRes"
	listResBuild := strings.Builder{}
	listResBuild.WriteString("g.Meta `mime:\"application/json\"` \n")
	listResBuild.WriteString("Rows   []*entity.")
	listResBuild.WriteString(table.ClassName)
	listResBuild.WriteString(" `json:\"rows\"` \n")
	listResBuild.WriteString("Total  int   `json:\"total\"` \n")
	listRes.TypeContent = listResBuild.String()
	apiSilce = append(apiSilce, listRes)
	//固定生成 列表 添加 编辑 删除
	//列表
	//api
	//添加
	addReq.TypeName = consts.AddMethod + table.ClassName + "Req"
	addReq.TypeContent = addReqBuild.String()
	addRes.TypeName = consts.AddMethod + table.ClassName + "Res"
	addRes.TypeContent = "g.Meta `mime:\"application/json\"` \n"

	//编辑
	updateReq.TypeName = consts.UpdateMethod + table.ClassName + "Req"
	updateReq.TypeContent = updateReqBuild.String()
	updateRes.TypeName = consts.UpdateMethod + table.ClassName + "Res"
	updateRes.TypeContent = "g.Meta `mime:\"application/json\"` \n"

	//获取编辑数据显示
	getUpdateReq.TypeName = consts.ListMethod + table.ClassName + "Req"

	getUpdateUrl = urlPath + "/{" + pkId + "}"
	getUpdateMeta := s.GetReqMeta(getUpdateUrl, consts.ListMethod, table.TableName)
	getUpdateReq.TypeContent = getUpdateMeta + getUpdateReqBuild.String()

	getUpdateRes.TypeName = consts.ListMethod + table.ClassName + "Res"
	getUpdateResBuild := strings.Builder{}
	getUpdateResBuild.WriteString("g.Meta `mime:\"application/json\"` \n")
	getUpdateResBuild.WriteString("*entity.")
	getUpdateResBuild.WriteString(table.ClassName)
	getUpdateResBuild.WriteString("\n")
	getUpdateRes.TypeContent = getUpdateResBuild.String()

	//删除
	deleteUrl = urlPath + "/{" + pkId + "}"
	deleteMeta := s.GetReqMeta(deleteUrl, consts.DeleteMethod, table.TableName)
	deleteReq.TypeContent = deleteMeta + deleteReqBuild.String()
	deleteReq.TypeName = consts.DeleteMethod + table.ClassName + "Req"
	deleteRes.TypeName = consts.DeleteMethod + table.ClassName + "Res"
	deleteRes.TypeContent = "g.Meta `mime:\"application/json\"` \n"

	apiSilce = append(apiSilce, addReq)
	apiSilce = append(apiSilce, addRes)
	apiSilce = append(apiSilce, updateReq)
	apiSilce = append(apiSilce, updateRes)
	apiSilce = append(apiSilce, deleteReq)
	apiSilce = append(apiSilce, deleteRes)
	apiSilce = append(apiSilce, getUpdateReq)
	apiSilce = append(apiSilce, getUpdateRes)

	params := g.Map{
		"packageName": consts.ApiPackageName,
		"apiSilce":    apiSilce,
	}
	src := consts.TemplateDir + "v1/v.tpl"
	dst := codeDir + "/api/" + consts.ApiPackageName + "/" + table.TableName + ".go"
	//写入api接口文件
	err = s.WriteDownFile(src, dst, params, "api")
	if err != nil {
		return
	}
	//格式化api文件
	lib.FmtGoFile(dst)
	//写入controller
	var tplCtrlSilce []model.TplCtrlContent
	nc := lib.StrFirstToUpper(table.TableName) + "Controller"
	tplCtrlSilce = append(tplCtrlSilce, model.TplCtrlContent{FuncName: consts.ListMethod + table.ClassName + "List", ServiceReq: listReq.TypeName, ServiceRes: listRes.TypeName, ServiceName: table.ClassName, ApiPackageName: consts.ApiPackageName, NameController: nc})
	tplCtrlSilce = append(tplCtrlSilce, model.TplCtrlContent{FuncName: "Add", ServiceReq: addReq.TypeName, ServiceRes: addRes.TypeName, ServiceName: table.ClassName, ApiPackageName: consts.ApiPackageName, NameController: nc})
	tplCtrlSilce = append(tplCtrlSilce, model.TplCtrlContent{FuncName: "Update", ServiceReq: updateReq.TypeName, ServiceRes: updateRes.TypeName, ServiceName: table.ClassName, ApiPackageName: consts.ApiPackageName, NameController: nc})
	tplCtrlSilce = append(tplCtrlSilce, model.TplCtrlContent{FuncName: "Delete", ServiceReq: deleteReq.TypeName, ServiceRes: deleteRes.TypeName, ServiceName: table.ClassName, ApiPackageName: consts.ApiPackageName, NameController: nc})
	tplCtrlSilce = append(tplCtrlSilce, model.TplCtrlContent{FuncName: consts.ListMethod + table.ClassName, ServiceReq: getUpdateReq.TypeName, ServiceRes: getUpdateRes.TypeName, ServiceName: table.ClassName, ApiPackageName: consts.ApiPackageName, NameController: nc})
	err = s.GenControllerFile(codeDir, table, tplCtrlSilce)
	if err != nil {
		return
	}
	//写入consts
	err = s.GenConstsFile(codeDir, table)
	if err != nil {
		return
	}
	//写入logic
	var tplLogicSilce []model.TplLogicContent
	var tplApiJsSilce []model.TplApiJsContent
	funcSilce := []string{"list", "add", "update", "delete", "get"}
	for _, val := range funcSilce {
		tplLogicSilce = append(tplLogicSilce, s.LogicTplContent(val, table, tableColumnMap))
		tplApiJsSilce = append(tplApiJsSilce, s.ApiJsTplContent(val, table))

	}
	err = s.GenLogicFile(codeDir, table, tplLogicSilce)
	if err != nil {
		return
	}
	//写入js
	err = s.GenApiJsFile(codeDir, table, tplApiJsSilce)
	if err != nil {
		return
	}
	//写入vue
	tvc := s.VueTplContent(funcSilce, table, tableColumnMap)
	err = s.GenVueFile(codeDir, table, tvc)
	return
}

func (s *sGenTable) GetStatusDictType(tableName, columnName string) (res string, ok bool) {
	if columnName == "status" {
		res = "sys_normal_disable"
		ok = true
		return
	}
	if res, ok = model.DictTypeMap[tableName+columnName]; ok {
		return
	}
	return
}
func (s *sGenTable) VueTplContent(funcSilce []string, table *entity.GenTable, tableColumnMap []*entity.GenTableColumn) (tvc model.TplVueContent) {
	tvc = model.TplVueContent{ClaseName: table.ClassName, TableComment: table.TableComment}

	build := strings.Builder{}
	paramBuild := strings.Builder{}
	rulesBuild := strings.Builder{}
	listColumnBuild := strings.Builder{}
	editColumnBuild := strings.Builder{}
	var dictTypeSilce []string
	tvc.ListParamQuery = "queryParams.value"
	for _, column := range tableColumnMap {
		//不处理create_by create_time update_time update_by
		if column.ColumnName == "create_by" && column.ColumnName == "update_by" && column.ColumnName == "update_time" {
			continue
		}
		//是主键并且自增
		if column.IsPk == "1" && column.IsIncrement == "1" {
			tvc.PKComment = column.ColumnComment
			tvc.PKJsonField = column.JsonField
		}
		//处理search
		if column.IsQuery == "1" {
			serchTpl := model.TplVueSearchContent{JsonField: column.JsonField, ColumnComment: column.ColumnComment}
			if v, ok := s.GetStatusDictType(table.TableName, column.ColumnName); ok { //判断是否是下拉框
				dictTypeSilce = append(dictTypeSilce, v)
				serchTpl.Content = "<el-select v-model=\"queryParams." + column.JsonField + "\" placeholder=\"" + column.ColumnComment + "\" clearable style=\"width: 200px\">\n               <el-option\n                  v-for=\"dict in " + v + "\"\n                  :key=\"dict.value\"\n                  :label=\"dict.label\"\n                  :value=\"dict.value\"\n               />\n            </el-select>"
			} else if column.ColumnName == "create_time" {
				serchTpl.Content = "                  <el-date-picker\n                     v-model=\"dateRange\"\n                     value-format=\"YYYY-MM-DD\"\n                     type=\"daterange\"\n                     range-separator=\"-\"\n                     start-placeholder=\"开始日期\"\n                     end-placeholder=\"结束日期\"\n                  ></el-date-picker>\n"
			} else {
				serchTpl.Content = "<el-input\n               v-model=\"queryParams." + column.JsonField + "\"\n               placeholder=\"请输入" + column.ColumnComment + "\"\n               clearable\n               style=\"width: 200px\"\n               @keyup.enter=\"handleQuery\"\n            />"
			}
			tvc.Search = append(tvc.Search, serchTpl)
			paramBuild.WriteString(column.JsonField + ": undefined,\n")
		}
		if column.ColumnName == "create_time" {
			tvc.ResetQuery = "  dateRange.value = [];\n"
			tvc.ListParamQuery = "proxy.addDateRange(queryParams.value, dateRange.value)\n"
			continue
		}
		//是否必填
		if column.IsRequired == "1" {
			rulesBuild.WriteString(column.JsonField + ": [{ required: true, ")
			if column.ColumnName == "email" {
				rulesBuild.WriteString(" type: \"email\", ")
			}
			if column.ColumnName == "phone" {
				rulesBuild.WriteString("pattern: /^1[3|4|5|6|7|8|9][0-9]\\d{8}$/,")
			}
			rulesBuild.WriteString("message: \"" + column.ColumnComment + "不能为空\", trigger: \"blur\" }],\n")
		}
		//处理email不必须
		if column.ColumnName == "email" && column.IsRequired != "1" {
			rulesBuild.WriteString(column.JsonField + ": [{  type: \"email\", ")
			rulesBuild.WriteString("message: \"" + column.ColumnComment + "不能为空\", trigger: \"blur\" }],\n")
		}
		//处理phone不必须
		if column.ColumnName == "phone" && column.IsRequired != "1" {
			rulesBuild.WriteString(column.JsonField + ": [{ pattern: /^1[3|4|5|6|7|8|9][0-9]\\d{8}$/,")
			rulesBuild.WriteString("message: \"" + column.ColumnComment + "不能为空\", trigger: \"blur\" }],\n")
		}
		build.WriteString(column.JsonField + ":")
		//重置字段
		if column.GoType == "string" || column.IsIncrement == "1" {
			if column.ColumnDef != "" {
				build.WriteString("\"" + column.ColumnDef + "\"")
			} else {
				build.WriteString("undefined")
			}
		} else {
			if column.ColumnDef != "" {
				build.WriteString(column.ColumnDef)
			} else {
				build.WriteString("0")
			}
		}
		build.WriteString(",\n")
		//处理列表
		if column.IsList == "1" {
			if column.DictType != "" {
				listColumnBuild.WriteString("         <el-table-column label=\"" + column.ColumnComment + "\" align=\"center\" prop=\"" + column.JsonField + "\">\n            <template #default=\"scope\">\n               <dict-tag :options=\"" + column.DictType + "\" :value=\"scope.row." + column.JsonField + "\" />\n            </template>\n         </el-table-column>\n")
			} else if column.ColumnType == "datetime" {
				listColumnBuild.WriteString("         <el-table-column label=\"" + column.ColumnComment + "\" align=\"center\" prop=\"" + column.JsonField + "\" width=\"180\">\n            <template #default=\"scope\">\n               <span>{{ parseTime(scope.row." + column.JsonField + ") }}</span>\n            </template>\n         </el-table-column>\n")
			} else {
				listColumnBuild.WriteString("<el-table-column label=\"" + column.ColumnComment + "\" align=\"center\" prop=\"" + column.JsonField + "\" />\n")
			}
		}
		//处理编辑添加
		if column.IsEdit == "1" && column.IsIncrement != "1" {
			if column.DictType == "" {
				editColumnBuild.WriteString("            <el-form-item label=\"" + column.ColumnComment + "\" prop=\"" + column.JsonField + "\">\n               <el-input v-model=\"form." + column.JsonField + "\" placeholder=\"请输入" + column.ColumnComment + "\" />\n            </el-form-item>\n")
			} else {
				editColumnBuild.WriteString("            <el-form-item label=\"" + column.ColumnComment + "\" prop=\"" + column.JsonField + "\">\n               <el-radio-group v-model=\"form." + column.JsonField + "\">\n                  <el-radio\n                     v-for=\"dict in " + column.DictType + "\"\n                     :key=\"dict.value\"\n                     :label=\"dict.value\"\n                  >{{ dict.label }}</el-radio>\n               </el-radio-group>\n            </el-form-item>")
			}
		}
	}
	tvc.EditColumn = editColumnBuild.String()
	tvc.ListColumn = listColumnBuild.String()
	tvc.ColumnReset = build.String()
	tvc.ColumnRules = rulesBuild.String()
	tvc.SearchParam = paramBuild.String()

	//引入字典类型
	tvc.DictType = s.GetDictTypeVue(dictTypeSilce)
	//引入方法
	var funcNameSilce []string
	for _, funcType := range funcSilce {
		funcNameSilce = append(funcNameSilce, funcType+table.ClassName)
	}
	tvc.ImportApiFunc = "import { " + strings.Join(funcNameSilce, ",") + " } from \"@/api/" + s.GetApiPath(table.TableName) + "\";"
	tvc.PermiPath = lib.GetPermiPath(table.TableName)
	return
}
func (s *sGenTable) GetDictTypeVue(dictType []string) string {
	if len(dictType) < 1 {
		return ""
	}
	str1 := strings.Join(dictType, ",")
	str2 := strings.Join(dictType, "\", \"")
	return "const { " + str1 + " } = proxy.useDict(\"" + str2 + "\");"
}

// 生成vue文件
func (s *sGenTable) GenVueFile(codeDir string, table *entity.GenTable, tvc model.TplVueContent) (err error) {
	params := g.Map{
		"SearchParam":    tvc.SearchParam,
		"ClassName":      table.ClassName,
		"VueSearchSilce": tvc.Search,
		"TableComment":   tvc.TableComment,
		"PKComment":      tvc.PKComment,
		"EditColumn":     tvc.EditColumn,
		"ListColumn":     tvc.ListColumn,
		"ColumnReset":    tvc.ColumnReset,
		"DictType":       tvc.DictType,
		"PKJsonField":    tvc.PKJsonField,
		"ImportApiFunc":  tvc.ImportApiFunc,
		"ColumnRules":    tvc.ColumnRules,
		"PermiPath":      tvc.PermiPath,
	}
	src := consts.TemplateDir + "vue/src/views/index.tpl"
	dst := codeDir + "/vue/src/views" + s.GetApiPath(table.TableName) + "/index.vue"
	//写入vue接口文件
	err = s.WriteDownFile(src, dst, params, "vue")
	return
}

// 生成consts文件
func (s *sGenTable) GenConstsFile(codeDir string, table *entity.GenTable) (err error) {
	params := g.Map{
		"TableName": table.ClassName,
	}
	src := consts.TemplateDir + "internal/consts/consts.tpl"
	dst := codeDir + "/internal/consts/" + table.TableName + ".go"
	//写入consts接口文件
	err = s.WriteDownFile(src, dst, params, "consts")
	//格式化文件
	if err == nil {
		lib.FmtGoFile(dst)
	}
	return
}

// 生成apijs文件
func (s *sGenTable) GenApiJsFile(codeDir string, table *entity.GenTable, tplApiJsContent []model.TplApiJsContent) (err error) {
	params := g.Map{
		"jsFuncSilce": tplApiJsContent,
	}
	src := consts.TemplateDir + "vue/src/api/js.tpl"
	dst := codeDir + "/vue/src/api/" + s.GetApiPath(table.TableName) + ".js"
	//写入api js接口文件
	err = s.WriteDownFile(src, dst, params, "js")

	return
}

// 生成文件
func (s *sGenTable) WriteDownFile(src, dst string, params g.Map, fileName string) (err error) {
	view := gview.New()
	content, err := view.Parse(context.TODO(), src, params)
	if err != nil {
		return
	}
	if s.IsPreview {
		fileCodeMap[fileName] = content
		return
	}
	//写入代码
	err = gfile.PutContents(dst, content)
	return
}

func (s *sGenTable) GetReqMeta(url, method, tableName string) (gMeta string) {
	build := strings.Builder{}
	build.WriteString("g.Meta `path:\"")
	build.WriteString(url)
	build.WriteString("\" method:\"")
	build.WriteString(method)
	build.WriteString("\" tags:\"")
	build.WriteString(method)
	build.WriteString(" ")
	build.WriteString(tableName)
	build.WriteString("\" summary:\"")
	build.WriteString(method)
	build.WriteString(" ")
	build.WriteString(tableName)
	build.WriteString("\"`")
	build.WriteString("\n")
	return build.String()
}
func (s *sGenTable) CreateDir(userId string) (dir string) {
	dir = consts.GenCodeDir + gtime.Now().Format("YmdHis") + userId
	if !gfile.IsDir(dir) {
		gfile.Mkdir(dir)
	}
	return
}

func (s *sGenTable) GetGenTableSilceByName(ctx context.Context, tableName []string) (tableSilce []*entity.GenTable, err error) {
	tableSilce = []*entity.GenTable{}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.GenTable.Ctx(ctx).WhereIn(dao.GenTable.Columns().TableName, tableName).Scan(&tableSilce)
		utility.WriteErrLog(ctx, err, "获取生成代码数据失败")
	})
	return
}

func (s *sGenTable) GetGenTableSilceByTableId(ctx context.Context, tableId int64) (tableSilce *entity.GenTable, err error) {
	tableSilce = &entity.GenTable{}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.GenTable.Ctx(ctx).Where(dao.GenTable.Columns().TableId, tableId).Scan(tableSilce)
		utility.WriteErrLog(ctx, err, "获取生成代码数据失败")
	})
	return
}
