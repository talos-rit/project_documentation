# Purchase options: Talos controller prototype 0.5

Two carts, both made on **23 November 2026** for these files, both **not paid** when this page was written. Prices in
USD as each cart showed them. Together they buy **two complete boards**, with spares of the cheap parts.

| Order | What it buys | Total |
| --- | --- | --- |
| JLCPCB | five bare boards, two of them assembled with every surface-mount part | **108.28** with a USD 9 coupon, **117.28** without |
| Digi-Key | the through-hole parts for two boards, with spares | **128.36** with shipping, estimated tariff and tax |
| **Both** | | **236.64**, or **245.64** without the JLCPCB coupon |

Not in either cart: the ESP32-S3-DevKitC-1 boards, the optional status LEDs, the PC power supply, the external
emergency-stop button and its cable, the rail's cable adapter, screws and standoffs, USB cables. [`parts.md`](parts.md)
says what fits.

## JLCPCB

The Gerber zip, `fab/jlcpcb_bom.csv` and `fab/jlcpcb_cpl.csv` of this folder, uploaded to JLCPCB's order page. Its
parts page read "29 parts detected, 29 parts confirmed", each line matched to the part number in the BOM.

| Line | USD |
| --- | --- |
| Five bare boards: two layers, 141 × 139.6 mm, green, 1.6 mm, HASL with lead | 14.55 |
| Economic assembly of two boards, top side, with the parts it places (29 lines, 241 placements a board) | 65.22 |
| Shipping, JLCPCB's estimate | 37.51 |
| **Without a coupon** | **117.28** |
| Coupon on the cart | −9.00 |
| **The cart's total** | **108.28** |

JLCPCB adds its own spares or minimum quantities on some lines (on this cart: the TVS diode 6 for 2, the chain fuse 5
for 2, the resettable fuse 6 for 2, the BAT54S 33 for 28, the 47 Ω 1 W resistor 22 for 14). The cart does not break
the assembly line down further.

## Digi-Key

`fab/digikey_cart_two_boards.csv`, uploaded with "Upload a list". Every line was available, none backordered.

| Qty | Digi-Key | Maker part | What | Unit | Line |
| --- | --- | --- | --- | --- | --- |
| 2 | 1195-2568-ND | HARTING 09665526616 | `J1` 50-pin D-sub socket, right angle | 18.20 | 36.40 |
| 2 | 732-4787-ND | Würth 649024227222 | `J2` ATX 24-pin header, right angle | 2.29 | 4.58 |
| 2 | F5197-ND | Littelfuse 178.6165.0001 | `F1` blade fuse holder | 5.75 | 11.50 |
| 5 | F4197-ND | Littelfuse 0287005.PXCN | `F1` fuse 5 A | 0.36 | 1.80 |
| 2 | 102-6203-ND | Same Sky TB007-508-02BE | `J4` screw terminal, 2 pins | 0.55 | 1.10 |
| 3 | Z2926-ND | Omron G2RL-2 DC12 | `K1` relay | 4.94 | 14.82 |
| 3 | SW414-ND | Omron B3F-4055 | `SW3` tactile switch | 0.52 | 1.56 |
| 2 | SW455-ND | Omron B32-1310 | `SW3` key top | 0.37 | 0.74 |
| 4 | S7020-ND | Sullins PPTC221LFBN-RC | `U1` 1×22 socket strips | 1.36 | 5.44 |
| 10 | 493-14474-ND | Nichicon UHV1V221MPD | `C1`, `C2` 220 µF 35 V | 0.372 | 3.72 |
| 1 | S1011EC-40-ND | Sullins PRPC040SAAN-RC | `JP1`, `J6` 1×40 pin header | 1.23 | 1.23 |
| 10 | S9001-ND | Sullins SPC02SYAN | `JP1` jumper shunt | 0.029 | 0.29 |
| 2 | 277-1266-ND | Phoenix Contact 1715789 | `J8` screw terminal, 8 pins | 4.59 | 9.18 |
| 5 | F5694CT-ND | Littelfuse 0466.500NRHF | `F2` spare chain fuses | 1.032 | 5.16 |
| | | | **Parts** | | **97.52** |
| | | | Shipping (FedEx Ground) | | 8.49 |
| | | | Tariff, Digi-Key's estimate | | 12.84 |
| | | | Sales tax | | 9.51 |
| | | | **The cart's total** | | **128.36** |

The supply header is on the cart twice, for two boards, with no spare; Digi-Key had 10 in stock that day.
