package authorization

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ServiceAuthorizationSuite struct{}

var _ = Suite(&ServiceAuthorizationSuite{})

func (suite *ServiceAuthorizationSuite) TestAddGrantType_ExpectSuccess(c *C) {

}
