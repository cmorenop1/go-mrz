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

var td3Tpl = []fieldParser{
	createFieldParser("documentCode", parseDocumentCodePassport, 0, 0, 2),
	createFieldParser("issuingState", parseState, 0, 2, 5),
	createFieldParser("lastName", parseLastName, 0, 5, 44),
	createFieldParser("firstName", parseFirstName, 0, 5, 44),
	createFieldParser("documentNumber", nil, 1, 0, 9),
	createFieldParser("documentNumberCheckDigit", nil, 1, 9, 10),
	createFieldParser("birthDate", parseDate, 1, 13, 19),
	createFieldParser("birthDateCheckDigit", nil, 1, 19, 20),
	createFieldParser("sex", parseSex, 1, 20, 21),
	createFieldParser("expirationDate", parseDate, 1, 21, 27),
	createFieldParser("expirationDateCheckDigit", nil, 1, 27, 28),
	createFieldParser("nationality", nil, 1, 10, 13),
	createFieldParser("personalNumber", parseText, 1, 28, 42),
	// createFieldParser("personalNumberCheckDigit", nil, 1, 28, 42),

	// createFieldParser("optional1", nil, 0, 15, 30),
	// createFieldParser("optional2", nil, 1, 18, 29),
}
