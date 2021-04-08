package mrz

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func createFieldParser(fieldName string, parser textParser, line, start, end uint) fieldParser {
	pos := &Position{
		Line:  line,
		Start: start,
		End:   end,
	}
	if parser == nil {
		parser = parseText
	}
	return func(lines []string) (*Field, error) {
		raw := getText(lines, pos)
		txt, err := parser(raw)
		return &Field{
			Name:  fieldName,
			Value: txt,
			Valid: err == nil,
			Ranges: []Range{
				{
					Raw:      raw,
					Position: pos,
				},
			},
			Position: pos,
		}, nil
	}
}

func parseText(src string) (string, error) {
	return clearText(src), nil
}

func parseDocumentCode(src string) (string, error) {
	if src[0] != 'A' && src[0] != 'C' && src[0] != 'I' {
		return src, fmt.Errorf(`invalid document code: %v. First character must be A, C or I`, src)
	}
	return parseText(src)
}

func parseDocumentNumber(src, checkDigit string) (string, error) {
	return "", nil
}

func parseFirstName(src string) (string, error) {
	arr := strings.Split(src, "<<")
	return parseText(arr[0])
}

func parseLastName(src string) (string, error) {
	arr := strings.Split(src, "<<")
	return parseText(arr[1])
}

func parseState(src string) (string, error) {
	src = clearText(src)
	state := StateMap[src]
	if state == "" {
		return src, fmt.Errorf(`invalid state code: %v`, src)
	}
	return state, nil
}

func parseDate(src string) (string, error) {
	matched, err := regexp.MatchString("^[0-9<]{4,6}$", src)
	if err != nil {
		return src, err
	}
	if !matched {
		return src, fmt.Errorf(`invalid date: %v`, src)
	}
	monthStr := src[2:4]
	month, _ := strconv.Atoi(monthStr)
	if monthStr != "<<" && (month < 1 || month > 12) {
		return src, fmt.Errorf(`invalid date month: %v`, src)
	}

	if len(src) == 6 {
		dayStr := src[4:]
		day, _ := strconv.Atoi(dayStr)
		if dayStr != "<<" && (day < 1 || day > 31) {
			return src, fmt.Errorf(`invalid date day: %v`, src)
		}
	}
	return src, nil
}

func parseSex(src string) (string, error) {
	switch src {
	case "M":
		return "male", nil
	case "F":
		return "female", nil
	case "<":
		return "notspecified", nil
	}
	return src, fmt.Errorf(`invalid sex: %v. must be 'M', 'F' or '<'.`, src)
}
