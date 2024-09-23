package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	tencentCaptcha := api2captcha.Tencent{
		AppId:   "2092215077",
		PageUrl: "http://lcec.lclog.cn/cargo/NewCargotracking?blno=BANR01XMHB0004&selectstate=BLNO",
	}

	req := tencentCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Println(err)
}
