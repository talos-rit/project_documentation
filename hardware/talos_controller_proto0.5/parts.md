# Parts list, Talos controller prototype 0.5

Everything for **two boards**, the quantity the carts on [`purchase_options.md`](purchase_options.md) were made for.
Prices are on that page only, as the carts of 23 November 2026 showed them. Part numbers were checked against the
makers' data sheets where this list says so; stock moves, so check it on the day.

The board is built in two steps: JLCPCB places every surface-mount part from its own catalogue (`fab/jlcpcb_bom.csv`,
`fab/jlcpcb_cpl.csv`), and you fit the through-hole parts (`fab/hand_assembled.csv`) from the Digi-Key list
(`fab/digikey_cart_two_boards.csv`).

## Placed by JLCPCB

29 lines, 241 placements a board, Economic assembly, top side. JLCPCB matches each line by its own part number (the
"JLCPCB Part #" column); the last column of `fab/jlcpcb_bom.csv` names the maker's part. Eight lines are Extended
parts, marked *E*.

| Designators | Part | JLCPCB | Notes |
| --- | --- | --- | --- |
| `U8` to `U14` | Texas Instruments DRV8874PWPR, motor driver, HTSSOP-16 with thermal pad | C1855818 *E* | one per motor; PH/EN mode; current limit 2.93 A set by 1.5 kΩ + 1 kΩ on its `IPROPI` pin |
| `U2`, `U6` | Nexperia 74HC595D,118, shift register, SOIC-16 | C5947 | direction lines and the multiplexer's select lines |
| `U3`, `U5` | Nexperia 74HC165D,653, shift register, SOIC-16 | C5613 *E* | fault, switch and e-stop inputs. **Nexperia only, no substitute**: the switch lines rely on its inputs tolerating a level above its supply through 1 kΩ |
| `U7` | Texas Instruments CD74HC4051M96, 8-channel analogue multiplexer, SOIC-16 | C132065 *E* | the seven motor currents and the motor rail into one ADC pin |
| `D29` to `D38`, `D45` to `D48` | Nexperia BAT54S,215, dual Schottky clamp, SOT-23 | C47546 *E* | the encoder lines' clamps to ground and 3.3 V |
| `Q3`, `Q4` | 2N7002, SOT-23 | C8545 | `Q3` the drivers' sleep gate (fed from the relay coil), `Q4` the status LEDs' level shifter |
| `D1` | SMBJ15A, TVS diode, SMB | C2760888 *E* | the motor rail's clamp |
| `D4` | 1N4148W, SOD-123 | C81598 | across the relay coil |
| `D20` | B5819W, Schottky, SOD-123 | C8598 | the board's 5 V to the DevKit's 5V pin |
| `F2` | Littelfuse 0466.500NRHF, 0.5 A fast, 1206 | C151133 *E* | the e-stop chain's fuse; spares are on the Digi-Key list |
| `F3` | Littelfuse MINISMDC150F-2, resettable 1.5 A hold, 1812 | C9308 *E* | the board's 5 V |
| `R84` to `R88`, `R113`, `R114` | 47 Ω 1 W, 2512 | C15261 *E* | encoder LED supply, one per motor; a short of the LED line puts 0.53 W in one |
| 16 lamps | green LED, 0805 | C2297 | the direction lamps and the two power lamps |
| 7 lamps | red LED, 0603 | C2286 | the fault lamps |
| capacitors | 10 µF 25 V 0805, 100 nF, 22 nF, 10 nF, 1 nF 0603 | C15850, C14663, C21122, C57112, C1588 | |
| resistors | 0603: 1k, 1.5k, 2.2k, 3k, 4.7k, 10k, 30k, 47k, 100k, 470 | see `fab/jlcpcb_bom.csv` | `R5` is 100 kΩ: the rail divider (÷ 11) keeps the multiplexer's input inside its 3.3 V supply up to 33 V |

## Fitted by hand: the Digi-Key list

Upload `fab/digikey_cart_two_boards.csv` with "Upload a list" (columns: quantity, part number, reference). Quantities
are for two boards; spares are noted.

