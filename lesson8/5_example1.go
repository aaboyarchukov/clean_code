package lesson8

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type BloomFilter struct {
	filter_len int
	filter     int64
}

// old name: CONST_17
// new name: HASH_1_KOEFF
const HASH_1_KOEFF int = 17

func (bf *BloomFilter) Hash1(s string) int {
	sum := 0
	for _, char := range s {
		code := int(char)
		sum += code * HASH_1_KOEFF
	}
	sum %= bf.filter_len
	return sum
}

// old name: CONST_223
// new name: HASH_2_KOEFF
const HASH_2_KOEFF int = 223

func (bf *BloomFilter) Hash2(s string) int {
	sum := 0
	for _, char := range s {
		code := int(char)
		sum += code * HASH_2_KOEFF
	}
	sum %= bf.filter_len
	return sum
}

func (bf *BloomFilter) Add(s string) {
	hash1 := bf.Hash1(s)
	hash2 := bf.Hash2(s)
	bf.filter |= 1 << hash1
	bf.filter |= 1 << hash2
}

func (bf *BloomFilter) IsValue(s string) bool {
	var mask int64
	mask |= 1 << bf.Hash1(s)
	mask |= 1 << bf.Hash2(s)

	if mask == bf.filter&mask {
		return true
	}

	return false
}

// old name: NUMBER
// new name: HASH_KOEFF
const HASH_KOEFF byte = 75

type NativeCache[T any] struct {
	size      int
	slots     []string
	fillSlots []bool
	hits      []int
	step      int
	values    []T
	cap       int
}

func Init[T any](sz int) NativeCache[T] {
	nc := NativeCache[T]{size: sz}
	nc.slots = make([]string, sz)
	nc.hits = make([]int, sz)
	nc.values = make([]T, sz)
	nc.fillSlots = make([]bool, sz)
	nc.step = 1
	return nc
}

func (nc *NativeCache[T]) HashFun(value string) int {
	if nc.size == 0 {
		return -1
	}

	var incx int
	var sum byte
	for i, item := range value {
		sum += byte(item) * (byte(i) + HASH_KOEFF)
	}

	incx = int(sum) % nc.size

	return incx
}

func (nc *NativeCache[T]) SeekSlot(value string) int {
	if nc.size == 0 {
		return -1
	}

	hash := nc.HashFun(value)

	if !nc.fillSlots[hash] {
		return hash
	}

	if nc.cap < nc.size {

		resultIndx, indx := hash, hash
		for nc.fillSlots[resultIndx] {
			indx += nc.step
			resultIndx = indx % nc.size
		}

		return resultIndx
	}

	if nc.cap >= nc.size {
		return nc.FindIndex(value)
	}

	return -1
}

func (nc *NativeCache[T]) MinHits() int {
	var resultIndx int
	min := math.MaxInt
	for indx := range nc.hits {
		tempElement := nc.hits[indx]
		if tempElement < min {
			min = tempElement
			resultIndx = indx
		}
	}

	return resultIndx
}

func (nc *NativeCache[T]) IsKey(key string) bool {
	if nc.size == 0 {
		return false
	}

	hash := nc.SeekSlot(key)

	return nc.fillSlots[hash] && nc.slots[hash] == key
}

func (nc *NativeCache[T]) Get(key string) (T, error) {

	var result T

	if !nc.IsKey(key) {
		return result, fmt.Errorf("key is not in array")
	}

	hash := nc.SeekSlot(key)

	if hash == -1 {
		return result, fmt.Errorf("size is zero")
	}

	result = nc.values[hash]
	nc.hits[hash] += 1

	return result, nil
}

func (nc *NativeCache[T]) Put(key string, value T) {
	if nc.cap >= nc.size {
		removedIndx := nc.MinHits()
		nc.fillSlots[removedIndx] = false
		nc.hits[removedIndx] = 0
		nc.slots[removedIndx] = ""
		nc.cap -= 1
	}

	hash := nc.SeekSlot(key)

	if !nc.fillSlots[hash] {
		nc.slots[hash] = key
		nc.values[hash] = value
		nc.fillSlots[hash] = true
		nc.cap += 1
	}

}

func (nc *NativeCache[T]) FindIndex(key string) int {
	for indx := range nc.slots {
		if nc.slots[indx] == key {
			return indx
		}
	}

	return -1
}

const MIN_EXCEL_ROW_LEN = 2

