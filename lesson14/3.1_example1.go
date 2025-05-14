package lesson14

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Report struct {
	ID   int
	Name string
}

func (report *Report) Print(computerHost string) {}

func (report *Report) SendByMail(destEmail string) {}

func (report *Report) Sign() {}

type Employee struct {
	EmployeeID int
	Name       string
	Post       string
}

func (employee *Employee) CountSalary(salaries []float64) map[string]float64  {}
func (employee *Employee) FormingReport(salaryInfo map[string]float64) Report {}

// формирование отчета о зарплате за определенный период
func (employee *Employee) PrepareSalaryReportByPeriod(period time.Time) Report {}
func (employee *Employee) ConvertToExcel(document Report) file.Excel           {}

type OS struct {
	OsID    int
	Name    string
	Version string
}

func (os *OS) GetVersion() {}

func (os *OS) GetOSParametrs() {}

func (os *OS) GetActiveThreads() {}

func (os *OS) ExecuteProgram(programID int) {}

func (os *OS) Start() {}

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

// отношение рабочей части
// (процесс, когда тендер находится в обработке) тендера
type WorkPartOfTender struct {
	WorkPartID              uint       `gorm:"primaryKey" json:"work_part_id"`
	TenderRefer             string     `gorm:"unique" json:"object_refer"`
	TenderID                Tender     `gorm:"foreignKey:TenderRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"object_id"`
	StuffReferVendor        *uint      `json:"stuff_refer_vendor"`
	StuffIDVendor           Stuff      `gorm:"foreignKey:StuffReferVendor;constraint:OnUpdate:CASCADE;" json:"stuff_id_vendor"`
	StageReferVendor        *uint      `json:"stage_refer_vendor"`
	StageVendorID           Stages     `gorm:"foreignKey:StageReferVendor;constraint:OnUpdate:CASCADE;" json:"stage_vendor_id"`
	StuffReferSeller        *uint      `json:"stuff_refer_seller"`
	StuffIDSeller           Stuff      `gorm:"foreignKey:StuffReferSeller;constraint:OnUpdate:CASCADE;" json:"stuff_id_seller"`
	StageReferSeller        *uint      `json:"stage_refer_seller"`
	StageSellerID           Stages     `gorm:"foreignKey:StageReferSeller;constraint:OnUpdate:CASCADE;" json:"stage_seller_id"`
	Status                  *string    `json:"status"`
	Deadline1               *time.Time `json:"deadline1"`
	Deadline2               *time.Time `json:"deadline2"`
	Deadline3               *time.Time `json:"deadline3"`
	InitiatorReferID        *string    `json:"initiator_refer_id"`
	IintiatorRefer          Companies  `gorm:"foreignKey:InitiatorReferID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"initiator_refer"`
	OrganizerReferID        *string    `json:"organizer_refer_id"`
	OrganizerRefer          Companies  `gorm:"foreignKey:OrganizerReferID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"organizer_refer"`
	InitiatorContactReferID *uint      `json:"initiator_contact_refer_id"`
	InitiatorContactRefer   Contact    `gorm:"foreignKey:InitiatorContactReferID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"initiator_contact_refer"`
	WinnerReferID           *string    `json:"winner_refer_id"`
	WinnerRefer             Companies  `gorm:"foreignKey:WinnerReferID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"winner_refer"`
}

