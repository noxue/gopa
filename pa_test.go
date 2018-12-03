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
	{"name":"text","selector":"(?i)<li[^>]+?><a href=\"[^\"]*?\" data-v-[^>]+?>([^>]*)?</a></li>"},
	{"name":"url","selector":"(?i)<li[^>]+?><a href=\"([^\"]*?)\" data-v-[^>]+?>[^>]*?</a></li>","do":"#,http://noxue.com#","replace":"","replaceText":"/aaaaaa"}
]
}`
	err := func()(err error){
		defer func() {
			if e:=recover(); e!=nil {
				err = e.(error)
			}
		}()
		pa:=NewPa()
		data:=pa.Get("https://noxue.com/a").Rules(rule).Do(func(data *Data) *Data {
			for i, _ := range data.data {
				data.data[i]["title"]+=" 不学网"
			}
			return data
		}).One()
		fmt.Println(data)
		return
	}()

	if err != nil {
		fmt.Println("出错了：",err)
	}
	fmt.Println("运行结束")
}
