package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	lemin := api2captcha.Lemin{
		CaptchaId: "CROPPED_d3d4d56_73ca4008925b4f83a8bed59c2dd0df6d",
		Url:       "http://sat2.aksigorta.com.tr",
		ApiServer: "api.leminnow.com",
	}

	req := lemin.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
