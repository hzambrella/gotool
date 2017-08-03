package proto

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type LogProto interface {
	ToMsg(i ...interface{})
	String() string
	GetMsg() []byte
	GetLoggerName() string
	GetTime() time.Time
	GetLevel() Level
	SetFormat(loggerNameFormat, timeFormat, levelFormat string)
	GetMoreInfo() map[string]interface{}
}

//const DefaultProteName = "default_logproto"
const (
	LevelDebug = Level(0)
	LevelInfo  = Level(1)
	LevelWarn  = Level(2)
	LevelError = Level(3)
	LevelFatal = Level(4)
)

type Level int

func (l Level) Int() int {
	return int(l)
}

func (l Level) String() string {
	//	fmt.Println(l)
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	}
	return strconv.Itoa(int(l))
}

type DefaultProto struct {
	LoggerName string
	Level      Level
	Time       time.Time
	Msg        []byte
	Format     []string
}

func (d *DefaultProto) ToMsg(i ...interface{}) {
	d.Msg = []byte(fmt.Sprint(i...))
}

func (d *DefaultProto) GetMsg() []byte {
	return d.Msg
}

func (d *DefaultProto) GetLoggerName() string {
	return d.LoggerName
}
func (d *DefaultProto) GetLevel() Level {
	return d.Level
}

func (d *DefaultProto) GetTime() time.Time {
	return d.Time
}

func (d *DefaultProto) SetFormat(loggerNameFormat, timeFormat, levelFormat string) {
	s := make([]string, 3)
	if len(loggerNameFormat) == 0 {
		loggerNameFormat = "[--%s--]"
	}

	// exp:2006-01-02
	if len(timeFormat) == 0 {
		timeFormat = "2006-01-02"
	}

	if len(levelFormat) == 0 {
		levelFormat = "[%s]"
	}

	s[0] = loggerNameFormat
	s[1] = timeFormat
	s[2] = levelFormat

	d.Format = s
}

func (d *DefaultProto) GetMoreInfo() map[string]interface{} {
	return nil
}

func (d *DefaultProto) String() string {

	if d.Format == nil {
		s := make([]string, 3)
		loggerNameFormat := "[--%s--]"
		timeFormat := "2006-01-02"
		levelFormat := "[%s]"
		s[0] = loggerNameFormat
		s[1] = timeFormat
		s[2] = levelFormat

		d.Format = s
	}

	s := d.Format
	loggerName := fmt.Sprintf(s[0], d.LoggerName)
	timeStr := d.Time.Format(s[1])
	levelStr := fmt.Sprintf(s[2], d.Level.String())
	return fmt.Sprintf("%s %s %s %s\n ", loggerName, timeStr, levelStr, strings.TrimSuffix(strings.TrimPrefix(string(d.Msg), "["), "]"))
}
