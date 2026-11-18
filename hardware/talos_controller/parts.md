# Parts list, Talos controller rev 0.5

Everything needed for **one board**, with spares of the cheap parts. Prices are not given: they move, and the vendor pages are the
source. Part numbers were checked against the makers' sheets during the design review; stock was not checked live, so expect to
substitute the odd line (the rules for that are at the end).

The board is built in two steps: JLCPCB solders most of the through-hole parts from its own catalogue (`fab/jlcpcb_bom.csv`,
`fab/jlcpcb_cpl.csv`), and you fit the rest (`fab/hand_assembled.csv`). The lists below follow that split: **always** is what you buy
whatever JLCPCB does; **fallback** is a part JLCPCB normally fits, to buy only for a line it could not place (it tells you at the
BOM step, and the hand list in this folder is where such a line ends up); **optional** needs no explanation. If you have the whole
board assembled by hand instead, buy the fallback lines too.

## Digi-Key

Upload `fab/digikey_upload.csv` with "Upload a list" (columns: quantity, part number, reference), then delete the fallback lines you
do not need. The Adafruit and SparkFun parts are on the same list; they can also be bought from the makers directly.

### Always

| Qty | Part number | What it is |
| --- | --- | --- |
| 1 | `DD50S565MTLF` | J1 D-sub DD-50 socket, right angle, solder tails, metal brackets (the robot cable); long lead time, about 15 weeks |
| 2 | `09670019985` | J1 4-40 hex screwlock kit, one per side |
| 2 | `ESP32-S3-DEVKITC-1-N16R8` | U1 ESP32-S3-DevKitC-1 (1 + spare). Any variant that is not a "...V" part fits: N16R8 (the one in production), N8R8, their 1U versions, and the old N8 and N8R2 (both end-of-life). **Stock was thin in November 2026 (Mouser had the N16R8 on backorder, Digi-Key had no N8R8): check several distributors and buy it first** |
| 8 | `LS7366R` | U13-U19 quadrature counters, DIP-14 (7 + spare); pushed into the sockets on the board |
| 2 | `1083` | U11, U12 Adafruit ADS1015 breakout, STEMMA QT version (plug-in) |
| 2 | `SN74HCT125N` | U22 pixel level shifter, DIP-14 (1 + spare); pushed into its socket. The `SN74AHCT125N` the board was drawn with fits as well; it was out of stock in November 2026 |
| 2 | `TLV431BILP` | U21 1.24 V shunt reference, TO-92: a fallback and spare, since JLCPCB fits U21 from its own stock. TI part only (pins 1 K, 2 A, 3 REF). The `TLV431ACLP` the board was drawn with came only in reels at Digi-Key in November 2026; the `TLV431BILP` is the same reference in a tighter grade, and any TI TLV431 in the LP (TO-92) package does |
| 20 | `BAT85S-TAP` | D1-D14, D23, D24 Schottky clamps, DO-35 (16 + spare); soldered by hand |
| 3 | `12986` | LED1-LED10 SparkFun 5 mm addressable pixels, pack of 5 (10 + 5 spare); soldered by hand |
| 4 | `PPTC221LFBN-RC` | JH1, JH2 1x22 female headers, the DevKit socket (2 + 2 spare); soldered by hand |
| 5 | `MFR-25FBF52-56K` | R61 56 k, upright (1 + spare); soldered by hand |
| 5 | `MFR-25FBF52-3K` | R64 3.0 k, upright (1 + spare); soldered by hand |
| 5 | `MFR-25FBF52-1K3` | R65 1.3 k, upright (1 + spare); soldered by hand |
| 1 | `0039300240` | J2 ATX 24-pin right-angle header, Molex Mini-Fit Jr 5569-24A2; soldered by hand |
| 1 | `0039300040` | J3 ATX12V 2x2 right-angle header, Molex 5569-04A2; soldered by hand |
| 1 | `XA1E-BV3SG01R` | SW1 IDEC 16 mm emergency stop, 1 NC, solder terminals; mounted through the board |
| 1 | `HAAV-27` | SW1 emergency-stop nameplate (confirm the code on the product page) |
| 0 | `XA9Z-LNW` | SW1 locking ring: Digi-Key has no such listing (November 2026), so it is off the upload list. Check on arrival that the switch came with its ring |
| 1 | `B32-1310` | SW2 key top, 12 x 12 mm, black |
| 3 | `028707.5PXCN` | F1 ATO fuse 7.5 A (1 + 2 spare) |
| 3 | `0287001.PXCN` | F2 ATO fuse 1 A (1 + 2 spare); never a 7.5 A blade in F2 |
| 2 | `SPC02SYAN` | JP1 jumper shunt (1 + spare) |
| 39 | `MFR-25FBF52-10K` | 10 k: the board's 25 if JLCPCB did not fit them, plus one per Pololu G2 driver if those are used (VREF current limit, up to 7; the Talos driver module needs none), plus spare |
| 2 | `1N5819G` | D21 Schottky (1, if JLCPCB did not fit it) and the alternative for the D22 link if the pixel rail reads above 5.25 V |

