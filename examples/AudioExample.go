package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	assetsDir := currentDir + "/assets"

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(assetsDir)

	//var buf bytes.Buffer

	fileName := assetsDir + "/" + "audio-en.mp3"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}

	//data := buf.Bytes()
	fileBase64Str := base64.StdEncoding.EncodeToString(bs)

	//fmt.Println(fileBase64Str)
	audio := api2captcha.Audio{
		Base64: fileBase64Str,
		Lang:   "en",
	}

	req := audio.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
