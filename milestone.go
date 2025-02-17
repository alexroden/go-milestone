package gomilestone

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	instance *Milestone
	once     sync.Once
	now      = time.Now
)

type Milestone struct {
	init    time.Time
	reports []*report
	*Opt
}

func Start() {
	once.Do(func() {
		instance = &Milestone{
			Opt: &Opt{},
		}
	})

	if instance.init.IsZero() {
		instance.init = now()
	}
}

func Step(opts ...Setter) error {
	if instance == nil || instance.init.IsZero() {
		return errors.New("Milestone not started. Call Start() first.")
	}

	for _, o := range opts {
		o(instance.Opt)
	}

	n := now()
	duration := n.Sub(instance.init)
	instance.reports = append(instance.reports, &report{
		Time:    int(duration.Seconds()),
		Message: instance.message,
	})

	instance.message = ""

	return nil
}

func Report() ([]*report, error) {
	if instance == nil || len(instance.reports) == 0 {
		fmt.Println("Milestone has nothing to report.")
		return nil, errors.New("Milestone has nothing to report.")
	}

	result := []*report{}
	for i, v := range instance.reports {
		r := v
		v.Step = i + 1

		result = append(result, r)
	}

	return result, nil
}

func GetInstance() *Milestone {
	return instance
}
