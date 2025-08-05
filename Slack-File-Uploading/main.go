package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	botToken := os.Getenv("BOT_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")

	api := slack.New(botToken)

	fileArr := []string{""}

	for _, filePath := range fileArr {
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Failed to read file %s: %v", filePath, err)
			continue
		}

		if len(fileData) == 0 {
			log.Printf("Skipping empty file %s", filePath)
			continue
		}

		filename := filepath.Base(filePath)

		params := slack.UploadFileV2Parameters{
			Channel:  channelID,
			Filename: filename,
			Reader:   bytes.NewReader(fileData), // use buffer here!
		}

		uploadedFile, err := api.UploadFileV2(params)
		if err != nil {
			log.Printf("Failed to upload file %s: %v", filename, err)
			continue
		}

		fileDetail, _, _, err := api.GetFileInfo(uploadedFile.ID, 0, 0)
		if err != nil {
			log.Printf("Failed to get uploaded file info: %v", err)
			continue
		}

		fmt.Printf("Uploaded: Name: %s, URL: %s\n", fileDetail.Name, fileDetail.URL)
	}
}
