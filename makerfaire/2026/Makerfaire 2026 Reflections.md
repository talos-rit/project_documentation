## Background: Makerfaire planning & application
We used mostly the same application responses as last year's team did, so see [2025 Fall Application Form](2025%20Fall%20Application%20Form.md) for reference.

As for planning, most of it was done in a post-demo scramble in the two weeks prior to Makerfaire, so see [2026_11_10_notes](2026_11_10_notes.md) and [2026_11_17_notes](2026_11_17_notes.md) for more info on that.

We also did some of our last-minute planning and comms through our Discord server. I made a Makerfaire channel for it where I put a bunch of reference stuff that came in handy, like the list of things to bring, the basic demo plan, the shifts, and our then-remaining questions. That's also where we kicked off the actual search for a new cart for Bluey, though ultimately we decided to keep it on the server rack cart and just use WD-40 on the wheels. I also posted all my configs in that channel which came in handy when debugging why Christine's laptop wasn't getting camera from Bluey (the pi-vision address was set to bingo.local still, from back during our two-bots-on-one-pi setup from the week prior to the sponsor demo).

## Technical Setup
In order to control the robots' speeds, we decided last minute on the night before to switch to my (Briggs's) speed slider fix branches for Commander and Pi-Vision (`feat/tracking-speed` in commander, `feat/live-tracking-speed`). That was as simple as changing the branch selections in Gaia.

That being said, my branch that implements per-axis speed limits (Pi-Vision `feat/travel-budget`) was still open and unmerged, which means we had to drive every axis but the base rotation as slowly as possible. For Bingo that meant driving *all* axes slowly since we couldn't limit the other ones individually, but for Bluey that meant a custom firmware build that massively lowered the max speed of the wrist pitch.

