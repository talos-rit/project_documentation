# Motor Ratios
This is a table of the mechanical aspects of the ER4pc robot (Bluey). Gear ratios found by counting teeth combined with datasheet. Motor ratios found in datasheet, also verified visually on robot.

Resolution figures represent degrees/millimeters traveled in one encoder step. 
    X1 represents only tracking the spaces in between the wheel (20 positions per wheel rotation).
    X4 represents tracking all 4 different positions of the quadrature system (80 positions per wheel rotation).

| Motor | Purpose | Gear Ratio | Motor Ratio | Bounds | Total Encoder Revolutions Over Bounds | X1 Resolution | X4 Resolution |
| - | - | - | - | - | - | - | - |
| 1 | Base | 24:120 | 127.1:1 | 310&deg; | 547.2 | 0.028&deg; | 0.0071&deg;
| 2 | Shoulder | 20:72 | 127.1:1 | +130&deg; -35&deg; | 209.7 | 0.039&deg; | 0.0098&deg;
| 3 | Elbow | 20:72 | 127.1:1 | &pm;130&deg; | 330.5 | 0.039&deg; | 0.0098&deg;
| 4 | Wrist 1 | 12:23 | 65.5:1 | &pm;130&deg;? | 125.5? | 0.14&deg; | 0.035&deg;
| 5 | Wrist 2 | 12:23 | 65.5:1 | &pm;130&deg;? | 125.5? | 0.14&deg; | 0.035&deg;
| 6 | Gripper | ??? | 19.5:1 | 65mm (75mm without pads) |
| 0 | Slide | 5mm per rev | 5.9:1 | 87mm | 102.7 | 0.042mm | 0.011mm
