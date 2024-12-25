package Error

type Msg string

const (
	EM_EmptyStr         Msg = ""
	EM_JsonEncode       Msg = "json encode failed."
	EM_JsonDecode       Msg = "json decode failed."
	EM_AesEnrypt        Msg = "aes encrypt failed."
	EM_AesDecrypt       Msg = "aes decrypt failed."
	EM_Md5Encrypt       Msg = "md5 encrypt failed."
	EM_Md5Decrypt       Msg = "md5 decrypt failed."
	EM_InvalidParamsNum Msg = "invalid param number."
	EM_CommonValidate   Msg = "common validate failed."
	EM_StrToNum         Msg = "string to number failed."
	EM_Db               Msg = "handle db failed."
)

type code int

// const (
// 	EC_Success                                 code = 0    // success
// 	EC_AbnormalCode                            code = 1001 // abnormal code
// 	EC_InvalidAuth                             code = 1002 // invalid auth
// 	EC_InvalidAgent                            code = 1003 // invalid agent
// 	EC_BadParams                               code = 1004 // bad params
// 	EC_NullData                                code = 1005 // null data
// 	EC_Account_PlayerLastBetRecordByGame       code = 2001 // account
// 	EC_GameId                                  code = 2002 // game id
// 	EC_Account_BetRecords                      code = 3001 // account
// 	EC_TimeFormat_BetRecords                   code = 3002 // time format
// 	EC_TimeZoneOverDays7_BetRecords            code = 3003 // time zone over 7 days
// 	EC_Account_GetBetRecordsByRoomId           code = 4001 // account
// 	EC_TimeFormat_GetBetRecordsByRoomId        code = 4002 // time format
// 	EC_TimeZoneOverDays7_GetBetRecordsByRoomId code = 4003 // time zone over 7 days
// 	EC_Account_GetLiveRoomId                   code = 5001 // account
// )

func (c code) FlagMap() map[int]string {
	return map[int]string{
		0:    "success",
		1001: "abnormal code",
		1002: "invalid auth",
		1003: "invalid agent",
		1004: "bad params",
		1005: "null data",
		2001: "account",
		2002: "game id",
		3001: "account",
		3002: "time format",
		3003: "time zone over 7 days",
		4001: "account",
		4002: "time format",
		4003: "time zone over 7 days",
		5001: "account",
	}
}

var Code = struct {
	Success                                 code // success
	AbnormalCode                            code // abnormal code
	InvalidAuth                             code // invalid auth
	InvalidAgent                            code // invalid agent
	BadParams                               code // bad params
	NullData                                code // null data
	Account_PlayerLastBetRecordByGame       code // account
	GameId                                  code // game id
	Account_BetRecords                      code // account
	TimeFormat_BetRecords                   code // time format
	TimeZoneOverDays7_BetRecords            code // time zone over 7 days
	Account_GetBetRecordsByRoomId           code // account
	TimeFormat_GetBetRecordsByRoomId        code // time format
	TimeZoneOverDays7_GetBetRecordsByRoomId code // time zone over 7 days
	Account_GetLiveRoomId                   code // account
}{
	0,    // success
	1001, // abnormal code
	1002, // invalid auth
	1003, // invalid agent
	1004, // bad params
	1005, // null data
	2001, // account
	2002, // game id
	3001, // account
	3002, // time format
	3003, // time zone over 7 days
	4001, // account
	4002, // time format
	4003, // time zone over 7 days
	5001, // account
}
