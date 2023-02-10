package nexus_test

import (
	"sync"
	"testing"
	"time"

	"github.com/atopx/nexus"
)

func TestNexus(t *testing.T) {
	testCount := 20
	n := nexus.NewNexus(nexus.WithStartTime(time.Now()))
	wg := new(sync.WaitGroup)
	wg.Add(testCount)
	for i := 0; i < testCount; i++ {
		go func(n *nexus.Nexus, seq int, wg *sync.WaitGroup) {
			part, _ := n.NextId()
			defer n.Recover(part)
			t.Logf("no.%d nexus part: %+v", seq, part)
			wg.Done()
		}(n, i, wg)
	}
	wg.Wait()
}
