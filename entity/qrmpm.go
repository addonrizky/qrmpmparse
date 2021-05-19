package entity

type RootID struct {
	Tag00 string `json:"00,omitempty"`
	Tag01 string `json:"01,omitempty"`
	Tag02 string `json:"02,omitempty"`
	Tag03 string `json:"03,omitempty"`
	Tag04 string `json:"04,omitempty"`
	Tag05 string `json:"05,omitempty"`
	Tag06 string `json:"06,omitempty"`
	Tag07 string `json:"07,omitempty"`
	Tag08 string `json:"08,omitempty"`
	Tag09 string `json:"09,omitempty"`
	Tag10 string `json:"10,omitempty"`
	Tag11 string `json:"11,omitempty"`
	Tag12 string `json:"12,omitempty"`
	Tag13 string `json:"13,omitempty"`
	Tag14 string `json:"14,omitempty"`
	Tag15 string `json:"15,omitempty"`
	Tag16 string `json:"16,omitempty"`
	Tag17 string `json:"17,omitempty"`
	Tag18 string `json:"18,omitempty"`
	Tag19 string `json:"19,omitempty"`
	Tag20 string `json:"20,omitempty"`
	Tag21 string `json:"21,omitempty"`
	Tag22 string `json:"22,omitempty"`
	Tag23 string `json:"23,omitempty"`
	Tag24 string `json:"24,omitempty"`
	Tag25 string `json:"25,omitempty"`
	Tag26 SubtagMerchantAccountInformation `json:"26,omitempty"`
	Tag27 SubtagMerchantAccountInformation `json:"27,omitempty"`
	Tag28 SubtagMerchantAccountInformation `json:"28,omitempty"`
	Tag29 SubtagMerchantAccountInformation `json:"29,omitempty"`
	Tag30 SubtagMerchantAccountInformation `json:"30,omitempty"`
	Tag31 SubtagMerchantAccountInformation `json:"31,omitempty"`
	Tag32 SubtagMerchantAccountInformation `json:"32,omitempty"`
	Tag33 SubtagMerchantAccountInformation `json:"33,omitempty"`
	Tag34 SubtagMerchantAccountInformation `json:"34,omitempty"`
	Tag35 SubtagMerchantAccountInformation `json:"35,omitempty"`
	Tag36 SubtagMerchantAccountInformation `json:"36,omitempty"`
	Tag37 SubtagMerchantAccountInformation `json:"37,omitempty"`
	Tag38 SubtagMerchantAccountInformation `json:"38,omitempty"`
	Tag39 SubtagMerchantAccountInformation `json:"39,omitempty"`
	Tag40 SubtagMerchantAccountInformation `json:"40,omitempty"`
	Tag41 SubtagMerchantAccountInformation `json:"41,omitempty"`
	Tag42 SubtagMerchantAccountInformation `json:"42,omitempty"`
	Tag43 SubtagMerchantAccountInformation `json:"43,omitempty"`
	Tag44 SubtagMerchantAccountInformation `json:"44,omitempty"`
	Tag45 SubtagMerchantAccountInformation `json:"45,omitempty"`
	Tag46 SubtagMerchantAccountInformation `json:"46,omitempty"`
	Tag47 SubtagMerchantAccountInformation `json:"47,omitempty"`
	Tag48 SubtagMerchantAccountInformation `json:"48,omitempty"`
	Tag49 SubtagMerchantAccountInformation `json:"49,omitempty"`
	Tag50 SubtagMerchantAccountInformation `json:"50,omitempty"`
	Tag51 SubtagMerchantAccountInformation `json:"51,omitempty"`
	Tag52 string `json:"52,omitempty"`
	Tag53 string `json:"53,omitempty"`
	Tag54 string `json:"54,omitempty"`
	Tag55 string `json:"55,omitempty"`
	Tag56 string `json:"56,omitempty"`
	Tag57 string `json:"57,omitempty"`
	Tag58 string `json:"58,omitempty"`
	Tag59 string `json:"59,omitempty"`
	Tag60 string `json:"60,omitempty"`
	Tag61 string `json:"61,omitempty"`
	Tag62 SubtagAdditionalData `json:"62,ompitempty"`
	Tag63 string `json:"63,omitempty"`
	Tag64 string `json:"64,omitempty"`
}

