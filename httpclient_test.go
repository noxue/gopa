package gopa

import (
	"fmt"
	"testing"
)

func TestGetHtml(t *testing.T) {
	hc := newHttpClient()
	html, err:=hc.getHtml("https://www.88dush.com/xiaoshuo/43/43849/11658052.html")
	checkErr(err)
	fmt.Println(html)
}
