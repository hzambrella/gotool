package datastore

import (
	"encoding/json"
	"io/ioutil"
)

type Data map[string]string

//
func Parse(src []byte) (Data, error) {
	d := make(Data)

	if err := json.Unmarshal(src, &d); err != nil {
		return nil, err
	}
	return d, nil
}

//
func (d Data) ToJson() []byte {
	result, err := json.MarshalIndent(d, "","	")
	if err != nil{
		panic(err)
	}
	return result
}

func (d Data) Get(key string) string {
	result, _ := d[key]
	return result
}

func (d Data) Set(key, val string) {
	d[key] = val
}

// read from file
func ParseDataFromFile(fileName string)(Data, error){
	data, err := ioutil.ReadFile(fileName)
	if err != nil{
		return nil, err
	}
	return Parse(data)
}

// save to file
func SaveDataToFile(fileName string, d Data) error{
	return ioutil.WriteFile(fileName, d.ToJson(), 0666)
}
