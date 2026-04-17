Zoom: https://rit.zoom.us/rec/share/wY8sWPnZKTQOt0_O-kLjQt-ZSnPXtsZnidp_qxQpHhnnJY0O1jlPfQKiwkGyDG04.tAEFCKs72oRdvjkV
## Key takeaways

- Midterm evaluation document reviewed and will serve as the final evaluation
- Demo meeting scheduled for next Thursday during senior project time
- PCNT driver merge is blocking multiple tasks including robot controller development
- Team discussed transitioning to PCB-based design for future iterations to simplify hardware assembly
- Documentation efforts prioritized for operator queue system and ESP driver
- Commander maturity has reached point where unit testing is becoming more feasible
- Poster design in progress for Imagine RIT event
- Hardware limitations identified with current motor hats exceeding 6 amp specifications

## Team updates


| **Speaker** |                                           **Completed**                                            |                                                  **In Progress**                                                   |                                                                 **Plan**                                                                  |                                     **Blockers**                                      |
|-------------|----------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| **Hiroto**  | - Fixed CI for PCNT driver- Completed robot controller rough draft- Shorted system issues resolved | - PCNT driver review (waiting on Jay)- Documentation for operator asynchronous queue system- README for ESP driver | - Merge PCNT driver once Jay completes comments- Create installation guide and links to documentation- Add rationale for hardware choices | - PCNT driver blocked by Jay's illness- Robot controller blocked by PCNT driver merge |
| **Ryan**    | - Testing stub for Commander merged- Bug fixes merged- Poster stub created                         | - Working on poster design- Object detection improvements                                                          | - Review poster pull request- Add diagrams to poster- Create screenshots of Commander interface                                           | N/A                                                                                   |
| **Connor**  | - Worked on poster layout- Meeting with Mike from CET about PCB design                             | - Poster documentation- Creating diagrams for data flow                                                            | - Add more diagrams to poster- Reduce text overlap- Create block diagrams for Commander                                                   | N/A                                                                                   |
| **Jay**     | N/A                                                                                                | - Script for demo- PCNT driver (2 open comments)                                                                   | - Complete PCNT driver review                                                                                                             | - Sick and unable to attend meeting- PCNT driver completion delayed                   |

## High risks

- PCNT driver merge delayed due to Jay's illness, blocking robot controller and encoder integration
- Current motor hats exceed 6 amp specifications, potentially losing motor functionality
- Two-axis system may not be completely working by demo
- Limited time remaining before demo and Imagine RIT
- Hardware testing infrastructure needs improvement
- Multiple ADS boards potentially broken, only one confirmed working

## Action items

- **Jay**
    - [ ] Complete PCNT driver review and resolve 2 open comments
    - [ ] Finish demo script
- **Hiroto**
    - [ ] Merge PCNT driver once Jay completes review
    - [ ] Document operator asynchronous queue system explaining 3-queue architecture
    - [ ] Create README for ESP driver with installation guide and documentation links
    - [ ] Add links to KICAD schematic for BOM generation
    - [ ] Write rationale document for hardware choices
- **Ryan**
    - [ ] Review poster pull request
    - [ ] Add bounding box demonstration to Commander screenshots
    - [ ] Help with ESP driver README and documentation
- **Connor**
    - [ ] Add diagrams to poster (data flow, block diagrams)
    - [ ] Reduce text and use phrases instead of full sentences on poster
    - [ ] Create Commander interface screenshots
    - [ ] Add block diagram showing major components
- **Team**
    - [ ] Write up meeting notes from hardware discussion with Mike about PCB design
    - [ ] Document future PCB design considerations for next team
    - [ ] Adopt open hardware license in addition to open source license
    - [ ] Request poster printing for Imagine RIT event
    - [ ] Set up classroom for Imagine RIT (open shades, prop door, position poster)
    - [ ] Prepare demo for next Thursday meeting with Malachowski and department head
    - [ ] Test ADS board functionality before demo
    - [ ] Discuss CE department member recruitment with Malachowski at demo meeting
    - [ ] Implement metrics tracking (PR count, issue closures, major changes, time estimates)
    - [ ] Consider adding strum.pack for serialization improvements
