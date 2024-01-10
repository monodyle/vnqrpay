package vnqrpay

type VietQRStatus int

const (
	VietQRNotSupported = -1
	VietQRReceiveOnly  = 0
	VietQRSupported    = 1
)

type QRProvider string

const (
	VietQRProvider QRProvider = "VIETQR"
	VNPayProvider  QRProvider = "VNPAY"
)

type Provider struct {
	FieldID string
	Name    QRProvider
	GUID    string
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

type QRPay struct {
	IsValid    bool
	Version    string
	InitMethod string
	Provider
	Merchant
}
