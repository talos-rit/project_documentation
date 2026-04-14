Zoom: https://rit.zoom.us/rec/share/nWGpes2xcDaEYxyvnJsDuA-gxrXx9-vk5GeVCn-Aei4I2W7kpEq1G9BZs4xETTiZ.k2ix6H9q_424A_Zr

## Key takeaways

- Mid-term evaluation documentation has been completed and submitted
- Maker Faire application submitted; awaiting confirmation email (event scheduled for November 15th, 9 AM - 5 PM at Gordon Field House)
- Milestone set for November 14th to have robots operational with manual control before Maker Faire
- Solution found for encoder voltage issue: using 4 volts instead of 5 volts to prevent diode burnout, eliminating need for resistor replacement
- Driver codebase being refactored from C-style to proper C++ (vectors, classes, removing custom linked lists)
- Multiple webcam support and operator connection improvements in progress
- Command naming convention updated: "command_val" renamed to "command_id" and "command_id" renamed to "message_id" for clarity

## Team updates

| **Speaker** |                                                                            **Completed**                                                                             |                                                                                                               **In Progress**                                                                                                               |                                                                                               **Plan**                                                                                               |                                                                                         **Blockers**                                                                                          |
|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Connor, Ryan**  | - Code cleanup and C++ modernization - Mid-term evaluation document submitted - Driver README created - Updated API naming conventions (command_val to command_id, command_id to message_id)  | - Speed control implementation (set speed works, get speed doesn't) - Refactoring operator code for command responses- Updating ICD documentation- Converting ACL code to use enum classes | - Template response structure for operator commands- Complete hardware implementation before Maker Faire- Refactor ICR driver to use proper C++ features- Custom AC commands to set motor velocities | - Manual interface locks out other controls during operation- Socket communication in different thread than serial |
| **Hiroto, Kai**  | - Multiple operator connections refactored - Fixed memory leak bug - Fixed scheduler subthread issue - Refactored asset files into one directory - Two connections completed | - Multiple webcam support implementation- Fixing Director compatibility with new connection tracking | - Complete multi-camera operator support- Refactor Director out of manual interface- Install dependencies and test new build  | - Director broken due to connection refactoring changes- Need to complete before Maker Faire deadline |
| **J**     | - Created schematics for hardware - Discovered other encoders are functional | - Hardware implementation tasks- Testing encoder with 4-volt solution - Working on 4-volt power supply solution for encoders- Implementing MCP limit switches and PCA motor controller driver | - Complete hardware setup for Maker Faire- Finish remaining hardware tasks | - Need voltage regulator circuit (shed only has 5V, 3.3V, and 15V supplies) |

## High risks

- Operator response architecture may require significant refactoring due to synchronous serial communication and manual interface limitations
- Director functionality currently broken after multi-connection refactoring; needs fixing before Maker Faire
- Voltage regulator circuit needed for encoders; custom solution required as standard supplies unavailable
- Two-week timeline to Maker Faire is tight for remaining implementation tasks
- Current sensing implementation blocked until other components completed (gripper disabled for safety)

## Action items

- **Team**
    - [ ] Schedule demo with Malikowsky around November 14th
    - [ ] Ensure all robots operational with manual control by November 14th
    - [ ] Bring router for Maker Faire demo setup
    - [ ] Follow up on Maker Faire confirmation email if not received by Friday
    - [ ] Test complete system before Maker Faire
- **J**
    - [ ] Create voltage regulator circuit for 4-volt encoder power supply
    - [ ] Implement MCP limit switches driver
    - [ ] Implement PCA motor controller driver
    - [ ] Test 4-volt encoder solution thoroughly
    - [ ] Verify all motor functionality except gripper
    - [ ] Research custom AC commands for motor velocity control
- **Connor, Ryan**
    - [ ] Review and merge set speed functionality
    - [ ] Complete operator command response template structure
    - [ ] Fix PM folder organization issue
    - [ ] Update ICD documentation with new naming conventions
    - [ ] Convert remaining ACL code to enum classes
    - [ ] Remove S-list linked list implementation
    - [ ] Complete individual evaluations survey
    - [ ] Complete remaining hardware implementation tasks (Deadline: before November 14th)
    - [ ] Implement CRC checking for command validation
- **Hiro, Kai**
    - [ ] Test and merge multi-connection branch
    - [ ] Refactor Director to work with new connection tracking (Deadline: before November 14th)
    - [ ] Complete multiple webcam support implementation (Deadline: before November 14th)
