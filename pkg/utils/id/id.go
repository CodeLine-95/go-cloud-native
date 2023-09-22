package id

import (
	"sync/atomic"
	"time"
)

const (
	IndexOffset     = 0
	TimestampOffset = 16
	ServerIDOffset  = 48
)

const (
	IndexBits     = 16
	TimestampBits = 32
	ServerIDBits  = 16
)

type ID uint64

type Maker interface {
	ServerID() uint32
	SetServerID(uint32)
	Make() ID
}

type idMaker struct {
	serverID uint32
	idx      uint32
	tm       int64
}

func NewMaker(serverID uint32) Maker {
	return &idMaker{
		serverID: serverID,
		idx:      0,
	}
}

func (im *idMaker) ServerID() uint32 {
	return im.serverID
}

func (im *idMaker) SetServerID(id uint32) {
	atomic.StoreUint32(&im.serverID, id)
}

func (im *idMaker) Make() ID {
	tm := time.Now().Unix()
	if atomic.SwapInt64(&im.tm, tm) != tm {
		atomic.StoreUint32(&im.idx, 0)
	}
	idx := atomic.AddUint32(&im.idx, 1)
	var id uint64 = 0
	id |= GetBits(uint64(idx), IndexBits, 0) << IndexOffset
	id |= GetBits(uint64(tm), TimestampBits, 0) << TimestampOffset
	id |= GetBits(uint64(im.serverID), ServerIDBits, 0) << ServerIDOffset
	return ID(id)
}

func Mask(bits uint64, offset uint64) uint64 {
	if bits+offset >= 64 {
		return ^((uint64(1) << offset) - 1)
	}
	return ((uint64(1) << bits) - 1) << offset
}

func GetBits(val uint64, bits uint64, offset uint64) uint64 {
	mask := Mask(bits, offset)
	return (val & mask) >> offset
}

func SetBits(oldVal uint64, bits uint64, offset uint64, val uint64) uint64 {
	return oldVal | GetBits(val, bits, 0)<<offset
}

func (id *ID) setBits(bits uint64, offset uint64, val uint64) {
	*id = ID(SetBits(uint64(*id), bits, offset, val))
}

func (id ID) getBits(bits uint64, offset uint64) uint64 {
	return GetBits(uint64(id), bits, offset)
}

func (id ID) GetIndex() uint16 {
	return uint16(id.getBits(ServerIDBits, ServerIDOffset))
}

func (id ID) GetTimestamp() uint64 {
	return id.getBits(TimestampBits, TimestampOffset)
}

func (id ID) GetServerID() uint16 {
	return uint16(id.getBits(ServerIDBits, ServerIDOffset))
}
