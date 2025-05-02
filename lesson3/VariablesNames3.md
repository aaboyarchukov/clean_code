# Имена переменных 3

Продолжаем впитывать рекомендации по правильному и чистому стилю кодирования.

1. Использование булевых переменных

Название булевых переменных должно однозначно передавать либо false, либо true. В качестве имен можно использовать префикс `is` - (`isComplete, isError`) 
Еще одно важное замечание: 
	Необходимо использовать только утвердительные имена ‼️ (не надо писать `notFound, notExist, notComplete` - это сбивает с толку)

Также лучше использовать и внедрять в свой код стандартные имена булевых переменных, которые применяются в различных практиках:

- done -- признак (флажок) окончательного завершения цикла или другой операции. Присвойте ему false до выполнения действия, и установите его в true после его завершения;

- success или ok -- признак успешного завершения операции. Присвойте ему false, если операция завершилась неудачей, и true, если операция выполнена успешно. В некоторых случаях, когда логика кода довольно сложная, замените success на более определенное имя, ясно определяющее смысл «успеха» -- например, processingComplete (processing_complete). Если «успех» подразумевает обнаружение конкретного значения, можете использовать переменную с именем found;

- found -- определение того, обнаружено ли некоторое значение. Установите его в false, если значение не обнаружено, и в true, как только значение найдено. Используйте переменную found при поиске значения в массиве, идентификатора пользователя в файле, определенного объекта в списке объектов и т. д.

- error -- признак ошибки. Присвойте ей значение false, если всё в порядке, и true в противном случае.

Примечание:
- Лучше не использовать переменную `flag` , так как само название не несет никакой статус переменной

Пример использования:

```go
// bad example ❌
userFlag := FindUser(userID)

// good example ✅
found := FindUser(userID)	
```

2. Индексы циклов

Здесь все достаточно просто, при написании циклов, индексы принято обозначать одной буквой в зависимости от уровня вложенности (1 -> 2 -> ... -> n) -> (i, j, k...).

Также, если идет перечисление объектов массива с использованием специальных конструкций (range), тогда лучше использовать наглядное имя. А также в случаях, когда вам необходимо произвести сложные вычисления внутри цикла, тогда тоже лучше выбирать более наглядное имя.

Например:

```go
// bad example ❌
for indexOfArray := 0; indexOfArray < arraySize; indexOfArray++ {
	// ...
} 

// good example ✅
for i := 0; i < arraySize; i++ {
	// ...
} 

// or

for indx := 0; indx < arraySize; indx++ {
	// ...
} 
```

Перечисления:

```go
// bad example ❌
for _, i := range items {
	// ...
}

// good example ✅
for _, item := range items {
	// ...
}
```

Но если при перечислении используете индекс, тогда его тоже лучше именовать одной буквой:

```go
// bad example ❌
for indexOfItems, itemOfItems := range items {
	// ...
}

// good example ✅
for i, item := range items {
	// ...
}

// or

for indx, item := range items {
	// ...
}
```

3. Антонимы

В определенных случаях лучше использовать антонимы для наглядности и ясности кода.

Например:

```go
begin -> end
first -> last
```

4. Временные переменные

Забудьте про именование `temp` или `x` - данные имена ничего не говорят о внутренней составляющей данных, которые вычисляются. Если локальные/временные переменные все-таки есть в каких-то частях программы, тогда их надо называть также осмыслено, как и остальные переменные.
Также, необходимо **избавляться** от временных переменных в вашем коде.

```go
// bad example ❌
for _, order := range orders {
	temp := order.OrderID
	// ...
}

// good example ✅
for _, order := range orders {
	orederID := order.OrderID
	// ...
}
```

5. Длина имени в зависимости от области видимости

