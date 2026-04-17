Zoom: https://rit.zoom.us/rec/share/aOsiU8Aim8w8-GPN5AhXAPl-1K89SDtkjfHAeDwbJN2kq1SxYk-fz5mjlZ9becMn.GXAhnfk1hQFHDMkb
## Key takeaways

- Code freeze scheduled for tomorrow to allow final integration testing and documentation focus
- Homing functionality works but current sensing implementation incomplete; only functional from correct side of limit switch
- Poster design reviewed with minor color scheme and text adjustments needed
- Video recording planned for Sunday at noon for IAB presentation
- Technical report due April 30th; team photo needed for Imagine event
- Industrial Advisory Board poster session confirmed for April 23rd evening
- Future work documentation being compiled with focus on onboarding new developers rather than simple handoff
- Windows compatibility issues identified with Commander; workaround using TK interface available
- Discussion of converting future work items into GitHub issues with "good first issue" tags

## Team updates


|    **Speaker**     |                                                                 **Completed**                                                                  |                                                **In Progress**                                                 |                                                                         **Plan**                                                                          |                                                                **Blockers**                                                                 |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| **Hiroto**         | - Homing functionality implemented- Object model class merged- Fixed left-right directional mapping issue                                      | - Current sensing for homing (not 100% working)- Testing integration                                           | - Code freeze tomorrow- Full integration tests tonight                                                                                                    | - Current sensing spike detection needs refinement for hard stop detection                                                                  |
| **Team (General)** | - Commander control working- Demo showcased initial movement- Architecture documentation for operator completed- Poster design draft completed | - Windows compatibility fixes in review- Onboarding documentation for new developers- Poster final adjustments | - Merge pending PRs before code freeze- Film short video Sunday noon- Create GitHub issues from future work items- Finalize poster with color/text tweaks | - PyQt6 installation issues on Windows without Visual Studio/MSVC- Need to resolve environment variable settings for better error reporting |
| **Connor**         | - Research computing exploration for server deployment                                                                                         | - Cloud project dealing with teammate issues- Thesis work                                                      | - Continue documentation work- Address Workday onboarding issues                                                                                          | - Workday account access problems preventing onboarding completion- Teammate merged large uncommented PR (19k+ lines) without proper review |

## High risks

- Current sensing implementation incomplete; homing only works from one side of limit switch
- Windows compatibility issues with Commander requiring Visual Studio/MSVC or fallback to TK interface
- Large unreviewed code merge in cloud project creating technical debt and integration challenges
- Potential latency issues if migrating to JavaScript framework for Commander interface

## Action items

- **Hiroto**
    - [ ] Complete integration tests tonight before code freeze
    - [ ] Review and merge GUI button fix PR
    - [ ] Test Windows compatibility fixes
- **Team**
    - [ ] Code freeze tomorrow morning/night
    - [ ] Film demonstration video Sunday at noon
    - [ ] Add GitHub link to poster
    - [ ] Update poster: change "Scorebot ERV" background from gray to orange
    - [ ] Remove periods from poster bullet points
    - [ ] Update system architecture diagram to reflect current operator/driver structure
    - [ ] Link operator architecture document in README
    - [ ] Create GitHub issues from future work document with "good first issue" tags
    - [ ] Take team photo at Imagine event or during video filming
    - [ ] Submit technical report by April 30th
    - [ ] Reorganize documentation structure: move future work to contributions folder
    - [ ] Add note about ESP web tool setup requiring Wi-Fi configuration changes
    - [ ] Document Windows setup issues and TK interface workaround
    - [ ] Consider RSVP for IAB roundtable lunch April 24th 12-1pm
    - [ ] Review and refine future work items for next team recommendations
- **Connor**
    - [ ] Resolve Workday account access issues with employer
    - [ ] Complete drug test for new job
    - [ ] Address cloud project PR review and architectural concerns
- **Documentation Team**
    - [ ] Continue onboarding documentation writing
    - [ ] Add Jay's photo to hardware handoff documentation
    - [ ] Create separate document for Malachowski outlining recommended project direction
    - [ ] Document digital twin integration challenges and architectural considerations
    - [ ] Add table of contents to relevant documentation files using proper heading structure
