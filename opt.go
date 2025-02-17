package gomilestone

type Setter func(*Opt)

type Opt struct {
	message string
}

func WithMessage(message string) Setter {
	return func(o *Opt) {
		o.message = message
	}
}
