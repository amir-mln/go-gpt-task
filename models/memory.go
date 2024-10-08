package models

import (
	"fmt"
	"strings"
)

type Memory struct {
	Type      string `json:"type" required:"true" description:"wether the memory is DDR4, DDR5 and so on"`
	Capacity  string `json:"capacity" required:"true" description:"the total capacity of the memory"`
	Frequency string `json:"frequency" required:"false" description:"the frequency of the memory, 6400 MHz"`
}

func (mem *Memory) Validate() error {
	if strings.Trim(mem.Type, " ") == "" || strings.Trim(mem.Capacity, " ") == "" {
		return fmt.Errorf("insufficient memory data, Type: %q and Capacity:%q", mem.Type, mem.Capacity)
	}

	return nil
}
