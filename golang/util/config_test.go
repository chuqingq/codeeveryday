package util

import (
	"log"
	"os"
	"testing"

	"github.com/bmizerany/assert"
)

const storeFile = "tmp_test.store"
const defaultStoreFile = "tmp_default.store"

// TestNotExistConfig
func TestNotExistConfig(t *testing.T) {
	log.Printf("TestConfig")
	// not exists config file
	_, err := NewConfig(storeFile)
	assert.NotEqual(t, err, nil)
}

// TestResetConfig reset from defaultStoreFile
func TestResetConfig(t *testing.T) {
	err := os.WriteFile(storeFile, []byte(`{}`), 0666)
	assert.Equal(t, err, nil)
	defer os.Remove(storeFile)

	c, err := NewConfig(storeFile)
	assert.Equal(t, err, nil)

	namevalue := c.Get("name").MustString()
	assert.Equal(t, namevalue, "")

	err = os.WriteFile(defaultStoreFile, []byte(`{"name": "myname"}`), 0666)
	assert.Equal(t, err, nil)
	defer os.Remove(defaultStoreFile)

	err = c.ResetFrom(defaultStoreFile)
	assert.Equal(t, err, nil)

	namevalue1 := c.Get("name").MustString()
	assert.Equal(t, namevalue1, "myname")
}

// Set, Get, GetDefaultInt, GetDefaultString, Remove
func TestConfig(t *testing.T) {
	err := os.WriteFile(storeFile, []byte(`{
			"name":"myname",
			"col1": {
				"key1": "value1",
				"structint": 2
			}
		}`), 0666)
	assert.Equal(t, err, nil)
	defer os.Remove(storeFile)

	c, err := NewConfig(storeFile)
	assert.Equal(t, err, nil)

	// get exit key
	namevalue := c.GetPath("name").MustString()
	assert.Equal(t, namevalue, "myname")

	// get not exist key
	value := c.GetPath("col1.key1").MustString()
	assert.Equal(t, value, "value1")

	// set key
	const SET_VALUE = "value2"
	c.SetPath("col1.key1", SET_VALUE)
	// c.GetPath("col1.key1").Set(SET_VALUE)
	// get exist key
	value = c.GetPath("col1.key1").MustString()
	assert.Equal(t, value, SET_VALUE)

	// replace key by GetDefaultInt
	c.SetPath("col1.key1", 10)
	valueInt := c.GetPath("col1.key1").MustInt(-1)
	assert.Equal(t, valueInt, 10)

	// open another
	// time.Sleep(time.Second * 20)
	conf2, err := NewConfig(storeFile)
	assert.Equal(t, err, nil)

	valueInt = conf2.GetPath("col1.key1").MustInt(-1)
	assert.Equal(t, valueInt, 10)
}
