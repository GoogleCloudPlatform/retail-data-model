// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.10
// source: api/enums/uom.proto

package enums

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Area int32

const (
	Area_AREA_OTHER    Area = 0
	Area_SQUARE_INCH   Area = 11
	Area_ACRE          Area = 12
	Area_HECTARE       Area = 13
	Area_SQR_CM        Area = 14
	Area_SQR_FOOT      Area = 15
	Area_SQR_METER     Area = 16
	Area_SQR_MILE      Area = 17
	Area_SQR_ROD       Area = 18
	Area_SQR_YARD      Area = 19
	Area_SQR_KILOMETER Area = 20
)

// Enum value maps for Area.
var (
	Area_name = map[int32]string{
		0:  "AREA_OTHER",
		11: "SQUARE_INCH",
		12: "ACRE",
		13: "HECTARE",
		14: "SQR_CM",
		15: "SQR_FOOT",
		16: "SQR_METER",
		17: "SQR_MILE",
		18: "SQR_ROD",
		19: "SQR_YARD",
		20: "SQR_KILOMETER",
	}
	Area_value = map[string]int32{
		"AREA_OTHER":    0,
		"SQUARE_INCH":   11,
		"ACRE":          12,
		"HECTARE":       13,
		"SQR_CM":        14,
		"SQR_FOOT":      15,
		"SQR_METER":     16,
		"SQR_MILE":      17,
		"SQR_ROD":       18,
		"SQR_YARD":      19,
		"SQR_KILOMETER": 20,
	}
)

func (x Area) Enum() *Area {
	p := new(Area)
	*p = x
	return p
}

func (x Area) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Area) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_uom_proto_enumTypes[0].Descriptor()
}

func (Area) Type() protoreflect.EnumType {
	return &file_api_enums_uom_proto_enumTypes[0]
}

func (x Area) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Area.Descriptor instead.
func (Area) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_uom_proto_rawDescGZIP(), []int{0}
}

type Capacity int32

const (
	Capacity_CAPACITY_OTHER   Capacity = 0
	Capacity_CUBIC_INCH       Capacity = 101
	Capacity_ACRE_FOOT        Capacity = 102
	Capacity_BOARD_FOOT       Capacity = 103
	Capacity_BUSHEL           Capacity = 104
	Capacity_CENTILITER       Capacity = 105
	Capacity_CORD             Capacity = 106
	Capacity_CUBIC_CENTIMETER Capacity = 107
	Capacity_CUBIC_FOOT       Capacity = 108
	Capacity_CUBIC_YARD       Capacity = 109
	Capacity_DECALITER        Capacity = 110
	Capacity_DECILITER        Capacity = 111
	Capacity_FLUID_DRAM       Capacity = 112
	Capacity_FLUID_OUNCE      Capacity = 113
	Capacity_GALLON           Capacity = 114
	Capacity_GILL             Capacity = 115
	Capacity_HECTOLITER       Capacity = 116
	Capacity_KILOLITER        Capacity = 117
	Capacity_LITER            Capacity = 118
	Capacity_MILLLITER        Capacity = 119
	Capacity_MINIM            Capacity = 120
	Capacity_PECK             Capacity = 121
	Capacity_PINT             Capacity = 122
	Capacity_QUART            Capacity = 123
)

// Enum value maps for Capacity.
var (
	Capacity_name = map[int32]string{
		0:   "CAPACITY_OTHER",
		101: "CUBIC_INCH",
		102: "ACRE_FOOT",
		103: "BOARD_FOOT",
		104: "BUSHEL",
		105: "CENTILITER",
		106: "CORD",
		107: "CUBIC_CENTIMETER",
		108: "CUBIC_FOOT",
		109: "CUBIC_YARD",
		110: "DECALITER",
		111: "DECILITER",
		112: "FLUID_DRAM",
		113: "FLUID_OUNCE",
		114: "GALLON",
		115: "GILL",
		116: "HECTOLITER",
		117: "KILOLITER",
		118: "LITER",
		119: "MILLLITER",
		120: "MINIM",
		121: "PECK",
		122: "PINT",
		123: "QUART",
	}
	Capacity_value = map[string]int32{
		"CAPACITY_OTHER":   0,
		"CUBIC_INCH":       101,
		"ACRE_FOOT":        102,
		"BOARD_FOOT":       103,
		"BUSHEL":           104,
		"CENTILITER":       105,
		"CORD":             106,
		"CUBIC_CENTIMETER": 107,
		"CUBIC_FOOT":       108,
		"CUBIC_YARD":       109,
		"DECALITER":        110,
		"DECILITER":        111,
		"FLUID_DRAM":       112,
		"FLUID_OUNCE":      113,
		"GALLON":           114,
		"GILL":             115,
		"HECTOLITER":       116,
		"KILOLITER":        117,
		"LITER":            118,
		"MILLLITER":        119,
		"MINIM":            120,
		"PECK":             121,
		"PINT":             122,
		"QUART":            123,
	}
)

func (x Capacity) Enum() *Capacity {
	p := new(Capacity)
	*p = x
	return p
}

func (x Capacity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Capacity) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_uom_proto_enumTypes[1].Descriptor()
}

