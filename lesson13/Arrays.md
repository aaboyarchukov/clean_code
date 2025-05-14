# Массивы

Продолжаем впитывать полезные рекомендации

1. Всегда визуально убеждайтесь в том, что:
	1. Индексы не выходят за пределы массива
	2. Индексы не пересекаются (при вложенных циклах, можно перепутать `mass[i]` и `mass[j]`)
	3. При работе с многомерными массивами проверяйте, что индексы идут в правильном порядке без перескакиваний
2. Проверьте граничные точки при обработке массива (начало, конец и середина)
3. Рассмотрите возможность замены на другую структуру данных или используйте массивы избегая произвольного доступа к индексам

Задания:

[6_example1.go](https://github.com/aaboyarchukov/clean_code/blob/master/lesson13/6_example1.go)

```go
for _, equipment := range equipments {
        error_delete := DataBase.DB.Delete(&models.Equipment{}, &models.Equipment{
            EquipmentID: equipment,
        }).Error
        if error_delete != nil {
            return false, error_delete
        }
    }
// в данном случае массив используются для перебора элементов с последующим 
// удалением из БД без прямой индексации

// old:
// for i := 0; i < len(peopleAges); i++ {
//  sumPeopleAges += peopleAges[i]
// }

// new:
for _, age := range peopleAges {
	sumPeopleAges += age
}
// в данном случае заменил обход массива с помощью индексации
// на обход без индексации

// old:
// for ind := 0; ind < len(arrRows); ind++ {
//  item := strings.TrimSpace(arrRows[ind])
//  arrItems := strings.Split(item, "=")
//  textMap[arrItems[0]] = arrItems[1]
// }

// new:
for _, item := range arrRows {
	item = strings.TrimSpace(item)
	arrItems := strings.Split(item, "=")
	// new
	var itemsList list.List
	for _, item := range arrItems {
		itemsList.PushBack(item)
	}

	listFront, listBack := itemsList.Front().Value, itemsList.Back().Value
	var mapKey, mapValue string = "", ""

	if listValue, ok := listFront.(string); ok {
		mapKey = listValue
	}
	if listValue, ok := listBack.(string); ok {
		mapValue = listValue
	}
	
	textMap[mapKey] = mapValue
}

// в данном случае заменил обход массива с помощью индексации
// на обход без индексации, а также внутри цикла мы создали связный список
// для того, чтобы поместить в него два значения, на которые разделился массив
// тогда мы избавимся от надобности использовать индексы, но код получается перегруженным

// old:
// for ind := 0; ind < len(phones_array); ind++ {
//  if phones_array[ind] != "" {
//      phones_result = append(phones_result, phones_array[ind])
//  }
// }

// new:
for _, phone := range phones_array {
	if phone != "" {
		phones_result = append(phones_result, phone)
	}
}
// в данном случае заменил обход массива с помощью индексации
// на обход без индексации
```
