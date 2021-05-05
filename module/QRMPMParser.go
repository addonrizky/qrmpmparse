package module

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/addonrizky/qrmpmparse/constant"
	"github.com/addonrizky/qrmpmparse/library"
	"github.com/addonrizky/qrmpmparse/entity"
	"github.com/snksoft/crc"
	"encoding/json"
)

func ExtractTagQR(QRString string) (entity.ExtractionCode, string, entity.RootID){
	currentIndex := 0
	tagMap := make(map[string]interface{})
	tagLabel := ""
	tagLengthString := ""
	tagLength := 0
	endIndex := 0
	tagValue := ""
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
			return extractionCode,"{}",tagObject
		}
		currentIndex = endIndex

		isHavingSubtag := library.StringInSlice(tagLabel, constant.GetRootTagHavingSubtag())

		var subtag map[string]string
		if isHavingSubtag {
			subtag = library.GetSubtag(tagValue)
			tagMap[tagLabel] = subtag
		} else {
			tagMap[tagLabel] = strings.ToUpper(tagValue)
		}
	}
	tagMapByte, _ := json.Marshal(tagMap)
	err = json.Unmarshal(tagMapByte, &tagObject)
	
	if err != nil {
		extractionCode = library.GetExtractionCode("E1", "invalid tag element")
		return extractionCode,"{}",tagObject
	}

	calculatedCRC := crc.CalculateCRC(crc.CCITT, []byte(QRString[:len(QRString)-4]))
	crc := fmt.Sprintf("%X", calculatedCRC)
	if crc != tagObject.Tag63 {
		extractionCode = library.GetExtractionCode("E2", "invalid CRC")
		return extractionCode,"{}",tagObject
	}
	

	extractionCode = library.GetExtractionCode("00", "extraction success")
	return extractionCode,string(tagMapByte),tagObject
}
