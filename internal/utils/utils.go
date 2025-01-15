package utils

import (
	"fmt"
	"log"
	"os"
	"time"
	"regexp"
	"strings"

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

type Tag struct {
	Name    string
	Content string
}

func ExtractTags(input string) ([]Tag, error) {
	// Define a regex to match tagged sections
	re := regexp.MustCompile(`(?s)\{\{(\w+)\}\}(.*?)\{\{/\w+\}\}`)

	// Find all matches
	matches := re.FindAllStringSubmatch(input, -1)

	if matches == nil {
		return nil, fmt.Errorf("no tagged sections found")
	}

	var tags []Tag

	// Iterate over matches
	for _, match := range matches {
		if len(match) >= 3 {
			openTag := match[1]
			closeTagPattern := fmt.Sprintf(`\{\{/%s\}\}`, openTag)
			closeTagRe := regexp.MustCompile(closeTagPattern)

			// Ensure the closing tag matches the opening tag
			if closeTagRe.MatchString(input) {
				tags = append(tags, Tag{
					Name:    openTag,
					Content: match[2],
				})
			}
		}
	}

	for i := range tags {
		tags[i].Content = strings.Trim(tags[i].Content, "\n")
	}

	return tags, nil
}