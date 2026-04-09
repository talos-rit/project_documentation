# For the next team

Hi! My name is J! I was in charge of the hardware development during the 2025-2026 team.

## Where we started

The initial project state of the hardware was a mess. Brooke from the 24-25 team was using a raspberry pi with a 4 ch motor driver hat for controlling the motors, which worked for their prototype, but for the long term we had issues with that motor driver.

ESP32 is really good at encoders with its integrated 240MHz pulse counter, and integrated library. However, it physically does not have enough pins to do all of the jobs that need to be done. We thought that maybe we use 2 different microcontrollers that can handle different tasks 


# Where we leave it
We unfortunatley added to the kill count metric quite a bit this year. Brooke killed a motor driver and an encoder, (which have since been fixed). We killed a motor driver, and then another motor driver, an esp32, and 3 ADS1015 modules.

### The E-stop Button
Originally, the 24-25 team used a direct 12 volt power source. Since we knew we needed logic voltage ofr running the microcontroller, it didn not make sense to get another whole power supply for that 3.3 or 5 volts, or use a step down converter. We found an old 'gaming' comupter in the GCCIS e-waste, which wouldn't boot we think due to graphics card shinanegans. We took the ATX poewr supply out of it which if you didn't know, has almost all of the common voltages an engineer would ever need, and some at significant ampacities. On the one we got, there are 5 seperate 12 volt lines that could each do like 10 ampsish. ATX is epic.

We decided that to actually do an E-stop, we just use it as the 'reverse on' button. When the estop is NOT triggered, the power supply is on. When the estop is triggered, the power supply turns off. simple. THERE IS A CATCH, when the estop is hit the voltage is drained from the lines not instantly, and has a logrithmic effect as the internal capacitors discharge. This doesn't have a noticable effect on the estop functionality, the robot stops moving effectivly instantly however the 12 volt rail after about 2 seconds still has like 5 volts.

### Makerfaire Incident
While we were testing the motor driver firmware during makerfaire, we didn't know that the slide base draws to much current than the motor driver can handle. To be technically correct, it barely can handle driving the slide base but when the slide base was stopped abruptly, there was too much bEMF and the diodes inside the TB6612FNG went smokey. This is why we decided to add an inductive kickback protection circuit, which is talked about in a lower section. Kill count: 2->3 :(

### Metal Enclosure Incident
For most of the project, we decided to disembowel the broken driver box made by eshed robotic (with SCORPOWER in the bottom left) to use to support the mecahnical connection to the 50pin connector going to the robot. We also wanted to mount the Estop button to it, but never got around to that. Right before spring break (cuz ofc)I had finally soldered together the current version of the hardware using perfboard/protoboard. We did this so that there would be no more accidental shorting incidents due to breadboards. We were testing all of the fucnctionality, so we didnt want to actually move the motors we just wanted to communicate over I2C so we disconnected the 12 volt power lines to the motor driver hat. To make sure these didnt short anywhere we just bent them back, but as the wire slowly straigtened out over the course of like 3 minutes, (while J was going for a drink for the vending machine) we are assuming that the 12 volt line shorted to the metal enclosure, which was shorted to one of the I2C lines(or just directly shorted to one of those lines). We came to this conclusion based of the things that started smoking as soon as J walked back into the room. (I think the robot hates me) 3 ADS1015 modules, the PCA9685 on the motorhat, and the esp32 (not visible damage, under the RF shield) all went smokey and were toast. ESP32 wouldn't flash, 2 of the ADS modules had visible thermal damage, and the motorhat and other ADS module didn't respond to I2Cdetect. We learned then, to tie down cables if theyre unused. Also, we took the ciurcuts out of the metal enclosure.
    

## The things we did to bluey
During our time fixing the problems that were left with the blue ER-4pc robot, we decided to replace the sinterd clay style encoder wheels with 3D printed wheels of the same resolution becasue the sintered clay wheels where stripped out and would not turn with the rotor. Unfortunately, the wavelength of light that the phototransistors use to detect if the encoder has turned is in the UV spectrum, which the resin 3D prints we made do not block UV light well becasue thats how they're cured when made. We had to paint the black with acrylic paint to work properly. If these are ever broken in the future, the file is [online](https://www.thingiverse.com/thing:3388595#google_vignette) called "Scorbot ER-4u encoder (Parametric design)"

### Circuit Protections
put a fuse somewhere! even if the inductive kickback is typically fine, there should be a secondary circuit protection system protecting the circuits from freak events. They might cost monney to replace, but theyre cheaper and easier than replacing modules, resoldering PCBs, or buying a new microcontroller.


### Inductive Kickback Protection
The inductive kickback protection designed is not correct. Originally, there are shottky or zener diodes meant to dissipate the excess voltage generatoed by the DC motors when the motor stop abruptly (back-electromotive force; bEMF), as we found that going from full throttle (using the TB6612) generated more than 15 volts in the opposite direction of motion. The SB550 shottky diodes are rated at 50 volts rather than the needed 13.5-15 volts. A new inductive kickback protection circuit should be designed. Realistcly, the slide base is the only motor that needs these more beefy protections, but they should be implemented on all motor regardless. The size and circuit may change depending on the motor driver you choose to switch to.

