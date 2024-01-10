package vnqrpay

type VietQRStatus int

type QRProvider string

type QRProviderGUID string

type QRFieldID string

type Provider struct {
	FieldID string
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

type Consumer struct {
	BankBin    string
	BankNumber string
}

type AdditionalDataID string

type AdditionalData struct {
	BillNumber    string
	MobileNumber  string
	Store         string
	LoyaltyNumber string
	Reference     string
	CustomerLabel string
	Terminal      string
	Purpose       string
	DataRequest   string
}

type QRPay struct {
	IsValid    bool
	Version    string
	InitMethod string
	Provider
	Merchant
	Consumer
	Category         string
	Currency         string
	Amount           string
	TipAndFeeType    string
	TipAndFeeAmount  string
	TipAndFeePercent string
	Nation           string
	City             string
	ZipCode          string
	AdditionalData
	CRC string
}
