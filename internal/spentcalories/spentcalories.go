package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию

	// 1. Разделение строки на слайс строк.
	info := strings.Split(data, ",")

	// 2. Проверка длины слайса.
	if len(info) != 3 {
		return 0, "", 0, errors.New("Неправильное количество параметров")
	}

	// 3. Преобразование первого элемента слайса (количество шагов) в тип int.
	countSteps, err := strconv.Atoi(info[0])
	if err != nil {
		return 0, "", 0, err
	}
	if countSteps <= 0 {
		return 0, "", 0, errors.New("Некорректное количество шагов")
	}

	// 4. Преобразование третьего элемента слайса в time.Duration.
	trainingDuration, err := time.ParseDuration(info[2])
	if err != nil {
		return 0, "", 0, err
	}
	if trainingDuration <= 0 {
		return 0, "", 0, errors.New("Ошибка временного интервала")
	}

	// 5. Вывод информации.
	trainingType := info[1]
	return countSteps, trainingType, trainingDuration, nil

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	// Вычисление дистанции.
	oneStep := height * stepLengthCoefficient // Рассчёт длины одного шага в метрах
	totalDistance := float64(steps) * oneStep // Дистанция за пройденное количество шагов в метрах
	return totalDistance / mInKm              // Возвращение пройденной дистанции в километрах

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	// 1. Проверка продолжительности duration.
	if duration <= 0 {
		return 0
	}

	// 2. Вычисление дистанции с помощью функции distance().
	total := distance(steps, height)

	// 3. Вычисление средней скорости.
	var speed float64
	if duration.Hours() > 0 {
		speed = total / duration.Hours()
	}
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	// 1. Получение значений из строки данных с помощью функции parseTraining().
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	// 2. Проверка вида тренировки и получение результата.
	var result string
	typeRun, typeWalk := "Бег", "Ходьба"
	rDistance := distance(steps, height)
	rSpeed := meanSpeed(steps, height, duration)
	runCalories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return "", err
	}
	walkCalories, err := WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return "", err
	}
	switch trainingType {
	case typeRun:
		result = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калроий: %.2f", trainingType, duration.Hours(), rDistance, rSpeed, runCalories)
	case typeWalk:
		result = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калроий: %.2f", trainingType, duration.Hours(), rDistance, rSpeed, walkCalories)
	default:
		result = "неизвестный тип тренировки"
	}

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	// 1. Проверка входных параметров на корректность.
	if (steps <= 0) || (weight <= 0) || (height <= 0) || (duration <= 0) {
		return 0, errors.New("Параметры функции RunningSpentCalories() указаны некорректно")
	}

	// 2. Рассчёт средней скорости с помощью функции meanSpeed().
	speed := meanSpeed(steps, height, duration)

	// 3. Рассчёт калорий.
	timeInM := duration.Minutes()                   // Перевод продолжительности тренировки в минуты
	calories := (weight * speed * timeInM) / minInH // Формуля для рассчёта калорий
	return calories, nil                            // Вывод значения при пробежке
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	// 1. Проверка входных параметров на корректность.
	if (steps <= 0) || (weight <= 0) || (height <= 0) || (duration <= 0) {
		return 0, errors.New("Параметры функции WalkingSpentCalories() указаны некорректно")
	}

	// 2. Рассчёт средней скорости с помощью функции meanSpeed().
	speed := meanSpeed(steps, height, duration)

	// 3. Рассчёт калорий.
	timeInM := duration.Minutes()                     // Перевод продолжительности тренировки в минуты
	calories := (weight * speed * timeInM) / minInH   // Формуля для рассчёта калорий
	return calories * walkingCaloriesCoefficient, nil // Вывод значения при ходьбе

}
