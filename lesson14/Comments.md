# Комментарии

Продолжаем впитывать полезные рекомендации.

Основная мысль комментариев в том, чтобы описывать, что делает ваш код, если по вашему коду непонятно, что он делает, тогда это указывает на то, что вам надо **ПЕРЕРАБОТАТЬ ВАШ КОД**, чтобы он был яснее.

В большинстве случаев комментарии засоряют ваш код и являются плохим стилем, лучше грамотно работайте с кодом, чем пишите комментарии.

Также стоит отметить то, что комментарии можно оставлять для уточнения, если имена ваших переменных/функций, ясны, но сложно читаемы на английском, в остальных случаях - переписывайте и улучшайте код, а не строчите комментарии.

Например:

```go
// bad example ❌ 
if ((employee.flags & HOURLY_FLAG) && (employee.age > 65)) ... 

// good example ✅ 
if (employee.isEligibleForFullBenefits())

// also good example ✅
// Проверить, положена ли работнику полная премия 
if (employee.isEligibleForFullBenefits())
```

Задания:

3.1. Прокомментируйте 7 мест в своём коде там, где это явно уместно

[3.1_example1.go](https://github.com/aaboyarchukov/clean_code/blob/master/lesson14/3.1_example1.go)

```go
// формирование отчета о зарплате за определенный период
func (employee *Employee) PrepareSalaryReportByPeriod(period time.Time) Report {}

// отношение между стадиями обработки и типами оборудований по торгам
type ProccesingStagesAndEquipmentTypes struct {
	// fields...
}

// отношение между оборудованием покупателей и поставщиков в тендере
type SellerAndVendorsEquipmentOfTender struct {
	// fields...
}

// история взаимодействия с компанией по тендеру
type HistoryInteractionWithCompanyOfTender struct {
	// fields...
}

// отношение описательной части рабочей части
// (процесс, когда тендер находится в обработке) тендера
type WorkPartDescriptionOfTender struct {
	// fields...
}

// отношение рабочей части
// (процесс, когда тендер находится в обработке) тендера
type WorkPartOfTender struct {
	// fields...
}

// формирование коммерческого предложения по торгу
func GetCommericalProposal(object_id string, columns data_analyze.ColumnsValues, ndsStatus string) (Answer_from_kp, error) {
 // logic...
}
```


3.2. Если вы раньше делали комментарии к коду, найдите 5 мест, где эти комментарии были излишни, удалите их и сделайте сам код более наглядным.

[3.2_example1.go](https://github.com/aaboyarchukov/clean_code/blob/master/lesson14/3.2_example1.go)

```go
// old:

func Contain(key string) string {
    // формируем словарь, в котором будем хранить данные,
    // записанные в файл
    textMap := make(map[string]string)

    // получаем путь до файла с информацией
    info, _ := os.LookupEnv("INFO_TXT")

    // получаем данные из файла
    data, err := os.ReadFile(info)

    if err != nil {
        log.Println(err)
        return ""
    }

  

    // переводим данные из массива байт в строку
    str := string(data)
    // получаем строки по разделителю из файла
    data2 := strings.Split(str, "---")
    // проходимся по каждой строке файла
    for ind := 0; ind < len(data2); ind++ {
        // получаем строку без пробелов
        r := strings.TrimSpace(data2[ind])
  
        // получаем два значение из строки по разделителю
        items := strings.Split(r, "=")

        // первый элемент записываем как ключ, а второй как значение
        textMap[items[0]] = items[1]
    }
    // возвращаем значение
    return textMap[key]
}

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
```

