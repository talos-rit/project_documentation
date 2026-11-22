#set page(paper: "us-letter", margin: 0.6in)
#set text(font: "Inter", fill: black, lang: "en")
#set par(leading: 0.22em, justify: false)

#let sign(icon, headline, body) = align(center + horizon, stack(
  spacing: 0.45in,
  image(icon, width: 3.7in),
  text(size: 80pt, weight: 900, tracking: -0.025em, headline),
  block(width: 6.6in, text(size: 26pt, weight: 500, par(leading: 0.45em, body))),
))

#sign("no-touch.svg", [Please don't touch the robot], [
  Robot arms can move without warning. \
  Ask us and we'll show you how it works!
])

#pagebreak()

#sign("live-video.svg", [No video is saved], [
  The cameras only help the robots aim. \
  Nothing is recorded or stored.
])
