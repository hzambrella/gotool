package adapter

import (
	"errors"
	"fmt"
	"silog/proto"
)

type Adapter interface {
	Put(proto proto.LogProto)
	Close()
}
type Adapters map[string]Adapter

var adapters = make(Adapters)

func GetAdapter(name string) (Adapter, error) {
	adapter, ok := adapters[name]
	if ok {
		return adapter, nil
	}
	return nil, errors.New("adapter  not found :" + name)
}

func Register(name string, adapter Adapter) {
	adapters[name] = adapter
}

// 处理适配器记录日志时产的错误，以println的方式输出至控制台。
// TODO: 将此错误写到系统日志当中。
func FailLog(log proto.LogProto, err error) {
	// log error
	failMsg := fmt.Sprintf("err:%s,len:%d", err.Error(), len(log.GetMsg()))
	fmt.Println(failMsg)
}
