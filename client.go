package deepl

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
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

	MinifyHTML     = true
	DontMinifyHTML = false
)

type TranslationResponse struct {
	Translations []TranslationResponseTranslations `json:"translations"`
}
type TranslationResponseTranslations struct {
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

type Translation struct {
	Texts     []string `json:"texts"`
	Source    string   `json:"source"`
	Target    string   `json:"target"`
	SplitSent string   `json:"split_sentences"`
	Preserve  string   `json:"preserve_formatting"`
	Formality string   `json:"formality"`
}

type DeepL struct {
	ApiKey     string
	ApiBaseURL string
	Minifier   *minify.M
}

func NewClient(apiKey string) *DeepL {

	minifier := minify.New()
	minifier.AddFunc("text/html", html.Minify)

	return &DeepL{
		ApiKey:     apiKey,
		ApiBaseURL: "https://api.deepl.com/v2",
		Minifier:   minifier,
	}
}

func (dl *DeepL) GetUsage() (UsageResponse, error) {
	var client = resty.New()
	resp, err := client.R().
		SetResult(&UsageResponse{}).
		SetAuthScheme("DeepL-Auth-Key").
		SetAuthToken(dl.ApiKey).
		Get(fmt.Sprintf("%s/usage", dl.ApiBaseURL))
	if err != nil {
		return UsageResponse{}, err
	}
	usage := resp.Result().(*UsageResponse)
	return *usage, nil
}

func (dl *DeepL) Translate(text, sourceLanguage, targetLanguage string) (translation Translation, err error) {
	splitSentencesOrNot := DontSplitSentences
	preserveFormattingOrNot := PreserveFormatting
	formalityLevel := FormalityMore
	resp, err := dl.translate(text, sourceLanguage, targetLanguage, splitSentencesOrNot, preserveFormattingOrNot, formalityLevel, MinifyHTML)
	if err != nil {
		return Translation{}, err
	}

	translation.Source = sourceLanguage
	translation.Target = targetLanguage
	translation.SplitSent = splitSentencesOrNot
	translation.Preserve = preserveFormattingOrNot
	translation.Formality = formalityLevel

	for _, t := range resp.Translations {
		translation.Texts = append(translation.Texts, t.Text)
	}

	return translation, nil
}

func (dl *DeepL) translate(text, sourceLanguage, targetLanguage, splitSentences, preserveFormatting, formality string, minifyHTML bool) (translations *TranslationResponse, err error) {
	if minifyHTML {
		text, err = dl.Minifier.String("text/html", text)
		if err != nil {
			return translations, fmt.Errorf("failed to minify text: %s", err)
		}
	}

	//if len(text) > 1024 {
	//	return translations, fmt.Errorf("text is too long, the DeepL API has a maximum of 1024 characters per request")
	//}

	var client = resty.New()
	//resp, err := client.R().
	//	SetResult(&TranslationResponse{}).
	//	SetAuthScheme("DeepL-Auth-Key").
	//	SetAuthToken(dl.ApiKey).
	//	SetQueryParams(map[string]string{
	//		"text":                    text,
	//		"source_lang":             sourceLanguage,
	//		"target_lang":             targetLanguage,
	//		"split_sentences":         splitSentences,
	//		"preserve_formatting":     preserveFormatting,
	//		"formality":               formality,
	//		"ignore_unsupported_lang": "true",
	//		"tag_handling":            TagHandlingHTML,
	//	}).
	//	Get(fmt.Sprintf("%s/translate", dl.ApiBaseURL))

	resp, err := client.R().
		SetResult(&TranslationResponse{}).
		SetAuthScheme("DeepL-Auth-Key").
		SetAuthToken(dl.ApiKey).
		SetFormData(map[string]string{
			"text":                    text,
			"source_lang":             sourceLanguage,
			"target_lang":             targetLanguage,
			"split_sentences":         splitSentences,
			"preserve_formatting":     preserveFormatting,
			"formality":               formality,
			"ignore_unsupported_lang": "true",
			"tag_handling":            TagHandlingHTML,
		}).
		Post(fmt.Sprintf("%s/translate", dl.ApiBaseURL))

	if err != nil {
		return translations, fmt.Errorf("failed to translate text: %s", err)
	}
	translations = resp.Result().(*TranslationResponse)
	return translations, nil
}
