Zoom: https://rit.zoom.us/rec/share/ABMHqz3BshMS9CajxW_xGnIY06h7sxnZ5OnT4z5ZzOL7m62vlQFjTwgBdp1MD3gn.jjP6VFZVpvSs_X2j

## Key takeaways

- Hardware freeze implemented on March 31st after successful testing of 4-axis motor control
- All four motors now working with TB6612FNG driver after resolving hardware issues
- Video streaming to Zoom successfully tested with acceptable latency
- Code coverage increased to 57% with 125 unit tests added
- Raspberry Pi required for webcam streaming; ESP32 lacks sufficient power for video processing
- GitHub Copilot token allowance reset at 8pm, team coordinated to maximize AI-assisted coding
- Homing sequence implementation planned but deferred to feature branch due to code freeze timeline

## Team updates


| **Speaker** |                                                                 **Completed**                                                                 |                                    **In Progress**                                    |                                           **Plan**                                            |                                               **Blockers**                                               |
|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------|
| **Hiroto**  | - Video streaming to Zoom working- Terminal UI performance optimization- Fixed PyTorch tensor to NumPy conversion issues                      | - Debugging error messages in detection models- Testing YOLO and MediaPipe models     | - Finalize Commander video streaming for demo- Test on different computers                    | - MediaPipe pose model freezes computer- Frame size zero error intermittently occurs                     |
| **Ryan**    | - Fixed motor driver wiring and direction control- Successfully tested all 4 motors- Implemented hardware with TB6612FNG drivers              | - Switching motor wires to correct directions- Creating hardware documentation README | - Set up Raspberry Pi for webcam streaming- Finalize hardware before freeze                   | - One TB6612FNG driver half-dead, one completely burnt- Overcurrent protection triggering intermittently |
| **Kai**     | - App settings refactoring merged- Removed deprecated MediaPipe pose detection option- Fixed circular import issues                           | - Implementing homing sequence for robot- Resolving model option circular imports     | - Complete homing feature in separate branch- Add gesture control for next team               | - Limit switch implementation complexity- Circular import problems with type hinting                     |
| **Connor**  | - GitHub Actions workflow for PyTest created- Unit test coverage increased from 30% to 57%- Added 125 total tests including integration tests | - Fixing failing test assertions- Implementing branch coverage testing                | - Achieve higher code coverage before code freeze April 10th- Add CI/CD pipeline improvements | - Some tests broke after app settings merge- GitHub Action initially failed locally                      |

## High risks

- Only one fully functional TB6612FNG motor driver remaining; two others damaged or partially functional
- MediaPipe pose model causes system freezes and needs to be deprecated
- Intermittent frame size zero error causing detection to stop, root cause unclear
- Overcurrent protection on motor drivers triggering unexpectedly during operation
- Code freeze deadline April 10th approaching with homing sequence incomplete
- Hardware must remain untouched after freeze to prevent breaking working configuration

## Action items

- **Ryan**
    - [ ] Switch motor wires to correct directions for proper movement
    - [ ] Create comprehensive hardware documentation README file
    - [ ] Set up Raspberry Pi for webcam streaming before Thursday demo
    - [ ] Desolder burnt TB6612FNG driver if backup needed
    - [ ] Document motor brake vs coast mode settings for future teams
- **Hiroto**
    - [ ] Test Commander on different laptops for demo reliability
    - [ ] Resolve divide by zero error in frame processing
    - [ ] Verify YOLO basic model bounding box drawing functionality
    - [ ] Test video streaming latency with PySide vs TK GUI
- **Kai**
    - [ ] Implement homing sequence in separate feature branch
    - [ ] Calculate encoder rotations and position tracking math
    - [ ] Document limit switch behavior and directional homing approach
    - [ ] Resolve remaining circular import issues with model options
    - [ ] Add feature documentation for gesture control recommendation
- **Connor**
    - [ ] Fix failing test assertions after app settings merge
    - [ ] Review and close irrelevant GitHub issues related to limit switches
    - [ ] Add integration tests for UI components
    - [ ] Implement JUnit XML reporting in GitHub Actions
    - [ ] Forward Maker Fair exhibitor email to team members
- **Team**
    - [ ] Run Commander demo on Thursday with robot control
    - [ ] Maintain hardware freeze after March 31st - no modifications allowed
    - [ ] Complete all features before code freeze on April 10th
    - [ ] Prepare for Maker Fair presentation with booth setup
    - [ ] Consider migrating from Python to C++ backend for performance improvements
