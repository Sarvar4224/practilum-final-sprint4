package spentcalories

import (
	"errors"
	"fmt"
	"log"
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
	newSlice := strings.Split(data, ",")
	if len(newSlice) != 3 {
		return 0, "", 0, errors.New("Неверное количество аргументов")
	}
	steps, err := strconv.Atoi(newSlice[0])
	if err != nil {
		return 0, newSlice[1], 0, err
	}

	if steps <= 0 {
		return 0, "", 0, errors.New("неверные шаги")
	}

	timeDuration, err := time.ParseDuration(newSlice[2])
	if err != nil {
		return 0, newSlice[1], 0, err
	}
	if timeDuration <= 0 {
		return 0, "", 0, errors.New("неверное время")
	}

	return steps, newSlice[1], timeDuration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	userDistance := (height * stepLengthCoefficient * float64(steps)) / mInKm
	return userDistance

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	userDistance := distance(steps, height)
	timeInHours := duration.Hours()
	return userDistance / timeInHours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activityType, timeDuration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	speed := meanSpeed(steps, height, timeDuration)
	distance := distance(steps, height)
	switch activityType {
	case "Бег":
		spentCalories, _ := RunningSpentCalories(steps, weight, height, timeDuration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, timeDuration.Hours(), distance, speed, spentCalories), nil
	case "Ходьба":
		spentCalories, _ := WalkingSpentCalories(steps, weight, height, timeDuration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, timeDuration.Hours(), distance, speed, spentCalories), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("Шагов мало")
	}
	if duration <= 0 || weight <= 0 || height <= 0 || steps <= 0 {
		return 0, errors.New("Неверные параметры")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("Шагов мало")
	}
	if duration <= 0 || weight <= 0 || height <= 0 || steps <= 0 {
		return 0, errors.New("Неверные параметры")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient, nil
}
