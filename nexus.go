package nexus

import (
	"errors"
	"net"
	"sync"
	"time"
)

// Nexus is a distributed unique Id generator.
type Nexus struct {
	mutex     *sync.Mutex
	start     uint64
	elapsed   uint64
	sequence  uint64
	machineId uint64
}

func NewNexus(opts ...Option) *Nexus {
	n := &Nexus{
		mutex:    new(sync.Mutex),
		sequence: 1<<BitLenSequence - 1,
	}
	for _, opt := range opts {
		opt(n)
	}
	if n.start == 0 {
		WithStartTime(time.Time{})(n)
	}
	if n.machineId == 0 {
		ip, err := getPrivateIPv4()
		if err != nil {
			panic(err)
		}
		WithIPv4(ip)(n)
	}
	return n
}

// NextId generates a next unique ID.
// After the Sonyflake time overflows, NextID returns an error.
func (n *Nexus) NextId() (*Part, error) {
	const maskSequence = 1<<BitLenSequence - 1

	n.mutex.Lock()
	defer n.mutex.Unlock()

	current := n.getCurrentElapsed(n.start)
	if n.elapsed < current {
		n.elapsed = current
		n.sequence = 0
	} else { // sf.elapsedTime >= current
		n.sequence = (n.sequence + 1) & maskSequence
		if n.sequence == 0 {
			n.elapsed++
			overtime := n.elapsed - current
			time.Sleep(n.getSleepTime((overtime)))
		}
	}
	id, err := n.toId()
	if err != nil {
		return nil, err
	}
	return getPart(id), nil
}

func (n *Nexus) Recover(p *Part) {
	putPart(p)
}

func (n *Nexus) setMachineId(ip net.IP) {
	n.machineId = uint64(ip[2])<<8 + uint64(ip[3])
}

func (n *Nexus) getCurrentElapsed(startTime uint64) uint64 {
	return n.toNexusTime(time.Now()) - startTime
}

func (n *Nexus) getSleepTime(overtime uint64) time.Duration {
	return time.Duration(overtime*TimeUnit) -
		time.Duration(time.Now().UTC().UnixNano()%TimeUnit)
}

func (n *Nexus) toNexusTime(t time.Time) uint64 {
	return uint64(t.UTC().UnixNano() / TimeUnit)
}

func (n *Nexus) toId() (uint64, error) {
	if n.elapsed >= 1<<BitLenTime {
		return 0, errors.New("over the time limit")
	}
	return uint64(n.elapsed)<<(BitLenSequence+BitLenMachineId) |
		uint64(n.sequence)<<BitLenMachineId |
		uint64(n.machineId), nil
}
