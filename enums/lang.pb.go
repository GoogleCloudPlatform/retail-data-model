// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.10
// source: api/enums/lang.proto

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

type Language int32

const (
	Language_UNKNOWN Language = 0
	Language_AA      Language = 2
	Language_AF      Language = 3
	Language_AK      Language = 4
	Language_SQ      Language = 5
	Language_AM      Language = 6
	Language_AR      Language = 7
	Language_AN      Language = 8
	Language_HY      Language = 9
	Language_AS      Language = 10
	Language_AE      Language = 12
	Language_AY      Language = 13
	Language_BM      Language = 15
	Language_BA      Language = 16
	Language_BE      Language = 18
	Language_BN      Language = 19
	Language_BI      Language = 20
	Language_BS      Language = 21
	Language_BR      Language = 22
	Language_BG      Language = 23
	Language_MY      Language = 24
	Language_CH      Language = 26
	Language_CE      Language = 27
	Language_CV      Language = 30
	Language_KW      Language = 31
	Language_CR      Language = 33
	Language_HR      Language = 34
	Language_DA      Language = 36
	Language_DZ      Language = 39
	Language_EN      Language = 40
	Language_EO      Language = 41
	Language_EE      Language = 43
	Language_FO      Language = 44
	Language_FJ      Language = 45
	Language_FR      Language = 47
	Language_GL      Language = 49
	Language_KA      Language = 50
	Language_DE      Language = 51
	Language_GN      Language = 53
	Language_GU      Language = 54
	Language_HA      Language = 56
	Language_HE      Language = 57
	Language_HZ      Language = 58
	Language_HO      Language = 60
	Language_HU      Language = 61
	Language_IA      Language = 62
	Language_ID      Language = 63
	Language_GA      Language = 65
	Language_IG      Language = 66
	Language_IO      Language = 68
	Language_IS      Language = 69
	Language_IT      Language = 70
	Language_IU      Language = 71
	Language_JA      Language = 72
	Language_KN      Language = 75
	Language_KR      Language = 76
	Language_KK      Language = 78
	Language_RW      Language = 81
	Language_KV      Language = 83
	Language_KG      Language = 84
	Language_KO      Language = 85
	Language_LG      Language = 90
	Language_LN      Language = 92
	Language_LO      Language = 93
	Language_LT      Language = 94
	Language_LU      Language = 95
	Language_LV      Language = 96
	Language_MK      Language = 98
	Language_MG      Language = 99
	Language_ML      Language = 101
	Language_MT      Language = 102
	Language_MI      Language = 103
	Language_MR      Language = 104
	Language_MH      Language = 105
	Language_MN      Language = 106
	Language_NA      Language = 107
	Language_ND      Language = 109
	Language_NE      Language = 110
	Language_NG      Language = 111
	Language_NB      Language = 112
	Language_NN      Language = 113
	Language_NO      Language = 114
	Language_NR      Language = 116
	Language_OJ      Language = 118
	Language_OM      Language = 120
	Language_OR      Language = 121
	Language_FA      Language = 125
	Language_PT      Language = 128
	Language_RM      Language = 130
	Language_RN      Language = 131
	Language_RU      Language = 133
	Language_SC      Language = 135
	Language_SE      Language = 137
	Language_SM      Language = 138
	Language_SG      Language = 139
	Language_SR      Language = 140
	Language_SN      Language = 142
	Language_ST      Language = 147
	Language_SU      Language = 149
	Language_SW      Language = 150
	Language_SS      Language = 151
	Language_SV      Language = 152
	Language_TA      Language = 153
	Language_TE      Language = 154
	Language_TH      Language = 156
	Language_TI      Language = 157
	Language_BO      Language = 158
	Language_TL      Language = 160
	Language_TN      Language = 161
	Language_TO      Language = 162
	Language_TR      Language = 163
	Language_TS      Language = 164
	Language_TW      Language = 166
	Language_TY      Language = 167
	Language_UK      Language = 169
	Language_UR      Language = 170
	Language_VE      Language = 172
	Language_VI      Language = 173
	Language_VO      Language = 174
	Language_WA      Language = 175
	Language_CY      Language = 176
	Language_WO      Language = 177
	Language_FY      Language = 178
	Language_XH      Language = 179
	Language_YI      Language = 180
	Language_YO      Language = 181
	Language_ZU      Language = 183
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0:   "UNKNOWN",
		2:   "AA",
		3:   "AF",
		4:   "AK",
		5:   "SQ",
		6:   "AM",
		7:   "AR",
		8:   "AN",
		9:   "HY",
		10:  "AS",
		12:  "AE",
		13:  "AY",
		15:  "BM",
		16:  "BA",
		18:  "BE",
		19:  "BN",
		20:  "BI",
		21:  "BS",
		22:  "BR",
		23:  "BG",
		24:  "MY",
		26:  "CH",
		27:  "CE",
		30:  "CV",
		31:  "KW",
		33:  "CR",
		34:  "HR",
		36:  "DA",
		39:  "DZ",
		40:  "EN",
		41:  "EO",
		43:  "EE",
		44:  "FO",
		45:  "FJ",
		47:  "FR",
		49:  "GL",
		50:  "KA",
		51:  "DE",
		53:  "GN",
		54:  "GU",
		56:  "HA",
		57:  "HE",
		58:  "HZ",
		60:  "HO",
		61:  "HU",
		62:  "IA",
		63:  "ID",
		65:  "GA",
		66:  "IG",
		68:  "IO",
		69:  "IS",
		70:  "IT",
		71:  "IU",
		72:  "JA",
		75:  "KN",
		76:  "KR",
		78:  "KK",
		81:  "RW",
		83:  "KV",
		84:  "KG",
		85:  "KO",
		90:  "LG",
		92:  "LN",
		93:  "LO",
		94:  "LT",
		95:  "LU",
		96:  "LV",
		98:  "MK",
		99:  "MG",
		101: "ML",
		102: "MT",
		103: "MI",
		104: "MR",
		105: "MH",
		106: "MN",
		107: "NA",
		109: "ND",
		110: "NE",
		111: "NG",
		112: "NB",
		113: "NN",
		114: "NO",
		116: "NR",
		118: "OJ",
		120: "OM",
		121: "OR",
		125: "FA",
		128: "PT",
		130: "RM",
		131: "RN",
		133: "RU",
		135: "SC",
		137: "SE",
		138: "SM",
		139: "SG",
		140: "SR",
		142: "SN",
		147: "ST",
		149: "SU",
		150: "SW",
		151: "SS",
		152: "SV",
		153: "TA",
		154: "TE",
		156: "TH",
		157: "TI",
		158: "BO",
		160: "TL",
		161: "TN",
		162: "TO",
		163: "TR",
		164: "TS",
		166: "TW",
		167: "TY",
		169: "UK",
		170: "UR",
		172: "VE",
		173: "VI",
		174: "VO",
		175: "WA",
		176: "CY",
		177: "WO",
		178: "FY",
		179: "XH",
		180: "YI",
		181: "YO",
		183: "ZU",
	}
	Language_value = map[string]int32{
		"UNKNOWN": 0,
		"AA":      2,
		"AF":      3,
		"AK":      4,
		"SQ":      5,
		"AM":      6,
		"AR":      7,
		"AN":      8,
		"HY":      9,
		"AS":      10,
		"AE":      12,
		"AY":      13,
		"BM":      15,
		"BA":      16,
		"BE":      18,
		"BN":      19,
		"BI":      20,
		"BS":      21,
		"BR":      22,
		"BG":      23,
		"MY":      24,
		"CH":      26,
		"CE":      27,
		"CV":      30,
		"KW":      31,
		"CR":      33,
		"HR":      34,
		"DA":      36,
		"DZ":      39,
		"EN":      40,
		"EO":      41,
		"EE":      43,
		"FO":      44,
		"FJ":      45,
		"FR":      47,
		"GL":      49,
		"KA":      50,
		"DE":      51,
		"GN":      53,
		"GU":      54,
		"HA":      56,
		"HE":      57,
		"HZ":      58,
		"HO":      60,
		"HU":      61,
		"IA":      62,
		"ID":      63,
		"GA":      65,
		"IG":      66,
		"IO":      68,
		"IS":      69,
		"IT":      70,
		"IU":      71,
		"JA":      72,
		"KN":      75,
		"KR":      76,
		"KK":      78,
		"RW":      81,
		"KV":      83,
		"KG":      84,
		"KO":      85,
		"LG":      90,
		"LN":      92,
		"LO":      93,
		"LT":      94,
		"LU":      95,
		"LV":      96,
		"MK":      98,
		"MG":      99,
		"ML":      101,
		"MT":      102,
		"MI":      103,
		"MR":      104,
		"MH":      105,
		"MN":      106,
		"NA":      107,
		"ND":      109,
		"NE":      110,
		"NG":      111,
		"NB":      112,
		"NN":      113,
		"NO":      114,
		"NR":      116,
		"OJ":      118,
		"OM":      120,
		"OR":      121,
		"FA":      125,
		"PT":      128,
		"RM":      130,
		"RN":      131,
		"RU":      133,
		"SC":      135,
		"SE":      137,
		"SM":      138,
		"SG":      139,
		"SR":      140,
		"SN":      142,
		"ST":      147,
		"SU":      149,
		"SW":      150,
		"SS":      151,
		"SV":      152,
		"TA":      153,
		"TE":      154,
		"TH":      156,
		"TI":      157,
		"BO":      158,
		"TL":      160,
		"TN":      161,
		"TO":      162,
		"TR":      163,
		"TS":      164,
		"TW":      166,
		"TY":      167,
		"UK":      169,
		"UR":      170,
		"VE":      172,
		"VI":      173,
		"VO":      174,
		"WA":      175,
		"CY":      176,
		"WO":      177,
		"FY":      178,
		"XH":      179,
		"YI":      180,
		"YO":      181,
		"ZU":      183,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_lang_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_api_enums_lang_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_lang_proto_rawDescGZIP(), []int{0}
}

