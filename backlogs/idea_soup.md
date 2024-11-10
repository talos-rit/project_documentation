# Idea Soup
This is a list of ideas that the robotic autonomous cameraperson could potenitally incorporate into its system. 


## Camera
- Camera considerations:
	- 1 Camera solution:
		- Have to figure out how to handle high resoultion streams
		- Tracking alogrithm can see exactly what is being captured
		- More complex video processing
	- 2 Camera solution:
		- Easier to handle high resolution video stream
		- Less complex video processing
		- Tracking alogrithm has to guess what is actually being captured (difficult to syncronize zoom/focus/offset)
    - 3 cameras
		- 2 bots and wide field of view cam to see the whole room
		- wide view cam as backup to bots
		- wide view cam as default for AI to fallback on if robots don't know that to focus on

- Camera meta data
    - confidence intervals for each camera feed
		- one for how good the shot is on a person
		- another for how "right" the shot is in terms of who is should be pointed to

- error detection from the camera feed


## AI Tracking
- AI Videography
	- AI controls which stream to capture (destructive editing)
		- low skill user
		- loss of event recording
		- live processing
	- AI gives user suggestions based on all footage coming in (not as wanted)
		- high skill user
		- post-processing
- AI-processes HD video stream or low quality video stream?
- A passive videographer
- Active noise cancellation


## Robot Behavior
- robot moves to look at object on table
	- like a classroom visualizer
	- visualizer mode?
	- rotating shot?
	- Marques Brownlee-like camera rotation shots
		- he has a cool demo

- save positions from manual interface

- digital twins of actual environments and use robot as scanning device

- interview professors for what they would want in their class room. 

- remembers the face of the primary presenter and only follows that face

## UI Interface
- Visualizer to see what the tracking algo sees
- 2D visualizers on the manual interface that displays the positions of the arms
- mockup fake people and fake arm to control
- draw "bounding box" on the live video
	- draw the center target


## Implementation

### Operator
- Coded control to prevent elbow from over-extending. 

### Director
- Add options for framing locations (center, left, top, topleft, etc.)


## Usability
- Low cost assembling, for flexibility of users?
- Control Talos from across the room (control over the network)