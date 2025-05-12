package lesson10

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func GenerateJWT(user models.Stuff, duration time.Duration) (string, error) {
	// old: token
	// new var token *jwt.Token
	var token *jwt.Token = jwt.New(jwt.SigningMethodHS256)

	// old: claims
	// new: var claims jwt.MapClaims
	var claims jwt.MapClaims = token.Claims.(jwt.MapClaims)

	claims["id"] = user.EmployeeID
	claims["role"] = user.Rights
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["iat"] = time.Now().Unix()

	secret, _ := os.LookupEnv("SECRET")
	signedToken, err_signed := token.SignedString([]byte(secret))
	if err_signed != nil {
		return signedToken, err_signed
	}

	return signedToken, nil

}

func Round(number float64, base uint) float64 {
	// old: diff
	// new: var diff float64
	var diff float64 = math.Pow(float64(10), float64(base))
	return math.Round(number*diff) / diff
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
	// old: 0
	// new: INDX_OF_VALUE_IN_MAP

	// old: object_id
	// new var tenderId string
	var tenderId string = files.Value["object_id"][INDX_OF_VALUE_IN_MAP]

	// old: equipment_type
	// new var tenderEquipmentType string
	var tenderEquipmentType string = files.Value["equipment_type"][INDX_OF_VALUE_IN_MAP]

	// old: nds
	// new: var ndsStatus string
	var ndsStatus string = files.Value["nds"][INDX_OF_VALUE_IN_MAP]

	type answerColumns struct {
		ColumnName string `json:"column_name"`
		IndRow     int    `json:"ind_row"`
		IndColumn  int    `json:"ind_column"`
	}
	strColumns := files.Value["columns"][INDX_OF_VALUE_IN_MAP]
	var columns []answerColumns
	err_unmurshull := json.Unmarshal([]byte(strColumns), &columns)
	if err_unmurshull != nil {
		return ctx.JSON(answerResponse{
			Ok:         false,
			StatusCode: 500,
			Errors:     err_unmurshull,
			Data:       Answer_from_kp{},
		})
	}

	resultColumns := data_analyze.ColumnsValues{
		EquipmentName: data_analyze.Point{
			Row:    columns[CELL_FOR_NAME].IndRow,
			Column: (columns[CELL_FOR_NAME].IndColumn),
		},
		EquipmentUnits: data_analyze.Point{
			Row:    columns[CELL_FOR_UNITS].IndRow,
			Column: (columns[CELL_FOR_UNITS].IndColumn),
		},
		EquipmentCount: data_analyze.Point{
			Row:    columns[CELL_FOR_COUNT].IndRow,
			Column: (columns[CELL_FOR_COUNT].IndColumn),
		},
		EquipmentDelivery: data_analyze.Point{
			Row:    columns[CELL_FOR_DELIVERY].IndRow,
			Column: (columns[CELL_FOR_DELIVERY].IndColumn),
		},
		EquipmentSpecification: data_analyze.Point{
			Row:    columns[CELL_FOR_SPECIFICATION].IndRow,
			Column: (columns[CELL_FOR_SPECIFICATION].IndColumn),
		},
		EquipmentArticleNumber: data_analyze.Point{
			Row:    columns[CELL_FOR_ARTICLE_NUMBER].IndRow,
			Column: (columns[CELL_FOR_ARTICLE_NUMBER].IndColumn),
		},
		EquipmentDeadLine: data_analyze.Point{
			Row:    columns[CELL_FOR_DEADLINE].IndRow,
			Column: (columns[CELL_FOR_DEADLINE].IndColumn),
		},
		EquipmentPaymentDate: data_analyze.Point{
			Row:    columns[CELL_FOR_PAYMENT_DATE].IndRow,
			Column: (columns[CELL_FOR_PAYMENT_DATE].IndColumn),
		},
		EquipmentPrice: data_analyze.Point{
			Row:    columns[CELL_FOR_PRICE].IndRow,
			Column: (columns[CELL_FOR_PRICE].IndColumn),
		},
		EquipmentCost: data_analyze.Point{
			Row:    columns[CELL_FOR_COST].IndRow,
			Column: (columns[CELL_FOR_COST].IndColumn),
		},
	}

	for _, file := range files.File["file"] {
		destination := fmt.Sprintf("%s%s", path, file.Filename)
		if save_file_err := ctx.SaveFile(file, destination); save_file_err != nil {
			log.Println("ошибка сохранения файла")
			log.Println(save_file_err)
			return save_file_err
		}
	}

	objects, err_get_kp := Get_KP(tenderId, resultColumns, ndsStatus)

	if err_get_kp != nil {
		fmt.Println("err_get_kp: ", err_get_kp)
		return ctx.JSON(answerResponse{
			Ok:         false,
			StatusCode: 500,
			Errors:     err_get_kp,
			Data:       Answer_from_kp{},
		})
	}

	add_vendor_err := database.Add_vendors(objects.WorkPartID, objects.CompanyInfo, tenderId, strings.ToLower(tenderEquipmentType), objects.Equipments_from_kp)
	if add_vendor_err != nil {
		fmt.Println("add_vendor_err: ", add_vendor_err)
		return ctx.JSON(answerResponse{
			Ok:         false,
			StatusCode: 500,
			Errors:     add_vendor_err,
			Data:       Answer_from_kp{},
		})
	}

	vendorEquipments, err_get_vendor_equipments := database.Get_equipments_for_company(strings.ToLower(tenderEquipmentType), tenderId, objects.CompanyInfo.CompanyID)
	if err_get_vendor_equipments != nil {
		fmt.Println("err_get_vendor_equipments: ", err_get_vendor_equipments)
		return ctx.JSON(answerResponse{
			Ok:         false,
			StatusCode: 500,
			Errors:     err_get_vendor_equipments,
			Data:       Answer_from_kp{},
		})
	}

	objects.Equipments_from_kp = vendorEquipments

	return ctx.JSON(answerResponse{
		Ok:         true,
		StatusCode: 200,
		Errors:     nil,
		Data:       objects,
	})
}
