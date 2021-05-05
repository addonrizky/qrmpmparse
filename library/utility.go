package library 

import (
	"strconv"
	"fmt"
	"strings"
	"github.com/addonrizky/qrmpmparse/entity"
)

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func GetSubtag(stringTagValue string) map[string]string {
	currentIndex := 0
	tagMap := make(map[string]string)
	tagLabel := ""
	tagLengthString := ""
	tagLength := 0
	endIndex := 0
	tagValue := ""
	var err error

	for currentIndex < len(stringTagValue){
		tagLabel = stringTagValue[currentIndex:currentIndex+2]
		tagLengthString = stringTagValue[currentIndex+2:currentIndex+4]
		tagLength, err = strconv.Atoi(tagLengthString)
		endIndex = currentIndex+4+tagLength
		tagValue = stringTagValue[currentIndex+4:endIndex]

		if err != nil {
			fmt.Println(err)
			return nil
		}
		currentIndex = endIndex

		tagMap[tagLabel] = strings.ToUpper(tagValue)
		
	}

	return tagMap
}

//GetExtractionResult, retrieve result extraction
func GetExtractionCode(code string, desc string) entity.ExtractionCode{
	extractionCode := entity.ExtractionCode{
		Code: code,
		Desc: desc,
	}
	return extractionCode
}