
## Maker Faire

Maker Faire is a similar exhibit style event at as Imagine RIT with more focus on engineering and makers space that has a few members coming from outside of the RIT organization. This event is in the fall semester around mid-November.

This event is really good for getting a practice at demoing our project and getting used to talking about this capstone project.

![[IMG_6311.jpg]]

## Demo setup Instructions

1. Turn on all raspberry pi's
2. Connect the router to both raspberry pi's with ethernet cords
3. SSH into unc(`ssh pi@talos.local`)
4. Run erv script on unc(`erv` alias can be used to start this from any location)
5. Connect laptop to the router run commander(ethernet recommended for reduced latency)
6. Verify camera for unc is working via the browser(http://talos.local:8889/camera)
7. Verify camera for bluey is working via the browser(http://bluey.local:8889/camera)
8. Make sure to modify the local config to connect commander with unc and bluey:
```
talos.local:
    socket_host: talos.local
    socket_port: 61616
    camera_index: rtsp://talos.local:8554/camera
    acceptable_box_percent: 0.4
    vertical_field_of_view: 48
    horizontal_field_of_view: 89
    confirmation_delay: 0.25
    command_delay: 0.25
    fps: 60
    frame_width: 500
    manual_only: false
bluey.local:
    socket_host: bluey.local
    socket_port: 61616
    camera_index: rtsp://bluey.local:8554/camera
    acceptable_box_percent: 0.4
    vertical_field_of_view: 48
    horizontal_field_of_view: 89
    confirmation_delay: 0.25
    command_delay: 0.25
    fps: 60
    frame_width: 500
    manual_only: false
```
9. Once the commander is open, click on "manage connections"
10. Click on "Connect" for `bluey.local` and wait for the camera to show up
11. Repeat for `talos.local`(make sure the camera feed does show as it can have trouble when multiple connections are opened too fast)
12. Close the manage window and show off the commander interface and move the robot around
13. You can also change between bluey and unc via the dropdown

## Few troubleshooting techniques
- You can't SSH
	- Some devices seems to not be able to discover local domain address when using wireless connection instead of wired, so plug in an ethernet dongle to try.
	- If above fails, you can view the ip address from the router's web interface
	- For macs, you can also share the internet connection and become its own dhcp server, but just use a non-eduroam wifi connection to share because eduroam prevents internet sharing.
- Unc sometimes jitters and does not move in correct direction
	- This can be solved via the pendant by pressing `run + 0 + Start`(homing)
	- Sometimes the home button on the commander works as well
- Unc does not respond to any commands that operator sends(you see the log on operator but no motion on the robot)
	- This is fixed by using the pendant and clicking on `control mode` to enable control mode
- The camera feed is not working
	- Double check that all of the services are running on the RPi
	- The camera-streamer.service should be running without issues
	- The mediamtx.service should be running without issues
	- Use `journalctl -u <service_name>` to see the logs of these services 
	- Use `sudo systemctl <start,restart,stop> <service_name>` to troubleshoot these services
	- Make sure the web cam is also plugged in to the raspberry pi and can be accessed via the camera index 0. If not, fix it by either editing the service file at `/etc/systemd/system/<service_name>` or correctly plugging the hardware
- The camera stream is poor
	- Sometimes the network is overloaded or commander is not able to allocate enough resource to receive the full bandwidth of the video decoding
	- Try stopping one of the robot's connection and see if that improves 
	- Make sure the browsers tab that is watching the live streaming are also closed

## Problems we faced during the fair
- Rain
	- can't stop the weather, but we got tarps to cover over the robot and carry back
- Autonomous control
	- There was simply too much people that is detected at once making it hard for the autonomous coontrol to decide which person to aim


## Recommendation for future Improvements
- Having a screen to block off from the background may be helpful
- We were also thinking of making a demo of a photoshoot booth to make the demo more interactive for the visitors
- Have some basic description and talking points ready, so you can talk with the visitors about what we worked on and the goals of the project.