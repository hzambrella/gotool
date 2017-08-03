package logz

import (
	"engine/datastore"
	"errors"
	"fmt"
	"os"
	"silog/adapter"
	"silog/adapter/console"
	"silog/proto"
	"strings"
	"time"
)

type Logz struct {
	Adapters   map[string]adapter.Adapter
	LoggerName string
	/*
		IsDebug    bool
		IsInfo     bool
		IsWarn     bool
		IsErr      bool
		IsFatal    bool
	*/
}

func NewLogz(logName string) *Logz {
	adapterNames, err := newLogAdapterRegister()
	if err != nil {
		panic(err)
	}

	a := make(map[string]adapter.Adapter, 0)
	for _, v := range adapterNames {
		adapter, err := adapter.GetAdapter(v)
		if err != nil {
			panic(err)
		}
		a[v] = adapter
	}
	return &Logz{a, logName}

}

func getAdaptersName() ([]string, error) {
	/*
	   {
	   	"adapters":"console,server"
	   }
	*/
	data, err := datastore.ParseDataFromFile(os.Getenv("LOGCFG") + "/log.cfg")
	if err != nil {
		return nil, err
	}

	adapterStr, ok := data["adapters"]
	if !ok || len(adapterStr) == 0 {
		return nil, errors.New("配置文件错误")
	}

	a := strings.Split(adapterStr, ",")
	return a, nil
}

func newLogAdapterRegister() ([]string, error) {
	adapterNames, err := getAdaptersName()
	if err != nil {
		return nil, err
	}

	for _, name := range adapterNames {
		switch name {
		case console.AdapterName:
			adapter.Register(console.AdapterName, console.New())
			/*
				case server.AdapterName:
					adapter.Register(server.AdapterName, server.New(etc.NewEtcByCfg(cfg)))
				}
			*/
		}
	}
	return adapterNames, nil
}

func (l *Logz) Debug(t ...interface{}) {
	p := &proto.DefaultProto{
		LoggerName: l.LoggerName,
		Level:      proto.LevelDebug,
		Time:       time.Now(),
	}
	p.ToMsg(t)
	l.Put(p)
}

func (l *Logz) Info(t ...interface{}) {
	p := &proto.DefaultProto{
		LoggerName: l.LoggerName,
		Level:      proto.LevelInfo,
		Time:       time.Now(),
	}
	p.ToMsg(t)
	l.Put(p)
}

func (l *Logz) Warn(t ...interface{}) {
	p := &proto.DefaultProto{
		LoggerName: l.LoggerName,
		Level:      proto.LevelWarn,
		Time:       time.Now(),
	}
	p.ToMsg(t)
	l.Put(p)
}

func (l *Logz) Error(t ...interface{}) {
	p := &proto.DefaultProto{
		LoggerName: l.LoggerName,
		Level:      proto.LevelError,
		Time:       time.Now(),
	}
	p.ToMsg(t)
	l.Put(p)
}

func (l *Logz) Fatal(t ...interface{}) {
	p := &proto.DefaultProto{
		LoggerName: l.LoggerName,
		Level:      proto.LevelFatal,
		Time:       time.Now(),
	}
	p.ToMsg(t)
	l.Put(p)
	l.Close()
	panic(p)
}

func (l *Logz) Put(p *proto.DefaultProto) {
	if p == nil {
		panic("logProto is nil")
	}
	for _, v := range l.Adapters {
		v.Put(p)
	}
}

func (l *Logz) Close() {
	for _, v := range l.Adapters {
		v.Close()
	}
	fmt.Println("logz closed!")
}

// Exit
// log an info level message, and close log, then call os.Exit(code)
//
// Param
// code -- code of os exit
// msg -- exit message
func (l *Logz) Exit(code int, t ...interface{}) {
	p := &proto.DefaultProto{
		LoggerName: l.LoggerName,
		Level:      proto.LevelInfo,
	}
	p.ToMsg(t)
	l.Put(p)
	l.Close()
	os.Exit(code)
}