CURRENTLY, we are not using the 'brake' configuration for each of the H-bridges inside the motor controller to limit the instantaneous inductive kickback as we've had issues worrying about it. Ideally the motors would brake themselves to maintain position when un-powered, this would have to be changed inside of the esp-driver repository, under the 'motorhat' component. 










## Logic Controller

### Microcontroler vs Raspberry Pi?
The previous team started the 'Ichor' project, named after an old greek autonoton, in line with the Talos name. This project was meant to controller the motors, read the encoders, and read current from the motors. The 24-25 team didnt get very far with this project though, the most achieved was back and fourth motion of the motors by manually connecting wires to the specific pins of the desired motor on the big 50 pin connector. Raspberry Pi worked for this, but we (25-26 team) found that the raspberry pi, though convenient for debugging through SSH, could not handle the intense frequency of pulses that are generated by the quadrature encoders. Through some simulations, we found that running at a decent speed each motors encoder generates ~3,000 interuppts per second. We decided to move to the ESP32 so that we had a super low level control. ESP32 runs freeRTOS by default! It also has a 4-8 pulse counter (PCNT) peripheral, which runs completely seperate from both CPU cores and can sense pulses up to 40MHz. The ESP32-S3 only has 4 PCNT units, but the OG ESP has 8.

### Help Im Out of Pins!
While I was designing the schematic for hooking up all the GPIO pins of the ESP32, I realized that there are not enough GPIOs to hook up all 6+1 encoders, 5 limit switches, and I2c for motors and current sensors. There are quite a few reserved pins on the ESP32, there is a set of SPI pins for flash-CPU memory control which CANNOT be touched, strapping pins for boot mode, and even though the UART pins can technically be used for GPIO, if they are they cant be used for USB debugging at the same time. How do we get around this? 2 esp32's can be hooked together for extra processing power and pins, for the encoder pulse counters and i2c busses! To do this we NEED to have a rock solid RTOS serial implementation though, as there could be some very nasty crashes if the closed loop motor control gets off syncronization.

## Things you should probably(definetly) implement first

### Motor Velocity Collection

We hadn't quite got to this during our time, but we were about to start implementing it. Just use an esp32 realtime timer and sample how many encoder pulses change per time increment. To note, the smaller the time increment the higher resolution of speeds, but the more processing power. Realisticly it might be prudent to use a significant amount of processing resources on this, because the smoother the motion of the arm, the smoother the camera motion. It matter even more when the camera is zoomed in on a distant subject.

### Closed Loop Motor Control
To accuratly position the DC motors, you'll need to implement a version of 'closed loop motor control.' This involves using the position and velocity (gotten from the encoders) to accuratley position the motor. You'll also need this for impact detection, as the current spikes so much during instantaneous start up, triggering the impact detector everytime. 

### Motion Profiler

For almost all times you want the robot to move, you dont want to immediatly hit the gas and jerk the arm. You want to slowly ramp up speed, hold a speed, then ramp down until you hit the target position. You'll need to implement closed loop motor control before this, and it probably should be part of the motor control component as its technically higher level than actually controlling the robot. This should be set based on the distance from the target position, but in our case the target position can be constantly changing so maybe look at a different algorithm for that. PID)position, integral, di controllers are a comparable algorithm in the motor driving industry.

### Impact/Overcurrent Detection
To actually detect an impact, you'll need to cross reference the speed/poisiton of the motor with the current that the motor is drawing. There are cases, like on startup, where the current spikes but isnt acutally an impact. Also cases where the current of the motor changes with position due to gravity, specificlly on bluey's 2nd axis(the shoulder). These should be accounted for in this component. Overcurrent detection is fundamentally different though, even though it shares the responsbility of immediatly stopping all motors if triggered. Overcurrent should fundamentally protect the circuitry from excess current and overheating, rather than keeping the robot human safe. Maybe implement it as a sepearte component as well? 

# Motor Driver
The previous team used a 4ch TB6612FNG motor driver hat controlled by a PCA9685 I2C to PWM driver. It worked for the prototype, and it worked for most of the motors on bluey, but we had issues with the slide base due to the moving mass of the robot. It drew to much current from the driver, burning it out.



## PCB

## New Driver Considerations
If we make a custom PCB, we can use surface mount components and still be 'hobbyist grade' as hobbyists can order the PCB and get the surface mount components pre-soldered on by the manufactor for a small extra cost. (JLCpcb, PCBway, there is also a pick and place shop in slaughter!) 

NEED to overspec! needs to feed AT LEAST 3 AMPS as the max peak of the slide base is 2.2sh amps. 
NEED to make sure inductive kickback is accounted for! 
SHOULD have inherent current sensing! this would be much simpler than doing all this shunt resistor stuff, and the shunt resistor wouldn't be high side or low side, it'd be before/inside the H bridge itself.

Potentially look at the STSPIN9P2 motor driver by ST, it seems to have most of the needed resources. Don't take our word for it though, do your own research before suggesting for purchase!


## Helpful Hints
1. TIE DOWN ALL UNUSED CABLES
1. Label everything, even if you think it would be obvious. 
    1. Just because you think it would, doesn't mean anyone else does too. Increase the bus factor!