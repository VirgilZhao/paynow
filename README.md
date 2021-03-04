# paynow
Singapore Paynow  generator for golang

# useage

```go
import "github.com/VirgilZhao/paynow"

qrcodeStr := paynow.GeneratePayNowString(paynow.Options{
 		UEN:             "12345678",
 		Editable:        false,
 		Expiry:          "20260304",
 		CompanyName:     "testcompany",
 		Amount:          "0.99",
 		ReferenceNumber: "testordernumber12345678",
 	})
```
Example will return qrcodeStr="00020101021226470009SG.PAYNOW010120208123456780301004082026030452040000530370254040.995802SG5911testcompany6009Singapore62270123testordernumber1234567863040047"
which you can send to front-end to show in QR format 

**Option Description**

| name | required |data type| desc |
| ---- | ---- | ---- | ---- |
| UEN | required | string | UEN number for merchant|
|Editable | optional | bool | true: client can edit the amount; false: client cannot edit the amount, default is false |
|Expiry | optional | string | payment expiry date, format is YYYYMMDD|
|CompanyName | optional | string | merchant name |
|Amount| required | string | pay amount, format as "99.99", "1.99"|
|ReferenceNumber| required | string | bill number/ID related to your own system|

# Data Structure Of The Generate QR Code String Content
**Paynow qrcode string data structure follows the SGQR specification, you can download the pdf in below link**
https://www.emvco.com/wp-content/uploads/documents/EMVCo-Merchant-Presented-QR-Specification-v1.1.pdf

# Credits
Original code reference from: https://gist.github.com/chengkiang/7e1c4899768245570cc49c7d23bc394c

