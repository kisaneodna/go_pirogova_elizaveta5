package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(data string) error {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат строки: ожидается 2 части, разделенных запятой")
	}

	stepsPart := parts[0]
	durationPart := parts[1]

	if strings.TrimSpace(stepsPart) != stepsPart {
		return fmt.Errorf("неверный формат шагов: '%s' содержит лишние пробелы", stepsPart)
	}

	stepsStr := strings.TrimSpace(stepsPart)
	durationStr := strings.TrimSpace(durationPart)

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("неверный формат шагов: %s", stepsStr)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным: %d", steps)
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("неверный формат продолжительности: %s", durationStr)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной: %v", duration)
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", errors.New("количество шагов должно быть положительным")
	}
	if ds.Duration <= 0 {
		return "", errors.New("продолжительность должна быть положительной")
	}
	if ds.Height <= 0 {
		return "", errors.New("рост должен быть положительным")
	}
	if ds.Weight <= 0 {
		return "", errors.New("вес должен быть положительным")
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчёта калорий: %w", err)
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	), nil
}
