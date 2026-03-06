package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var supportedTypes = map[string]bool{
	"Ходьба": true,
	"Бег":    true,
}

type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат строки: ожидается 3 части, разделенных запятой")
	}

	stepsStr := strings.TrimSpace(parts[0])
	trainingType := strings.TrimSpace(parts[1])
	durationStr := strings.TrimSpace(parts[2])

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

	t.Steps = steps
	t.TrainingType = trainingType // Может быть любым с точки зрения парсинга
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 {
		return "", errors.New("количество шагов должно быть положительным")
	}
	if t.Duration <= 0 {
		return "", errors.New("продолжительность должна быть положительной")
	}
	if t.Height <= 0 {
		return "", errors.New("рост должен быть положительным")
	}
	if t.Weight <= 0 {
		return "", errors.New("вес должен быть положительным")
	}

	if !supportedTypes[t.TrainingType] {
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories,
	), nil
}
