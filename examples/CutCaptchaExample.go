package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	cutCaptcha := api2captcha.CutCaptcha{
		MiseryKey:  "a1488b66da00bf332a1488993a5443c79047e752",
		Url:        "https://filecrypt.co/Container/237D4D0995.html",
		DataApiKey: "SAb83IIB",
	}

	req := cutCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
