# Talos API ICD

## Changelog

| Date | Name | Description |
|---|---|---|
| 2024-09-26 | Brooke Leinberger    | Init |
| 2024-10-03 | Brooke Leinberger    | Device IDs are redundant; Subbed out for ids that are unique message-to-message; Added reserved flag field to wrapper |
| 2024-10-21 | Devan Kavalchek      | Add home and fix unsigned integers that should be signed. |
| 2024-12-03 | Noah Carney          | Added set speed command structure and assigned it values |
| 2025-02-06 | Alex Vernes          | Added idle mode calls |
| 2025-02-25 | Brooke Leinberger    | Added hardware specific hook; Revised idle mode commands; Removed hardware coupling with speed commands |
| 2025-02-27 | Alex Vernes          | Reorganizing command order; added Cartesian Move |

Note: 
- Type names are given by the stdint.h header file in the C standard.
- All data is in big endian format

## Command Wrapper

| Arg           | Type | Description |
|---|---|---|
| Command ID    | UINT32    | Unique ID for individual commands |
| *RESERVED*    | UINT16    | *RESERVED* |
| Command Value | UINT16    | Command for device to carry out |
| Length        | UINT16    | Length of Payload |
| Payload       | UINT8[]   | Command Info |
| CRC           | UINT16    | Checksum |

*Note*: Command IDs are unique to the indiviual commands. If the director were to send out two identical commands, they would share a Command Value and payload, but would have unique Command IDs. Each command exchange will have 2 messages: a command request message, and a command return message. Half of the ID's will therefore be return messages, and should correspond to a specific command request message. Specifically, the upper 31 bits should denote the overall command exchange, while the least significant bit is a '0' for command, and a '1' for response (i.e. command issued with command ID 0 gets a return ID of 1, command ID of 54 gets a return ID of 55, etc).

## Handshake

### Brief
This command is used by the Talos Command Interface (primarily the Director, but also the manual interface) on the Operator in order to prepare for giving further commands to the operator. Upon the Operator receiving this command, it responds with the version of the Talos Operator software it's running. The Command Interface can then use this information to determine the best commands to give this unit, or decide if it's incompatible.

### Send
**Command Value**: 0x0000

| Arg           | Type | Description |
|---|---|---|
| OID | UINT16 | Unique ID for Operator-Arm model version connection |

### Receive
**Command Value**: 0x8000

| Arg | Type | Description |
|---|---|---|
| Return Code   | UINT16 | |
| Major Version | UINT8  | |
| Minor Version | UINT8  | |


## Polar Pan (Discrete)

### Send
**Command Value**: 0x0001

| Arg | Type | Description |
|---|---|---|
| Delta Azimuth     | INT32 | Requested change in azimuth |
| Delta Altitude    | INT32 | Requested change in altitude |
| Delay (ms)        | UINT32 | How long to wait until executing pan |
| Time              | UINT32 | How long the pan should take to execute |

### Receive
**Command Value**: 0x8001

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error

## Home

### Send
**Command Value**: 0x0002

| Arg | Type | Description |
|---|---|---|
| Delay (ms) | UINT32 | How long to wait until executing home |

### Receive
**Command Value**: 0x8002
| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error


## Polar Pan (Continuous START)
Starts/maintains a continuous polar pan rotation.

### Send

**Command Value**: 0x0003

| Arg | Type | Description |
|---|---|---|
| Moving Azimuth     | INT8 | -1, 0, or 1 |
| Moving Altitude    | INT8 | -1, 0, or 1 |

 The values in the body describe whether or not the arm is rotating in a given direction. 
 1 rotates counter-clockwise along the axis of movement, -1 rotates clockwise along the axis of 
 movement and 0 means no rotation.

### Receive
**Command Value**: 0x8003

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error


## Polar Pan (Continuous STOP)
Stops a continuous polar pan rotation.

### Send
**Command Value**: 0x0004

**No body sent**

### Receive
**Command Value**: 0x8004

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error


## Cartesian Move (Discrete)

### Send
**Command Value**: 0x0005

| Arg | Type | Description |
|---|---|---|
| Delta X       | INT32 | Tenths of millimeters on X-axis |
| Delta Y       | INT32 | Tenths of millimeters on Y-axis |
| Delta Z       | INT32 | Tenths of millimeters on Z-axis |
| Delay (ms)    | UINT32 | How long to wait until executing pan |
| Time          | UINT32 | How long the pan should take to execute |

### Receive
**Command Value**: 0x8005

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error


## Cartesian Move (Continuous START)
Starts/maintains a continuous cartesian movement.

### Send

**Command Value**: 0x0006

| Arg | Type | Description |
|---|---|---|
| Moving X     | INT8 | -1, 0, or 1 |
| Moving Y    | INT8 | -1, 0, or 1 |
| Moving Z     | INT8 | -1, 0, or 1 |

 The values in the body describe whether or not the arm is rotating in a given direction. 
 1 rotates counter-clockwise along the axis of movement, -1 rotates clockwise along the axis of 
 movement and 0 means no rotation.

### Receive
**Command Value**: 0x8006

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error


