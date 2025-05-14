# Время жизни переменных

Продолжаем впитывать рекомендации по курсу

1. Минимизируйте область видимости переменных

При использовании переменных надо стараться минимизировать область их видимости, таким образом вы сокращается время их жизни.

2. Группируйте связанные команды

При работе с одной переменной, необходимо группировать связанные с ней команды, которые используются, иначе, если между командами будет большое "окно", ваша переменная будет более уязвима к изменениям и в итоге это приведет к неверной обработке данных

Задания:

[3_example1.go](https://github.com/aaboyarchukov/clean_code/blob/master/lesson11/3_example1.go)

```go
// old:

// const HASH_1_KOEFF int = 17
func (bf *BloomFilter) Hash1(s string) int {

    // new:
    const HASH_1_KOEFF int = 17
    for _, char := range s {
        code := int(char)
        sum += code * HASH_1_KOEFF
    }
	// ...
}
// в данном примере я снизил область видимости константы, так как
// она используется только в одной функции,
// расположив ее максимально близко к месту использования

// const HASH_2_KOEFF int = 223
func (bf *BloomFilter) Hash2(s string) int {

    // new:
    const HASH_2_KOEFF int = 223
    for _, char := range s {
        code := int(char)
        sum += code * HASH_2_KOEFF
    }
	// ...
}
// в данном примере я снизил область видимости константы, так как
// она используется только в одной функции,
// расположив ее максимально близко к месту использования

// old:

// const (
//  CELL_FOR_NAME = iota
//  CELL_FOR_UNITS
//  CELL_FOR_COUNT
//  CELL_FOR_DELIVERY
//  CELL_FOR_SPECIFICATION
//  CELL_FOR_ARTICLE_NUMBER
//  CELL_FOR_DEADLINE
//  CELL_FOR_PAYMENT_DATE
//  CELL_FOR_PRICE
//  CELL_FOR_COST
// )
// const INDX_OF_VALUE_IN_MAP = 0
func Get_kp(ctx *fiber.Ctx) error {
    // some logic...
    // new:
    const (
        CELL_FOR_NAME = iota
        CELL_FOR_UNITS
        CELL_FOR_COUNT
        CELL_FOR_DELIVERY
        CELL_FOR_SPECIFICATION
        CELL_FOR_ARTICLE_NUMBER
        CELL_FOR_DEADLINE
        CELL_FOR_PAYMENT_DATE
        CELL_FOR_PRICE
        CELL_FOR_COST
    )
    const INDX_OF_VALUE_IN_MAP = 0
    // some logic...
}
// в данном примере я снизил область видимости константы, так как
// она используется только в одной функции,
// расположив ее максимально близко к месту использования


func Get_data_for_tables(ctx *fiber.Ctx) error {

  
	// some logic
    object_info, err_get_object_info := database.Get_object_info(object_id)

    // new:
    if err_get_object_info != nil {
        return ctx.JSON(answer{
            Code:   fiber.ErrBadRequest.Code,
            Status: "error",
            Error:  err_get_object_info,
            Body:   response{},
        })
    }

    work_part, err_get_work_part := database.Get_object_work_part(object_id)

    // new:
    if err_get_work_part != nil {
        return ctx.JSON(answer{
            Code:   fiber.ErrBadRequest.Code,
            Status: "error",
            Error:  err_get_work_part,
            Body:   response{},
        })
    }

    winner_info, err_get_winner_info := database.Get_company(winner_id)

    // new:
    if err_get_winner_info != nil {
        return ctx.JSON(answer{
            Code:   fiber.ErrBadRequest.Code,
            Status: "error",
            Error:  err_get_winner_info,
            Body:   response{},
        })
    }

    types_equipment, err_get_types := database.Get_types_of_equipments_of_object(object_id)

    // new:
    if err_get_types != nil {
        return ctx.JSON(answer{
            Code:   fiber.ErrBadRequest.Code,
            Status: "error",
            Error:  err_get_types,
            Body:   response{},
        })
    }

  

    

        vendors_array, err_get_vendors := database.Get_vendors(work_part.Work_PathID)

        // new:
        if err_get_vendors != nil {
            return ctx.JSON(answer{
                Code:   fiber.ErrBadRequest.Code,
                Status: "error",
                Error:  err_get_vendors,
                Body:   response{},
            })
        }

    // old:
    // if err_get_object_info != nil {
    //  return ctx.JSON(answer{
    //      Code:   fiber.ErrBadRequest.Code,
    //      Status: "error",
    //      Error:  err_get_object_info,
    //      Body:   response{},
    //  })
    // }
    // if err_get_winner_info != nil {
    //  return ctx.JSON(answer{
    //      Code:   fiber.ErrBadRequest.Code,
    //      Status: "error",
    //      Error:  err_get_winner_info,
    //      Body:   response{},
    //  })
    // }
    // if err_get_types != nil {
    //  return ctx.JSON(answer{
    //      Code:   fiber.ErrBadRequest.Code,
    //      Status: "error",
    //      Error:  err_get_types,
    //      Body:   response{},
    //  })
    // }
    // if err_get_vendors != nil {
    //  return ctx.JSON(answer{
    //      Code:   fiber.ErrBadRequest.Code,
    //      Status: "error",
    //      Error:  err_get_vendors,
    //      Body:   response{},
    //  })
    // }

    return ctx.JSON(answer{
        Code:   200,
        Status: "success",
        Body: response{
            Winner: winner_info,
            Object_info: object{
                FullName:   object_info.FullName,
                FullAdress: object_info.FullAdress,
                Date:       object_info.Delivery_term,
            },
            Work_part:      work_part,
            EquipmentTypes: data_array,
        },
    })
}

// в данных примерах я сгруппировал обработку ошибок
// близко к объявлению переменной с ошибкой,
// тем самым я уменьшил окно уязвимости между объявлением переменных с ошибками
// и их обработкой
```