func (Capacity) Type() protoreflect.EnumType {
	return &file_api_enums_uom_proto_enumTypes[1]
}

func (x Capacity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Capacity.Descriptor instead.
func (Capacity) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_uom_proto_rawDescGZIP(), []int{1}
}

type Count int32

const (
	Count_COUNT_OTHER Count = 0
	Count_EACH        Count = 201
	Count_BAG         Count = 202
	Count_BOWL        Count = 203
	Count_BOX         Count = 204
	Count_BUCKET      Count = 205
	Count_BUNDLE      Count = 206
	Count_CARTON      Count = 207
	Count_CASE        Count = 208
	Count_DOZEN       Count = 209
	Count_KIT         Count = 210
	Count_LOT         Count = 211
	Count_PACK        Count = 212
	Count_PAIR        Count = 213
	Count_PEICE       Count = 214
	Count_PERSON      Count = 215
	Count_PLATE       Count = 216
	Count_ROLL        Count = 217
	Count_SET_2       Count = 218
	Count_SET_3       Count = 219
	Count_SET_4       Count = 220
	Count_SET_5       Count = 221
	Count_SET_6       Count = 222
	Count_SET_7       Count = 223
	Count_SET_8       Count = 224
	Count_SET_9       Count = 225
	Count_SET_10      Count = 226
	Count_SHEET       Count = 227
	Count_SINGLE      Count = 228
	Count_TUBE        Count = 229
)

// Enum value maps for Count.
var (
	Count_name = map[int32]string{
		0:   "COUNT_OTHER",
		201: "EACH",
		202: "BAG",
		203: "BOWL",
		204: "BOX",
		205: "BUCKET",
		206: "BUNDLE",
		207: "CARTON",
		208: "CASE",
		209: "DOZEN",
		210: "KIT",
		211: "LOT",
		212: "PACK",
		213: "PAIR",
		214: "PEICE",
		215: "PERSON",
		216: "PLATE",
		217: "ROLL",
		218: "SET_2",
		219: "SET_3",
		220: "SET_4",
		221: "SET_5",
		222: "SET_6",
		223: "SET_7",
		224: "SET_8",
		225: "SET_9",
		226: "SET_10",
		227: "SHEET",
		228: "SINGLE",
		229: "TUBE",
	}
	Count_value = map[string]int32{
		"COUNT_OTHER": 0,
		"EACH":        201,
		"BAG":         202,
		"BOWL":        203,
		"BOX":         204,
		"BUCKET":      205,
		"BUNDLE":      206,
		"CARTON":      207,
		"CASE":        208,
		"DOZEN":       209,
		"KIT":         210,
		"LOT":         211,
		"PACK":        212,
		"PAIR":        213,
		"PEICE":       214,
		"PERSON":      215,
		"PLATE":       216,
		"ROLL":        217,
		"SET_2":       218,
		"SET_3":       219,
		"SET_4":       220,
		"SET_5":       221,
		"SET_6":       222,
		"SET_7":       223,
		"SET_8":       224,
		"SET_9":       225,
		"SET_10":      226,
		"SHEET":       227,
		"SINGLE":      228,
		"TUBE":        229,
	}
)

func (x Count) Enum() *Count {
	p := new(Count)
	*p = x
	return p
}

func (x Count) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Count) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_uom_proto_enumTypes[2].Descriptor()
}

func (Count) Type() protoreflect.EnumType {
	return &file_api_enums_uom_proto_enumTypes[2]
}

func (x Count) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Count.Descriptor instead.
func (Count) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_uom_proto_rawDescGZIP(), []int{2}
}

type Distance int32

const (
	Distance_DISTANCE_OTHER Distance = 0
	Distance_INCH           Distance = 301
	Distance_CENTIMETER     Distance = 302
	Distance_DECAMETER      Distance = 303
	Distance_DECIMETER      Distance = 304
	Distance_FATHOM         Distance = 305
	Distance_FOOT           Distance = 306
	Distance_FURLONG        Distance = 307
	Distance_HECTOMETER     Distance = 308
	Distance_KILOMETER      Distance = 309
	Distance_METER          Distance = 310
	Distance_MILE           Distance = 311
	Distance_MILLIMETER     Distance = 312
	Distance_NAUTICAL_MILE  Distance = 313
	Distance_ROD            Distance = 314
	Distance_YARD           Distance = 315
)

// Enum value maps for Distance.
var (
	Distance_name = map[int32]string{
		0:   "DISTANCE_OTHER",
		301: "INCH",
		302: "CENTIMETER",
		303: "DECAMETER",
		304: "DECIMETER",
		305: "FATHOM",
		306: "FOOT",
		307: "FURLONG",
		308: "HECTOMETER",
		309: "KILOMETER",
		310: "METER",
		311: "MILE",
		312: "MILLIMETER",
		313: "NAUTICAL_MILE",
		314: "ROD",
		315: "YARD",
	}
	Distance_value = map[string]int32{
		"DISTANCE_OTHER": 0,
		"INCH":           301,
		"CENTIMETER":     302,
		"DECAMETER":      303,
		"DECIMETER":      304,
		"FATHOM":         305,
		"FOOT":           306,
		"FURLONG":        307,
		"HECTOMETER":     308,
		"KILOMETER":      309,
		"METER":          310,
		"MILE":           311,
		"MILLIMETER":     312,
		"NAUTICAL_MILE":  313,
		"ROD":            314,
		"YARD":           315,
	}
)

