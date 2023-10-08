package lib

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/mod/modfile"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// 查找某字符串值是否在切片中
func InSliceString(v string, m *[]string) bool {
	for _, value := range *m {
		if value == v {
			return true
		}
	}
	return false
}

// 查找某值是否在切片中
func InSliceInt64(v int64, m *[]int64) bool {
	for _, value := range *m {
		if value == v {
			return true
		}
	}
	return false
}

// 转换为大驼峰命名法则
// 首字母大写，“_” 忽略后大写
func StrFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}

// sha1加密
func Sha1En(data string) string {
	t := gsha1.Encrypt(data)
	return t
}

// 对字符串进行MD5哈希
func Md5En(data string) string {
	t, _ := gmd5.Encrypt(data)
	return t
}

// GetClientIp 获取客户端IP
func GetClientIp(ctx context.Context) string {
	return g.RequestFromCtx(ctx).GetClientIp()
}

func ParamStrToSlice(str, sep string) (res []int64) {
	arr := strings.Split(str, sep)
	for _, s := range arr {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			continue
		}
		res = append(res, i)
	}
	return
}

// 转换为大驼峰命名法则
// 首字母大写，“_” 忽略后大写
func StrFirstToUpperS(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}
func GetPrefixName(name string) string {
	if name == "" {
		return ""
	}
	temp := strings.Split(name, "_")
	return temp[0]
}

func GetPermiPath(name string) string {
	if name == "" {
		return ""
	}
	temp := strings.Split(name, "_")
	if len(temp) < 2 {
		return temp[0]
	}
	return temp[1]
}

// 读取go.mod module
func GetGodModule() string {
	goModFilePathData, err := os.ReadFile("go.mod")
	if err != nil {
		return ""
	}
	modFile, err := modfile.Parse("go.mod", goModFilePathData, nil)
	if err != nil {
		return ""
	}
	return modFile.Module.Mod.String()
}

// 格式化go文件
func FmtGoFile(path string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "go fmt "+path)
	} else {
		cmd = exec.Command("go fmt ", path)
	}
	if err := cmd.Start(); err != nil { // 运行命令
		log.Fatal(err)
	}
}

// 截取子串位置
func SubStr(str, substr string) string {
	if str == "" {
		return ""
	}
	n := UnicodeIndex(str, substr)
	if n < 0 {
		return str
	} else if n == 0 {
		return ""
	}
	s := []rune(str)
	return string(s[0:n])
}

// 子串在字符串的字节位置
func UnicodeIndex(str, substr string) int {
	result := strings.Index(str, substr)
	if result > 0 {
		prefix := []byte(str)[0:result]
		rs := []rune(string(prefix))
		result = len(rs)
	}
	return result
}
