package snowflake

import (
	"errors"
	"math/rand"
	"os"
)

type Snowflake interface {
	GetNextId() (int64, error)
}

func InitByMachineId(machineId int64) (Snowflake, error) {
	return initWithCAS(machineId)
}

func InitByRandom() (Snowflake, error) {
	machineId := rand.Int63n(machineIdMax + 1)
	return initWithCAS(machineId)
}

func InitByK8sStatefulSet() (Snowflake, error) {
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		return nil, errors.New("HOSTNAME sys env is not defined")
	}
	machineId, err := extractMachineId(hostname)
	if err != nil {
		return nil, err
	}
	return initWithCAS(machineId)
}
