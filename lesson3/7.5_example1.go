package lesson3

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func (l *OrderedList[T]) Count() int {
	count := 0
	// 7.5
	// old name: tempNode
	// new name: currentNode
	currentNode := l.head

	for currentNode != nil {
		count++
		currentNode = currentNode.next
	}

	return count
}

func (restaurant *Restaurant) printWaiterStats() {
	fmt.Println("\nСтатистика работы официантов:")

	for waiterID, stats := range restaurant.waiterStats {
		stats.mu.Lock()
		// 7.5
		// we could remove all locale variable
		// before
		// tablesCount := len(stats.tablesServed)
		// ordersCount := stats.ordersCount
		// stats.mu.Unlock()

		// fmt.Printf("Официант #%d: обслужил %d столов, принял %d заказов\n",
		// 	waiterID, tablesCount, ordersCount)

		// after
		stats.mu.Unlock()

		fmt.Printf("Официант #%d: обслужил %d столов, принял %d заказов\n",
			waiterID, len(stats.tablesServed), stats.ordersCount)
	}
}

func Get_kp(ctx *fiber.Ctx) error {
	type answerResponse struct {
		Ok         bool           `json:"ok"`
		StatusCode int            `json:"code"`
		Errors     error          `json:"errors"`
		Data       Answer_from_kp `json:"data"`
	}
	path, ok := os.LookupEnv("PATH_KP_FILES")

	if !ok {
		return fmt.Errorf("error with reading .env virables")
	}
	files, err_multi := ctx.MultipartForm()

	if err_multi != nil {
		return err_multi
	}

	for _, file := range files.File["file"] {
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
	}

	// ...

	return ctx.JSON(answerResponse{
		Ok:         true,
		StatusCode: 200,
		Errors:     nil,
		Data:       objects,
	})
}

func Add_vendors(work_part_id uint, company models.Companies, object_id string, equipment_type string,
	equipments []models.Equipment) error {

	equipments_to_company := make([]models.Company_to_equipment, 0, len(equipments))
	equipments_to_object := make([]models.Equipment_to_object, 0, len(equipments))
	equipments_to_types := make([]models.Equipment_to_type, 0, len(equipments))
	is_repeat_company, err_repeat_company := Is_repeat_company(company.CompanyID)
	if err_repeat_company != nil {
		return err_repeat_company
	}

	if !is_repeat_company {
		err_add_company := Add_company(company)
		if err_add_company != nil {
			return err_add_company
		}
	}

	equipment_type_id, err_get_equipment_type := Get_equipment_type_id(equipment_type)
	if err_get_equipment_type != nil {
		return err_get_equipment_type
	}

	for _, equipment := range equipments {
		// 7.5
		// we could remove temp_equipment_id

		// before
		// var equipment_id uint

		// temp_equipment_id, err_get_equipment_id := Add_equipment(equipment)
		// if err_get_equipment_id != nil {
		// 	return err_get_equipment_id
		// }

		// equipment_id = temp_equipment_id

		// after

		equipment_id, err_get_equipment_id := Add_equipment(equipment)
		if err_get_equipment_id != nil {
			return err_get_equipment_id
		}

		equipments_to_types = append(equipments_to_types, models.Equipment_to_type{
			EquipmentReferID:      equipment_id,
			Equipment_typeReferID: equipment_type_id,
		})

		equipments_to_company = append(equipments_to_company, models.Company_to_equipment{
			CompanyID:        company,
			EquipmentReferID: equipment_id,
		})

		equipments_to_object = append(equipments_to_object, models.Equipment_to_object{
			ObjectReferID:    object_id,
			EquipmentReferID: equipment_id,
		})

	}

	if len(equipments_to_company) != 0 {
		object_equipments_company := DataBase.DB.Create(&equipments_to_company)
		if object_equipments_company.Error != nil {
			return object_equipments_company.Error
		}
	}

	if len(equipments_to_object) != 0 {
		object_equipments := DataBase.DB.Create(&equipments_to_object)
		if object_equipments.Error != nil {
			return object_equipments.Error
		}
	}

	if len(equipments_to_types) != 0 {
		object_types := DataBase.DB.Create(&equipments_to_types)
		if object_types.Error != nil {
			return object_types.Error
		}
	}

	is_vendor_exist, err_exist_vendor := Is_vendor_exist(work_part_id, company.CompanyID)
	if err_exist_vendor != nil {
		return err_exist_vendor
	}

	if !is_vendor_exist {
		object_vendor := DataBase.DB.Create(&models.Work_part_description{
			WorkPartReferID: work_part_id,
			CompanyReferID:  &company.CompanyID,
		})
		if object_vendor.Error != nil {
			return object_vendor.Error
		}
	}

	return nil
}

