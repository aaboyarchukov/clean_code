# Переменные и их значения

Продолжаем впитывать полезные рекомендации по стилю кодирования.

1. Всегда явно объявляйте переменные

Явное объявление, это представление переменной вместе с типом, который данная переменная представляет.
Например:

```go
// bad example ❌
speedKmph := 0

// good example ✅
var speedKmph float32
```

2. Всегда инициализируйте переменные

Переменным необходимо при объявлении давать начальное значение, то есть инициализировать

3. Используйте переменные однократно

Пример:

```go
// bad example ❌
var speedKmph float32 = 0
var timeHours int32 = 0
var distanceKm float32 = 0
// ...
distanceKm = speedKmph * float32(timeHours)

// good example ✅
var speedKmph float32 = 0
var timeHours int32 = 0
// ...

var distanceKm float32 = speedKmph * float32(timeHours)
```

4. Изменять аргументы функций/методов - плохой стиль
5. Инициализируйте все атрибуты и поля класса в конструкторе
6. Завершение работы с переменными

После завершения работы с переменными необходимо присваивать им "недопустимые" значения.

7. Не оставляйте переменные не используемыми
8. Переменные и циклы

Переменные, которые будут использоваться в телах цикла, лучше инициализировать перед циклом.
Счетчики инициализируются в заголовке цикла.
Не забывайте обнулять аккумулятор перед его следующим использованием.
9. Проверяйте инварианты (истинные условия) в коде на недопустимые значения

Задания:

[6_example1.go](https://github.com/aaboyarchukov/clean_code/blob/master/lesson10/6_example1.go)

```go
token - var token *jwt.Token
// ясно выразил переменную

claims - var claims jwt.MapClaims
// ясно выразил переменную

diff - var diff float64
// ясно выразил переменную

object_id - var tenderId string
// ясно выразил переменную

equipment_type - var tenderEquipmentType string
// ясно выразил переменную

nds - var ndsStatus string
// ясно выразил переменную

strColumns - var columnsFromFile string
// ясно выразил переменную

resultColumns - var resultColumns data_analyze.ColumnsValues
// ясно выразил переменную

destination - var fileDestination string
// ясно выразил переменную

flagStr - var flagStr bool
// ясно выразил переменную

currentNode := l.head
deleted := false
for currentNode != nil {
	// ...
}
// здесь статус удаления отслеживаетс только один раз
// поэтому объявляем ее перед циклом

currentNode - var currentNode *Node[T]
// ясно выразил переменную

deleted - var deleted bool
// ясно выразил переменную

node - var addNode Node[T]
// ясно выразил переменную

asc - var asc bool
// ясно выразил переменную

desc - var desc bool
// ясно выразил переменную
```