func (x Distance) Enum() *Distance {
	p := new(Distance)
	*p = x
	return p
}

func (x Distance) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Distance) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_uom_proto_enumTypes[3].Descriptor()
}

func (Distance) Type() protoreflect.EnumType {
	return &file_api_enums_uom_proto_enumTypes[3]
}

func (x Distance) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Distance.Descriptor instead.
func (Distance) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_uom_proto_rawDescGZIP(), []int{3}
}

type Screen int32

const (
	Screen_SCREEN_OTHER Screen = 0
	Screen_PX           Screen = 401
	Screen_PT           Screen = 402
	Screen_EM           Screen = 403
	Screen_PCT          Screen = 404
	Screen_IN           Screen = 405
)

// Enum value maps for Screen.
var (
	Screen_name = map[int32]string{
		0:   "SCREEN_OTHER",
		401: "PX",
		402: "PT",
		403: "EM",
		404: "PCT",
		405: "IN",
	}
	Screen_value = map[string]int32{
		"SCREEN_OTHER": 0,
		"PX":           401,
		"PT":           402,
		"EM":           403,
		"PCT":          404,
		"IN":           405,
	}
)

func (x Screen) Enum() *Screen {
	p := new(Screen)
	*p = x
	return p
}

func (x Screen) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Screen) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_uom_proto_enumTypes[4].Descriptor()
}

func (Screen) Type() protoreflect.EnumType {
	return &file_api_enums_uom_proto_enumTypes[4]
}

func (x Screen) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Screen.Descriptor instead.
func (Screen) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_uom_proto_rawDescGZIP(), []int{4}
}

type Weight int32

const (
	Weight_WEIGHT_OTHER         Weight = 0
	Weight_OUNCE                Weight = 501
	Weight_CENTIGRAM            Weight = 502
	Weight_DECAGRAM             Weight = 503
	Weight_DECIGRAM             Weight = 504
	Weight_DRAM                 Weight = 505
	Weight_DRAM_APOTHECARY      Weight = 506
	Weight_GRAIN                Weight = 507
	Weight_GRAIN_APOTHECARY     Weight = 508
	Weight_GRAIN_TROY           Weight = 509
	Weight_GRAM                 Weight = 510
	Weight_HECTOGRAM            Weight = 511
	Weight_HUNDRED_WEIGHT_LONG  Weight = 512
	Weight_HUNDRED_WEIGHT_SHORT Weight = 513
	Weight_KILOGRAM             Weight = 514
	Weight_MILLIGRAM            Weight = 515
	Weight_OUNCE_APOTHECARY     Weight = 516
	Weight_OUNCE_TROY           Weight = 517
	Weight_PENNY_WEIGHT_TROY    Weight = 518
	Weight_POUND                Weight = 519
	Weight_POUND_APOTHECARY     Weight = 520
	Weight_POUND_TROY           Weight = 521
	Weight_SCRUPLE_APOTHECARY   Weight = 522
	Weight_STONE                Weight = 523
	Weight_TON                  Weight = 524
	Weight_TON_LONG             Weight = 525
	Weight_TON_SHORT            Weight = 526
	Weight_TONNE                Weight = 527
)

// Enum value maps for Weight.
var (
	Weight_name = map[int32]string{
		0:   "WEIGHT_OTHER",
		501: "OUNCE",
		502: "CENTIGRAM",
		503: "DECAGRAM",
		504: "DECIGRAM",
		505: "DRAM",
		506: "DRAM_APOTHECARY",
		507: "GRAIN",
		508: "GRAIN_APOTHECARY",
		509: "GRAIN_TROY",
		510: "GRAM",
		511: "HECTOGRAM",
		512: "HUNDRED_WEIGHT_LONG",
		513: "HUNDRED_WEIGHT_SHORT",
		514: "KILOGRAM",
		515: "MILLIGRAM",
		516: "OUNCE_APOTHECARY",
		517: "OUNCE_TROY",
		518: "PENNY_WEIGHT_TROY",
		519: "POUND",
		520: "POUND_APOTHECARY",
		521: "POUND_TROY",
		522: "SCRUPLE_APOTHECARY",
		523: "STONE",
		524: "TON",
		525: "TON_LONG",
		526: "TON_SHORT",
		527: "TONNE",
	}
	Weight_value = map[string]int32{
		"WEIGHT_OTHER":         0,
		"OUNCE":                501,
		"CENTIGRAM":            502,
		"DECAGRAM":             503,
		"DECIGRAM":             504,
		"DRAM":                 505,
		"DRAM_APOTHECARY":      506,
		"GRAIN":                507,
		"GRAIN_APOTHECARY":     508,
		"GRAIN_TROY":           509,
		"GRAM":                 510,
		"HECTOGRAM":            511,
		"HUNDRED_WEIGHT_LONG":  512,
		"HUNDRED_WEIGHT_SHORT": 513,
		"KILOGRAM":             514,
		"MILLIGRAM":            515,
		"OUNCE_APOTHECARY":     516,
		"OUNCE_TROY":           517,
		"PENNY_WEIGHT_TROY":    518,
		"POUND":                519,
		"POUND_APOTHECARY":     520,
		"POUND_TROY":           521,
		"SCRUPLE_APOTHECARY":   522,
		"STONE":                523,
		"TON":                  524,
		"TON_LONG":             525,
		"TON_SHORT":            526,
		"TONNE":                527,
	}
)

