package lesson3

import (
	"log"
	"time"
)

func Put_Data_From_Sbis_To_Tables(response_data *api.Result, folder_name string) (string, error) {
	var returnObjectID string
	dateNow := time.Now()

	for i := 0; i < len(response_data.Tenders); i++ {

		docs_size := len(response_data.Tenders[i].Docs)
		docs := make([]string, 0, docs_size)
		names_docs := make([]string, 0, docs_size)
		for j := 0; j < docs_size; j++ {
			docs = append(docs, response_data.Tenders[i].Docs[j].External_url)
			names_docs = append(names_docs, response_data.Tenders[i].Docs[j].Description)
		}

		description := models.Description{
			Number:          response_data.Tenders[i].Number,
			WorkName:        response_data.Tenders[i].Name,
			FullName:        response_data.Tenders[i].Name,
			FullAdress:      response_data.Tenders[i].Lots[0].Delivery_place,
			WorkAdress:      response_data.Tenders[i].Lots[0].Delivery_place,
			Link_to_sbis:    response_data.Tenders[i].Tender_sbis_url,
			Link_to_zakupki: response_data.Tenders[i].Tender_url,
			ContractNumber:  response_data.Tenders[i].Lots[0].Contract_number,
			ContractLink:    response_data.Tenders[i].Lots[0].Contract_url,
			ContractObject:  response_data.Tenders[i].Lots[0].Contract_subject,
			Name_Documents:  names_docs,
			Documents:       docs,
			Delivery_term:   response_data.Tenders[i].Lots[0].Delivery_term,
		}

		// 7.1
		// old name: is_repeat_object_desc
		// new name: repeatDescription
		referObjectID, repeatDescription, errRepeatObjectDesc := Is_repeat_object_desc(description)
		if errRepeatObjectDesc != nil {
			return "", errRepeatObjectDesc
		}

		if repeatDescription {
			returnObjectID = referObjectID
			break
		}

		prevObjectID, errGetPrevID := Get_Last_objectID()

		if errGetPrevID != nil {
			log.Println(errGetPrevID)
			return "", errGetPrevID
		}

		objectID, errGetID := data_processing.Generate_ObjectID(prevObjectID)

		if errGetID != nil {
			log.Println(errGetID)
			return "", errGetID
		}

		var work_path models.Work_Part

		returnObjectID = objectID

		object := models.Object{
			ObjectID:   objectID,
			FolderName: folder_name,
		}

		DataBase.DB.Create(&object)

		description.ObjectRefer = objectID
		DataBase.DB.Create(&description)

		winner := models.Companies{
			CompanyID:  response_data.Tenders[i].Winner_inn,
			Full_name:  response_data.Tenders[i].Winner_full_name,
			Short_name: response_data.Tenders[i].Winner_name,
			KPP:        response_data.Tenders[i].Winner_kpp,
			OGRN:       response_data.Tenders[i].Winner_ogrn,
			Contract:   response_data.Tenders[i].Lots[0].Contract_url,
		}

		if response_data.Tenders[i].Winner_inn != "" && response_data.Tenders[i].Winner_inn != " " {

			// 7.1
			// old name: is_repeat_company
			// new name: repeatCompany
			repeat, errRepeatCompany := Is_repeat_company(winner.CompanyID)

			if errRepeatCompany != nil {
				return "", errRepeatCompany
			}

			if !repeat {
				DataBase.DB.Create(&winner)
				winner_contact := models.Contact{
					CompanyRefer: &winner.CompanyID,
					Phones:       response_data.Tenders[i].Winner_phone,
					Email:        response_data.Tenders[i].Winner_email,
				}

				DataBase.DB.Create(&winner_contact)
			}

			work_path.WinnerReferID = &winner.CompanyID

		}

		initiator := models.Companies{
			CompanyID:  response_data.Tenders[i].Initiator_inn,
			Full_name:  response_data.Tenders[i].Initiator_full_name,
			Short_name: response_data.Tenders[i].Initiator_name,
			KPP:        response_data.Tenders[i].Initiator_kpp,
			OGRN:       response_data.Tenders[i].Initiator_ogrn,
		}

		// 7.1
		// old name: is_initiator_repeat
		// new name: repeat
		repeat, err_repeat_initiator := Is_repeat_company(initiator.CompanyID)
		if err_repeat_initiator != nil {
			return "", err_repeat_initiator
		}

		initiator_contact := models.Contact{
			CompanyRefer: &initiator.CompanyID,
			FIO:          response_data.Tenders[i].Contact_person_name,
			Phones:       []string{response_data.Tenders[i].Contact_phone},
			Email:        []string{response_data.Tenders[i].Contact_email},
		}

		if !repeat {
			DataBase.DB.Create(&initiator)
			DataBase.DB.Create(&initiator_contact)
		}

		get_initiator_contact, err_get_initiator_contact := Get_contact(initiator)
		if err_get_initiator_contact != nil {
			return "", err_get_initiator_contact
		}

		// organizer := models.Companies{
		// 	CompanyID:  response_data.Tenders[i].Organizer_inn,
		// 	Short_name: response_data.Tenders[i].Organizer_name,
		// 	Full_name:  response_data.Tenders[i].Organizer_full_name,
		// 	OGRN:       response_data.Tenders[i].Organizer_ogrn,
		// 	KPP:        response_data.Tenders[i].Organizer_kpp,
		// }

		// DataBase.DB.Create(&organizer)

		status := "active"

		work_path.ObjectRefer = objectID
		work_path.Status = &status
		work_path.InitiatorReferID = &initiator.CompanyID
		work_path.InitiatorContactReferID = &get_initiator_contact.UserID

		DataBase.DB.Create(&work_path)

	}

	SetObjectDate(returnObjectID, dateNow)

	return returnObjectID, nil

}

