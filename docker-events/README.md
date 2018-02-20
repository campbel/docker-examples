# Introduction
The `docker events` API is useful for monitoring the activity of the docker daemon. This example runs a container that will print the events, JSON encoded, to standard out. If a logging driver is specified, these events can be forwarded to your logging service. The demo is configured for splunk, with the following variables needing to be defined `DEMO_SPLUNK_URL`, `DEMO_SPLUNK_TOKEN` & `DEMO_SPLUNK_INDEX`.

# Run
Start the event logger with Docker compose.
```
docker-compose up
```
In a second terminal, run the sample application
```
docker-compose -f ../example-voting-app/docker-compose.yml up -d
```
The events related to starting the example-vpting-app should now be visible in splunk.
