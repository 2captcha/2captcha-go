package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	capy := api2captcha.Capy{
		SiteKey:   "PUZZLE_Abc1dEFghIJKLM2no34P56q7rStu8v",
		Url:       "https://www.mysite.com/captcha/",
		ApiServer: "https://jp.api.capy.me/",
	}

	req := capy.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
