package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	captcha := api2captcha.MTCaptcha{
		SiteKey: "MTPublic-KzqLY1cKH",
		Url:     "https://2captcha.com/demo/mtcaptcha",
	}
	req := captcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println(token)
	fmt.Println(captchaId)
	fmt.Println(err)
}
