package internal

import (
	"log"
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
