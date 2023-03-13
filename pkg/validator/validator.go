package validator

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type Validator struct {
	messages []string
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Check(condition bool, message string) *Validator {
	if condition {
		v.messages = append(v.messages, message)
	}

	return v
}

func (v *Validator) Verdict() error {
	if v.messages != nil {
		return fiber.NewError(http.StatusBadRequest,
			strings.Join(v.messages, "; "))
	}

	return nil
}
