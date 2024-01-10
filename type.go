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

type Merchant struct {
	ID   string
	Name string
}

type BankKey string

const (
	ABBankKey                BankKey = "abbank"
	ACBKey                   BankKey = "acb"
	AgribankKey              BankKey = "agribank"
	BacABankKey              BankKey = "bacabank"
	BaoVietBankKey           BankKey = "baoviet"
	BIDCKey                  BankKey = "bidc"
	BIDVKey                  BankKey = "bidv"
	CakeKey                  BankKey = "cake"
	CBBankKey                BankKey = "cbbank"
	CIMBKey                  BankKey = "cimb"
	COOPBankKey              BankKey = "coopbank"
	DBSBankKey               BankKey = "dbsbank"
	DongABankKey             BankKey = "dongabank"
	EXIMBankKey              BankKey = "eximbank"
	GPBankKey                BankKey = "gpbank"
	HDBankKey                BankKey = "hdbank"
	HongLeongBankKey         BankKey = "hongleongbank"
	HSBCKey                  BankKey = "hsbc"
	IBKHCMKey                BankKey = "ibkhcm"
	IBKHNKey                 BankKey = "ibkhn"
	IndovinaBankKey          BankKey = "indovinabank"
	KasikornBankKey          BankKey = "kasikorn"
	KienLongBankKey          BankKey = "kienlongbank"
	KookminBankHCMKey        BankKey = "kookminhcm"
	KookminBankHNKey         BankKey = "kookminhn"
	LienVietPostBankKey      BankKey = "lienvietpostbank"
	MBBankKey                BankKey = "mbbank"
	MSBKey                   BankKey = "msb"
	NamABankKey              BankKey = "namabank"
	NCBKey                   BankKey = "ncb"
	NonghyupBankHNKey        BankKey = "nonghyup"
	OCBKey                   BankKey = "ocb"
	OceanBankKey             BankKey = "oceanbank"
	PGBankKey                BankKey = "pgbank"
	PublicBankKey            BankKey = "publicbank"
	PVCOMBankKey             BankKey = "pvcombank"
	SacombankKey             BankKey = "sacombank"
	SaigonBankKey            BankKey = "saigonbank"
	SCBKey                   BankKey = "scb"
	SEABankKey               BankKey = "seabank"
	SHBKey                   BankKey = "shb"
	ShinhanBankKey           BankKey = "shinhanbank"
	StandardCharteredBankKey BankKey = "standardcharteredbank"
	TechcombankKey           BankKey = "techcombank"
	TimoKey                  BankKey = "timo"
	TPBankKey                BankKey = "tpbank"
	UBankKey                 BankKey = "ubank"
	UnitedOverseasBankKey    BankKey = "uob"
	VIBKey                   BankKey = "vib"
	VietABankKey             BankKey = "vietabank"
	VietBankKey              BankKey = "vietbank"
	VietCapitalBankKey       BankKey = "banviet"
	VietcomBankKey           BankKey = "vietcombank"
	ViettinBankKey           BankKey = "vietinbank"
	VPBankKey                BankKey = "vpbank"
	VRBKey                   BankKey = "vrb"
	WooriBankKey             BankKey = "wooribank"
)

type BankCode string

const (
	ABBankCode                BankCode = "ABB"
	ACBCode                   BankCode = "ACB"
	AgribankCode              BankCode = "AGRIBANK"
	BacABankCode              BankCode = "BAB"
	BaoVietBankCode           BankCode = "BAOVIETBANK"
	BIDCCode                  BankCode = "BIDC"
	BIDVCode                  BankCode = "BID"
	CakeCode                  BankCode = "CAKE"
	CBBankCode                BankCode = "VNCB"
	CIMBCode                  BankCode = "CIMB"
	COOPBankCode              BankCode = "COOPBANK"
	DBSBankCode               BankCode = "DBS"
	DongABankCode             BankCode = "DONGABANK"
	EXIMBankCode              BankCode = "EIB"
	GPBankCode                BankCode = "GPBANK"
	HDBankCode                BankCode = "HDB"
	HongLeongBankCode         BankCode = "HLB"
	HSBCCode                  BankCode = "HSBC"
	IBKHCMCode                BankCode = "IBKHCM"
	IBKHNCode                 BankCode = "IBKHN"
	IndovinaBankCode          BankCode = "IVB"
	KasikornBankCode          BankCode = "KBANK"
	KienLongBankCode          BankCode = "KLB"
	KookminBankHCMCode        BankCode = "KBHCM"
	KookminBankHNCode         BankCode = "KBHN"
	LienVietPostBankCode      BankCode = "LPB"
	MBBankCode                BankCode = "MBB"
	MSBCode                   BankCode = "MSB"
	NamABankCode              BankCode = "NAB"
	NCBCode                   BankCode = "NVB"
	NonghyupBankHNCode        BankCode = "NONGHYUP"
	OCBCode                   BankCode = "OCB"
	OceanBankCode             BankCode = "OCEANBANK"
	PGBankCode                BankCode = "PGB"
	PublicBankCode            BankCode = "PBVN"
	PVCOMBankCode             BankCode = "PVCOMBANK"
	SacombankCode             BankCode = "STB"
	SaigonBankCode            BankCode = "SGB"
	SCBCode                   BankCode = "SCB"
	SEABankCode               BankCode = "SSB"
	SHBCode                   BankCode = "SHB"
	ShinhanBankCode           BankCode = "SVB"
	StandardCharteredBankCode BankCode = "SC"
	TechcombankCode           BankCode = "TCB"
	TimoCode                  BankCode = "TIMO"
	TPBankCode                BankCode = "TPB"
	UBankCode                 BankCode = "UBANK"
	UnitedOverseasBankCode    BankCode = "UOB"
	VIBCode                   BankCode = "VIB"
	VietABankCode             BankCode = "VAB"
	VietBankCode              BankCode = "VBB"
	VietCapitalBankCode       BankCode = "BVB"
	VietcomBankCode           BankCode = "VCB"
	ViettinBankCode           BankCode = "CTG"
	VPBankCode                BankCode = "VPB"
	VRBCode                   BankCode = "VRB"
	WooriBankCode             BankCode = "WRB"
)

type Bank struct {
	Key             BankKey
	Code            BankCode
	Name            string
	ShortName       string
	bin             string
	VietQRStatus    int
	LookupSupported int
	SwiftCode       string
	Keywords        string
}

type QRPay struct {
	IsValid    bool
	Version    string
	InitMethod string
	Provider
	Merchant
}
