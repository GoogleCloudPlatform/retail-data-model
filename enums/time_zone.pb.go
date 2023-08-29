// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.10
// source: api/enums/time_zone.proto

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

type TimeZone int32

const (
	TimeZone_UTC0000                                  TimeZone = 0
	TimeZone_UTCMinus1200InternationalDateLineWest    TimeZone = 1
	TimeZone_UTCMinus1100MidwayIsland_Samoa           TimeZone = 2
	TimeZone_UTCMinus1000Hawaii                       TimeZone = 3
	TimeZone_UTCMinus0900Alaska                       TimeZone = 4
	TimeZone_UTCMinus0800PacificTime                  TimeZone = 5
	TimeZone_UTCMinus0800Tijuana_BajaCalifornia       TimeZone = 6
	TimeZone_UTCMinus0700Arizona                      TimeZone = 7
	TimeZone_UTCMinus0700MountainTime                 TimeZone = 8
	TimeZone_UTCMinus0700Chihuahua_Lapaz_Mazatlan     TimeZone = 9
	TimeZone_UTCMinus0600CentralAmerica               TimeZone = 10
	TimeZone_UTCMinus0600CentralTime                  TimeZone = 11
	TimeZone_UTCMinus0600Guadalajara_MexicoCity       TimeZone = 12
	TimeZone_UTCMinus0600Saskatchewan                 TimeZone = 13
	TimeZone_UTCMinus0500Bogota_Lima_Quito_Riobranco  TimeZone = 14
	TimeZone_UTCMinus0500EasternTime                  TimeZone = 15
	TimeZone_UTCMinus0500Indiana                      TimeZone = 16
	TimeZone_UTCMinus0400AtlanticTime                 TimeZone = 17
	TimeZone_UTCMinus0400Lapaz                        TimeZone = 18
	TimeZone_UTCMinus0400Manaus                       TimeZone = 19
	TimeZone_UTCMinus0400Santiago                     TimeZone = 20
	TimeZone_UTCMinus0430Caracas                      TimeZone = 21
	TimeZone_UTCMinus0330Newfoundland                 TimeZone = 22
	TimeZone_UTCMinus0300Brasilia                     TimeZone = 23
	TimeZone_UTCMinus0300BuenosAires_Georgetown       TimeZone = 24
	TimeZone_UTCMinus0300Greenland                    TimeZone = 25
	TimeZone_UTCMinus0300Montevideo                   TimeZone = 26
	TimeZone_UTCMinus0200MidAtlantic                  TimeZone = 27
	TimeZone_UTCMinus0100Azores                       TimeZone = 28
	TimeZone_UTCMinus0100CapeVerdIs                   TimeZone = 29
	TimeZone_UTC0000Casablanca_Montrovia_Reykjavik    TimeZone = 30
	TimeZone_UTC0000Dublin_Edinburgh_Lisbon_London    TimeZone = 31
	TimeZone_UTCPlus0100Amsterdam_Berlin_Bern_Rome    TimeZone = 32
	TimeZone_UTCPlus0100Belgrade_Bratislava_Budapest  TimeZone = 33
	TimeZone_UTCPlus0100Brussels_Copenhagen_Madrid    TimeZone = 34
	TimeZone_UTCPlus0100Sarajevo_Skopje_Warsaw_Zagreb TimeZone = 35
	TimeZone_UTCPlus0100WestCentralAfrica             TimeZone = 36
	TimeZone_UTCPlus0200Amman                         TimeZone = 37
	TimeZone_UTCPlus0200Athens_Bucharest_Istanbul     TimeZone = 38
	TimeZone_UTCPlus0200Beirut                        TimeZone = 39
	TimeZone_UTCPlus0200Minsk                         TimeZone = 40
	TimeZone_UTCPlus0200Cairo                         TimeZone = 41
	TimeZone_UTCPlus0200Harare_Pretoria               TimeZone = 42
	TimeZone_UTCPlus0200Helsinki_Kyiv_Riga_Vilnius    TimeZone = 43
	TimeZone_UTCPlus0200Jerusalem                     TimeZone = 44
	TimeZone_UTCPlus0200Windhoek                      TimeZone = 45
	TimeZone_UTCPlus0300Baghdad                       TimeZone = 46
	TimeZone_UTCPlus0300Kuwait_Riyadh                 TimeZone = 47
	TimeZone_UTCPlus0300Moscow_StPetersburg_Volgograd TimeZone = 48
	TimeZone_UTCPlus0300Nairobi                       TimeZone = 49
	TimeZone_UTCPlus0300Tbilisi                       TimeZone = 50
	TimeZone_UTCPlus0330Tehran                        TimeZone = 51
	TimeZone_UTCPlus0400AbuDhabi_Muscat               TimeZone = 52
	TimeZone_UTCPlus0400Baku                          TimeZone = 53
	TimeZone_UTCPlus0400CaucasusStandardTime          TimeZone = 54
	TimeZone_UTCPlus0400Yerevan                       TimeZone = 55
	TimeZone_UTCPlus0430Kabul                         TimeZone = 56
	TimeZone_UTCPlus0500Ekaterinburg                  TimeZone = 57
	TimeZone_UTCPlus0500Islamabad_Karachi_Tashkent    TimeZone = 58
	TimeZone_UTCPlus0530Chennai_Kolkata_Mumbai        TimeZone = 59
	TimeZone_UTCPlus0530SriJayawardenepura            TimeZone = 60
	TimeZone_UTCPlus0545Kathmandu                     TimeZone = 61
	TimeZone_UTCPlus0600Almaty_Novosibirsk            TimeZone = 62
	TimeZone_UTCPlus0600Astana_Dhaka                  TimeZone = 63
	TimeZone_UTCPlus0630Yangon                        TimeZone = 64
	TimeZone_UTCPlus0700Bangkok_Hanoi_Jakarta         TimeZone = 65
	TimeZone_UTCPlus0700Krasnoyarsk                   TimeZone = 66
	TimeZone_UTCPlus0800Beijing_Chongqing_HongKong    TimeZone = 67
	TimeZone_UTCPlus0800Irkutsk_UlaanBataar           TimeZone = 68
	TimeZone_UTCPlus0800KualaLumpur_Singapore         TimeZone = 69
	TimeZone_UTCPlus0800Perth                         TimeZone = 70
	TimeZone_UTCPlus0800Taipei                        TimeZone = 71
	TimeZone_UTCPlus0900Osaka_Sapporo_Tokyo           TimeZone = 72
	TimeZone_UTCPlus0900Seoul                         TimeZone = 73
	TimeZone_UTCPlus0900Yakutsk                       TimeZone = 74
	TimeZone_UTCPlus0930Adelaide                      TimeZone = 75
	TimeZone_UTCPlus0930Darwin                        TimeZone = 76
	TimeZone_UTCPlus1000Brisbane                      TimeZone = 77
	TimeZone_UTCPlus1000Canberra_Melbourne_Sydney     TimeZone = 78
	TimeZone_UTCPlus1000Guam_PortMoresby              TimeZone = 79
	TimeZone_UTCPlus1000Hobart                        TimeZone = 80
	TimeZone_UTCPlus1000Vladivostok                   TimeZone = 81
	TimeZone_UTCPlus1100Magadan_Solomonis             TimeZone = 82
	TimeZone_UTCPlus1200Auckland_Wellington           TimeZone = 83
	TimeZone_UTCPlus1200Fiji_Kamchatka_MarshallIs     TimeZone = 84
	TimeZone_UTCPlus1300Nuku_alofa                    TimeZone = 85
)

