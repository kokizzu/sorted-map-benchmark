package googleBtree_test

import (
	"testing"

	"github.com/google/btree"
	"modernc.org/benchmark-ordered-map/fixture"
)

type googItem fixture.Item

func (a googItem) Less(b btree.Item) bool {
	return a.Key < b.(googItem).Key
}

const btreeDegree = 32

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := btree.New(btreeDegree)
		for i := 0; i < len(fixture.TestData); i++ {
			tree.ReplaceOrInsert(googItem(fixture.TestData[i]))
		}
	}
}

func BenchmarkSortedInsert_ReplaceOrInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := btree.New(btreeDegree)
		for i := 0; i < len(fixture.SortedTestData); i++ {
			tree.ReplaceOrInsert(googItem(fixture.SortedTestData[i]))
		}
	}
}

var (
	ksink fixture.Key
	vsink fixture.Value
)

func BenchmarkIterate(b *testing.B) {
	tree := btree.New(btreeDegree)
	for i := 0; i < len(fixture.TestData); i++ {
		tree.ReplaceOrInsert(googItem(fixture.TestData[i]))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		tree.Ascend(func(i btree.Item) bool {
			ksink = i.(googItem).Key
			vsink = i.(googItem).Value
			return true
		})
	}
}
