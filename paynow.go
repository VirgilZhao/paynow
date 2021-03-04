package paynow

import "fmt"

type Options struct {
	UEN             string
	Editable        bool
	Expiry          string
	CompanyName     string
	Amount          string
	ReferenceNumber string
}

type RootObject struct {
	DataObjects []DataObject
}

type DataObject struct {
	ID    string
	Name  string
	Value interface{}
}

func (obj DataObject) getString() string {
	valueStr := ""
	switch obj.Value.(type) {
	case string:
		valueStr = obj.Value.(string)
		break
	case []DataObject:
		subObjects := obj.Value.([]DataObject)
		for _, subObj := range subObjects {
			valueStr += subObj.getString()
		}
		break
	}
	lengthStr := fmt.Sprintf("%02d", len(valueStr))
	return obj.ID + lengthStr + valueStr
}

func (obj DataObject) getCRCString(value string) string {
	value += obj.ID + "04" // CRC require length "04"
	data := []byte(value)
	checkSum := crc16(data)
	valueStr := fmt.Sprintf("%04X", checkSum)
	return obj.ID + "04" + valueStr
}

func (rt *RootObject) getString() string {
	valueStr := ""
	for _, dataObject := range rt.DataObjects {
		if dataObject.ID == "63" {
			valueStr += dataObject.getCRCString(valueStr)
		} else {
			valueStr += dataObject.getString()
		}
	}
	return valueStr
}

func getPayNowDataObject(options Options) *RootObject {
	payloadFormatIndicator := DataObject{
		ID:    "00",
		Name:  "Payload Format Indicator",
		Value: "01",
	}
	pointOfInitiationMethod := DataObject{
		ID:    "01",
		Name:  "Point of Initiation Method",
		Value: "12",
	}
	payNowIndicator := DataObject{
		ID:    "00",
		Name:  "PayNow Indicator",
		Value: "SG.PAYNOW",
	}
	mobileOrUenAccount := DataObject{
		ID:    "01",
		Name:  "Mobile Or UEN Account",
		Value: "2",
	}
	uenAccount := DataObject{
		ID:    "02",
		Name:  "UEN Account Number",
		Value: options.UEN,
	}
	editable := DataObject{
		ID:   "03",
		Name: "payment amount editable",
	}
	if options.Editable {
		editable.Value = "1"
	} else {
		editable.Value = "0"
	}
	merchantAccountInfoTemplateValues := []DataObject{
		payNowIndicator,
		mobileOrUenAccount,
		uenAccount,
		editable,
	}
	if len(options.Expiry) == 8 {
		expiry := DataObject{
			ID:    "04",
			Name:  "Expiry Date",
			Value: options.Expiry,
		}
		merchantAccountInfoTemplateValues = append(merchantAccountInfoTemplateValues, expiry)
	}
	merchantAccountInfoTemplate := DataObject{
		ID:    "26",
		Name:  "Merchant Account Info Template",
		Value: merchantAccountInfoTemplateValues,
	}
	merchantCategoryCode := DataObject{
		ID:    "52",
		Name:  "Merchant Category Code",
		Value: "0000",
	}
	currency := DataObject{
		ID:    "53",
		Name:  "Currency",
		Value: "702",
	}
	transactionAmount := DataObject{
		ID:    "54",
		Name:  "Transaction Amount",
		Value: options.Amount,
	}
	countryCode := DataObject{
		ID:    "58",
		Name:  "Country Code",
		Value: "SG",
	}
	objects := []DataObject{
		payloadFormatIndicator,
		pointOfInitiationMethod,
		merchantAccountInfoTemplate,
		merchantCategoryCode,
		currency,
		transactionAmount,
		countryCode,
	}
	if len(options.CompanyName) > 0 {
		companyName := DataObject{
			ID:    "59",
			Name:  "Company Name",
			Value: options.CompanyName,
		}
		objects = append(objects, companyName)
	}
	merchantCity := DataObject{
		ID:    "60",
		Name:  "Merchant City",
		Value: "Singapore",
	}
	objects = append(objects, merchantCity)
	referenceNumber := DataObject{
		ID:    "01",
		Name:  "Bill/Reference Number",
		Value: options.ReferenceNumber,
	}
	addiontialDataFields := DataObject{
		ID:    "62",
		Name:  "Additional Data Fields",
		Value: []DataObject{referenceNumber},
	}
	objects = append(objects, addiontialDataFields)
	crc := DataObject{
		ID:    "63",
		Name:  "CRC",
		Value: "",
	}
	objects = append(objects, crc)
	rootObject := &RootObject{DataObjects: objects}
	return rootObject
}

func GeneratePayNowString(options Options) string {
	rootObject := getPayNowDataObject(options)
	return rootObject.getString()
}
