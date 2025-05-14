package lesson13

import (
	"container/list"
	"log"
	"os"
	"strings"
)

func DeleteAllOfObject(object_id string) (bool, error) {
	equipments, err_get_equipments := GetEquipmentsOfObject(object_id)
	if err_get_equipments != nil {
		return false, err_get_equipments
	}

	// array example
	for _, equipment := range equipments {
		error_delete := DataBase.DB.Delete(&models.Equipment{}, &models.Equipment{
			EquipmentID: equipment,
		}).Error

		if error_delete != nil {
			return false, error_delete
		}
	}
	// to delete equipment without object:
	// delete from equipment
	// where equipment_id in
	// (select equipment.equipment_id from equipment
	// 		left join equipment_to_objects
	// 			ON equipment.equipment_id = equipment_to_objects.equipment_refer_id
	// 			where equipment_to_objects.equipment_refer_id IS NULL);

	return true, nil
}

func SumPeopleAges(peopleAges []int) int {
	var sumPeopleAges int = 0
	// old:
	// for i := 0; i < len(peopleAges); i++ {
	// 	sumPeopleAges += peopleAges[i]
	// }

	// new:
	for _, age := range peopleAges {
		sumPeopleAges += age
	}
	return sumPeopleAges
}

func Contain(value string) string {
	textMap := make(map[string]string)
	filePath, _ := os.LookupEnv("INFO_TXT")

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return ""
	}

	stringData := string(fileData)
	arrRows := strings.Split(stringData, "---")
	// old:
	// for ind := 0; ind < len(arrRows); ind++ {
	// 	item := strings.TrimSpace(arrRows[ind])
	// 	arrItems := strings.Split(item, "=")
	// 	textMap[arrItems[0]] = arrItems[1]
	// }

	// new:
	for _, item := range arrRows {
		item = strings.TrimSpace(item)
		arrItems := strings.Split(item, "=")
		// new
		var itemsList list.List
		for _, item := range arrItems {
			itemsList.PushBack(item)
		}

		listFront, listBack := itemsList.Front().Value, itemsList.Back().Value
		var mapKey, mapValue string = "", ""

		if listValue, ok := listFront.(string); ok {
			mapKey = listValue
		}

		if listValue, ok := listBack.(string); ok {
			mapValue = listValue
		}

		textMap[mapKey] = mapValue
	}

	return textMap[value]
}

func RowWithPhonesToArray(rowWithPhones string) []string {
	phones_array := strings.Split(rowWithPhones, " ")
	var phones_result []string
	// old:
	// for ind := 0; ind < len(phones_array); ind++ {
	// 	if phones_array[ind] != "" {
	// 		phones_result = append(phones_result, phones_array[ind])
	// 	}

	// }

	// new:
	for _, phone := range phones_array {
		if phone != "" {
			phones_result = append(phones_result, phone)
		}

	}

	return phones_result
}
