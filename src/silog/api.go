package silog

import "silog/logz"

func NewLogz(name string) Log {
	return logz.NewLogz(name)
}
