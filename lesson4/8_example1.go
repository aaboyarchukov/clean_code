package lesson4

import (
	"constraints"
	"encoding/json"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func GetLinkedList(values []int) *LinkedList {
	// 8
	// old name: resultLL
	// new name: linkedList
	var linkedList LinkedList // resulting linked list
	for _, value := range values {
		linkedList.AddInTail(Node{
			value: value,
		})
	}
	return &linkedList
}

// 8
// old name: l1, l2
// new name: firstList, secondList
func EqualLists(firstList *LinkedList, secondList *LinkedList) bool {
	// equals len and elements
	if firstList.head == nil &&
		secondList.head == nil {
		return true
	}

	if firstList.head.value != secondList.head.value {
		return false
	}
	if firstList.tail.value != secondList.tail.value {
		return false
	}

	// 8
	// old name: countLL1, counLL2
	// new name: firstListSize, secondListSize
	firstListSize, secondListSize := firstList.Count(), secondList.Count()

	if firstListSize == secondListSize {
		// 8
		// old name: tempLL1, tempLL2
		// new name: currNodeFirstList, currNodeSecondList
		currNodeFirstList, currNodeSecondList := firstList.head, secondList.head

		for currNodeFirstList != nil && currNodeSecondList != nil {
			if currNodeFirstList.value != currNodeSecondList.value {
				return false
			}
			currNodeFirstList = currNodeFirstList.next
			currNodeSecondList = currNodeSecondList.next
		}

		return true
	}

	return false
}

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head *Node[T]
	tail *Node[T]
	// base       []T
	_ascending bool
}

func (l *OrderedList[T]) Count() int {
	// 8
	// old name: count
	// new name: nodesCount
	nodesCount := 0

	// 8
	// old name: tempNode
	// new name: currentNode
	currentNode := l.head

	for currentNode != nil {
		nodesCount++
		currentNode = currentNode.next
	}

	return nodesCount
}

func (l *OrderedList[T]) ToArray() []T {
	// 8
	// old name: result
	// new name: arrayFromList
	arrayFromList := make([]T, 0, l.Count())

	// 8
	// old name: tempNode
	// new name: currentNode
	currentNode := l.head

	for currentNode != nil {
		arrayFromList = append(arrayFromList, currentNode.value)
		currentNode = currentNode.next
	}

	return arrayFromList
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

	// 8
	// old name: object_id_struct
	// new name: requestStruct
	requestStruct := request{}

	// 8
	// old name: id
	// new name: requestBody
	requestBody := ctx.Request().Body()

	if err_json_unmarshall := json.Unmarshal(requestBody, &requestStruct); err_json_unmarshall != nil {
		return ctx.JSON(answer{
			Code:   fiber.ErrBadRequest.Code,
			Status: "error",
			Error:  err_json_unmarshall,
			Body:   response{},
		})
	}

	// 8
	// old name: object_id
	// new name: tenderID
	tenderID := requestStruct.ID

	object_info, err_get_object_info := database.Get_object_info(tenderID)
	if err_get_object_info != nil {
		return ctx.JSON(answer{
			Code:   fiber.ErrBadRequest.Code,
			Status: "error",
			Error:  err_get_object_info,
			Body:   response{},
		})
	}

	work_part, err_get_work_part := database.Get_object_work_part(tenderID)
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

	// 8
	// old name: types_equipment
	// new name: tenderEquipmentTypes
	tenderEquipmentTypes, errGetTypes := database.Get_types_of_equipments_of_object(tenderID)
	if errGetTypes != nil {
		return ctx.JSON(answer{
			Code:   fiber.ErrBadRequest.Code,
			Status: "error",
			Error:  errGetTypes,
			Body:   response{},
		})
	}

	// 8
	// old name: data_array
	// new name: equipmentTypesGroup
	equipmentTypesGroup := make([]equipment_types, 0, len(tenderEquipmentTypes))
	for _, equipment_type := range tenderEquipmentTypes {
		seller_equipments, err_get_equipments := database.Get_equipments_of_type_and_object(tenderID, equipment_type.EquipmentTypeID)
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
			vendor_equipments_ids, err_get_vendor_equipments := database.GetVendorsForEquipment(seller_equipment, tenderID)
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

		// 8
		// old name: vendors_array
		// new name: tenderVendors
		tenderVendors, errGetVendors := database.Get_vendors(work_part.Work_PathID)
		if errGetVendors != nil {
			return ctx.JSON(answer{
				Code:   fiber.ErrBadRequest.Code,
				Status: "error",
				Error:  errGetVendors,
				Body:   response{},
			})
		}

		// 8
		// old name: result_vendors
		// new name: vendorsGroup
		vendorsGroup := make([]companyVendor, 0, len(tenderVendors))

		for _, vendor := range tenderVendors {
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

			vendorCompany := companyVendor{
				CompanyInfo: vendor_info,
				CompanyID:   vendor,
				Equipments:  vendor_equipments,
			}

			vendorsGroup = append(vendorsGroup, vendorCompany)

		}

		equipmentTypesData := equipment_types{
			EquipmentType:   equipment_type.Name,
			EquipmentTypeID: equipment_type.EquipmentTypeID,
			Seller:          sellers,
			Vendors:         vendorsGroup,
		}

		equipmentTypesGroup = append(equipmentTypesGroup, equipmentTypesData)

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
			EquipmentTypes: equipmentTypesGroup,
		},
	})

}
