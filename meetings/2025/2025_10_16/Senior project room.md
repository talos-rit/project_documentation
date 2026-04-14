Zoom: https://rit.zoom.us/rec/share/FtbgyvSpK_uZIfki2bGC8O50Vdf_o5UaMWosW2Y6Fu94fRG8KEpYunx9vv-rbtih.lWSeMSqW9ew5EddD
## Key takeaways

- The Adafruit order has arrived and delivered to the team room
- Operator code has been formatted to a standard, improving readability
- Speed command implementation is in progress for operator but blocked for commander
- Commander updates include camera feed through WiFi, UI enhancements with dropdown menu, and config improvements
- Current risks include termination issues in Commander and hardware development timeline
- The team needs to update the ICD (Interface Control Document) to accurately reflect current functionality
- A midterm retrospective meeting is scheduled for Monday, October 20th from 4:15-5:00 PM
## Team updates

| **Speaker**  |                                                                              **Completed**                                                                               |                             **In Progress**                              |                                                                        **Plan**                                                                        |                                   **Blockers**                                   |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------|
| Ryan, Connor | - Formatted operator code to standard- Created speed command                                                                                                             | - Working on data return flow from operator                              | - Update ICD to reflect actual functionality- Implement data return flow from operator- Implement PolarPan in Driver                                   | - Hardware availability- Complex threading model in operator                     |
| Hiro, Kai    | - Implemented camera feed through WiFi- Fixed README to remove ActiveMQ- Created UI enhancements with dropdown menu- Added config enhancement with local config override | - Working on termination issues- Created repository for camera streaming | - Document Commander architecture- Implement optional dependencies enhancement- Reorganize project directory- Multi-camera and operator implementation | - Tkinter termination issues- Coupling issues with config system                 |
| J            | - Created basic schematic for current sensing- Installed optical encoder                                                                                                 | - Working on hardware implementation                                     | - Complete hardware assembly- Test optical encoder                                                                                                     | - Understanding high side vs low side measurement- Need ferrules and heat shrink |

## High risks

- Termination issues in Commander - Python automatically detects and cleans up resources but doesn't do it on the first termination command
- Hardware development timeline and scope
- Current sensing implementation (high side vs low side understanding)
- ICD documentation inaccuracies - doesn't match actual implementation
- Complex threading model in operator making return flow implementation difficult
## Action items

- **Ryan, Connor**
    - [ ] Update ICD to accurately reflect current functionality
    - [ ] Implement data return flow from operator
    - [ ] Implement speed command with individual axis support
- **Hiro, Kai**
    - [ ] Create architecture diagram for Commander showing multiplicity handling
    - [ ] Fix dropdown menu functionality
    - [ ] Implement optional dependencies enhancement
    - [ ] Reorganize project directory structure
- **J**
    - [ ] Obtain ferrules and heat shrink for connections
    - [ ] Test optical encoder with oscilloscope
    - [ ] Complete hardware assembly in the gutted box
- **Team**
    - [ ] Attend midterm retrospective meeting (Monday, October 20th, 4:15-5:00 PM)
    - [ ] Request future deliveries be sent to the SE office for easier access
