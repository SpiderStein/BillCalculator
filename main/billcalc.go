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
		consumption := waterCmd.Float64("c", -1, billcalc.WaterCmdArgHelp)
		if flag.Parse(); *consumption < 0 {
			break
		}
		bill := billcalc.CalcWaterBill(*consumption)
		fmt.Printf("The water bill is the following: %v", bill)
		return
	case "electricity":
		electricityCmd := flag.NewFlagSet("electricity", flag.ContinueOnError)
		consumption := electricityCmd.Float64("c", -1, billcalc.ElectricityCmdArgHelp)
		if flag.Parse(); *consumption < 0 {
			break
		}
		bill := billcalc.CalcElectricityBill(*consumption)
		fmt.Printf("The electricity bill is the following: %v", bill)
		return
	}

	fmt.Println(billcalc.Help)
}
