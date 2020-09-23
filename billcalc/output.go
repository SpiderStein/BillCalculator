package billcalc

// Help is the output to be shown when billcalc executable gets no/invalid arguments
const Help = `
Bill Calculator

*	Calculations include VAT if it's required.

Usage: billcalc <bill-type> [arguments]

The bill types and their arguments are the following:
*	water -c <m3>
*	electricity -c <1000imp/kWh => what's written in the meter>`

// WaterCFlagHelp is the help output of flag "c" when flag "bill-type" is assigned to the value "water"
const WaterCFlagHelp = "Water consumption in m3 unit"

// ElectricityCFlagHelp is the help output of flag "c" when flag "bill-type" is assigned to the value "electricity"
const ElectricityCFlagHelp = "Electricity consumption in 1000imp/kWh => what's written in the meter"
