package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go-sprint5-final/internal/personaldata"
	"go-sprint5-final/internal/spentenergy"
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
		return errors.New("no data for conversion")
	}

	dataParse := strings.Split(datastring, ",")
	if len(dataParse) != 3 {
		return errors.New("the data has not been converted correctly")
	}

	//Шаги
	steps, err := strconv.Atoi(dataParse[0])
	if err != nil {
		return err
	}
	if steps < 0 {
		return errors.New("negative value Steps")
	}
	//Сохряняем значение шагов в структуру Training
	t.Steps = steps

	//Тип тренировки
	if dataParse[1] == "Бег" {
		t.TrainingType = "Бег"
	} else if dataParse[1] == "Ходьба" {
		t.TrainingType = "Ходьба"
	} else {
		t.TrainingType = "Неизвестный тип тренировки"
		return errors.New("unknown training type")
	}

	//Длительность
	duration, err := time.ParseDuration(dataParse[2])
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("negative value Duration")
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
		return "", errors.New("zero value Duration")
	}

	meanSpeed := spentenergy.MeanSpeed(steps, duration)

	var calories float64
	switch t.TrainingType {
	case "Бег":
		calories = spentenergy.RunningSpentCalories(steps, t.Weight, duration)
	case "Ходьба":
		calories = spentenergy.WalkingSpentCalories(steps, t.Weight, t.Height, duration)
	default:
		return "", errors.New("unknown training type")
	}
	if calories == 0 {
		return "", errors.New("zero value calories")
	}

	title := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", t.TrainingType, duration.Hours(), distance, meanSpeed, calories)

	return title, nil
}
