// Package billcalc provides billcalc executable, its cfg with a matching struct
package billcalc

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Cfg contains billcalc executable's cfg values
type Cfg struct {
	Water struct {
		LowRateCommission  float64
		HighRateCommission float64
	}
	Electricity struct {
		CommissionRate float64
	}
}

// GetCfg gets the cfg from the path that's written in the env var "billcalcCfgPath".
// Panics if cannot find/read the cfg file, or if unmarshaling it fails.
func GetCfg() *Cfg {
	cfgPath := os.Getenv("billcalcCfgPath")
	unmarshaledCfg, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		panic(`cfg is not found.
		` + err.Error())
	}
	var cfg = &Cfg{}
	if err := json.Unmarshal(unmarshaledCfg, cfg); err != nil {
		panic(err.Error())
	}
	return cfg
}
