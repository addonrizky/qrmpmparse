package library 

import (
	"strconv"
	"fmt"
	"github.com/addonrizky/qrmpmparse/entity"
	"github.com/addonrizky/qrmpmparse/constant"
	"encoding/json"
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

		tagMap[tagLabel] = tagValue
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

func IsMerchantAccountInfoExist(tagMapByte []byte) bool {
	isMerchantAccountInfoExist := false
	c := make(map[string]json.RawMessage)
    err := json.Unmarshal(tagMapByte, &c)
    if err != nil {
        return isMerchantAccountInfoExist
    }
	k := make([]string, len(c))
	i := 0
	for s, _ := range c {
		k[i] = s
		if isMerchantAccountInfoExist == false {
			isMerchantAccountInfoExist = StringInSlice(s, constant.GetMerchantAccountInfoTag())
		}
		i++;
	}
    
	return isMerchantAccountInfoExist
}