package main

import (
	"encoding/base64"
	"fmt"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
	helper "github.com/2captcha/2captcha-go/examples/internal"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	assetsDir := helper.GetExamplesAssetsDir(os.Args[0])

	bodyFile := assetsDir + "/vk.jpg"

	bodyBs := helper.ReadFile2BiteSlice(bodyFile)

	body := base64.StdEncoding.EncodeToString(bodyBs)

	vk := api2captcha.Vk{
		Body:  body,
		Steps: "[5,12,22,24,21,23,10,7,2,8,19,18,8,24,21,22,11,14,16,5,18,20,4,21,12,6,0,0,11,12,8,20,19,3,14,8,9,13,16,24,18,3,2,23,8,12,6,1,11,0,20,15,19,22,17,24,8,0,12,5,19,14,11,6,7,14,23,24,23,20,4,20,6,12,4,17,4,18,6,20,17,5,23,7,10,2,8,9,5,4,17,24,11,14,4,10,12,22,21,2]",
	}

	req := vk.ToRequest("vkimage")

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
