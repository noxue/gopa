package gopa

import (
	"fmt"
	"testing"
)

func TestData_Rules(t *testing.T) {

	rule:=`{
	"all":true,
	"rules":[
	{"name":"title","selector":"<title.*?>(.*)</title>"},
	{"name":"text","selector":"(?i)<li[^>]+?><a href=\"[^\"]*?\" Data-v-[^>]+?>([^>]*)?</a></li>"},
	{"name":"url","selector":"(?i)<li[^>]+?><a href=\"([^\"]*?)\" Data-v-[^>]+?>[^>]*?</a></li>","do":"#,http://noxue.com#","replace":"","replaceText":"/aaaaaa"}
]
}`

	err := func()(err error){
		// 类似java的try catch机制
		defer func() {
			if e:=recover(); e!=nil {
				err = e.(error)
			}
		}()

		data:=NewPa().Get("https://noxue.com/a").Rules(rule).Do(func(data *Data) *Data {
			for i, _ := range data.Data {
				data.Data[i]["title"]+=" 不学网"
			}
			return data
		})
		fmt.Println(data.All())
		fmt.Println(data.One())
		return
	}()

	if err != nil {
		fmt.Println("出错了：",err)
	}
	fmt.Println("运行结束")
}
