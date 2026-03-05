Zoom: https://rit.zoom.us/rec/share/HnlQCQBDV4lIvHkZ3T2e_lNjyTYQ01BY0c_1TjtZ_ntZnQ-aOZKzivHMoZgo5Cet.IpDjokiwTNRbLN7j

## Key takeaways

- Successfully tested and implemented ADS controllers for encoder reading
- Added YAML source to app settings for better configuration management
- Implemented CLI argument parsing with Pydantic integration
- Created functionality to toggle bounding boxes via command line arguments
- Discussed hot reloading functionality for app settings
- Addressed issues with hardware connections and wiring for the ESP32 board
- Explored singleton pattern implementation for better code organization
## Team updates


| **Speaker** |                                         **Completed**                                         |             **In Progress**              |                                **Plan**                                |                     **Blockers**                     |
|-------------|-----------------------------------------------------------------------------------------------|------------------------------------------|------------------------------------------------------------------------|------------------------------------------------------|
| **Hiroto**  | - Added VCAM functionality- Implemented CLI argument parsing- Added toggle for bounding boxes | - Working on app settings improvements   | - Complete PR for CLI argument parsing- Add more configuration options | - Some issues with PySide and Ctrl-C handling        |
| **Jay**     | - Tested ADS controllers- Wired up encoders                                                   | - Testing hardware connections           | - Test multiple ADS controllers simultaneously                         | - Mislabeled board with incorrect ground connections |
| **Tyler**   | - Added YAML source to app settings                                                           | - Working on hot reloading functionality | - Implement watchdog for file changes                                  | - N/A                                                |

## High risks

- PySide GUI has random crashes that are difficult to reproduce
- Hardware issues with the ESP32 board having mislabeled pins and ground connections
- Circular imports could cause problems with the singleton implementation
- Multiple termination handlers being called causing segmentation faults
## Action items

- **Hiroto**
    - [ ] Complete PR for CLI argument parsing
    - [ ] Fix issues with PySide and Ctrl-C handling
    - [ ] Update README with instructions for virtual camera setup
- **Jay**
    - [ ] Test multiple ADS controllers simultaneously
    - [ ] Reconstruct hardware setup with correct pin connections
    - [ ] Return Thursday to test the updated hardware setup
- **Tyler**
    - [ ] Implement watchdog for file changes
    - [ ] Complete hot reloading functionality
    - [ ] Add YAML settings source to app settings
- **Ryan**
    - [ ] Test VCAM functionality on Linux
    - [ ] Test if OBS works properly with the virtual camera setup