### Fallback: only for a line JLCPCB did not place

| Qty | Part number | What it is |
| --- | --- | --- |
| 9 | `ED14DT` | DIP-14 sockets for U13-U19 and U22 (8 + spare), only if JLCPCB did not fit its own |
| 2 | `MCP23017-E/SP` | U9, U10 I/O expanders, SPDIP-28, only if JLCPCB did not solder them |
| 2 | `LM393P` | U20 dual comparator, DIP-8 (1 + spare), only if JLCPCB did not solder it; one maker only (TI LM393P, ST LM393N or onsemi LM393NG) |
| 3 | `G2RL-1A-E DC12` | K1, K2 relays, 16 A, 12 V coil, only if JLCPCB did not fit them (the third is a spare) |
| 1 | `1715721` | J4 screw terminal 2-pos 5.08 mm (external e-stop), only if JLCPCB did not fit one |
| 1 | `1715789` | J8 screw terminal 8-pos 5.08 mm (the rail cable), only if JLCPCB did not fit one |
| 2 | `B3F-4055` | SW2 MOTORS ON tactile switch 12 x 12 mm (1 + spare), only if JLCPCB did not fit it |
| 2 | `178.6165.0001` | F1, F2 ATO fuse holders, only if JLCPCB did not fit them |
| 2 | `RUEF300` | F3 PTC 3 A hold (1 + spare), only if JLCPCB did not fit the Bourns MF-R300 |
| 2 | `RXEF050` | F4 PTC 0.5 A hold (1 + spare), only if JLCPCB did not fit it |
| 4 | `EEU-FR1E102` | C42-C45 1000 uF 25 V, 10 mm can, 5 mm pitch, only if JLCPCB did not fit them |
| 1 | `EEU-FR1E101` | C31 100 uF 25 V, 6.3 mm can, 2.5 mm pitch, only if JLCPCB did not fit one |
| 1 | `1.5KE16A` | D19 TVS 16 V, only if JLCPCB did not fit it |
| 2 | `1N4007G` | D20, D25 rectifier diodes (silicon, never a Schottky), only if JLCPCB did not fit them |
| 4 | `WP710A10GD` | D16, D17 3 mm green lamps (2 + 2 spare), only if JLCPCB did not fit its own |
| 2 | `WP710A10YD` | D18 3 mm yellow lamp (1 + spare), only if JLCPCB did not fit its own |
| 1 | `PRPC040SAAN-RC` | 1x40 breakaway pin header for J7 and JP1, only if JLCPCB did not fit them |
| 16 | `PPTC081LFBN-RC` | JH3-JH16 1x8 female headers, the driver sockets (14 + 2 spare), only if JLCPCB did not fit them |
| 6 | `PPTC061LFBN-RC` | JH17-JH20 1x6 female headers, the ADC sockets (4 + 2 spare), only if JLCPCB did not fit them |
| 17 | `FG28C0G1H102JNT06` | 1 nF 50 V C0G, 5 mm pitch (14 + spare), only if JLCPCB did not fit its own |
| 32 | `FG28X7R1H104KNT06` | 100 nF 50 V X7R, 5 mm pitch (28 + spare), only if JLCPCB did not fit its own |
| 27 | `MFR-25FBF52-1K` | 1 k (23 + spare), only if JLCPCB did not fit them; R55-R60 and R69 are fitted upright |
| 10 | `MFR-25FBF52-47R` | 47 ohm encoder LED series resistors (7 + spare), only if JLCPCB did not fit them |
| 12 | `MFR-25FBF52-2K2` | 2.2 k (6 + spare), only if JLCPCB did not fit them; R66 and R67 are fitted upright |
| 5 | `MFR-25FBF52-39K` | R46 39 k (1 + spare), only if JLCPCB did not fit it |
| 5 | `MFR-25FBF52-30K` | R48 30 k (1 + spare), only if JLCPCB did not fit it |
| 5 | `MFR-25FBF52-330R` | R53, R70 330 ohm (2 + spare), only if JLCPCB did not fit them |
| 5 | `MFR-25FBF52-100K` | R62 100 k, upright (1 + spare), only if JLCPCB did not fit it |
| 5 | `MFR-25FBF52-22K` | R63 22 k, upright (1 + spare), only if JLCPCB did not fit it |
| 5 | `MFR-25FBF52-4K7` | R68 4.7 k, upright (1 + spare), only if JLCPCB did not fit it |

### Optional

| Qty | Part number | What it is |
| --- | --- | --- |
| 10 | `MFR-25FBF52-39R` | 39 ohm alternative for the seven encoder LED resistors, if the encoder LED current is low on your robot |
| 10 | `MFR-25FBF52-33R` | 33 ohm alternative for the seven encoder LED resistors |
| 1 | `FX0380` | inline 5 x 20 mm fuse holder for an external e-stop cable (Bulgin) |
| 2 | `0217.500MXP` | 0.5 A fast 5 x 20 mm fuse for that holder (1 + spare) |

## Driver modules

