# Имена переменных 4

Продолжаем дальше изучать рекомендации!

1. Избегайте похожих имен по написанию, но разных по значению

При написании кода необходимо избегать ситуаций, когда вы называете переменные практически одинаково, но смысл в них вы закладываете разный, например:

```go
// bad example ❌
value
insertValue
insertStringValue
insertIntValue
fileRow
fileString
indexOfArray
indeOfList
// ... etc
```

2. Избегайте указания цифр в имени

Также не нужно указывать цифры в названии переменных, например:

```go
// bad example ❌
index10
value0
user00
```

3. Избегайте указания ключевых слов и типов, которые определены в языке, как имя переменной

В данном случае также не нужно называть переменные, как ключевые слова в вашем языке, которые принадлежат типам, функциям, методам и т.д.

```go
// bad example ❌
interface, error, int, var
```

4. Избегайте не информативных имен

Также при выборе имени для переменной не стоит давать полностью неинформативное имя:

```go
// bad example ❌
table, value, nameString, numberInt, array, object, result 
```

5. Не используйте имена со скрытом смыслом, который понимаете только вы

Старайтесь внедрять имена, которые будут однозначно понятны всем, а не какие-нибудь специфические, которые понятны только вам (например транскрипции русских слов или ваши сокращения, аббревиатуры):

```go
// bad example ❌
andrey, snils, polis, hp, pc, kp
```

6. Используйте спецификаторы в конце имени и не используйте в имени информацию о типе и области видимости

В переменных, которые отражают вычисления и различные операции, можно указывать спецификаторы типа `Total, Sum, Average, Max, Min, Record, String или Pointer` в конце имени, так оно будет более информативно и понятно, за исключением `Num, number` 

```go
// good example ✅
ageMin, ageMax, scoreTotal, heightAverage
```

А также лучше не указывать в именах информации о типе (если они точно им не являются) и области видимости:

```go
// bad example ❌
phoneNumbersList, phoneNumbersArray

// good example ✅
phoneNumbers
```

7. Уточняйте в именах единицы измерения, если они подразумеваются 

Также всегда важно указывать в коде конкретные единицы измерения, ели этого требует программа и ее контекст, а также прописывать в документации

```go
// bad example ❌
speed

// good example ✅
speedKmps, speedMps, speedMph
```

Задание:

Найдите 12 примеров имён в вашем коде, которые следует избегать, исправьте, и выложите на гитхаб в формате "было - стало" (с учётом контекста).

[8_example1.go](https://github.com/aaboyarchukov/clean_code/blob/master/lesson4/8_example1.go)

```go

resultLL - linkedList
// итоговый связный список, который мы получаем
context:
	var linkedList LinkedList // resulting linked list

l1 - firstList
// первый связный список, котоырый используется в функции
context:
	func EqualLists(firstList *LinkedList, secondList *LinkedList) bool

l2 - secondList
// второй связный список, котоырый используется в функции
context:
	func EqualLists(firstList *LinkedList, secondList *LinkedList) bool

countLL1 - firstListSize
// размер первого списка
context:
	firstListSize, secondListSize := firstList.Count(), secondList.Count()

counLL2 - secondListSize
// размер второго списка
context:
	firstListSize, secondListSize := firstList.Count(), secondList.Count()

tempLL1 - currNodeFirstList
// текущий узел первого списка
context:
	currNodeFirstList, currNodeSecondList := firstList.head, secondList.head

tempLL2 - currNodeSecondList
// текущий узел второго списка
context:
	currNodeFirstList, currNodeSecondList := firstList.head, secondList.head

count - nodesCount
// количество узлов в списке
context:
	func (l *OrderedList[T]) Count() int {
		nodesCount := 0
		// ...
	}

tempNode - currentNode
// текущий узел списка
context:
	currentNode := l.head

result - arrayFromList
// итоговый массив из списка
context:
	func (l *OrderedList[T]) ToArray() []T {
		arrayFromList := make([]T, 0, l.Count())
		// ...
	}

tempNode - currentNode
// текущий узел списка
context:
	currentNode := l.head

object_id_struct - requestStruct
// переменная структурного типа для запроса
context:
	type request struct {
        ID string `json:"id"`
    }
    requestStruct := request{}


id - requestBody
// тело запроса
context:
	requestBody := ctx.Request().Body()

object_id - tenderID
// идентификационный номер торга с торговой площадки
context:
	tenderID := requestStruct.ID

types_equipment - tenderEquipmentTypes
// все типы для всех оборудований по данному торгу
context:
	tenderEquipmentTypes, errGetTypes := 
		database.Get_types_of_equipments_of_object(tenderID)

data_array - equipmentTypesGroup
// группа типов для оборудований из торга
context:
	equipmentTypesGroup := make([]equipment_types, 0, len(tenderEquipmentTypes))

vendors_array - tenderVendors
// все поставщики для тендера
context:
	tenderVendors, errGetVendors := database.Get_vendors(work_part.Work_PathID)

result_vendors - vendorsGroup
// группа поставщиков для торга
context:
	vendorsGroup := make([]companyVendor, 0, len(tenderVendors))
	
```
