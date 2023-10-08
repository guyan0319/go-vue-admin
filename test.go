package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gogf/gf/v2/frame/g"

	"golang.org/x/mod/modfile"
)

func main() {
	var aa map[string][]string
	aa = make(map[string][]string, 0)
	aa["list"] = append(aa["list"], "fasdf")
	fmt.Println(aa)

	//getModule()
	//err := gcompress.ZipPath("temp/down/202309051513511/", consts.GenCodeDir+"1/goadmin.zip")
	//err := gcompress.ZipPath("temp/down/202309051513511/", consts.GenCodeZipDir+"1/govueadmin.zip")
	//err := gcompress.ZipPath("temp/down/202309051513511/", consts.GenCodeDir+"1/goadmin.zip")
	//fmt.Println(err)

	//userId := "2"
	//path := consts.GenCodeDir + gtime.Now().Format("YmdHis") + userId
	//if !gfile.IsDir(path) {
	//	fmt.Println(path)
	//	gfile.Mkdir(path)
	//}
	//gfile.Remove(path)
	//gfile.Remove("temp/202309032304332")
	//test()
	//view := gview.New()
	//view.SetConfigWithMap(g.Map{
	//	"Paths": []string{"template"},
	//	//"DefaultFile": "index.html",
	//	"Delimiters": []string{"{{", "}}"},
	//	//"Data": g.Map{
	//	//	"name":    "gf",
	//	//	"version": "1.10.0",
	//	//},
	//})
	//params := g.Map{
	//	"name": "gffasdfas",
	//}
	////
	//result, err := view.Parse(context.TODO(), "index.tpl", params)
	//////result, err := view.ParseContent(context.TODO(), "hello {{.name}}, version: {{.version}}")
	//if err != nil {
	//	panic(err)
	//}
	////下载
	//
	//fmt.Println(result)
}
func test() {
	// 绑定全局的模板函数
	g.View().BindFunc("hello", funcHello)

	// 普通方式传参
	parsed1, err := g.View().ParseContent(context.TODO(), `{{hello "GoFrame"}}`, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(parsed1))

	// 通过管道传参
	parsed2, err := g.View().ParseContent(context.TODO(), `{{"GoFrame" | hello}}`, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(parsed2))
}

// 用于测试的带参数的内置函数
func funcHello(name string) string {
	return fmt.Sprintf(`Hello %s`, name)
}
func getModule() (path string) {
	goModFilePathData, _ := os.ReadFile("go.mod")
	modFile, _ := modfile.Parse("go.mod", goModFilePathData, nil)
	fmt.Println("require如下：")
	fmt.Println("")
	fmt.Println(modFile.Module.Mod)
	//for _, r := range modFile.Require {
	//	fmt.Println(r)
	//	//fmt.Println(r.Mod.Path + "@" + r.Mod.Version)
	//}
	fmt.Println("------------------------")
	fmt.Println("replace如下：")
	fmt.Println("")
	for _, r := range modFile.Replace {
		fmt.Println("老：" + r.Old.Path + "@" + r.Old.Version)
		fmt.Println("新：" + r.New.Path + "@" + r.New.Version)
		fmt.Println("")
	}
	return
}
