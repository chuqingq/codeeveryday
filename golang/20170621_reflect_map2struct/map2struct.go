package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strconv"

	"./mypackage"
)

func main() {
	cfg := &mypackage.MyConfig{}
	values := map[string]string{
		"listenip":               "127.0.0.1",
		"listenport":             "1234",
		"keepalive/idle":         "123",
		"keepalive/count":        "456",
		"keepalive/abc/abcvalue": "abcvalue",
	}

	err := Unmarshal(values, cfg)
	fmt.Printf("cfg: %+v, err: %v\n", cfg, err)
}

const defaultTag = "cfg"

func Unmarshal(values map[string]string, v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				fmt.Printf("runtime.Error panic\n")
				panic(r)
			}
			if err2, ok := r.(error); ok {
				err = err2
			} else {
				err = errors.New(r.(string))
			}
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

	return setValue("", values, rv)
}

func setValue(prefix string, values map[string]string, value reflect.Value) (err error) {
	fmt.Printf("value prefix: %v, %v\n", prefix, value.IsValid())
	tv := value.Type()
	for i := 0; i < value.NumField(); i++ {
		if tagv := tv.Field(i).Tag.Get(defaultTag); tagv != "" {
			fmt.Printf("tagv: %v\n", tagv)
			field := value.Field(i)
			fmt.Printf("field kind: %v\n", field.Kind())
			if field.Kind() == reflect.Struct {
				if err := setValue(prefix+tagv+"/", values, field); err != nil {
					return err
				}
			} else {
				if !field.CanSet() {
					continue
				}
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

