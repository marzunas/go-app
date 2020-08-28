package conf

import (
	"testing"
)

func TestGetConfig_success(t *testing.T) {
	var cfg = GetConfig()
	if cfg.Port == 0 || cfg.HashWaitPeriod == 0 {
		t.Errorf("incorrect configuration")
	}
}

func TestGetConfig_singleton(t *testing.T) {
	var cfg1 = GetConfig()
	var cfg2 = GetConfig()
	cfg1.Port = 9090
	if cfg1.Port != cfg2.Port {
		t.Errorf("multiple config instances were created")
	}

}