// формирование коммерческого предложения по торгу
func GetCommericalProposal(object_id string, columns data_analyze.ColumnsValues, ndsStatus string) (Answer_from_kp, error) {
	work_part_id, err_get_work_part_id := database.Get_work_path_id(object_id)
	if err_get_work_part_id != nil {
		return Answer_from_kp{}, err_get_work_part_id
	}

	indexColumns := []data_analyze.Point{
		columns.EquipmentName,
		columns.EquipmentUnits,
		columns.EquipmentCount,
		columns.EquipmentDelivery,
		columns.EquipmentSpecification,
		columns.EquipmentArticleNumber,
		columns.EquipmentDeadLine,
		columns.EquipmentPaymentDate,
		columns.EquipmentPrice,
		columns.EquipmentCost,
	}

	path, ok := os.LookupEnv("PATH_KP_FILES")

	if !ok {
		return Answer_from_kp{}, fmt.Errorf("error with reading .env virables")
	}

	files, err_read_dir := os.ReadDir(path)

	if err_read_dir != nil {
		log.Println(err_read_dir)
		return Answer_from_kp{}, err_read_dir
	}

	if len(files) == 0 {
		return Answer_from_kp{}, fmt.Errorf("there are no files in directory")
	}

	file := files[0]
	name := fmt.Sprintf("%v%v", path, file.Name())
	current_file, err := excelize.OpenFile(name)

	defer func() {
		if err := current_file.Close(); err != nil {
			log.Println(err)
		}
		os.Remove(name)
	}()

	if err != nil {
		return Answer_from_kp{}, err
	}

	sheetName := current_file.GetSheetList()
	if len(sheetName) == 0 {
		return Answer_from_kp{}, fmt.Errorf("empty excel file")
	}

	rows, err_get_rows := current_file.GetRows(sheetName[0])
	if err_get_rows != nil {
		return Answer_from_kp{}, err_get_rows
	}

	var result_company models.Companies

	for ind := 0; ind < len(rows); ind++ {
		if len(rows[ind]) == 0 {
			continue
		}

		needed_ind := 0
		for rows[ind][needed_ind] == "" || rows[ind][needed_ind] == " " {
			needed_ind++
		}

		if len(rows[ind]) != 0 && data_analyze.Find_row(rows[ind][needed_ind]) != "" {
			info, err_get_info := data_analyze.Get_info_KP(rows[ind][needed_ind])
			if err_get_info != nil {
				return Answer_from_kp{}, err_get_info
			}
			result_company = info.CompanyInfo
			break
		}

	}

	var begin_rows int
	for _, ind := range indexColumns {
		if ind.Row != -1 {
			begin_rows = ind.Row + 1
			break
		}
	}

	equipments := []models.Equipment{}

	for ind := begin_rows; ind < len(rows); ind++ {
		var tempEquipment models.Equipment

		if !Is_correct_row(rows[ind], indexColumns, rows[begin_rows-1]) {
			continue
		}

		var count float64
		if columns.EquipmentCount.Column == -1 {
			count = 1
		} else {
			prepare_count, err_prepare_count := data_analyze.Trim_symbols(rows[ind][columns.EquipmentCount.Column])
			prepare_count = edit_numbers.FormatStrNumber(prepare_count)
			if err_prepare_count != nil {
				return Answer_from_kp{}, err_prepare_count
			}

			if prepare_count == "" || prepare_count == " " {
				continue
			}
			tempCount, err_convert_count := strconv.ParseFloat(prepare_count, 64)
			if err_convert_count != nil {
				return Answer_from_kp{}, err_convert_count
			}

			count = tempCount
		}

		var cost float64

		if columns.EquipmentPrice.Column == -1 && columns.EquipmentCost.Column != -1 {
			prepare_cost, err_prepare_cost := data_analyze.Trim_symbols(rows[ind][columns.EquipmentCost.Column])
			prepare_cost = edit_numbers.FormatStrNumber(prepare_cost)
			if err_prepare_cost != nil {
				return Answer_from_kp{}, err_prepare_cost
			}

			tempCost, err_convert_cost := strconv.ParseFloat(prepare_cost, 64)
			if err_convert_cost != nil {
				return Answer_from_kp{}, err_convert_cost
			}

			switch ndsStatus {
			case "with_nds":
				cost = edit_numbers.Round(tempCost/1.2, 2)
			}
		} else {
			prepare_price, err_prepare_price := data_analyze.Trim_symbols(rows[ind][columns.EquipmentPrice.Column])
			if err_prepare_price != nil {
				return Answer_from_kp{}, err_prepare_price
			}

			prepare_price = edit_numbers.FormatStrNumber(prepare_price)
			tempPrice, err_temp_price := strconv.ParseFloat(prepare_price, 64)
			if err_temp_price != nil {
				return Answer_from_kp{}, err_temp_price
			}

			switch ndsStatus {
			case "with_nds":
				cost = edit_numbers.Round((tempPrice/1.2)*count, 2)
			case "without_nds":
				cost = edit_numbers.Round((tempPrice)*count, 2)
			}
		}

		tempEquipment.Count = count
		tempEquipment.NewCount = count
		tempEquipment.Price_per_unit_now = edit_numbers.Round((cost / (count)), 2)
		tempEquipment.New_Price_per_unit_now = edit_numbers.Round((cost / (count)), 2)
		tempEquipment.Cost = cost
		tempEquipment.NewCost = cost
		tempEquipment.NDS = 1.2
		tempEquipment.EquipmentKind = "vendor"

		if columns.EquipmentName.Column != -1 {
			tempEquipment.Resource_name = rows[ind][columns.EquipmentName.Column]
		}
		if columns.EquipmentSpecification.Column != -1 {
			tempEquipment.Specifications = rows[ind][columns.EquipmentSpecification.Column]
		}
		if columns.EquipmentArticleNumber.Column != -1 {
			tempEquipment.ArticleNumber = rows[ind][columns.EquipmentArticleNumber.Column]
		}
		if columns.EquipmentUnits.Column != -1 {
			tempEquipment.Units = rows[ind][columns.EquipmentUnits.Column]
		}
		if columns.EquipmentDelivery.Column != -1 {
			tempEquipment.Delivery = rows[ind][columns.EquipmentDelivery.Column]
		}
		if columns.EquipmentDeadLine.Column != -1 {
			tempEquipment.Deadline = rows[ind][columns.EquipmentDeadLine.Column]
		}
		if columns.EquipmentPaymentDate.Column != -1 {
			tempEquipment.PaymentDay = rows[ind][columns.EquipmentPaymentDate.Column]
		}

		equipments = append(equipments, tempEquipment)

	}

	return Answer_from_kp{
		CompanyInfo:        result_company,
		WorkPartID:         work_part_id,
		CompanyID:          result_company.CompanyID,
		Equipments_from_kp: equipments,
	}, nil

}
