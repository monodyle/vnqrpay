package vnqrpay

const (
	VietQRNotSupported = -1
	VietQRReceiveOnly  = 0
	VietQRSupported    = 1
)

const (
	VietQRProvider QRProvider = "VIETQR"
	VNPayProvider  QRProvider = "VNPAY"
)

const (
	FieldVersion          QRFieldID = "00"
	FieldInitMethod       QRFieldID = "01"
	FieldVNPayQR          QRFieldID = "26"
	FieldVietQR           QRFieldID = "38"
	FieldCategory         QRFieldID = "52"
	FieldCurrency         QRFieldID = "53"
	FieldAmount           QRFieldID = "54"
	FieldTipAndFeeType    QRFieldID = "55"
	FieldTipAndFeeAmount  QRFieldID = "56"
	FieldTipAndFeePercent QRFieldID = "56"
	FieldNation           QRFieldID = "58"
	FieldMerchantName     QRFieldID = "59"
	FieldCity             QRFieldID = "60"
	FieldZipCode          QRFieldID = "61"
	FieldAdditionalData   QRFieldID = "62"
	FieldCRC              QRFieldID = "63"
)

const (
	VNPayGUID  QRProviderGUID = "A000000775"
	VietQRGUID QRProviderGUID = "A000000727"
)

const (
	ADID_BillNumber                    AdditionalDataID = "01"
	ADID_MobileNumber                  AdditionalDataID = "02"
	ADID_StoreLabel                    AdditionalDataID = "03"
	ADID_LoyaltyNumber                 AdditionalDataID = "04"
	ADID_ReferenceLabel                AdditionalDataID = "05"
	ADID_CustomerLabel                 AdditionalDataID = "06"
	ADID_TerminalPurpose               AdditionalDataID = "07"
	ADID_PurposeOfTransaction          AdditionalDataID = "08"
	ADID_AdditionalConsumerDataRequest AdditionalDataID = "09"
)

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

const (
	ABBankSwiftCode                BankSwiftCode = "ABBKVNVX"
	ACBSwiftCode                   BankSwiftCode = "ASCBVNVX"
	AgribankSwiftCode              BankSwiftCode = "VBAAVNVX"
	BacABankSwiftCode              BankSwiftCode = "NASCVNVX"
	BaoVietBankSwiftCode           BankSwiftCode = "BVBVVNVX"
	BIDVSwiftCode                  BankSwiftCode = "BIDVVNVX"
	CBBankSwiftCode                BankSwiftCode = "GTBAVNVX"
	CIMBSwiftCode                  BankSwiftCode = "CIBBVNVN"
	DBSBankSwiftCode               BankSwiftCode = "DBSSVNVX"
	DongABankSwiftCode             BankSwiftCode = "EACBVNVX"
	EXIMBankSwiftCode              BankSwiftCode = "EBVIVNVX"
	GPBankSwiftCode                BankSwiftCode = "GBNKVNVX"
	HDBankSwiftCode                BankSwiftCode = "HDBCVNVX"
	HongLeongBankSwiftCode         BankSwiftCode = "HLBBVNVX"
	HSBCSwiftCode                  BankSwiftCode = "HSBCVNVX"
	KasikornBankSwiftCode          BankSwiftCode = "KASIVNVX"
	KienLongBankSwiftCode          BankSwiftCode = "KLBKVNVX"
	LienVietPostBankSwiftCode      BankSwiftCode = "LVBKVNVX"
	MBBankSwiftCode                BankSwiftCode = "MSCBVNVX"
	MSBSwiftCode                   BankSwiftCode = "MCOBVNVX"
	NamABankSwiftCode              BankSwiftCode = "NAMAVNVX"
	NCBSwiftCode                   BankSwiftCode = "NVBAVNVX"
	OCBSwiftCode                   BankSwiftCode = "ORCOVNVX"
	OceanBankSwiftCode             BankSwiftCode = "OCBKUS3M"
	PGBankSwiftCode                BankSwiftCode = "PGBLVNVX"
	PublicBankSwiftCode            BankSwiftCode = "VIDPVNVX"
	PVCOMBankSwiftCode             BankSwiftCode = "WBVNVNVX"
	SacombankSwiftCode             BankSwiftCode = "SGTTVNVX"
	SaigonBankSwiftCode            BankSwiftCode = "SBITVNVX"
	SCBSwiftCode                   BankSwiftCode = "SACLVNVX"
	SEABankSwiftCode               BankSwiftCode = "SEAVVNVX"
	SHBSwiftCode                   BankSwiftCode = "SHBAVNVX"
	ShinhanBankSwiftCode           BankSwiftCode = "SHBKVNVX"
	StandardCharteredBankSwiftCode BankSwiftCode = "SCBLVNVX"
	TechcombankSwiftCode           BankSwiftCode = "VTCBVNVX"
	TPBankSwiftCode                BankSwiftCode = "TPBVVNVX"
	VIBSwiftCode                   BankSwiftCode = "VNIBVNVX"
	VietABankSwiftCode             BankSwiftCode = "VNACVNVX"
	VietBankSwiftCode              BankSwiftCode = "VNTTVNVX"
	VietCapitalBankSwiftCode       BankSwiftCode = "VCBCVNVX"
	VietcomBankSwiftCode           BankSwiftCode = "BFTVVNVX"
	ViettinBankSwiftCode           BankSwiftCode = "ICBVVNVX"
	VPBankSwiftCode                BankSwiftCode = "VPBKVNVX"
)

