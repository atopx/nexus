package nexus

import (
	"sync"
)

type Part struct {
	Id        uint64 `json:"id"`
	Msb       uint64 `json:"msb"`
	Elapsed   uint64 `json:"elapsed"`
	Sequence  uint64 `json:"sequence"`
	MachineId uint64 `json:"machine_id"`
}

func (p *Part) init(id uint64) {
	p.Id = id
	p.Msb = p.Id >> BaseBitLen
	p.Elapsed = p.Id >> (BitLenSequence + BitLenMachineId)
	p.Sequence = p.Id & MaskSequence >> BitLenMachineId
	p.MachineId = p.Id & MaskMachineId
}

var pool = sync.Pool{New: func() any {
	return new(Part)
}}

func getPart(id uint64) *Part {
	part := pool.Get().(*Part)
	part.init(id)
	return part
}

func putPart(part *Part) {
	part.Id = 0
	part.Msb = 0
	part.Elapsed = 0
	part.Sequence = 0
	part.MachineId = 0
	pool.Put(part)
}
