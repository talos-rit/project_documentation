# Purchase options: Talos controller and driver modules

Prices in USD, from carts built on 18 November 2026. Nothing on this page has been ordered yet.

Every option buys the same hardware, enough for one robot with spares: **5 controller boards** (the minimum order) and
**10 driver modules**. The options differ in who solders them.

Both boards are first prototypes: reviewed, not yet built or powered. If the driver modules misbehave on the bench, a
commercial driver (Pololu G2 18v17, about $45 each) fits the same sockets.

## The three configurations

| | A: assembled | B: all by hand | C: modules assembled |
| --- | --- | --- | --- |
| Controller boards | JLCPCB assembles 2, 3 stay bare | 5 bare | 5 bare |
| Driver modules | JLCPCB assembles all 10 | 10 bare | JLCPCB assembles all 10 |
| Team solders on the controller | 36 parts, about 200 joints | about 200 parts, about 820 joints | about 200 parts, about 820 joints |
| Team solders on each module | 16 header pins | 23 surface-mount parts and 16 pins | 16 header pins |
| Extra tools | none | stencil, solder paste, hot air | none |
| Orders to place | JLCPCB, Digi-Key | JLCPCB, Digi-Key, Mouser, LCSC | JLCPCB, Digi-Key, Mouser |
| **Total** | **$513** | **$548 to $551, plus import duties** | **$521** |

**Recommended: A.** It is the cheapest, needs the least soldering, and delivers two assembled controller boards. B is the
dearest, and each module's main chip has to be soldered with paste and hot air, where a bad joint cannot be seen.

## Cost by order

| Order | A | B | C |
| --- | --- | --- | --- |
| JLCPCB: boards, assembly, shipping | 214.98 | 97.96 | 136.70 |
| Digi-Key | 236.60 | 174.32 | 174.32 |
| Mouser | none | 192.94 | 148.50 |
| LCSC | none | 21.85 to 24.85, plus import duties | none |
| Cables, mounting hardware, enclosure (estimate) | 61 | 61 | 61 |
| **Total** | **512.58** | **about 548 to 551, plus import duties** | **about 520.52** |

- Digi-Key and Mouser figures include shipping and tariff as their carts showed them. Digi-Key's include sales tax. Mouser shows
  sales tax only at checkout, so 8 % is assumed, the rate in Digi-Key's carts.
- A's JLCPCB figure includes a $20 coupon the cart applied. Without it, C would be about $12 cheaper than A.
- Not included: a PC power supply, a Raspberry Pi, an external emergency-stop button and bench instruments, which the lab has.

## What is in each order

### JLCPCB

| Line | A | B | C |
| --- | --- | --- | --- |
| Controller boards, 5, rev 0.5, lead-free HASL | 34.25 | 34.25 | 34.25 |
| Assembly of 2 controller boards, parts included | 91.12 | | |
| Driver-module boards, 10, ENIG | 22.77 | 22.97 | 22.77 |
| Assembly of 10 driver modules, parts included | 42.04 | | 42.04 |
| Solder-paste stencil for the modules | | 3.10 | |
| Shipping (estimate) | 44.80 | 37.64 | 37.64 |
| Coupon | −20.00 | | |
| **Total** | **214.98** | **97.96** | **136.70** |

The files to upload are in [`fab/`](fab/) for the controller and in [`../talos_driver_module/fab/`](../talos_driver_module/fab/)
for the module.

### Parts carts

| Cart | Options | Lines | Parts | Shipping, tariff, tax | Total | List |
| --- | --- | --- | --- | --- | --- | --- |
| Digi-Key | A | 22 | 204.18 | 32.42 | 236.60 | [`orders/digikey_A.csv`](orders/digikey_A.csv), [cart](https://digikey.com/short/5q3j0jj4) |
| Digi-Key | B, C | 34 | 151.05 | 23.27 | 174.32 | [`orders/digikey_B_C.csv`](orders/digikey_B_C.csv), [cart](https://digikey.com/short/4vjz8z0f) |
| Mouser | B, C | 18 | 114.53 | 33.97 | 148.50 | [`orders/mouser_B_C.csv`](orders/mouser_B_C.csv) |
| Mouser, added for B: 12 gate drivers and solder paste | B | 2 | 41.15 | 3.29 | 44.44 | the same file, marked "B only" |
| LCSC: the other driver-module parts | B | 12 | 15.34 | 6.51 to 9.51, plus import duties | 21.85 to 24.85 | [`orders/lcsc_B_modules.csv`](orders/lcsc_B_modules.csv) |

Small items bought locally, estimated: USB data cables 10, panel-mount USB extension 8, M3 screws, washers and standoffs 18,
3D-printed enclosure 15, hook-up wire 5, cable adapter for the slide rail 5.

## Why assembled is cheaper

JLCPCB's assembly price includes the parts, which it buys at volume prices.

- **Controller:** assembling two boards costs $91. Building one by hand instead means buying about $86 more parts at retail.
- **Driver modules:** assembling ten costs $42. The loose parts, stencil, paste and shipping to build ten by hand cost about $70.

## Price check across suppliers

Every loose part (66 lines) was priced at Digi-Key, Mouser and LCSC, and each split was confirmed in a real cart.
[`orders/price_check.csv`](orders/price_check.csv) has every line.

- Five lines are $155 of the $282 hand-build parts list and have few sellers: the emergency-stop switch, the counter chips,
  the current-measurement boards, the status lights and the switch's nameplate.
- **Mouser** is cheaper on 18 lines of the hand-build list and has the emergency-stop switch in stock. Splitting B's and C's
  controller parts between Digi-Key and Mouser saves about $8. For A, one Digi-Key order is about $3 cheaper than a split.
- **LCSC** charged more in the cart than it quoted (eleven controller lines: $105, against $95 at Digi-Key), takes 10 to 18
  business days and leaves import duties to the buyer. It is used only for B's small module parts, where it is much cheaper.

## Timing

- JLCPCB estimated shipping on 21 November 2026, then 2 to 4 days in transit.
- The counter chips ship from their maker about 6 days after the Digi-Key order.
- **A:** Digi-Key does not stock the emergency-stop switch (estimated 30 December 2026). Until it arrives, an external
  emergency-stop button on the board's terminal does the same job. The robot cable connector was expected on 24 November.
- **B and C:** everything else is in stock except one 100 µF capacitor.
- **B:** the LCSC parcel takes 10 to 18 business days.