func (x Weight) Enum() *Weight {
	p := new(Weight)
	*p = x
	return p
}

func (x Weight) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Weight) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_uom_proto_enumTypes[5].Descriptor()
}

func (Weight) Type() protoreflect.EnumType {
	return &file_api_enums_uom_proto_enumTypes[5]
}

func (x Weight) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Weight.Descriptor instead.
func (Weight) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_uom_proto_rawDescGZIP(), []int{5}
}

type Packaging int32

const (
	Packaging_PACKAGE_OTHER    Packaging = 0
	Packaging_PACKAGE_BOX      Packaging = 600
	Packaging_PACKAGE_CYLINDER Packaging = 601
	Packaging_PACKAGE_ENVELOPE Packaging = 602
)

// Enum value maps for Packaging.
var (
	Packaging_name = map[int32]string{
		0:   "PACKAGE_OTHER",
		600: "PACKAGE_BOX",
		601: "PACKAGE_CYLINDER",
		602: "PACKAGE_ENVELOPE",
	}
	Packaging_value = map[string]int32{
		"PACKAGE_OTHER":    0,
		"PACKAGE_BOX":      600,
		"PACKAGE_CYLINDER": 601,
		"PACKAGE_ENVELOPE": 602,
	}
)

func (x Packaging) Enum() *Packaging {
	p := new(Packaging)
	*p = x
	return p
}

func (x Packaging) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Packaging) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_uom_proto_enumTypes[6].Descriptor()
}

func (Packaging) Type() protoreflect.EnumType {
	return &file_api_enums_uom_proto_enumTypes[6]
}

func (x Packaging) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Packaging.Descriptor instead.
func (Packaging) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_uom_proto_rawDescGZIP(), []int{6}
}

type Time int32

const (
	Time_TIME_OTHER Time = 0
	Time_SECONDS    Time = 701
	Time_MINUTES    Time = 702
	Time_HOURS      Time = 703
	Time_DAYS       Time = 704
	Time_WEEKS      Time = 705
	Time_MONTHS     Time = 706
	Time_YEARS      Time = 707
)

// Enum value maps for Time.
var (
	Time_name = map[int32]string{
		0:   "TIME_OTHER",
		701: "SECONDS",
		702: "MINUTES",
		703: "HOURS",
		704: "DAYS",
		705: "WEEKS",
		706: "MONTHS",
		707: "YEARS",
	}
	Time_value = map[string]int32{
		"TIME_OTHER": 0,
		"SECONDS":    701,
		"MINUTES":    702,
		"HOURS":      703,
		"DAYS":       704,
		"WEEKS":      705,
		"MONTHS":     706,
		"YEARS":      707,
	}
)

func (x Time) Enum() *Time {
	p := new(Time)
	*p = x
	return p
}

func (x Time) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Time) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_uom_proto_enumTypes[7].Descriptor()
}

func (Time) Type() protoreflect.EnumType {
	return &file_api_enums_uom_proto_enumTypes[7]
}

func (x Time) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Time.Descriptor instead.
func (Time) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_uom_proto_rawDescGZIP(), []int{7}
}

type Clothing int32

const (
	Clothing_XXS    Clothing = 0
	Clothing_XS     Clothing = 1
	Clothing_S      Clothing = 2
	Clothing_M      Clothing = 3
	Clothing_XL     Clothing = 4
	Clothing_XXL    Clothing = 5
	Clothing_XXXL   Clothing = 6
	Clothing_XXXXL  Clothing = 7
	Clothing_XXXXXL Clothing = 8
)

// Enum value maps for Clothing.
var (
	Clothing_name = map[int32]string{
		0: "XXS",
		1: "XS",
		2: "S",
		3: "M",
		4: "XL",
		5: "XXL",
		6: "XXXL",
		7: "XXXXL",
		8: "XXXXXL",
	}
	Clothing_value = map[string]int32{
		"XXS":    0,
		"XS":     1,
		"S":      2,
		"M":      3,
		"XL":     4,
		"XXL":    5,
		"XXXL":   6,
		"XXXXL":  7,
		"XXXXXL": 8,
	}
)

func (x Clothing) Enum() *Clothing {
	p := new(Clothing)
	*p = x
	return p
}

func (x Clothing) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Clothing) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_uom_proto_enumTypes[8].Descriptor()
}

func (Clothing) Type() protoreflect.EnumType {
	return &file_api_enums_uom_proto_enumTypes[8]
}

func (x Clothing) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Clothing.Descriptor instead.
func (Clothing) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_uom_proto_rawDescGZIP(), []int{8}
}

var File_api_enums_uom_proto protoreflect.FileDescriptor

