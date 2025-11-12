Zoom: [https://rit.zoom.us/rec/share/GJvcwqoeLWz8cMX_5DlGCGPf-_5DtqeNySISLLEVwJal6tHa-k7aw8NjA0Orb2o.lwF0mt21vrIOGJpb](https://rit.zoom.us/rec/share/GJvcwqoeLWz8cMX_5DlGCGPf-_5DtqeNySISLLEVwJal6tHa-k7aw8NjA0Orb2o.lwF0mt21vrIOGJpb)
## Key takeaways

- The team discussed implementing gRPC as a better communication method than their current approach, noting it would be easier to extend and maintain
- They identified an issue with manual mode preventing other commands from running simultaneously
- The team discovered they can run macros while in manual mode, which could solve their control issues
- They explored how to access and manipulate encoder values directly
- There was discussion about implementing a software interrupt for the e-stop functionality
- The team worked on debugging issues with the detection process and camera functionality
## Team updates

| **Speaker**  | **Completed** |                                   **In Progress**                                   |                  **Plan**                   |                          **Blockers**                           |
|--------------|---------------|-------------------------------------------------------------------------------------|---------------------------------------------|-----------------------------------------------------------------|
| Kai & Hiroto | N/A           | - Working on camera and detection process debugging                                 | N/A                                         | - Unexplained program crashes when setting object model to none |
| Hiroto       | N/A           | - Investigating gRPC implementation- Created a separate repository for gRPC testing | - Continue refining gRPC implementation     | N/A                                                             |
| Conor        | N/A           | - Working on unit testing for Operator- Reviewing code changes                      | N/A                                         | N/A                                                             |
| J            | N/A           | - Working on e-stop wiring- Testing robot control commands                          | - Investigate ACL scripts for robot control | N/A                                                             |

## High risks

- Manual mode prevents running other commands simultaneously, limiting robot control capabilities
- Unexplained program crashes when changing detection models
- Memory leaks occurring during detection process debugging
- Potential issues with coordinate system when mixing manual mode and macro operations
## Action items

- **Hiroto**
    - [ ] Continue developing gRPC implementation
    - [ ] Share gRPC repository with team members for review
- **Ryan & Connor**
    - [ ] Investigate ACL scripts for robot control
    - [ ] Create a "serial loader" for easier macro management
    - [ ] Test moving to encoder values directly
- **Hiroto & Kai**
    - [ ] Fix detection process termination issues
    - [ ] Resolve camera functionality problems
- **Team**
    - [ ] Implement software interrupt for e-stop functionality
    - [ ] Evaluate gRPC performance compared to current implementation
    - [ ] Determine approach for handling robot movement while maintaining ability to run other commands
    - [ ] Install KiCAD to review schematics
