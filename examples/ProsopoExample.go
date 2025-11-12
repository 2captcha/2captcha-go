package main

import (
	"fmt"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	prosopo := api2captcha.Prosopo{
		SiteKey: "5EZVvsHMrKCFKp5NYNoTyDjTjetoVo1Z4UNNbTwJf1GfN6Xm",
		Url:     "https://www.twickets.live/",
	}

	req := prosopo.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