var file_api_enums_uom_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x75, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x75, 0x6f, 0x6d, 0x2a, 0xa3,
	0x01, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x52, 0x45, 0x41, 0x5f,
	0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x51, 0x55, 0x41, 0x52,
	0x45, 0x5f, 0x49, 0x4e, 0x43, 0x48, 0x10, 0x0b, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x43, 0x52, 0x45,
	0x10, 0x0c, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45, 0x43, 0x54, 0x41, 0x52, 0x45, 0x10, 0x0d, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x51, 0x52, 0x5f, 0x43, 0x4d, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x51, 0x52, 0x5f, 0x46, 0x4f, 0x4f, 0x54, 0x10, 0x0f, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x51, 0x52,
	0x5f, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x51, 0x52, 0x5f,
	0x4d, 0x49, 0x4c, 0x45, 0x10, 0x11, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x51, 0x52, 0x5f, 0x52, 0x4f,
	0x44, 0x10, 0x12, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x51, 0x52, 0x5f, 0x59, 0x41, 0x52, 0x44, 0x10,
	0x13, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x51, 0x52, 0x5f, 0x4b, 0x49, 0x4c, 0x4f, 0x4d, 0x45, 0x54,
	0x45, 0x52, 0x10, 0x14, 0x2a, 0xe1, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x50, 0x41, 0x43, 0x49, 0x54, 0x59, 0x5f, 0x4f, 0x54,
	0x48, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x55, 0x42, 0x49, 0x43, 0x5f, 0x49,
	0x4e, 0x43, 0x48, 0x10, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x43, 0x52, 0x45, 0x5f, 0x46, 0x4f,
	0x4f, 0x54, 0x10, 0x66, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x5f, 0x46, 0x4f,
	0x4f, 0x54, 0x10, 0x67, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x55, 0x53, 0x48, 0x45, 0x4c, 0x10, 0x68,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x45, 0x4e, 0x54, 0x49, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x10, 0x69,
	0x12, 0x08, 0x0a, 0x04, 0x43, 0x4f, 0x52, 0x44, 0x10, 0x6a, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x55,
	0x42, 0x49, 0x43, 0x5f, 0x43, 0x45, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x6b,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x55, 0x42, 0x49, 0x43, 0x5f, 0x46, 0x4f, 0x4f, 0x54, 0x10, 0x6c,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x55, 0x42, 0x49, 0x43, 0x5f, 0x59, 0x41, 0x52, 0x44, 0x10, 0x6d,
	0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x43, 0x41, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x10, 0x6e, 0x12,
	0x0d, 0x0a, 0x09, 0x44, 0x45, 0x43, 0x49, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x10, 0x6f, 0x12, 0x0e,
	0x0a, 0x0a, 0x46, 0x4c, 0x55, 0x49, 0x44, 0x5f, 0x44, 0x52, 0x41, 0x4d, 0x10, 0x70, 0x12, 0x0f,
	0x0a, 0x0b, 0x46, 0x4c, 0x55, 0x49, 0x44, 0x5f, 0x4f, 0x55, 0x4e, 0x43, 0x45, 0x10, 0x71, 0x12,
	0x0a, 0x0a, 0x06, 0x47, 0x41, 0x4c, 0x4c, 0x4f, 0x4e, 0x10, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x47,
	0x49, 0x4c, 0x4c, 0x10, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x45, 0x43, 0x54, 0x4f, 0x4c, 0x49,
	0x54, 0x45, 0x52, 0x10, 0x74, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x49, 0x4c, 0x4f, 0x4c, 0x49, 0x54,
	0x45, 0x52, 0x10, 0x75, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x10, 0x76, 0x12,
	0x0d, 0x0a, 0x09, 0x4d, 0x49, 0x4c, 0x4c, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x10, 0x77, 0x12, 0x09,
	0x0a, 0x05, 0x4d, 0x49, 0x4e, 0x49, 0x4d, 0x10, 0x78, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x45, 0x43,
	0x4b, 0x10, 0x79, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x54, 0x10, 0x7a, 0x12, 0x09, 0x0a,
	0x05, 0x51, 0x55, 0x41, 0x52, 0x54, 0x10, 0x7b, 0x2a, 0xeb, 0x02, 0x0a, 0x05, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4f, 0x54, 0x48, 0x45,
	0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x04, 0x45, 0x41, 0x43, 0x48, 0x10, 0xc9, 0x01, 0x12, 0x08,
	0x0a, 0x03, 0x42, 0x41, 0x47, 0x10, 0xca, 0x01, 0x12, 0x09, 0x0a, 0x04, 0x42, 0x4f, 0x57, 0x4c,
	0x10, 0xcb, 0x01, 0x12, 0x08, 0x0a, 0x03, 0x42, 0x4f, 0x58, 0x10, 0xcc, 0x01, 0x12, 0x0b, 0x0a,
	0x06, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x10, 0xcd, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x42, 0x55,
	0x4e, 0x44, 0x4c, 0x45, 0x10, 0xce, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x41, 0x52, 0x54, 0x4f,
	0x4e, 0x10, 0xcf, 0x01, 0x12, 0x09, 0x0a, 0x04, 0x43, 0x41, 0x53, 0x45, 0x10, 0xd0, 0x01, 0x12,
	0x0a, 0x0a, 0x05, 0x44, 0x4f, 0x5a, 0x45, 0x4e, 0x10, 0xd1, 0x01, 0x12, 0x08, 0x0a, 0x03, 0x4b,
	0x49, 0x54, 0x10, 0xd2, 0x01, 0x12, 0x08, 0x0a, 0x03, 0x4c, 0x4f, 0x54, 0x10, 0xd3, 0x01, 0x12,
	0x09, 0x0a, 0x04, 0x50, 0x41, 0x43, 0x4b, 0x10, 0xd4, 0x01, 0x12, 0x09, 0x0a, 0x04, 0x50, 0x41,
	0x49, 0x52, 0x10, 0xd5, 0x01, 0x12, 0x0a, 0x0a, 0x05, 0x50, 0x45, 0x49, 0x43, 0x45, 0x10, 0xd6,
	0x01, 0x12, 0x0b, 0x0a, 0x06, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x10, 0xd7, 0x01, 0x12, 0x0a,
	0x0a, 0x05, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x10, 0xd8, 0x01, 0x12, 0x09, 0x0a, 0x04, 0x52, 0x4f,
	0x4c, 0x4c, 0x10, 0xd9, 0x01, 0x12, 0x0a, 0x0a, 0x05, 0x53, 0x45, 0x54, 0x5f, 0x32, 0x10, 0xda,
	0x01, 0x12, 0x0a, 0x0a, 0x05, 0x53, 0x45, 0x54, 0x5f, 0x33, 0x10, 0xdb, 0x01, 0x12, 0x0a, 0x0a,
	0x05, 0x53, 0x45, 0x54, 0x5f, 0x34, 0x10, 0xdc, 0x01, 0x12, 0x0a, 0x0a, 0x05, 0x53, 0x45, 0x54,
	0x5f, 0x35, 0x10, 0xdd, 0x01, 0x12, 0x0a, 0x0a, 0x05, 0x53, 0x45, 0x54, 0x5f, 0x36, 0x10, 0xde,
	0x01, 0x12, 0x0a, 0x0a, 0x05, 0x53, 0x45, 0x54, 0x5f, 0x37, 0x10, 0xdf, 0x01, 0x12, 0x0a, 0x0a,
	0x05, 0x53, 0x45, 0x54, 0x5f, 0x38, 0x10, 0xe0, 0x01, 0x12, 0x0a, 0x0a, 0x05, 0x53, 0x45, 0x54,
	0x5f, 0x39, 0x10, 0xe1, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x53, 0x45, 0x54, 0x5f, 0x31, 0x30, 0x10,
	0xe2, 0x01, 0x12, 0x0a, 0x0a, 0x05, 0x53, 0x48, 0x45, 0x45, 0x54, 0x10, 0xe3, 0x01, 0x12, 0x0b,
	0x0a, 0x06, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x10, 0xe4, 0x01, 0x12, 0x09, 0x0a, 0x04, 0x54,
	0x55, 0x42, 0x45, 0x10, 0xe5, 0x01, 0x2a, 0xf2, 0x01, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x49, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f,
	0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x04, 0x49, 0x4e, 0x43, 0x48, 0x10,
	0xad, 0x02, 0x12, 0x0f, 0x0a, 0x0a, 0x43, 0x45, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x54, 0x45, 0x52,
	0x10, 0xae, 0x02, 0x12, 0x0e, 0x0a, 0x09, 0x44, 0x45, 0x43, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52,
	0x10, 0xaf, 0x02, 0x12, 0x0e, 0x0a, 0x09, 0x44, 0x45, 0x43, 0x49, 0x4d, 0x45, 0x54, 0x45, 0x52,
	0x10, 0xb0, 0x02, 0x12, 0x0b, 0x0a, 0x06, 0x46, 0x41, 0x54, 0x48, 0x4f, 0x4d, 0x10, 0xb1, 0x02,
	0x12, 0x09, 0x0a, 0x04, 0x46, 0x4f, 0x4f, 0x54, 0x10, 0xb2, 0x02, 0x12, 0x0c, 0x0a, 0x07, 0x46,
	0x55, 0x52, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0xb3, 0x02, 0x12, 0x0f, 0x0a, 0x0a, 0x48, 0x45, 0x43,
	0x54, 0x4f, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0xb4, 0x02, 0x12, 0x0e, 0x0a, 0x09, 0x4b, 0x49,
	0x4c, 0x4f, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0xb5, 0x02, 0x12, 0x0a, 0x0a, 0x05, 0x4d, 0x45,
	0x54, 0x45, 0x52, 0x10, 0xb6, 0x02, 0x12, 0x09, 0x0a, 0x04, 0x4d, 0x49, 0x4c, 0x45, 0x10, 0xb7,
	0x02, 0x12, 0x0f, 0x0a, 0x0a, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10,
	0xb8, 0x02, 0x12, 0x12, 0x0a, 0x0d, 0x4e, 0x41, 0x55, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x4d,
	0x49, 0x4c, 0x45, 0x10, 0xb9, 0x02, 0x12, 0x08, 0x0a, 0x03, 0x52, 0x4f, 0x44, 0x10, 0xba, 0x02,
	0x12, 0x09, 0x0a, 0x04, 0x59, 0x41, 0x52, 0x44, 0x10, 0xbb, 0x02, 0x2a, 0x48, 0x0a, 0x06, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x43, 0x52, 0x45, 0x45, 0x4e, 0x5f,
	0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x02, 0x50, 0x58, 0x10, 0x91, 0x03,
	0x12, 0x07, 0x0a, 0x02, 0x50, 0x54, 0x10, 0x92, 0x03, 0x12, 0x07, 0x0a, 0x02, 0x45, 0x4d, 0x10,
	0x93, 0x03, 0x12, 0x08, 0x0a, 0x03, 0x50, 0x43, 0x54, 0x10, 0x94, 0x03, 0x12, 0x07, 0x0a, 0x02,
	0x49, 0x4e, 0x10, 0x95, 0x03, 0x2a, 0xe6, 0x03, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x10, 0x0a, 0x0c, 0x57, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x05, 0x4f, 0x55, 0x4e, 0x43, 0x45, 0x10, 0xf5, 0x03, 0x12, 0x0e,
	0x0a, 0x09, 0x43, 0x45, 0x4e, 0x54, 0x49, 0x47, 0x52, 0x41, 0x4d, 0x10, 0xf6, 0x03, 0x12, 0x0d,
	0x0a, 0x08, 0x44, 0x45, 0x43, 0x41, 0x47, 0x52, 0x41, 0x4d, 0x10, 0xf7, 0x03, 0x12, 0x0d, 0x0a,
	0x08, 0x44, 0x45, 0x43, 0x49, 0x47, 0x52, 0x41, 0x4d, 0x10, 0xf8, 0x03, 0x12, 0x09, 0x0a, 0x04,
	0x44, 0x52, 0x41, 0x4d, 0x10, 0xf9, 0x03, 0x12, 0x14, 0x0a, 0x0f, 0x44, 0x52, 0x41, 0x4d, 0x5f,
	0x41, 0x50, 0x4f, 0x54, 0x48, 0x45, 0x43, 0x41, 0x52, 0x59, 0x10, 0xfa, 0x03, 0x12, 0x0a, 0x0a,
	0x05, 0x47, 0x52, 0x41, 0x49, 0x4e, 0x10, 0xfb, 0x03, 0x12, 0x15, 0x0a, 0x10, 0x47, 0x52, 0x41,
	0x49, 0x4e, 0x5f, 0x41, 0x50, 0x4f, 0x54, 0x48, 0x45, 0x43, 0x41, 0x52, 0x59, 0x10, 0xfc, 0x03,
	0x12, 0x0f, 0x0a, 0x0a, 0x47, 0x52, 0x41, 0x49, 0x4e, 0x5f, 0x54, 0x52, 0x4f, 0x59, 0x10, 0xfd,
	0x03, 0x12, 0x09, 0x0a, 0x04, 0x47, 0x52, 0x41, 0x4d, 0x10, 0xfe, 0x03, 0x12, 0x0e, 0x0a, 0x09,
	0x48, 0x45, 0x43, 0x54, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x10, 0xff, 0x03, 0x12, 0x18, 0x0a, 0x13,
	0x48, 0x55, 0x4e, 0x44, 0x52, 0x45, 0x44, 0x5f, 0x57, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x4c,
	0x4f, 0x4e, 0x47, 0x10, 0x80, 0x04, 0x12, 0x19, 0x0a, 0x14, 0x48, 0x55, 0x4e, 0x44, 0x52, 0x45,
	0x44, 0x5f, 0x57, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x10, 0x81,
	0x04, 0x12, 0x0d, 0x0a, 0x08, 0x4b, 0x49, 0x4c, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x82, 0x04,
	0x12, 0x0e, 0x0a, 0x09, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x83, 0x04,
	0x12, 0x15, 0x0a, 0x10, 0x4f, 0x55, 0x4e, 0x43, 0x45, 0x5f, 0x41, 0x50, 0x4f, 0x54, 0x48, 0x45,
	0x43, 0x41, 0x52, 0x59, 0x10, 0x84, 0x04, 0x12, 0x0f, 0x0a, 0x0a, 0x4f, 0x55, 0x4e, 0x43, 0x45,
	0x5f, 0x54, 0x52, 0x4f, 0x59, 0x10, 0x85, 0x04, 0x12, 0x16, 0x0a, 0x11, 0x50, 0x45, 0x4e, 0x4e,
	0x59, 0x5f, 0x57, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x54, 0x52, 0x4f, 0x59, 0x10, 0x86, 0x04,
	0x12, 0x0a, 0x0a, 0x05, 0x50, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x87, 0x04, 0x12, 0x15, 0x0a, 0x10,
	0x50, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x41, 0x50, 0x4f, 0x54, 0x48, 0x45, 0x43, 0x41, 0x52, 0x59,
	0x10, 0x88, 0x04, 0x12, 0x0f, 0x0a, 0x0a, 0x50, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x54, 0x52, 0x4f,
	0x59, 0x10, 0x89, 0x04, 0x12, 0x17, 0x0a, 0x12, 0x53, 0x43, 0x52, 0x55, 0x50, 0x4c, 0x45, 0x5f,
	0x41, 0x50, 0x4f, 0x54, 0x48, 0x45, 0x43, 0x41, 0x52, 0x59, 0x10, 0x8a, 0x04, 0x12, 0x0a, 0x0a,
	0x05, 0x53, 0x54, 0x4f, 0x4e, 0x45, 0x10, 0x8b, 0x04, 0x12, 0x08, 0x0a, 0x03, 0x54, 0x4f, 0x4e,
	0x10, 0x8c, 0x04, 0x12, 0x0d, 0x0a, 0x08, 0x54, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x10,
	0x8d, 0x04, 0x12, 0x0e, 0x0a, 0x09, 0x54, 0x4f, 0x4e, 0x5f, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x10,
	0x8e, 0x04, 0x12, 0x0a, 0x0a, 0x05, 0x54, 0x4f, 0x4e, 0x4e, 0x45, 0x10, 0x8f, 0x04, 0x2a, 0x5e,
	0x0a, 0x09, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x0a, 0x0d, 0x50,
	0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0b, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x4f, 0x58, 0x10, 0xd8, 0x04,
	0x12, 0x15, 0x0a, 0x10, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x59, 0x4c, 0x49,
	0x4e, 0x44, 0x45, 0x52, 0x10, 0xd9, 0x04, 0x12, 0x15, 0x0a, 0x10, 0x50, 0x41, 0x43, 0x4b, 0x41,
	0x47, 0x45, 0x5f, 0x45, 0x4e, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x45, 0x10, 0xda, 0x04, 0x2a, 0x6e,
	0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x4f,
	0x54, 0x48, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x07, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44,
	0x53, 0x10, 0xbd, 0x05, 0x12, 0x0c, 0x0a, 0x07, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x10,
	0xbe, 0x05, 0x12, 0x0a, 0x0a, 0x05, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x10, 0xbf, 0x05, 0x12, 0x09,
	0x0a, 0x04, 0x44, 0x41, 0x59, 0x53, 0x10, 0xc0, 0x05, 0x12, 0x0a, 0x0a, 0x05, 0x57, 0x45, 0x45,
	0x4b, 0x53, 0x10, 0xc1, 0x05, 0x12, 0x0b, 0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x53, 0x10,
	0xc2, 0x05, 0x12, 0x0a, 0x0a, 0x05, 0x59, 0x45, 0x41, 0x52, 0x53, 0x10, 0xc3, 0x05, 0x2a, 0x5b,
	0x0a, 0x08, 0x43, 0x6c, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x58,
	0x53, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x58, 0x53, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x53,
	0x10, 0x02, 0x12, 0x05, 0x0a, 0x01, 0x4d, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x58, 0x4c, 0x10,
	0x04, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x58, 0x4c, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x58, 0x58,
	0x58, 0x4c, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x58, 0x58, 0x58, 0x58, 0x4c, 0x10, 0x07, 0x12,
	0x0a, 0x0a, 0x06, 0x58, 0x58, 0x58, 0x58, 0x58, 0x4c, 0x10, 0x08, 0x42, 0x53, 0x0a, 0x17, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2d,
	0x64, 0x61, 0x74, 0x61, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_enums_uom_proto_rawDescOnce sync.Once
	file_api_enums_uom_proto_rawDescData = file_api_enums_uom_proto_rawDesc
)

