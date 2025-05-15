# Плохие комментарии

Продолжаем разбирать рекомендации.

Важное правило: отсутствие комментариев лучше, чем иметь плохие комментарии!!!

1. Связь между кодом и комментарием должна быть ясной
2. Не пишите комментарии на скорую руку - если нужны комментарии, то уделите им достаточно времени
3. Пишите достоверные комментарии - не сочиняйте
4. Не пишите в комментариях "шум" (очевидные вещи)
5. Позиционные маркеры в виде /////////////// и других, не выделяйте так комментарии, редкость комментариев в ходе - лучшее их выделение
6. Если возникает потребность прокомментировать закрывающую скобку - перефакторите функцию и уменьшите ее
7. Не надо писать мемуары в комментариях, комментарии не должны быть больше функций
8. Не оставляйте в комментариях исторические факты и подобное
9. Не пишите комментарии, в которых описана не локальная информация
10. Удалите закомментированный код
11. Обязательные комментарии - плохой стиль
12. Хорошее название функции или переменной - лучше любых комментариев

Задания:

[13_example1.go](https://github.com/aaboyarchukov/clean_code/blob/master/lesson16/13_example1.go)

```go
// ПУНКТ 1

// old: стадии оборудования
// отношение между стадиями обработки и типами оборудований по торгам
type ProccesingStagesAndEquipmentTypes struct {
    // fields...
}
// комментарий неявный, сделал его явным

// old: оборудования покупателей и поставщиков
// отношение между оборудованием покупателей и поставщиков в тендере
type SellerAndVendorsEquipmentOfTender struct {
    // fields...
}
// комментарий неявный, сделал его явным

// old: история компании
// история взаимодействия с компанией по тендеру
type HistoryInteractionWithCompanyOfTender struct {
    // fields...
}
// комментарий неявный, сделал его явным

// old: описание рабочей части
// отношение описательной части рабочей части
// (процесс, когда тендер находится в обработке) тендера
type WorkPartDescriptionOfTender struct {
    // fields...
}
// комментарий неявный, сделал его явным

// ПУНКТ 2

func (bf *BloomFilter) Hash2(s string) int {
    sum := 0
    for _, char := range s {
        code := int(char)
        // old:
        // сумма битов строки, умноженная на определенное число
        sum += code * HASH_2_KOEFF
    }
    sum %= bf.filter_len
    return sum
}
// данный комментарий здесь ни к чему, его надо удалить

// ПУНКТ 3

// old:
// реализация фильтра Блюма, который гаранитрует со 100% вероятностью,
// что элемент есть в множестве
type BloomFilter struct {
    filter_len int
    filter     int64
}
// комментарий не достоверный, либо надо описать правильно, либо удалить
// я удалю, так как название структуры говорит за себя, что это такое,
// а подробнее можно почитать устройство в интернете

// ПУНКТ 4

// old:
// присваивание константе значение
const HASH_1_KOEFF int = 17
// комментарий содержит шум - надо удалять комментарий

// ПУНКТ 5

// old:
//////////////////////COMMENTS/////////////////////////
// настройка ручек сервера
func Setup_Routes(app *fiber.App) {
	// logic
}
// присутствует ненужный маркер, а также комментарий - их надо удалить

// ПУНКТ 6

func GetDataForTables() {

    // logic...
    for _, equipment_type := range types_equipment {

        for _, seller_equipment := range seller_equipments {

            for _, vendorEquipmentId := range vendor_equipments_ids {

            } // закрытие первого внутреннего цикла в первом внутреннем цикле


        } // закрытие первого внутреннего цикла


        for _, vendor := range vendors_array {

        } // закрытие второго внутреннего цикла

    }

    // logic...

}
// здесь присутвуют комментарии при закрытых скобках - их надо удалить

// ПУНКТ 7 и 8

// данная функция высчитывает хэш для значения
// которое будет помещено в фильтр Блюма
// который является способом нахождения значения
// в множестве за счет следующего принципа...
func (bf *BloomFilter) Hash2(s string) int {

}
// комментарий слишком нагружен, он здесь и не нужен и его надо удалить

// ПУНКТ 9

func (l *LinkedList) Delete(n int, all bool) {
    // old:
    // проверим три случая при удалении объекта
    // из связного списка: начало, конец и середина
    if l.head == nil {
        return
    }

    tempNode := l.head
    var prev *Node

    if l.Count() == 1 && tempNode.value == n {
        l.Clean()
        return
    }


    // new:
    // в данном цикле мы будем проверять три случая при удалении объекта
    // из связного списка: начало, конец и середина
    for tempNode != nil {
        deleted := false
        if tempNode.value == n && tempNode == l.head {
            l.head = tempNode.next
            deleted = true
        } else if tempNode.value == n && tempNode == l.tail {
            prev.next = nil
            l.tail = prev
            deleted = true
        } else if tempNode.value == n {
            prev.next = tempNode.next
            deleted = true
        }
        if !all && deleted {
            return
        }
        if !deleted {
            prev = tempNode
        }
        tempNode = tempNode.next
    }
}
// комментарий не является локальным, так как проверка трех случаев идет позже
// перенесем комментарий ближе к точке использования

// ПУНКТ 11

// func Get_equipments_of_type_with_company_and_object(object_id string, type_name string) (gorm.DB, error) {
	// logic...
//  }
// удалил закомментированный код

// func Get_Object(ctx *fiber.Ctx) models.Description{} {
// удалил закомментированный код

func Update_seller_and_vendors(ctx *fiber.Ctx) error {
    // old:

    // err_add := database.Add_seller_and_vendor_equipment_relation(request.SellerEquipment, request.VendorsEquipment)

    // if err_add != nil {

    //  return ctx.JSON(answer{

    //      Ok:    false,

    //      Code:  fiber.ErrBadRequest.Code,

    //      Error: err_add,

    //  })

    // }

}

// удалил закомментированный код

// func Analyze_columns(columns []string, row uint) (ColumnsValues, error) {
// }
// удалил закомментированный код
```

