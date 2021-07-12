package module

import (
	"fmt"
	"strconv"
	"github.com/addonrizky/qrmpmparse/constant"
	"github.com/addonrizky/qrmpmparse/library"
	"github.com/addonrizky/qrmpmparse/entity"
	"github.com/snksoft/crc"
	"encoding/json"
)

func ExtractTagQR(QRString string) (entity.ExtractionCode, string, entity.RootID, string){
	currentIndex := 0
	tagMap := make(map[string]interface{})
	tagLabel := ""
	tagLengthString := ""
	tagLength := 0
	endIndex := 0
	tagValue := ""
	tagSpecial62 := ""
	var err error
	var tagObject entity.RootID
	var extractionCode entity.ExtractionCode

	for currentIndex < len(QRString){
		tagLabel = QRString[currentIndex:currentIndex+2]
		tagLengthString = QRString[currentIndex+2:currentIndex+4]
		tagLength, err = strconv.Atoi(tagLengthString)
		endIndex = currentIndex+4+tagLength
		tagValue = QRString[currentIndex+4:endIndex]

		if err != nil {
			extractionCode = library.GetExtractionCode("E0", "error parsing taglength")
			return extractionCode,"{}",tagObject,tagSpecial62
		}
		currentIndex = endIndex

		isHavingSubtag := library.StringInSlice(tagLabel, constant.GetRootTagHavingSubtag())

		if tagLabel == "62" {
			tagSpecial62 = tagValue
		}

		var subtag map[string]string
		if isHavingSubtag {
			subtag = library.GetSubtag(tagValue)
			tagMap[tagLabel] = subtag
		} else {
			tagMap[tagLabel] = tagValue
		}
	}

	tagMapByte, _ := json.Marshal(tagMap)
	err = json.Unmarshal(tagMapByte, &tagObject)

	//check existence of tag 02-51, merchant account informations
	isMerchantAccountInfoExist := false
	isMerchantAccountInfoExist = library.IsMerchantAccountInfoExist(tagMapByte)

	if err != nil {
		extractionCode = library.GetExtractionCode("E1", "invalid tag element")
		return extractionCode,"{}",tagObject,tagSpecial62
	}

	calculatedCRC := crc.CalculateCRC(crc.CCITT, []byte(QRString[:len(QRString)-4]))
	crc := fmt.Sprintf("%X", calculatedCRC)
	crc = library.Lpad(crc,"0",4)
	if crc != tagObject.Tag63 {
		extractionCode = library.GetExtractionCode("E2", "invalid CRC")
		return extractionCode,"{}",tagObject,tagSpecial62
	}

	if isMerchantAccountInfoExist == false {
		extractionCode = library.GetExtractionCode("E3", "no existence tag of merchant account info tag02-51")
		return extractionCode,"{}",tagObject,tagSpecial62
	}
	
	extractionCode = library.GetExtractionCode("00", "extraction success")
	return extractionCode,string(tagMapByte),tagObject,tagSpecial62
}