// Enum value maps for TimeZone.
var (
	TimeZone_name = map[int32]string{
		0:  "UTC0000",
		1:  "UTCMinus1200InternationalDateLineWest",
		2:  "UTCMinus1100MidwayIsland_Samoa",
		3:  "UTCMinus1000Hawaii",
		4:  "UTCMinus0900Alaska",
		5:  "UTCMinus0800PacificTime",
		6:  "UTCMinus0800Tijuana_BajaCalifornia",
		7:  "UTCMinus0700Arizona",
		8:  "UTCMinus0700MountainTime",
		9:  "UTCMinus0700Chihuahua_Lapaz_Mazatlan",
		10: "UTCMinus0600CentralAmerica",
		11: "UTCMinus0600CentralTime",
		12: "UTCMinus0600Guadalajara_MexicoCity",
		13: "UTCMinus0600Saskatchewan",
		14: "UTCMinus0500Bogota_Lima_Quito_Riobranco",
		15: "UTCMinus0500EasternTime",
		16: "UTCMinus0500Indiana",
		17: "UTCMinus0400AtlanticTime",
		18: "UTCMinus0400Lapaz",
		19: "UTCMinus0400Manaus",
		20: "UTCMinus0400Santiago",
		21: "UTCMinus0430Caracas",
		22: "UTCMinus0330Newfoundland",
		23: "UTCMinus0300Brasilia",
		24: "UTCMinus0300BuenosAires_Georgetown",
		25: "UTCMinus0300Greenland",
		26: "UTCMinus0300Montevideo",
		27: "UTCMinus0200MidAtlantic",
		28: "UTCMinus0100Azores",
		29: "UTCMinus0100CapeVerdIs",
		30: "UTC0000Casablanca_Montrovia_Reykjavik",
		31: "UTC0000Dublin_Edinburgh_Lisbon_London",
		32: "UTCPlus0100Amsterdam_Berlin_Bern_Rome",
		33: "UTCPlus0100Belgrade_Bratislava_Budapest",
		34: "UTCPlus0100Brussels_Copenhagen_Madrid",
		35: "UTCPlus0100Sarajevo_Skopje_Warsaw_Zagreb",
		36: "UTCPlus0100WestCentralAfrica",
		37: "UTCPlus0200Amman",
		38: "UTCPlus0200Athens_Bucharest_Istanbul",
		39: "UTCPlus0200Beirut",
		40: "UTCPlus0200Minsk",
		41: "UTCPlus0200Cairo",
		42: "UTCPlus0200Harare_Pretoria",
		43: "UTCPlus0200Helsinki_Kyiv_Riga_Vilnius",
		44: "UTCPlus0200Jerusalem",
		45: "UTCPlus0200Windhoek",
		46: "UTCPlus0300Baghdad",
		47: "UTCPlus0300Kuwait_Riyadh",
		48: "UTCPlus0300Moscow_StPetersburg_Volgograd",
		49: "UTCPlus0300Nairobi",
		50: "UTCPlus0300Tbilisi",
		51: "UTCPlus0330Tehran",
		52: "UTCPlus0400AbuDhabi_Muscat",
		53: "UTCPlus0400Baku",
		54: "UTCPlus0400CaucasusStandardTime",
		55: "UTCPlus0400Yerevan",
		56: "UTCPlus0430Kabul",
		57: "UTCPlus0500Ekaterinburg",
		58: "UTCPlus0500Islamabad_Karachi_Tashkent",
		59: "UTCPlus0530Chennai_Kolkata_Mumbai",
		60: "UTCPlus0530SriJayawardenepura",
		61: "UTCPlus0545Kathmandu",
		62: "UTCPlus0600Almaty_Novosibirsk",
		63: "UTCPlus0600Astana_Dhaka",
		64: "UTCPlus0630Yangon",
		65: "UTCPlus0700Bangkok_Hanoi_Jakarta",
		66: "UTCPlus0700Krasnoyarsk",
		67: "UTCPlus0800Beijing_Chongqing_HongKong",
		68: "UTCPlus0800Irkutsk_UlaanBataar",
		69: "UTCPlus0800KualaLumpur_Singapore",
		70: "UTCPlus0800Perth",
		71: "UTCPlus0800Taipei",
		72: "UTCPlus0900Osaka_Sapporo_Tokyo",
		73: "UTCPlus0900Seoul",
		74: "UTCPlus0900Yakutsk",
		75: "UTCPlus0930Adelaide",
		76: "UTCPlus0930Darwin",
		77: "UTCPlus1000Brisbane",
		78: "UTCPlus1000Canberra_Melbourne_Sydney",
		79: "UTCPlus1000Guam_PortMoresby",
		80: "UTCPlus1000Hobart",
		81: "UTCPlus1000Vladivostok",
		82: "UTCPlus1100Magadan_Solomonis",
		83: "UTCPlus1200Auckland_Wellington",
		84: "UTCPlus1200Fiji_Kamchatka_MarshallIs",
		85: "UTCPlus1300Nuku_alofa",
	}
	TimeZone_value = map[string]int32{
		"UTC0000":                                  0,
		"UTCMinus1200InternationalDateLineWest":    1,
		"UTCMinus1100MidwayIsland_Samoa":           2,
		"UTCMinus1000Hawaii":                       3,
		"UTCMinus0900Alaska":                       4,
		"UTCMinus0800PacificTime":                  5,
		"UTCMinus0800Tijuana_BajaCalifornia":       6,
		"UTCMinus0700Arizona":                      7,
		"UTCMinus0700MountainTime":                 8,
		"UTCMinus0700Chihuahua_Lapaz_Mazatlan":     9,
		"UTCMinus0600CentralAmerica":               10,
		"UTCMinus0600CentralTime":                  11,
		"UTCMinus0600Guadalajara_MexicoCity":       12,
		"UTCMinus0600Saskatchewan":                 13,
		"UTCMinus0500Bogota_Lima_Quito_Riobranco":  14,
		"UTCMinus0500EasternTime":                  15,
		"UTCMinus0500Indiana":                      16,
		"UTCMinus0400AtlanticTime":                 17,
		"UTCMinus0400Lapaz":                        18,
		"UTCMinus0400Manaus":                       19,
		"UTCMinus0400Santiago":                     20,
		"UTCMinus0430Caracas":                      21,
		"UTCMinus0330Newfoundland":                 22,
		"UTCMinus0300Brasilia":                     23,
		"UTCMinus0300BuenosAires_Georgetown":       24,
		"UTCMinus0300Greenland":                    25,
		"UTCMinus0300Montevideo":                   26,
		"UTCMinus0200MidAtlantic":                  27,
		"UTCMinus0100Azores":                       28,
		"UTCMinus0100CapeVerdIs":                   29,
		"UTC0000Casablanca_Montrovia_Reykjavik":    30,
		"UTC0000Dublin_Edinburgh_Lisbon_London":    31,
		"UTCPlus0100Amsterdam_Berlin_Bern_Rome":    32,
		"UTCPlus0100Belgrade_Bratislava_Budapest":  33,
		"UTCPlus0100Brussels_Copenhagen_Madrid":    34,
		"UTCPlus0100Sarajevo_Skopje_Warsaw_Zagreb": 35,
		"UTCPlus0100WestCentralAfrica":             36,
		"UTCPlus0200Amman":                         37,
		"UTCPlus0200Athens_Bucharest_Istanbul":     38,
		"UTCPlus0200Beirut":                        39,
		"UTCPlus0200Minsk":                         40,
		"UTCPlus0200Cairo":                         41,
		"UTCPlus0200Harare_Pretoria":               42,
		"UTCPlus0200Helsinki_Kyiv_Riga_Vilnius":    43,
		"UTCPlus0200Jerusalem":                     44,
		"UTCPlus0200Windhoek":                      45,
		"UTCPlus0300Baghdad":                       46,
		"UTCPlus0300Kuwait_Riyadh":                 47,
		"UTCPlus0300Moscow_StPetersburg_Volgograd": 48,
		"UTCPlus0300Nairobi":                       49,
		"UTCPlus0300Tbilisi":                       50,
		"UTCPlus0330Tehran":                        51,
		"UTCPlus0400AbuDhabi_Muscat":               52,
		"UTCPlus0400Baku":                          53,
		"UTCPlus0400CaucasusStandardTime":          54,
		"UTCPlus0400Yerevan":                       55,
		"UTCPlus0430Kabul":                         56,
		"UTCPlus0500Ekaterinburg":                  57,
		"UTCPlus0500Islamabad_Karachi_Tashkent":    58,
		"UTCPlus0530Chennai_Kolkata_Mumbai":        59,
		"UTCPlus0530SriJayawardenepura":            60,
		"UTCPlus0545Kathmandu":                     61,
		"UTCPlus0600Almaty_Novosibirsk":            62,
		"UTCPlus0600Astana_Dhaka":                  63,
		"UTCPlus0630Yangon":                        64,
		"UTCPlus0700Bangkok_Hanoi_Jakarta":         65,
		"UTCPlus0700Krasnoyarsk":                   66,
		"UTCPlus0800Beijing_Chongqing_HongKong":    67,
		"UTCPlus0800Irkutsk_UlaanBataar":           68,
		"UTCPlus0800KualaLumpur_Singapore":         69,
		"UTCPlus0800Perth":                         70,
		"UTCPlus0800Taipei":                        71,
		"UTCPlus0900Osaka_Sapporo_Tokyo":           72,
		"UTCPlus0900Seoul":                         73,
		"UTCPlus0900Yakutsk":                       74,
		"UTCPlus0930Adelaide":                      75,
		"UTCPlus0930Darwin":                        76,
		"UTCPlus1000Brisbane":                      77,
		"UTCPlus1000Canberra_Melbourne_Sydney":     78,
		"UTCPlus1000Guam_PortMoresby":              79,
		"UTCPlus1000Hobart":                        80,
		"UTCPlus1000Vladivostok":                   81,
		"UTCPlus1100Magadan_Solomonis":             82,
		"UTCPlus1200Auckland_Wellington":           83,
		"UTCPlus1200Fiji_Kamchatka_MarshallIs":     84,
		"UTCPlus1300Nuku_alofa":                    85,
	}
)

