package go_oauth2_server

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {TestingT(t)}

type GenerateSuite struct {}

var _  = Suite(&GenerateSuite{})


func (suite *GenerateSuite) TestGenerator_Random_ExpectNotEqual(c *C) {
	generator := &Generator{10}

	val1 := generator.Generate()
	val2 := generator.Generate()

	c.Assert(val1 == val2, Equals, false)
}

func (suite *GenerateSuite) TestGenerate_Length_ExpectCorrect (c *C) {
	generator := &Generator{10}

	val := generator.Generate()

	c.Assert(val, HasLen, 10)
}