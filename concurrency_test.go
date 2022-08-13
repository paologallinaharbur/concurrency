package awesomeProject

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

const maxCounter = 10

type data struct {
	m       sync.Mutex
	counter int
}

func (d *data) isSafeToIncrease() bool {
	d.m.Lock()
	defer d.m.Unlock()

	return d.counter < maxCounter
}

func (d *data) increase() {
	d.m.Lock()
	defer d.m.Unlock()

	d.counter += 1
}

func (d *data) getCounter() int {
	d.m.Lock()
	defer d.m.Unlock()

	return d.counter
}

func TestExample1000Times(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprint(i), TestExample)
	}
}

func TestExample(t *testing.T) {

	d := &data{
		m: sync.Mutex{},
	}

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go increaseNTimes(wg, d, 10)
	go increaseNTimes(wg, d, 10)
	wg.Wait()

	require.Equal(t, maxCounter, d.getCounter(), "we should have reached the maximum allowed")
}

func increaseNTimes(wg *sync.WaitGroup, d *data, n int) {
	defer wg.Done()

	for i := 0; i < n; i++ {
		if d.isSafeToIncrease() {
			d.increase()
		}

	}
	return
}
