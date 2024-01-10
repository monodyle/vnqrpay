package vnqrpay

import (
	"errors"
	"strconv"
	"strings"
)

func MakeCRC(content string) string {
	sum := CRC16CCITT(content)
	hex := "0000" + strconv.FormatInt(int64(sum), 16)
	CRC := strings.ToLower(hex[len(hex)-4:])
	return CRC
}

func VerifyCRC(content string) bool {
	check := content[:len(content)-4]
	crcCode := strings.ToUpper(content[len(content)-4:])

	return MakeCRC(check) == crcCode
}

func Parse(content string) (QRPay, error) {
	qr := QRPay{}

	if len(content) < 4 {
		return qr, errors.New("invalid QR: too short")
	}

	if !VerifyCRC(content) {
		return qr, errors.New("invalid QR: CRC is invalid")
	}

	var id QRFieldID
	id = QRFieldID(content[:2])
	length, err := strconv.Atoi(content[2:4])
	if err != nil {
		return qr, err
	}
	value := content[4 : 4+length]
	nextValue := content[4+length:]

	switch id {
	case FieldVersion:
		qr.Version = value
		break
	}

	return qr, nil
}
