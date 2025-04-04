package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go-sprint5-final/internal/personaldata"
	"go-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// Структура DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	if len(datastring) == 0 {
		return errors.New("no data for conversion")
	}

	dataParse := strings.Split(datastring, ",")
	if len(dataParse) != 2 {
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
	ds.Steps = steps

	//Длительность
	duration, err := time.ParseDuration(dataParse[1])
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("negative value Duration")
	}
	//Сохряняем значение длительности в структуру Training
	ds.Duration = duration

	return nil
}

// Метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	duration := ds.Duration
	if duration == 0 {
		return "", errors.New("zero value Duration")
	}

	steps := ds.Steps
	distance := spentenergy.Distance(steps)

	calories := spentenergy.WalkingSpentCalories(steps, ds.Weight, ds.Height, duration)
	if calories == 0 {
		return "", errors.New("zero value calories")
	}

	title := fmt.Sprintf("Количество шагов: %v.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, calories)

	return title, nil
}
