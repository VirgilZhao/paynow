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
	ID        string
	Name      string
	MaxLength int
	Value     interface{}
}

func (obj DataObject) getString() string {
	valueStr := ""
	switch obj.Value.(type) {
	case string:
		valueStr = obj.Value.(string)
		if len(valueStr) > obj.MaxLength {
			panic(fmt.Sprintf("id %s value %s is out of max length %d", obj.ID, valueStr, obj.MaxLength))
		}
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
		ID:        "00",
		Name:      "Payload Format Indicator",
		MaxLength: 2,
		Value:     "01",
	}
	pointOfInitiationMethod := DataObject{
		ID:        "01",
		Name:      "Point of Initiation Method",
		MaxLength: 2,
		Value:     "12",
	}
	payNowIndicator := DataObject{
		ID:        "00",
		Name:      "PayNow Indicator",
		MaxLength: 32,
		Value:     "SG.PAYNOW",
	}
	mobileOrUenAccount := DataObject{
		ID:        "01",
		Name:      "Mobile Or UEN Account",
		MaxLength: 1,
		Value:     "2",
	}
	uenAccount := DataObject{
		ID:        "02",
		MaxLength: 13,
		Name:      "UEN Account Number",
		Value:     options.UEN,
	}
	editable := DataObject{
		ID:        "03",
		MaxLength: 1,
		Name:      "payment amount editable",
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
			ID:        "04",
			Name:      "Expiry Date",
			MaxLength: 8,
			Value:     options.Expiry,
		}
		merchantAccountInfoTemplateValues = append(merchantAccountInfoTemplateValues, expiry)
	}
	merchantAccountInfoTemplate := DataObject{
		ID:        "26",
		Name:      "Merchant Account Info Template",
		MaxLength: 99,
		Value:     merchantAccountInfoTemplateValues,
	}
	merchantCategoryCode := DataObject{
		ID:        "52",
		Name:      "Merchant Category Code",
		MaxLength: 4,
		Value:     "0000",
	}
	currency := DataObject{
		ID:        "53",
		Name:      "Currency",
		MaxLength: 3,
		Value:     "702",
	}
	transactionAmount := DataObject{
		ID:        "54",
		Name:      "Transaction Amount",
		MaxLength: 13,
		Value:     options.Amount,
	}
	countryCode := DataObject{
		ID:        "58",
		Name:      "Country Code",
		MaxLength: 2,
		Value:     "SG",
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
			ID:        "59",
			Name:      "Company Name",
			MaxLength: 25,
			Value:     options.CompanyName,
		}
		objects = append(objects, companyName)
	}
	merchantCity := DataObject{
		ID:        "60",
		Name:      "Merchant City",
		MaxLength: 15,
		Value:     "Singapore",
	}
	objects = append(objects, merchantCity)
	referenceNumber := DataObject{
		ID:        "01",
		Name:      "Bill/Reference Number",
		MaxLength: 25,
		Value:     options.ReferenceNumber,
	}
	additionalDataFields := DataObject{
		ID:        "62",
		Name:      "Additional Data Fields",
		MaxLength: 99,
		Value:     []DataObject{referenceNumber},
	}
	objects = append(objects, additionalDataFields)
	crc := DataObject{
		ID:        "63",
		Name:      "CRC",
		MaxLength: 4,
		Value:     "",
	}
	objects = append(objects, crc)
	rootObject := &RootObject{DataObjects: objects}
	return rootObject
}

func GeneratePayNowString(options Options) string {
	rootObject := getPayNowDataObject(options)
	return rootObject.getString()
}
