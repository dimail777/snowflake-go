package snowflake

import "testing"

func TestK8SMachineIdExtraction(t *testing.T) {
	machineId, err := extractMachineId("Abc_Zyz-39-123")
	if err != nil {
		t.Fatalf("snowflake error extractMachineId : %s", err.Error())
	}
	if machineId != 123 {
		t.Fatalf("snowflake error extractMachineId : %s", "machineId != 123")
	}
}

func TestSnowflake(t *testing.T) {
	snowflake, err := InitByRandom()
	if err != nil {
		t.Fatalf("snowflake error init : %s", err.Error())
	}
	var prev int64 = 0
	for i := 0; i < 10000000; i++ {
		id, err := snowflake.GetNextId()
		if err != nil {
			t.Fatalf("snowflake error next : %s", err.Error())
		}
		if id < prev {
			t.Fatalf("snowflake error unique: %d - %d", id, prev)
		}
	}
}