## Cartesian Move (Continuous STOP)
Stops a continuous cartesian move.

### Send
**Command Value**: 0x0007

**No body sent**

### Receive
**Command Value**: 0x8007

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error


## Execute Hardware Operation
Some operations require high coupling with the specifics of the hardware (e.g. axis-by-axis positions).
Such operations should be defined by a separate companion ICD, to avoid coupling the high level API with the hardware

### Send
**Command Value**: 0x0008

| Arg | Type | Description |
|---|---|---|
| Subcommand Value | UINT16 | Command for function in hardware specific ICD |
| RESERVED | UINT32 | RESERVED |
| Payload | UINT8[] | Payload defined by hardware specific ICD |

### Receive
**Command Value**: 0x8008

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error |
| Subcommand Value | UINT16 | Command for function in hardware specific ICD |
| Payload | UINT8[] | Payload defined by hardware specific ICD |


## Get Speed
Command to get the speed of all axes on Talos

### Send
**Command Value**: 0x0009
No body

### Receive
**Command Value**: 0x8009
| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error
| Speed | UINT8 | What to the default speed of the Talos is |


## Set Speed
Command to set the speed of all axes on Talos to the received number (uint8)

### Send
**Command Value**: 0x000A
| Arg | Type | Description |
|---|---|---|
| Speed | UINT8 | The speed of all axes to set on Talos |

### Receive
**Command Value**: 0x800A
| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error


## Save Position
Saves a position. If this command is sent with a name that already exists, it will be overwritten with the new arguments.
If reference is an empty string (length of 0), the default value will be used (can be reconfigured, unconfigured default is empty string).
If the reference string is empty, the Anchor value is ignored and the position is always treated as if anchor is set to false.

### Send
**Command Value**: 0x000B

| Arg | Type | Description |
|---|---|---|
| Name len | UINT8 | length of name field; must be more than zero |
| Name | CHAR[] | Name descriptor for the position (non null terminated) |
| Anchor | BOOLEAN | Whether the position will move relative to the parent position (0x01 for True, 0x00 for False) |
| Parent len | UINT8 | Length of the parent array; if 0 |
| Parent | CHAR[] | Another previously saved position to act as a parent (refernce) position |

### Receive
**Command Value**: 0x800B

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error |

## Delete Position
Given a position name, deletes that position information. 

### Send
**Command Value**: 0x000C

| Arg | Type | Description |
|---|---|---|
| Name len | UINT8 | length of name field; must be more than zero |
| Name | CHAR[] | Name descriptor for the position (non null terminated) |

### Receive
**Command Value**: 0x800C

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error

## Go to Position
Move to a pre-defined position. 

### Send
**Command Value**: 0x000D

| Arg | Type | Description |
|---|---|---|
| Name len | UINT8 | length of name field; must be more than zero |
| Name | CHAR[] | Name descriptor for the position (non null terminated) |

### Receive
**Command Value**: 0x800D

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error

## Set Polar Position
Defines a position in terms of polar coordinates

### Send
**Command Value**: 0x000E

| Arg | Type | Description |
|---|---|---|
| Name len | UINT8 | length of name field; must be more than zero |
| Name | CHAR[] | Name descriptor for the position (non null terminated)|
| Delta | INT32 | Tenths of degrees on delta axis |
| Azimuth | INT32 | Tenths of degrees on azimuth axis |
| Radius | INT32 | Tenths of distance to extend outwards |

### Receive
**Command Value**: 0x800E
| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error

## Get Polar Position
Returns the polar coordinates of a named position

### Send
**Command Value**: 0x000F
| Arg | Type | Description |
|---|---|---|
| Name len | UINT8 | length of name field; must be more than zero |
| Name | CHAR[] | Name descriptor for the position (non null terminated)|


### Receive
**Command Value**: 0x800F
| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error |
| Delta | INT32 | Tenths of degrees on delta axis |
| Azimuth | INT32 | Tenths of degrees on azimuth axis |
| Radius | INT32 | Tenths of distance to extend outwards |


## Set Cartesian Position
Defines a position in terms of cartesian coordinates

### Send
**Command Value**: 0x0010

| Arg | Type | Description |
|---|---|---|
| Name len | UINT8 | length of name field; must be more than zero |
| Name | CHAR[] | Name descriptor for the position (non null terminated)|
| X | INT32 | Tenths of millimeters on X-axis |
| Y | INT32 | Tenths of millimeters on Y-axis |
| Z | INT32 | Tenths of millimeters on Z-axis |

### Receive
**Command Value**: 0x8010
| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error |

## Get Cartesian Position
Returns the cartesian coordinates of a named position

### Send
**Command Value**: 0x0011

| Arg | Type | Description |
|---|---|---|
| Name len | UINT8 | length of name field; must be more than zero |
| Name | CHAR[] | Name descriptor for the position (non null terminated)|

### Receive
**Command Value**: 0x8011
| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error |
| X | INT32 | Tenths of millimeters on X-axis |
| Y | INT32 | Tenths of millimeters on Y-axis |
| Z | INT32 | Tenths of millimeters on Z-axis |
