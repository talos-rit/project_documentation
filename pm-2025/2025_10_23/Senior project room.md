Zoom: https://rit.zoom.us/rec/share/LNXuGsZIfWmvifQoH-8Al4AoNMkKBpPRjHJ3WBh1X9SX2uTfX6gThKGbMzgh7gpL.iDffCnOXIfIVyLCJ

## Key takeaways

- Encoders are working with 4 discrete intervals for direction and step detection; UV blocking issue resolved by painting resin prints black
- New agile methodology adopted: feature-driven development with milestone-based releases, expected completion dates on issues, and weekly issue reviews
- Maker Faire demo scheduled for November 15th as first milestone deadline
- Hardware goal for Maker Faire: manual motor control with limit switches and e-stop (encoders and current sensing excluded from this milestone)
- Commander goal: support multiple robot connections with ability to swap between them for manual control
- Bi-weekly email updates to sponsor identified as improvement area
- Logic level shifters needed for 5V to 3.3V I2C communication (Hiroto has spares available)
## Team updates

| **Speaker**      |                                                                                                              **Completed**                                                                                                              |                                        **In Progress**                                         |                                                                                                                **Plan**                                                                                                                 |                                                                  **Blockers**                                                                   |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| **J**            | - Encoders working with 4 discrete intervals- Solved UV blocking issue by painting resin prints black- Researched GPIO expander signal routing- Created schematic and uploaded to Discord/GitHub- Breadboarded current sensing hardware | - Testing current sensing implementation- Soldering I2C address jumpers on motor driver boards | - Implement test programs for motors and encoder feedback- Disassemble encoders to replace 50 ohm resistors with 200-1000 ohm resistors- Wire limit switches and e-stop- Focus on manual motor control without encoders for Maker Faire | - Logic level shifter needed for 5V to 3.3V I2C conversion (resolved- Hiroto has spares)                                                        |
| **Ryan, Connor** | - Discovered manual mode blocks serial line and limits commands- Identified 500ms timeout limitation in operator- Fixed memory leaks and toggle dropdown menu issues                                                                    | - Researching ACML functions for velocity control- Exploring alternative motor control methods | - Prioritize command queue during polar pan operations- Implement I2C communication with hardware- Work on firmware with Hiroto- Get return feedback working                                                                            | - Manual mode prevents other commands during 500ms polar pan period- Need to determine if velocity control via ACML is feasible                 |
| **Kai, Hiroto**  | - Fixed memory leaks- Fixed dropdown menu responsiveness issues- Worked on config for multiple connection management                                                                                                                    | - Developing UI for managing multiple connections- Adding connection validation                | - Connect UI to actual connection logic- Add error messaging for failed connections- Support two robot connections with swap capability- Display video feeds from multiple robots- Implement variable FPS based on performance metrics  | - High CPU usage causing UI freezing (frame update rate hard-coded, not using config constant)- Need to add performance metrics and measurement |

## High risks

- Current sensing implementation is untested and represents significant additional work (excluded from Maker Faire milestone)
- Encoder implementation requires disassembly of all 6 encoders to replace resistors, with risk of damaging diodes if pulsed incorrectly
- Commander UI freezing issues due to overloaded event loop and high CPU usage on some machines
- Polar pan manual mode blocks other serial commands for up to 500ms, requiring command prioritization solution
- E-stop hardware not yet functional (button doesn't press properly)
- Tight timeline for Maker Faire demo (less than one month)
## Action items

- **J**
    - [ ] Bring logic level shifters tomorrow before 9AM class
    - [ ] Solder I2C address jumpers on motor driver boards
    - [ ] Disassemble and modify encoders (replace 50 ohm resistors with 200-1000 ohm)
    - [ ] Wire limit switches and grounds
    - [ ] Implement e-stop hardware solution
    - [ ] Create test programs for motor control
    - [ ] Create issues for Maker Faire milestone by Tuesday
    - [ ] Bring breadboarded hardware tomorrow
- **Ryan, Connor**
    - [ ] Work on I2C communication implementation with Hiroto
    - [ ] Research and test ACML velocity control functions
    - [ ] Implement command prioritization for polar pan operations
    - [ ] Stub out I2C communication code
    - [ ] Create issues for Maker Faire milestone by Tuesday
- **Kai, Hiroto**
    - [ ] Complete multiple connection UI implementation
    - [ ] Add connection validation and error messaging
    - [ ] Fix hard-coded frame update constant to use config file
    - [ ] Implement performance metrics tracking
    - [ ] Add variable FPS based on execution time
    - [ ] Enable video feed display for multiple robots
    - [ ] Create issues for Maker Faire milestone by Tuesday
- **Team**
    - [ ] Apply for Maker Faire by November 1st
    - [ ] Send email to Malachowski about Maker Faire participation and new methodology (after Thursday next week)
    - [ ] Update methodology documentation in project files
    - [ ] Update issue templates in all repos with difficulty level and expected completion date fields
    - [ ] Create milestone board/view in GitHub with labels
    - [ ] Review open issues weekly starting next Thursday
    - [ ] Create poster for Maker Faire (review previous team's poster first)
    - [ ] Implement bi-weekly email updates to sponsor going forward
    - [ ] Complete retrospective rubric and comments
    - [ ] Demo MVP of UNK robot at Maker Faire as interactive component
    - [ ] Achieve manual control of Bluey robot with limit switches and e-stop by November 15th
