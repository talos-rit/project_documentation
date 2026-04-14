Zoom: https://rit.zoom.us/rec/share/oCBMHT9JpLa_r9uPAC2BnMROb8egfBYxUUXVsL9ztXgTuWqoeeoAX8-PTu1kfEOx.r2qv5ChctuWllyTQ

## Key takeaways

- The team is working on a robotic camera system with two models (Bluey and Commander)
- They discovered their development process resembles Feature-Driven Development (FDD)
- A motor control driver burned out at Maker Faire due to back EMF from motors
- The team is focusing on improving hardware safety and reliability for Milestone 2
- They plan to implement parallel homing to improve efficiency
- They need to add an open source license to the repository (likely MIT)
- A graduate student may contribute to the project, focusing on computer vision
## Team updates


| **Speaker** |                       **Completed**                       |                                             **In Progress**                                              |                  **Plan**                  | **Blockers** |
|-------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------|--------------------------------------------|--------------|
| **Hiroto**  | - Worked on Commander- Presented at Maker Faire           | - Working on spike for implementing different communications- Working across both commander and operator | - Complete documentation for Milestone 1.5 | N/A          |
| **Kai**     | - Worked on Commander                                     | - Learning more about operator                                                                           | - Documentation for Milestone 1.5          | N/A          |
| **Ryan**    | - Worked on operator- Worked on PyInstaller for Commander | - Working on driver                                                                                      | - Complete a homing cycle for Milestone 2  | N/A          |
| **Connor**  | - Worked on operator- Worked on PyInstaller for Commander | - Working on unit test suite for the operator                                                            | - Documentation for Milestone 1.5          | N/A          |
| **Jay**     | - Worked on hardware- Demonstrated at Maker Faire         | - Hardware improvements                                                                                  | - Documentation for Milestone 1.5          | N/A          |

## High risks

- Back EMF from motors causing hardware damage (burned out motor control driver)
- Current sensing and human safety features need implementation
- Slide-based motor drawing too much current (7 amps with 2 amp jumps)
- Limited camera capabilities affecting video quality and framing options
## Action items

- **Team**
    - [ ] Add an open source license to the repository (likely MIT)
    - [ ] Consult with Benili about open source licensing options
    - [ ] Implement back EMF protection for motors
    - [ ] Complete documentation for Milestone 1.5
    - [ ] Implement current sensing for human safety features
    - [ ] Research parallel homing implementation for improved efficiency
    - [ ] Refactor operator code to improve structure
- **Ryan**
    - [ ] Complete a homing cycle for Milestone 2
    - [ ] Investigate parallel homing possibilities
- **Connor**
    - [ ] Complete unit test suite for the operator
- **Hiroto**
    - [ ] Continue spike for implementing different communications
