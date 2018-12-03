package pachong

import (
	"github.com/noxue/transcode"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type httpClient struct {
	client  *http.Client
}

func newHttpClient() (request *httpClient) {
	client:=&http.Client{}
	request = &httpClient{
		client:client,
	}
	return
}

func (this *httpClient)get(site string) (b []byte,err error){
	defer func() {
		if e:=recover(); e!=nil{
			err = e.(error)
		}
	}()
	req, err := http.NewRequest("GET", site, nil)
	res,err:=this.client.Do(req)
	checkErr(err)

	defer res.Body.Close()

	b,err=ioutil.ReadAll(res.Body)
	checkErr(err)
	return
}

func (this *httpClient)getHtml(site string) (html string, err error) {
	defer func() {
		if e:=recover(); e!=nil{
			err = e.(error)
		}
	}()
	bs,err:= this.get(site)
	checkErr(err)
	html = string(bs)
	html = this.decode(html)
	return
}

func (this *httpClient)decode(content string)(html string){
	html = content
	// 获取网页编码
	re := regexp.MustCompile(`(?i)<meta.+?charset.*?=(.+?)[/'"><]`)
	arr:=re.FindStringSubmatch(html)
	// 如果没找到网页编码，就不转换
	if len(arr)!=2 {
		return
	}
	// 如果编码是gbk/gb2312就转换成utf-8
	arr[1] = strings.ToUpper(arr[1])
	if strings.Contains(arr[1],"GB") {
		html = transcode.FromString(html).Decode(arr[1]).ToString()
	}
	return
}