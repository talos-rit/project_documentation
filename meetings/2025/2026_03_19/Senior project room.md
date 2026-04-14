Zoom: https://rit.zoom.us/rec/share/CqZ5aaIi0F2L8W4xpxiEot-0pCjOIcGL0BpRbQ4RQxoCrnnQXwMCexWAyNc3bU7x.ay_OFmQU2Q2o7ndW

## Key takeaways

- Hardware failure occurred before spring break: 12-volt power rail shorted to metal enclosure, destroying 3 current sensor modules and a microcontroller
- Team is under feature freeze; focusing on documentation and polishing existing features rather than adding new functionality
- Code freeze scheduled for April 10th (excluding hot patches and documentation updates)
- Demo meeting scheduled for March 26th, Thursday at 5pm in senior project room
- Only two motor axes can be demonstrated due to hardware limitations; team will focus on documentation for future teams
- Milestone 3 end date set at Imagine RIT
- Government shutdown affecting FEMA funding has frozen $625 million pot for World Cup host cities, including Massachusetts' $46 million application

## Team updates


|    **Speaker**    |                                                                                 **Completed**                                                                                 |                                                **In Progress**                                                |                                                             **Plan**                                                             |                                                    **Blockers**                                                    |
|-------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------|
| **Hiroto**        | - PCNT review finishing up- Renamed Glue Code to Robot Control- Created Milestone 2 release for Commander- Set up DeepWiki for ESP driver- Sent demo meeting scheduling email | - Writing commits for review comments- Application settings auto-reloading feature- More Commander unit tests | - Schedule demo meeting with Malachowski- Finalize ESP driver documentation- Complete robot control implementation               | - No common meeting times available for all stakeholders- Hardware destroyed (3 sensor modules, 1 microcontroller) |
| **Team Member 2** | - Arguments feature completed- Optional boundary boxes drawing implemented- Set up SonarCube for static analysis on Commander                                                 | - Robot control/glue code development- Commander unit tests (mostly auto-generated)- Budget completed         | - Work on robot control this weekend- Implement socket commands- Add encoder, limit switch, and overcurrent e-stop functionality | - Waiting for hardware setup                                                                                       |
| **Team Member 3** | - Started operator documentation                                                                                                                                              | - Documentation for operators- Project documentation work                                                     | - Focus entirely on documentation for next few weeks- Complete onboarding documentation for ESP driver                           | - N/A                                                                                                              |

## High risks

- Insufficient hardware components for full demonstration (only 2 motor drivers and 2 current sensor modules available)
- Limited time to complete comprehensive documentation before semester end
- Potential inability to demonstrate full robot functionality at demo meeting
- Feature freeze may limit ability to address newly discovered issues
- Team may need to request additional funding for replacement hardware

## Action items

- **Hiroto**
    - [ ] Find Malachowski to discuss demo meeting availability
    - [ ] Send calendar invite for March 26th demo meeting at 5pm
    - [ ] Include Manili, Larry Kaiser, and department head in meeting invite
    - [ ] Create issues for Milestone 3 planning items
    - [ ] Add README documentation for ESP driver
    - [ ] Document hardware selection rationale
    - [ ] Write up ideas for new motor drivers
- **Team Member 2**
    - [ ] Complete robot control/glue code implementation this weekend
    - [ ] Implement socket commands (start, stop, speed control)
    - [ ] Add encoder integration
    - [ ] Implement limit switch functionality
    - [ ] Add overcurrent e-stop feature
    - [ ] Create PR for current unit tests
    - [ ] Generate additional unit tests for untested modules
- **Team Member 3**
    - [ ] Continue documentation work for operators
    - [ ] Create onboarding documentation for ESP driver
    - [ ] Document flashing procedures
    - [ ] Add datasheet links to KiCAD schematics
    - [ ] Export BOM from KiCAD
- **All Team Members**
    - [ ] Complete midterm review form by Tuesday meeting
    - [ ] Review midterm review on Thursday meeting
    - [ ] Request last year's technical report for reference
    - [ ] Create poster for presentation (avoid text-heavy design, use diagrams)
    - [ ] Prepare video demo as backup for presentation
    - [ ] Consider creating architecture diagram for poster
    - [ ] Finalize tech stack documentation
- **Malachowski**
    - [ ] Confirm availability for March 26th demo meeting at 5pm
    - [ ] Review demo meeting invite and respond
