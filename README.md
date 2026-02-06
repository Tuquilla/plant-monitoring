# Plant Monitoring
Simple plant monitoring that subscribes to a NATS and get updates from different sensor connected to plant soil.

# Architecture
**sensor** ----> **esp32** --_message_--> **nats** --_message_--> **plant-monitoring**
