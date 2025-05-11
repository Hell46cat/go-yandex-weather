package main

import (
	"context"
	"fmt"
	"log"
	"time"
	
	weather "github.com/Hell46cat/go-yandex-weather"  // Импортируем пакет
)

func main() {
	// Создаем клиент с API-ключом
	client := weather.NewClient("ваш-api-ключ")
	
	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Координаты Москвы
	lat, lon := 55.7558, 37.6176
	
	// Получаем текущую погоду
	resp, err := client.GetCurrent(ctx, lat, lon, "ru_RU")
	if err != nil {
		log.Fatalf("Ошибка получения погоды: %v", err)
	}
	
	// Выводим основную информацию
	fact := resp.Fact  // Используем resp вместо weather
	
	fmt.Printf("Координаты: %.4f, %.4f\n", resp.Info.Lat, resp.Info.Lon)
	fmt.Printf("Температура: %d°C\n", fact.Temp)
	fmt.Printf("Ощущается как: %d°C\n", fact.FeelsLike)
	fmt.Printf("Погодные условия: %s\n", weather.TranslateCondition(fact.Condition))  // Вызываем через пакет
	fmt.Printf("Ветер: %s %.1f м/с\n", weather.TranslateWindDirection(fact.WindDir), fact.WindSpeed)  // Вызываем через пакет
	fmt.Printf("Влажность: %d%%\n", fact.Humidity)
	
	// Проверка на валидность значения давления
	pressureText := "Нет данных"
	if resp.Info.DefPressureMm > 0 {
		pressureText = fmt.Sprintf("%d мм рт.ст.", resp.Info.DefPressureMm)
	}
	fmt.Printf("Давление: %s\n", pressureText)
	
	// Получаем прогноз на 3 дня
	forecast, err := client.GetForecast(ctx, lat, lon, "ru_RU", 3, true)
	if err != nil {
		log.Fatalf("Ошибка получения прогноза: %v", err)
	}
	
	// Выводим прогноз на 3 дня
	fmt.Println("\nПрогноз на 3 дня:")
	for i, day := range forecast.Forecasts {
		fmt.Printf("\nДень %d (%s):\n", i+1, day.Date)
		fmt.Printf("Температура днем: %d°C, %s\n", 
			day.Parts.Day.TempAvg, 
			weather.TranslateCondition(day.Parts.Day.Condition))  // Вызываем через пакет
		fmt.Printf("Температура ночью: %d°C, %s\n", 
			day.Parts.Night.TempAvg, 
			weather.TranslateCondition(day.Parts.Night.Condition))  // Вызываем через пакет
	}
}