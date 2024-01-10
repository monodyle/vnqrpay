package vnqrpay

import (
	"errors"
	"strconv"
	"strings"
)

func MakeCRC(content string) string {
	sum := CRC16CCITT(content)
	hex := "0000" + strconv.FormatInt(int64(sum), 16)
	result := strings.ToUpper(hex[len(hex)-4:])
	return result
}

func VerifyCRC(content string) bool {
	value := content[:len(content)-4]
	extracted := strings.ToUpper(content[len(content)-4:])
	crc := MakeCRC(value)

	return crc == extracted
}

func sliceContent(content string) (string, string, string, error) {
	id := content[:2]
	length, err := strconv.Atoi(content[2:4])
	if err != nil {
		return id, "", "", err
	}
	value := content[4 : 4+length]
	nextValue := content[4+length:]
	return id, value, nextValue, nil
}

func (qr *QRPay) Parse(content string) error {
	if len(content) < 4 {
		return errors.New("invalid QR: too short")
	}

	qr.IsValid = VerifyCRC(content)

	if !qr.IsValid {
		return errors.New("invalid QR: CRC is invalid")
	}

	return nil
}

func (qr *QRPay) ParseQRContent(content string) error {
	id, value, nextValue, err := sliceContent(content)
	if err != nil {
		return err
	}

	switch QRFieldID(id) {
	case FieldVersion:
		qr.Version = value
	case FieldInitMethod:
		qr.InitMethod = value
	case FieldVietQR:
		qr.Provider.FieldID = QRFieldID(id)
	case FieldVNPayQR:
		qr.Provider.FieldID = QRFieldID(id)
		err := qr.ParseProviderInfo(value)
		if err != nil {
			return err
		}
	case FieldCategory:
		qr.Category = value
	case FieldCurrency:
		qr.Currency = value
	case FieldAmount:
		qr.Amount = value
	case FieldTipAndFeeAmount:
		qr.TipAndFeeAmount = value
	case FieldTipAndFeeType:
		qr.TipAndFeeType = value
	case FieldTipAndFeePercent:
		qr.TipAndFeePercent = value
	case FieldNation:
		qr.Nation = value
	case FieldMerchantName:
		qr.Merchant.Name = value
	case FieldCity:
		qr.City = value
	case FieldZipCode:
		qr.ZipCode = value
	case FieldAdditionalData:
		err := qr.ParseAdditionalData(value)
		if err != nil {
			return err
		}
	case FieldCRC:
		qr.CRC = value
	}

	if len(nextValue) > 4 {
		return qr.Parse(nextValue)
	}

	return nil
}

func (qr *QRPay) ParseProviderInfo(content string) error {
	id, value, nextValue, err := sliceContent(content)
	if err != nil {
		return err
	}

	switch QRProviderFieldID(id) {
	case ProviderFieldGUID:
		qr.Provider.GUID = QRProviderGUID(value)
	case ProviderFieldData:
		if qr.Provider.GUID == VNPayGUID {
			qr.Provider.Name = VNPayProvider
			qr.Merchant.ID = value
		} else if qr.Provider.GUID == VietQRGUID {
			qr.Provider.Name = VietQRProvider
			err := qr.ParseVietQRConsumer(value)
			if err != nil {
				return err
			}
		}
	case ProviderFieldService:
		qr.Provider.Service = value
	}

	if len(nextValue) > 4 {
		qr.ParseProviderInfo(nextValue)
	}
	return nil
}

func (qr *QRPay) ParseVietQRConsumer(content string) error {
	id, value, nextValue, err := sliceContent(content)
	if err != nil {
		return err
	}

	switch VietQRConsumerID(id) {
	case VietQRBankBin:
		qr.Consumer.BankBin = value
	case VietQRBankNumber:
		qr.Consumer.BankNumber = value
	}

	if len(nextValue) > 4 {
		qr.ParseVietQRConsumer(nextValue)
	}
	return nil
}

func (qr *QRPay) ParseAdditionalData(content string) error {
	id, value, nextValue, err := sliceContent(content)
	if err != nil {
		return err
	}

	switch AdditionalDataID(id) {
	case AdditionalDataPurposeOfTransaction:
		qr.AdditionalData.Purpose = value
	case AdditionalDataBillNumber:
		qr.AdditionalData.BillNumber = value
	case AdditionalDataMobileNumber:
		qr.AdditionalData.MobileNumber = value
	case AdditionalDataReferenceLabel:
		qr.AdditionalData.Reference = value
	case AdditionalDataStoreLabel:
		qr.AdditionalData.Store = value
	case AdditionalDataTerminalPurpose:
		qr.AdditionalData.Terminal = value
	}

	if len(nextValue) > 4 {
		qr.ParseAdditionalData(nextValue)
	}
	return nil
}