var File_api_enums_lang_proto protoreflect.FileDescriptor

var file_api_enums_lang_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x6c, 0x61, 0x6e, 0x67,
	0x2a, 0xaf, 0x08, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x41,
	0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x46, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x4b,
	0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x51, 0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x4d,
	0x10, 0x06, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x52, 0x10, 0x07, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x4e,
	0x10, 0x08, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x59, 0x10, 0x09, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x53,
	0x10, 0x0a, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x45, 0x10, 0x0c, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x59,
	0x10, 0x0d, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x4d, 0x10, 0x0f, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x41,
	0x10, 0x10, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x45, 0x10, 0x12, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x4e,
	0x10, 0x13, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x49, 0x10, 0x14, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x53,
	0x10, 0x15, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x52, 0x10, 0x16, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x47,
	0x10, 0x17, 0x12, 0x06, 0x0a, 0x02, 0x4d, 0x59, 0x10, 0x18, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x48,
	0x10, 0x1a, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x45, 0x10, 0x1b, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x56,
	0x10, 0x1e, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x57, 0x10, 0x1f, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x52,
	0x10, 0x21, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x52, 0x10, 0x22, 0x12, 0x06, 0x0a, 0x02, 0x44, 0x41,
	0x10, 0x24, 0x12, 0x06, 0x0a, 0x02, 0x44, 0x5a, 0x10, 0x27, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x4e,
	0x10, 0x28, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x4f, 0x10, 0x29, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x45,
	0x10, 0x2b, 0x12, 0x06, 0x0a, 0x02, 0x46, 0x4f, 0x10, 0x2c, 0x12, 0x06, 0x0a, 0x02, 0x46, 0x4a,
	0x10, 0x2d, 0x12, 0x06, 0x0a, 0x02, 0x46, 0x52, 0x10, 0x2f, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x4c,
	0x10, 0x31, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x41, 0x10, 0x32, 0x12, 0x06, 0x0a, 0x02, 0x44, 0x45,
	0x10, 0x33, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x4e, 0x10, 0x35, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x55,
	0x10, 0x36, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x41, 0x10, 0x38, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x45,
	0x10, 0x39, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x5a, 0x10, 0x3a, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x4f,
	0x10, 0x3c, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x55, 0x10, 0x3d, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x41,
	0x10, 0x3e, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x44, 0x10, 0x3f, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x41,
	0x10, 0x41, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x47, 0x10, 0x42, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x4f,
	0x10, 0x44, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x53, 0x10, 0x45, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x54,
	0x10, 0x46, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x55, 0x10, 0x47, 0x12, 0x06, 0x0a, 0x02, 0x4a, 0x41,
	0x10, 0x48, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x4e, 0x10, 0x4b, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x52,
	0x10, 0x4c, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x4b, 0x10, 0x4e, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x57,
	0x10, 0x51, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x56, 0x10, 0x53, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x47,
	0x10, 0x54, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x4f, 0x10, 0x55, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x47,
	0x10, 0x5a, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x4e, 0x10, 0x5c, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x4f,
	0x10, 0x5d, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x54, 0x10, 0x5e, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x55,
	0x10, 0x5f, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x56, 0x10, 0x60, 0x12, 0x06, 0x0a, 0x02, 0x4d, 0x4b,
	0x10, 0x62, 0x12, 0x06, 0x0a, 0x02, 0x4d, 0x47, 0x10, 0x63, 0x12, 0x06, 0x0a, 0x02, 0x4d, 0x4c,
	0x10, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4d, 0x54, 0x10, 0x66, 0x12, 0x06, 0x0a, 0x02, 0x4d, 0x49,
	0x10, 0x67, 0x12, 0x06, 0x0a, 0x02, 0x4d, 0x52, 0x10, 0x68, 0x12, 0x06, 0x0a, 0x02, 0x4d, 0x48,
	0x10, 0x69, 0x12, 0x06, 0x0a, 0x02, 0x4d, 0x4e, 0x10, 0x6a, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x41,
	0x10, 0x6b, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x44, 0x10, 0x6d, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x45,
	0x10, 0x6e, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x47, 0x10, 0x6f, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x42,
	0x10, 0x70, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x4e, 0x10, 0x71, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x4f,
	0x10, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x52, 0x10, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4a,
	0x10, 0x76, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4d, 0x10, 0x78, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x52,
	0x10, 0x79, 0x12, 0x06, 0x0a, 0x02, 0x46, 0x41, 0x10, 0x7d, 0x12, 0x07, 0x0a, 0x02, 0x50, 0x54,
	0x10, 0x80, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x52, 0x4d, 0x10, 0x82, 0x01, 0x12, 0x07, 0x0a, 0x02,
	0x52, 0x4e, 0x10, 0x83, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x52, 0x55, 0x10, 0x85, 0x01, 0x12, 0x07,
	0x0a, 0x02, 0x53, 0x43, 0x10, 0x87, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x53, 0x45, 0x10, 0x89, 0x01,
	0x12, 0x07, 0x0a, 0x02, 0x53, 0x4d, 0x10, 0x8a, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x53, 0x47, 0x10,
	0x8b, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x53, 0x52, 0x10, 0x8c, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x53,
	0x4e, 0x10, 0x8e, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x53, 0x54, 0x10, 0x93, 0x01, 0x12, 0x07, 0x0a,
	0x02, 0x53, 0x55, 0x10, 0x95, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x53, 0x57, 0x10, 0x96, 0x01, 0x12,
	0x07, 0x0a, 0x02, 0x53, 0x53, 0x10, 0x97, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x53, 0x56, 0x10, 0x98,
	0x01, 0x12, 0x07, 0x0a, 0x02, 0x54, 0x41, 0x10, 0x99, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x54, 0x45,
	0x10, 0x9a, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x54, 0x48, 0x10, 0x9c, 0x01, 0x12, 0x07, 0x0a, 0x02,
	0x54, 0x49, 0x10, 0x9d, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x42, 0x4f, 0x10, 0x9e, 0x01, 0x12, 0x07,
	0x0a, 0x02, 0x54, 0x4c, 0x10, 0xa0, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x54, 0x4e, 0x10, 0xa1, 0x01,
	0x12, 0x07, 0x0a, 0x02, 0x54, 0x4f, 0x10, 0xa2, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x54, 0x52, 0x10,
	0xa3, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x54, 0x53, 0x10, 0xa4, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x54,
	0x57, 0x10, 0xa6, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x54, 0x59, 0x10, 0xa7, 0x01, 0x12, 0x07, 0x0a,
	0x02, 0x55, 0x4b, 0x10, 0xa9, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x55, 0x52, 0x10, 0xaa, 0x01, 0x12,
	0x07, 0x0a, 0x02, 0x56, 0x45, 0x10, 0xac, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x56, 0x49, 0x10, 0xad,
	0x01, 0x12, 0x07, 0x0a, 0x02, 0x56, 0x4f, 0x10, 0xae, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x57, 0x41,
	0x10, 0xaf, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x43, 0x59, 0x10, 0xb0, 0x01, 0x12, 0x07, 0x0a, 0x02,
	0x57, 0x4f, 0x10, 0xb1, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x46, 0x59, 0x10, 0xb2, 0x01, 0x12, 0x07,
	0x0a, 0x02, 0x58, 0x48, 0x10, 0xb3, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x59, 0x49, 0x10, 0xb4, 0x01,
	0x12, 0x07, 0x0a, 0x02, 0x59, 0x4f, 0x10, 0xb5, 0x01, 0x12, 0x07, 0x0a, 0x02, 0x5a, 0x55, 0x10,
	0xb7, 0x01, 0x42, 0x4c, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x01, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x72, 0x6d, 0x63,
	0x67, 0x75, 0x69, 0x6e, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2d,
	0x64, 0x61, 0x74, 0x61, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_enums_lang_proto_rawDescOnce sync.Once
	file_api_enums_lang_proto_rawDescData = file_api_enums_lang_proto_rawDesc
)

func file_api_enums_lang_proto_rawDescGZIP() []byte {
	file_api_enums_lang_proto_rawDescOnce.Do(func() {
		file_api_enums_lang_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_enums_lang_proto_rawDescData)
	})
	return file_api_enums_lang_proto_rawDescData
}

var file_api_enums_lang_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_enums_lang_proto_goTypes = []interface{}{
	(Language)(0), // 0: google.retail.enums.lang.Language
}
var file_api_enums_lang_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_enums_lang_proto_init() }
func file_api_enums_lang_proto_init() {
	if File_api_enums_lang_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_enums_lang_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_enums_lang_proto_goTypes,
		DependencyIndexes: file_api_enums_lang_proto_depIdxs,
		EnumInfos:         file_api_enums_lang_proto_enumTypes,
	}.Build()
	File_api_enums_lang_proto = out.File
	file_api_enums_lang_proto_rawDesc = nil
	file_api_enums_lang_proto_goTypes = nil
	file_api_enums_lang_proto_depIdxs = nil
}
