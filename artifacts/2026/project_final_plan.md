# Robotic Autonomous Cameraperson Refinement Project Plan

## Project Goals & Scope
This research project repurposes legacy robotic hardware, the ScorBot ER-4pc and ER-V, into an autonomous camera platform. The project's goal is to automate camera operations which typically require a person: panning, tilting, and framing subjects. By utilizing two robotic arms, two camera feeds can be integrated and combined, autonomously and dynamically deciding which feed to highlight. This project includes ideating, research, implementation, and refinement of software that combines machine learning, computer vision, and other technologies to guide camera movements. The previous two iterations of this project focused on subject tracking and a new controller for the ER-4pc. The plan for this year includes refining subject tracking and implementing predictive movements, a digital twin to use for testing and debugging, and an updated controller user interface for more intuitive design. This year’s main goal is multi-camera integration.
### High Level Domain Model
![[Domain Model.drawio.png]]
## Planned Milestones and Major Deliverables
We plan to use Scrum, working in two week sprints. We will order our backlog based on priority and chronological order to ensure we are doing tasks sequentially as they need to be completed.\
Our design cycle is research → analysis & design → development → testing → refinement\
Assign each major artifact & project goal a deadline, determine during backlog refinement & sprint planning what needs to be prioritized each sprint.

| Milestone / Deliverable | Description                                                                                       | Target Date                                            |
| ----------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------ |
| Bluey Hardware          | Finalize Bluey (ER4pc) driver and make it functional                                              | 11/14                                                  |
| Computer Vision         | System should be successfully tracking subjects and moving the Scorbot arms accordingly to follow | 11/14                                                  |
| Maker Faire             | Showcase our work so far -- can use as a dry run for Imagine                                      | 11/21                                                  |
| Commander UI            | Commander UI should be more intuitive and display more helpful indicators & metrics               | 11/30                                                  |
| Multi-Cam               | The two video feeds should dynamically switch between each other on Commander                     | 11/30                                                  |
| Feature Freeze #1       | Feature freeze to focus on finalizing semester 1 docs                                             | 11/30                                                  |
| Feature Freeze #2       | Feature freeze to focus on finalizing final docs                                                  | 04/11                                                  |
| Stakeholder Showcases   | Showcasing the current state of the project. With Profs Kiser & Meneeley                          | TO DO: one mid-late semester ~mid-nov; 2 in the spring |
| Imagine RIT             | Final showcase                                                                                    | 04/21                                                  |

We plan to have the project in a working state with the initial scope of basic subject detection and following completed by the Maker Faire. We set deadlines for deliverables a week before the Maker Faire so we have buffer time to prepare.\
We currently only have major milestones and deliverables planned through the fall semester, but plan to re-evaluate our progress and scope towards the end of the semester to plan for the spring.

### Standards and Quality Practices
- **Testing** → Components will be tested using unit and integration tests. Manual testing will also happen continuously across all components for the whole project.
- **Deployment Verification** → CI set up to run unit tests in all repositories when code is pushed.
- **Deadlines** → Research, implementation, and testing will all be time-boxed. If work is not being completed within the timeframe, it will be reassessed to either adjust the timeframe or adjust the scope of the item itself.
- **Documentation** → Daily/weekly progress will be documented in [meeting notes](meetings/2026/fall_semester/meeting_notes), but any major developments will be documented in the proper location in this repository so future teams can easily find old research, resources, and status. Items that seem self-explanatory but would not be to someone unfamiliar with the project should be documented in an accessible and logical location.
### Tools
We plan to use GitHub projects to keep track of our project tasks: https://github.com/orgs/talos-rit/projects/9 \
GitHub projects will allow us to tag issues to specific repos & bugs as tasks to complete by individual members. We will customize our project board to have a Product Backlog, Spring Backlog, In-Progress, Ready for Review/Test, Blocked, and Done.

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
6. As an end user, I would like to be able to confirm position and basic operation. ("Does my robot work?")
## Communication & Stakeholder Management Plan
Each Thursday, we meet with our sponsor/coach to give status updates and receive feedback on our progress and future plans and goals for the project. We also set aside time to bring up any major risks we have encountered, both physical robot-related risks and project risks, so that we can discuss potential mitigations and solutions.\
As a part of these weekly meetings, we go over our weekly 4-up, which details individual progress, current risks as of that week, immediate next plans, and any needs of the team from the sponsor/coach.
## Risk Management Plan
- Implement overcurrent protection on Bluey
- Be careful with the webcams since the robots can break them

## Metrics 
These metrics will be compiled & analyzed during each sprint retrospective.

### Process Metrics
1. **Working hours** → number of hours worked each week by each member
2. **Velocity** → average amount of work completed per sprint
3. **PR confidence** →  scale 1-5 how confident you are in quality of PRs
4. **Estimation accuracy** → planned vs actual effort for each story
5. **Kill count** → number of hardware devices broken during development

### Project Metrics
1. **Earned value** → estimated percentage of total project completion
2. **Ease of Onboarding** → how easy it is to setup the robot based on docs instructions, as estimated by a team member acting as evaluator/devil's advocate

### Software Functionality Metrics
1. **Predictive Accuracy** → measurement of how accurate the computer vision predictive model is
2. **End User Experience** → have volunteers with a feasible use-case test the arm and evaluate it based on:
	- Ease of use
	- Practicality
	- Whether or not they would use it day-to-day
3. **Camera to Arm Latency** → measurement of how long it takes for camera readings to pass to the arm and for the arm to take action
4. **Time to Detect Person** → how long the model takes to recognize a person in frame

### Product Metrics
- Test Coverage → percentage of code branches/lines under test

## Roles
| Name              | Role                      |
| ----------------- | ------------------------- |
| Aleck Hernandez   | Hardware Lead             |
| Bevan Neiberg     | AI / Computer Vision Lead |
| Briggs Tucker     | Communications Lead       |
| Christine Morgado | Scrum Master              |
| Jacob Odle        | Controls Lead             |
| Zoe Rizzo         | Scribe                    |

While most people have some aspect of the project they plan to focus on, everyone will be collaborating and kept in the loop about each part of the project. The Scrum Master will ensure we are kept on track to meet deadlines. We will have both all-hands meetings for all team members, and sub-team meetings to work on specific aspects of the project.