# Imagine Emailer
This is used to send emails using AWS Simple Email Service (SES). Because we were using a router, this runs on a pi connected to both the internal and external network. It listens for API calls and takes the name, email, and image and sends the email over SES. `aws configure` is required to run before this server.

It is setup to be used on the `imagine-rit` branch of commander.
