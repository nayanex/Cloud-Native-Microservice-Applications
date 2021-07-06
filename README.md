# nd064_C1

## Exercise: Endpoints for Application Status
This exercise aims to extend a Python Flask web application with status and metrics endpoints.

## Exercise: Application Logging
Logging is a core factor in increasing the visibility and transparency of an application. When in troubleshooting or debugging scenarios, it is paramount to pin-point the functionality that impacted the service. This exercise will focus on bringing the logging capabilities to an application.

At this stage, you have extended the Hello World application to handle different endpoints. Once an endpoint is reached, a log line should be recorded showcasing this operation.

In this exercise, you need to further develop the Hello World application collect logs, with the following requirements:

* A log line should be recorded the timestamp and the requested endpoint e.g. `"{{TIMESTAMP}}, {{ ENDPOINT_NAME}} endpoint was reached"`
* The logs should be stored in a file with the name app.log. Refer to the logging Python module for more details.
* Enable the collection of Python logs at the DEBUG level. Refer to the logging Python module for more details.

Note: For the environment setup, follow the instructions in the previous exercise.

Tips: If you get stuck, feel free to check the solution video where detailed operations are demonstrated.
