package models

import "fmt"

type Processor struct {
	Part
	Cores     uint   `json:"cores" required:"true" description:"the total  number of cores of the processor"`
	Threads   uint   `json:"threads" required:"true" description:"the total number of threads of the processor"`
	Frequency string `json:"frequency" required:"false" description:"the clock speed of the processor"`
}

func (proc *Processor) Validate() error {
	if err := proc.Part.Validate(); err != nil {
		return fmt.Errorf("for processor part got: %w", err)
	}

	if proc.Cores == 0 || proc.Threads == 0 {
		return fmt.Errorf(
			"number of Cores and Threads must be specified. Cores:%d, Threads:%d",
			proc.Cores,
			proc.Threads,
		)
	}

	return nil
}
