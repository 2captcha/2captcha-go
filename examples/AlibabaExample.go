package main

import (
	"fmt"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	alibaba := api2captcha.Alibaba{
		SceneId: "bxs_login",
		Prefix:  "a",
		Url:     "https://example.com/page-with-alibaba",
	}

	req := alibaba.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
