package mrz

import (
	"errors"
	"regexp"
	"strings"
)

func Parse(data string) (*Document, error) {
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		matched, err := regexp.MatchString("[A-Z0-9<]+", line)
		if err != nil {
			return nil, err
		}
		if !matched {
			return nil, errors.New(`lines must be composed of only alphanumerical characters and "<"`)
		}
	}

	switch len(lines) {
	case 2, 3:
		switch len(lines[0]) {
		case 30:
			return parseTD1(lines)
		case 36:
			matched, err := regexp.MatchString("[0-9]+", lines[0][30:])
			if err != nil {
				return nil, err
			}
			if matched {
				return parseFrenchID(lines)
			}
			return parseTD2(lines)
		case 44:
			return parseTD3(lines)
		case 9:
			return parseSwissDriveLicense(lines)
		default:
			return nil, errors.New("unrecognized document format")
		}
	}
	return nil, errors.New("invalid line length")
}

func parseTD1(lines []string) (*Document, error) {
	fields, summary, valid, err := parseFields(lines, td1Tpl)
	if err != nil {
		return nil, err
	}
	return &Document{
		Format:  "TD1",
		Fields:  fields,
		Summary: summary,
		Valid:   valid,
	}, nil
}

func parseTD2(lines []string) (*Document, error) {
	fields, summary, valid, err := parseFields(lines, td2Tpl)
	if err != nil {
		return nil, err
	}
	return &Document{
		Format:  "TD2",
		Fields:  fields,
		Summary: summary,
		Valid:   valid,
	}, nil
}

func parseTD3(lines []string) (*Document, error) {
	fields, summary, valid, err := parseFields(lines, td3Tpl)
	if err != nil {
		return nil, err
	}
	return &Document{
		Format:  "TD3",
		Fields:  fields,
		Summary: summary,
		Valid:   valid,
	}, nil
}

func parseSwissDriveLicense(lines []string) (*Document, error) {
	return nil, nil
}

func parseFrenchID(lines []string) (*Document, error) {
	return nil, nil
}

func parseFields(lines []string, tpl []fieldParser) ([]*Field, map[string]interface{}, bool, error) {
	var (
		fields  []*Field
		summary = map[string]interface{}{}
	)
	valid := true
	for _, parse := range tpl {
		f, err := parse(lines)
		if err != nil {
			return nil, nil, false, err
		}
		fields = append(fields, f)
		valid = valid && f.Valid
		if valid {
			summary[f.Name] = f.Value
		}
	}
	return fields, summary, valid, nil
}
