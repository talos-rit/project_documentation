Zoom: https://rit.zoom.us/rec/share/H8tCh2t-8B7GkfTc8sczOS4Afa_ekNhusTNhkUFAB4v9VecaVXF3axNzvO4s8cJ8.WcSGq0eOdKVl68bp
## Key takeaways

- ADS1015 and pulse counter drivers have been merged
- Hardware is mostly complete, but there was thermal runaway with a motor driver
- Limit switches currently need to be in parallel due to PIN count constraints
- Commander updates by QT have been merged
- Pi VCAM has been merged, allowing direct video streaming
- App settings have been merged
- The team smoked an ADC during testing
- Milestone 2 is ending tomorrow (Friday)
- The team needs to plan a demo and schedule Milestone 3 planning

## Team updates


| **Speaker** |                                                                       **Completed**                                                                        |                             **In Progress**                              |                                                  **Plan**                                                   |                                    **Blockers**                                     |
|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------|
| **Hiroto**  | - ADS1015 merged- PCNT (pulse counter) drivers in review- Commander updates by QT merged- Pi VCAM merged- App settings merged- PiSci switch bug fix merged | - Main controller application on ESP- Refactoring object model class     | - Send out email for project demo- Schedule Milestone 3 planning for Tuesday in two weeks                   | - Time constraints                                                                  |
| **Team**    | - Hardware mostly complete- Limit switches approved- One automated test implemented- DeepWiki has indexed project documentation                            | - Glue code for ESP- Homing functionality- Dynamic scheduling for frames | - Get two axes working on Bluey- Implement homing for two axes- Create standalone backend detection process | - Need new motor drivers- DB6 cable (wrong direction)- Slide-based hardware on hold |

## High risks

- ESP feature completeness by end of Milestone 2 (tomorrow)
- Limited time for Milestone 3 work (approximately 4 weeks)
- Motor drivers failing (one burned out during testing)
- ADC components smoking/shorting during testing

## Action items

- **Hiroto**
    - [ ] Send out email with when-to-meet for project demo
    - [ ] Find space for demo
    - [ ] Contact Malachowski about getting Manili and others to attend demo
- **Team**
    - [ ] Complete team surveys
    - [ ] Prepare for Milestone 3 planning on Tuesday in two weeks
    - [ ] Complete retrospective document when returning next week
    - [ ] Research and purchase new motor drivers (possibly DRV8871)
    - [ ] Source correct DB6 cable or gender adapter
    - [ ] Debug streaming issues
    - [ ] Update documentation
- **Jay**
    - [ ] Fix CLI arguments
- **Ryan**
    - [ ] Update app settings manager
