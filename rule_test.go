package gopa

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T)  {
	ruleStr := `[
	{"name":"title","selector":"(.*)","remove":"p","do":"#,http://noxue.com#"},
	{"name":"url","selector":"(.*)111","remove":"p1","do":"#,http://noxue.com11#"},
	{"name":"content","selector":"(.*)222","remove":"p2","do":"#,http://noxue.com22#"}
]`
	fmt.Println(parseRule(ruleStr))
}
