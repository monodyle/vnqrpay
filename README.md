# Vietnam QRPay Golang

<img src=".github/vietqr.png" height="25" /> <img src=".github/vnpay.png" height="19" />

<img src=".github/qr.png" height="128" align="right" />

The library helps you to encode/decode QR code of VietQR/VNPay. Converted to go from the original typescript repository [xuannghia/vietnam-qr-pay]

[xuannghia/vietnam-qr-pay]: https://github.com/xuannghia/vietnam-qr-pay/

## Features

### Build QRCode

#### Static VietQR

```go
qr, err := vnqrpay.CreateVietQR(vnqrpay.VietQROptions{
  Consumer: vnqrpay.Consumer{
    BankBin:    vnqrpay.Banks[vnqrpay.BIDVKey].Bin,
    BankNumber: "6254046383",
  },
})
if err != nil {
  log.Fatal(err)
}

content, err := qr.Build()
if err != nil {
  log.Fatal(err)
}

fmt.Println(content)
```

#### Dynamic VietQR
```go
qr, err := vnqrpay.CreateVietQR(vnqrpay.VietQROptions{
  Consumer: vnqrpay.Consumer{
    BankBin:    vnqrpay.Banks[vnqrpay.BIDVKey].Bin,
    BankNumber: "6254046383",
  },
  Amount:  "24000",
  Purpose: "morning coffee",
})
if err != nil {
  log.Fatal(err)
}

content, err := qr.Build()
if err != nil {
  log.Fatal(err)
}

fmt.Println(content)
```

#### VNPay

```go
qr, err := vnqrpay.CreateVNPay(vnqrpay.VNPayOptions{
  Merchant: vnqrpay.Merchant{
    ID:   "0123456789",
    Name: "COMPANY",
  },
  AdditionalData: vnqrpay.AdditionalData{
    Store:    "YOUR COMPANY",
    Terminal: "YC01",
  },
})
if err != nil {
  log.Fatal(err)
}

content, err := qr.Build()
if err != nil {
  log.Fatal(err)
}

fmt.Println(content)
```

#### Build your own

```go
var qr vnqrpay.QRPay = vnqrpay.QRPay{
  ...
}

content, err := qr.Build()
if err != nil {
  log.Fatal(err)
}

fmt.Println(content)
```

### Parse QRCode

```go
content := "00020101021138540010A00000072701240006970418011062540463830208QRIBFTTA53037045802VN6304AB32"
qr, err := vnqrpay.Parse(content)
if err != nil {
  log.Fatal(err)
}
if !qr.IsValid {
  log.Fatal("invalid qr content")
}

fmt.Println(qr.Consumer.BankNumber) // 6254046383
```

## Credits

This product was built as part of my hobby of exploring and learning.

Special thanks go to
- [@xuannghia](https://github.com/xuannghia/) and his awesome repository [xuannghia/vietnam-qr-pay]
