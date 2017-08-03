package datastore

import (
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	src := []byte(`{"key":"val"}`)
	data, err := Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	val := data.Get("key")
	if val != "val" {
		t.Fatal(val, data)
	}

	// test ToJson
	out := data.ToJson()
	if out == nil{
		t.Fatal("to json fail")
	}

	fmt.Println(string(out))


	// test file
	data.Set("test", "test")
	if err := SaveDataToFile("test", data); err != nil{
	}
	data1, err := ParseDataFromFile("test")
	if err != nil{
		t.Fatal(err)
	}

	if data1.Get("test") != "test"{
		t.Fatal(data, data1)
	}

	fmt.Println(data1)
}
