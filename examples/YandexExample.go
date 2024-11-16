package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	yandex := api2captcha.Yandex{
		SiteKey: "Y5Lh0tiycconMJGsFd3EbbuNKSp1yaZESUOIHfeV",
		Url:     "https://rutube.ru",
	}

	req := yandex.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
