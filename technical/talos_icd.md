# Talos API ICD

## Changelog

| Date | Name | Description |
|---|---|---|
| 2024-09-26 | Brooke Leinberger | Init |
| 2024-10-03 | Brooke Leinberger | Device IDs are redundant; Subbed out for ids that are unique message-to-message; Added reserved flag field to wrapper |
| 2024-10-21 | Devan Kavalchek   | Add home and fix unsigned integers that should be signed. |
| 2024-12-03 | Noah Carney       | Added set speed command structure and assigned it values |

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


## Save Position
Saves a position [FINISH ME]

### Send
**Command Value**: 0x0006

| Arg | Type | Description |
|---|---|---|
| Name | CHAR[] | Name descriptor for the position |
| Reference | UINT32 | Another previously saved position to act as a reference position
| Anchor | BOOLEAN | Whether the position will move relative to the reference position

### Receive
**Command Value**: 0x8006

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error


## Delete Position
Given a position name, deletes that position information. 

### Send
**Command Value**: 0x0007

| Arg | Type | Description |
|---|---|---|
| Name | CHAR[] | Name of the position to be deleted |

### Receive
**Command Value**: 0x8007

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error

## Go to Position
Move to a pre-defined position. 

### Send
**Command Value**: 0x0008

| Arg | Type | Description |
|---|---|---|
| Name | CHAR[] | Name of the position to move to |

### Receive
**Command Value**: 0x8008

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error

## Get Current Position
Gets the current position of all axes

### Send
**Command Value**: 0x0009

**no body sent**

### Receive
**Command Value**: 0x8009

| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error |
| Axis 1 position | INT32 | Current position of Axis 1 |
| Axis 2 position | INT32 | Current position of Axis 2 |
| Axis 3 position | INT32 | Current position of Axis 3 |
| Axis 4 position | INT32 | Current position of Axis 4 |
| Axis 5 position | INT32 | Current position of Axis 5 |
<!-- There are 6 axes right? -->

## Set Polar Position
Defines a position in terms of polar coordinates

### Send
**Command Value**: 0x000A

| Arg | Type | Description |
|---|---|---|
| Name | CHAR[] | Name descriptor for the position |
| Delta | INT32 | Degrees on delta axis |
| Azimuth | INT32 | Degrees on azimuth axis |
| Radius | INT32 | Distance to extend outwards |

### Receive
**Command Value**: 0x800A
| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error


## Set Cartesian Position

### Send
**Command Value**: 0x000B

| Arg | Type | Description |
|---|---|---|
| Name | CHAR[] | Name descriptor for the position |
| X | INT32 | Degrees on X-axis |
| Y | INT32 | Degrees on Y-axis |
| Z | INT32 | Degrees on Z-axis |

### Receive
**Command Value**: 0x800B
| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error


## Set Speed

### Brief
Command to set the speed of all axes on the scorbot to the received number (uint8)

### Send
**Command Value**: 0x0005
| Arg | Type | Description |
|---|---|---|
| Speed | UINT8 | What to set the speed of all axes to on the scorbot |
| Axis | UINT8  | Axis to change (for all axes send 0)

### Receive
**Command Value**: 0x8005
| Arg | Type | Description |
|---|---|---|
| Return Code | UINT16 | Reports success/error