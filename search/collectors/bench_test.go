package collectors

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/blevesearch/bleve/search"
	"golang.org/x/net/context"
)

func benchHelper(numOfMatches int, collector search.Collector, b *testing.B) {
	matches := make(search.DocumentMatchCollection, 0, numOfMatches)
	for i := 0; i < numOfMatches; i++ {
		matches = append(matches, &search.DocumentMatch{
			ID:    strconv.Itoa(i),
			Score: rand.Float64(),
		})
	}

	b.ResetTimer()

	for run := 0; run < b.N; run++ {
		searcher := &stubSearcher{
			matches: matches,
		}
		err := collector.Collect(context.Background(), searcher)
		if err != nil {
			b.Fatal(err)
		}
	}
}
