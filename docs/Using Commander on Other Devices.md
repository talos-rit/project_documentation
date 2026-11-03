# Using Commander on Other Devices
If you want to use Commander on a device other than the one hosting its web server, you'll need to

*(TODO: describe the steps needed to make commander accessible over LAN; also mention that making it accessible over WAN (e.g., thru Tailscale) is out of scope for this guide)*

## Limitations
While you can control the robot perfectly fine from other devices through this method, you may not be able to make full use of the [virtual camera support](Streaming%20and%20Recording%20Video.md). This is because the virtual webcam software runs on the server, not the device you're accessing Commander from. As a result, you cannot record or stream the video feed on this device. Of course, you can still use the virtual camera on the device hosting Commander, but you won't be able to work with it remotely.
*(TODO: clean up this wording)*