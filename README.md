# Paddle Demultiplexer
[Paddle only allows one webhook URL](https://jasminek.net/blog/post/paddle_multiple_webhooks/), so we need a demultiplexer.

## How it works
This tool will run on a free tier Google AppEngine instance and listen for webhooks.
Once it receives a webhook it will look into its Body and redirect the payload to a specified URL based on `subscription_plan_id` (also known as `Plan ID`).
The path is not configurable, meaning that all your apps must have the same URL structure for Paddle webhooks (e.g. `/paddle/webhook`).

This tool doesn't validate the payload, it's up to your app to do that.

## Setup
- [Create a new Google AppEngine project](https://console.cloud.google.com/projectcreate), let's call it `your_project`.
- Edit `main.go` and enter your webhook URLs along with their matching Plan IDs. You can find your Plan IDs [here](https://vendors.paddle.com/subscriptions/plans).
- Edit the `Makefile` and set `PROJECT` to `your_project`.`
- Type `make deploy`.
- [Tell Paddle](https://vendors.paddle.com/alerts-webhooks) to send webhooks to `https://your_project.appspot.com/your_webhook_path`.
