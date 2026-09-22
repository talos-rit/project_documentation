# Robotic Autonomous Cameraperson Refinement Project Plan

## Project Goals & Scope
This research project repurposes legacy robotic hardware, the ScorBot ER-4pc and ER-V, into an autonomous camera platform. The project's goal is to automate camera operations which typically require a person: panning, tilting, and framing subjects. By utilizing two robotic arms, two camera feeds can be integrated and combined, autonomously and dynamically deciding which feed to highlight. This project includes ideating, research, implementation, and refinement of software that combines machine learning, computer vision, and other technologies to guide camera movements. The previous two iterations of this project focused on subject tracking and a new controller for the ER-4pc. The plan for this year includes refining subject tracking and implementing predictive movements, a digital twin to use for testing and debugging, and an updated controller user interface for more intuitive design. This year’s main goal is multi-camera integration.
### High Level Domain Model
![](https://github.com/talos-rit/project_documentation/blob/master/artifacts/2026/Domain%20Model.drawio.png)
## Planned Milestones and Major Deliverables

| Milestone / Deliverable | Description                                                                                       | Target Date |
| ----------------------- | ------------------------------------------------------------------------------------------------- | ----------- |
| Bluey Hardware          | Finalize Bluey (ER4pc) driver and make it functional                                              | 11/14       |
| Computer Vision         | System should be successfully tracking subjects and moving the Scorbot arms accordingly to follow | 11/14       |
| Maker Faire             | Showcase our work so far -- can use as a dry run for Imagine                                      | 11/21       |
| Commander UI            | Commander UI should be more intuitive and display more helpful indicators & metrics               | 11/30       |
| Multi-Cam               | The two video feeds should dynamically switch between each other on Commander                     | 11/30       |
| Imagine RIT             | Final showcase                                                                                    | 04/21       |

We plan to have the project in a working state with the initial scope of basic subject detection and following completed by the Maker Faire. We set deadlines for deliverables a week before the Maker Faire so we have buffer time to prepare.\
We currently only have major milestones and deliverables planned through the fall semester, but plan to re-evaluate our progress and scope towards the end of the semester to plan for the spring.
## Initial Requirements & User Stories
### Functional Requirements
1. Both robots must be able to move and record
2. Both robots must be able to identify the lecturer and other subjects
3. Both robots must be able to follow the lecturer
4. Both robots must be able to be accessed by a user friendly UI
5. A lecture must be able to be outputted, which is a combination of video streams from the cameras on both robots
### Nonfunctional Requirements
1. Project must be open source
2. Project must be have an easy set up time for someone with a basic technical background
### User Stories
1. As the end user, I want to use 2 robots to capture myself in various angles while giving lectures in order to improve how my lecture is presented.
2. As the end user, I want to have a software that is given the video feeds from both robots to automatically figure out the best shot from both of them and output an edited lecture video, so that I can focus on my lecture as opposed to editing videos myself. 
3. As the end user, I want to walk around my classroom and have the camera follow me so that I can focus on my lecture and not have to move the camera at the same time.
4. As the end user, I want to easily access the Robot UIs from my device so that I can manage my videos and cameras easily. 
5. As the end user, I want to save the videos captured by the robots easily so that I can use them in the future.
## Communication & Stakeholder Management Plan
Each Thursday, we meet with our sponsor/coach to give status updates and receive feedback on our progress and future plans and goals for the project. We also set aside time to bring up any major risks we have encountered, both physical robot-related risks and project risks, so that we can discuss potential mitigations and solutions.\
As a part of these weekly meetings, we go over our weekly 4-up, which details individual progress, current risks as of that week, immediate next plans, and any needs of the team from the sponsor/coach.
## Risk Management Plan
- Implement overcurrent protection on Bluey
- Be careful with the webcams since the robots can break them