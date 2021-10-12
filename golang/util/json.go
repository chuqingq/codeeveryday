package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	sjson "github.com/bitly/go-simplejson"
)

// Message comm使用的消息
type Json struct {
	sjson.Json
}

func NewJson() *Json {
	return &Json{}
}

// Set 支持v是string、bool、int、map[string]interface{}、[]interface{}
func (m *Json) SetPath(path string, v interface{}) {
	m.Json.SetPath(strings.Split(path, "."), v)
}

func (m *Json) GetPath(path string) *Json {
	if path == "" {
		return m
	}
	return &Json{*m.Json.GetPath(strings.Split(path, ".")...)}
}

// func (j *Json) String() (string, error)
// func (j *Json) MustString(args ...string) string

// func (j *Json) Bool() (bool, error)
// func (j *Json) MustBool(args ...bool) bool

// func (j *Json) Int() (int, error)
// func (j *Json) MustInt(args ...int) int

func (m *Json) Array() ([]Json, error) {
	arr, err := m.Json.Array()
	if arr == nil || err != nil {
		return nil, err
	}
	marray := make([]Json, len(arr), len(arr))
	for i, a := range arr {
		marray[i].Json.SetPath([]string{}, a)
	}
	return marray, nil
}

func (m *Json) MustArray() []Json {
	arr, err := m.Array()
	if err != nil {
		return nil
	}
	return arr
}

// func (j *Json) Map() (map[string]interface{}, error)
// func (j *Json) MustMap(args ...map[string]interface{}) map[string]interface{}
// TODO 是否要提供一个map[string]Json的方法？

// 其他类型参考：https://pkg.go.dev/github.com/bitly/go-simplejson

// Unmarshal 把m解析到v上。类似json.Unmarshal()
func (m *Json) Unmarshal(v interface{}) error {
	b := m.ToBytes()
	return json.Unmarshal(b, v)
}

func (m *Json) ToInterface(v interface{}) error {
	return m.Unmarshal(v)
}

// ToBytes Message转成[]byte
func (m *Json) ToBytes() []byte {
	b, err := m.EncodePretty()
	if err != nil {
		log.Printf("messagep[%v].EncodePretty() error: %v", m, err)
	}
	return b
}

// ToString Message转成string
func (m *Json) ToString() string {
	return string(m.ToBytes())
}

// JsonFromBytes 字节数组转成Message
func JsonFromBytes(data []byte) (*Json, error) {
	m, err := sjson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return &Json{*m}, nil
}

// JsonFromString 字符串转成Message
func JsonFromString(s string) (*Json, error) {
	m, err := JsonFromBytes([]byte(s))
	return m, err
}

// JsonFromInterface
func JsonFromInterface(v interface{}) *Json {
	str := InterfaceToJsonString(v)
	m, _ := JsonFromString(str)
	return m
}

// JsonFromFile 从filepath读取Message
func JsonFromFile(filepath string) (*Json, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	m, err := JsonFromBytes(b)
	return m, err
}

// ToFile 把Message保存到filepath
func (m *Json) ToFile(filepath string) error {
	b, err := m.EncodePretty()
	if err != nil {
		return err
	}
	const defaultFileMode = 0644
	return ioutil.WriteFile(filepath, b, defaultFileMode)
}

// interface{}和json string之间的转换

func InterfaceToJsonString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func InterfaceFromJsonString(s string, v interface{}) error {
	err := json.Unmarshal([]byte(s), v)
	if err != nil {
		return err
	}
	return nil
}
