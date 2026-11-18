# Parts list, Talos driver module rev 0.1

Everything for **one module**. The assembly service fits every part except the pin strips; `fab/jlcpcb_bom.csv` is the same list in
the form JLCPCB reads, with the part number it matches on. Prices move, so none are given per line: ten assembled modules were quoted at
USD 49.22 on 13 November 2026 (lead-free HASL, before shipping), about USD 5 each.

## Fitted by the assembly service

| Ref | Qty | Value | Part | Package | JLCPCB part |
| --- | --- | --- | --- | --- | --- |
| U1 | 1 | DRV8701E | Texas Instruments DRV8701ERGER, H-bridge gate driver | VQFN-24, 4 × 4 mm | C90964 |
| Q1, Q2 | 2 | AO4882 | Alpha & Omega AO4882, dual N-channel MOSFET, 40 V | SOIC-8 | C477943 |
| R1 | 1 | 20 mΩ 1 % 2 W | current-sense shunt, metal element | 2512 | C105360 |
| R2 | 1 | 1.5 kΩ 1 % | | 0603 | C22843 |
| R3A | 1 | 100 kΩ 1 % | | 0603 | C25803 |
| R4 | 1 | 12 kΩ 1 % | | 0603 | C22790 |
| R5, R7 | 2 | 10 kΩ 1 % | | 0603 | C25804 |
| R6 | 1 | 33 kΩ 1 % | | 0603 | C4216 |
| R8 | 1 | 22 Ω 1 % | | 0603 | C23345 |
| C1, C2 | 2 | 10 µF 50 V X5R | | 1206 | C13585 |
| C3, C4, C9, C10, C11 | 5 | 100 nF 50 V X7R | | 0603 | C14663 |
| C5, C6, C7 | 3 | 1 µF 50 V X5R | | 0603 | C15849 |
| C8, C12 | 2 | 1 nF 50 V X7R | | 0603 | C1588 |

R3B is an empty position (section 5 of the guide). TP1 to TP6 are bare test pads.

## Fitted by hand

| Ref | Qty | Part |
| --- | --- | --- |
| J1 | 2 | 1 × 8 male pin strip, 2.54 mm pitch, straight (for example XFCN PZ254V-11-08P, LCSC C492407; any standard 0.64 mm square-pin strip cut to eight) |

## What the parts do

| Part | Job |
| --- | --- |
| U1 | drives the four MOSFET gates, limits the current, reports faults, makes its own 3.3 V |
| Q1, Q2 | one half-bridge each: Q1 switches OUTA, Q2 switches OUTB |
| R1 | the motor current returns through it; its voltage is the current-sense signal |
| R2, C8 | pass the shunt voltage to the CS pin and filter it |
| R3A (or R3B) | adds the 50 mV offset to CS |
| R4, R5, C9 | set the current limit |
| R6 | sets how fast the MOSFETs switch |
| R7 | pull-up for the current-limit indicator on test pad TP3 |
| R8, C12 | a small filter at the driver's current-sense input; R8 becomes 100 Ω only if the short-circuit test asks for it |
| C1, C2, C10, C11 | supply capacitors at the MOSFETs |
| C3 to C7 | the driver's own supply, charge pump and regulator capacitors |

## Substitution rules

- **U1** has no substitute. It must be the DRV8701**E** (direction and PWM inputs), not the DRV8701P.
- **Q1, Q2**: keep the AO4882 if you can. A substitute must be a dual N-channel MOSFET in SOIC-8 with the same pin order
  (S2, G2, S1, G1 on pins 1 to 4, D1 on 5 and 6, D2 on 7 and 8), 40 V or more, ±20 V gate rating, and an on-resistance of 20 mΩ or
  less at 10 V gate drive. Copies sold under the same name differ in on-resistance: read the sheet.
- **R1** must be a 2 W (or more) metal-element shunt in 2512, 1 %. Its value sets both the current-sense scale and the current limit.
- **R4, R5** set the current limit together; change neither without recalculating it.
- **C1, C2** must be rated 50 V; a lower-voltage part loses much of its capacitance at 12 V.
- The other resistors and capacitors are ordinary parts: same value, same size, 1 % for the resistors, 50 V for the capacitors.
