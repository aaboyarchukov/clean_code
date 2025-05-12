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

	// old: strColumns
	// new var columnsFromFile string
	var columnsFromFile string = files.Value["columns"][INDX_OF_VALUE_IN_MAP]
	var columns []answerColumns
	err_unmurshull := json.Unmarshal([]byte(columnsFromFile), &columns)
	if err_unmurshull != nil {
		return ctx.JSON(answerResponse{
			Ok:         false,
			StatusCode: 500,
			Errors:     err_unmurshull,
			Data:       Answer_from_kp{},
		})
	}

	// old: resultColumns
	// new: var resultColumns data_analyze.ColumnsValues
	var resultColumns data_analyze.ColumnsValues = data_analyze.ColumnsValues{
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
		// old: destination
		// new: var fileDestination string
		var fileDestination string = fmt.Sprintf("%s%s", path, file.Filename)
		if save_file_err := ctx.SaveFile(file, fileDestination); save_file_err != nil {
			log.Println("ошибка сохранения файла")
			log.Println(save_file_err)
			return save_file_err
		}
	}

	tenders, err_get_kp := Get_KP(tenderId, resultColumns, ndsStatus)

	if err_get_kp != nil {
		fmt.Println("err_get_kp: ", err_get_kp)
		return ctx.JSON(answerResponse{
			Ok:         false,
			StatusCode: 500,
			Errors:     err_get_kp,
			Data:       Answer_from_kp{},
		})
	}

	add_vendor_err := database.Add_vendors(tenders.WorkPartID, tenders.CompanyInfo, tenderId, strings.ToLower(tenderEquipmentType), tenders.Equipments_from_kp)
	if add_vendor_err != nil {
		fmt.Println("add_vendor_err: ", add_vendor_err)
		return ctx.JSON(answerResponse{
			Ok:         false,
			StatusCode: 500,
			Errors:     add_vendor_err,
			Data:       Answer_from_kp{},
		})
	}

	vendorEquipments, err_get_vendor_equipments := database.Get_equipments_for_company(strings.ToLower(tenderEquipmentType), tenderId, tenders.CompanyInfo.CompanyID)
	if err_get_vendor_equipments != nil {
		fmt.Println("err_get_vendor_equipments: ", err_get_vendor_equipments)
		return ctx.JSON(answerResponse{
			Ok:         false,
			StatusCode: 500,
			Errors:     err_get_vendor_equipments,
			Data:       Answer_from_kp{},
		})
	}

	tenders.Equipments_from_kp = vendorEquipments

	return ctx.JSON(answerResponse{
		Ok:         true,
		StatusCode: 200,
		Errors:     nil,
		Data:       tenders,
	})
}

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {

	var valueStr1, valueStr2 string
	// old: flagStr
	// new: var flagStr bool
	var flagStr bool = false

	if value, ok := any(v1).(string); ok {
		valueStr1 = strings.Trim(value, " ")
		flagStr = true
	}

	if value, ok := any(v2).(string); ok {
		valueStr2 = strings.Trim(value, " ")
		flagStr = true
	}

	switch flagStr {
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

func (l *OrderedList[T]) Delete(n T) {
	if l.head == nil {
		return
	}

	if l.Count() == 1 {
		l.Clear(l._ascending)
		return
	}

	// old: currentNode
	// new: var currentNode *Node[T]
	var currentNode *Node[T] = l.head
	var deleted bool = false
	for currentNode != nil {
		compareNode := l.Compare(currentNode.value, n)
		if compareNode == 0 && currentNode == l.head {
			l.head = currentNode.next
			if l.head != nil {
				l.head.prev = nil
			}
			deleted = true
		} else if compareNode == 0 && currentNode == l.tail {
			l.tail = currentNode.prev
			if l.tail != nil {
				l.tail.next = nil
			}
			deleted = true
		} else if compareNode == 0 {
			prevNode := currentNode.prev
			nextNode := currentNode.next
			prevNode.next = nextNode
			nextNode.prev = prevNode
			deleted = true
		}

		if deleted {
			break
		}
		currentNode = currentNode.next
	}
}

func (l *OrderedList[T]) Add(item T) {
	// old: node
	// new: var addNode Node[T]
	var addNode Node[T] = Node[T]{
		value: item,
	}

	size := l.Count()

	if size == 0 {
		l.head = &addNode
		l.tail = &addNode
		return
	}

	compareHead, compareTail := l.Compare(addNode.value, l.head.value), l.Compare(addNode.value, l.tail.value)

	if (l._ascending && (compareTail == 1 || compareTail == 0)) ||
		(!l._ascending && (compareTail == -1 || compareTail == 0)) {
		l.tail.next = &addNode
		addNode.prev = l.tail
		l.tail = &addNode
		return
	}
	if (l._ascending && (compareHead == -1 || compareHead == 0)) ||
		(!l._ascending && (compareHead == 1 || compareHead == 0)) {
		addNode.next = l.head
		l.head.prev = &addNode
		l.head = &addNode
		return
	}

	left, right := l.head, l.head.next
	for right != nil {
		compareNodeAndLeft := l.Compare(addNode.value, left.value)
		compareNodeAndRight := l.Compare(addNode.value, right.value)

		// old: asc
		// new: var asc bool
		var asc bool = (compareNodeAndLeft == 1 || compareNodeAndLeft == 0) &&
			(compareNodeAndRight == -1 || compareNodeAndRight == 0)

		// old: desc
		// new: var desc bool
		var desc bool = (compareNodeAndLeft == -1 || compareNodeAndLeft == 0) &&
			(compareNodeAndRight == 1 || compareNodeAndRight == 0)

		if (l._ascending && asc) || (!l._ascending && desc) {
			addNode.prev = left
			addNode.next = right
			left.next = &addNode
			right.prev = &addNode
			return
		}

		left = left.next
		right = right.next
	}
}
