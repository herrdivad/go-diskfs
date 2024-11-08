package partition

import (
	"github.com/herrdivad/go-diskfs/partition/part"
	"github.com/herrdivad/go-diskfs/util"
)

// Table reference to a partitioning table on disk
type Table interface {
	Type() string
	Write(util.File, int64) error
	GetPartitions() []part.Partition
	Repair(diskSize uint64) error
	Verify(f util.File, diskSize uint64) error
	UUID() string
}
