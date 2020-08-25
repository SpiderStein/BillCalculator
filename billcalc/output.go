package billcalc

// Help is the output to be shown when billcalc executable gets no/invalid arguments
const Help = `
Bill Calculator

*	Calculations include VAT if it's needed.

Usage: billcalc <bill-type> [arguments-if-needed]

The bill types and their arguments are the following:
*	water -c <m3>
*	electricity -c <1000imp/kWh => what's written in the meter>`

// WaterCmdArgHelp is the help output of water cmd 'c' argument
const WaterCmdArgHelp = "Water consumption in m3 unit"

// ElectricityCmdArgHelp is the help output of electricity cmd 'c' argument
const ElectricityCmdArgHelp = "Electricity consumption in 1000imp/kWh => what's written in the meter"
