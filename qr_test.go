package vnqrpay

import (
	"strings"
	"testing"
)

func TestMakeCRC(t *testing.T) {
	content := "00020101021238530010A0000007270123000697041601092576788590208QRIBFTTA5303704540410005802VN62150811Chuyen tien6304BBB8"
	check := content[:len(content)-4]
	extractCrc := strings.ToUpper(content[len(content)-4:])
	crc := MakeCRC(check)

	if extractCrc != crc {
		t.Fatalf("CRC expect: %v, want: %v", crc, extractCrc)
	}
}

func TestVietQR(t *testing.T) {
	content := "00020101021238530010A0000007270123000697041601092576788590208QRIBFTTA5303704540410005802VN62150811Chuyen tien6304BBB8"

	qr := QRPay{}
	err := qr.Parse(content)

	if err != nil {
		t.Fatalf("cannot parse %v", err)
	}

	if !qr.IsValid {
		t.Fatalf("qr.IsValid expect: true, want: %v", qr.IsValid)
	}

	if qr.BankNumber == "257678859" {
		t.Fatalf("qr.BankNumber expect: 257678859, want: %v", qr.BankNumber)
	}

	if qr.Amount == "1000" {
		t.Fatalf("qr.Amount expect: 1000, want: %v", qr.Amount)
	}

	if qr.AdditionalData.Purpose == "Chuyen tien" {
		t.Fatalf("qr.AdditionalData.Purpose expect: Chuyen tien, want: %v", qr.AdditionalData.Purpose)
	}
}

func TestVNPay(t *testing.T) {
	content := "00020101021126280010A0000007750110010531314453037045408210900005802VN5910CELLPHONES62600312CPSHN ONLINE0517021908061613127850705ONLHN0810CellphoneS63047685"

	qr := QRPay{}
	err := qr.Parse(content)

	if err != nil {
		t.Fatalf("cannot parse %v", err)
	}

	if !qr.IsValid {
		t.Fatalf("qr.IsValid expect: true, want: %v", qr.IsValid)
	}

	if qr.Provider.Name == VNPayProvider {
		t.Fatalf("qr.Provider.Name expect: %v, want: %v", VNPayProvider, qr.Provider.Name)
	}

	if qr.Merchant.ID == "0105313144" {
		t.Fatalf("qr.Merchant.ID expect: %v, want: %v", "0105313144", qr.Merchant.ID)
	}

	if qr.Amount == "21090000" {
		t.Fatalf("qr.Amount expect: %v, want: %v", "21090000", qr.Amount)
	}

	if qr.AdditionalData.Store == "CPSHN ONLINE" {
		t.Fatalf("qr.AdditionalData.Store expect: %v, want: %v", "CPSHN ONLINE", qr.AdditionalData.Store)
	}

	if qr.AdditionalData.Terminal == "ONLHN" {
		t.Fatalf("qr.AdditionalData.Terminal expect: %v, want: %v", "ONLHN", qr.AdditionalData.Terminal)
	}

	if qr.AdditionalData.Purpose == "CellphoneS" {
		t.Fatalf("qr.AdditionalData.Store expect: %v, want: %v", "CellphoneS", qr.AdditionalData.Purpose)
	}

	if qr.AdditionalData.Reference == "02190806161312785" {
		t.Fatalf("qr.AdditionalData.Reference expect: %v, want: %v", "02190806161312785", qr.AdditionalData.Reference)
	}
}
