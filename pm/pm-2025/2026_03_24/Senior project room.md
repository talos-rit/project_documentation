Zoom: https://rit.zoom.us/rec/share/wXXP8okkG2kyARvT5IusdLn55UmnRQeXXDibNz2nP3TUuNMMzcgUnxFiPwDlNswj.Qh9F_PMRM7rhl74m

## Key takeaways

- Team discussed motor driver specifications and PCB design considerations for the robot control system
- Current motor drivers are under-specced; identified need for drivers capable of handling 6 amps at 12 volts
- Explored using dual ESP32 configuration to address GPIO pin limitations and pulse counter requirements
- Identified critical safety concerns: need for fuses and improved flyback voltage protection to prevent component damage
- Discussed switching from Raspberry Pi to ESP32 due to interrupt handling limitations
- Reviewed project rubrics and documentation requirements for milestone evaluation
- Meeting with sponsor rescheduled from original date to Thursday

## Team updates


| **Speaker** |                                         **Completed**                                          |                            **In Progress**                            |                                   **Plan**                                   |                       **Blockers**                        |
|-------------|------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------|------------------------------------------------------------------------------|-----------------------------------------------------------|
| **Hiroto**  | - Motor control code for motors 1 and 2- Debugger launch configurations updated                | - Testing pulse counter functionality- Current sensing implementation | - Refine motor control logic- Address encoder integration                    | - CUDA out of memory errors on research cluster           |
| **Connor**  | - Investigated Raspberry Pi interrupt limitations- Researched ESP32 pulse counter capabilities | - Working on connection management code- App settings implementation  | - Complete API integration- Test multi-robot communication                   | - Raspberry Pi cannot handle required interrupt frequency |
| **Ryan**    | - Rubric documentation review                                                                  | - Filling out project rubrics- Documentation updates                  | - Complete milestone documentation                                           | N/A                                                       |
| **Mike**    | - Provided PCB design consultation                                                             | N/A                                                                   | - Review schematic improvements- Provide design best practices documentation | N/A                                                       |

## High risks

- Motor drivers currently under-specced and prone to failure; multiple components already damaged due to voltage spikes
- No proper fusing system in place, leading to catastrophic failures when shorts occur
- Flyback voltage protection inadequate; diodes rated at 20V insufficient for 12V system with inductive kickback
- Current sensing implementation incomplete, preventing human-safe operation requirements
- Limited GPIO pins on single ESP32 may require dual-processor architecture
- Lack of consistent unit testing framework, particularly for hardware interfaces
- Documentation accessibility issues; information exists but not easily accessible to stakeholders

## Action items

- **Hiroto**
    - [ ] Merge pulse counter code after conflict resolution
    - [ ] Test current sensing with shunt resistors
    - [ ] Investigate VLLM memory issues on research cluster
    - [ ] Update motor control implementation for dual ESP32 configuration
- **Connor**
    - [ ] Complete connection management refactoring
    - [ ] Fix app settings fetch implementation
    - [ ] Test multi-robot communication functionality
    - [ ] Implement robot ID handling for add vs edit operations
- **Ryan**
    - [ ] Complete project rubric documentation
    - [ ] Add comments about DeepWiki incorporation for documentation accessibility
    - [ ] Document lessons learned regarding scope changes and milestone management
    - [ ] Update metrics documentation with kill count tracking
- **Team**
    - [ ] Research and select appropriate motor drivers capable of 6+ amps at 12V
    - [ ] Design fusing system for main power and individual components
    - [ ] Implement proper flyback voltage protection with appropriate Zener diodes
    - [ ] Evaluate dual ESP32 architecture for GPIO expansion
    - [ ] Add current sensing circuitry for safety compliance
    - [ ] Review and improve PCB schematic readability and labeling
    - [ ] Investigate through-hole vs surface-mount component trade-offs for hobbyist accessibility
    - [ ] Create comprehensive testing strategy for hardware components
    - [ ] Prepare for Thursday sponsor meeting with updated demo materials
- **Mike**
    - [ ] Send PCB design best practices PowerPoint notes
    - [ ] Provide feedback on schematic improvements for readability
- **Future considerations**
    - [ ] Explore dedicated DSP chips for high-frequency encoder sampling
    - [ ] Consider PCB design with both hobbyist and production versions
    - [ ] Implement verification and validation testing procedures
    - [ ] Address deployment strategy for Commander system
