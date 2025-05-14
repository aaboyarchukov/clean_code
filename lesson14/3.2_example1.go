package lesson14

import (
	"log"
	"os"
	"strings"
)

// old:
// func Contain(key string) string {
// 	// формируем словарь, в котором будем хранить данные,
// 	// записанные в файл
// 	textMap := make(map[string]string)

// 	// получаем путь до файла с информацией
// 	info, _ := os.LookupEnv("INFO_TXT")

// 	// получаем данные из файла
// 	data, err := os.ReadFile(info)
// 	if err != nil {
// 		log.Println(err)
// 		return ""
// 	}

// 	// переводим данные из массива байт в строку
// 	str := string(data)

// 	// получаем строки по разделителю из файла
// 	data2 := strings.Split(str, "---")

// 	// проходимся по каждой строке файла
// 	for ind := 0; ind < len(data2); ind++ {
// 		// получаем строку без пробелов
// 		r := strings.TrimSpace(data2[ind])

// 		// получаем два значение из строки по разделителю
// 		items := strings.Split(r, "=")

// 		// первый элемент записываем как ключ, а второй как значение
// 		textMap[items[0]] = items[1]
// 	}

// 	// возвращаем значение
// 	return textMap[key]
// }

// поиск выражения по ключу в файле
func FindSentenceByKeyInFile(key string) string {
	textStorage := make(map[string]string)
	pathToFile, _ := os.LookupEnv("INFO_TXT")

	fileDataBytes, err := os.ReadFile(pathToFile)
	if err != nil {
		log.Println(err)
		return ""
	}

	fileDataString := string(fileDataBytes)
	fileRows := strings.Split(fileDataString, "---")
	for _, row := range fileRows {
		row = strings.TrimSpace(row)
		pairKeyAndValue := strings.Split(row, "=")
		var keyOfPair, valueOfPair string = pairKeyAndValue[0], pairKeyAndValue[1]
		textStorage[keyOfPair] = valueOfPair
	}

	return textStorage[key]
}
