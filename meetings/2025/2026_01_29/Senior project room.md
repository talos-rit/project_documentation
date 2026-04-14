Zoom: https://rit.zoom.us/rec/share/ug2G0JJuq2Hvddc31DKm2SPcr0HePeTg69E_zzYkXxsQC3eJ7S16CHBVxpmo-0ig.h7VQf89umZ93XNb5
## Key takeaways

- The team decided to switch from Raspberry Pi to ESP32 for motor control due to encoder pulse detection limitations
- Successfully read current measurements using the analog-digital converter chip
- Implemented PIAV library which significantly reduced latency issues in Commander
- Started developing ESP driver project for motor control
- Registered for Imagine RIT event
- Tested the speed of Pi with and without MCP I/O expander; Pi kept up without MCP but had limitations with it
## Team updates


| **Speaker** |                                                                 **Completed**                                                                 |                                     **In Progress**                                     |                                **Plan**                                 |                      **Blockers**                      |
|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------|-------------------------------------------------------------------------|--------------------------------------------------------|
| **Hiroto**  | - Started implementing ESP driver- Set up ESP dev environment- Tested Pi speed with/without MCP                                               | - Developing ESP driver- Converting driver code to C                                    | - Continue ESP driver development- Implement I2C bus class              | - RIT declined Swain purchase due to card restrictions |
| **Connor**  | - Implemented Commander single speaker feature- Found and implemented PIAV library to replace OpenCV- Fixed latency issues in video streaming | - Working on Conan integration for dependencies- Fixing CMake for dependencies          | - Test PIAV implementation- Profile Commander with scaling              | - N/A                                                  |
| **Kai**     | - Made progress on PiSide 6 migration- Created terminal_talos.py script                                                                       | - Working on UI styling issues                                                          | - Fix dark mode issues in UI                                            | - Some UI elements not displaying correctly            |
| **Ryan**    | - Successfully read current measurements with ADS- Tested current sensing                                                                     | - Working on current measurement calibration- Developing code to capture current spikes | - Finish PCA implementation- Start on socket side of ESP implementation | - N/A                                                  |

## High risks

- Redesigning the driver component for ESP32 instead of Raspberry Pi
- Potential integration issues between Raspberry Pi and ESP32
- Time constraints for implementing the new ESP-based solution
- UI styling issues in Commander with PiSide 6 migration
## Action items

- **Hiroto**
    - [ ] Continue developing ESP driver
    - [ ] Implement I2C bus class
    - [ ] Create GitHub pages website for ESP flasher
    - [ ] Create bill of materials for ESP implementation
- **Connor**
    - [ ] Test PIAV implementation
    - [ ] Profile Commander with scaling
    - [ ] Fix UI styling issues in Commander
    - [ ] Merge open PRs for Commander
- **Kai**
    - [ ] Fix dark mode issues in PiSide 6 migration
    - [ ] Complete profile setup for Imagine RIT
- **Ryan**
    - [ ] Finish PCA implementation
    - [ ] Start working on socket side of ESP implementation
    - [ ] Calibrate current measurements
    - [ ] Develop code to capture current spikes
- **All team members**
    - [ ] Complete individual registration for Imagine RIT
    - [ ] Fill out time cards