func (x TimeZone) Enum() *TimeZone {
	p := new(TimeZone)
	*p = x
	return p
}

func (x TimeZone) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeZone) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_time_zone_proto_enumTypes[0].Descriptor()
}

func (TimeZone) Type() protoreflect.EnumType {
	return &file_api_enums_time_zone_proto_enumTypes[0]
}

func (x TimeZone) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeZone.Descriptor instead.
func (TimeZone) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_time_zone_proto_rawDescGZIP(), []int{0}
}

var File_api_enums_time_zone_proto protoreflect.FileDescriptor

var file_api_enums_time_zone_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2a, 0xf2, 0x14, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f,
	0x6e, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x54, 0x43, 0x30, 0x30, 0x30, 0x30, 0x10, 0x00, 0x12,
	0x29, 0x0a, 0x25, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x31, 0x32, 0x30, 0x30, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6e, 0x65, 0x57, 0x65, 0x73, 0x74, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x54,
	0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x31, 0x31, 0x30, 0x30, 0x4d, 0x69, 0x64, 0x77, 0x61, 0x79,
	0x49, 0x73, 0x6c, 0x61, 0x6e, 0x64, 0x5f, 0x53, 0x61, 0x6d, 0x6f, 0x61, 0x10, 0x02, 0x12, 0x16,
	0x0a, 0x12, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x31, 0x30, 0x30, 0x30, 0x48, 0x61,
	0x77, 0x61, 0x69, 0x69, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e,
	0x75, 0x73, 0x30, 0x39, 0x30, 0x30, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x10, 0x04, 0x12, 0x1b,
	0x0a, 0x17, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x38, 0x30, 0x30, 0x50, 0x61,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x10, 0x05, 0x12, 0x26, 0x0a, 0x22, 0x55,
	0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x38, 0x30, 0x30, 0x54, 0x69, 0x6a, 0x75, 0x61,
	0x6e, 0x61, 0x5f, 0x42, 0x61, 0x6a, 0x61, 0x43, 0x61, 0x6c, 0x69, 0x66, 0x6f, 0x72, 0x6e, 0x69,
	0x61, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30,
	0x37, 0x30, 0x30, 0x41, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x61, 0x10, 0x07, 0x12, 0x1c, 0x0a, 0x18,
	0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x37, 0x30, 0x30, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x10, 0x08, 0x12, 0x28, 0x0a, 0x24, 0x55, 0x54,
	0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x37, 0x30, 0x30, 0x43, 0x68, 0x69, 0x68, 0x75, 0x61,
	0x68, 0x75, 0x61, 0x5f, 0x4c, 0x61, 0x70, 0x61, 0x7a, 0x5f, 0x4d, 0x61, 0x7a, 0x61, 0x74, 0x6c,
	0x61, 0x6e, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73,
	0x30, 0x36, 0x30, 0x30, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x41, 0x6d, 0x65, 0x72, 0x69,
	0x63, 0x61, 0x10, 0x0a, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73,
	0x30, 0x36, 0x30, 0x30, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x10,
	0x0b, 0x12, 0x26, 0x0a, 0x22, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x36, 0x30,
	0x30, 0x47, 0x75, 0x61, 0x64, 0x61, 0x6c, 0x61, 0x6a, 0x61, 0x72, 0x61, 0x5f, 0x4d, 0x65, 0x78,
	0x69, 0x63, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x10, 0x0c, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x54, 0x43,
	0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x36, 0x30, 0x30, 0x53, 0x61, 0x73, 0x6b, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x77, 0x61, 0x6e, 0x10, 0x0d, 0x12, 0x2b, 0x0a, 0x27, 0x55, 0x54, 0x43, 0x4d, 0x69,
	0x6e, 0x75, 0x73, 0x30, 0x35, 0x30, 0x30, 0x42, 0x6f, 0x67, 0x6f, 0x74, 0x61, 0x5f, 0x4c, 0x69,
	0x6d, 0x61, 0x5f, 0x51, 0x75, 0x69, 0x74, 0x6f, 0x5f, 0x52, 0x69, 0x6f, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x6f, 0x10, 0x0e, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73,
	0x30, 0x35, 0x30, 0x30, 0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x10,
	0x0f, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x35, 0x30,
	0x30, 0x49, 0x6e, 0x64, 0x69, 0x61, 0x6e, 0x61, 0x10, 0x10, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x54,
	0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x34, 0x30, 0x30, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74,
	0x69, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x10, 0x11, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x54, 0x43, 0x4d,
	0x69, 0x6e, 0x75, 0x73, 0x30, 0x34, 0x30, 0x30, 0x4c, 0x61, 0x70, 0x61, 0x7a, 0x10, 0x12, 0x12,
	0x16, 0x0a, 0x12, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x34, 0x30, 0x30, 0x4d,
	0x61, 0x6e, 0x61, 0x75, 0x73, 0x10, 0x13, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x54, 0x43, 0x4d, 0x69,
	0x6e, 0x75, 0x73, 0x30, 0x34, 0x30, 0x30, 0x53, 0x61, 0x6e, 0x74, 0x69, 0x61, 0x67, 0x6f, 0x10,
	0x14, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x34, 0x33,
	0x30, 0x43, 0x61, 0x72, 0x61, 0x63, 0x61, 0x73, 0x10, 0x15, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x54,
	0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x33, 0x33, 0x30, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x10, 0x16, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x54, 0x43, 0x4d,
	0x69, 0x6e, 0x75, 0x73, 0x30, 0x33, 0x30, 0x30, 0x42, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x69, 0x61,
	0x10, 0x17, 0x12, 0x26, 0x0a, 0x22, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x33,
	0x30, 0x30, 0x42, 0x75, 0x65, 0x6e, 0x6f, 0x73, 0x41, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x47, 0x65,
	0x6f, 0x72, 0x67, 0x65, 0x74, 0x6f, 0x77, 0x6e, 0x10, 0x18, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x54,
	0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x33, 0x30, 0x30, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x6c,
	0x61, 0x6e, 0x64, 0x10, 0x19, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75,
	0x73, 0x30, 0x33, 0x30, 0x30, 0x4d, 0x6f, 0x6e, 0x74, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x10,
	0x1a, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x32, 0x30,
	0x30, 0x4d, 0x69, 0x64, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x10, 0x1b, 0x12, 0x16,
	0x0a, 0x12, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x30, 0x31, 0x30, 0x30, 0x41, 0x7a,
	0x6f, 0x72, 0x65, 0x73, 0x10, 0x1c, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x54, 0x43, 0x4d, 0x69, 0x6e,
	0x75, 0x73, 0x30, 0x31, 0x30, 0x30, 0x43, 0x61, 0x70, 0x65, 0x56, 0x65, 0x72, 0x64, 0x49, 0x73,
	0x10, 0x1d, 0x12, 0x29, 0x0a, 0x25, 0x55, 0x54, 0x43, 0x30, 0x30, 0x30, 0x30, 0x43, 0x61, 0x73,
	0x61, 0x62, 0x6c, 0x61, 0x6e, 0x63, 0x61, 0x5f, 0x4d, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x76, 0x69,
	0x61, 0x5f, 0x52, 0x65, 0x79, 0x6b, 0x6a, 0x61, 0x76, 0x69, 0x6b, 0x10, 0x1e, 0x12, 0x29, 0x0a,
	0x25, 0x55, 0x54, 0x43, 0x30, 0x30, 0x30, 0x30, 0x44, 0x75, 0x62, 0x6c, 0x69, 0x6e, 0x5f, 0x45,
	0x64, 0x69, 0x6e, 0x62, 0x75, 0x72, 0x67, 0x68, 0x5f, 0x4c, 0x69, 0x73, 0x62, 0x6f, 0x6e, 0x5f,
	0x4c, 0x6f, 0x6e, 0x64, 0x6f, 0x6e, 0x10, 0x1f, 0x12, 0x29, 0x0a, 0x25, 0x55, 0x54, 0x43, 0x50,
	0x6c, 0x75, 0x73, 0x30, 0x31, 0x30, 0x30, 0x41, 0x6d, 0x73, 0x74, 0x65, 0x72, 0x64, 0x61, 0x6d,
	0x5f, 0x42, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x5f, 0x42, 0x65, 0x72, 0x6e, 0x5f, 0x52, 0x6f, 0x6d,
	0x65, 0x10, 0x20, 0x12, 0x2b, 0x0a, 0x27, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x31,
	0x30, 0x30, 0x42, 0x65, 0x6c, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x42, 0x72, 0x61, 0x74, 0x69,
	0x73, 0x6c, 0x61, 0x76, 0x61, 0x5f, 0x42, 0x75, 0x64, 0x61, 0x70, 0x65, 0x73, 0x74, 0x10, 0x21,
	0x12, 0x29, 0x0a, 0x25, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x31, 0x30, 0x30, 0x42,
	0x72, 0x75, 0x73, 0x73, 0x65, 0x6c, 0x73, 0x5f, 0x43, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x4d, 0x61, 0x64, 0x72, 0x69, 0x64, 0x10, 0x22, 0x12, 0x2c, 0x0a, 0x28, 0x55,
	0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x31, 0x30, 0x30, 0x53, 0x61, 0x72, 0x61, 0x6a, 0x65,
	0x76, 0x6f, 0x5f, 0x53, 0x6b, 0x6f, 0x70, 0x6a, 0x65, 0x5f, 0x57, 0x61, 0x72, 0x73, 0x61, 0x77,
	0x5f, 0x5a, 0x61, 0x67, 0x72, 0x65, 0x62, 0x10, 0x23, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x54, 0x43,
	0x50, 0x6c, 0x75, 0x73, 0x30, 0x31, 0x30, 0x30, 0x57, 0x65, 0x73, 0x74, 0x43, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61, 0x10, 0x24, 0x12, 0x14, 0x0a, 0x10, 0x55,
	0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x32, 0x30, 0x30, 0x41, 0x6d, 0x6d, 0x61, 0x6e, 0x10,
	0x25, 0x12, 0x28, 0x0a, 0x24, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x32, 0x30, 0x30,
	0x41, 0x74, 0x68, 0x65, 0x6e, 0x73, 0x5f, 0x42, 0x75, 0x63, 0x68, 0x61, 0x72, 0x65, 0x73, 0x74,
	0x5f, 0x49, 0x73, 0x74, 0x61, 0x6e, 0x62, 0x75, 0x6c, 0x10, 0x26, 0x12, 0x15, 0x0a, 0x11, 0x55,
	0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x32, 0x30, 0x30, 0x42, 0x65, 0x69, 0x72, 0x75, 0x74,
	0x10, 0x27, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x32, 0x30,
	0x30, 0x4d, 0x69, 0x6e, 0x73, 0x6b, 0x10, 0x28, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x54, 0x43, 0x50,
	0x6c, 0x75, 0x73, 0x30, 0x32, 0x30, 0x30, 0x43, 0x61, 0x69, 0x72, 0x6f, 0x10, 0x29, 0x12, 0x1e,
	0x0a, 0x1a, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x32, 0x30, 0x30, 0x48, 0x61, 0x72,
	0x61, 0x72, 0x65, 0x5f, 0x50, 0x72, 0x65, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x10, 0x2a, 0x12, 0x29,
	0x0a, 0x25, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x32, 0x30, 0x30, 0x48, 0x65, 0x6c,
	0x73, 0x69, 0x6e, 0x6b, 0x69, 0x5f, 0x4b, 0x79, 0x69, 0x76, 0x5f, 0x52, 0x69, 0x67, 0x61, 0x5f,
	0x56, 0x69, 0x6c, 0x6e, 0x69, 0x75, 0x73, 0x10, 0x2b, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x54, 0x43,
	0x50, 0x6c, 0x75, 0x73, 0x30, 0x32, 0x30, 0x30, 0x4a, 0x65, 0x72, 0x75, 0x73, 0x61, 0x6c, 0x65,
	0x6d, 0x10, 0x2c, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x32,
	0x30, 0x30, 0x57, 0x69, 0x6e, 0x64, 0x68, 0x6f, 0x65, 0x6b, 0x10, 0x2d, 0x12, 0x16, 0x0a, 0x12,
	0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x33, 0x30, 0x30, 0x42, 0x61, 0x67, 0x68, 0x64,
	0x61, 0x64, 0x10, 0x2e, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30,
	0x33, 0x30, 0x30, 0x4b, 0x75, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x52, 0x69, 0x79, 0x61, 0x64, 0x68,
	0x10, 0x2f, 0x12, 0x2c, 0x0a, 0x28, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x33, 0x30,
	0x30, 0x4d, 0x6f, 0x73, 0x63, 0x6f, 0x77, 0x5f, 0x53, 0x74, 0x50, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x62, 0x75, 0x72, 0x67, 0x5f, 0x56, 0x6f, 0x6c, 0x67, 0x6f, 0x67, 0x72, 0x61, 0x64, 0x10, 0x30,
	0x12, 0x16, 0x0a, 0x12, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x33, 0x30, 0x30, 0x4e,
	0x61, 0x69, 0x72, 0x6f, 0x62, 0x69, 0x10, 0x31, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x54, 0x43, 0x50,
	0x6c, 0x75, 0x73, 0x30, 0x33, 0x30, 0x30, 0x54, 0x62, 0x69, 0x6c, 0x69, 0x73, 0x69, 0x10, 0x32,
	0x12, 0x15, 0x0a, 0x11, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x33, 0x33, 0x30, 0x54,
	0x65, 0x68, 0x72, 0x61, 0x6e, 0x10, 0x33, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x54, 0x43, 0x50, 0x6c,
	0x75, 0x73, 0x30, 0x34, 0x30, 0x30, 0x41, 0x62, 0x75, 0x44, 0x68, 0x61, 0x62, 0x69, 0x5f, 0x4d,
	0x75, 0x73, 0x63, 0x61, 0x74, 0x10, 0x34, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x54, 0x43, 0x50, 0x6c,
	0x75, 0x73, 0x30, 0x34, 0x30, 0x30, 0x42, 0x61, 0x6b, 0x75, 0x10, 0x35, 0x12, 0x23, 0x0a, 0x1f,
	0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x34, 0x30, 0x30, 0x43, 0x61, 0x75, 0x63, 0x61,
	0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x10,
	0x36, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x34, 0x30, 0x30,
	0x59, 0x65, 0x72, 0x65, 0x76, 0x61, 0x6e, 0x10, 0x37, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x54, 0x43,
	0x50, 0x6c, 0x75, 0x73, 0x30, 0x34, 0x33, 0x30, 0x4b, 0x61, 0x62, 0x75, 0x6c, 0x10, 0x38, 0x12,
	0x1b, 0x0a, 0x17, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x35, 0x30, 0x30, 0x45, 0x6b,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x62, 0x75, 0x72, 0x67, 0x10, 0x39, 0x12, 0x29, 0x0a, 0x25,
	0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x35, 0x30, 0x30, 0x49, 0x73, 0x6c, 0x61, 0x6d,
	0x61, 0x62, 0x61, 0x64, 0x5f, 0x4b, 0x61, 0x72, 0x61, 0x63, 0x68, 0x69, 0x5f, 0x54, 0x61, 0x73,
	0x68, 0x6b, 0x65, 0x6e, 0x74, 0x10, 0x3a, 0x12, 0x25, 0x0a, 0x21, 0x55, 0x54, 0x43, 0x50, 0x6c,
	0x75, 0x73, 0x30, 0x35, 0x33, 0x30, 0x43, 0x68, 0x65, 0x6e, 0x6e, 0x61, 0x69, 0x5f, 0x4b, 0x6f,
	0x6c, 0x6b, 0x61, 0x74, 0x61, 0x5f, 0x4d, 0x75, 0x6d, 0x62, 0x61, 0x69, 0x10, 0x3b, 0x12, 0x21,
	0x0a, 0x1d, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x35, 0x33, 0x30, 0x53, 0x72, 0x69,
	0x4a, 0x61, 0x79, 0x61, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x65, 0x70, 0x75, 0x72, 0x61, 0x10,
	0x3c, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x35, 0x34, 0x35,
	0x4b, 0x61, 0x74, 0x68, 0x6d, 0x61, 0x6e, 0x64, 0x75, 0x10, 0x3d, 0x12, 0x21, 0x0a, 0x1d, 0x55,
	0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x36, 0x30, 0x30, 0x41, 0x6c, 0x6d, 0x61, 0x74, 0x79,
	0x5f, 0x4e, 0x6f, 0x76, 0x6f, 0x73, 0x69, 0x62, 0x69, 0x72, 0x73, 0x6b, 0x10, 0x3e, 0x12, 0x1b,
	0x0a, 0x17, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x36, 0x30, 0x30, 0x41, 0x73, 0x74,
	0x61, 0x6e, 0x61, 0x5f, 0x44, 0x68, 0x61, 0x6b, 0x61, 0x10, 0x3f, 0x12, 0x15, 0x0a, 0x11, 0x55,
	0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x36, 0x33, 0x30, 0x59, 0x61, 0x6e, 0x67, 0x6f, 0x6e,
	0x10, 0x40, 0x12, 0x24, 0x0a, 0x20, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x37, 0x30,
	0x30, 0x42, 0x61, 0x6e, 0x67, 0x6b, 0x6f, 0x6b, 0x5f, 0x48, 0x61, 0x6e, 0x6f, 0x69, 0x5f, 0x4a,
	0x61, 0x6b, 0x61, 0x72, 0x74, 0x61, 0x10, 0x41, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x54, 0x43, 0x50,
	0x6c, 0x75, 0x73, 0x30, 0x37, 0x30, 0x30, 0x4b, 0x72, 0x61, 0x73, 0x6e, 0x6f, 0x79, 0x61, 0x72,
	0x73, 0x6b, 0x10, 0x42, 0x12, 0x29, 0x0a, 0x25, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30,
	0x38, 0x30, 0x30, 0x42, 0x65, 0x69, 0x6a, 0x69, 0x6e, 0x67, 0x5f, 0x43, 0x68, 0x6f, 0x6e, 0x67,
	0x71, 0x69, 0x6e, 0x67, 0x5f, 0x48, 0x6f, 0x6e, 0x67, 0x4b, 0x6f, 0x6e, 0x67, 0x10, 0x43, 0x12,
	0x22, 0x0a, 0x1e, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x38, 0x30, 0x30, 0x49, 0x72,
	0x6b, 0x75, 0x74, 0x73, 0x6b, 0x5f, 0x55, 0x6c, 0x61, 0x61, 0x6e, 0x42, 0x61, 0x74, 0x61, 0x61,
	0x72, 0x10, 0x44, 0x12, 0x24, 0x0a, 0x20, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x38,
	0x30, 0x30, 0x4b, 0x75, 0x61, 0x6c, 0x61, 0x4c, 0x75, 0x6d, 0x70, 0x75, 0x72, 0x5f, 0x53, 0x69,
	0x6e, 0x67, 0x61, 0x70, 0x6f, 0x72, 0x65, 0x10, 0x45, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x54, 0x43,
	0x50, 0x6c, 0x75, 0x73, 0x30, 0x38, 0x30, 0x30, 0x50, 0x65, 0x72, 0x74, 0x68, 0x10, 0x46, 0x12,
	0x15, 0x0a, 0x11, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x38, 0x30, 0x30, 0x54, 0x61,
	0x69, 0x70, 0x65, 0x69, 0x10, 0x47, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75,
	0x73, 0x30, 0x39, 0x30, 0x30, 0x4f, 0x73, 0x61, 0x6b, 0x61, 0x5f, 0x53, 0x61, 0x70, 0x70, 0x6f,
	0x72, 0x6f, 0x5f, 0x54, 0x6f, 0x6b, 0x79, 0x6f, 0x10, 0x48, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x54,
	0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x39, 0x30, 0x30, 0x53, 0x65, 0x6f, 0x75, 0x6c, 0x10, 0x49,
	0x12, 0x16, 0x0a, 0x12, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x39, 0x30, 0x30, 0x59,
	0x61, 0x6b, 0x75, 0x74, 0x73, 0x6b, 0x10, 0x4a, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x54, 0x43, 0x50,
	0x6c, 0x75, 0x73, 0x30, 0x39, 0x33, 0x30, 0x41, 0x64, 0x65, 0x6c, 0x61, 0x69, 0x64, 0x65, 0x10,
	0x4b, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x30, 0x39, 0x33, 0x30,
	0x44, 0x61, 0x72, 0x77, 0x69, 0x6e, 0x10, 0x4c, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x54, 0x43, 0x50,
	0x6c, 0x75, 0x73, 0x31, 0x30, 0x30, 0x30, 0x42, 0x72, 0x69, 0x73, 0x62, 0x61, 0x6e, 0x65, 0x10,
	0x4d, 0x12, 0x28, 0x0a, 0x24, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x31, 0x30, 0x30, 0x30,
	0x43, 0x61, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x61, 0x5f, 0x4d, 0x65, 0x6c, 0x62, 0x6f, 0x75, 0x72,
	0x6e, 0x65, 0x5f, 0x53, 0x79, 0x64, 0x6e, 0x65, 0x79, 0x10, 0x4e, 0x12, 0x1f, 0x0a, 0x1b, 0x55,
	0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x31, 0x30, 0x30, 0x30, 0x47, 0x75, 0x61, 0x6d, 0x5f, 0x50,
	0x6f, 0x72, 0x74, 0x4d, 0x6f, 0x72, 0x65, 0x73, 0x62, 0x79, 0x10, 0x4f, 0x12, 0x15, 0x0a, 0x11,
	0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x31, 0x30, 0x30, 0x30, 0x48, 0x6f, 0x62, 0x61, 0x72,
	0x74, 0x10, 0x50, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x31, 0x30,
	0x30, 0x30, 0x56, 0x6c, 0x61, 0x64, 0x69, 0x76, 0x6f, 0x73, 0x74, 0x6f, 0x6b, 0x10, 0x51, 0x12,
	0x20, 0x0a, 0x1c, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x31, 0x31, 0x30, 0x30, 0x4d, 0x61,
	0x67, 0x61, 0x64, 0x61, 0x6e, 0x5f, 0x53, 0x6f, 0x6c, 0x6f, 0x6d, 0x6f, 0x6e, 0x69, 0x73, 0x10,
	0x52, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x31, 0x32, 0x30, 0x30,
	0x41, 0x75, 0x63, 0x6b, 0x6c, 0x61, 0x6e, 0x64, 0x5f, 0x57, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x74, 0x6f, 0x6e, 0x10, 0x53, 0x12, 0x28, 0x0a, 0x24, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73,
	0x31, 0x32, 0x30, 0x30, 0x46, 0x69, 0x6a, 0x69, 0x5f, 0x4b, 0x61, 0x6d, 0x63, 0x68, 0x61, 0x74,
	0x6b, 0x61, 0x5f, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x6c, 0x49, 0x73, 0x10, 0x54, 0x12,
	0x19, 0x0a, 0x15, 0x55, 0x54, 0x43, 0x50, 0x6c, 0x75, 0x73, 0x31, 0x33, 0x30, 0x30, 0x4e, 0x75,
	0x6b, 0x75, 0x5f, 0x61, 0x6c, 0x6f, 0x66, 0x61, 0x10, 0x55, 0x42, 0x4c, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x72, 0x6d, 0x63, 0x67, 0x75, 0x69, 0x6e, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_enums_time_zone_proto_rawDescOnce sync.Once
	file_api_enums_time_zone_proto_rawDescData = file_api_enums_time_zone_proto_rawDesc
)

func file_api_enums_time_zone_proto_rawDescGZIP() []byte {
	file_api_enums_time_zone_proto_rawDescOnce.Do(func() {
		file_api_enums_time_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_enums_time_zone_proto_rawDescData)
	})
	return file_api_enums_time_zone_proto_rawDescData
}

var file_api_enums_time_zone_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_enums_time_zone_proto_goTypes = []interface{}{
	(TimeZone)(0), // 0: google.retail.enums.time.TimeZone
}
var file_api_enums_time_zone_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_enums_time_zone_proto_init() }
func file_api_enums_time_zone_proto_init() {
	if File_api_enums_time_zone_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_enums_time_zone_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_enums_time_zone_proto_goTypes,
		DependencyIndexes: file_api_enums_time_zone_proto_depIdxs,
		EnumInfos:         file_api_enums_time_zone_proto_enumTypes,
	}.Build()
	File_api_enums_time_zone_proto = out.File
	file_api_enums_time_zone_proto_rawDesc = nil
	file_api_enums_time_zone_proto_goTypes = nil
	file_api_enums_time_zone_proto_depIdxs = nil
}
