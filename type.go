package vnqrpay

type VietQRStatus int

type QRProvider string

type QRProviderFieldID string

type QRProviderGUID string

type QRFieldID string

type Provider struct {
	FieldID QRFieldID
	Name    QRProvider
	GUID    QRProviderGUID
	Service string
}

type Merchant struct {
	ID   string
	Name string
}

type BankKey string

type BankCode string

type BankSwiftCode string

type Bank struct {
	Key             BankKey
	Code            BankCode
	Name            string
	ShortName       string
	Bin             string
	VietQRStatus    VietQRStatus
	LookupSupported int
	SwiftCode       BankSwiftCode
	Keywords        string
}

type VietQRConsumerID string

type Consumer struct {
	BankBin    string
	BankNumber string
}

type AdditionalDataID string

type AdditionalData struct {
	Store         string
	Terminal      string
	BillNumber    string
	MobileNumber  string
	LoyaltyNumber string
	Reference     string
	CustomerLabel string
	Purpose       string
	DataRequest   string
}

type QRPay struct {
	IsValid          bool
	Version          string
	InitMethod       string
	Provider         Provider
	Merchant         Merchant
	Consumer         Consumer
	Category         string
	Currency         string
	Amount           string
	TipAndFeeType    string
	TipAndFeeAmount  string
	TipAndFeePercent string
	Nation           string
	City             string
	ZipCode          string
	AdditionalData   AdditionalData
	CRC              string
}

type VietQRService string

type VietQROptions struct {
	Consumer Consumer
	Amount   string
	Purpose  string
	Service  VietQRService
}

type VNPayOptions struct {
	Merchant       Merchant
	AdditionalData AdditionalData
}
