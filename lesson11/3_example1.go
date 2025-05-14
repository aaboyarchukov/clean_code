package lesson11

import (
	"encoding/json"
	"fmt"
	"os"
)

// old:
// const HASH_1_KOEFF int = 17

func (bf *BloomFilter) Hash1(s string) int {
	sum := 0
	// new:
	const HASH_1_KOEFF int = 17
	for _, char := range s {
		code := int(char)
		sum += code * HASH_1_KOEFF
	}
	sum %= bf.filter_len
	return sum
}

// old:
// const HASH_2_KOEFF int = 223

func (bf *BloomFilter) Hash2(s string) int {
	sum := 0
	const HASH_2_KOEFF int = 223
	for _, char := range s {
		code := int(char)
		sum += code * HASH_2_KOEFF
	}
	sum %= bf.filter_len
	return sum
}

// old:
// const (
// 	CELL_FOR_NAME = iota
// 	CELL_FOR_UNITS
// 	CELL_FOR_COUNT
// 	CELL_FOR_DELIVERY
// 	CELL_FOR_SPECIFICATION
// 	CELL_FOR_ARTICLE_NUMBER
// 	CELL_FOR_DEADLINE
// 	CELL_FOR_PAYMENT_DATE
// 	CELL_FOR_PRICE
// 	CELL_FOR_COST
// )

// const INDX_OF_VALUE_IN_MAP = 0

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

	object_id := files.Value["object_id"][INDX_OF_VALUE_IN_MAP]
	equipment_type := files.Value["equipment_type"][INDX_OF_VALUE_IN_MAP]
	nds := files.Value["nds"][INDX_OF_VALUE_IN_MAP]

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
	// new:
	if err_get_object_info != nil {
		return ctx.JSON(answer{
			Code:   fiber.ErrBadRequest.Code,
			Status: "error",
			Error:  err_get_object_info,
			Body:   response{},
		})
	}

	work_part, err_get_work_part := database.Get_object_work_part(object_id)
	// new:
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
	// new:
	if err_get_winner_info != nil {
		return ctx.JSON(answer{
			Code:   fiber.ErrBadRequest.Code,
			Status: "error",
			Error:  err_get_winner_info,
			Body:   response{},
		})
	}

	types_equipment, err_get_types := database.Get_types_of_equipments_of_object(object_id)
	// new:
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

			temp_item := equipmentSellerAndVendor{
				SellerEquipment:  seller_equipment,
				VendorEquipments: result_vendors_equipments,
			}

			sellers = append(sellers, companySeller{
				CompanyInfo: company_for_seller,
				CompanyID:   company_for_seller.CompanyID,
				Equipments:  temp_item,
			})
		}

		vendors_array, err_get_vendors := database.Get_vendors(work_part.Work_PathID)
		// new:
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

		}

		temp_data_item := equipment_types{
			EquipmentType:   equipment_type.Name,
			EquipmentTypeID: equipment_type.EquipmentTypeID,
			Seller:          sellers,
			Vendors:         result_vendors,
		}

		data_array = append(data_array, temp_data_item)

	}

	// old:
	// if err_get_object_info != nil {
	// 	return ctx.JSON(answer{
	// 		Code:   fiber.ErrBadRequest.Code,
	// 		Status: "error",
	// 		Error:  err_get_object_info,
	// 		Body:   response{},
	// 	})
	// }
	// if err_get_winner_info != nil {
	// 	return ctx.JSON(answer{
	// 		Code:   fiber.ErrBadRequest.Code,
	// 		Status: "error",
	// 		Error:  err_get_winner_info,
	// 		Body:   response{},
	// 	})
	// }
	// if err_get_types != nil {
	// 	return ctx.JSON(answer{
	// 		Code:   fiber.ErrBadRequest.Code,
	// 		Status: "error",
	// 		Error:  err_get_types,
	// 		Body:   response{},
	// 	})
	// }
	// if err_get_vendors != nil {
	// 	return ctx.JSON(answer{
	// 		Code:   fiber.ErrBadRequest.Code,
	// 		Status: "error",
	// 		Error:  err_get_vendors,
	// 		Body:   response{},
	// 	})
	// }
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
