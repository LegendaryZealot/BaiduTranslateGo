package main

import (
	"config"
	"encoding/json"
	"fmt"
	"math/rand"
	"md5"
	"os"
	"request"
)

type translateResultItem struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type translateResult struct {
	From        string                `json:"from"`
	To          string                `json:"to"`
	TransResult []translateResultItem `json:"trans_result"`
	Err         string                `json:"error_code"`
	ErrMsg      string                `json:"error_msg"`
}

var (
	url       = "http://api.fanyi.baidu.com/api/trans/vip/translate?q=%v&from=%v&to=%v&appid=%v&salt=%v&sign=%v"
	resultFmt = "%v->%v:\n%v:%v\n%v:%v"
)

var result translateResult

func main() {
	fmt.Println(fmt.Sprintf("version:%v", config.GetVersion()))
	var osArgsLen = len(os.Args)
	if 1 == osArgsLen {
		fmt.Println("[use]:cmd targetlanuage content")
		fmt.Println("[eg]:BaiduTranslateGo zh English")
		fmt.Println("[output]:\n\tversion:0.01\n\ten->zh:\n\tzh:英语\n\ten:English")
		fmt.Println("[language list]")
		fmt.Println("\ten:English")
		fmt.Println("\tzh:中文")
		fmt.Println("\tspa:Español")
		fmt.Println("\tjp:日本語")
		fmt.Println("\tkor:한국어")
		fmt.Println("\tth:ภาษาไทย")
		fmt.Println("\t...")
		return
	}
	from := "auto"
	to := os.Args[1]
	q := os.Args[2]
	salt := rand.Intn(10000)
	sign := md5.Encryption(fmt.Sprintf("%v%v%v%v", config.GetAppid(), q, salt, config.GetKey()))
	bytes, err := request.Get(fmt.Sprintf(url, q, from, to, config.GetAppid(), salt, sign))
	if nil != err {
		fmt.Println("err:%v", err)
	}
	err = json.Unmarshal(bytes, &result)
	if nil != err {
		fmt.Println("err:%v", err)
	}
	if 0 != len(result.Err) {
		fmt.Println(fmt.Sprintf("err:%v,%v", result.Err, result.ErrMsg))
		return
	}
	resultLen := len(result.TransResult)
	if resultLen < 1 {
		fmt.Println("result is empty")
	}
	for i := 0; i < resultLen; i++ {
		fmt.Println(fmt.Sprintf(resultFmt, result.From, result.To, result.To, result.TransResult[i].Dst, result.From, result.TransResult[i].Src))
	}
}
