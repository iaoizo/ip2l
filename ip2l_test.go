package ip2lo

import (
	"fmt"
	"testing"
)

func Test_parse(t *testing.T) {
	//测试
	ip2L := NewIp2L()
	parse, err := ip2L.Parse("218.88.126.63")
	if err != nil {
		return
	}
	fmt.Printf("%#v", parse)
}
