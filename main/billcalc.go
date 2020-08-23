package main

import (
	"flag"
	"fmt"

	"github.com/steidan/BillCalculator/billcalc"
)

func main() {
	selectedBillType := flag.Arg(0)
	switch selectedBillType {
	case "water":
		waterCmd := flag.NewFlagSet("water", flag.ContinueOnError)
		waterCmd.Float64("c", 7.385, "Water consumption in m3 unit")
	case "electricity":
		electricityCmd := flag.NewFlagSet("electricity", flag.ContinueOnError)
		electricityCmd.Float64("c", 7.385, "Electricity consumption in 1000imp/kWh => what's written in the meter")
	default:
		fmt.Println(billcalc.Help)
		return
	}
	// cfg := billcalc.GetCfg()
}
