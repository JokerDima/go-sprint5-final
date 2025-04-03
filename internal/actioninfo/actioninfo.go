package actioninfo

import (
	"fmt"
	_ "go-sprint5-final/internal/daysteps"
	_ "go-sprint5-final/internal/trainings"
)

// Интерфейс DataParser
type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

// Функция Info()
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		actionInfo, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(actionInfo)
	}
}
