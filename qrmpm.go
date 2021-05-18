package main

import (
	"fmt"
	"github.com/addonrizky/qrmpmparse/module"
	"github.com/addonrizky/qrmpmparse/entity"
)

func main() {
	var tagObject entity.RootID
	extractionCode, tags, tagObject, tagSpecial62:= module.ExtractTagQR("00020101021226690017ID.CO.GANESHA.WWW011893600161970729023802157100024722912310303UMI51420014ID.CO.QRIS.WWW0213ID710002472290303UMI5204481253033605405300005502015802ID5919Warung Nenek Hafizh6013Tangerang Sel610529723622007169473872618926109630491A1")
	
	fmt.Println(extractionCode)
	fmt.Println(tagObject.Tag62)
	fmt.Println(tagSpecial62)
	fmt.Println(tags)
	//fmt.Println(len(tagObject))
}