func file_api_enums_uom_proto_rawDescGZIP() []byte {
	file_api_enums_uom_proto_rawDescOnce.Do(func() {
		file_api_enums_uom_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_enums_uom_proto_rawDescData)
	})
	return file_api_enums_uom_proto_rawDescData
}

var file_api_enums_uom_proto_enumTypes = make([]protoimpl.EnumInfo, 9)
var file_api_enums_uom_proto_goTypes = []interface{}{
	(Area)(0),      // 0: google.retail.enums.uom.Area
	(Capacity)(0),  // 1: google.retail.enums.uom.Capacity
	(Count)(0),     // 2: google.retail.enums.uom.Count
	(Distance)(0),  // 3: google.retail.enums.uom.Distance
	(Screen)(0),    // 4: google.retail.enums.uom.Screen
	(Weight)(0),    // 5: google.retail.enums.uom.Weight
	(Packaging)(0), // 6: google.retail.enums.uom.Packaging
	(Time)(0),      // 7: google.retail.enums.uom.Time
	(Clothing)(0),  // 8: google.retail.enums.uom.Clothing
}
var file_api_enums_uom_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_enums_uom_proto_init() }
func file_api_enums_uom_proto_init() {
	if File_api_enums_uom_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_enums_uom_proto_rawDesc,
			NumEnums:      9,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_enums_uom_proto_goTypes,
		DependencyIndexes: file_api_enums_uom_proto_depIdxs,
		EnumInfos:         file_api_enums_uom_proto_enumTypes,
	}.Build()
	File_api_enums_uom_proto = out.File
	file_api_enums_uom_proto_rawDesc = nil
	file_api_enums_uom_proto_goTypes = nil
	file_api_enums_uom_proto_depIdxs = nil
}
