package spentenergy

import (
	"errors"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be positive")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be positive")
	}
	if height <= 0 {
		return 0, errors.New("height must be positive")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}

	speed := MeanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be positive")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be positive")
	}
	if height <= 0 {
		return 0, errors.New("height must be positive")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}

	speed := MeanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil
}
