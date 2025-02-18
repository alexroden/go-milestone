package gomilestone

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type MilestoneSuite struct {
	suite.Suite
}

func (s *MilestoneSuite) IntiateSteps() {
	Start()

	time.Sleep(1 * time.Second)
	Step()

	time.Sleep(1 * time.Second)
	Step()
}

func (s *MilestoneSuite) TestSteps() {
	s.IntiateSteps()

	m := GetInstance()

	s.NotEmpty(m.reports)
	s.Equal([]*report{
		{Time: 1},
		{Time: 2},
	}, m.reports)
}

func (s *MilestoneSuite) TestReport() {
	s.IntiateSteps()

	r, err := Report()
	s.NoError(err)

	s.IsType([]*report{}, r)
	s.Len(r, 2)
	s.Equal([]*report{
		{Step: 1, Time: 1},
		{Step: 2, Time: 2},
	}, r)
}

func (s *MilestoneSuite) TestMessage() {
	message := "Hello there!"

	Start()

	time.Sleep(1 * time.Second)
	Step(WithMessage(message))

	time.Sleep(1 * time.Second)
	Step()

	r, err := Report()
	s.NoError(err)

	s.IsType([]*report{}, r)
	s.Equal([]*report{
		{Step: 1, Time: 1, Message: message},
		{Step: 2, Time: 2},
	}, r)
}

func (s *MilestoneSuite) TestEnd() {
	s.IntiateSteps()
	s.Len(instance.reports, 2)

	End()
	s.Empty(instance.reports)
	s.True(instance.init.IsZero())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMilestoneSuite(t *testing.T) {
	suite.Run(t, new(MilestoneSuite))
}
