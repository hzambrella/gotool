package main

import (
	"fmt"
	"loghz"

	"github.com/Unknwon/goconfig"
)

func main() {
	c, err := goconfig.LoadConfigFile("db.cfg")
	if err != nil {
		fmt.Println(err)
	}
	logz := loghz.NewLogDebug(true)
	logz.Println(*c)

	v, err := c.GetValue("master", "dsn")
	if err != nil {
		fmt.Println(err)
	}
	logz.Println(v)

}
