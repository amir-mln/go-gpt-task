package models

import (
	"fmt"
	"strings"
)

type Memory struct {
	Type     string `json:"type" description:"wether the memory is DDR4, DDR5 and so on"`
	Capacity string `json:"capacity" description:"the total capacity of the memory"`
}

func (mem *Memory) Validate() error {
	if strings.Trim(mem.Capacity, " ") == "" {
		return fmt.Errorf("insufficient memory data, Capacity:%q", mem.Capacity)
	}

	return nil
}
