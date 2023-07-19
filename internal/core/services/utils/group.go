package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ConvertStringSliceToUUIDSlice(stringSlice []string) ([]uuid.UUID, error) {
	uuidSlice := make([]uuid.UUID, len(stringSlice))

	for i, str := range stringSlice {
		u, err := uuid.Parse(str)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("failed to parse UUID at index %d: %v", i, err))
		}
		uuidSlice[i] = u
	}

	return uuidSlice, nil
}