| Qty | Item | What it is |
| --- | --- | --- |
| 5 or 7 | Talos driver module, [`../talos_driver_module/`](../talos_driver_module/README.md) | one per channel: five for the arm (U2 to U6); the gripper (U7) and the rail (U8) need two more. Ordered from JLCPCB, assembled, with the files in that folder; two 1x8 pin strips (on the Digi-Key list) are soldered to each and plug into the sockets on the board. A first prototype: designed and reviewed, not yet built or tested |
| instead | Pololu 2991, G2 High-Power Motor Driver 18v17 | the commercial driver the sockets were drawn for. It fits the same sockets and ships with its pin strips: the alternative if the Talos modules misbehave on the bench |

The Talos module has its current limit on board (3.6 A typical). A Pololu G2 has none below its 40 A default, so every fitted G2
gets a 10 kΩ resistor (on the Digi-Key list) between its VREF pad and the GND pad beside it: about 3.7 A. Never less than 10 kΩ,
and never 15 kΩ on the rail's module.

## Adafruit and SparkFun

On the Digi-Key list as `1083` (ADS1015 breakout, STEMMA QT version, two) and `12986` (5 mm addressable pixels, three packs of five).
Pick the Adafruit `1083`, not another maker's ADS1015 board: the pin rows must match the two 1x6 sockets. Check a pixel pack on
arrival before soldering any: the leads along the row must read DIN, VDD, GND, DOUT (GND and DOUT are the long leads); the
footprint's square pad is DIN.

## Elsewhere (lab stock, Amazon, any supplier)

| Qty | Item | Notes |
| --- | --- | --- |
| 1 | ATX PC power supply | any ATX supply with a 24-pin main lead and a 4-pin "CPU 12 V" lead (or the latched half of a 4+4). 12 V enters the board only through the 4-pin lead; a supply of 300 W or more has all the headroom this board can use |
| 2 | micro-USB **data** cables, 0.3 to 1 m | one from the Pi to the DevKit's "UART" socket (the link), one for the "USB" socket (console and logs). Charge-only cables do not work |
| 1 kit | M3 hardware | seven M3 pan-head screws with DIN 125 washers for the mounting holes H1 to H4, H9, H10 and H11 (a nylon washer, or none, at H10 and H11 where copper is close), seven standoffs of 15 mm or taller (the e-stop body reaches 12.6 mm below the board), two M3 screws with nuts for the D-sub bracket holes SH1 and SH2, a nylon screw or nylon washer at SH1 (a motor track runs close to it) |
| 1 | 80 mm 12 V fan with a 3-pin plug | optional, for the header J7 |
| 1 | rail pigtail | a female D9 to bare wires, about 0.5 m, to bring the slide rail's cable to the screw terminal J8 (pin table in the README) |
| 1 | 30 mm of solid 20 AWG insulated wire | the link in the external e-stop terminal J4 when no external button is fitted |
| 1 | external emergency-stop button, optional | a proper e-stop with an NC contact pair, 2 m of 18 AWG cable; fused as the README describes |
| 1 | enclosure | the board is designed to be its own front panel for the e-stop; the team prints its own enclosure. Keep it plastic or insulate a metal one from the board's underside |

## Rules when a part is out of stock

- **The DevKit:** any `ESP32-S3-DevKitC-1` that is not a part ending in V: N16R8, N8R8, the 1U versions (external antenna socket),
  or the old N8 and N8R2. A part ending in V (N16R8V, N32R16V) runs two of the pins the board uses at 1.8 V; it does not work here.
- **D20 and D25:** any maker's 1N4007 (silicon), never a Schottky; the relay hold time depends on it.
- **F2:** a 1 A blade fuse, never the 7.5 A of F1; F2 protects the e-stop chain and the external cable.
- **U21:** a TI TLV431 in the LP (TO-92) package, `TLV431BILP` or `TLV431ACLP`. Another maker's TLV431 only if its datasheet puts the anode on the middle lead of the TO-92.
- **U22:** `SN74HCT125N` or `SN74AHCT125N`, DIP-14. Both read the 3.3 V data line as a valid high from a 5 V supply; a plain 74HC125 is not guaranteed to.
- **U20:** one of TI LM393P, ST LM393N or onsemi LM393NG, no unbranded clone: the supply-loss detection depends on its documented behaviour.
- **C42 to C45:** 1000 µF 25 V with a 10 mm can and 5 mm pitch (12.5 mm cans do not fit). **C31:** 100 µF 25 V, 6.3 mm can, 2.5 mm pitch; any maker.
- **The clamps D1 to D14, D23, D24:** BAT85. The BAT42 that JLCPCB stocks is accepted for D15 only.
- **Resistors:** 1/4 W metal film, 1 %, 0207 body. The sense resistors R61, R64, R65, R66, R67 must stay 1 %.
- **Relays:** Omron G2RL-1A-E DC12; the inrush-rated G5RL-1A-E-HR DC12 is the drop-in alternative (same footprint and coil).
- **The 1x8 and 1x6 female headers:** any 2.54 mm strip 8.5 mm high with 3 A contacts; the 1x22 for the DevKit must be a single strip.
