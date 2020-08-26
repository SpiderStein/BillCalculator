package billcalc

// CalcWaterBill calculates and returns the water bill.
// The consumption is in m3 unit.
func CalcWaterBill(consumption, lowCommissionConsumption, lowCommission, highCommission float64) float64 {
	if consumption > lowCommissionConsumption {
		highCommissionConsumption := consumption - lowCommissionConsumption
		return lowCommission*lowCommissionConsumption + highCommission*highCommissionConsumption
	}
	return consumption * lowCommission
}

// CalcElectricityBill calculates and returns the water bill.
// The consumption is in 1000imp/kWh.
func CalcElectricityBill(consumption, commission float64) float64 {
	return commission * consumption
}
