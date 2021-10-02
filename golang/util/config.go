package util

import (
	"os"
)

type Config struct {
	*Json
	filePath        string
	defaultFilePath string
}

func NewConfig(filePath, defaultFilePath string) (*Config, error) {
	m, err := JsonFromFile(filePath)
	if err != nil {
		m, err = JsonFromFile(defaultFilePath)
		if err != nil {
			return nil, err
		}
		m.ToFile(filePath)
	}
	return &Config{
		Json:            m,
		filePath:        filePath,
		defaultFilePath: defaultFilePath,
	}, nil
}

// 恢复默认设置
func (c *Config) Reset() error {
	var err error
	c.Json, err = JsonFromFile(c.defaultFilePath)
	if err != nil {
		return err
	}
	c.Save()
	return nil
}

// // configSave 模块级配置保存
// func configSave(path string, val interface{}) error {
// 	str := util.ToJson(val)
// 	util.D().Printf("configSave(%v) %v", path, str)
// 	msg, _ := util.MessageFromString(str)
// 	defaultConfig.Set(path, msg.Map("", nil)) // TODO
// 	return nil
// }

func (c *Config) SaveInterface(path string, v interface{}) {
	m := JsonFromInterface(v)
	c.Set(path, m)
}

// // configLoad 模块级配置加载
// func configLoad(path string, val interface{}) {
// 	log.Printf("defaultConfig: %v", defaultConfig)
// 	msg := defaultConfig.Get(path)
// 	msg.Unmarshal(val)
// }

/*
// 等同于GetPath() + ToInterface()
func (c *Config) LoadStruct(path string, v interface{}) {
	m := c.Get(path)
	m.ToInterface(v)
}
*/

// Set 设置值。Message.Set + save。v支持string/int/bool等，如果是复合值，需要是Map/MustMap()
func (c *Config) SetPath(path string, v interface{}) {
	c.Json.SetPath(path, v)
	c.Save()
}

// 说明：读取直接使用Message的String/Int/Bool等

func (c *Config) Save() error {
	return c.ToFile(c.filePath)
}

func (c *Config) Remove() {
	os.Remove(c.filePath)
}
