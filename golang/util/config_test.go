package util

import (
	"log"
	"testing"
)

const storeFile = "test.store"
const defaultStoreFile = "default.store"

// Set, Get, GetDefaultInt, GetDefaultString, Remove
func TestConfig(t *testing.T) {
	log.Printf("TestConfig")
	conf, err := NewConfig(storeFile, defaultStoreFile)
	if err != nil {
		t.Fatalf("NewStore error: %v", err)
	}
	defer conf.Remove()
	// get not exist key
	value := conf.GetPath("col1.key1").MustString("12")
	if value != "12" {
		t.Fatalf("GetDefaultString error")
	}
	// set key
	conf.SetPath("col1.key1", "value1")
	// exist key
	value = conf.GetPath("col1.key1").MustString("34")
	if value != "value1" {
		t.Fatalf("store.Get(exist key) error: %v", value)
	}
	// replace key
	conf.SetPath("col1.key1", "value2")
	value = conf.GetPath("col1.key1").MustString("56")
	if value != "value2" {
		t.Fatalf("store.Get(replaced key) error: %v", value)
	}
	// replace key by GetDefaultInt
	conf.SetPath("col1.key1", 10)
	valueInt := conf.GetPath("col1.key1").MustInt(-1)
	if valueInt != 10 {
		t.Fatalf("store.Get(replaced key) error: %v", valueInt)
	}
	// open another
	conf2, err := NewConfig(storeFile, defaultStoreFile)
	if err != nil {
		t.Fatalf("NewStore(another store) error: %v", err)
	}
	valueInt = conf2.GetPath("col1.key1").MustInt(-1)
	if valueInt != 10 {
		t.Fatalf("store.Get(replaced key) error: %v", valueInt)
	}
	log.Printf("TestStore success")
}

// SetJson, GetJson
// func TestStoreJson(t *testing.T) {
// 	store, err := NewStore(storeFile, defaultStoreFile)
// 	if err != nil {
// 		t.Fatalf("NewStore error: %v", err)
// 	}
// 	defer store.Remove()
// 	err = store.SetJSON("col1.key1", `{"key2":128}`)
// 	if err != nil {
// 		t.Fatalf("SetJSON error: %v", err)
// 	}
// 	// Get
// 	v2 := store.Get("col1.key1.key2")
// 	if v2.MustInt(0) != 128 {
// 		t.Fatalf("Get.MustInt error: %v", v2.MustInt(0))
// 	}
// 	// GetDefaultInt
// 	v := store.GetDefaultInt("col1.key1.key2", -1)
// 	if v != 128 {
// 		t.Fatalf("GetDefaultInt after SetJSON error: %v", v)
// 	}
// 	// GetJSOM
// 	v1 := store.GetJSON("col1.key1")
// 	log.Printf("GetJSON: %v", v1)
// }

func TestConfigReset(t *testing.T) {
	store, err := NewConfig(storeFile, defaultStoreFile)
	if err != nil {
		t.Fatalf("newStore error: %v", err)
	}
	defer store.Remove()
	// Set
	store.SetPath("IMCM.Net.ComPort", "123")
	// Get
	v := store.GetPath("IMCM.Net.ComPort").MustString("default")
	if v != "123" {
		t.Fatalf("GetDefaultString error: %v, should be %v", v, "123")
	}
	// Reset
	err = store.Reset()
	if err != nil {
		t.Fatalf("Reset error: %v", err)
	}
	// Get
	v2 := store.GetPath("IMCM.Net.ComPort").MustString("default2")
	if v2 != "COM1" {
		t.Fatalf("GetDefaultString error: %v, should be %v", v, "COM1")
	}
}
