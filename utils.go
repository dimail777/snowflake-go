package snowflake

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

func nowEpochTime() int64 {
	return timeMax & (time.Now().UnixNano()/int64(time.Millisecond) - snowflakeEpoch)
}

func extractMachineId(hostname string) (int64, error) {
	pod := regexp.MustCompile("^([0-9a-zA-Z_-]+)-([0-9]+)$")
	m := pod.FindStringSubmatch(hostname)
	if len(m) == 0 {
		return 0, errors.New("wrong format of replica-set name - " + hostname)
	}
	number := m[2]
	if number == "" {
		return 0, errors.New("wrong format of replica-set name - " + hostname)
	}
	machineId, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return 0, errors.New("wrong format of replica-set name - " + hostname)
	}
	return machineId, nil
}

const timeBits int = 41
const machineIdBits int = 10
const sequenceNoBits int = 12

const machineIdShift = sequenceNoBits
const timeShift = sequenceNoBits + machineIdBits

const timeMax int64 = (1 << timeBits) - 1
const machineIdMax int64 = (1 << machineIdBits) - 1
const sequenceNoMax int64 = (1 << sequenceNoBits) - 1

const timeChipTolerance int64 = 100
const snowflakeEpoch int64 = 1672531200000 // 2023-01-01 00:00:00.000 UTC
