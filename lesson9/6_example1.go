package lesson9

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const LINE_SEPARATOR = "---"
const KEY_VALUE_SEPARATOR = "="
const (
	INDX_KEY = iota
	INDX_VALUE
)

func Contain(value string) string {
	textMap := make(map[string]string)
	filePath, _ := os.LookupEnv("INFO_TXT")

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return ""
	}

	stringData := string(fileData)
	arrRows := strings.Split(stringData, LINE_SEPARATOR)
	for _, item := range arrRows {
		item = strings.TrimSpace(item)
		pairItems := strings.Split(item, KEY_VALUE_SEPARATOR)
		pairKey := pairItems[INDX_KEY]
		pairValue := pairItems[INDX_VALUE]
		textMap[pairKey] = pairValue
	}

	return textMap[value]
}

func (l *LinkedList) Delete(n int, all bool) {
	if l.head == nil {
		return
	}

	tempNode := l.head
	var prev *Node
	var equalValue bool = tempNode.value == n
	var onlyOneNodeInLinkedList bool = l.Count() == 1

	// old: tempNode.value == n && l.Count() == 1
	if onlyOneNodeInLinkedList && equalValue {
		l.Clean()
		return
	}

	for tempNode != nil {
		var deleted bool = false
		// old: tempNode.value == n && tempNode == l.head
		// tempNode.value == n && tempNode == l.tail
		// tempNode.value == n
		equalValue = tempNode.value == n
		var equalHead bool = tempNode == l.head
		var equalTail bool = tempNode == l.tail

		if equalValue && equalHead {
			l.head = tempNode.next
			deleted = true
		} else if equalValue && equalTail {
			prev.next = nil
			l.tail = prev
			deleted = true
		} else if equalValue {
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

// old: '.'
const FLOAT_NUMBER_SEPARATOR rune = '.'

func FormatStrNumber(strNumber string) string {
	countDot := 0
	for _, symbol := range strNumber {
		// old: string(symbol) == "."
		// symbol == FLOAT_NUMBER_SEPARATOR
		if symbol == FLOAT_NUMBER_SEPARATOR {
			countDot++
		}
	}

	return strings.Replace(strNumber, ".", "", countDot-1)
}

const ERROR_RATE float64 = 0.0001

func EqualFloat(firstNumber, secondNumber float64) bool {
	// old: firstNumber == secondNumber
	difference := firstNumber - secondNumber
	return difference < ERROR_RATE
}

const RETURN_WRONG_FLOAT float64 = -1
const ERROR_TEXT string = "devide by zero"

func SinWithTwoTriangleSides(firstSide, secondSide int) (float64, error) {
	if secondSide == 0 {
		return RETURN_WRONG_FLOAT, fmt.Errorf(ERROR_TEXT)
	}

	floatFirstSide := float64(firstSide)
	floatSecondSide := float64(secondSide)
	sin := floatFirstSide / floatSecondSide

	return sin, nil

}
