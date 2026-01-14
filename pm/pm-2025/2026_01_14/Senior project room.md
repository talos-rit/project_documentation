Zoom: https://rit.zoom.us/rec/share/UFGRDuXuDdTYfnxqc2li4JOCWz-kw9dEiTGZT18gADd4T0ellVR3nIcNbPdk5l7p.dhilUMkhHqD7x1Jx
## Key takeaways

- Team decided to use MIT license for the project after discussing licensing options
- Milestone 1.5 is mostly complete with a few remaining items
- Milestone 2 will end before spring break (March 5th)
- The team needs to refactor the operator code, which will require architectural planning
- Documentation will be moved to backlog until after refactoring is complete
- The team will dedicate next week to planning and creating tickets for upcoming work
## Team updates


| **Speaker** |              **Completed**               |                   **In Progress**                   |                       **Plan**                       |    **Blockers**    |
|-------------|------------------------------------------|-----------------------------------------------------|------------------------------------------------------|--------------------|
| **Hiroto**  | - Initial unit tests                     | - Documentation for Commander                       | - Set up dev container- Work on operator refactoring | N/A                |
| **Ryan**    | - Fixed driver and made it more readable | - Working on PCA driver issue- CMake implementation | - Test PCA driver with tester script                 | - PCA driver issue |
| **Connor**  | - Worked on hardware schematics          | - Desoldering top hat                               | - Continue hardware work                             | N/A                |
| **Jay**     | - Investigated gRPC spike                | - Working on operator architecture                  | - Create document on operator architecture needs     | N/A                |

## High risks

- PCA driver issue is blocking progress on milestone 1.5
- Both object detection models (YOLO and MediaPipe) are problematic and may need to be restructured
- Operator code needs significant refactoring, potentially a complete rewrite
## Action items

- **Team**
    - [ ] Discuss open source license with Prof. Meneely
    - [ ] Plan architecture for operator refactoring next week
    - [ ] Break down refactoring tasks into tickets
    - [ ] Create a document outlining operator architecture needs
- **Ryan**
    - [ ] Test PCA driver on Raspberry Pi
    - [ ] Fix driver issues
    - [ ] Complete CMake implementation
- **Connor**
    - [ ] Complete desoldering of top hat
    - [ ] Work on hardware schematics
- **Jay**
    - [ ] Create document on operator architecture requirements
    - [ ] Investigate task queuing systems
- **Hiroto**
    - [ ] Set up dev container for operator
    - [ ] Work on refactoring operator code
