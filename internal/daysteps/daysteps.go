package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию

	// 1. Разделение строки на слайс строк
	stepAndTime := strings.Split(data, ",")
	// 2. Проверка длинны слайса
	if len(stepAndTime) != 2 {
		return 0, 0, fmt.Errorf("Длина строки %s не соответсвует формату", data)
	}

	// 3. Преобразование первого элемента слайса (количество шагов) в тип int.
	step, err := strconv.Atoi(stepAndTime[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка при конвертации: %v", err)
	}

	// 4. Проверка информации о количестве шагов
	if step <= 0 {
		return 0, 0, fmt.Errorf("Отрицательное количество шагов: %v", err)
	}

	// 5. Преобразование второго элемента слайса в time.Duration.
	time, err := time.ParseDuration(stepAndTime[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка при выполнении time.ParseDuration: %v", err)
	}
	if time <= 0 {
		return 0, 0, fmt.Errorf("Ошибка! Отрицательное время: %v", err)
	}

	// 6. Если всё прошло без ошибок, верните количество шагов, продолжительность и nil (для ошибки).
	return step, time, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	// 1. Получение данные о прогулке с помощью функции parsePackage().
	steps, walkDuration, err := parsePackage(data)
	if err != nil {
		log.Println("Возникла ошибка при парсинге строки: ", err)
		return ""
	}

	// 2. Повторная проверка информации о количестве шагов.
	if steps <= 0 {
		return ""
	}

	// 3. Вычисление дистанции в метрах.
	distince := float64(steps) * stepLength

	// 4. Перевод дистанции в километры.
	distinceKm := distince / mInKm

	// 5. Вычисление количества калорий, потраченных на прогулке.
	// Функция WalkingSpentCalories() определена в пакете spentcalories.
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkDuration)
	if err != nil {
		log.Println("Ошибка при расчёте калорий", err)
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distinceKm, calories)
	return result
}
