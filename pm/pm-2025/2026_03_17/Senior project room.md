Zoom: https://rit.zoom.us/rec/share/GYm8eP6pyb3pbdVW7Od8-rfsfn4F7Jmh2_jzUxgsHO0Uo_Wok7kCl80VMuGZIcRk.iZH3_8F-hrbqDlwH

## Key takeaways

- Team is 2 weeks away from code freeze (April 10th), with final technical report due April 30th
- Operator project will no longer be in active development but cannot be fully deprecated yet due to pending ESP ACL implementation
- Commander needs comprehensive documentation including architecture diagrams for the queue system
- Hardware demo requires replacing ADS components; multiple units were damaged during testing
- All team members will contribute equally to the final presentation, poster, and video deliverables
- A Computer Engineering student may join next meeting to discuss potential PCB design opportunities

## Team updates


| **Speaker** |                                **Completed**                                |                      **In Progress**                       |                                                                                                              **Plan**                                                                                                               |                **Blockers**                 |
|-------------|-----------------------------------------------------------------------------|------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------|
| **Hiroto**  | N/A                                                                         | ESP driver developmentPulse counter workRobot control code | Complete glue code and pulse counterReview and finish scheduled itemsFix ADS driver bug (5-second fix moving static variable to handle)Hardware integration and testingFinish hardware for demo by replacing damaged ADS components | N/A                                         |
| **Ryan**    | Started ESP driver branch (not PR yet)Restructured Commander test framework | Commander documentationUnit testing framework setup        | Create README for CommanderWrite architecture diagram and documentation for queue systemImplement unit tests for CommanderAdd docstrings to codeLink schematic to datasheets in hardware documentation                              | Need to determine what to include in README |
| **Connor**  | N/A                                                                         | PySide6 GUI improvements                                   | Review and improve PySide6 UIAdd docstrings to codeWork on PCB schematics in extra timeCreate initial poster draftCollaborate on video production                                                                                   | N/A                                         |

## High risks

- Multiple ADS components damaged and need replacement before demo
- ESP ACL functionality not yet implemented, blocking full deprecation of Operator
- Extensive documentation gaps across multiple components (Operator, Commander, ESP driver)
- Tight timeline with only 3.5 weeks until code freeze and multiple deliverables due simultaneously

## Action items

- **Hiroto**
    - [ ] Complete ESP driver glue code and pulse counter implementation (April 10th)
    - [ ] Fix ADS driver bug - move static variable to handle
    - [ ] Replace damaged ADS components for hardware demo
    - [ ] Complete hardware integration and testing
    - [ ] Add unit tests to ESP driver
    - [ ] Create release version for ESP driver once glue code is finished
- **Ryan**
    - [ ] Create README for Commander with comprehensive documentation
    - [ ] Write architecture diagram explaining queue system for Commander
    - [ ] Document Operator project with memo on reading and README update
    - [ ] Implement unit testing for Commander
    - [ ] Add docstrings throughout codebase
    - [ ] Link schematic components to datasheets and parts in hardware documentation
    - [ ] Generate BOM from schematic
- **Connor**
    - [ ] Review and beautify PySide6 UI
    - [ ] Add docstrings to PySide6 code
    - [ ] Work on PCB schematics documentation in available time
    - [ ] Create initial draft of project poster (April 10th)
    - [ ] Collaborate on video script and production
- **All team members**
    - [ ] Contribute equally to final presentation preparation
    - [ ] Participate in video filming session (schedule TBD)
    - [ ] Complete poster collaboratively (April 10th)
    - [ ] Submit technical report (April 30th)
    - [ ] Prepare for code freeze (April 10th)
    - [ ] Meet with Computer Engineering student next meeting to discuss PCB design possibilities
- **Future considerations**
    - [ ] Implement zero MQ for object detection
    - [ ] Clean up object detection code
    - [ ] Add app settings editing functionality to PySide6
    - [ ] Fully deprecate Operator once ESP ACL is implemented
    - [ ] Create custom enclosure design
    - [ ] Research new motor drivers for next team
