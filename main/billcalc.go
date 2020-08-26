package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/steidan/BillCalculator/billcalc"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(billcalc.Help)
		}
	}()

	cfg := billcalc.GetCfg()
	selectedBillType := os.Args[1]
	switch selectedBillType {
	case "water":
		waterCmd := flag.NewFlagSet("water", flag.ContinueOnError)
		consumption := waterCmd.Float64("c", 0, billcalc.WaterCmdArgHelp)
		if err := waterCmd.Parse(os.Args[2:4]); *consumption < 0 || err != nil {
			break
		}
		bill := billcalc.CalcWaterBill(*consumption, cfg.Water.LowCommissionConsumption, cfg.Water.LowCommission, cfg.Water.HighCommission)
		fmt.Printf("The water bill is the following: %v\n", bill)
		return
	case "electricity":
		electricityCmd := flag.NewFlagSet("electricity", flag.ContinueOnError)
		consumption := electricityCmd.Float64("c", 0, billcalc.ElectricityCmdArgHelp)
		if err := electricityCmd.Parse(os.Args[2:4]); *consumption < 0 || err != nil {
			break
		}
		bill := billcalc.CalcElectricityBill(*consumption, cfg.Electricity.Commission)
		fmt.Printf("The electricity bill is the following: %v\n", bill)
		return
	}

	fmt.Println(billcalc.Help)
}
