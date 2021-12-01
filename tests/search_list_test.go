package tests

import (
	"strings"
	
	chk "gopkg.in/check.v1"
	
	"github.com/Virepri/adventofcode-2021/util"
)

type SearchListTestSuite struct {}
var _ = chk.Suite(&SearchListTestSuite{})

func (l *SearchListTestSuite) TestInsert(c *chk.C) {
	list := util.SearchList{}
	list.Insert("a", util.LinkedSearchFront)
	list.Insert("b", util.LinkedSearchFront)
	list.Insert("c", util.LinkedSearchFront)
	list.Insert("d", util.LinkedSearchFront)
	
	c.Assert(list.Len(), chk.Equals, uint(4))
	shouldContain := strings.Split("abcd", "")
	it := list.Iterator(true)
	
	for k,v := range shouldContain {
		c.Assert(it.Value(), chk.Equals, v)
		c.Assert(it.Next(), chk.Equals, k < len(shouldContain) - 1)
	}
}

func (l *SearchListTestSuite) TestSearch(c *chk.C) {
	list := util.SearchList{}
	list.Insert("a", util.LinkedSearchFront)
	list.Insert("b", util.LinkedSearchBack)
	list.Insert("c", util.LinkedSearchBack)
	list.Insert("d", util.LinkedSearchBack)
	
	matches, mIdx := list.Search(util.LinkedSearchFront) // test front
	c.Assert(len(matches), chk.Equals, 1)
	c.Assert(len(mIdx), chk.Equals, 1)
	c.Assert(matches[0], chk.Equals, "a")
	c.Assert(mIdx[0], chk.Equals, uint(0))
	
	matches, mIdx = list.Search(util.LinkedSearchBack) // test back
	c.Assert(len(matches), chk.Equals, 1)
	c.Assert(len(mIdx), chk.Equals, 1)
	c.Assert(matches[0], chk.Equals, "d")
	c.Assert(mIdx[0], chk.Equals, uint(3))
	
	matchFunc := util.CurryLinkedSearchEq("c", util.SearchMatchNoContinue)
	matches, mIdx = list.Search(matchFunc)
	c.Assert(len(matches), chk.Equals, 1)
	c.Assert(len(mIdx), chk.Equals, 1)
	c.Assert(matches[0], chk.Equals, "c")
	c.Assert(mIdx[0], chk.Equals, uint(2))
	
	matchFunc = util.CurryLinkedSearchIndex(1, util.SearchContinue)
	matches, mIdx = list.Search(matchFunc)
	c.Assert(len(matches), chk.Equals, 1)
	c.Assert(len(mIdx), chk.Equals, 1)
	c.Assert(matches[0], chk.Equals, "b")
	c.Assert(mIdx[0], chk.Equals, uint(1))
}

func (l *SearchListTestSuite) TestDelete(c *chk.C) {
	list := util.SearchList{}
	list.Insert("a", util.LinkedSearchFront)
	list.Insert("b", util.LinkedSearchFront)
	list.Insert("c", util.LinkedSearchFront)
	list.Insert("d", util.LinkedSearchFront)
	c.Assert(list.Len(), chk.Equals, uint(4))
	
	list.Remove(util.CurryLinkedSearchEq("b", util.SearchMatchNoContinue))
	c.Assert(list.Len(), chk.Equals, uint(3))
	
	shouldContain := strings.Split("acd", "")
	it := list.Iterator(true)
	
	for k,v := range shouldContain {
		c.Assert(it.Value(), chk.Equals, v)
		c.Assert(it.Next(), chk.Equals, k < len(shouldContain) - 1)
	}
	
	list.Remove(func(next interface{}) util.LinkedSearchStatus {
		if strings.Compare("b", next.(string)) >= 0 {
			return util.SearchMatch
		}
		
		return util.SearchContinue
	})
	
	shouldContain = strings.Split("a", "")
	it = list.Iterator(true)
	
	for k,v := range shouldContain {
		c.Assert(it.Value(), chk.Equals, v)
		c.Assert(it.Next(), chk.Equals, k < len(shouldContain) - 1)
	}
}
