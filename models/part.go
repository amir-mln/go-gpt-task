package models

import (
	"fmt"
	"strings"
)

type Part struct {
	Brand string `json:"brand" required:"true" description:"the brand of the manufacturer"`
	Model string `json:"model" required:"true" description:"the model of the part"`
}

func (pa *Part) Validate() error {
	if strings.Trim(pa.Brand, " ") == "" || strings.Trim(pa.Model, " ") == "" {
		return fmt.Errorf("insufficient part data with Model as %q and Brand as %q", pa.Model, pa.Brand)
	}

	return nil
}