func Get_companies(ctx *fiber.Ctx) error {
	type object_companies struct {
		Object_id string             `json:"object_id"`
		Vendors   []models.Companies `json:"vendors"`
		Seller    models.Companies   `json:"seller"`
		Initiator models.Companies   `json:"initiator"`
		Organizer models.Companies   `json:"organizer"`
	}

	type answer struct {
		Code   uint               `json:"code"`
		Status string             `json:"status"`
		Error  error              `json:"error"`
		Body   []object_companies `json:"body"`
	}

	work_parts, err_get_objects := database.Get_work_parts()

	if err_get_objects != nil {
		return ctx.JSON(
			answer{
				Code:   500,
				Status: "error",
				Error:  err_get_objects,
				Body:   []object_companies{},
			},
		)
	}

	result := make([]object_companies, 0, len(work_parts))

	for _, work_part := range work_parts {
		// 7.5
		// old name: item
		// new name: objectCompanies
		var objectCompanies object_companies
		objectCompanies.Object_id = work_part.ObjectRefer
		var winner, initiator, organizer models.Companies

		if work_part.WinnerReferID == nil {
			winner = models.Companies{}
		} else {
			get_winner, err_get_winner := database.Get_company(*work_part.WinnerReferID)
			if err_get_winner != nil {
				return ctx.JSON(
					answer{
						Code:   500,
						Status: "error",
						Error:  err_get_objects,
						Body:   []object_companies{},
					},
				)
			}

			winner = get_winner
		}

		if work_part.InitiatorReferID == nil {
			initiator = models.Companies{}
		} else {
			get_initiator, err_get_initiator := database.Get_company(*work_part.InitiatorReferID)
			if err_get_initiator != nil {
				return ctx.JSON(
					answer{
						Code:   500,
						Status: "error",
						Error:  err_get_objects,
						Body:   []object_companies{},
					},
				)
			}

			initiator = get_initiator
		}

		if work_part.OrganizerReferID == nil {
			organizer = models.Companies{}
		} else {
			get_organaizer, err_get_organizer := database.Get_company(*work_part.OrganizerReferID)
			if err_get_organizer != nil {
				return ctx.JSON(
					answer{
						Code:   500,
						Status: "error",
						Error:  err_get_objects,
						Body:   []object_companies{},
					},
				)
			}

			organizer = get_organaizer
		}

		get_vendors_id, err_get_vendors := database.Get_object_vendors(work_part.Work_PathID)
		if err_get_vendors != nil {
			return ctx.JSON(
				answer{
					Code:   500,
					Status: "error",
					Error:  err_get_objects,
					Body:   []object_companies{},
				},
			)
		}

		vendors := make([]models.Companies, 0, len(get_vendors_id))

		for _, vendor_id := range get_vendors_id {
			get_vendor_company, err_get_vendor_company := database.Get_company(vendor_id)
			if err_get_vendor_company != nil {
				return ctx.JSON(
					answer{
						Code:   500,
						Status: "error",
						Error:  err_get_objects,
						Body:   []object_companies{},
					},
				)
			}

			vendors = append(vendors, get_vendor_company)
		}

		objectCompanies.Seller = winner
		objectCompanies.Initiator = initiator
		objectCompanies.Organizer = organizer
		objectCompanies.Vendors = vendors

		result = append(result, objectCompanies)
	}
	return ctx.JSON(
		answer{
			Code:   200,
			Status: "success",
			Error:  nil,
			Body:   result,
		},
	)

}

