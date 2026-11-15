
**Position Title:** Junior Embedded Engineer for Senior Project
### Project Overview and Existing Team
**Project Sponsor:** Samuel Malachowsky
*Note: Since this is a research-based project, the project sponsor and project coach are the same person.*

**Project Summary:** This research project repurposes legacy robotic hardware, the ScorBot ER-4pc and ER-V (nicknamed Bluey and Bingo), into an autonomous camera platform. The project's goal is to automate camera operations which typically require a person: panning, tilting, and framing subjects. By utilizing two robotic arms, two camera feeds can be integrated and combined, autonomously and dynamically deciding which feed to highlight. This project includes ideating, research, implementation, and refinement of software that combines machine learning, computer vision, and other technologies to guide camera movements. The previous two iterations of this project focused on subject tracking and a new controller for the ER-4pc. The plan for this year includes refining subject tracking and implementing predictive movements, a digital twin to use for testing and debugging, and an updated controller user interface for more intuitive design. This year’s main goal is multi-camera integration.

**Team Roles:** Our team is comprised of six members: hardware lead, control lead, computer vision lead, communications lead, scrum master, and scribe. Most members specialize in one aspect of the project, while some support multiple areas.
### What Has Been Accomplished
#### Subject Tracking
Talos is currently able to differentiate between subjects in frame. The user can select which subject to track using a dropdown, along with a zoom toggle and slider to indicate if and how much to zoom in on the subject.
#### Commander User Interface
The commander user interface, which is what users use to send controls to the robots, has been created using React. It includes two feeds (one for each robot), a tracking panel to select the subject, a zoom toggle, manual controls, and a settings panel to configure the streams. 
#### Robot Hardware
Bluey’s six axes can now all be driven through the ESP32 which translates commands from the Raspberry Pi and aims for feature parity with the original controller. Each axis is driven only up to a soft limit derived from its actual physical range of motion so movement always stays within safe bounds.
#### Digital Twin
A digital twin was created as an accurate telemetry informed mirror of Bingo's arm positions. The majority of it's functionality has been ported to commander, but it served as a good test bed to get more commands running.
### Second Semester Plan
The main goal for the second semester is multi-camera integration so Talos can dynamically switch between the two feeds without a user manually selecting one. Subject tracking will be enhanced by doing research into predicting a subject's future movement, and implementing identity recognition for reacquisition. Camera framing will be improved by implementing a shot library and adding support for different framing modes and camera views. For hardware, a custom controller PCB is being designed for feature parity between the two robot arms.
### Technical Description
We are specifically looking for a co-op interested in embedded and computer engineering. This role involves working with an ESP32 via ESP-IDF using I2C to safely control a robotic arm.
#### Roles and Responsibilities
- Maintain and extend the PCA9685 motor driver, encoder reading, limit-switch handling, and ADS1015/ADS1115 current sensing
- Overhaul ESP driver unit testing efforts, write and run on-device unit test suite, and bench-test on the real arms
- Set and tune per-axis soft limits and homing behavior
- Work with the hardware lead to bring up the custom controller PCB and achieve feature parity between the two arms
- Maintain the ESP32 ↔ Commander/Pi socket protocol and document any changes
- Attend all full-team meetings and work alongside senior project team members to accomplish tasks
### Preferred Candidate Attributes
**Qualifications:**
- Completed SWEN-340 (Software Design for Computing Systems) or equivalent experience
- Experience with C and Python
- Basic understanding of Linux file systems
- Interest in embedded engineering
- Strong problem solving, analytical thinking, and communication skills
**Preferred Qualifications:**
- Previous experience with Raspberry Pis
- Previous experience with embedded real-time operating systems
- Previous experience with concurrency
- Exposure to Docker and containerization
- Exposure to OpenCV