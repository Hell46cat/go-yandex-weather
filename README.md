# Go Yandex Weather API

![GitHub](https://img.shields.io/github/license/Hell46cat/go-yandex-weather)
![Go Version](https://img.shields.io/github/go-mod/go-version/Hell46cat/go-yandex-weather)

Простая и удобная библиотека-клиент для работы с API Яндекс.Погоды на Go. Библиотека предоставляет доступ к текущей погоде и прогнозам на несколько дней с минимальным количеством кода.

## Особенности

- ✅ Простой и понятный API
- ✅ Поддержка контекстов для управления таймаутами
- ✅ Полное покрытие API Яндекс.Погоды
- ✅ Переводы погодных условий и направлений ветра
- ✅ Удобное форматирование информации о погоде
## Навигация
- [Установка](#Установка)
- [Быстрый старт](#Быстрый-старт)
- [Инициализация клиента](#Инициализация-клиента)
- [Получение текущей погоды](#Получение-текущей-погоды)
- [Получение прогноза на несколько дней](#Получени-прогноза-на-несколько-дней)
- [Форматирование данных о погоде](#Форматирование-данных-о-погоде)
- [Структуры данных](#Структуры-данных)
- [Переводы погодных условий](#Переводы-погодных-условий)
- [Примеры использования](#Примеры-использования)
## Установка

```bash
go get -u github.com/Hell46cat/go-yandex-weather
```

## Быстрый старт

```
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/Hell46cat/go-yandex-weather"
)

func main() {
    // Создаем клиент с API-ключом
    client := weather.NewClient("ваш-api-ключ")
    
    // Координаты Москвы
    lat, lon := 55.7558, 37.6176
    
    // Создаем контекст с таймаутом
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Получаем текущую погоду
    resp, err := client.GetCurrent(ctx, lat, lon, "ru_RU")
    if err != nil {
        log.Fatalf("Ошибка получения погоды: %v", err)
    }
    
    // Выводим основную информацию
    fact := resp.Fact
    fmt.Printf("Температура: %d°C, %s\n", 
        fact.Temp, 
        weather.TranslateCondition(fact.Condition))
}
```

## Инициализация клиента
```
// Стандартная инициализация
client := weather.NewClient("ваш-api-ключ")

// Если требуется кастомный HTTP-клиент
httpClient := &http.Client{
    Timeout: 15 * time.Second,
}
client := &weather.Client{
    APIKey:     "ваш-api-ключ",
    HTTPClient: httpClient,
}
```
## Получение текущей погоды
```
// Создаем контекст
ctx := context.Background()

// Получаем текущую погоду для Санкт-Петербурга
resp, err := client.GetCurrent(ctx, 59.9386, 30.3141, "ru_RU")
if err != nil {
    log.Fatalf("Ошибка получения погоды: %v", err)
}

// Работаем с данными
fmt.Printf("Погода: %d°C, %s\n", 
    resp.Fact.Temp, 
    weather.TranslateCondition(resp.Fact.Condition))
fmt.Printf("Ветер: %s %.1f м/с\n", 
    weather.TranslateWindDirection(resp.Fact.WindDir), 
    resp.Fact.WindSpeed)
```
## Получение прогноза на несколько дней
```
// Получаем прогноз на 5 дней с почасовыми данными
forecast, err := client.GetForecast(ctx, lat, lon, "ru_RU", 5, true)
if err != nil {
    log.Fatalf("Ошибка получения прогноза: %v", err)
}

// Выводим прогноз на каждый день
for i, day := range forecast.Forecasts {
    fmt.Printf("\nДень %d (%s):\n", i+1, day.Date)
    
    // Информация о дне
    dayPart := day.Parts.Day
    fmt.Printf("Днем: %d°C (%d°C...%d°C), %s\n", 
        dayPart.TempAvg, 
        dayPart.TempMin, 
        dayPart.TempMax, 
        weather.TranslateCondition(dayPart.Condition))
    
    // Информация о ночи
    nightPart := day.Parts.Night
    fmt.Printf("Ночью: %d°C (%d°C...%d°C), %s\n", 
        nightPart.TempAvg, 
        nightPart.TempMin, 
        nightPart.TempMax, 
        weather.TranslateCondition(nightPart.Condition))
    
    // Информация о восходе и закате
    fmt.Printf("Восход: %s, Закат: %s\n", 
        day.Sunrise, 
        day.Sunset)
    
    // Если были запрошены почасовые данные
    if len(day.Hours) > 0 {
        fmt.Println("Почасовой прогноз (избранные часы):")
        // Выводим каждый третий час для компактности
        for i := 0; i < len(day.Hours); i += 3 {
            hour := day.Hours[i]
            fmt.Printf("  %s:00 - %d°C, %s\n", 
                hour.Hour, 
                hour.Temp, 
                weather.TranslateCondition(hour.Condition))
        }
    }
}
```
## Форматирование данных о погоде
```
// Получаем текущую погоду
resp, err := client.GetCurrent(ctx, lat, lon, "ru_RU")
if err != nil {
    log.Fatalf("Ошибка получения погоды: %v", err)
}

// Используем встроенную функцию форматирования
formattedWeather := weather.FormatCurrentWeather(resp, "ru_RU")
fmt.Println(formattedWeather)

// Также можно форматировать данные по-своему
fact := resp.Fact
fmt.Printf("🌡 Сейчас: %d°C (%s)\n", 
    fact.Temp, 
    weather.TranslateCondition(fact.Condition))
fmt.Printf("💨 Ветер: %s %.1f м/с\n", 
    weather.TranslateWindDirection(fact.WindDir), 
    fact.WindSpeed)
```
## Структуры данных
Полное описание всех структур и полей можно найти в файле models.go
## Переводы погодных условий
Библиотека предоставляет функции для перевода кодов погодных условий и направления ветра:
```
// Погодные условия
condition := weather.TranslateCondition("clear") // "ясно"
condition := weather.TranslateCondition("partly-cloudy") // "малооблачно"
condition := weather.TranslateCondition("overcast") // "пасмурно"

// Направления ветра
windDir := weather.TranslateWindDirection("n") // "северный"
windDir := weather.TranslateWindDirection("sw") // "юго-западный"
```
## Примеры использования
Примеры находятся в директории examples: examples/simple/main.go - простой пример получения погоды