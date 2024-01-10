package vnqrpay

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

type QRPay struct {
	IsValid    bool
	Version    string
	InitMethod string
	Provider
}
