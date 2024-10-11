package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func GetAssetsDir(currentDir string) string {
	currentDir, err := filepath.Abs(filepath.Dir(currentDir))
	assetsDir := currentDir + "/assets"

	if err != nil {
		log.Fatal(err)
	}
	return assetsDir
}

func ReadFile2BiteSlice(fileName string) []byte {

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
