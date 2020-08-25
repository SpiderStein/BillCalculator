package billcalc

// CalcWaterBill calculates and returns the water bill.
// The consumption is in m3 unit.
func CalcWaterBill(consumption float64) float64 {
	const (
		lowCommissionConsumption float64 = 3.5
		lowCommission            float64 = 7.385
		highCommission           float64 = 13.461
	)

	if consumption > lowCommissionConsumption {
		highCommissionConsumption := consumption - lowCommissionConsumption
		return lowCommission*lowCommissionConsumption + highCommission*highCommissionConsumption
	}

	return consumption * lowCommission
}

// CalcElectricityBill calculates and returns the water bill.
// The consumption is in 1000imp/kWh.
func CalcElectricityBill(consumption float64) float64 {
	const commission float64 = 0.4484
	return commission * consumption
}
