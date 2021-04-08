package mrz

var td1Tpl = []fieldParser{
	createFieldParser("documentCode", parseDocumentCode, 0, 0, 2),
	createFieldParser("issuingState", parseState, 0, 2, 5),
	createFieldParser("documentNumber", nil, 0, 5, 14),
	createFieldParser("documentNumberCheckDigit", nil, 0, 14, 15),
	createFieldParser("birthDate", parseDate, 1, 0, 6),
	createFieldParser("birthDateCheckDigit", nil, 1, 6, 7),
	createFieldParser("sex", parseSex, 1, 7, 8),
	createFieldParser("expirationDate", parseDate, 1, 8, 14),
	createFieldParser("expirationDateCheckDigit", nil, 1, 14, 15),
	createFieldParser("nationality", nil, 1, 15, 18),
	createFieldParser("lastName", parseLastName, 2, 0, 30),
	createFieldParser("firstName", parseFirstName, 2, 0, 30),
	createFieldParser("optional1", nil, 0, 15, 30),
	createFieldParser("optional2", nil, 1, 18, 29),
}

var td2Tpl = []fieldParser{}

var td3Tpl = []fieldParser{}
