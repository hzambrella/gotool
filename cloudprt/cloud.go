package cloudprt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"unicode"
)

const (
	printContentUrl = `http://open.printcenter.cn:8080/addOrder`
	queryOrderUrl   = `http://open.printcenter.cn:8080/queryOrder`
	getPrtStatusUrl = `http://open.printcenter.cn:8080/queryPrinterStatus`
)

// 打印机类型
const (
	//s1打印机
	S1 = iota
	//usb打印机
	USB
)

var (
	ErrWrongType = errors.New("invalid prt type")
)

type Prt struct {
	// 打印机类型
	// 0- S1打印机
	// 1- USB打印机
	Type int `json:"type"`
	// 打印机编号
	DeviceNo string `json:"deviceNo"`
	// 密钥
	Key string `json:"key"`
}

// 响应
type PrtResp struct {
	//响应码
	ResponseCode string `json:"responseCode"`
	// 消息
	Msg string `json:"msg"`
	// 订单索引
	OrderIndex string `json:"orderindex"`
}

type QueryResp struct {
	//响应码
	ResponseCode string `json:"responseCode"`
	// 消息
	Msg string `json:"msg"`
}

func NewPrt(prtType int, deviceNo, key string) (*Prt, error) {
	if prtType < 0 || prtType > 1 {
		return nil, ErrWrongType
	}
	return &Prt{Type: prtType, DeviceNo: deviceNo, Key: key}, nil
}

// 接口1：打印内容

/*
----------S1小票机返回的结果有如下几种：----------
{"responseCode":0,"msg":"订单添加成功，打印完成","orderindex":"xxxxxxxxxxxxxxxxxx"}
{"responseCode":1,"msg":"订单添加成功，正在打印中","orderindex":"xxxxxxxxxxxxxxxxxx"}
{"responseCode":2,"msg":"订单添加成功，但是打印机缺纸，无法打印","orderindex":"xxxxxxxxxxxxxxxxxx"}
{"responseCode":3,"msg":"订单添加成功，但是打印机不在线","orderindex":"xxxxxxxxxxxxxxxxxx"}
----------以上情况无须再次发送订单;下面的情况需要进行错误处理----------
{"responseCode":10,"msg":"内部服务器错误；"}
{"responseCode":11,"msg":"参数不正确；"}
{"responseCode":12,"msg":"打印机未添加到服务器；"}
{"responseCode":13,"msg":"未添加为订单服务器；"}
{"responseCode":14,"msg":"订单服务器和打印机不在同一个组；"}
{"responseCode":15,"msg":"订单已经存在，不能再次打印；"}
*/

// 打印
// content:打印内容
// 返回Resp。若code>=10,会返回错误,需要处理
// code<10，缺纸打印机没开之类的，无需重新发送订单。
func (prt *Prt) PrintString(content string) (*PrtResp, error) {
	response, err := http.PostForm(printContentUrl,
		url.Values{"deviceNo": {prt.DeviceNo}, "key": {prt.Key}, "printContent": {content}, "times": {strconv.Itoa(1)}})

	if err != nil {
		return nil, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		if body == nil {
			return nil, errors.New("resp body is nil")
		}

		var resp PrtResp
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, err
		}

		code, err := strconv.Atoi(resp.ResponseCode)
		if err != nil {
			return nil, err
		}
		var minErrCode int

		switch prt.Type {
		case S1:
			minErrCode = 10
		case USB:
			minErrCode = 2
		default:
			return nil, ErrWrongType
		}

		if code >= minErrCode {
			return &resp, errors.New(resp.Msg)
		} else {
			return &resp, nil
		}
	}
}

// 查询订单打印的状态
// 参数
// orderindex -- 打印的订单号，在PrintContent后返回
func (prt *Prt) QueryOrder(orderindex string) (*QueryResp, error) {
	response, err := http.PostForm(queryOrderUrl,
		url.Values{"deviceNo": {prt.DeviceNo}, "key": {prt.Key}, "orderindex": {orderindex}})

	if err != nil {
		return nil, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		if body == nil {
			return nil, errors.New("resp body is nil")
		}

		var resp QueryResp
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, err
		}

		code, err := strconv.Atoi(resp.ResponseCode)
		if err != nil {
			return nil, err
		}

		minErrCode := 2
		if code >= minErrCode {
			return &resp, errors.New(resp.Msg)
		} else {
			return &resp, nil
		}
	}
}

// 查询打印机状态
func (prt *Prt) GetPrtStatus() (*QueryResp, error) {
	response, err := http.PostForm(getPrtStatusUrl,
		url.Values{"deviceNo": {prt.DeviceNo}, "key": {prt.Key}})

	if err != nil {
		return nil, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		if body == nil {
			return nil, errors.New("body is nil")
		}

		var resp QueryResp
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, err
		}

		code, err := strconv.Atoi(resp.ResponseCode)
		if err != nil {
			return nil, err
		}
		minErrCode := 2
		if code >= minErrCode {
			return &resp, errors.New(resp.Msg)
		} else {
			return &resp, nil
		}
	}
}

// 对小票的商品部分格式化
//for ( GOODS.... ){
//formatName(buf,name,price,num,total)
//}
func formatName(buf *bytes.Buffer, name string, priceStr, numStr, totalNowStr string) {
	var weight int = 0
	var supWeight int = 0
	for _, r := range name {
		if unicode.Is(unicode.Scripts["Han"], r) {
			weight += weightHan
		} else {
			weight += weightOther
		}
	}
	fmt.Println("weight:", weight)

	if weight <= maxWeight {
		supWeight = maxWeight - weight
		if supWeight > 0 {
			for i := 0; i < supWeight; i++ {
				name = name + " "
			}
		}
		buf.WriteString(fmt.Sprintf("%s", name))
		buf.WriteString(fmt.Sprintf("%5s\t\t\t%2s\t%5s\n", priceStr, numStr, totalNowStr))

	} else {
		buf.WriteString(fmt.Sprintf("%s\n", name))
		buf.WriteString(fmt.Sprintf("%s%5s\t\t\t%2s\t%5s\n", createSpace(maxWeight), priceStr, numStr, totalNowStr))
	}
}

// 产生空格
func createSpace(len int) string {
	var space string
	for i := 0; i < maxWeight; i++ {
		space = space + " "
	}
	return space
}
