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
|Amount| required | string | pay amount, format as "99.99", "1.99", "1.80", "2.00"|
|ReferenceNumber| required | string | bill number/ID related to your own system|

# Data Structure Of The Generate QR Code String Content
**Paynow qrcode string data structure follows the SGQR specification, you can download the pdf in below link**
https://www.emvco.com/wp-content/uploads/documents/EMVCo-Merchant-Presented-QR-Specification-v1.1.pdf

**Data Structure Introduction**

>Usage Example "00020101021226470009SG.PAYNOW010120208123456780301004082026030452040000530370254040.995802SG5911testcompany6009Singapore62270123testordernumber1234567863040047"

A data object is represented as an ID/Length/Value combination, example "000201", 00 is the ID, 02 is the value length, 01 is the value.

A entire data message is the combination of data objects, below table shows the specification of all data objects

| name | ID | length | value | example | Comment |
| ---- | ---- | ---- | ---- | ---- | ---- |
 |Payload Format Indicator | 00 | 02 |  01 |  000201 |static value|
 |Point of Initiation Method | 01 | 02 | 12 | 010212 | value 11 represents static QR code, 12 represents dynamic QR code|
 |Merchant Account Infomation| 26 | up to 99 | |26470009SG.PAYNOW0101202081234567803010040820260304 | paynow account info |
 |Merchant Category Code | 52 | 04 | 52040000 | not used, set value as 0000 |
 |Merchant Currency | 53 | 03 | 702 | 5303702 | SG currency static value |
 |Transaction Amount | 54 | up to 13 |  | 54040.99 | amount |
 |Country Code | 58 | 02 | SG | 5802SG | static value for singapore |
 |Merchant Name|59|up to 25| |5911testcompany| company name |
 |Merchant City|60|up to 15| Singapore|6009Singapore| stative value for singapore|
 |Additional Data Field Template| 62| up to 99| |62270123testordernumber12345678| additional data attached|
 |CRC | 63 | 04 | |63040047| CRC16 sign, sign string includes 6304|
 
 **Merchant Account Infomation Desciption**
 
 | Name|ID|length|Example|Comment|
 | ----| ---- | ---- | ---- | ---- |
 | Globally Unique Identifier | 00 | up to 32 | 0009SG.PAYNOW| static value|
 | Account Type | 01 | 01 | 01012 | value 0 for mobile, value 2 for UEN |
 | UEN | 02 | up to 13 |  020812345678 | UEN number takes 10, suffix takes 3 |
 | Payment Amount Editable | 03 | 01 | 1 | value 1 payment is editable, value 0 is not editable |
 | Expiry Date | 04 | 08 | 040820260304 | expire date format YYYYMMDD| 
 
 **Additional Data Field Description**
 
 | Name|ID|length|Example|Comment|
 | ----| ---- | ---- | ---- | ---- |
 |Bill Number | 01 | up to 25 | 0123testordernumber12345678 | reference order number/ID|


# Credits
Original code reference from: https://gist.github.com/chengkiang/7e1c4899768245570cc49c7d23bc394c

