package console

import (
	"fmt"
	"silog/adapter"
	"silog/proto"
)

const (
	AdapterName = "console"
)

type consoleAdapter struct {
	buffer chan proto.LogProto
	exit   chan bool
	end    chan bool
}

// put a log protocol to log queue
func (conA *consoleAdapter) Put(log proto.LogProto) {
	if log == nil {
		panic("argument is nil")
	}
	conA.buffer <- log
}

func (conA *consoleAdapter) Close() {
	conA.exit <- true

	// wait for log flush.
	<-conA.end
}

func (conA *consoleAdapter) run() {
	isExit := false
	for {
		// if adapter is closed and flush done, exit running.
		if isExit && len(conA.buffer) == 0 {
			conA.end <- true
			fmt.Println("console adapter is closed!")
			return
			//break
		}

		select {
		case log := <-conA.buffer:
			fmt.Println(log)
		case <-conA.exit:
			isExit = true
		}
	}
}

func New() adapter.Adapter {
	cAdapter := &consoleAdapter{
		make(chan proto.LogProto, 100),
		make(chan bool, 1),
		make(chan bool),
	}
	go cAdapter.run()
	fmt.Println("you have new a console log")
	return cAdapter
}
