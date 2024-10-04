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

	assetsDir := getAssetsDir(os.Args[0])

	fileName := assetsDir + "/" + "audio-en.mp3"

	bs := readFile2BiteSlice(fileName)

	if bs == nil {
		return
	}

	fileBase64Str := base64.StdEncoding.EncodeToString(bs)

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

func readFile2BiteSlice(fileName string) []byte {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return nil
	}
	return bs
}

func getAssetsDir(currentDir string) string {
	currentDir, err := filepath.Abs(filepath.Dir(currentDir))
	assetsDir := currentDir + "/assets"

	if err != nil {
		log.Fatal(err)
	}
	return assetsDir
}
