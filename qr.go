package vnqrpay

import (
	"errors"
	"fmt"
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

func Parse(content string) (qr *QRPay, err error) {
	qr = &QRPay{}

	if len(content) < 4 {
		return qr, errors.New("invalid QR: too short")
	}

	qr.IsValid = VerifyCRC(content)

	if !qr.IsValid {
		return qr, errors.New("invalid QR: CRC is invalid")
	}

	qr.parseQRContent(content)

	return
}

func (qr *QRPay) Build() (content string, err error) {
	version := combineFieldData(FieldVersion, defaultStrValue(qr.Version, "01"))
	initMethod := combineFieldData(FieldInitMethod, defaultStrValue(qr.Version, "11"))
	guid := combineFieldData(QRFieldID(ProviderFieldGUID), string(qr.Provider.GUID))

	providerDataContent := ""
	if qr.Provider.GUID == VietQRGUID {
		bankBin := combineFieldData(QRFieldID(VietQRBankBin), qr.Consumer.BankBin)
		bankNumber := combineFieldData(QRFieldID(VietQRBankNumber), qr.Consumer.BankNumber)
		providerDataContent = bankBin + bankNumber
	} else if qr.Provider.GUID == VNPayGUID {
		providerDataContent = defaultStrValue(qr.Merchant.ID, "")
	}
	provider := combineFieldData(QRFieldID(ProviderFieldData), providerDataContent)
	service := combineFieldData(QRFieldID(ProviderFieldService), qr.Provider.Service)
	providerData := combineFieldData(qr.Provider.FieldID, guid+provider+service)

	category := combineFieldData(FieldCategory, qr.Category)
	currency := combineFieldData(FieldCurrency, defaultStrValue(qr.Currency, "704"))
	amount := combineFieldData(FieldAmount, qr.Amount)
	tipAndFeeType := combineFieldData(FieldTipAndFeeType, qr.TipAndFeeType)
	tipAndFeeAmount := combineFieldData(FieldTipAndFeeAmount, qr.TipAndFeeAmount)
	tipAndFeePercent := combineFieldData(FieldTipAndFeePercent, qr.TipAndFeePercent)
	nation := combineFieldData(FieldNation, defaultStrValue(qr.Nation, "VN"))
	merchantName := combineFieldData(FieldMerchantName, qr.Merchant.Name)
	city := combineFieldData(FieldCity, qr.City)
	zipCode := combineFieldData(FieldZipCode, qr.ZipCode)

	buildNumber := combineFieldData(QRFieldID(AdditionalDataBillNumber), qr.AdditionalData.BillNumber)
	mobileNumber := combineFieldData(QRFieldID(AdditionalDataMobileNumber), qr.AdditionalData.MobileNumber)
	storeLabel := combineFieldData(QRFieldID(AdditionalDataStoreLabel), qr.AdditionalData.Store)
	loyaltyNumber := combineFieldData(QRFieldID(AdditionalDataLoyaltyNumber), qr.AdditionalData.LoyaltyNumber)
	reference := combineFieldData(QRFieldID(AdditionalDataReferenceLabel), qr.AdditionalData.Reference)
	customerLabel := combineFieldData(QRFieldID(AdditionalDataCustomerLabel), qr.AdditionalData.CustomerLabel)
	terminal := combineFieldData(QRFieldID(AdditionalDataTerminalLabel), qr.AdditionalData.Terminal)
	purpose := combineFieldData(QRFieldID(AdditionalDataPurposeOfTransaction), qr.AdditionalData.Purpose)
	dataRequest := combineFieldData(QRFieldID(AdditionalDataAdditionalConsumerDataRequest), qr.AdditionalData.DataRequest)

	additionalDataContent := buildNumber + mobileNumber + storeLabel + loyaltyNumber + reference + customerLabel + terminal + purpose + dataRequest
	additionalData := combineFieldData(FieldAdditionalData, additionalDataContent)

	content = version + initMethod + providerData + category + currency + amount + tipAndFeeType + tipAndFeeAmount + tipAndFeePercent + nation + merchantName + city + zipCode + additionalData + string(FieldCRC) + "04"
	crc := MakeCRC(content)
	content += crc

	return
}

func CreateVietQR(opts VietQROptions) (qr *QRPay, err error) {
	//
	return
}

func CreateVNPay(opts VNPayOptions) (qr *QRPay, err error) {
	//
	return
}

func (qr *QRPay) parseQRContent(content string) error {
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
		err = qr.parseProviderInfo(value)
		if err != nil {
			return err
		}
	case FieldVNPayQR:
		qr.Provider.FieldID = QRFieldID(id)
		err = qr.parseProviderInfo(value)
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
		err = qr.parseAdditionalData(value)
		if err != nil {
			return err
		}
	case FieldCRC:
		qr.CRC = value
	}

	if len(nextValue) > 4 {
		qr.parseQRContent(nextValue)
		return err
	}

	return err
}

func defaultStrValue(val, defaultVal string) string {
	if val == "" {
		return defaultVal
	}
	return val
}

func combineFieldData(id QRFieldID, value string) string {
	if len(id) != 2 || len(value) == 0 {
		return ""
	}
	length := fmt.Sprintf("%02d", len(value))
	return fmt.Sprintf("%v%v%v", id, length, value)
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

func (qr *QRPay) parseProviderInfo(content string) error {
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
			err := qr.parseVietQRConsumer(value)
			if err != nil {
				return err
			}
		}
	case ProviderFieldService:
		qr.Provider.Service = value
	}

	if len(nextValue) > 4 {
		qr.parseProviderInfo(nextValue)
	}
	return nil
}

func (qr *QRPay) parseVietQRConsumer(content string) error {
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
		qr.parseVietQRConsumer(nextValue)
	}
	return nil
}

func (qr *QRPay) parseAdditionalData(content string) error {
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
	case AdditionalDataTerminalLabel:
		qr.AdditionalData.Terminal = value
	}

	if len(nextValue) > 4 {
		qr.parseAdditionalData(nextValue)
	}
	return nil
}
