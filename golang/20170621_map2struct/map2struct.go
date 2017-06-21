package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

type myConfig struct {
	ListenIP   string `cfg:"listenip"`
	ListenPort int    `cfg:"listenport"`
	Keepalive  struct {
		Idle  int `cfg:"idle"`
		Count int `cfg:"count"`
		Abc   struct {
			AbcValue string `cfg:"abcvalue"`
		} `cfg:"abc"`
	} `cfg:"keepalive"`
}

func main() {
	cfg := &myConfig{}
	values := map[string]string{
		"listenip":               "127.0.0.1",
		"listenport":             "1234",
		"keepalive/idle":         "123",
		"keepalive/count":        "456",
		"keepalive/abc/abcvalue": "abcvalue",
	}

	Unmarshal(values, cfg)
	fmt.Printf("cfg: %+v\n", cfg)
}

const defaultTag = "cfg"

func Unmarshal(values map[string]string, v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New(fmt.Sprintf("nil or not ptr"))
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("only accept struct; got %T", v))
	}

	return setValue("", values, &rv)
}

func setValue(prefix string, values map[string]string, value *reflect.Value) (err error) {
	fmt.Printf("value prefix: %v\n", prefix)
	tv := value.Type()
	for i := 0; i < value.NumField(); i++ {
		if tagv := tv.Field(i).Tag.Get(defaultTag); tagv != "" {
			fmt.Printf("tagv: %v\n", tagv)
			field := value.Field(i)
			if field.Kind() == reflect.Struct {
				if err := setValue(prefix+tagv+"/", values, &field); err != nil {
					return err
				}
			} else {
				key := prefix + tagv
				value, ok := values[key]
				if !ok {
					continue
				}

				switch field.Kind() {
				case reflect.String:
					field.SetString(value)
				case reflect.Int64:
					i, _ := strconv.Atoi(value)
					field.SetInt(int64(i))
				case reflect.Int:
					i, _ := strconv.Atoi(value)
					field.SetInt(int64(i))
				}
			}
		}
	}
	return nil
}

// func ConvertConfig(values map[string]interface{}, cfg interface{}) error {
// 	jsonBytes, err := json.Marshal(values)
// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(jsonBytes, cfg)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
