package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	keyCaptcha := api2captcha.KeyCaptcha{
		UserId:         10,
		SessionId:      "493e52c37c10c2bcdf4a00cbc9ccd1e8",
		WebServerSign:  "9006dc725760858e4c0715b835472f22",
		WebServerSign2: "2ca3abe86d90c6142d5571db98af6714",
		Url:            "https://www.keycaptcha.ru/demo-magnetic/",
	}

	req := keyCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
