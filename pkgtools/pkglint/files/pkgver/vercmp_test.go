package pkgver

import (
	"gopkg.in/check.v1"
	"testing"
)

type Suite struct{}

func Test(t *testing.T) {
	check.Suite(&Suite{})
	check.TestingT(t)
}

func (s *Suite) Test_newVersion(c *check.C) {
	c.Check(newVersion("5.0"), check.DeepEquals, &version{[]int{5, 0, 0}, 0})
	c.Check(newVersion("5.0nb5"), check.DeepEquals, &version{[]int{5, 0, 0}, 5})
	c.Check(newVersion("0.0.1-SNAPSHOT"), check.DeepEquals, &version{[]int{0, 0, 0, 0, 1, 19, 14, 1, 16, 19, 8, 15, 20}, 0})
	c.Check(newVersion("1.0alpha3"), check.DeepEquals, &version{[]int{1, 0, 0, -3, 3}, 0})
	c.Check(newVersion("1_0alpha3"), check.DeepEquals, &version{[]int{1, 0, 0, -3, 3}, 0})
	c.Check(newVersion("2.5beta"), check.DeepEquals, &version{[]int{2, 0, 5, -2}, 0})
	c.Check(newVersion("20151110"), check.DeepEquals, &version{[]int{20151110}, 0})
	c.Check(newVersion("0"), check.DeepEquals, &version{[]int{0}, 0})
	c.Check(newVersion("nb1"), check.DeepEquals, &version{nil, 1})
	c.Check(newVersion("1.0.1a"), check.DeepEquals, &version{[]int{1, 0, 0, 0, 1, 1}, 0})
	c.Check(newVersion("1.0.1z"), check.DeepEquals, &version{[]int{1, 0, 0, 0, 1, 26}, 0})
	c.Check(newVersion("0pre20160620"), check.DeepEquals, &version{[]int{0, -1, 20160620}, 0})
	c.Check(newVersion("3.5.DEV1710"), check.DeepEquals, &version{[]int{3, 0, 5, 0, 4, 5, 22, 1710}, 0})
}

func (s *Suite) Test_Compare(c *check.C) {
	var versions = [][]string{
		{"0pre20160620"},
		{"0"},
		{"nb1"},
		{"0.0.1-SNAPSHOT"},
		{"1.0alpha"},
		{"1.0alpha3"},
		{"1", "1.0", "1.0.0"},
		{"1.0nb1"},
		{"1.0nb2", "1_0nb2"},
		{"1.0.1a", "1.0.a1", "1.0.aa"},
		{"1.0.1z"},
		{"1.0.11", "1.0.k"},
		{"2.0pre", "2.0rc"},
		{"2.0", "2.0pl"},
		{"2.0.1nb4"},
		{"2.0.1nb17"},
		{"2.5beta"},
		{"5.0"},
		{"5.0nb5"},
		{"5.5", "5.005"},
		{"20151110"},
	}

	checkVersion := func(i int, iversion string, j int, jversion string) {
		actual := Compare(iversion, jversion)
		switch {
		case i < j && !(actual < 0):
			c.Check([]interface{}{i, iversion, j, jversion, actual}, check.DeepEquals, []interface{}{i, iversion, j, jversion, "<0"})
		case i == j && !(actual == 0):
			c.Check([]interface{}{i, iversion, j, jversion, actual}, check.DeepEquals, []interface{}{i, iversion, j, jversion, "==0"})
		case i > j && !(actual > 0):
			c.Check([]interface{}{i, iversion, j, jversion, actual}, check.DeepEquals, []interface{}{i, iversion, j, jversion, ">0"})
		}
	}

	for i, iversions := range versions {
		for j, jversions := range versions {
			for _, iversion := range iversions {
				for _, jversion := range jversions {
					checkVersion(i, iversion, j, jversion)
				}
			}
		}
	}
}
