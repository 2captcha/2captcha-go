package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	atbCaptcha := api2captcha.AtbCAPTCHA{
		AppId:     "af23e041b22d000a11e22a230fa8991c",
		Url:       "https://www.playzone.vip/",
		ApiServer: "https://cap.aisecurius.com",
	}

	req := atbCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
