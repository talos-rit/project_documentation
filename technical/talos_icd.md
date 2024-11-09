# Talos API ICD

## Changelog

| Date | Name | Description |
|---|---|---|
| 2024-09-26 | Brooke Leinberger | Init |
| 2024-10-03 | Brooke Leinberger | Device IDs are redundant; Subbed out for ids that are unique message-to-message; Added reserved flag field to wrapper |
| 2024-10-21 | Devan Kavalchek   | Add home and fix unsigned integers that should be signed. |

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
