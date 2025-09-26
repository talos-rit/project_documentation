# Project Plan

# Project Goals & Scope Statement

This research project seeks to repurpose the ScorBot ER-4pc and ER-V, an educational robotic arm, into an autonomous robotic cameraperson. While basic videography techniques are well understood, capturing high-quality footage in dynamic environments like classrooms or live presentations requires constant operator attention. The project's goal is to automate camera operations—such as panning, tilting, zooming, and framing subjects—to create more engaging recordings. This project includes the ideation, research, and development of software that integrates machine learning, computer vision, and other technologies to guide camera movements. The previous iteration of this project completed the full control of the ER-V with a basic tracking algorithm and a basic driver for starting to control the ER-4pc without its original controller. Our scope plans to further the research by completing the functionality of the ER-4pc and refactoring the controller to interface with both robots simultaneously. We will also be modifying the algorithm to use better framing and videography techniques, as well as working with both camera angles to swap between. Additionally, we plan to refactor existing code in order to increase the maintainability of the system.

# Planned Milestones & Major Deliverables

1. Improved camera framing with the ER-V  
     
2. ER-4pc working with a custom controller to at least the same degree as the ER-V  
     
3. 2 arms moving 2 cameras that can swap between the different camera views and perform basic videography techniques.

# Initial requirements & User Stories

The main user profile for this camera system is a highly skilled lecturing professor who wants to give a motivational speech and lecture to their students online in a zoom meeting and students offline projected on a screen in a large lecture hall. To improve the delivery of his speech, he wanted a great video of himself giving the speech. All of the requirements for the project stem from this user profile and their needs.

Professor User Stories:

1. As a professor, I want to download and run the application so that I can use the features.  
2. As a professor, I want to connect to the robot from my device so that I can run the program on the robot.  
3. As a professor, I want to save the video captured by the robots so that I can use it in the future.  
4. As a professor, I want to use 2 robots to capture different views so that I can use video editing software to swap between angles.  
5. As a professor, I want to have the robot automatically figure out the best shot from both of the robots, so that I can focus on my speech.  
6. As a professor, I want to stream the video data captured by the robots to zoom so that my online audience can see me with high quality footage.

Developer User Stories

1. As a developer, I want to be able to package my software so that it can be installed easily by other developers and used.  
2. As a developer, I want to be able to make my software extendable so that it can be used with other software tools or customizable by other developers.  
3. As a developer, I want to clean up development tools used by the team so that the development speed and maintenance of software is the best it can be.

## System Requirements

### Fix hardware for Ichor

- Fix encoders(6,4)  
- Fix end stops

### Develop Driver for Ichor

- Ability to drive motor simultaneously   
- Ability to handle current surge  
- Ability to read encoder position sensing  
- Ability to return motor states and position

### Commander/Director

- Control multiple bots simultaneously  
- Receive multiple video streams  
  - Process video streams  
- Package Application  
  - Use pyinstaller  
  - Make the application more portable  
- Clean development tools  
- Virtual Webcam  
  - Able to stream the final video result live on another application

# Work Breakdown & Subteams

## Commander Team

Kai & Hiro

## Driver/Hardware Team

J, Ryan & Connor

# Communication and Stakeholder Management Plan

Every Thursday, we meet with our coach, giving status updates on our project and collecting feedback about potential features for the future of the project. Additionally, any concerns about major risks of the project, given the nature of its research-oriented development, are brought up during the meeting and discussed to resolve how to handle the risk.

Sponsor meetings are scheduled as needed based on needs and/or scope changes to the project.

As part of the weekly meetings, 4-Up document is presented to the coach, detailing the progress of the teams, the current risks of the project, the near future plans for the team, and the needs of the team from the sponsor/coach.

