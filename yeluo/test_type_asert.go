// 创建人：lg
// 创建时间：2021/5/13
package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

// https://segmentfault.com/a/1190000039980197
func GetPtrValueV2(raw interface{}) (interface{}, error) {
	// 处理各种类型的nil
	vi := reflect.ValueOf(raw)
	if vi.IsNil() { // 判断值为nil
		return nil, nil
	}
	switch i := raw.(type) {
	case nil:
		return nil, nil
	case *string:
		return *i, nil
	case *int:
		return *i, nil
	case *int64:
		return *i, nil
	case *float32:
		return *i, nil
	case *float64:
		return *i, nil
	case *time.Time:
		return *i, nil
	default:
		return nil, errors.New("unsupported type")
	}
}

func main() {
	var (
		pi *int
	)
	for i:=0; i<10;i++ {

	}

	s := "hello"
	ps := &s
	vi, _ := GetPtrValueV2(pi)
	fmt.Println(vi)
	vs, _ := GetPtrValueV2(ps)
	fmt.Println(vs)
}