We also encountered a last-minute problem with the elbow and its encoder. For whatever reason, the encoder count was drifting upwards until eventually, the elbow's "home" position was pointed straight upwards. We couldn't fix that in time for the Faire, so we just disabled Elbow in the firmware and re-flashed. See also: [ESP-Driver issue #27](https://github.com/talos-rit/esp-driver/issues/27)

On the bright side, the amount of times we had to re-flash the firmware made us realize that "Maybe there's a better way of doing this". Sure enough, there totally is. A few ways, actually. We plan on investigating those after thanksgiving break. (Context: right now changing the config means re-building the firmware which takes 5 minutes on Aleck's laptop. Not ideal when you're in a time crunch!) See also: [project docs issue #88](https://github.com/talos-rit/project_documentation/issues/88)
For the sake of thorough-ness: the axis 2 encoder is still bad, meaning we also couldn't use shoulder for Bluey.

Bingo used elbow, shoulder, and wrist pitch. Though (due to the lack of soft limits on elbow position) in practice Bingo will never actually use its wrist pitch when the other two are enabled. Fix is in Pi-Vision `feat/travel-budget`, as mentioned before. It's just not been tested enough to be used during an event this big/important.

As for the software on the Pi's: same stack on both of them. The only thing that's different is the identification of their serial & USB devices. Everything else was identical, kept exactly as deployed by Gaia.

The Gaia and Commander configs I (Briggs) was using are available in `./makerfaire-configs-2026-11-20/`, so feel free to explore those. Christine's configs were slightly different, we didn't copy them onto her computer, but they served as a good reference.

## What Our Demo Was
![440](2026_makerfaire_booth.jpg)
We got a corner booth which, for our purposes, is pretty good. It meant people had the room to actually move around and see the bots tracking them. We put Commander up on Christine's TV, hooked up to her laptop. We also had my gamepad (xbox controller) there which helped a lot with fixing aiming.

Which brings me to..

## Problems
So, there are quite a few things that weren't optimal about our booth.

1. People would just walk by, not stopping at our booth, and the robots would track them.

There's not really anything you can do about this aside from trying to get a space that is more nestled away. We were on a corner, true, but it was an *interior* corner. On a major thoroughfare. Lots of people walked by, and not all of them cared about robots. Fair enough. But Bluey and Bingo would still track them regardless.

Which leads into:

2. When leaving our booth, the robots would keep tracking them…

As in, once they were done, the robots just kept following them until they hit their limits. And then, because there were other booths in that direction, the robots would just stay locked over there until you turn off auto-track.

This led to us having to turn off auto-track with the gamepad, then manually jog the robots back to roughly home, then turn auto-track back on.

3. Before the Faire opened, we had a weird issue with Bluey's ESP32

It just stopped sending telemetry entirely. This might have been the problem from our sponsor demo a few weeks prior, since an RTS pulse (which resets the ESP) seemingly fixed it. The issue didn't pop back up for the rest of the day (TODO: DID IT?).

4. A weird split-brain problem with Bingo's controller and Operator/Pi-Vision

Sometimes Bingo would fault due to the lack of soft limits on elbow count mentioned before, and occasionally it would even get *stuck* in that state. The pendant wouldn't say anything about a fault, and it'd even let you home the thing (homing isn't blocked by faults in any of our software apparently). But Commander greyed out all movements due to a perceived fault. The only solution I could find was to restart the entire Pi stack.

5. The robots move quite slowly

We were on my branch for letting auto-track take the speed set in Commander's slider, but that requires being at the laptop, which is often far away and requires stepping over cables and between booth visitors. The solution involves the gamepad: instead of the D-pad just duplicating the left sticks (wrist pitch and base rotation), make it increase speed for the currently-selected robot by counts of 5. I will be implementing this soon.

6. There were lots of faults, and clearing faults also requires the laptop

Good thing we're freeing up the D-pad. One of the buttons on it can be made to clear faults/re-enable control. Will also be done soon.

7. **People didn't understand it.**

This is the big one. At a first glance, at a maker fair, it's hard to know what these two robot arms are intended for. Seeing both camera feeds on Commander helps, but still, people will get close enough to the robots that even normal movements don't move them outside of the centering box. So the robots don't move and they're confused over what the deal is.

8. **People come in groups.**

And yet, Pi-Vision still just latches on to the first person it sees, or sometimes just seemingly random people. It's never very clear who's trying to be tracked. If a family arrives all at once, and they're all roughly in the center, then the robots won't move until the correct person moves. And how would any of us know who the 'correct person' is?

## Solutions
- Print out a sign, or use a custom build of Commander, that makes it clear that "**THE ROBOTS MOVE TO KEEP YOU IN THIS BOX**" (i.e., the centering tolerance box).
	- Put it on or very close to the screen, *not* on the robots. People are usually looking at the screen where they see themselves with a weird skeleton superimposed - that's probably also part of why they didn't realize they were being tracked. There's lots going on and they just don't notice the motion, or don't realize that it's automatic (especially if someone's holding the controller at the time).
- Use a custom Pi-Vision that makes it clearer who the tracking subject is.
	- For instance, make the tracking subject be a different color bounding box. Or just don't even draw bounding boxes/pose skeletons for anyone *but* the tracking subject.
	- Or add a massive arrow above their head like the Sims.
- Investigate robot faults for both robots.
	- These problems we had today might be hard to reproduce since the ESP32 problem only happened once, and Bingo's split-brain persistent faults are intermittent. But if we ever encounter either of them again, we should save & analyze those logs!
	- One guess on Bingo: maybe if two faults come in within a very short time window, something locks up somewhere in our stack? Not sure. Total guesswork.
- Make Bluey & Bingo return to center every so often.
	- Either base it off how small someone is in frame (i.e., how far away they are), or lack of relative motion, or something along those lines. Basically, make it so that they don't wander off looking way out in the distance for someone who's not at the booth.
	- This would have to be a custom build of Pi-Vision, probably?
- **Change their soft limits!**
	- This is probably the easiest win for both of them. Just set their soft limits in pi-vision for joint counts (i.e., position) to be a certain level such that, when people walk away from the booth in either direction, the robots lose them and (ideally) return to center after a certain amount of time without a subject.

## Other Reflections
**Get there early:** I got there at 7:20 and others started arriving soon after. We managed to find things to do for the entirety of those ~90 minutes until the Faire opened. Tests to run and things to fine-tune. For instance, fine-tuning those extra-conservative soft limits to just a 'cone' in front of the booth where the actual visitors would be. (This also makes parking easy, though it's not that bad even way closer to 9. But the lot was basically empty at 7!)

**People don't like being recorded:** So print out signs saying that you aren't doing that. I printed out 3 one-sided copies of [talos-booth-signs](talos-booth-signs.pdf):

```bash
lp -n 3 -o sides=one-sided makerfaire/2026/makerfaire-2026-signs/talos-booth-signs.pdf
```

(You don't have to use CLI to print obviously, but it's a massive flex if you do. Not sure if this works on Linux, and it most certainly won't work on Windows.)

Our signs also cover telling people not to touch the robot. That sign is far clearer to read, and it just so happens that nobody even tried touching them.

Caveat about the signs: we often noticed people leaning in to read them, perhaps assuming they contained information about our project. We probably should have put bigger signage somewhere about what our project was - even just the posters from previous years.

**People are interested in the project:** So print out QR codes to the GitHub repo too! You could also put them on the booth sign PDFs I attached above, if you want.

**You can't prepare a speech:** As RIT students, we're all well-versed in the elevator pitch, so I highly recommend you just ad-lib it. Not to say you should make it up from scratch each time - you should probably come in with *some* sort of plan for what you're going to say - but you can't be planning to say that exact same thing to every person. Some kids will show up, and they just want to see the robots track them. That's perfectly fine. They don't need to know that the motor hat has a servo on it with a really low max frequency which is why Bluey is actually slightly more annoying than Bingo, despite Bingo's controller having a much louder fan than the one in Bluey's power supply. That sort of detail can be reserved for the nerdier visitors!

For reference, my bit was something like this:

> Yeah, so these are old educational robot arms from the '80s or '90s. A lot of them are sitting in storage closets at schools and universities collecting dust, and a lot of them can be found on eBay for relatively cheap. They used to be $12k, now you can find them for $300 or so. We're a team of software engineers working on a senior project where we open up these robots so that anyone can interface with them way more easily. Right now we're running software that makes them track people - all it's trying to do is center your eyes in that blue box on the screen. So if you move until your face is out of that box, the robots will move to compensate. The software for figuring out where your face is runs on the Pi 4's, running at about 4 frames per second. We moved them to the Pi's to minimize latency - it's all a closed loop. Essentially, when you move out of that blue box, we look at what direction you moved in and tell the robot controller to move the cameras to compensate.
> 
> The controllers, by the way, are hard to come by. A lot of people just threw them out. You can see in this big black box, it's old and heavy and slow, that's what Bingo's controller looks like. Over there in that shoebox is something our team's worked on - it's modern hardware that's essentially mimicking that old controller. We use an ESP32 on our custom firmware, communicating with the exact same serial protocol at 9600 baud, plus some motor hats and other stuff to drive the motors.

But obviously the younger visitors might not care about the frames per second of pi-vision or what the baud rate of the custom controller is. So tweak accordingly. And mention that we're open-source, etc. - plus whatever your team is focusing in on. If the slider still isn't working by next year, perhaps your team is working on the slider. Talk about that! Especially with the guests around our age - they're often interested in the details, so talk about what struggles you've had, etc. And listen to them, too. Lots of people with interesting background go to Makerfaire!

## A Note on Logistics
All-in, we spent ~1hr moving the robot from the project room in Golisano all the way to the field house on Friday evening - we started a little after 6PM and managed to have the robots in-place a little before 7PM.

On Saturday evening, we did even better and managed to pack up our whole booth in ~50 minutes, including moving the robots back to GCCIS and taking down the (admittedly very few) signs we had, etc.

Expect it to be a very loud, unpleasant hour. The carts make the loudest noise imaginable. This could perhaps be mitigated by removing some of the loose objects on the bottom of Bluey's server rack cart, but that's still only partial. Even Bingo's cabinet made a ruckus.

We were lucky enough to be able to move them outside, which probably was better for both our ears and avoiding being a public nuisance. Hopefully it isn't snowing for future teams…

## One Last Problem
At ~4pm, Bingo stopped working. We're currently not entirely sure of the cause, but it seems like some sort of split-brain problem between Operator and the controller. This will hopefully be fixed by ImagineRIT & by next MakerFaire.

**UPDATE:** On Monday, 2026-11-23, I was able to get Bingo working. The fix was to plug in the serial cable. Our best guess is that it simply got loose, perhaps from someone bumping into the cart at Makerfaire, causing the occasional jog prompt to get lost. Operator did not have a way to recover from even just one lost prompt, so it would simply stop working and not recover until Operator is restarted, potentially also needing a controller restart if it's in a jog prompt mode so that their state machines match.

An alternative explanation would be that it's simply due to the long nature of the Faire; something overheated after the first 7+ hours of running causing just the right thing to break for the most mysterious, incomprehensible bug possible. But unfortunately the serial link seems more likely.

In our defense, it'd been a long day for all of us.
