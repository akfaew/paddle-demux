# Paddle Demultiplexer
Paddle only allows one webhook URL, so we need a Paddle Demultiplexer.
This tool is designed to run as a Google AppEngine project on the free tier.

## Setup
- Edit `main.go` and enter your webhook URLs along with their matching Subscription Plan IDs.
- Create a new Google AppEngine project.
- Edit the `PROJECT` variable in `Makefile`.
- Type `make deploy`.
