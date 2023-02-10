package nexus

import (
	"net"
	"time"
)

type Option func(n *Nexus)

func WithStartTime(start time.Time) Option {
	return func(n *Nexus) {
		// boundary value handling for time
		now := time.Now()
		if start.After(now) {
			start = now
		}
		if start.IsZero() {
			n.start = 1
		}
		n.start = n.toNexusTime(start)
	}
}

func WithIPv4(ip net.IP) Option {
	return func(n *Nexus) {
		n.setMachineId(ip)
	}
}
