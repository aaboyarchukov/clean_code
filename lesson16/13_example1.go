package lesson16

import "encoding/json"

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
	filter     int64
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
	app.Post("api/post_sbis_api", routes_handlers.Get_Data_From_Sbis)
	app.Post("api/new_objects", database.Get_Data_About_New_Object)
	app.Post("api/docs", database.Get_Docs)
	app.Get("api/folders_name", database.Get_Folders_Name)
	app.Post("api/dadata", routes_handlers.Get_DaData)
	app.Post("api/edit_equipment", routes_handlers.Equipment_Edit)
	app.Post("api/get_object", routes_handlers.Get_object)
	app.Get("api/get_catalog_equipment", routes_handlers.Get_Catalog)
	app.Post("api/add_catalog_equipment", database.Add_catalog_equipment)
	app.Post("api/add_contact", database.Add_contact)
	app.Post("api/excel_files_objects", excel.GetObjectFiles)
	app.Post("api/excel_files_equipment", excel.GetEquipmentFiles)
	app.Get("api/get_contacts", database.Get_Contacts)
	app.Get("api/get_equip_id", database.Get_Last_Equip_Id)
	app.Post("api/kp", routes_handlers.Get_columns)
	app.Post("api/kp_with_columns", excel.Get_kp)
	app.Post("api/get_data_for_table", routes_handlers.Get_data_for_tables)
	app.Post("api/delete_object", routes_handlers.Delete_object)
	app.Get("api/companies", routes_handlers.Get_companies)
	app.Post("api/company", routes_handlers.Get_company)
	app.Get("api/all_companies", routes_handlers.Get_all_companies)
	app.Post("api/update_seller_and_vendors", routes_handlers.Update_seller_and_vendors)
	app.Post("api/update_equipments_price", routes_handlers.UpdateEquipmentsPrice)
	app.Post("api/update_equipments_count", routes_handlers.UpdateEquipmentCount)
	app.Post("api/delete_vendor_to_seller_relation", routes_handlers.Delete_Vendor_To_Seller_Relation)
	app.Post("api/delete_equipments", routes_handlers.DeleteEquipments)
	app.Post("api/get_kp_file", excel.FormingExcelFile)
	app.Post("api/get_events", routes_handlers.GetEvents)
	app.Post("api/set_object_date", routes_handlers.SetObjectDate)
	app.Post("api/get_object_date", routes_handlers.GetObjectDate)
	app.Post("api/actions", routes_handlers.GetActivities)
	app.Post("api/set_actions", routes_handlers.SetActivities)
	app.Post("api/get_winner", routes_handlers.GetWinner)
	app.Post("api/get_request_kp_file", excel.FormingExcelFileForRequest)
	app.Post("api/get_defend_kp_file", excel.FormingExcelFileForDefend)
	app.Post("api/defence", routes_handlers.GetDefence)
	app.Post("api/set_defend", routes_handlers.SetDefend)
	app.Post("api/login", routes_handlers.LoginUser)
	app.Post("api/search_object", routes_handlers.SearchObject)
}

// присутствует ненужный маркер, а также комментарий - их надо удалить

// ПУНКТ 6

