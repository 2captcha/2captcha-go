package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	geeTestV4 := api2captcha.GeeTestV4{
		CaptchaId: "72bf15796d0b69c43867452fea615052",
		Url:       "https://mysite.com/captcha.html",
		Challenge: "12345678abc90123d45678ef90123a456b",
		ApiServer: "api-na.geetest.com",
	}

	req := geeTestV4.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
