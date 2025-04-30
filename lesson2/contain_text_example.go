package lesson2

import (
	"log"
	"os"
	"strings"
)

func Contain(value string) string {
	// 6.2
	// use Map in context variable name
	textMap := make(map[string]string)
	filePath, _ := os.LookupEnv("INFO_TXT")

	// 6.2
	// old name: fileData
	// new name: fileBytes
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return ""
	}

	// 6.2
	// old name: stringData
	// new name: fileString
	fileString := string(fileBytes)

	// 6.2
	// old name: arrRows
	// new name: arrayOfString
	arrayOfString := strings.Split(fileString, "---")
	for _, arrayItem := range arrayOfString {
		// 6.2
		// old name: item
		// new name: arrayItem
		arrayItem = strings.TrimSpace(arrayItem)

		// 6.2
		// old name: arrItems
		// new name: splitArray
		splitArray := strings.Split(arrayItem, "=")
		textMap[splitArray[0]] = splitArray[1]
	}

	return textMap[value]
}
