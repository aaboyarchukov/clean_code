package lesson15

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"
)

// формирование отчета о зарплате за определенный период
func (employee *Employee) PrepareSalaryReportByPeriod(period time.Time) Report {}

// отношение между стадиями обработки и типами оборудований по торгам
type ProccesingStagesAndEquipmentTypes struct {
	ID                   uint `gorm:"primaryKey"`
	TenderReferID        string
	TenderID             Tender `gorm:"foreignKey:TenderReferID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EquipmentTypeReferID uint
	TypeID               Equipment_type `gorm:"foreignKey:TypeReferID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Stages               pq.StringArray `gorm:"type:text[]"`
}

// отношение между оборудованием покупателей и поставщиков в тендере
type SellerAndVendorsEquipmentOfTender struct {
	ID                   uint `gorm:"primaryKey"`
	SellerEquipmentID    uint
	SellerEquipmentRefer Equipment `gorm:"foreignKey:SellerEquipmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VendorEquipmentID    uint
	VendorEquipmentRefer Equipment `gorm:"foreignKey:VendorEquipmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// история взаимодействия с компанией по тендеру
type HistoryInteractionWithCompanyOfTender struct {
	HistoryID      uint `gorm:"primaryKey"`
	CompanyReferID string
	CompanyID      Companies `gorm:"foreignKey:CompanyReferID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt      *time.Time
	History        string
}

// отношение описательной части рабочей части
// (процесс, когда тендер находится в обработке) тендера
type WorkPartDescriptionOfTender struct {
	ID              uint `gorm:"primaryKey"`
	WorkPartReferID uint
	WorkPartRefer   Work_Part `gorm:"foreignKey:WorkPartReferID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CompanyReferID  *string
	CompanyRefer    Companies `gorm:"foreignKey:CompanyReferID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Defence         string
	KPDocumentPath  string
}

// мы формируем коммерческое предложение по торгу
func GetCommericalProposal(object_id string, columns data_analyze.ColumnsValues, ndsStatus string) (Answer_from_kp, error) {
}

// мы ищем текстовые выражения по ключу в файле
func FindSentenceByKeyInFile(key string) string {
	textStorage := make(map[string]string)
	pathToFile, _ := os.LookupEnv("INFO_TXT")

	fileDataBytes, err := os.ReadFile(pathToFile)
	if err != nil {
		log.Println(err)
		return ""
	}

	fileDataString := string(fileDataBytes)
	fileRows := strings.Split(fileDataString, "---")
	for _, row := range fileRows {
		// Данный вызов TrimSpace очень важен
		// здесь мы удаляем пробелы с концов строки,
		// чтобы далее она правильно интерпритировалась в массив
		row = strings.TrimSpace(row)
		pairKeyAndValue := strings.Split(row, "=")
		var keyOfPair, valueOfPair string = pairKeyAndValue[0], pairKeyAndValue[1]
		textStorage[keyOfPair] = valueOfPair
	}

	return textStorage[key]
}

type userManagerApi struct {
	// реализация сервера, который служит заглушкой
	// для еще не реализованных запросов
	user_manager_v1.UnimplementedUserManagerServer
	userManager UserManager
}

func main() {
	// load env variables
	MustSetUpEnv()

	log := SetupLoger("local")

	authConfig := auth_config.MustLoad()
	authApllication := auth_app.New(log, authConfig.GRPC.Port)
	authApllication.GRPCServer.MustRun()
	// TODO: все упаковать в gorutines и добавить канал,
	// который ожидает сигнала по завершению
	// затем GracefulShotDowmn
}

// Не используйте этот метод при реализации обработчика запросов на сервере,
// так как он некорректно работает с базой данных
func Add_catalog_equipment(ctx *fiber.Ctx) error {
	type item struct {
		Name string `json:"name"`
	}

	req_body := ctx.Request().Body()

	req_result := item{}

	err_json_unmarshall := json.Unmarshal(req_body, &req_result)

	if err_json_unmarshall != nil {
		log.Println(err_json_unmarshall)
	}

	result := models.Equipment_type{
		Name: req_result.Name,
	}

	DataBase.DB.Create(&result)

	return ctx.JSON("Success")

}

func (userManager *userManagerApi) GetUserProfile(
	ctx context.Context,
	request *user_manager_v1.GetUserProfileRequest,
) (*user_manager_v1.UserProfile, error) {
	// TODO: сделать валидацию данных
	user, errGetUser := userManager.userManager.GetUserProfile(ctx, request.GetUserId())
	if errGetUser != nil {
		return nil, errGetUser
	}

	return &user_manager_v1.UserProfile{
		Email:       user.Login,
		Name:        user.Name,
		Surname:     user.Surname,
		Sex:         user.Sex,
		Height:      user.Height,
		Weight:      user.Weight,
		DateBirthMs: user.BirthDateMs,
	}, nil
}

func (userManager *userManagerApi) GetMatchStatistic(
	context.Context,
	*user_manager_v1.GetUserStatisticRequest,
) (*user_manager_v1.UserStatistic, error) {
	// TODO: реализовать метод получения статистики по матчу
	return &user_manager_v1.UserStatistic{}, nil
}

func (userManager *userManagerApi) GetUserMatchesStatistic(
	context.Context,
	*user_manager_v1.GetUserStatisticRequest,
) (*user_manager_v1.UserMatchesStatistic, error) {
	// TODO: реализовать метод получения статистики по матчам
	return &user_manager_v1.UserMatchesStatistic{}, nil
}

func (userManager *userManagerApi) GetLeagueSchedule(
	context.Context,
	*user_manager_v1.GetLeagueScheduleRequest,
) (*user_manager_v1.LeagueSchedule, error) {
	// TODO: реализовать метод получения расписания лиги
	return &user_manager_v1.LeagueSchedule{}, nil
}

func (userManager *userManagerApi) GetLeagueStanding(
	context.Context,
	*user_manager_v1.GetLeagueStandingRequest,
) (*user_manager_v1.LeagueStanding, error) {
	// TODO: реализовать метод получения турнирной таблицы турнира
	return &user_manager_v1.LeagueStanding{}, nil
}