Здесь все просто - чем шире область видимости, чем длиннее и информативнее должно быть имя. Об этом упоминалось на прошлых уроках [[Имена переменных 2#^f52e9f|Имена переменных 2]].

Задания:

7.1. Приведите пять примеров правильного именования булевых переменных в вашем коде в формате "было - стало".

[**7.1_example1.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson3/7.1_example1.go)

```go
is_repeat_object_desc - repeat
// повторяется ли такое же описание объекта
// в данном случае использую repeat, так как 
// мы вкладываем контекст названия функции Is_repeat_object_desc
// данное имя более точно отображает статус, чемм found

is_repeat_company - repeat
// повторяется ли такая компания в базе данных
// в данном случае использую repeat, так как 
// мы вкладываем контекст названия функции Is_repeat_company
// данное имя более точно отображает статус, чем found

is_initiator_repeat - repeat
// в данном случае использую repeat, так как 
// мы вкладываем контекст названия функции Is_repeat_company
// данное имя более точно отображает статус, чем found

is_repeat_company - repeat
// повторяется ли такая компания в базе данных
// в данном случае использую repeat, так как 
// мы вкладываем контекст названия функции Is_repeat_company
// данное имя более точно отображает статус, чем found

is_vendor_exist - exist
// существует ли такой поставщик в базе данных
// в данном случае использую exist, так как 
// мы вкладываем контекст названия функции Is_vendor_exist
// данное имя более точно отображает статус, чем found

existLogin - exist
// повторяется ли такая компания в базе данных
// в данном случае использую exist, так как 
// мы вкладываем контекст названия функции IsLoginExist
// данное имя более точно отображает статус, чем found

// по хорошему на этапе проектирования данных функций 
// необходимо было по-другому назвать их
// например: FoundObjectDesc или FoundCompany

deleted - deleted
// статус узла (удален или нет)
// в данном случае использую deleted,
// а не done, так как данное имя больше передает контекста

flagStr - isStr
// статус значения - является ли оно строкой
```


7.2. Найдите несколько подходящих случаев, когда в вашем коде можно использовать типичные имена булевых переменных.

[**7.2_example1.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson3/7.2_example1.go)

```go
// 7.2
// name: ok

// ...
path, ok := os.LookupEnv("PATH_KP_FILES")
if !ok {
	return Answer_from_kp{}, fmt.Errorf("error with reading .env virables")
}
// ...

// 7.2

// name: success
var result T
success := nd.IsKey(key)
if !success {
	return result, fmt.Errorf("key is not in array")
}

// 7.2
// name: ok

if value, ok := any(v1).(string); ok {
	valueStr1 = strings.Trim(value, " ")
	isStr = true
}

// 7.2
// name: ok

if value, ok := any(v2).(string); ok {
	valueStr2 = strings.Trim(value, " ")
	isStr = true
}

```

7.3. Проверьте, правильно ли вы даёте имена индексам циклов. Попробуйте найти случай, когда вместо i j k нагляднее использовать более выразительное имя.

[**7.3_example1.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson3/7.3_example1.go)

```go
// правильность индексов в циклах

// 7.3
for i := 0; i < restaurant.waiters; i++ {
	restaurant.wg.Add(1)
	go restaurant.waiter(i + 1)
}

// 7.3
for i := 0; i < restaurant.chefs; i++ {
	restaurant.wg.Add(1)
	go restaurant.chef(i + 1)
}

// пример более выразительных индексов

// 7.3
// в данном случае идет подсчет статистики по каждому официанту,
// у каждого работника есть свой id,
// тогда более наглядно назвать индекс - waiterID,
// так как мы их используем как идентифицируещие номера официантов
for waiterID, stats := range restaurant.waiterStats {
	stats.mu.Lock()
	tablesCount := len(stats.tablesServed)
	ordersCount := stats.ordersCount
	stats.mu.Unlock()
	fmt.Printf("Официант #%d: обслужил %d столов, принял %d заказов\n",
		waiterID, tablesCount, ordersCount)
}
```

7.4. Попробуйте найти в своих решениях два-три случая, когда можно использовать пары имён - антонимы.

[**7.4_example1.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson3/7.4_example1.go)

```go
// 7.4
start, err_start := time.Parse("2006-01-02 15:04:05-07", req_result.Start)
if err_start != nil {
	return ctx.JSON(err_start)
}

// 7.4
end, err_end := time.Parse("2006-01-02 15:04:05-07", req_result.End)
if err_end != nil {
	return ctx.JSON(err_end)
}

// 7.4

var beginRows int
for _, ind := range indexColumns {
	if ind.Row != -1 {
		begin_rows = ind.Row + 1
		break
	}
}

// 7.4
var endRows int = len(rows)
equipments := []models.Equipment{}
for ind := beginRows; ind < endRows; ind++ {
	// ...
}

```

7.5. Всем ли временным переменным в вашем коде присвоены выразительные имена? Найдите несколько случаев, когда временные переменные надо переименовать, и поищите, возможно, от некоторых временных переменных вам получится вообще полностью избавиться.

[**7.5_example1.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson3/7.5_example1.go)

```go
// rename

// 7.5
// old name: tempNode
// new name: currentNode

currentNode := l.head
for currentNode != nil {
	count++
	currentNode = currentNode.next
}

// 7.5
// old name: item
// new name: objectCompanies

var objectCompanies object_companies
objectCompanies.Object_id = work_part.ObjectRefer

// 7.5
// old name: temp_item
// new name: sellerAndVendorEquipments

sellerAndVendorEquipments := equipmentSellerAndVendor{
	SellerEquipment:  seller_equipment,
	VendorEquipments: result_vendors_equipments,

}

// 7.5
// old name: temp_vendor
// new name: vendorCompany

vendorCompany := companyVendor{
	CompanyInfo: vendor_info,
	CompanyID:   vendor,
	Equipments:  vendor_equipments,
}

// 7.5
// old name: temp_data_item
// new name: equipmentTypesData

equipmentTypesData := equipment_types{
	EquipmentType:   equipment_type.Name,
	EquipmentTypeID: equipment_type.EquipmentTypeID,
	Seller:          sellers,
	Vendors:         result_vendors,

}

// remove

// 7.5
// we could remove all locale variable

// before

tablesCount := len(stats.tablesServed)
ordersCount := stats.ordersCount
stats.mu.Unlock()
fmt.Printf("Официант #%d: обслужил %d столов, принял %d заказов\n",
	waiterID, tablesCount, ordersCount)

// after
stats.mu.Unlock()
fmt.Printf("Официант #%d: обслужил %d столов, принял %d заказов\n",
	waiterID, len(stats.tablesServed), stats.ordersCount)

// 7.5
// we could remove variable destinaiton

// before
destination := fmt.Sprintf("%s%s", path, file.Filename)
if save_file_err := ctx.SaveFile(file, destination); save_file_err != nil {
	log.Println("ошибка сохранения файла")
	log.Println(save_file_err)
	return save_file_err
}

// after
if save_file_err := ctx.SaveFile(file,
	fmt.Sprintf("%s%s", path, file.Filename),
); save_file_err != nil {
	log.Println("ошибка сохранения файла")
	log.Println(save_file_err)
	return save_file_err
}

// 7.5
// we could remove temp_equipment_id

// before

var equipment_id uint
temp_equipment_id, err_get_equipment_id := Add_equipment(equipment)
if err_get_equipment_id != nil {
	return err_get_equipment_id
}
equipment_id = temp_equipment_id

// after
equipment_id, err_get_equipment_id := Add_equipment(equipment)
if err_get_equipment_id != nil {
	return err_get_equipment_id
}

```