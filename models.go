// Package weather предоставляет клиент для API Яндекс.Погоды
package weather

// WeatherResponse представляет ответ от API Яндекс.Погоды
type WeatherResponse struct {
	Now        int64      `json:"now"`        // Текущее время в Unix формате
	NowDt      string     `json:"now_dt"`     // Текущее время в ISO формате
	Info       Info       `json:"info"`       // Информация о запросе
	Fact       Fact       `json:"fact"`       // Фактическая погода
	Forecasts  []Forecast `json:"forecasts"`  // Прогнозы на несколько дней
}

// Info содержит информацию о запросе
type Info struct {
	Lat            float64 `json:"lat"`            // Широта
	Lon            float64 `json:"lon"`            // Долгота
	URL            string  `json:"url"`            // URL для получения прогноза в браузере
	DefPressureMm  int     `json:"def_pressure_mm"` // Стандартное давление в мм рт.ст.
	DefPressurePa  int     `json:"def_pressure_pa"` // Стандартное давление в гПа
	TzInfo         TzInfo  `json:"tzinfo"`         // Информация о часовом поясе
}

// TzInfo содержит информацию о часовом поясе
type TzInfo struct {
	Name   string `json:"name"`   // Название часового пояса
	Abbr   string `json:"abbr"`   // Аббревиатура часового пояса
	Dst    bool   `json:"dst"`    // Переход на летнее время
	Offset int    `json:"offset"` // Смещение от UTC в секундах
}

// Fact содержит информацию о фактической погоде
type Fact struct {
	Temp         int     `json:"temp"`         // Температура (°C)
	FeelsLike    int     `json:"feels_like"`   // Ощущаемая температура (°C)
	Icon         string  `json:"icon"`         // Код иконки погоды
	Condition    string  `json:"condition"`    // Код погодного описания
	WindSpeed    float64 `json:"wind_speed"`   // Скорость ветра (м/с)
	WindGust     float64 `json:"wind_gust"`    // Скорость порывов ветра (м/с)
	WindDir      string  `json:"wind_dir"`     // Направление ветра
	Humidity     int     `json:"humidity"`     // Влажность воздуха (%)
	Daytime      string  `json:"daytime"`      // Светлое или темное время суток (d/n)
	Polar        bool    `json:"polar"`        // Признак полярного дня или ночи
	Season       string  `json:"season"`       // Время года
	ObsTime      int64   `json:"obs_time"`     // Время наблюдения
	PrecType     int     `json:"prec_type"`    // Тип осадков
	PrecStrength float64 `json:"prec_strength"` // Сила осадков
	Cloudness    float64 `json:"cloudness"`    // Облачность
}

// Forecast содержит прогноз погоды на день
type Forecast struct {
	Date       string    `json:"date"`        // Дата прогноза в формате ГГГГ-ММ-ДД
	DateTs     int64     `json:"date_ts"`     // Дата прогноза в Unix формате
	Week       int       `json:"week"`        // Номер недели
	Sunrise    string    `json:"sunrise"`     // Время восхода Солнца
	Sunset     string    `json:"sunset"`      // Время заката Солнца
	MoonCode   int       `json:"moon_code"`   // Код фазы Луны
	MoonText   string    `json:"moon_text"`   // Текстовое описание фазы Луны
	Parts      Parts     `json:"parts"`       // Прогноз по частям дня
	Hours      []Hour    `json:"hours"`       // Почасовой прогноз
}

// Parts содержит прогноз погоды по частям дня
type Parts struct {
	Night      Part `json:"night"`      // Ночь
	Morning    Part `json:"morning"`    // Утро
	Day        Part `json:"day"`        // День
	Evening    Part `json:"evening"`    // Вечер
	DayShort   Part `json:"day_short"`  // Световой день (короткий вариант)
	NightShort Part `json:"night_short"` // Ночь (короткий вариант)
}

// Part содержит прогноз погоды на часть дня
type Part struct {
	TempMin        int     `json:"temp_min"`         // Минимальная температура (°C)
	TempMax        int     `json:"temp_max"`         // Максимальная температура (°C)
	TempAvg        int     `json:"temp_avg"`         // Средняя температура (°C)
	FeelsLike      int     `json:"feels_like"`       // Ощущаемая температура (°C)
	Icon           string  `json:"icon"`             // Код иконки погоды
	Condition      string  `json:"condition"`        // Код погодного описания
	Daytime        string  `json:"daytime"`          // Светлое или темное время суток (d/n)
	Polar          bool    `json:"polar"`            // Признак полярного дня или ночи
	WindSpeed      float64 `json:"wind_speed"`       // Скорость ветра (м/с)
	WindGust       float64 `json:"wind_gust"`        // Скорость порывов ветра (м/с)
	WindDir        string  `json:"wind_dir"`         // Направление ветра
	PressureMm     int     `json:"pressure_mm"`      // Давление (мм рт.ст.)
	PressurePa     int     `json:"pressure_pa"`      // Давление (гПа)
	Humidity       int     `json:"humidity"`         // Влажность воздуха (%)
	PrecMm         float64 `json:"prec_mm"`          // Количество осадков (мм)
	PrecPeriod     int     `json:"prec_period"`      // Период осадков (минуты)
	PrecProbability int     `json:"prec_probability"` // Вероятность осадков (%)
	PrecType       int     `json:"prec_type"`        // Тип осадков
	PrecStrength   float64 `json:"prec_strength"`    // Сила осадков
	Cloudness      float64 `json:"cloudness"`        // Облачность
}

// Hour содержит почасовой прогноз погоды
type Hour struct {
	Hour            string  `json:"hour"`             // Час в формате ЧЧ
	HourTs          int64   `json:"hour_ts"`          // Час в Unix формате
	Temp            int     `json:"temp"`             // Температура (°C)
	FeelsLike       int     `json:"feels_like"`       // Ощущаемая температура (°C)
	Icon            string  `json:"icon"`             // Код иконки погоды
	Condition       string  `json:"condition"`        // Код погодного описания
	Cloudness       float64 `json:"cloudness"`        // Облачность
	PrecType        int     `json:"prec_type"`        // Тип осадков
	PrecStrength    float64 `json:"prec_strength"`    // Сила осадков
	IsThunder       bool    `json:"is_thunder"`       // Признак грозы
	WindDir         string  `json:"wind_dir"`         // Направление ветра
	WindSpeed       float64 `json:"wind_speed"`       // Скорость ветра (м/с)
	WindGust        float64 `json:"wind_gust"`        // Скорость порывов ветра (м/с)
	PressureMm      int     `json:"pressure_mm"`      // Давление (мм рт.ст.)
	PressurePa      int     `json:"pressure_pa"`      // Давление (гПа)
	Humidity        int     `json:"humidity"`         // Влажность воздуха (%)
	PrecMm          float64 `json:"prec_mm"`          // Количество осадков (мм)
	PrecPeriod      int     `json:"prec_period"`      // Период осадков (минуты)
	PrecProbability int     `json:"prec_probability"` // Вероятность осадков (%)
}