type SubtagMerchantAccountInformation struct {
	Tag00 string `json:"00"`
	Tag01 string `json:"01"`
	Tag02 string `json:"02"`
	Tag03 string `json:"03"`
}

type SubtagAdditionalData struct {
	Tag01 string `json:"01"`
	Tag02 string `json:"02"`
	Tag03 string `json:"03"`
	Tag04 string `json:"04"`
	Tag05 string `json:"05"`
	Tag06 string `json:"06"`
	Tag07 string `json:"07"`
	Tag08 string `json:"08"`
	Tag09 string `json:"09"`
	Tag10 string `json:"10"`
	//----------------------
	Tag11 string `json:"11"`
	Tag12 string `json:"12"`
	Tag13 string `json:"13"`
	Tag14 string `json:"14"`
	Tag15 string `json:"15"`
	Tag16 string `json:"16"`
	Tag17 string `json:"17"`
	Tag18 string `json:"18"`
	Tag19 string `json:"19"`
	Tag20 string `json:"20"`
	//----------------------
	Tag21 string `json:"21"`
	Tag22 string `json:"22"`
	Tag23 string `json:"23"`
	Tag24 string `json:"24"`
	Tag25 string `json:"25"`
	Tag26 string `json:"26"`
	Tag27 string `json:"27"`
	Tag28 string `json:"28"`
	Tag29 string `json:"29"`
	Tag30 string `json:"30"`
	//----------------------
	Tag31 string `json:"31"`
	Tag32 string `json:"32"`
	Tag33 string `json:"33"`
	Tag34 string `json:"34"`
	Tag35 string `json:"35"`
	Tag36 string `json:"36"`
	Tag37 string `json:"37"`
	Tag38 string `json:"38"`
	Tag39 string `json:"39"`
	Tag40 string `json:"40"`
	//---------------------
	Tag41 string `json:"41"`
	Tag42 string `json:"42"`
	Tag43 string `json:"43"`
	Tag44 string `json:"44"`
	Tag45 string `json:"45"`
	Tag46 string `json:"46"`
	Tag47 string `json:"47"`
	Tag48 string `json:"48"`
	Tag49 string `json:"49"`
	Tag50 string `json:"50"`
	//----------------------
	Tag51 string `json:"51"`
	Tag52 string `json:"52"`
	Tag53 string `json:"53"`
	Tag54 string `json:"54"`
	Tag55 string `json:"55"`
	Tag56 string `json:"56"`
	Tag57 string `json:"57"`
	Tag58 string `json:"58"`
	Tag59 string `json:"59"`
	Tag60 string `json:"60"`
	//----------------------
	Tag61 string `json:"61"`
	Tag62 string `json:"62"`
	Tag63 string `json:"63"`
	Tag64 string `json:"64"`
	Tag65 string `json:"65"`
	Tag66 string `json:"66"`
	Tag67 string `json:"67"`
	Tag68 string `json:"68"`
	Tag69 string `json:"69"`
	Tag70 string `json:"70"`
	//-----------------------
	Tag71 string `json:"71"`
	Tag72 string `json:"72"`
	Tag73 string `json:"73"`
	Tag74 string `json:"74"`
	Tag75 string `json:"75"`
	Tag76 string `json:"76"`
	Tag77 string `json:"77"`
	Tag78 string `json:"78"`
	Tag79 string `json:"79"`
	Tag80 string `json:"80"`
	//-----------------------
	Tag81 string `json:"81"`
	Tag82 string `json:"82"`
	Tag83 string `json:"83"`
	Tag84 string `json:"84"`
	Tag85 string `json:"85"`
	Tag86 string `json:"86"`
	Tag87 string `json:"87"`
	Tag88 string `json:"88"`
	Tag89 string `json:"89"`
	Tag90 string `json:"90"`
	//----------------------
	Tag91 string `json:"91"`
	Tag92 string `json:"92"`
	Tag93 string `json:"93"`
	Tag94 string `json:"94"`
	Tag95 string `json:"95"`
	Tag96 string `json:"96"`
	Tag97 string `json:"97"`
	Tag98 string `json:"98"`
	Tag99 string `json:"99"`
}

