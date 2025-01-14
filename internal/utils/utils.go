package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"path/filepath"
	"github.com/google/uuid"
)

func GetTimeStampString() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func GetAvaFolder() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	ava_folder := filepath.Join(homeDir, ".ava")

	// create the folder if it doesn't exist
	os.MkdirAll(ava_folder, 0755)

	return ava_folder
}

func GetNewUuid() string {
	return uuid.New().String()
}