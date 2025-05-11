package weather

// ConditionTranslations содержит переводы погодных условий
var ConditionTranslations = map[string]string{
	"clear":                  "ясно",
	"partly-cloudy":          "малооблачно",
	"cloudy":                 "облачно с прояснениями",
	"overcast":               "пасмурно",
	"drizzle":                "морось",
	"rain":                   "дождь",
	"heavy-rain":             "сильный дождь",
	"showers":                "ливень",
	"wet-snow":               "дождь со снегом",
	"light-snow":             "небольшой снег",
	"light-rain":             "небольшой дождь",
	"snow":                   "снег",
	"snow-showers":           "снегопад",
	"hail":                   "град",
	"thunderstorm":           "гроза",
	"thunderstorm-with-rain": "дождь с грозой",
	"thunderstorm-with-hail": "гроза с градом",
}

// WindDirectionTranslations содержит переводы направлений ветра
var WindDirectionTranslations = map[string]string{
	"nw": "северо-западный",
	"n":  "северный",
	"ne": "северо-восточный",
	"e":  "восточный",
	"se": "юго-восточный",
	"s":  "южный",
	"sw": "юго-западный",
	"w":  "западный",
	"c":  "штиль",
}

// TranslateCondition переводит погодные условия с английского на русский
func TranslateCondition(condition string) string {
	translated, ok := ConditionTranslations[condition]
	if !ok {
		return condition
	}
	return translated
}

// TranslateWindDirection переводит направление ветра с английского на русский
func TranslateWindDirection(direction string) string {
	translated, ok := WindDirectionTranslations[direction]
	if !ok {
		return direction
	}
	return translated
}
