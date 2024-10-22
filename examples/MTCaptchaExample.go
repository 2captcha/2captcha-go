package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	mtCaptcha := api2captcha.MTCaptcha{
		SiteKey: "MTPublic-KzqLY1cKH",
		Url:     "https://2captcha.com/demo/mtcaptcha",
	}

	req := mtCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
