package generator

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func BannerList() ([]string, error) {
	files, err := os.ReadDir("banners")
	if err != nil {
		log.Printf("Failed to read directory: %v", err)
		return nil, err
	}

	var banners []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".txt" {
			bannerName := strings.TrimSuffix(file.Name(), ".txt")
			banners = append(banners, bannerName)
		}
	}

	return banners, nil
}
