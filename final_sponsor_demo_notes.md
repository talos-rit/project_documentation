# Notes from end of Milestone 3 Demo
Attendees
1. Malachowsky
1. DQ
!. Kiser
1. Meenely
1. Ryan Y
1. J R.
1. Connor 
1. Kai F.
1. Hiro T.


Ryan Demonstrates hardware hookups

Professor M. asked about limit siwtches, Ryan explains how they are all tied together.

Current sense limit switch homing questions from malachowsky


Malachowsky asked about the viability of the high school tech teacher being able to port the esp-driver to their own robot.
This is something we need to continuously monitor.


After this, Hiro demoed the TUI on the projector.
port 61616 arbitrraily defined? or is it a standard? it is, for ActiveMQ. we should change that default port in operator.

DQ asked about coordinating system for the robot.
Hiro: ACL uses cartesian coordinates and polar coordinates.


Media MTX stream over network and onbaord virtual camera were shown.

Malachowsky still wants a digital twin of the robot arm! Still in future.md 

Future plans for director to predict motoion of subject based on pose?

Malachowsky suggests rather than tracking the largest or smallest bounding box, track a person with a token, DQ chimed in to track the person that raises their hand. DQ suggests some type of object.

Malachowsky asked if the model can persistant recognize that each bounding box is the same person from the previous frame. Hiro said we use a hybrid apporach, where every few seconds the persistance is cleared, but otherwise the model can track persistance between people.

DQ suggests tracking audio sounds to find who is talking. DQ suggest simplyfying tracking system, so that the system works well the first time and it is easy for end users to understand. 

DQ suggest that the YOLO tracking be able to accept input on what to track, like telling the yolo model to track whiteboards!


Open-source-ness, how might an interested devloper get started with helping on this project? 

Malachowsky suggest having other roles read and touch the other roles documentation and code.

Malachowsky learned about the deepwiki connection through the github.

Malachowsky asks team to put togehter a list of nearterm midterm and long term goals to set a firmer scope for the next team. 

Malachowsky Move limit switches to daugther board.

Contribution pathway should be carfeully considered. Have a boacklog of things that should be done, mark some as "good for a newbie". COMMUNITY ENGAGEMENT is a big deal! fostering a community around a program or idea. Meenely asked about First robotics?

Malachowsky concerned with current draw with overspeccing the motor drivers, due to human safe aspect.

Malachowsky wants hiro to add the camera mount he designed and 3D printed on bluey to the project repo.


Malachowsky asks us to make a prompt for the next team to narrow down what the project actually needs. Basiclly write the intro segment that we got for the new team.


Malachowsky needs an inventory of the hardware both on the robot and extra.