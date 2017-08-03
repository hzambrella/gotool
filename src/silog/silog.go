package silog

func init() {
}

/*
const (
	TIME_FORMAT = "2006-01-02 15:04"
)
*/

type Log interface {
	Debug(t ...interface{})
	Info(t ...interface{})
	Warn(t ...interface{})
	Error(t ...interface{})
	Fatal(t ...interface{})
	// 关闭日志器
	// 当不再使用日志时，应调用关闭，以使用日志的缓冲得到输出。
	Close()

	// 退出操作，执行关闭日志操作并调用os.Exit
	// 对于可执行程序来说，日志的退出意味着程序的退出。
	//
	// 参数
	//  code -- 退出码值，由os.Exit调用
	//  msg -- 记录的消息，级别是一个Info级别.
	//
	Exit(code int, msg ...interface{})
}