func Get_row_for_analyze(rows [][]string) ([]string, uint, error) {
	if len(rows) == 0 {
		return []string{}, 0, fmt.Errorf("empty rows")
	}

	result_ind := 0

	result := []string{}
	for ind := range rows {
		// old: 2
		// new: MIN_EXCEL_ROW_LEN
		if len(rows[ind]) > MIN_EXCEL_ROW_LEN {
			result = rows[ind]
			result_ind = ind
			break
		}
	}

	return result, uint(result_ind), nil
}

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
	object_id := files.Value["object_id"][INDX_OF_VALUE_IN_MAP]
	equipment_type := files.Value["equipment_type"][INDX_OF_VALUE_IN_MAP]
	nds := files.Value["nds"][INDX_OF_VALUE_IN_MAP]

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
			// old: 0
			// new: CELL_FOR_NAME
			Row:    columns[CELL_FOR_NAME].IndRow,
			Column: (columns[CELL_FOR_NAME].IndColumn),
		},
		EquipmentUnits: data_analyze.Point{
			// old: 1
			// new: CELL_FOR_UNITS
			Row:    columns[CELL_FOR_UNITS].IndRow,
			Column: (columns[CELL_FOR_UNITS].IndColumn),
		},
		EquipmentCount: data_analyze.Point{
			// old: 2
			// new: CELL_FOR_COUNT
			Row:    columns[CELL_FOR_COUNT].IndRow,
			Column: (columns[CELL_FOR_COUNT].IndColumn),
		},
		EquipmentDelivery: data_analyze.Point{
			// old: 3
			// new: CELL_FOR_DELIVERY
			Row:    columns[CELL_FOR_DELIVERY].IndRow,
			Column: (columns[CELL_FOR_DELIVERY].IndColumn),
		},
		EquipmentSpecification: data_analyze.Point{
			// old: 4
			// new: CELL_FOR_SPECIFICATION
			Row:    columns[CELL_FOR_SPECIFICATION].IndRow,
			Column: (columns[CELL_FOR_SPECIFICATION].IndColumn),
		},
		EquipmentArticleNumber: data_analyze.Point{
			// old: 5
			// new: CELL_FOR_ARTICLE_NUMBER
			Row:    columns[CELL_FOR_ARTICLE_NUMBER].IndRow,
			Column: (columns[CELL_FOR_ARTICLE_NUMBER].IndColumn),
		},
		EquipmentDeadLine: data_analyze.Point{
			// old: 6
			// new: CELL_FOR_DEADLINE
			Row:    columns[CELL_FOR_DEADLINE].IndRow,
			Column: (columns[CELL_FOR_DEADLINE].IndColumn),
		},
		EquipmentPaymentDate: data_analyze.Point{
			// old: 7
			// new: CELL_FOR_PAYMENT_DATE
			Row:    columns[CELL_FOR_PAYMENT_DATE].IndRow,
			Column: (columns[CELL_FOR_PAYMENT_DATE].IndColumn),
		},
		EquipmentPrice: data_analyze.Point{
			// old: 8
			// new: CELL_FOR_PRICE
			Row:    columns[CELL_FOR_PRICE].IndRow,
			Column: (columns[CELL_FOR_PRICE].IndColumn),
		},
		EquipmentCost: data_analyze.Point{
			// old: 9
			// new: CELL_FOR_COST
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

	objects, err_get_kp := Get_KP(object_id, resultColumns, nds)

	if err_get_kp != nil {
		fmt.Println("err_get_kp: ", err_get_kp)
		return ctx.JSON(answerResponse{
			Ok:         false,
			StatusCode: 500,
			Errors:     err_get_kp,
			Data:       Answer_from_kp{},
		})
	}

	add_vendor_err := database.Add_vendors(objects.WorkPartID, objects.CompanyInfo, object_id, strings.ToLower(equipment_type), objects.Equipments_from_kp)
	if add_vendor_err != nil {
		fmt.Println("add_vendor_err: ", add_vendor_err)
		return ctx.JSON(answerResponse{
			Ok:         false,
			StatusCode: 500,
			Errors:     add_vendor_err,
			Data:       Answer_from_kp{},
		})
	}

	vendorEquipments, err_get_vendor_equipments := database.Get_equipments_for_company(strings.ToLower(equipment_type), object_id, objects.CompanyInfo.CompanyID)
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
