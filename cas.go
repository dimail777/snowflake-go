package snowflake

import (
	"errors"
	"strconv"
	"sync/atomic"
	"unsafe"
)

type snowflake struct {
	machineId int64
	snapshot  unsafe.Pointer
}

type snapshot struct {
	sequenceNo int64
	unixTime   int64
}

func initWithCAS(machineId int64) (*snowflake, error) {
	if machineId < 0 || machineId > machineIdMax {
		return nil, errors.New("machineId out of 0-" + strconv.FormatInt(machineIdMax, 10))
	}
	sh := &snapshot{sequenceNo: -1, unixTime: nowEpochTime()}
	return &snowflake{machineId: machineId, snapshot: unsafe.Pointer(sh)}, nil
}

func (s *snowflake) GetNextId() (int64, error) {
	nextSnapshot, err := s.nextSnapshot()
	if err != nil {
		return 0, err
	}
	return (nextSnapshot.unixTime << timeShift) | (s.machineId << machineIdShift) | nextSnapshot.sequenceNo, nil
}

func (s *snowflake) nextSnapshot() (*snapshot, error) {
	for {
		pointer := atomic.LoadPointer(&s.snapshot)
		newTime := nowEpochTime()
		nextSnapshot, err := (*snapshot)(pointer).nextSnapshot(newTime)
		if err != nil {
			return nil, err
		}
		swapped := false
		if nextSnapshot != nil {
			swapped = atomic.CompareAndSwapPointer(&s.snapshot, pointer, unsafe.Pointer(nextSnapshot))
		}
		if swapped {
			return nextSnapshot, err
		}
	}
}

func (s *snapshot) nextSnapshot(newTime int64) (*snapshot, error) {
	if newTime+timeChipTolerance < s.unixTime {
		return nil, errors.New("clock error")
	}
	if newTime > s.unixTime {
		return &snapshot{sequenceNo: 0, unixTime: newTime}, nil
	} else if sequenceNoMax == s.sequenceNo {
		return nil, nil
	} else {
		return &snapshot{sequenceNo: s.sequenceNo + 1, unixTime: s.unixTime}, nil
	}
}
