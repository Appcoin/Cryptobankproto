package order

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type OrderSuite struct{}

var _ = Suite(&OrderSuite{})

func (s *OrderSuite) TestTotalPrice(c *C) {
	// c.Assert(cur.Symbol, Equals, "BTC")
}
