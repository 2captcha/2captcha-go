package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	hCaptcha := api2captcha.HCaptcha{
		SiteKey: "c0421d06-b92e-47fc-ab9a-5caa43c04538",
		Url:     "https://2captcha.com/demo/hcaptcha",
	}

	req := hCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
