# How to Locally Test Commander (to the greatest extent possible)
> [!note] 
> This guide is written by human hands to make sure it's actually useful and adequately thorough! The target audience is moderately technical (able to use a terminal); you needn't be a programmer to follow along.
> 
> _**Disclaimer:** This guide works best on macOS or Linux, but should be followable on Windows too, albeit perhaps with minor adaptation of certain commands. I tried my best to find those alternatives where applicable!_

> [!hint] Hint
> If you don't want/need any explanation, feel free to just skip to the code block sections to see the concrete steps!

This guide explains how to run and test Commander's WebUI locally, including as many features as possible (basically everything except the actual robot movement commands). It assumes you already have `commander-web` in a runnable state on your machine; see [Commander's README](https://github.com/talos-rit/commander/blob/main/README.md) for how to do that. The tl;dr is:

```bash
git clone https://github.com/talos-rit/commander
cd commander
uv sync --extra yolo # we should only need the yolo models now that mediapipe is deprecated

# build commander-web
cd web
bun install # use whatever package manager you want; npm works fine, it's just slower
bun run build # again, feel free to use npm or pnpm here
cd ..

# now optionally run it to make sure it works
uv run commander-web
```

*(**TODO:** streamline this in the future, that's too many steps...)*

Then you would open `http://localhost:8000` in your browser of choice and configure a robot. *Typically* that would be an actual Scorbot setup with Operator + cam_streamer + pi-vision, and if you've already got Commander set up, it would indeed be linked with that stack. If that's the case, just Add a new Robot; otherwise, follow along!

In either case, Commander-Web doesn't need to be running right now since we need to go set up other things, so feel free to shut it down if you wish.

## A Short Primer
As mentioned above, the typical production setup for Commander would include links to the following:
- **Operator:** Exposes a standardized HTTP API for interacting with the Scorbot
	- Runs on a Raspberry Pi that's plugged into (1) a webcam (*for Cam Streamer*) and (2) the Scorbot controller box via USB-Serial cable (*for Operator*).
	- Commander uses this to control the arms. It gets subject location data from pi-vision, determines a bounding box for the subject, and (if that bounding box is outside the box of acceptability) sends movement commands to the robot.
- **MediaMTX:** Generates an [RTSP (Real-Time Streaming Protocol)](https://en.wikipedia.org/wiki/Real-Time_Streaming_Protocol) server on port 8554.
	- Runs on the same Raspberry Pi as Operator & Cam Streamer.
	- This is set up during Cam Streamer's setup process; see [Cam Streamer's README](https://github.com/talos-rit/cam_streamer/blob/main/README.md) for info about that process.
	- Commander uses this RTSP feed to display the camera feed from the robot's webcam.
- **Cam Streamer (FFmpeg Script):** An FFmpeg script that publishes an H.264 feed from the robot's webcam to the RTSP port.
	- Runs on the same Raspberry Pi as MediaMTX.
	- Commander uses this in the same way as MediaMTX; they work together to expose and provide video to that feed.
- **Pi Vision:** A Python script that runs YOLO models for subject detection and pose estimation.
	- Can run on *anything*, but preferably runs on the same Pi as everything else to minimize latency.
	- Commander receives position + pose information from Pi Vision, which it uses to determine whether (& where) the robot needs to move. It also uses that information to draw pose and bboxes over the RTSP video feed.

> [!question] So, what needs to change to test vision stuff completely locally?

The short version is to run everything besides Operator on your local machine.

The real question is, *how?* After all, most of the documentation for these repos assumes you're working with a remote Pi, so it can be unclear how to actually set all that up.

The answer depends - do you want to use a premade test video, or your webcam? Webcam is recommended for ad hoc testing, but this guide will explain how to do both.

## Local Webcam
> [!question] What do we need?

Clearly, we need:
- To expose an RTSP server on port 8554
- To publish your webcam to that port
- To run Pi Vision locally

Cam_Streamer's scripts won't work well for this because they assume the target is running Linux and your computer *probably* isn't running Linux. And even if it is on Linux, is it using Arm64 Linux? Probably not! So cam_streamer probably won't work for you.

The simplest way to accomplish these things is by addressing these requirements literally. Keep in mind, Cam_Streamer is just some setup scripts. We can emulate that pretty trivially!

### RTSP Feed Setup
This is the important part, where we essentially recreate cam_streamer from first principles.

#### MediaMTX
First, install MediaMTX; it's still the easiest way to set up an RTSP server. This changes based on your OS.

**On MacOS:**
```bash
# assuming you have homebrew installed:
brew install mediamtx
brew services start mediamtx
```

**On Linux/Windows:**
1. Download the appropriate binary from [MediaMTX's Github Releases](https://github.com/bluenviron/mediamtx/releases)
2. Extract the `.tar.gz`/`.zip` file and `cd` into it
3. Run it
	1. Windows: `.\mediamtx.exe`
	2. Linux: `./mediamtx` (you may have to give yourself execute perms w/ `chmod +x mediamtx`)

I haven't tested on Win/Linux, but presumably these run in the foreground. If so, open a new terminal for the next steps.

*(Note: You could also use Docker, or `yay` if you're on Arch (btw), but that is outside the scope of this guide.)*

#### FFmpeg
Again, the installation process depends on OS:
```shell
# on macOS:
brew install ffmpeg

# on Linux:
sudo apt-get update && sudo apt-get install -y ffmpeg 

# on Windows:
winget install Gyan.FFmpeg
```

Now you need to figure out which device is your video device in order to start sending video to the RTSP server. This depends *heavily* on operating system!

##### MacOS:
```bash
ffmpeg -f avfoundation -list_devices true -i ""
```

On my MacBook, that yielded the following:

![[FFmpeg Video Device Selection MacOS.png]]

We're interested in this bit here:

```bash
[AVFoundation indev @ 0x146004770] AVFoundation video devices:
[AVFoundation indev @ 0x146004770] [0] MacBook Pro Camera
[AVFoundation indev @ 0x146004770] [1] OBS Virtual Camera
[AVFoundation indev @ 0x146004770] [2] MacBook Pro Desk View Camera
[AVFoundation indev @ 0x146004770] [3] Capture screen 0
```

My main webcam is MacBook Pro Camera, so my video device will be device 0.

Now, use FFmpeg to stream video to the RTSP feed:

```bash
ffmpeg -f avfoundation -framerate 30 -video_size 1280x720 -i "0:none" \
  -an -pix_fmt yuv420p -c:v libx264 -preset ultrafast -tune zerolatency -g 30 -bf 0 \
  -f rtsp -rtsp_transport tcp rtsp://127.0.0.1:8554/camera
```

**What this command does:**
- **`-f avfoundation`:** Input format will be Apple's [`avfoundation`](https://developer.apple.com/av-foundation/) (normally this can be auto-detected, but we might as well specify it ourselves)
	- _**This is one of the two FFmpeg parameters that actually changes between OSes!**_[^3]
- **`-framerate 30 -video_size 1280x720`:** The RTSP feed will be 720p30
- **`-i "0:none"`** Sets the input video device to `0` and audio device to `none`
	- _**This is one of the two FFmpeg parameters that actually changes between OSes!**_[^3]
- **`-an`:** Disable audio stream entirely
- **`-pix_fmt yuv420p`:** Nerd stuff (sets the [pixel format](https://en.wikipedia.org/wiki/Pixel_format) to [yuv420p](https://en.wikipedia.org/wiki/Y%E2%80%B2UV))
- **`-c:v libx264`:** Use software encoding (if your computer can handle it, which it probably can!)
- **`-preset ultrafast -tune zerolatency -g 30 -bf 0`:** These are encoder options, aka more nerd stuff. For the curious:
	- **`-preset ultrafast`:** We need real-time video so let's not care about quality and focus on speed
	- **`-tune zerolatency`:** Tunes various encoder settings to further focus on speed
	- **`-g 30`:** Sets the Group of Picture (GOP) size to 30, AKA the maximum gap between keyframes, measured in frames
		- Keyframes are i-frames that start a Group of Pictures
		- i-frames are coded only from their own pixels, as opposed to p-frames which predict from earlier frames[^1] and b-frames which are *bi-predictive* and may use both earlier and later frames for their compression[^2]
		- So this setting basically means, "The stream should have a real, full keyframe *at least* once per second", which is important because *you can only start viewing a stream (or recover after a lost packet) on a keyframe!*
	- **`-bf 0`:** Turns off b-frames since they are not suited to low latency encoding/streaming
- **`-f rtsp`:** Sets the file format to the correct one for RTSP
- **`-rtsp_transport tcp`:** Sets the transport protocol to TCP instead of UDP, which is the default
	- UDP is the default protocol for video packets since it's normally okay to drop a few frames here and there
	- The RTSP control connection is **already** TCP, so this option basically just says, "Interleave the video feed with the existing control connection so that nothing is sent as a separate UDP stream"
	- We want to use TCP because, if we were using UDP, one lost packet would leave the decoder wrong and errored until the next keyframe, and as mentioned earlier we have the keyint/GOP set to 1 second! So we use TCP to avoid losing these h.264 reference frames.
- **`rtsp://127.0.0.1:8554/camera`:** This is the output location; for us, that's the MediaMTX feed

##### Linux:
The device is usually `/dev/video0`:
```bash
ffmpeg -f v4l2 -list_formats all -i /dev/video0
```

If so, start streaming video to the MediaMTX RTSP feed:

```bash
ffmpeg -f v4l2 -input_format yuyv422 -framerate 30 -video_size 1280x720 -i /dev/video0 \
  -an -pix_fmt yuv420p -c:v libx264 -preset ultrafast -tune zerolatency -g 30 -bf 0 \
  -f rtsp -rtsp_transport tcp rtsp://127.0.0.1:8554/camera
```

You can see the [[#MacOS]] section above for info on what all of that means. Notably, the only part that changes from the mac command are the input format ([Video4Linux](https://en.wikipedia.org/wiki/Video4Linux) instead of avfoundation) and input device (`/dev/video0` instead of `"0:none"`). Everything else is the same.

##### Windows:
This is the name of your webcam. You can find it via this command:
```powershell
ffmpeg -list_devices true -f dshow -i dummy
```

If your webcam is named `Integrated Camera`, you'd use the following command:

```powershell
ffmpeg -f dshow -framerate 30 -video_size 1280x720 -i video="Integrated Camera" `
  -an -pix_fmt yuv420p -c:v libx264 -preset ultrafast -tune zerolatency -g 30 -bf 0 `
  -f rtsp -rtsp_transport tcp rtsp://127.0.0.1:8554/camera
```

Be sure to adapt it for whatever the name of your webcam actually is!

### Pi Vision
Now, you should have an RTSP feed streaming from your webcam. You can verify by opening it in your favorite video player e.g., MPV or VLC, or by adding it as a source in OBS, or by trying to add it into Commander, which **should accept it and show you the video feed at this point!**

If you're fine with running no pose estimation, you can actually use Commander's built-in subject detection as of writing. However, do keep in mind that doing so may not be supported into the future. But we want to test ***as much of Commander as is possible without a robot***, so we'll be setting up pose estimation in this guide too.

Thankfully, Pi Vision is a platform-agnostic script. Start by cloning it in yet another new terminal and installing its requirements:

```bash
git clone https://github.com/talos-rit/pi-vision
cd pi-vision
pip install -r requirements.txt onnxruntime
```

Unfortunately, that doesn't make things perfectly straightforward. This is because you'll need the pose model. There are two ways of getting it. The simpler way is copying it from the Pi, which should be as easy as...

```bash
scp pi@bluey.local:/home/pi/test/tracker/yolo11n-pose-320.onnx .
```

However, if you don't have access to the Pi (as is likely if you are following this guide), you'll have to download the model yourself:

```bash
pip install ultralytics
yolo export model=yolo11n-pose.pt format=onnx imgsz=320 dynamic=False
mv yolo11n-pose.onnx yolo11n-pose-320.onnx
```
*(Note: you may need to change the `mv` command if you're on Windows, although powershell should still support that syntax... Alternatively, you could just rename the file in File Explorer.)*

Now, you should be able to run pi-vision!:

```bash
python main.py --camera rtsp://127.0.0.1:8554/camera --port 5050 \
  --model ./yolo11n-pose-320.onnx --edge-control
```
*(Again, change paths and syntax as needed for Windows.)*

We're using `--edge-control` here so that Commander's status poll succeeds (to avoid warnings in commander's UI). Auto-tracking will still fail, of course, but it's one less warning to worry about.

Now, pi-vision should be running! (Currently, you can verify this by going to `http://0.0.0.0:5050` in your browser, which is pi-vision's own web UI)

So let's put it all together...

### Commander WebUI Configuration
Now it's time to start Commander in our final terminal[^4]:

```
cd commander
uv run commander-web
```

Now, navigate to `http://localhost:8000` in your browser and go to Settings (the cog wheel at the top right) then click "Add robot". Set the following:

- Operator Host: `localhost`
- Port: (can stay as `61616`; we aren't using operator so it doesn't matter)
- Camera: `rtsp://localhost:8554/camera`
- PiVision URL: `http://localhost:5050`
- FPS: 30

As of writing, that looks like the following:
![[Commander Localhost Robot Setup.png|464]]

Once these are filled out appropriately, click "Add robot" at the bottom then, under Cameras, set the dropdown for "Camera 1 (primary)" to be `localhost`. Then click Save. These can be seen as steps 2 and 3 in the image below:
![[Commander Localhost Setup.png|402]]
*(Again, the UI may have changed since this guide was created; the important part is setting the active camera (if needed) and Saving.)*

Now you may exit settings. You should be able to see yourself, complete with pose estimation! Try clicking "Pause PiVision" at the bottom; you should see the bounding box and pose skeleton disappear[^5]. You can also try other stuff, like loading a local subject tracking model by picking a model size on the slider in the debug rail. In fact, you can even have that local model running *on top of pi-vision* (at time of writing)! It draws both sets of bboxes. Interesting.

If you wanted to, you could probably hook up a second webcam to your computer to test camera switching. This would involve creating a new MediaMTX server (though i have no clue how to do this), starting a new ffmpeg process to push the second webcam's video feed to it, and starting a new pi-vision instance. All of these would be running on new ports, of course. This is purely theoretical though; i haven't tried such a thing yet.


## Video on Disk

*(TODO: can this be done with video still? i doubt pi-vision would work with it, tho commander's old UI used to...)*

[^1]: (the simple version is "only changing the pixels that have changed since the previous frame", the more complex version involves motion vectors and lots of compression nonsense that i won't pretend to understand)

[^2]: (don't ask me to explain this one)

[^3]: unless you get a weird bug with pix_fmt needing to be mjpeg on linux or something along those lines, but that's out of scope here...

[^4]: one could even call it the terminal terminal... (i'm sorry)

[^5]: assuming bboxes and skeletons are still enabled by default when you're reading this!
