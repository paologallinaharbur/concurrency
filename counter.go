package awesomeProject

import "sync"

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
