package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	text := api2captcha.Text{
		Text: "If tomorrow is Saturday, what day is today?",
	}

	req := text.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