func GetDataForTables() {
	// logic...
	for _, equipment_type := range types_equipment {
		seller_equipments, err_get_equipments := database.Get_equipments_of_type_and_object(object_id, equipment_type.EquipmentTypeID)
		if err_get_equipments != nil {
			return ctx.JSON(answer{
				Code:   fiber.ErrBadRequest.Code,
				Status: "error",
				Error:  err_get_equipments,
				Body:   response{},
			})
		}

		sellers := make([]companySeller, 0, len(seller_equipments))
		for _, seller_equipment := range seller_equipments {
			company_for_seller, err_get_company := database.Get_equipment_company(seller_equipment.EquipmentID)
			if err_get_company != nil {
				return ctx.JSON(answer{
					Code:   fiber.ErrBadRequest.Code,
					Status: "error",
					Error:  err_get_company,
					Body:   response{},
				})
			}
			vendor_equipments_ids, err_get_vendor_equipments := database.GetVendorsForEquipment(seller_equipment, object_id)
			if err_get_vendor_equipments != nil {
				return ctx.JSON(answer{
					Code:   fiber.ErrBadRequest.Code,
					Status: "error",
					Error:  err_get_vendor_equipments,
					Body:   response{},
				})
			}

			result_vendors_equipments := make([]equipmentForVendor, 0, len(vendor_equipments_ids))
			for _, vendorEquipmentId := range vendor_equipments_ids {
				vendorEquipmentCompanyInfo, err_get_info := database.Get_equipment_company(uint(vendorEquipmentId))
				if err_get_info != nil {
					return ctx.JSON(answer{
						Code:   fiber.ErrBadRequest.Code,
						Status: "error",
						Error:  err_get_info,
						Body:   response{},
					})
				}

				vendorEquipment, err_get_vendor_equipment := database.GetEquipment(uint(vendorEquipmentId))

				if err_get_vendor_equipment != nil {
					return ctx.JSON(answer{
						Code:   fiber.ErrBadRequest.Code,
						Status: "error",
						Error:  err_get_vendor_equipment,
						Body:   response{},
					})
				}

				result_vendors_equipments = append(result_vendors_equipments,
					equipmentForVendor{
						CompanyInfo: vendorEquipmentCompanyInfo,
						Equipment:   vendorEquipment,
					})
			} // закрытие первого внутреннего цикла в первом внутреннем цикле

			temp_item := equipmentSellerAndVendor{
				SellerEquipment:  seller_equipment,
				VendorEquipments: result_vendors_equipments,
			}

			sellers = append(sellers, companySeller{
				CompanyInfo: company_for_seller,
				CompanyID:   company_for_seller.CompanyID,
				Equipments:  temp_item,
			})
		} // закрытие первого внутреннего цикла

		vendors_array, err_get_vendors := database.Get_vendors(work_part.Work_PathID)
		if err_get_vendors != nil {
			return ctx.JSON(answer{
				Code:   fiber.ErrBadRequest.Code,
				Status: "error",
				Error:  err_get_vendors,
				Body:   response{},
			})
		}

		result_vendors := make([]companyVendor, 0, len(vendors_array))

		for _, vendor := range vendors_array {
			vendor_info, err_get_vendor_info := database.Get_company(vendor)
			if err_get_vendor_info != nil {
				return ctx.JSON(answer{
					Code:   fiber.ErrBadRequest.Code,
					Status: "error",
					Error:  err_get_vendor_info,
					Body:   response{},
				})
			}

			vendor_equipments, err_get_vendor_equipments := database.Get_equipments_for_company(equipment_type.Name, object_id, vendor)

			if err_get_vendor_equipments != nil {
				return ctx.JSON(answer{
					Code:   fiber.ErrBadRequest.Code,
					Status: "error",
					Error:  err_get_vendor_equipments,
					Body:   response{},
				})
			}

			if len(vendor_equipments) == 0 {
				continue
			}

			temp_vendor := companyVendor{
				CompanyInfo: vendor_info,
				CompanyID:   vendor,
				Equipments:  vendor_equipments,
			}

			result_vendors = append(result_vendors, temp_vendor)

		} // закрытие второго внутреннего цикла

		temp_data_item := equipment_types{
			EquipmentType:   equipment_type.Name,
			EquipmentTypeID: equipment_type.EquipmentTypeID,
			Seller:          sellers,
			Vendors:         result_vendors,
		}

		data_array = append(data_array, temp_data_item)

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
	sum := 0
	for _, char := range s {
		code := int(char)
		sum += code * HASH_2_KOEFF
	}
	sum %= bf.filter_len
	return sum
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
// 	object := DataBase.DB.Model(&models.Equipment{}).
// 		Joins("LEFT JOIN company_to_equipments ON equipment.equipment_id = company_to_equipments.equipment_refer_id").
// 		Joins("JOIN equipment_to_types ON equipment.equipment_id = equipment_to_types.equipment_refer_id").
// 		Joins("JOIN equipment_types ON equipment_types.equipment_type_id = equipment_to_types.equipment_type_refer_id").
// 		Joins("JOIN equipment_to_objects ON equipment.equipment_id = equipment_to_objects.equipment_refer_id").
// 		Where("equipment_types.name = ?", type_name).
// 		Where("equipment_to_objects.object_refer_id = ?", object_id)
// 	if object.Error != nil {
// 		return gorm.DB{}, object.Error
// 	}

// 	return *object, nil
// }
// удалил закомментированный код

// func Get_Object(ctx *fiber.Ctx) models.Description{} {
// 	type object_id struct {
// 		Id string `json:"id"`
// 	}

// 	req_body := ctx.Request().Body()
// 	resp_body := object_id{}

// 	err_json_unmarshall := json.Unmarshal(req_body, &resp_body)

// 	if err_json_unmarshall != nil {
// 		log.Println(err_json_unmarshall)
// 	}

// 	fields := `descriptions.*`
// 	join_object := models.Description{}
// 	error_new_object := DataBase.DB.Model(&models.Object{}).Select(fields).
// 		Where("object_id = ?", resp_body.Id).
// 		Joins("JOIN descriptions ON descriptions.object_refer = objects.object_id").
// 		Scan(&join_object).Error

// 	if error_new_object != nil {
// 		log.Println(error_new_object)
// 	}

// 	return ctx.JSON(join_object)

// }
// удалил закомментированный код

func Update_seller_and_vendors(ctx *fiber.Ctx) error {
	type answer struct {
		Ok    bool  `json:"ok"`
		Code  int   `json:"code"`
		Error error `json:"error"`
	}

	request := struct {
		SellerEquipment  models.Equipment   `json:"seller_equipment"`
		VendorsEquipment []models.Equipment `json:"vendors_equipment"`
		ID               string             `json:"object_id"`
	}{}

	if err_unmarshal := json.Unmarshal(
		ctx.Request().Body(),
		&request); err_unmarshal != nil {
		return ctx.JSON(answer{
			Ok:    false,
			Code:  fiber.ErrBadRequest.Code,
			Error: err_unmarshal,
		})
	}

	// old:
	// err_add := database.Add_seller_and_vendor_equipment_relation(request.SellerEquipment, request.VendorsEquipment)
	// if err_add != nil {
	// 	return ctx.JSON(answer{
	// 		Ok:    false,
	// 		Code:  fiber.ErrBadRequest.Code,
	// 		Error: err_add,
	// 	})
	// }

	err_add := database.Add_equipment_and_vendors_relation(request.SellerEquipment, request.ID, request.VendorsEquipment)
	if err_add != nil {
		return ctx.JSON(answer{
			Ok:    false,
			Code:  fiber.ErrBadRequest.Code,
			Error: err_add,
		})
	}

	return ctx.JSON(answer{
		Ok:    true,
		Code:  200,
		Error: nil,
	})
}

// удалил закомментированный код

// func Analyze_columns(columns []string, row uint) (ColumnsValues, error) {
// 	needed_columns, err_get_needed_columns := database.Get_needed_columns()
// 	if err_get_needed_columns != nil {
// 		return ColumnsValues{}, err_get_needed_columns
// 	}

// 	equipment_columns := ColumnsValues{}
// 	result := make([]Point, 0, len(needed_columns))

// 	for _, needed_column := range needed_columns {
// 		for ind := 0; ind < len(columns); ind++ {
// 			if columns[ind] == "" {
// 				ind++
// 			}

// 			target_column := columns[ind]

// 			is_match, err_is_match := database.Is_columns_match(strings.ToLower(needed_column.ColumnName), strings.ToLower(target_column))
// 			if err_is_match != nil {
// 				return ColumnsValues{}, err_is_match
// 			}

// 			if is_match {
// 				result = append(result, Point{
// 					Row:    int(row),
// 					Column: ind,
// 				})
// 				break
// 			}
// 		}
// 	}

// 	if len(result) < 3 {
// 		return equipment_columns, fmt.Errorf("some columns not math with needed columns")
// 	}

// 	equipment_columns.EquipmentName = result[0]
// 	equipment_columns.EquipmentCount = result[1]
// 	equipment_columns.EquipmentCost = result[2]
// 	fmt.Printf("%+v", equipment_columns)
// 	return equipment_columns, nil

// }
// удалил закомментированный код
