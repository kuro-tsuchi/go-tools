package helper

import (
	"net"
	"regexp"
	"time"
)

var (
	TArr 		 TsArr
	TStr 		 TsStr // transfer
	TInt         TsInt
	TFloat 		 TsFloat
	TFile		 TsFile
	TUri         TsUri
	THash        TsHash
	TCallFunc    TsCallFunc
	TDebug       TsDebug
	TOs          TsOs
	TConv        TsConvert
	TTime        TsTime
	TUuid        TsUuid
	TCorn        TsCorn
	TJson		 TsJson
	TEncrypt	 TsEncrypt
	TPrivyCiders []*net.IPNet
	NowTime 				 = time.Now()
	RegDNSName               = regexp.MustCompile(PatternDNSName)
	RegMultiByte             = regexp.MustCompile(PatternMultibyte)
	RegFloat                 = regexp.MustCompile(PatternFloat)
	RegFullWidth             = regexp.MustCompile(PatternFullwidth)
	RegHalfWidth             = regexp.MustCompile(PatternHalfWidth)
	RegAlphaLower            = regexp.MustCompile(PatternAlphaLower)
	RegAlphaUpper            = regexp.MustCompile(PatternAlphaUpper)
	RegChineseAll            = regexp.MustCompile(PatternChineseAll)
	RegChineseName           = regexp.MustCompile(PatternChineseName)
	RegEmail                 = regexp.MustCompile(PatternEmail)
	RegMobileCN              = regexp.MustCompile(PatternMobileCN)
	RegTelephone             = regexp.MustCompile(PatternTelephone)
	RegPhone                 = regexp.MustCompile(PatternPhone)
	RegCreditNo              = regexp.MustCompile(PatternCreditNo)
	RegDatetime              = regexp.MustCompile(PatternDatetime)
	RegAlphaNumeric          = regexp.MustCompile(PatternAlphaNumeric)
	RegRGBColor              = regexp.MustCompile(PatternRGBColor)
	RegWhitespaceAll         = regexp.MustCompile(PatternWhitespaceAll)
	RegWhitespaceHas         = regexp.MustCompile(PatternWhitespaceHas)
	RegBase64                = regexp.MustCompile(PatternBase64)
	RegBase64Image           = regexp.MustCompile(PatternBase64Image)
	RegMd5                   = regexp.MustCompile(PatternMd5)
	RegSha1                  = regexp.MustCompile(PatternSha1)
	RegSha256                = regexp.MustCompile(PatternSha256)
	RegSha512                = regexp.MustCompile(PatternSha512)
	// CreditArea ???????????????
	CreditArea = map[string]string{
		"11": "??????",
		"12": "??????",
		"13": "??????",
		"14": "??????",
		"15": "?????????",
		"21": "??????",
		"22": "??????",
		"23": "?????????",
		"31": "??????",
		"32": "??????",
		"33": "??????",
		"34": "??????",
		"35": "??????",
		"36": "??????",
		"37": "??????",
		"41": "??????",
		"42": "??????",
		"43": "??????",
		"44": "??????",
		"45": "??????",
		"46": "??????",
		"50": "??????",
		"51": "??????",
		"52": "??????",
		"53": "??????",
		"54": "??????",
		"61": "??????",
		"62": "??????",
		"63": "??????",
		"64": "??????",
		"65": "??????",
		"71": "??????",
		"81": "??????",
		"82": "??????",
		"91": "??????"}
)