| Qty | Digi-Key | Maker, part | Designator, what it is |
| --- | --- | --- | --- |
| 2 | `1195-2568-ND` | HARTING 09 66 552 6616 | `J1` 50-pin D-sub socket, right angle, threaded flange (the robot cable). No screwlock kits needed |
| 2 | `732-4787-ND` | Würth Elektronik 649024227222 | `J2` ATX 24-pin header, 4.2 mm, right angle, plastic pegs (the PC supply) |
| 2 | `F5197-ND` | Littelfuse 178.6165.0001 | `F1` ATO blade fuse holder |
| 5 | `F4197-ND` | Littelfuse 0287005.PXCN | `F1` ATO fuse **5 A** (2 + 3 spare) |
| 2 | `102-6203-ND` | Same Sky TB007-508-02BE | `J4` screw terminal, 2 pins, 5.08 mm (the external e-stop) |
| 3 | `Z2926-ND` | Omron G2RL-2 DC12 | `K1` relay, two poles, 8 A, 12 V coil (2 + 1 spare) |
| 3 | `SW414-ND` | Omron B3F-4055 | `SW3` MOTORS ON tactile switch, 12 × 12 mm (2 + 1 spare) |
| 2 | `SW455-ND` | Omron B32-1310 | `SW3` key top, 12 × 12 mm, black |
| 4 | `S7020-ND` | Sullins PPTC221LFBN-RC | `U1` 1×22 female socket strips for the DevKit, two a board |
| 10 | `493-14474-ND` | Nichicon UHV1V221MPD | `C1`, `C2` 220 µF 35 V, radial, 8 mm can, 3.5 mm lead pitch (4 + spares) |
| 1 | `S1011EC-40-ND` | Sullins PRPC040SAAN-RC | 1×40 pin header strip: cut 2-pin pieces for `JP1` and `J6` (four pieces for two boards) |
| 10 | `S9001-ND` | Sullins SPC02SYAN | `JP1` jumper shunt (2 + spares) |
| 2 | `277-1266-ND` | Phoenix Contact 1715789 (MKDS 1,5/ 8-5,08) | `J8` screw terminal, 8 pins, 5.08 mm (the rail's cable) |
| 5 | `F5694CT-ND` | Littelfuse 0466.500NRHF | spare chain fuses `F2` (the factory solders one on each board; a short in the e-stop wiring blows it; replaced with an iron) |

Per board that is twelve positions (the DevKit's two socket strips counted as one) and 163 solder joints (`fab/hand_assembled.csv` lists them with their joints).

## Not on either list

- **The ESP32-S3-DevKitC-1**, one per board, plugged into `U1`'s sockets. Which variants fit, and which must not be
  used, is in the [README](README.md), "What you need".
- **The optional status LEDs** `D21` to `D28`, `D49`, `D50` (2.0 × 2.0 mm, WS2812 type, fitted by hand with hot air;
  the README says which parts fit and what they need).
- The ATX PC power supply, the external emergency-stop button and its cable, the rail's adapter from its D9 plug to
  `J8`, six M3 screws, washers and standoffs, two micro-USB data cables.

## Makes that fit, and rules when substituting

- **`F1`: a 5 A ATO blade fuse, never larger.** One pin of the 24-pin connector is rated 6 A and the pins are not
  designed to share current. The label on the board reads "5A MAX".
- **`J2`, the 24-pin supply header.** The footprint is Molex's 5569-24A2 pattern (Mini-Fit Jr, right angle, with
  pegs). Three makes fit it:
  - **Würth Elektronik 649024227222** (Digi-Key `732-4787-ND`), on the list. Its data sheet gives the footprint's
    hole pattern value for value, 6 A at most per contact; it does not state that the part mates with the PC
    supply's plug, so **try the plug in a loose header before soldering**. Digi-Key showed 10 in stock with a 14-week
    lead time on 23 November 2026, and the list asks for exactly two.
  - **Molex 39-30-1240** (`0039301240`, Digi-Key `WM1362-ND`): on the same Molex drawing (55690002-SD) as the part
    the footprint was drawn for, with the same hole layout; its housing is rated UL94 V-2 instead of V-0. Molex ships
    it bulk-packed: look at the pins when it arrives.
  - **Molex 39-30-0240** (`0039300240`, Digi-Key `WM21362-ND`), the part the footprint was drawn for (out of stock at
    Digi-Key in November 2026).

  Whatever the make, **its pegs must be plastic**: 12 V and e-stop copper pass 0.25 to 0.4 mm from the peg holes, and
  a metal board-lock would join them. The same holds for the fuse holder's peg.
- **`J1`, the robot connector.** Two makes fit:
  - **HARTING 09 66 552 6616** (Digi-Key `1195-2568-ND`), on the list. Its flange is threaded (no screwlock kits) and
    reaches below the board's top face, so **the board's bottom edge steps back 2.8 mm over 69 mm under the
    connector**. Its fit was checked against HARTING's 3D model, not a dimensioned drawing or a real part: **seat one
    in a bare board before soldering.**
  - **Amphenol DD50S565MTLF** (Digi-Key `DD50S565MTLF-ND`), the part the footprint was drawn for, with two 4-40
    screwlock kits (`1195-2607-ND`, one per side): its flange stands in front of the edge and fits either way. It was
    backordered at Digi-Key in November 2026.
- **`J4`:** the footprint is Phoenix Contact's MKDS 1,5/2-5,08 (1715721): two pins 5.08 mm apart in 1.3 mm holes.
  The Same Sky TB007-508-02BE on the list has the same pitch; its drawing was not checked, so check the pins on
  arrival. If neither fits, two wires can be soldered in its place.
- **`J8`:** Phoenix Contact 1715789; its footprint was checked against Phoenix's data sheet (8 poles at 5.08 mm,
  1.3 mm holes). Wire entry towards the board's edge.
- **`C1`, `C2`:** any 220 µF 35 V aluminium electrolytic, radial, 8 mm can, 3.5 mm lead pitch, 11.5 mm high or more,
  105 °C (for example Nichicon UHV1V221MPD, Panasonic EEU-FR1V221 or Rubycon 35ZLH220MEFC8X11.5). Nichicon marks the
  UHV1V221MPD "Not For New Designs"; it was in stock.
- **`K1`:** Omron G2RL-2 DC12 (two change-over poles, 8 A, 12 V coil). The board uses both poles: one for the motor
  rail, one for the latch.
- **`U3`, `U5`:** Nexperia 74HC165 only (above).
- **The DevKit:** any `ESP32-S3-DevKitC-1` that is not a part ending in V.
