package paggr

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

func Mdn(fl validator.FieldLevel) bool {
	m := fl.Field().String()
	if _, err := strconv.ParseFloat(m, 64); err != nil {
		return false
	}
	if len(m) < 5 || len(m) > 16 {
		return false
	} else if !strings.HasPrefix(m, "+6288") && !strings.HasPrefix(m, "6288") &&
		!strings.HasPrefix(m, "088") && !strings.HasPrefix(m, "88") {
		return false
	}
	return true
}

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("mdn", Mdn)
	return validate
}

func NormalizeMDN(mdn string) string {
	if _, err := strconv.ParseFloat(mdn, 64); err != nil {
		return ""
	} else if len(mdn) > 15 {
		return ""
	}

	if strings.HasPrefix(mdn, "+6288") {
		return strings.Replace(mdn, "+62", "62", 1)
	} else if strings.HasPrefix(mdn, "088") && len(mdn) >= 8 && len(mdn) <= 13 {
		return strings.Replace(mdn, "0", "62", 1)
	} else if strings.HasPrefix(mdn, "88") && len(mdn) >= 7 && len(mdn) <= 12 {
		return fmt.Sprintf("62%s", mdn)
	} else if strings.HasPrefix(mdn, "6288") && len(mdn) >= 8 && len(mdn) <= 15 {
		return mdn
	}
	return ""
}
