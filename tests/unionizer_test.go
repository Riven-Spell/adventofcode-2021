package tests

import (
	chk "gopkg.in/check.v1"
	
	"github.com/Virepri/adventofcode-2021/util"
)

type UnionizerTestSuite struct {}
var _ = chk.Suite(&UnionizerTestSuite{})

func (*UnionizerTestSuite) TestAddEmpty(c *chk.C) {
	u := util.Unionizer{}

	// add 2 and 5
	u.AddItemsA([]int{2,5})

	// check there's no union
	c.Assert(len(u.GetUnion()), chk.Equals, 0)
}

func (*UnionizerTestSuite) TestNoUnion(c *chk.C) {
	u := util.Unionizer{}

	// add 2 and 5
	u.AddItemsA([]int{2,5})
	// add a non-value
	u.AddItemsB([]int{0})
	// ensure that nothing is present
	c.Assert(u.Len(), chk.Equals, 0)
}

func (*UnionizerTestSuite) TestBasicUnion(c *chk.C) {
	u := util.Unionizer{}

	u.AddItemsA([]int{5,2})
	u.AddItemsB([]int{2,8})
	c.Assert(u.Len(), chk.Equals, 1)
}

func (*UnionizerTestSuite) TestRemoveItems(c *chk.C) {
	u := util.Unionizer{}

	u.AddItemsA([]int{0,1})
	c.Assert(u.Len(), chk.Equals, 2)
	u.RemoveItems([]int{1})
	c.Assert(u.Len(), chk.Equals, 1)
}