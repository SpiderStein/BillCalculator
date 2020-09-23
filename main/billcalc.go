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
		waterFlagSet := flag.NewFlagSet("water", flag.ContinueOnError)
		cFlag := waterFlagSet.Float64("c", 0, billcalc.WaterCFlagHelp)
		if err := waterFlagSet.Parse(os.Args[2:4]); *cFlag < 0 || err != nil {
			break
		}
		bill := billcalc.CalcWaterBill(*cFlag, cfg.Water.LowCommissionConsumption, cfg.Water.LowCommission, cfg.Water.HighCommission)
		fmt.Printf("The water bill is the following: %v\n", bill)
		return
	case "electricity":
		electricityFlagSet := flag.NewFlagSet("electricity", flag.ContinueOnError)
		cFlag := electricityFlagSet.Float64("c", 0, billcalc.ElectricityCFlagHelp)
		if err := electricityFlagSet.Parse(os.Args[2:4]); *cFlag < 0 || err != nil {
			break
		}
		bill := billcalc.CalcElectricityBill(*cFlag, cfg.Electricity.Commission)
		fmt.Printf("The electricity bill is the following: %v\n", bill)
		return
	}

	fmt.Println(billcalc.Help)
}
