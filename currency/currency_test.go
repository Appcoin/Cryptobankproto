package currency

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type CurrencySuite struct{}

var _ = Suite(&CurrencySuite{})

func (s *CurrencySuite) TestFromString(c *C) {
	cur := FromString("BTC")
	c.Assert(cur.Symbol, Equals, "BTC")
}

func (s *CurrencySuite) TestPairConcat(c *C) {
	pair := NewPair("BTC", "DCR")
	c.Assert(pair.First.Symbol, Equals, "BTC")
	c.Assert(pair.Second.Symbol, Equals, "DCR")
	c.Assert(pair.Concat("_"), Equals, "BTC_DCR")
	c.Assert(pair.Concat(""), Equals, "BTCDCR")
}