var Banks = map[BankKey]Bank{
	ABBankKey: {
		Key:             ABBankKey,
		Code:            ABBankCode,
		Name:            "Ngân hàng TMCP An Bình",
		Bin:             "970425",
		ShortName:       "AB Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       ABBankSwiftCode,
		Keywords:        "anbinh",
	},
	ACBKey: {
		Key:             ACBKey,
		Code:            ACBCode,
		Name:            "Ngân hàng TMCP Á Châu",
		Bin:             "970416",
		ShortName:       "ACB",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       ACBSwiftCode,
		Keywords:        "achau",
	},
	AgribankKey: {
		Key:             AgribankKey,
		Code:            AgribankCode,
		Name:            "Ngân hàng Nông nghiệp và Phát triển Nông thôn Việt Nam",
		Bin:             "970405",
		ShortName:       "Agribank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       AgribankSwiftCode,
		Keywords:        "nongnghiep, nongthon, agribank, agri",
	},
	BacABankKey: {
		Key:             BacABankKey,
		Code:            BacABankCode,
		Name:            "Ngân hàng TMCP Bắc Á",
		Bin:             "970409",
		ShortName:       "BacA Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       BacABankSwiftCode,
		Keywords:        "baca, NASB",
	},
	BaoVietBankKey: {
		Key:             BaoVietBankKey,
		Code:            BaoVietBankCode,
		Name:            "Ngân hàng TMCP Bảo Việt",
		Bin:             "970438",
		ShortName:       "BaoViet Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       BaoVietBankSwiftCode,
		Keywords:        "baoviet, BVB",
	},
	BIDCKey: {
		Key:          BIDCKey,
		Code:         BIDCCode,
		Name:         "Ngân hàng TMCP Đầu tư và Phát triển Campuchia",
		ShortName:    "BIDC",
		VietQRStatus: VietQRNotSupported,
	},
	BIDVKey: {
		Key:             BIDVKey,
		Code:            BIDVCode,
		Name:            "Ngân hàng TMCP Đầu tư và Phát triển Việt Nam",
		Bin:             "970418",
		ShortName:       "BIDV",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       BIDVSwiftCode,
	},
	CakeKey: {
		Key:             CakeKey,
		Code:            CakeCode,
		Name:            "Ngân hàng số CAKE by VPBank - Ngân hàng TMCP Việt Nam Thịnh Vượng",
		Bin:             "546034",
		ShortName:       "CAKE by VPBank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
	},
	CBBankKey: {
		Key:             CBBankKey,
		Code:            CBBankCode,
		Name:            "Ngân hàng Thương mại TNHH MTV Xây dựng Việt Nam",
		Bin:             "970444",
		ShortName:       "CB Bank",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
		SwiftCode:       CBBankSwiftCode,
		Keywords:        "xaydungvn, xaydung",
	},
	CIMBKey: {
		Key:             CIMBKey,
		Code:            CIMBCode,
		Name:            "Ngân hàng TNHH MTV CIMB Việt Nam",
		Bin:             "422589",
		ShortName:       "CIMB Bank",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
		SwiftCode:       CIMBSwiftCode,
		Keywords:        "cimbvn",
	},
	COOPBankKey: {
		Key:             COOPBankKey,
		Code:            COOPBankCode,
		Name:            "Ngân hàng Hợp tác xã Việt Nam",
		Bin:             "970446",
		ShortName:       "Co-op Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		Keywords:        "hoptacxa, coop",
	},
	DBSBankKey: {
		Key:             DBSBankKey,
		Code:            DBSBankCode,
		Name:            "NH TNHH MTV Phát triển Singapore - Chi nhánh TP. Hồ Chí Minh",
		Bin:             "796500",
		ShortName:       "DBS Bank",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 0,
		SwiftCode:       DBSBankSwiftCode,
		Keywords:        "dbshcm",
	},
	DongABankKey: {
		Key:             DongABankKey,
		Code:            DongABankCode,
		Name:            "Ngân hàng TMCP Đông Á",
		Bin:             "970406",
		ShortName:       "DongA Bank",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
		SwiftCode:       DongABankSwiftCode,
		Keywords:        "donga, DAB",
	},
	EXIMBankKey: {
		Key:             EXIMBankKey,
		Code:            EXIMBankCode,
		Name:            "Ngân hàng TMCP Xuất Nhập khẩu Việt Nam",
		Bin:             "970431",
		ShortName:       "Eximbank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       EXIMBankSwiftCode,
	},
	GPBankKey: {
		Key:             GPBankKey,
		Code:            GPBankCode,
		Name:            "Ngân hàng Thương mại TNHH MTV Dầu Khí Toàn Cầu",
		Bin:             "970408",
		ShortName:       "GPBank",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
		SwiftCode:       GPBankSwiftCode,
		Keywords:        "daukhi",
	},
	HDBankKey: {
		Key:             HDBankKey,
		Code:            HDBankCode,
		Name:            "Ngân hàng TMCP Phát triển TP. Hồ Chí Minh",
		Bin:             "970437",
		ShortName:       "HDBank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       HDBankSwiftCode,
	},
	HongLeongBankKey: {
		Key:             HongLeongBankKey,
		Code:            HongLeongBankCode,
		Name:            "Ngân hàng TNHH MTV Hong Leong Việt Nam",
		Bin:             "970442",
		ShortName:       "HongLeong Bank",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
		SwiftCode:       HongLeongBankSwiftCode,
		Keywords:        "HLBVN",
	},
	HSBCKey: {
		Key:             HSBCKey,
		Code:            HSBCCode,
		Name:            "Ngân hàng TNHH MTV HSBC (Việt Nam)",
		Bin:             "458761",
		ShortName:       "HSBC Vietnam",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
		SwiftCode:       HSBCSwiftCode,
	},
	IBKHCMKey: {
		Key:             IBKHCMKey,
		Code:            IBKHCMCode,
		Name:            "Ngân hàng Công nghiệp Hàn Quốc - Chi nhánh TP. Hồ Chí Minh",
		Bin:             "970456",
		ShortName:       "IBK HCM",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 0,
		Keywords:        "congnghiep",
	},
	IBKHNKey: {
		Key:             IBKHNKey,
		Code:            IBKHNCode,
		Name:            "Ngân hàng Công nghiệp Hàn Quốc - Chi nhánh Hà Nội",
		Bin:             "970455",
		ShortName:       "IBK HN",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 0,
		Keywords:        "congnghiep",
	},
	IndovinaBankKey: {
		Key:             IndovinaBankKey,
		Code:            IndovinaBankCode,
		Name:            "Ngân hàng TNHH Indovina",
		Bin:             "970434",
		ShortName:       "Indovina Bank",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
	},
	KasikornBankKey: {
		Key:             KasikornBankKey,
		Code:            KasikornBankCode,
		Name:            "Ngân hàng Đại chúng TNHH KASIKORNBANK - CN TP. Hồ Chí Minh",
		Bin:             "668888",
		ShortName:       "Kasikornbank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       KasikornBankSwiftCode,
	},
	KienLongBankKey: {
		Key:             KienLongBankKey,
		Code:            KienLongBankCode,
		Name:            "Ngân hàng TMCP Kiên Long",
		Bin:             "970452",
		ShortName:       "KienlongBank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       KienLongBankSwiftCode,
	},
	KookminBankHCMKey: {
		Key:             KookminBankHCMKey,
		Code:            KookminBankHCMCode,
		Name:            "Ngân hàng Kookmin - Chi nhánh TP. Hồ Chí Minh",
		Bin:             "970463",
		ShortName:       "Kookmin Bank HCM",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 0,
	},
	KookminBankHNKey: {
		Key:             KookminBankHNKey,
		Code:            KookminBankHNCode,
		Name:            "Ngân hàng Kookmin - Chi nhánh Hà Nội",
		Bin:             "970462",
		ShortName:       "Kookmin Bank HN",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 0,
	},
	LienVietPostBankKey: {
		Key:             LienVietPostBankKey,
		Code:            LienVietPostBankCode,
		Name:            "Ngân hàng TMCP Bưu Điện Liên Việt",
		Bin:             "970449",
		ShortName:       "LienVietPostBank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       LienVietPostBankSwiftCode,
		Keywords:        "lienvietbank",
	},
	MBBankKey: {
		Key:             MBBankKey,
		Code:            MBBankCode,
		Name:            "Ngân hàng TMCP Quân đội",
		Bin:             "970422",
		ShortName:       "MB Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       MBBankSwiftCode,
	},
	MSBKey: {
		Key:             MSBKey,
		Code:            MSBCode,
		Name:            "Ngân hàng TMCP Hàng Hải",
		Bin:             "970426",
		ShortName:       "MSB",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       MSBSwiftCode,
		Keywords:        "hanghai",
	},
	NamABankKey: {
		Key:             NamABankKey,
		Code:            NamABankCode,
		Name:            "Ngân hàng TMCP Nam Á",
		Bin:             "970428",
		ShortName:       "Nam A Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       NamABankSwiftCode,
		Keywords:        "namabank",
	},
	NCBKey: {
		Key:             NCBKey,
		Code:            NCBCode,
		Name:            "Ngân hàng TMCP Quốc Dân",
		Bin:             "970419",
		ShortName:       "NCB Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       NCBSwiftCode,
		Keywords:        "quocdan",
	},
	NonghyupBankHNKey: {
		Key:             NonghyupBankHNKey,
		Code:            NonghyupBankHNCode,
		Name:            "Ngân hàng Nonghyup - Chi nhánh Hà Nội",
		Bin:             "801011",
		ShortName:       "Nonghyup Bank",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 0,
	},
	OCBKey: {
		Key:             OCBKey,
		Code:            OCBCode,
		Name:            "Ngân hàng TMCP Phương Đông",
		Bin:             "970448",
		ShortName:       "OCB Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       OCBSwiftCode,
		Keywords:        "phuongdong",
	},
	OceanBankKey: {
		Key:             OceanBankKey,
		Code:            OceanBankCode,
		Name:            "Ngân hàng Thương mại TNHH MTV Đại Dương",
		Bin:             "970414",
		ShortName:       "Ocean Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       OceanBankSwiftCode,
		Keywords:        "daiduong",
	},
	PGBankKey: {
		Key:             PGBankKey,
		Code:            PGBankCode,
		Name:            "Ngân hàng TMCP Xăng dầu Petrolimex",
		Bin:             "970430",
		ShortName:       "PG Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       PGBankSwiftCode,
	},
	PublicBankKey: {
		Key:             PublicBankKey,
		Code:            PublicBankCode,
		Name:            "Ngân hàng TNHH MTV Public Việt Nam",
		Bin:             "970439",
		ShortName:       "Public Bank Vietnam",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
		SwiftCode:       PublicBankSwiftCode,
		Keywords:        "publicvn",
	},
	PVCOMBankKey: {
		Key:             PVCOMBankKey,
		Code:            PVCOMBankCode,
		Name:            "Ngân hàng TMCP Đại Chúng Việt Nam",
		Bin:             "970412",
		ShortName:       "PVcomBank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       PVCOMBankSwiftCode,
		Keywords:        "daichung",
	},
	SacombankKey: {
		Key:             SacombankKey,
		Code:            SacombankCode,
		Name:            "Ngân hàng TMCP Sài Gòn Thương Tín",
		Bin:             "970403",
		ShortName:       "Sacombank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       SacombankSwiftCode,
	},
	SaigonBankKey: {
		Key:             SaigonBankKey,
		Code:            SaigonBankCode,
		Name:            "Ngân hàng TMCP Sài Gòn Công Thương",
		Bin:             "970400",
		ShortName:       "SaigonBank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       SaigonBankSwiftCode,
		Keywords:        "saigoncongthuong, saigonbank",
	},
	SCBKey: {
		Key:             SCBKey,
		Code:            SCBCode,
		Name:            "Ngân hàng TMCP Sài Gòn",
		Bin:             "970429",
		ShortName:       "SCB",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       SCBSwiftCode,
		Keywords:        "saigon",
	},
	SEABankKey: {
		Key:             SEABankKey,
		Code:            SEABankCode,
		Name:            "Ngân hàng TMCP Đông Nam Á",
		Bin:             "970440",
		ShortName:       "SeABank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       SEABankSwiftCode,
	},
	SHBKey: {
		Key:             SHBKey,
		Code:            SHBCode,
		Name:            "Ngân hàng TMCP Sài Gòn - Hà Nội",
		Bin:             "970443",
		ShortName:       "SHB",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       SHBSwiftCode,
		Keywords:        "saigonhanoi, sghn",
	},
	ShinhanBankKey: {
		Key:             ShinhanBankKey,
		Code:            ShinhanBankCode,
		Name:            "Ngân hàng TNHH MTV Shinhan Việt Nam",
		Bin:             "970424",
		ShortName:       "Shinhan Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       ShinhanBankSwiftCode,
	},
	StandardCharteredBankKey: {
		Key:             StandardCharteredBankKey,
		Code:            StandardCharteredBankCode,
		Name:            "Ngân hàng TNHH MTV Standard Chartered Bank Việt Nam",
		Bin:             "970410",
		ShortName:       "Standard Chartered Vietnam",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
		SwiftCode:       StandardCharteredBankSwiftCode,
	},
	TechcombankKey: {
		Key:             TechcombankKey,
		Code:            TechcombankCode,
		Name:            "Ngân hàng TMCP Kỹ thương Việt Nam",
		Bin:             "970407",
		ShortName:       "Techcombank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       TechcombankSwiftCode,
	},
	TimoKey: {
		Key:             TimoKey,
		Code:            TimoCode,
		Name:            "Ngân hàng số Timo by Bản Việt Bank",
		Bin:             "963388",
		ShortName:       "Timo",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 0,
		Keywords:        "banviet",
	},
	TPBankKey: {
		Key:             TPBankKey,
		Code:            TPBankCode,
		Name:            "Ngân hàng TMCP Tiên Phong",
		Bin:             "970423",
		ShortName:       "TPBank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       TPBankSwiftCode,
		Keywords:        "tienphong",
	},
	UBankKey: {
		Key:             UBankKey,
		Code:            UBankCode,
		Name:            "Ngân hàng số Ubank by VPBank - Ngân hàng TMCP Việt Nam Thịnh Vượng",
		Bin:             "546035",
		ShortName:       "Ubank by VPBank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
	},
	UnitedOverseasBankKey: {
		Key:             UnitedOverseasBankKey,
		Code:            UnitedOverseasBankCode,
		Name:            "Ngân hàng United Overseas Bank Việt Nam",
		Bin:             "970458",
		ShortName:       "United Overseas Bank Vietnam",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
	},
	VIBKey: {
		Key:             VIBKey,
		Code:            VIBCode,
		Name:            "Ngân hàng TMCP Quốc tế Việt Nam",
		Bin:             "970441",
		ShortName:       "VIB",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       VIBSwiftCode,
		Keywords:        "quocte",
	},
	VietABankKey: {
		Key:             VietABankKey,
		Code:            VietABankCode,
		Name:            "Ngân hàng TMCP Việt Á",
		Bin:             "970427",
		ShortName:       "VietABank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       VietABankSwiftCode,
	},
	VietBankKey: {
		Key:             VietBankKey,
		Code:            VietBankCode,
		Name:            "Ngân hàng TMCP Việt Nam Thương Tín",
		Bin:             "970433",
		ShortName:       "VietBank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       VietBankSwiftCode,
		Keywords:        "vietnamthuongtin, vnthuongtin",
	},
	VietCapitalBankKey: {
		Key:             VietCapitalBankKey,
		Code:            VietCapitalBankCode,
		Name:            "Ngân hàng TMCP Bản Việt",
		Bin:             "970454",
		ShortName:       "Viet Capital Bank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       VietCapitalBankSwiftCode,
		Keywords:        "banviet",
	},
	VietcomBankKey: {
		Key:             VietcomBankKey,
		Code:            VietcomBankCode,
		Name:            "Ngân hàng TMCP Ngoại Thương Việt Nam",
		Bin:             "970436",
		ShortName:       "Vietcombank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       VietcomBankSwiftCode,
	},
	ViettinBankKey: {
		Key:             ViettinBankKey,
		Code:            ViettinBankCode,
		Name:            "Ngân hàng TMCP Công thương Việt Nam",
		Bin:             "970415",
		ShortName:       "VietinBank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       ViettinBankSwiftCode,
		Keywords:        "viettin",
	},
	VPBankKey: {
		Key:             VPBankKey,
		Code:            VPBankCode,
		Name:            "Ngân hàng TMCP Việt Nam Thịnh Vượng",
		Bin:             "970432",
		ShortName:       "VPBank",
		VietQRStatus:    VietQRSupported,
		LookupSupported: 1,
		SwiftCode:       VPBankSwiftCode,
		Keywords:        "vnthinhvuong",
	},
	VRBKey: {
		Key:             VRBKey,
		Code:            VRBCode,
		Name:            "Ngân hàng Liên doanh Việt - Nga",
		Bin:             "970421",
		ShortName:       "VietNgaBank",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
		Keywords:        "vietnam-russia, vrbank",
	},
	WooriBankKey: {
		Key:             WooriBankKey,
		Code:            WooriBankCode,
		Name:            "Ngân hàng TNHH MTV Woori Việt Nam",
		Bin:             "970457",
		ShortName:       "Woori Bank",
		VietQRStatus:    VietQRReceiveOnly,
		LookupSupported: 1,
	},
}
