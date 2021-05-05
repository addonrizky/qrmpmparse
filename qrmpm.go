package main

import (
	"fmt"
	"github.com/addonrizky/qrmpmparse/module"
	"github.com/addonrizky/qrmpmparse/entity"
)

func main() {
	var tagObject entity.RootID
	extractionCode, tags, tagObject:= module.ExtractTagQR("00020101021226650013ID.CO.BRI.WWW011893600002011067409602150000010033400000303UME52045021530336054045000550202560410005802ID5920TOKO PRABHAS GEMBIRA6013JAKARTA BARAT6105171446222011800169100000185578963041516")
	
	fmt.Println(extractionCode)
	fmt.Println(tagObject.Tag60)
	
	fmt.Println(tags)
}