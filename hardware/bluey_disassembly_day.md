J Russ
2/27/2026

# Bluey Disassembly

## Why
During development testing of the ESP32 pulse counter, encoders on axis 4 and 5 were overvolted and burnt out. I (J) decided that it would be best to trek it accros campus to the shed to both:
1. replace the burnt out phototransistor assmeplies
1. replace the shittium encoder wheels with 3D printed+painted wheels on axis 1 2 and 3 which cannot be accessed without substantial disassembly effort 

I know the encoders were burnt up becasue I was measuring voltage in both directions across the diodes. Earlier in the semester I did an experiment with one of the newer OPB628 phototransistors using the schematic from the robot (PC510) and found that if I drove the Vcc with ~4 volts it would keep the diodes at acceptable current levels while not undervolting the output transistor. The sponsor had mentioned that I cannot shcange the 51ohm resistor on the PC510, becasue then it would not be backwards compatible with other scorbots that end users may use.

## Experiment 1

To figure out how Eshed drove the encoders, I'm going to measure the voltage out of the ER4U controller box.

Somehow the ER4U is driving the PC510 at a steady 5 volts???? I measuerd between pin33 (axis1 GND) and pin 17 (axis1 VLED)

the H22LOI phototransistor that comes stock on the ER4pc -- that also might be a source of error, not the ER4u-- is rated for 50mA MAX, with a forward voltage of 1.7 volts.

THERES AN EXTRA 47ohm RESISTOR IN THE BOX

## Experiment 2

Now that I know theres actually 100 ohms of resistance, I made a spice model to simulate the current draw across the diodes. 35.7 mA is what the model says.

## Experiment 3 

To make absolutley sure, I made a breadboarded version of my spice model. I use a resistor ladder, into the diodes
`5V -> 47ohm -> 47ohm -> 4.7ohm -> diode1 -> diode2 -> GND`
The current drawn on the power supply shows that the current as 22mA, and if I take out one of the 47ohms it goes up to 33mA.

This was different from what the simulation showed, so I made sure my parameters were right in spice. I adjusted the Emission Coefficient and Saturation Current to be in line with an infrared diode, (from  Gemini advice). The currents were much more similar at 24mA, and 46mA(?) without a 47ohm. 

## Results

Im going to add 47 ohms in from of all of the VLED supply pins, and use 5 volts. 


## Hooking up to "dead" encoders
THEYRE ALIVE!!!!!! with the extra 47 ohm resistor, a ran 5 volts through the VLED and held my phone camera up the the diodes. I saw a purple dot!!!


## Replacing encoder wheels
Axis 1 was easy. I pulled off the black back cover and had access to the encoder. I noticed that the shaft through the encoder wheel was spinning when I rotated the bot, but the wheel itself did not spin. I replaced it without closing the top

Axis 2 and 3 I had to remove from the wall carefully so that I could access the encoder compartment, and I observed the same non-rotation of the encoder wheel. It seems the old encoder wheels have cracks where the grub screw affixes to the shaft, so theres no pressure keeping it on the shaft.