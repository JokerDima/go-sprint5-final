package trainings

import (
	"errors"
	"fmt"
	"go-sprint5-final/internal/personaldata"
	"go-sprint5-final/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

// Структура Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	if len(datastring) == 0 {
		return errors.New("error data. No data for conversion")
	}

	dataParse := strings.Split(datastring, ",")
	if len(dataParse) != 3 {
		return errors.New("error conversion. The data has not been converted correctly")
	}

	//Шаги
	steps, err := strconv.Atoi(dataParse[0])
	if err != nil {
		return err
	}
	if steps < 0 {
		return errors.New("error data. Negative value Steps")
	}
	//Сохряняем значение шагов в структуру Training
	t.Steps = steps

	//Тип тренировки
	var trainingType string
	switch dataParse[1] {
	case "Бег":
		trainingType = "Бег"
	case "Ходьба":
		trainingType = "Ходьба"
	default:
		trainingType = "Неизвестный тип тренировки"
		return errors.New("error data. Unknown training type")
	}
	//Сохряняем значение тренировки в структуру Training
	t.TrainingType = trainingType

	//Длительность
	duration, err := time.ParseDuration(dataParse[2])
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("error data. Negative value Duration")
	}
	//Сохряняем значение длительности в структуру Training
	t.Duration = duration

	return nil
}

// Метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	steps := t.Steps
	distance := spentenergy.Distance(steps)
	duration := t.Duration
	if duration == 0 {
		return "", errors.New("error data. Zero value Duration")
	}

	meanSpeed := spentenergy.MeanSpeed(steps, duration)
	trainingType := t.TrainingType
	if trainingType == "Неизвестный тип тренировки" {
		return "", errors.New("error type. Unknown training type")
	}

	var calories float64
	switch trainingType {
	case "Бег":
		calories = spentenergy.RunningSpentCalories(steps, t.Weight, duration)
	case "Ходьба":
		calories = spentenergy.WalkingSpentCalories(steps, t.Weight, t.Height, duration)
	}
	if calories == 0 {
		return "", errors.New("error data. Zero value calories")
	}

	title := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", trainingType, duration.Hours(), distance, meanSpeed, calories)

	return title, nil
}
