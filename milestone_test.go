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
	s.True(m.reports[0].Time > 1000 && m.reports[0].Time < 2000)
	s.True(m.reports[1].Time > 2000 && m.reports[1].Time < 3000)
}

func (s *MilestoneSuite) TestReport() {
	s.IntiateSteps()

	r, err := Report()
	s.NoError(err)

	s.IsType([]*report{}, r)
	s.Len(r, 2)
	s.Equal(1, r[0].Step)
	s.True(r[0].Time > 1000 && r[0].Time < 2000)
	s.Equal(2, r[1].Step)
	s.True(r[1].Time > 2000 && r[1].Time < 3000)
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
	s.Equal(1, r[0].Step)
	s.True(r[0].Time > 1000 && r[0].Time < 2000)
	s.Equal(message, r[0].Message)
	s.Equal(2, r[1].Step)
	s.True(r[1].Time > 2000 && r[1].Time < 3000)
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