func Get_data_for_tables(ctx *fiber.Ctx) error {
	type equipmentForVendor struct {
		CompanyInfo models.Companies `json:"company_info"`
		Equipment   models.Equipment `json:"equipment"`
	}
	type equipmentSellerAndVendor struct {
		SellerEquipment  models.Equipment     `json:"seller_equipment"`
		VendorEquipments []equipmentForVendor `json:"vendor_equipments"`
	}

	type companySeller struct {
		CompanyInfo models.Companies         `json:"company_info"`
		CompanyID   string                   `json:"company_id"`
		Equipments  equipmentSellerAndVendor `json:"equipments"`
	}

	type companyVendor struct {
		CompanyInfo models.Companies   `json:"company_info"`
		CompanyID   string             `json:"company_id"`
		Equipments  []models.Equipment `json:"equipments"`
	}

	type object struct {
		FullName   string `json:"full_name"`
		FullAdress string `json:"full_adress"`
		Date       string `json:"date"`
	}

	type equipment_types struct {
		EquipmentType   string `json:"equipment_type"`
		EquipmentTypeID uint   `json:"equipment_type_id"`
		Seller          []companySeller
		Vendors         []companyVendor `json:"vendors"`
	}

	type response struct {
		Work_part      models.Work_Part  `json:"work_part"`
		Object_info    object            `json:"object"`
		EquipmentTypes []equipment_types `json:"equipment_types"`
		Winner         models.Companies  `json:"winner"`
	}

	type answer struct {
		Code   int      `json:"code"`
		Status string   `json:"status"`
		Error  error    `json:"error"`
		Body   response `json:"body"`
	}

	type request struct {
		ID string `json:"id"`
	}

	object_id_struct := request{}

	id := ctx.Request().Body()

	if err_json_unmarshall := json.Unmarshal(id, &object_id_struct); err_json_unmarshall != nil {
		return ctx.JSON(answer{
			Code:   fiber.ErrBadRequest.Code,
			Status: "error",
			Error:  err_json_unmarshall,
			Body:   response{},
		})
	}

	object_id := object_id_struct.ID

	object_info, err_get_object_info := database.Get_object_info(object_id)
	if err_get_object_info != nil {
		return ctx.JSON(answer{
			Code:   fiber.ErrBadRequest.Code,
			Status: "error",
			Error:  err_get_object_info,
			Body:   response{},
		})
	}

	work_part, err_get_work_part := database.Get_object_work_part(object_id)
	if err_get_work_part != nil {
		return ctx.JSON(answer{
			Code:   fiber.ErrBadRequest.Code,
			Status: "error",
			Error:  err_get_work_part,
			Body:   response{},
		})
	}

	winner_id := ""
	if work_part.WinnerReferID != nil {
		winner_id = *work_part.WinnerReferID
	}

	winner_info, err_get_winner_info := database.Get_company(winner_id)
	if err_get_winner_info != nil {
		return ctx.JSON(answer{
			Code:   fiber.ErrBadRequest.Code,
			Status: "error",
			Error:  err_get_winner_info,
			Body:   response{},
		})
	}

	types_equipment, err_get_types := database.Get_types_of_equipments_of_object(object_id)
	if err_get_types != nil {
		return ctx.JSON(answer{
			Code:   fiber.ErrBadRequest.Code,
			Status: "error",
			Error:  err_get_types,
			Body:   response{},
		})
	}

	data_array := make([]equipment_types, 0, len(types_equipment))
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
			}

			// 7.5
			// old name: temp_item
			// new name: sellerAndVendorEquipments
			sellerAndVendorEquipments := equipmentSellerAndVendor{
				SellerEquipment:  seller_equipment,
				VendorEquipments: result_vendors_equipments,
			}

			sellers = append(sellers, companySeller{
				CompanyInfo: company_for_seller,
				CompanyID:   company_for_seller.CompanyID,
				Equipments:  sellerAndVendorEquipments,
			})
		}

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

			// 7.5
			// old name: temp_vendor
			// new name: vendorCompany
			vendorCompany := companyVendor{
				CompanyInfo: vendor_info,
				CompanyID:   vendor,
				Equipments:  vendor_equipments,
			}

			result_vendors = append(result_vendors, vendorCompany)

		}

		// 7.5
		// old name: temp_data_item
		// new name: equipmentTypesData
		equipmentTypesData := equipment_types{
			EquipmentType:   equipment_type.Name,
			EquipmentTypeID: equipment_type.EquipmentTypeID,
			Seller:          sellers,
			Vendors:         result_vendors,
		}

		data_array = append(data_array, equipmentTypesData)

	}

	return ctx.JSON(answer{
		Code:   200,
		Status: "success",
		Body: response{
			Winner: winner_info,
			Object_info: object{
				FullName:   object_info.FullName,
				FullAdress: object_info.FullAdress,
				Date:       object_info.Delivery_term,
			},
			Work_part:      work_part,
			EquipmentTypes: data_array,
		},
	})

}
