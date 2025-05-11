package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// BaseURL - базовый URL API Яндекс.Погоды
const BaseURL = "https://api.weather.yandex.ru/v2/forecast"

// Client представляет клиент для работы с API Яндекс.Погоды
type Client struct {
	// APIKey - ключ API Яндекс.Погоды
	APIKey     string
	// HTTPClient - HTTP-клиент для выполнения запросов
	HTTPClient *http.Client
}

// NewClient создает новый экземпляр клиента API Яндекс.Погоды
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// GetForecast получает прогноз погоды по заданным координатам
// lat и lon - географические координаты (широта и долгота)
// lang - язык ответа (например, "ru_RU" или "en_US")
// limit - количество дней в прогнозе (от 1 до 7)
// hours - включать ли почасовой прогноз
func (c *Client) GetForecast(ctx context.Context, lat, lon float64, lang string, limit int, hours bool) (*WeatherResponse, error) {
	// Формируем URL с параметрами запроса
	url := fmt.Sprintf("%s?lat=%.6f&lon=%.6f", BaseURL, lat, lon)
	
	// Добавляем язык, если указан
	if lang != "" {
		url = fmt.Sprintf("%s&lang=%s", url, lang)
	}
	
	// Добавляем лимит дней, если указан
	if limit > 0 {
		url = fmt.Sprintf("%s&limit=%d", url, limit)
	}
	
	// Добавляем флаг почасового прогноза
	if hours {
		url = fmt.Sprintf("%s&hours=true", url)
	} else {
		url = fmt.Sprintf("%s&hours=false", url)
	}
	
	// Создаем HTTP-запрос
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}
	
	// Добавляем заголовок с ключом API
	req.Header.Set("X-Yandex-API-Key", c.APIKey)
	
	// Выполняем запрос
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()
	
	// Проверяем код ответа
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("неуспешный код ответа: %d, тело: %s", resp.StatusCode, string(bodyBytes))
	}
	
	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
	}
	
	// Разбираем JSON в структуру
	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON: %w", err)
	}
	
	return &weather, nil
}

// GetCurrent получает текущую погоду по координатам
func (c *Client) GetCurrent(ctx context.Context, lat, lon float64, lang string) (*WeatherResponse, error) {
	// Для получения только текущей погоды устанавливаем лимит в 1 день без почасовых данных
	return c.GetForecast(ctx, lat, lon, lang, 1, false)
}

// FormatCurrentWeather форматирует текущую погоду в удобочитаемый вид
func FormatCurrentWeather(weather *WeatherResponse, lang string) string {
	if weather == nil {
		return "Нет данных о погоде"
	}
	
	// Определяем название места
	city := "неизвестном месте"
	
	fact := weather.Fact
	
	var result string
	
	// Определяем язык вывода
	if lang == "ru_RU" || lang == "" {
		// Русская локализация
		result += fmt.Sprintf("Погода в %s:\n", city)
		result += fmt.Sprintf("Температура: %d°C\n", fact.Temp)
		result += fmt.Sprintf("Ощущается как: %d°C\n", fact.FeelsLike)
		result += fmt.Sprintf("Погодные условия: %s\n", TranslateCondition(fact.Condition))
		result += fmt.Sprintf("Ветер: %s %.1f м/с\n", TranslateWindDirection(fact.WindDir), fact.WindSpeed)
		result += fmt.Sprintf("Влажность: %d%%\n", fact.Humidity)
		
		// Проверка на валидность значения давления
		pressureText := "Нет данных"
		if weather.Info.DefPressureMm > 0 {
			pressureText = fmt.Sprintf("%d мм рт.ст.", weather.Info.DefPressureMm)
		}
		result += fmt.Sprintf("Давление: %s\n", pressureText)
	} else {
		// Английская локализация
		result += fmt.Sprintf("Weather in %s:\n", city)
		result += fmt.Sprintf("Temperature: %d°C\n", fact.Temp)
		result += fmt.Sprintf("Feels like: %d°C\n", fact.FeelsLike)
		result += fmt.Sprintf("Condition: %s\n", fact.Condition)
		result += fmt.Sprintf("Wind: %s %.1f m/s\n", fact.WindDir, fact.WindSpeed)
		result += fmt.Sprintf("Humidity: %d%%\n", fact.Humidity)
		
		// Проверка на валидность значения давления
		pressureText := "No data"
		if weather.Info.DefPressureMm > 0 {
			pressureText = fmt.Sprintf("%d mmHg", weather.Info.DefPressureMm)
		}
		result += fmt.Sprintf("Pressure: %s\n", pressureText)
	}
	
	return result
}