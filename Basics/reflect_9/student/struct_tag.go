package student

import (
	"fmt"
	"reflect"

	// 编解码的包基本都在encoding
	"encoding/json"
)

// 结构体标签
type Student struct {
	Name   string `info:"name" doc:"我的名字"`
	Gender string `info:"gender"`
}

func FindTag(stu any) {
	t := reflect.TypeOf(stu)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("field:%v\n", t.Field(i).Tag.Get("info"))
	}
}

// 结构体标签的应用，json编码
// 结构体编码为json格式时会把key值替换为标签里的值
type Movie struct {
	Name   string   `json:"name"`
	Time   int      `json:"time"`
	Actors []string `json:"actors"`
}

func MovieInfo(movie any) []byte {
	JsonObj, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json编码失败")
		return nil
	}
	fmt.Printf("JsonObj:%s\n", JsonObj)
	return JsonObj
}

func MovieCollect(JsonByte any) (m Movie) {
	res := Movie{}
	value, ok := JsonByte.([]byte)
	if ok {
		err := json.Unmarshal(value, &res)
		if err != nil {
			fmt.Println("json解码失败")
			return
		}
	} else {
		fmt.Println("Argument should be Jsonbyte")
		return
	}
	fmt.Printf("struct:%v\n", res)
	return res
}
