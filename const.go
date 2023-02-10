package nexus

import "time"

const (
	BaseBitLen      = 63                                       // base bit length
	BitLenTime      = 39                                       // bit length of time
	BitLenSequence  = 8                                        // bit length of sequence number
	BitLenMachineId = BaseBitLen - BitLenTime - BitLenSequence // bit length of machine id
)

const (
	// TimeUnit nsec, i.e. 10 msec
	TimeUnit      = 1e7
	MaskMachineId = 1<<BitLenMachineId - 1
	MaskSequence  = (1<<BitLenSequence - 1) << BitLenMachineId
)

type Config struct {
	StartTime      time.Time
	MachineId      func() (uint64, error)
	CheckMachineId func(uint64) bool
}
