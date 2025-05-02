package lesson3

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func GetEvents(ctx *fiber.Ctx) error {
	type answer struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Start       string `json:"start"`
		End         string `json:"end"`
		Description string `json:"description"`
	}

	type request struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}

	var req_result request

	if err_json_unmarshall := json.Unmarshal(ctx.Request().Body(), &req_result); err_json_unmarshall != nil {
		return ctx.JSON(err_json_unmarshall)
	}

	// 7.4
	start, err_start := time.Parse("2006-01-02 15:04:05-07", req_result.Start)
	if err_start != nil {
		return ctx.JSON(err_start)
	}
	// 7.4
	end, err_end := time.Parse("2006-01-02 15:04:05-07", req_result.End)
	if err_end != nil {
		return ctx.JSON(err_end)
	}

	var resp_result []answer
	events := database.GetEvents(start, end)
	for idx, event := range events {
		winner_id, err_get_winner_id := database.Get_object_winner(event.ObjectRefer)
		winner, err_get_winner := database.Get_company(winner_id)
		if err_get_winner != nil || err_get_winner_id != nil {
			return ctx.JSON(err_get_winner.Error() + err_get_winner_id.Error())
		}

		description := fmt.Sprintf("Название: %s Победитель: %s", event.FullName, winner.Short_name)
		resultDate := event.Data.Format("2006-01-02")
		resp_result = append(resp_result, answer{
			Id:          idx + 1,
			Title:       event.ObjectRefer,
			Start:       resultDate,
			End:         resultDate,
			Description: description,
		})
	}

	return ctx.JSON(resp_result)
}

func Get_KP(object_id string, columns data_analyze.ColumnsValues, ndsStatus string) (Answer_from_kp, error) {
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

	// 7.4
	var beginRows int
	for _, ind := range indexColumns {
		if ind.Row != -1 {
			begin_rows = ind.Row + 1
			break
		}
	}

	// 7.4
	var endRows int = len(rows)

	equipments := []models.Equipment{}

	for ind := beginRows; ind < endRows; ind++ {
		var tempEquipment models.Equipment

		if !Is_correct_row(rows[ind], indexColumns, rows[beginRows-1]) {
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
