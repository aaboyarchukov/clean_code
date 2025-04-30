package lesson2

import (
	"backend/excel"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func Get_columns(ctx *fiber.Ctx) error {
	type answerColumns struct {
		ColumnName string `json:"column_name"`
		IndRow     uint   `json:"ind_row"`
		IndColumn  int    `json:"ind_column"`
	}

	type answer struct {
		Status string          `json:"status"`
		Code   int             `json:"code"`
		Error  error           `json:"error"`
		Body   []answerColumns `json:"body"`
	}

	// 6.4
	// old name: path
	// new name: filePath
	// old name: ok
	// new name: notEmptyPathENV
	filePath, notEmptyPathENV := os.LookupEnv("PATH_KP_FILES")

	if !notEmptyPathENV {
		return ctx.JSON(
			answer{
				Status: "failed",
				Code:   fiber.ErrBadRequest.Code,
				Error:  fmt.Errorf("error with reading .env virables"),
				Body:   []answerColumns{},
			},
		)
	}

	filesSend, errMulti := ctx.MultipartForm()

	if errMulti != nil {
		return ctx.JSON(
			answer{
				Status: "failed",
				Code:   fiber.ErrBadRequest.Code,
				Error:  errMulti,
				Body:   []answerColumns{},
			},
		)
	}

	// object_id := filesSend.Value["object_id"][0]
	// equipment_type := filesSend.Value["equipment_type"][0]

	for _, file := range filesSend.File["file"] {
		destination := fmt.Sprintf("%s%s", filePath, file.Filename)
		if save_file_err := ctx.SaveFile(file, destination); save_file_err != nil {
			return ctx.JSON(
				answer{
					Status: "failed",
					Code:   fiber.ErrBadRequest.Code,
					Error:  save_file_err,
					Body:   []answerColumns{},
				},
			)
		}
	}

	// 6.4
	// old name: files
	// new name: filesFromDir
	filesFromDir, errReadDir := os.ReadDir(filePath)

	if errReadDir != nil {
		return ctx.JSON(
			answer{
				Status: "failed",
				Code:   fiber.ErrBadRequest.Code,
				Error:  errReadDir,
				Body:   []answerColumns{},
			},
		)
	}

	if len(filesFromDir) == 0 {
		return ctx.JSON(
			answer{
				Status: "failed",
				Code:   fiber.ErrBadRequest.Code,
				Error:  fmt.Errorf("there are no files in directory"),
				Body:   []answerColumns{},
			},
		)
	}

	// 6.4
	// old name: file
	// new name: firstFile
	firstFile := filesFromDir[0]

	// 6.4
	// old name: name
	// new name: nameOfFile
	nameOfFile := fmt.Sprintf("%v%v", filePath, firstFile.Name())
	currentFile, err := excelize.OpenFile(nameOfFile)
	if err != nil {
		return ctx.JSON(
			answer{
				Status: "failed",
				Code:   fiber.ErrBadRequest.Code,
				Error:  err,
				Body:   []answerColumns{},
			},
		)
	}

	defer func() {
		if err := currentFile.Close(); err != nil {
			log.Println(err)
		}
		os.Remove(nameOfFile)
	}()

	sheetName := currentFile.GetSheetList()
	if len(sheetName) == 0 {
		return ctx.JSON(
			answer{
				Status: "failed",
				Code:   fiber.ErrBadRequest.Code,
				Error:  fmt.Errorf("empty excel file"),
				Body:   []answerColumns{},
			},
		)
	}

	// 6.4
	// old name: rows
	// new name: rowsFromFile
	rowsFromFile, errGetRows := currentFile.GetRows(sheetName[0])
	if errGetRows != nil {
		return ctx.JSON(
			answer{
				Status: "failed",
				Code:   fiber.ErrBadRequest.Code,
				Error:  errGetRows,
				Body:   []answerColumns{},
			},
		)
	}

	rowForAnalyze, rowInd, errGetRow := excel.Get_row_for_analyze(rowsFromFile)
	if errGetRow != nil {
		return ctx.JSON(
			answer{
				Status: "failed",
				Code:   fiber.ErrBadRequest.Code,
				Error:  errGetRow,
				Body:   []answerColumns{},
			},
		)
	}

	// 6.4
	// old name: columns
	// new name: columnsFromAnswer
	var columnsForAnswer []answerColumns
	for columnInd, columnName := range rowForAnalyze {
		if columnName == "" || columnName == " " {
			continue
		}

		tempColumn := answerColumns{
			ColumnName: columnName,
			IndColumn:  columnInd,
			IndRow:     rowInd,
		}
		columnsForAnswer = append(columnsForAnswer, tempColumn)
	}

	return ctx.JSON(
		answer{
			Status: "success",
			Code:   200,
			Error:  nil,
			Body:   columnsFromAnswer,
		},
	)
}
