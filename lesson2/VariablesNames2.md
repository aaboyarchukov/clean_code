# Имена переменных 2
Далее изучим еще одни рекомендации:

1. Имена, объявляемых переменных, должны быть на подходящем уровне абстракции

То есть, в названиях должен быть отражен этот уровень - имя должно быть конкретизировано.

Например:
В задании нам необходимо посчитать количество синих машин, проезжающих через определенный перекресток. В плохом примере имя подобрано недостаточно хорошо, поскольку оно отражает уровень абстракции выше, чем необходимо в задаче:

```go
// bad example ❌
var countCars int

// good example ✅
var countBlueCars int	
```


2. Использовать имена из пространства решения и также применять технические термины

При объявлении названий переменных, функций и методов лучше будет вложить в имя технический термин, для ясности кода для программистов, поскольку лишь по названию функции разработчик сразу получит более широкий контекст и более ясное понимание об реализации данной функции.

Например:
В задании нам необходимо найти человека определенного роста, который стоит в упорядоченной последовательности из человек. Из задачи понятно, что надо использовать бинарный поиск для более высокой эффективности, поскольку обычный поиск занял бы `O(n)`, а бинарный - `O(log(n))`:

```go
// bad example ❌
func FindPerson(height int, peopleHeights []int) (int, error) {...}

// good example ✅
func PersonBinarySearch(height int, peopleHeights []int) (int, error) {...}
```

3. Вкладывайте в имена контекст.

Для большей ясности кода и более высокой скорости его понимания, необходимо в имена вкладывать контекст, если того требует ситуация.

Например:
В задании вам необходимо определить структуру для инициализации хранилища в приложении (БД). Для этого вы выбрали подходящую систему и реализовывайте именно для нее определенную структуру и интерфейс. Лучше будет вложить в контекст названия выбранную вами СУБД (в данном случае).

```go
// bad example ❌
type Storage struct {
	// fields...
}

// good example ✅
type PostgresStorage struct {
	// fields...
}
```


4. Делайте имена более выразительными и читабельными.

Имена ваших переменных должны быть ясными и понятными для читателя. Если имя переменной трудно произносимо, тогда что-то не так

Например:
Для вашего приложения необходимо разработать конфиг, с помощью которого вы будете работать с переменными среды. В данном случае будет лишним сокращение количества букв в слове, ведь и без этого полное слово читается хорошо.

```go
// bad example ❌
var appCnfg *Config

// good example ✅
var appConfig *Config
```

5. Длина имени должна быть в диапазоне от 8 до 20 символов.

Кроме длины важно учесть область, в которой переменная живет, чем она меньше, тем короче название.

Например:

```go
// bad example ❌
for indexOfList := 0; indexOfList < lenOfList; indexOfList++ {
	// some code...
}

// good example ✅
for indx := 0; indx < lenOfList; indx++ {
	// some code...
}

// or

// good example ✅
for i := 0; i < lenOfList; i++ {
	// some code...
}
```

Также еще один пример:

```go
// bad example ❌
var nameOfUser string

// good example ✅
var userName string
```

Еще один пример:

```go
// bad example ❌
var hp []byte

// good example ✅
var hashingPassword []byte
```


Пример ответа:

```
oldName - newName
// info about variable
```

Задания:

6.1 Разберите свой код, и сделайте пять примеров, где можно более наглядно учесть в именах переменных уровни абстракции.

[**relations_example.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson2/relations_example.go)

```go
object - updateCompanyObject
// обновление объекта конкретной компании в базе данных

equipments - getEquipments
// получение всех оборудований

result - getEquipmentsIDs
// идентифицируещие номера оборудований

object - getObjectEquipments
// поиск всех оборудований для определенного объекта

object - updateWinnerCompany
// обновление компании победителя

result - getActions
// получение всех действий

object - getActionsObject
// объект, в котором хранятся все действия
```


6.2. Приведите четыре примера, где вы в качестве имён переменных использовали или могли бы использовать технические термины из информатики.

[**contain_example.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson2/contain_example.go)

```go
textMap - textMap // (use Map in context variable name)
// результирующий словарь из исходного текста

fileData - fileBytes // (string)
// считывание файла и получение массива байт данных

stringData - fileString // (string)
// преобразование полученных байт из файла в строку

arrRows - arrayOfString // (string, array)
// разделение строки на массив из строк

item - arrayItem // (array)
// получение одного экземпляра из массива

arrItems - splitArray // (array, split)
// разделение экземпляра из массива на массив из строк
```

[**linked_list_example.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson2/linked_list_example.go)

```go
resultLL - linkedList // (linked list)
// формирование связного списка из значений 
```

6.3. Придумайте или найдите в своём коде три примера, когда имена переменных даны с учётом контекста (функции, метода, класса).

[**server_example.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson2/server_example.go)

```go
// перечислим примеры имен с контекстом
errValidate, errLogin, errRegister, userID
```

[**jwt_example.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson2/jwt_example.go)

```go
// перечислим примеры имен с контекстом
signedToken, errSignedToken, secretKey
```

6.4. Найдите пять имён переменных в своём коде, длины которых не укладываются в 8-20 символов, и исправьте, чтобы они укладывались в данный диапазон.

[**login_user_example.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson2/login_user_example.go)

```go
err - errJSONUnmarshal
```

[**get_object_date_example.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson2/get_object_date_example.go)

```go
result - requestBody
// тело запроса

err - errJSONUnmarshal
// ошибка при парсинге объекта

date - objectDate
// дата принадлежащая объекту
```

[**get_tables_columns_example.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson2/get_tables_columns_example.go)

```go
path - filePath
// путь к файлу

ok - notEmptyPathENV
// статус, сообщающий о пустоте переменной в env

files - filesFromDir
// файлы из директории

file - firstFile
// первый файл из директории

name - nameOfFile
// имя файла

rows - rowsFromFile
// строки, считанные из файла

columns - columnsForAnswer
// колонки для ответа на запрос
```