func Add_vendors(work_part_id uint, company models.Companies, object_id string, equipment_type string,
	equipments []models.Equipment) error {

	equipments_to_company := make([]models.Company_to_equipment, 0, len(equipments))
	equipments_to_object := make([]models.Equipment_to_object, 0, len(equipments))
	equipments_to_types := make([]models.Equipment_to_type, 0, len(equipments))

	// 7.1
	// old name: is_repeat_company
	// new name: repeat
	repeat, err_repeat_company := Is_repeat_company(company.CompanyID)
	if err_repeat_company != nil {
		return err_repeat_company
	}

	if !repeat {
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
		// is_repeat, err_repeat := Check_repeat_equipment(equipment)
		// if err_repeat != nil {
		// 	return err_repeat
		// }

		var equipment_id uint

		temp_equipment_id, err_get_equipment_id := Add_equipment(equipment)
		if err_get_equipment_id != nil {
			return err_get_equipment_id
		}

		equipment_id = temp_equipment_id

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

	// 7.1
	// old name: is_vendor_exist
	// new name: exist
	exist, err_exist_vendor := Is_vendor_exist(work_part_id, company.CompanyID)
	if err_exist_vendor != nil {
		return err_exist_vendor
	}

	if !exist {
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

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Delete(n int, all bool) {
	if l.head == nil {
		return
	}

	tempNode := l.head
	var prev *Node

	if l.Count() == 1 && tempNode.value == n {
		l.Clean()
		return
	}

	for tempNode != nil {
		// 7.1
		// old name: deleted
		// new name: deleted
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


package logic

import (
	"context"
	"fmt"
	"log/slog"
	auth_jwt "system_of_monitoring_statistics/services/auth/internal/jwt"
	"system_of_monitoring_statistics/services/auth/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserProvider interface {
	GetUser(ctx context.Context, login string) (models.User, error)
	IsLoginExist(ctx context.Context, login string) (bool, error)
}

type UserSaver interface {
	SaveUser(ctx context.Context,
		email string,
		password string,
		name string,
		surname string,
		sex string,
		height int64,
		weight int64,
		birth_date int64) (int64, error)
}

type Auth struct {
	log          *slog.Logger
	userProvider UserProvider
	userSaver    UserSaver
}

func New(log *slog.Logger,
	userProvider UserProvider,
	userSaver UserSaver,
) *Auth {
	return &Auth{
		log:          log,
		userProvider: userProvider,
		userSaver:    userSaver,
	}
}

func (a *Auth) Login(ctx context.Context,
	email string,
	password string,
) (string, error) {
	const operation string = "logic.Login"

	log := a.log.With(
		slog.String("operation", operation),
	)

	log.Info("attempting getting user")

	userData, errGetUserData := a.userProvider.GetUser(ctx, email)
	if errGetUserData != nil {
		// TODO: add specific error processing
		log.Warn("user not found", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(errGetUserData.Error()),
		})

		log.Error("user not found", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(errGetUserData.Error()),
		})

		return "", fmt.Errorf("%s: %w", operation, errGetUserData)
	}

	errComparePasswords := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if errComparePasswords != nil {
		log.Info("invalid credentials", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(errComparePasswords.Error()),
		})

		return "", fmt.Errorf("%s: %w", operation, errComparePasswords)
	}

	// generate jwt
	token, errGetToken := auth_jwt.GenerateJWT(userData, time.Hour*8)
	if errGetToken != nil {
		log.Error("err generate jwt", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(errGetToken.Error()),
		})
	}

	log.Info("user logged in successfully")

	return token, nil
}

func (a *Auth) Register(ctx context.Context,
	email string,
	password string,
	name string,
	surname string,
	sex string,
	height int64,
	weight int64,
	birthDate int64,
) (int64, error) {
	const operation string = "logic.Register"

	log := a.log.With(
		slog.String("operation", operation),
	)

	log.Info("starting loging user")
	log.Info("check login")

	// 7.1
	// old name: existLogin
	// new name: exist
	exist, errExistLogin := a.userProvider.IsLoginExist(ctx, email)

	if exist != nil {
		// TODO: add specific error processing
		log.Error("error with sql row", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(errExistLogin.Error()),
		})

		return -1, fmt.Errorf("%s: %w", operation, errExistLogin)
	}

	if existLogin {
		log.Info("login alredy exist")
		return -1, fmt.Errorf("%s: login exist", operation)
	}

	log.Info("generating password")
	userPass, errGeneratePass := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errGeneratePass != nil {
		log.Error("error with sql row", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(errGeneratePass.Error()),
		})
		return -1, fmt.Errorf("%s: %w", operation, errGeneratePass)
	}

	log.Info("saving user")
	newUserID, errSaveUser := a.userSaver.SaveUser(
		ctx, email, string(userPass),
		name, surname, sex, height,
		weight, birthDate,
	)
	if errSaveUser != nil {
		log.Error("error save user", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(errSaveUser.Error()),
		})
		return -1, fmt.Errorf("%s: %w", operation, errSaveUser)
	}

	log.Info("user saved successfuly")

	return newUserID, nil
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

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {

	var valueStr1, valueStr2 string
	
	// 7.1
	// old name: flagStr
	// new name: isStr
	isStr := false

	if value, ok := any(v1).(string); ok {
		valueStr1 = strings.Trim(value, " ")
		isStr = true
	}

	if value, ok := any(v2).(string); ok {
		valueStr2 = strings.Trim(value, " ")
		isStr = true
	}

	switch isStr {
	case false:
		if v1 < v2 {
			return -1
		}
		if v1 > v2 {
			return +1
		}

	case true:
		if valueStr1 < valueStr2 {
			return -1
		}
		if valueStr1 > valueStr2 {
			return +1
		}
	}

	return 0
}