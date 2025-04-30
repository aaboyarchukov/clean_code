package lesson2

func UpdateCompanyData(companyID string, columnName string, companyData interface{}) (bool, error) {
	// old name: object
	// new name: updateCompanyObject
	updateCompanyObject := DataBase.DB.Model(&models.Companies{}).
		Where(&models.Companies{
			CompanyID: companyID,
		}).Update(columnName, companyData)
	if err := updateCompanyObject.Error; err != nil {
		return false, err
	}

	return true, nil
}

func GetEquipmentsOfObject(objectID string) ([]uint, error) {
	// old name: equipments
	// new name: getEquipments
	var getEquipments []models.Equipment_to_object

	// old name: result
	// new name: getEquipmentsIDs
	var getEquipmentsIDs []uint

	// old name: object
	// new name: getObjectEquipments
	getObjectEquipments := DataBase.DB.Model(&models.Equipment_to_object{}).Select("*").
		Where(&models.Equipment_to_object{
			ObjectReferID: objectID,
		}).
		Scan(&getEquipments)

	if err := getObjectEquipments.Error; err != nil {
		return []uint{}, err
	}

	for _, equipment := range getEquipments {
		getEquipmentsIDs = append(getEquipmentsIDs, equipment.EquipmentReferID)
	}

	return getEquipmentsIDs, nil
}

func UpdateWinner(objectID string, winnerID string) error {
	// old name: object
	// new name: updateWinnerCompany
	updateWinnerCompany := DataBase.DB.Model(&models.Work_Part{}).Where("work_parts.object_refer = ?", objectID).
		Update("winner_refer_id", winnerID)
	if err := updateWinnerCompany.Error; err != nil {
		return err
	}

	return nil
}

func GetActions(objectID string) ([]models.ActivitiesAndTypes, error) {
	// old name: result
	// new name: getActions
	var getActions []models.ActivitiesAndTypes

	// old name: object
	// new name: getActionsObject
	getActionsObject := DataBase.DB.Model(&models.ActivitiesAndTypes{}).Select("activities_and_types.*").
		Where("activities_and_types.object_refer_id = ?", objectID).Scan(&getActions)

	if err := getActionsObject.Error; err != nil {
		return getActions, err
	}

	return getActions, nil
}
