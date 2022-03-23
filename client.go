package deepl

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	SourceBulgarian  = "BG"
	SourceCzech      = "CS"
	SourceDanish     = "DA"
	SourceGerman     = "DE"
	SourceGreek      = "EL"
	SourceEnglish    = "EN"
	SourceSpanish    = "ES"
	SourceEstonian   = "ET"
	SourceFinnish    = "FI"
	SourceFrench     = "FR"
	SourceHungarian  = "HU"
	SourceItalian    = "IT"
	SourceJapanese   = "JA"
	SourceLithuanian = "LT"
	SourceLatvian    = "LV"
	SourceDutch      = "NL"
	SourcePolish     = "PL"
	SourcePortuguese = "PT"
	SourceRomanian   = "RO"
	SourceRussian    = "RU"
	SourceSlovak     = "SK"
	SourceSlovenian  = "SL"
	SourceSwedish    = "SV"
	SourceChinese    = "ZH"

	TargetBulgarian           = "BG"
	TargetCzech               = "CS"
	TargetDanish              = "DA"
	TargetGerman              = "DE"
	TargetGreek               = "EL"
	TargetEnglishBritish      = "EN-GB"
	TargetEnglishAmerican     = "EN-US"
	TargetSpanish             = "ES"
	TargetEstonian            = "ET"
	TargetFinnish             = "FI"
	TargetFrench              = "FR"
	TargetHungarian           = "HU"
	TargetItalian             = "IT"
	TargetJapanese            = "JA"
	TargetLithuanian          = "LT"
	TargetLatvian             = "LV"
	TargetDutch               = "NL"
	TargetPolish              = "PL"
	TargetPortuguese          = "PT-PT"
	TargetPortugueseBrazilian = "PT-BR"
	TargetRomanian            = "RO"
	TargetRussian             = "RU"
	TargetSlovak              = "SK"
	TargetSlovenian           = "SL"
	TargetSwedish             = "SV"
	TargetChinese             = "ZH"

	SplitSentences         = "1"
	DontSplitSentences     = "0"
	PreserveFormatting     = "1"
	DontPreserveFormatting = "0"
	FormalityDefault       = "default"
	FormalityMore          = "more"
	FormalityLess          = "less"
	TagHandlingXML         = "xml"
	TagHandlingHTML        = "html"
)

type TranslationResponse struct {
	Translations []Translations `json:"translations"`
}
type Translations struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}

type UsageResponse struct {
	CharacterCount int   `json:"character_count"`
	CharacterLimit int64 `json:"character_limit"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

var client = resty.New()
var apiBaseURL = "https://api-free.deepl.com/v2"

func GetUsage() {
	resp, err := client.R().SetResult(&UsageResponse{}).Get(fmt.Sprintf("%s/usage", apiBaseURL))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.Result().(*UsageResponse))
}
