package suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HAProxySuite struct {
	*RodSuite
}

func NewHAProxySuite() *HAProxySuite {
	return &HAProxySuite{
		RodSuite: NewRodSuite(haproxySuiteName),
	}
}

func (s *HAProxySuite) TestOneFactorScenario() {
	suite.Run(s.T(), NewOneFactorScenario())
}

func (s *HAProxySuite) TestTwoFactorTOTPScenario() {
	suite.Run(s.T(), NewTwoFactorTOTPScenario())
}

func (s *HAProxySuite) TestCustomHeaders() {
	suite.Run(s.T(), NewCustomHeadersScenario())
}

func TestHAProxySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping suite test in short mode")
	}

	suite.Run(t, NewHAProxySuite())
}
