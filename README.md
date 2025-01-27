Echo server to mimic Datadog's statsd daemon. Run it and then run your app, with DD pointed at port 8125, and this server will echo out the DD metrics payloads (in their raw TCP form :) )
