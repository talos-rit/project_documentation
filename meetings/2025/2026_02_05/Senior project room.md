Zoom: https://rit.zoom.us/rec/share/PzXYD_5NNgtTp0XUuwwWziur5TzonoUwEe4FTfsi9ogdbhZjWO9e_G29EqO8_xDF.0EQfoD5fpLzuBSJs

## Key takeaways

- The team discussed ESP driver development and hardware issues
- PCA driver has been completed and is in review, almost merged
- Commander WebRTC features are in review, enabling remote video feed viewing
- The team spent significant time troubleshooting serial communication issues between an ESP32 board and a robot
- Hardware driver installation for Mac was identified and resolved (Chinese CH343 driver)
- Team members worked on configuring and testing oscilloscope readings to debug communication problems
## Team updates


| **Speaker** |                                                         **Completed**                                                          |            **In Progress**            | **Plan** |             **Blockers**              |
|-------------|--------------------------------------------------------------------------------------------------------------------------------|---------------------------------------|----------|---------------------------------------|
| **Hiroto**  | - PCA driver reported to ESP (in review, almost merged)- Figured out calibration for current sense with uneven shunt resistors | - Commander WebRTC features in review | - N/A    | - ESP requires hardware redevelopment |
| **Ryan**    | - N/A                                                                                                                          | - Working on ESP driver development   | - N/A    | - N/A                                 |
| **Kai**     | - N/A                                                                                                                          | - Onboarded to ESP                    | - N/A    | - N/A                                 |
| **Connor**  | - Downloaded ESP                                                                                                               | - Working on ESP onboarding           | - N/A    | - N/A                                 |

## High risks

- ESP requires hardware redevelopment (ongoing risk)
- Serial communication issues between ESP32 and robot
- Driver compatibility issues on Mac systems
- Potential issues with RX/TX pin configuration
## Action items

- **Hiroto**
    - [ ] Complete onboarding to ESP
    - [ ] Work on ESP driver development
- **Kai**
    - [ ] Work on ESP driver development this weekend
    - [ ] Implement alert ready functionality for the ADS1015 driver
- **Team**
    - [ ] Add CH343 driver installation to documentation for Mac users
    - [ ] Update documentation with clearer ESP setup instructions
    - [ ] Implement handshake mechanism between devices
    - [ ] Resolve serial communication issues with ESP32
