package pachong

import (
	"encoding/json"
	"regexp"
	"strings"
)

type ruleNode struct {
	Name        string // 采集后组成json的名称
	Selector    string // 要采集的内容的正则
	Replace     string // 要替换的内容满足的正则
	ReplaceText string // 要替换的内容
	// 获取到的内容做处理，如：采集的相对路径变绝对路径
	// 例如：###,http://noxue.com###
	// 表示逗号之后的内容中包含逗号之前的内容的部分用采集到的内容替换
	// 比如：采集到的内容是 /2018/logo.jpg ,那么经过处理后的网址就是 http://noxue.com/2018/logo.jpg
	Do string
}

type ruleType struct {
	All   bool // 是否匹配全部，默认是简单模式
	Rules []ruleNode
}

func parseRule(ruleStr string) (rule ruleType) {
	err := json.Unmarshal([]byte(ruleStr), &rule)
	checkErr(err)
	return
}


// 把text内容 按照replace指定的正则匹配的所有内容替换成replaceText
func (this *Data) replace(text, replace, replaceText string) string {
	if replace == "" {
		return text
	}
	re, err := regexp.Compile(replace)
	checkErr(err)
	return re.ReplaceAllString(text, replaceText)
}

// 替换字符串，content是匹配到的内容，doStr是替换规则字符串
func (this *Data) do(content, doStr string) string {
	// 如果无需处理
	if len(doStr) == 0 {
		return content
	}
	arr := strings.SplitN(doStr, ",", 2)
	if len(arr) < 2 {
		return content
	}
	return strings.Replace(arr[1], arr[0], content, -1)
}
