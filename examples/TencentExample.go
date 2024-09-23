package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	//"06c869c488704f62826181f2562ac999"
	//argsWithProg := os.Args
	client := api2captcha.NewClient(os.Args[1])

	captcha := api2captcha.MTCaptcha{
		SiteKey: "MTPublic-KzqLY1cKH",
		Url:     "https://2captcha.com/demo/mtcaptcha",
	}
	req := captcha.ToRequest()
	code, code1, err := client.Solve(req)

	fmt.Println(code)
	fmt.Println(code1)
	fmt.Println(err)
}
