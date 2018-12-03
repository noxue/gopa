package gopa

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"regexp"
)

type DoFunc func(data *Data) *Data

type Pa struct {
	hc   *httpClient
	data Data
}

type Data struct {
	b    []byte
	Html string
	Data []map[string]string
	Rule ruleType
}

func NewPa() *Pa {
	return &Pa{
		hc: newHttpClient(),
	}
}

func (this *Data)Download(savePath string) {
	err:=ioutil.WriteFile(savePath,this.b,0666)
	checkErr(err)
}

func (this *Pa) Get(site string) (data *Data) {
	b, err := this.hc.get(site)
	checkErr(err)
	data = &Data{
		b:    b,
		Html: this.hc.decode(string(b)),
	}
	return
}

func (this *Data) ToBytes() (bs []byte) {
	bs = this.b
	return
}

func (this *Data) ToString() (html string) {
	html = this.Html
	return
}

func (this *Data) Rules(ruleStr string) (*Data) {

	this.Rule = parseRule(ruleStr)
	for _, r := range this.Rule.Rules {

		// name不能为空
		if r.Name == "" {
			panic(errors.New("规则中缺少name，请检查规则是否错误"))
		}

		re, err := regexp.Compile(r.Selector)
		checkErr(err)
		content := this.Html

		// 如果不匹配全部
		if this.Rule.All == false {
			if len(this.Data) == 0 {
				this.Data = append(this.Data, make(map[string]string))
			}
			ss := re.FindStringSubmatch(content)
			if len(ss) != 2 {
				this.Data[0][r.Name] = ""
			}
			this.Data[0][r.Name] = this.do(ss[1], r.Do)
			this.Data[0][r.Name] = this.replace(this.Data[0][r.Name], r.Replace, r.ReplaceText)
			continue
		}

		//下面是匹配全部
		ss := re.FindAllStringSubmatch(content, -1)

		if len(ss) == 0 {
			continue
		}

		if len(this.Data) == 0 {
			for i := 0; i < len(ss); i++ {
				this.Data = append(this.Data, make(map[string]string))
			}
		}

		l := len(this.Data)
		for i, v := range ss {
			if len(v) != 2 {
				continue
			}
			// 如果大小不够，就添加一个，防止越界
			if i > l-1 {
				this.Data = append(this.Data, make(map[string]string))
			}
			this.Data[i][r.Name] = this.do(v[1], r.Do)
			this.Data[i][r.Name] = this.replace(this.Data[i][r.Name], r.Replace, r.ReplaceText)
		}
	}
	return this
}

func (this *Data) Do(doFunc DoFunc) *Data {
	return doFunc(this)
}

func (this *Data) String() string {
	str := "["
	for i, v := range this.Data {
		str += fmt.Sprintf("\n\t[%d] =>(\n", i)
		for k, v1 := range v {
			str += fmt.Sprintf("\t\t[%s] => %s,\n", k, v1)
		}
		str += "\t),\n"
	}
	str += "]"
	return str
}

// 返回第一条采集到的结果
func (this *Data) One() map[string]string{
	if len(this.Data) == 0{
		panic(errors.New("内容采集失败，没有任何结果"))
	}
	return this.Data[0]
}

// 返回所有采集到的结果
func (this *Data) All() []map[string]string{
	return this.Data
}
