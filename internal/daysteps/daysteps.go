package daysteps

import (
	"errors"
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
	newSlice := strings.Split(data, ",")
	if len(newSlice) != 2 {
		return 0, 0, errors.New("Неверное количество аргументов")
	}
	steps, err := strconv.Atoi(newSlice[0])
	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, errors.New("неверные шаги")
	}

	timeDuration, err := time.ParseDuration(newSlice[1])
	if err != nil {
		return 0, 0, err
	}
	if timeDuration <= 0 {
		return 0, 0, errors.New("неверное время")
	}

	return steps, timeDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, timeDuration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	//fmt.Println(timeDuration)
	distance := (float64(steps) * stepLength) / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, timeDuration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
