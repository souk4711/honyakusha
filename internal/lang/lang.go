package lang

import (
	"strings"

	"github.com/Xuanwo/go-locale"
)

var (
	// https://github.com/soimort/translate-shell/wiki/Languages
	data = map[string]Lang{
		"af":       {Code: "af", Name: "Afrikaans", LocaleName: "Afrikaans"},
		"am":       {Code: "am", Name: "Amharic", LocaleName: "አማርኛ"},
		"ar":       {Code: "ar", Name: "Arabic", LocaleName: "العربية"},
		"as":       {Code: "as", Name: "Assamese", LocaleName: "অসমীয়া"},
		"ay":       {Code: "ay", Name: "Aymara", LocaleName: "Aymar aru"},
		"az":       {Code: "az", Name: "Azerbaijani", LocaleName: "Azərbaycanca"},
		"ba":       {Code: "ba", Name: "Bashkir", LocaleName: "Башҡортса"},
		"be":       {Code: "be", Name: "Belarusian", LocaleName: "беларуская"},
		"bg":       {Code: "bg", Name: "Bulgarian", LocaleName: "български"},
		"bho":      {Code: "bho", Name: "Bhojpuri", LocaleName: "भोजपुरी"},
		"bm":       {Code: "bm", Name: "Bambara", LocaleName: "Bamanankan"},
		"bn":       {Code: "bn", Name: "Bengali", LocaleName: "বাংলা"},
		"bo":       {Code: "bo", Name: "Tibetan", LocaleName: "བོད་ཡིག"},
		"br":       {Code: "br", Name: "Breton", LocaleName: "Brezhoneg"},
		"bs":       {Code: "bs", Name: "Bosnian", LocaleName: "Bosanski"},
		"ca":       {Code: "ca", Name: "Catalan", LocaleName: "Català"},
		"ceb":      {Code: "ceb", Name: "Cebuano", LocaleName: "Cebuano"},
		"chr":      {Code: "chr", Name: "Cherokee", LocaleName: "ᏣᎳᎩ"},
		"ckb":      {Code: "ckb", Name: "Kurdish (Central)", LocaleName: "سۆرانی"},
		"co":       {Code: "co", Name: "Corsican", LocaleName: "Corsu"},
		"cs":       {Code: "cs", Name: "Czech", LocaleName: "Čeština"},
		"cv":       {Code: "cv", Name: "Chuvash", LocaleName: "Чӑвашла"},
		"cy":       {Code: "cy", Name: "Welsh", LocaleName: "Cymraeg"},
		"da":       {Code: "da", Name: "Danish", LocaleName: "Dansk"},
		"de":       {Code: "de", Name: "German", LocaleName: "Deutsch"},
		"doi":      {Code: "doi", Name: "Dogri", LocaleName: "डोगरी"},
		"dv":       {Code: "dv", Name: "Dhivehi", LocaleName: "ދިވެހި"},
		"dz":       {Code: "dz", Name: "Dzongkha", LocaleName: "རྫོང་ཁ"},
		"ee":       {Code: "ee", Name: "Ewe", LocaleName: "Eʋegbe"},
		"el":       {Code: "el", Name: "Greek", LocaleName: "Ελληνικά"},
		"en":       {Code: "en", Name: "English", LocaleName: "English"},
		"en-GB":    {Code: "en-GB", Name: "English (British)", LocaleName: "English (British)"},
		"en-US":    {Code: "en-US", Name: "English (American)", LocaleName: "English (American)"},
		"eo":       {Code: "eo", Name: "Esperanto", LocaleName: "Esperanto"},
		"es":       {Code: "es", Name: "Spanish", LocaleName: "Español"},
		"et":       {Code: "et", Name: "Estonian", LocaleName: "Eesti"},
		"eu":       {Code: "eu", Name: "Basque", LocaleName: "Euskara"},
		"fa":       {Code: "fa", Name: "Persian", LocaleName: "فارسی"},
		"fi":       {Code: "fi", Name: "Finnish", LocaleName: "Suomi"},
		"fj":       {Code: "fj", Name: "Fijian", LocaleName: "Vosa Vakaviti"},
		"fo":       {Code: "fo", Name: "Faroese", LocaleName: "Føroyskt"},
		"fr":       {Code: "fr", Name: "French", LocaleName: "Français"},
		"fr-CA":    {Code: "fr-CA", Name: "French (Canadian)", LocaleName: "Français canadien"},
		"fy":       {Code: "fy", Name: "Frisian", LocaleName: "Frysk"},
		"ga":       {Code: "ga", Name: "Irish", LocaleName: "Gaeilge"},
		"gd":       {Code: "gd", Name: "Scots Gaelic", LocaleName: "Gàidhlig"},
		"gl":       {Code: "gl", Name: "Galician", LocaleName: "Galego"},
		"gn":       {Code: "gn", Name: "Guarani", LocaleName: "Avañe'ẽ"},
		"gom":      {Code: "gom", Name: "Konkani", LocaleName: "कोंकणी"},
		"gu":       {Code: "gu", Name: "Gujarati", LocaleName: "ગુજરાતી"},
		"ha":       {Code: "ha", Name: "Hausa", LocaleName: "Hausa"},
		"haw":      {Code: "haw", Name: "Hawaiian", LocaleName: "ʻŌlelo Hawaiʻi"},
		"he":       {Code: "he", Name: "Hebrew", LocaleName: "עִבְרִית"},
		"hi":       {Code: "hi", Name: "Hindi", LocaleName: "हिन्दी"},
		"hmn":      {Code: "hmn", Name: "Hmong", LocaleName: "Hmoob"},
		"hr":       {Code: "hr", Name: "Croatian", LocaleName: "Hrvatski"},
		"hsb":      {Code: "hsb", Name: "Upper Sorbian", LocaleName: "Hornjoserbšćina"},
		"ht":       {Code: "ht", Name: "Haitian Creole", LocaleName: "Kreyòl Ayisyen"},
		"hu":       {Code: "hu", Name: "Hungarian", LocaleName: "Magyar"},
		"hy":       {Code: "hy", Name: "Armenian", LocaleName: "Հայերեն"},
		"id":       {Code: "id", Name: "Indonesian", LocaleName: "Bahasa Indonesia"},
		"ie":       {Code: "ie", Name: "Interlingue", LocaleName: "Interlingue"},
		"ig":       {Code: "ig", Name: "Igbo", LocaleName: "Igbo"},
		"ikt":      {Code: "ikt", Name: "Inuinnaqtun", LocaleName: "Inuinnaqtun"},
		"ilo":      {Code: "ilo", Name: "Ilocano", LocaleName: "Ilokano"},
		"is":       {Code: "is", Name: "Icelandic", LocaleName: "Íslenska"},
		"it":       {Code: "it", Name: "Italian", LocaleName: "Italiano"},
		"iu":       {Code: "iu", Name: "Inuktitut", LocaleName: "ᐃᓄᒃᑎᑐᑦ"},
		"iu-Latn":  {Code: "iu-Latn", Name: "Inuktitut (Latin)", LocaleName: "Inuktitut"},
		"ja":       {Code: "ja", Name: "Japanese", LocaleName: "日本語"},
		"jv":       {Code: "jv", Name: "Javanese", LocaleName: "Basa Jawa"},
		"ka":       {Code: "ka", Name: "Georgian", LocaleName: "ქართული"},
		"kk":       {Code: "kk", Name: "Kazakh", LocaleName: "Қазақ тілі"},
		"kl":       {Code: "kl", Name: "Greenlandic", LocaleName: "Kalaallisut"},
		"km":       {Code: "km", Name: "Khmer", LocaleName: "ភាសាខ្មែរ"},
		"kn":       {Code: "kn", Name: "Kannada", LocaleName: "ಕನ್ನಡ"},
		"ko":       {Code: "ko", Name: "Korean", LocaleName: "한국어"},
		"kri":      {Code: "kri", Name: "Krio", LocaleName: "Krio"},
		"ku":       {Code: "ku", Name: "Kurdish (Northern)", LocaleName: "Kurmancî"},
		"ky":       {Code: "ky", Name: "Kyrgyz", LocaleName: "Кыргызча"},
		"la":       {Code: "la", Name: "Latin", LocaleName: "Latina"},
		"lb":       {Code: "lb", Name: "Luxembourgish", LocaleName: "Lëtzebuergesch"},
		"lg":       {Code: "lg", Name: "Luganda", LocaleName: "Luganda"},
		"ln":       {Code: "ln", Name: "Lingala", LocaleName: "Lingála"},
		"lo":       {Code: "lo", Name: "Lao", LocaleName: "ລາວ"},
		"lt":       {Code: "lt", Name: "Lithuanian", LocaleName: "Lietuvių"},
		"lus":      {Code: "lus", Name: "Mizo", LocaleName: "Mizo ṭawng"},
		"lv":       {Code: "lv", Name: "Latvian", LocaleName: "Latviešu"},
		"lzh":      {Code: "lzh", Name: "Chinese (Literary)", LocaleName: "文言"},
		"mai":      {Code: "mai", Name: "Maithili", LocaleName: "मैथिली"},
		"mg":       {Code: "mg", Name: "Malagasy", LocaleName: "Malagasy"},
		"mhr":      {Code: "mhr", Name: "Eastern Mari", LocaleName: "Олык марий"},
		"mi":       {Code: "mi", Name: "Maori", LocaleName: "Māori"},
		"mk":       {Code: "mk", Name: "Macedonian", LocaleName: "Македонски"},
		"ml":       {Code: "ml", Name: "Malayalam", LocaleName: "മലയാളം"},
		"mn":       {Code: "mn", Name: "Mongolian", LocaleName: "Монгол"},
		"mn-Mong":  {Code: "mn-Mong", Name: "Mongolian (Traditional)", LocaleName: "ᠮᠣᠩᠭᠣᠯ"},
		"mni-Mtei": {Code: "mni-Mtei", Name: "Meiteilon", LocaleName: "ꯃꯤꯇꯩꯂꯣꯟ"},
		"mr":       {Code: "mr", Name: "Marathi", LocaleName: "मराठी"},
		"mrj":      {Code: "mrj", Name: "Hill Mari", LocaleName: "Кырык мары"},
		"ms":       {Code: "ms", Name: "Malay", LocaleName: "Bahasa Melayu"},
		"mt":       {Code: "mt", Name: "Maltese", LocaleName: "Malti"},
		"my":       {Code: "my", Name: "Myanmar", LocaleName: "မြန်မာစာ"},
		"ne":       {Code: "ne", Name: "Nepali", LocaleName: "नेपाली"},
		"nl":       {Code: "nl", Name: "Dutch", LocaleName: "Nederlands"},
		"no":       {Code: "no", Name: "Norwegian", LocaleName: "Norwegian"},
		"nso":      {Code: "nso", Name: "Sepedi", LocaleName: "Sepedi"},
		"ny":       {Code: "ny", Name: "Chichewa", LocaleName: "Nyanja"},
		"oc":       {Code: "oc", Name: "Occitan", LocaleName: "Occitan"},
		"om":       {Code: "om", Name: "Oromo", LocaleName: "Afaan Oromoo"},
		"or":       {Code: "or", Name: "Odia", LocaleName: "ଓଡ଼ିଆ"},
		"otq":      {Code: "otg", Name: "Querétaro Otomi", LocaleName: "Hñąñho"},
		"pa":       {Code: "pa", Name: "Punjabi", LocaleName: "ਪੰਜਾਬੀ"},
		"pap":      {Code: "pap", Name: "Papiamento", LocaleName: "Papiamentu"},
		"pl":       {Code: "pl", Name: "Polish", LocaleName: "Polski"},
		"prs":      {Code: "prs", Name: "Dari", LocaleName: "دری"},
		"ps":       {Code: "ps", Name: "Pashto", LocaleName: "پښتو"},
		"pt":       {Code: "pt", Name: "Portuguese", LocaleName: "Português"},
		"pt-BR":    {Code: "pt-BR", Name: "Portuguese (Brazilian)", LocaleName: "Português Brasileiro"},
		"pt-PT":    {Code: "pt-PT", Name: "Portuguese (European)", LocaleName: "Português Europeu"},
		"qu":       {Code: "qu", Name: "Quechua", LocaleName: "Runasimi"},
		"rm":       {Code: "rm", Name: "Romansh", LocaleName: "Rumantsch"},
		"ro":       {Code: "ro", Name: "Romanian", LocaleName: "Română"},
		"ru":       {Code: "ru", Name: "Russian", LocaleName: "Русский"},
		"rw":       {Code: "rw", Name: "Ikinyarwanda", LocaleName: "Ikinyarwanda"},
		"sa":       {Code: "sa", Name: "Sanskrit", LocaleName: "संस्कृतम्"},
		"sah":      {Code: "sah", Name: "Yakut", LocaleName: "Sakha"},
		"sd":       {Code: "sd", Name: "Sindhi", LocaleName: "سنڌي"},
		"si":       {Code: "si", Name: "Sinhala", LocaleName: "සිංහල"},
		"sk":       {Code: "sk", Name: "Slovak", LocaleName: "Slovenčina"},
		"sl":       {Code: "sl", Name: "Slovenian", LocaleName: "Slovenščina"},
		"sm":       {Code: "sm", Name: "Samoan", LocaleName: "Gagana Sāmoa"},
		"sn":       {Code: "sn", Name: "Shona", LocaleName: "chiShona"},
		"so":       {Code: "so", Name: "Somali", LocaleName: "Soomaali"},
		"sq":       {Code: "sq", Name: "Albanian", LocaleName: "Shqip"},
		"sr":       {Code: "sr", Name: "Serbian", LocaleName: "Српски"},
		"sr-Cyrl":  {Code: "sr-Cyrl", Name: "Serbian (Cyrillic)", LocaleName: "Српски"},
		"sr-Latn":  {Code: "sr-Latn", Name: "Serbian (Latin)", LocaleName: "Srpski"},
		"st":       {Code: "st", Name: "Sesotho", LocaleName: "Sesotho"},
		"su":       {Code: "su", Name: "Sundanese ", LocaleName: "Basa Sunda"},
		"sv":       {Code: "sv", Name: "Swedish", LocaleName: "Svenska"},
		"sw":       {Code: "sw", Name: "Swahili", LocaleName: "Kiswahili"},
		"ta":       {Code: "ta", Name: "Tamil", LocaleName: "தமிழ்"},
		"te":       {Code: "te", Name: "Telugu", LocaleName: "తెలుగు"},
		"tg":       {Code: "tg", Name: "Tajik", LocaleName: "Тоҷикӣ"},
		"th":       {Code: "th", Name: "Thai", LocaleName: "ไทย"},
		"ti":       {Code: "ti", Name: "Tigrinya", LocaleName: "ትግርኛ"},
		"tk":       {Code: "tk", Name: "Turkmen", LocaleName: "Türkmen"},
		"tl":       {Code: "tl", Name: "Filipino", LocaleName: "Filipino"},
		"tlh-Latn": {Code: "tlh-Latn", Name: "Klingon", LocaleName: "tlhIngan Hol"},
		"tn":       {Code: "tn", Name: "Setswana", LocaleName: "Setswana"},
		"to":       {Code: "to", Name: "Tongan", LocaleName: "Lea faka-Tonga"},
		"tr":       {Code: "tr", Name: "Turkish", LocaleName: "Türkçe"},
		"ts":       {Code: "ts", Name: "Tsonga", LocaleName: "Xitsonga"},
		"tt":       {Code: "tt", Name: "Tatar", LocaleName: "татарча"},
		"tw":       {Code: "tw", Name: "Twi", LocaleName: "Twi"},
		"ty":       {Code: "ty", Name: "Tahitian", LocaleName: "Reo Tahiti"},
		"udm":      {Code: "udm", Name: "Udmurt", LocaleName: "Удмурт"},
		"ug":       {Code: "ug", Name: "Uyghur", LocaleName: "ئۇيغۇر تىلى"},
		"uk":       {Code: "uk", Name: "Ukrainian", LocaleName: "Українська"},
		"ur":       {Code: "ur", Name: "Urdu", LocaleName: "اُردُو"},
		"uz":       {Code: "uz", Name: "Uzbek", LocaleName: "Oʻzbek tili"},
		"vi":       {Code: "vi", Name: "Vietnamese", LocaleName: "Tiếng Việt"},
		"vo":       {Code: "vo", Name: "Volapük", LocaleName: "Volapük"},
		"wo":       {Code: "wo", Name: "Wolof", LocaleName: "Wollof"},
		"xh":       {Code: "xh", Name: "Xhosa", LocaleName: "isiXhosa"},
		"yi":       {Code: "yi", Name: "Yiddish", LocaleName: "ייִדיש"},
		"yo":       {Code: "yo", Name: "Yoruba", LocaleName: "Yorùbá"},
		"yua":      {Code: "yua", Name: "Yucatec Maya", LocaleName: "Màaya T'àan"},
		"yue":      {Code: "yue", Name: "Cantonese", LocaleName: "粵語"},
		"zh":       {Code: "zh", Name: "Chinese", LocaleName: "中文"},
		"zh-CN":    {Code: "zh-CN", Name: "Chinese (Simplified)", LocaleName: "简体中文"},
		"zh-TW":    {Code: "zh-TW", Name: "Chinese (Traditional)", LocaleName: "繁體中文"},
		"zu":       {Code: "zu", Name: "Zulu", LocaleName: "isiZulu"},
	}
)

type Lang struct {
	Code       string
	Name       string
	LocaleName string
}

func (l *Lang) Code_639_1() string {
	if idx := strings.Index(l.Code, "-"); idx == -1 {
		return l.Code
	} else {
		return l.Code[0:idx]
	}
}

func AutoDetect() string {
	tag, err := locale.Detect()
	if err != nil {
		return ""
	}

	code := tag.String()
	if _, ok := data[code]; ok {
		return code
	}

	if idx := strings.Index(code, "-"); idx != -1 {
		if _, ok := data[code[0:idx]]; ok {
			return code[0:idx]
		}
	}

	return ""
}

func Query(code string) Lang {
	return data[code]
}
