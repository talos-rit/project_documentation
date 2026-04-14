Zoom: https://rit.zoom.us/rec/share/k0kCkmYyzHljid5GTH9dCG9DeQ7lgByAZ9yOjSbGywgSYTc8uc2TcLeCW-ZEl3fe.mNYFMOCeMdMWpFN2

## Key takeaways

- The team discussed motor driver options for the robot project, comparing L298 and TB6612 drivers
- Current measurements showed the motors draw about 0.8 amps normally, potentially up to 1.4 amps when stalled
- The team needs to create an inventory spreadsheet of existing hardware and purchase additional components
- Development environment setup for Operator was discussed, with challenges for macOS compatibility identified
- The team is working on refactoring the connection management to handle multiple robots(commander)
## Team updates

| **Speaker** |                           **Completed**                           |                         **In Progress**                          |                                  **Plan**                                   |                                                         **Blockers**                                                         |
|-------------|-------------------------------------------------------------------|------------------------------------------------------------------|-----------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------|
| **Hiroto**  | N/A                                                               | - Working on refactoring Commander                               | N/A                                                                         | N/A                                                                                                                          |
| Kai         | N/A                                                               | - Working on refactoring the configuration file on commander     | N/A                                                                         | - Refactoring work could be a blocker for config but it is relatively refactored for usage, so it should be fine to continue |
| **Connor**  | - Tested motor current draw (0.8 amps)                            | - Setting up development environment for Operator                | - Create inventory spreadsheet- Compare motor driver options                | - Hardware procurement                                                                                                       |
| Ryan        | - Merged PR that Brooke started- Merged pull request for Operator | - Working on connection management- Implementing scheduler class | - Separate image processing from object detection for better UI performance | N/A                                                                                                                          |

## High risks

- Development environment compatibility issues - Operator code only works on Linux, not macOS
- Hardware procurement delays could impact development progress
- Multiple processes for image processing and object detection might overload the system
## Action items

- **Connor**
    - [ ] Create Excel spreadsheet to compare different motor driver options
    - [ ] Take inventory of existing hardware components
    - [ ] Figure out development environment setup for Operator
- **Connor?**
    - [ ] Approve pull request for driver
- **Hiroto**
    - [ ] Work on scheduler class implementation
    - [ ] Explore separating image processing from object detection pipeline
    - [ ] Implement UI improvements (convert buttons to dropdown menu)
- **J**
    - [ ] Purchase motor drivers (likely TB6612 hat)
    - [ ] Determine if all motors need to run simultaneously (affects power requirements)
    - [ ] Update Malkovsky about hardware needs (budget of ~$200)
