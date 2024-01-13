package vnqrpay

import (
	"log"
	"strings"
	"testing"
)

func TestMakeCRC(t *testing.T) {
	log.Println("TestMakeCRC")
	content := "00020101021138540010A00000072701240006970418011062540463830208QRIBFTTA53037045802VN6304AB32"
	check := content[:len(content)-4]
	extractCrc := strings.ToUpper(content[len(content)-4:])
	crc := MakeCRC(check)

	if extractCrc != crc {
		t.Fatalf("CRC expect: %v, got: %v", crc, extractCrc)
	}
}

func TestVerifyCRC(t *testing.T) {
	log.Println("TestVerifyCRC")
	content := "00020101021138540010A00000072701240006970418011062540463830208QRIBFTTA53037045802VN6304AB32"
	if !VerifyCRC(content) {
		t.Fatal("Invalid CRC")
	}
}

func TestVietQR(t *testing.T) {
	log.Println("TestVietQR")
	content := "00020101021138540010A00000072701240006970418011062540463830208QRIBFTTA53037045802VN6304AB32"

	qr, err := Parse(content)
	if err != nil {
		t.Fatalf("cannot parse %v", err)
	}

	if !qr.IsValid {
		t.Fatalf("qr.IsValid expect: true, got: %v", qr.IsValid)
	}

	if qr.Consumer.BankNumber != "6254046383" {
		t.Fatalf("qr.Consumer.BankNumber expect: 6254046383, got: %v", qr.Consumer.BankNumber)
	}
}

func TestDynamicVietQR(t *testing.T) {
	log.Println("TestDynamicVietQR")
	content := "00020101021238540010A00000072701240006970418011062540463830208QRIBFTTA53037045405240005802VN62180814a cup of cofee63043301"

	qr, err := Parse(content)

	if err != nil {
		t.Fatalf("cannot parse %v", err)
	}

	if !qr.IsValid {
		t.Fatalf("qr.IsValid expect: true, got: %v", qr.IsValid)
	}

	if qr.Consumer.BankNumber != "6254046383" {
		t.Fatalf("qr.Consumer.BankNumber expect: 6254046383, got: %v", qr.Consumer.BankNumber)
	}

	if qr.Amount != "24000" {
		t.Fatalf("qr.Amount expect: 24000, got: %v", qr.Amount)
	}

	if qr.AdditionalData.Purpose != "a cup of cofee" {
		t.Fatalf("qr.Amount expect: a cup of cofee, got: %v", qr.Amount)
	}
}

func TestVNPay(t *testing.T) {
	log.Println("TestVNPay")
	content := "00020101021126280010A0000007750110010531314453037045408210900005802VN5910CELLPHONES62600312CPSHN ONLINE0517021908061613127850705ONLHN0810CellphoneS63047685"

	qr, err := Parse(content)

	if err != nil {
		t.Fatalf("cannot parse %v", err)
	}

	if !qr.IsValid {
		t.Fatalf("qr.IsValid expect: true, got: %v", qr.IsValid)
	}

	if qr.Provider.Name != VNPayProvider {
		t.Fatalf("qr.Provider.Name expect: %v, got: %v", VNPayProvider, qr.Provider.Name)
	}

	if qr.Merchant.ID != "0105313144" {
		t.Fatalf("qr.Merchant.ID expect: %v, got: %v", "0105313144", qr.Merchant.ID)
	}

	if qr.Amount != "21090000" {
		t.Fatalf("qr.Amount expect: %v, got: %v", "21090000", qr.Amount)
	}

	if qr.AdditionalData.Store != "CPSHN ONLINE" {
		t.Fatalf("qr.AdditionalData.Store expect: %v, got: %v", "CPSHN ONLINE", qr.AdditionalData.Store)
	}

	if qr.AdditionalData.Terminal != "ONLHN" {
		t.Fatalf("qr.AdditionalData.Terminal expect: %v, got: %v", "ONLHN", qr.AdditionalData.Terminal)
	}

	if qr.AdditionalData.Purpose != "CellphoneS" {
		t.Fatalf("qr.AdditionalData.Store expect: %v, got: %v", "CellphoneS", qr.AdditionalData.Purpose)
	}

	if qr.AdditionalData.Reference != "02190806161312785" {
		t.Fatalf("qr.AdditionalData.Reference expect: %v, got: %v", "02190806161312785", qr.AdditionalData.Reference)
	}
}

func TestCreateVietQRCode(t *testing.T) {
	log.Println("TestCreateVietQRCode")
	qr, err := CreateVietQR(VietQROptions{
		Consumer: Consumer{
			BankBin:    Banks[BIDVKey].Bin,
			BankNumber: "6254046383",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	content, err := qr.Build()
	if err != nil {
		t.Fatal(err)
	}

	if content != "00020101021138540010A00000072701240006970418011062540463830208QRIBFTTA53037045802VN6304AB32" {
		t.Fatalf("expect: %v, got: %v", "00020101021138540010A00000072701240006970418011062540463830208QRIBFTTA53037045802VN6304AB32", qr.AdditionalData.Reference)
	}
}

func TestCreateDynamicVietQRCode(t *testing.T) {
	log.Println("TestCreateVietQRCode")
	qr, err := CreateVietQR(VietQROptions{
		Consumer: Consumer{
			BankBin:    Banks[BIDVKey].Bin,
			BankNumber: "6254046383",
		},
		Amount:  "24000",
		Purpose: "a cup of cofee",
	})
	if err != nil {
		t.Fatal(err)
	}

	content, err := qr.Build()
	if err != nil {
		t.Fatal(err)
	}

	if content != "00020101021138540010A00000072701240006970418011062540463830208QRIBFTTA53037045405240005802VN62180814a cup of cofee6304C455" {
		t.Fatalf("expect: %v, got: %v", "00020101021138540010A00000072701240006970418011062540463830208QRIBFTTA53037045405240005802VN62180814a cup of cofee6304C455", content)
	}
}
