package gopa

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"regexp"
)

type DoFunc func(data *Data, value string) string

type Pa struct {
	hc   *httpClient
	data *Data
}

type Data struct {
	b    []byte
	Url  string // 当前请求的地址
	Html string
	Data []map[string]string
	Rule Rule
	// 根据do指定字符串执行指定函数，比如：用户传入自动下载图片的函数
	doFuncs map[string]DoFunc
}

func NewPa() *Pa {
	return &Pa{
		hc: newHttpClient(),
	}
}

func (this *Data) Download(savePath string) {
	err := ioutil.WriteFile(savePath, this.b, 0666)
	checkErr(err)
}

func (this *Pa) Get(site string) *Data {
	b, err := this.hc.get(site)
	checkErr(err)
	this.data = &Data{
		doFuncs:make(map[string]DoFunc),
		Url:  site,
		b:    b,
		Html: this.hc.decode(string(b)),
	}
	return this.data
}

func (this *Data) ToBytes() (bs []byte) {
	bs = this.b
	return
}

func (this *Data) ToString() (html string) {
	html = this.Html
	return
}

// 设置函数，执行规则中do指定的字符串与name对应的函数
func (this *Data) SetDoFunc(name string, doFunc DoFunc) *Data {
	if name==""{
		panic(errors.New("设置的DoFunc函数为指定Name"))
	}
	this.doFuncs[name] = doFunc
	return this
}

func (this *Data) Rules(rule Rule) (*Data) {
	// 如果有被调用过，清空数据，不影响本次结果，让rules函数可以多次调用
	if len(this.Data) != 0 {
		this.Data = *&[]map[string]string{}
	}
	this.Rule = rule
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
				checkErr(errors.New(r.Name+"规则没匹配到内容，请检查规则是否正确"))
			}
			if r.SubSelector != "" {
				ss = this.subSelector(ss[1], r.SubSelector, false)[0]
			}
			this.Data[0][r.Name] = this.do(ss[1], r.Do)
			for _, v1 := range r.Replace {
				this.Data[0][r.Name] = this.replace(this.Data[0][r.Name], v1.Selector, v1.Text)
			}
			continue
		}

		//下面是匹配全部
		ss := re.FindAllStringSubmatch(content, -1)

		if len(ss) == 0 || len(ss[0]) != 2 {
			checkErr(errors.New(r.Name + "规则出错，正则没获取到内容，请检查：" + r.SubSelector))
		}

		if r.SubSelector != "" {
			ss = this.subSelector(ss[0][1], r.SubSelector, true)
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
			for _, v1 := range r.Replace {
				this.Data[i][r.Name] = this.replace(this.Data[i][r.Name], v1.Selector, v1.Text)
			}
		}
	}
	return this
}

func (this *Data) subSelector(text string, subSelector string, all bool) (ret [][]string) {

	re, err := regexp.Compile(subSelector)
	checkErr(err)
	if all {
		ret = re.FindAllStringSubmatch(text, -1)
		return
	}
	ret = append(ret, re.FindStringSubmatch(text))
	return
}

func (this *Data) Do(doFunc DoFunc,value string) string {
	return doFunc(this,value)
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
func (this *Data) One() map[string]string {
	if len(this.Data) == 0 {
		panic(errors.New("内容采集失败，没有任何结果"))
	}
	return this.Data[0]
}

// 返回所有采集到的结果
func (this *Data) All() []map[string]string {
	return this.Data
}
