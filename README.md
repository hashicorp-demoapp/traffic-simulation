# Traffic Simulation

[![CircleCI](https://circleci.com/gh/hashicorp-demoapp/traffic-simulation.svg?style=svg)](https://circleci.com/gh/hashicorp-demoapp/traffic-simulation)  
Docker Image: [https://hub.docker.com/repository/docker/hashicorpdemoapp/traffic-simulation](https://hub.docker.com/repository/docker/hashicorpdemoapp/traffic-simulation)  

## Parameters

The following parameters can be set when running traffic sumulations. All parameters are set as environment variables

### BASE_URI

Base URI for all requests

**default**: http://localhost:18080"

### TIMEOUT

Timeout for each scenario

**default**: 60*time.Second

### DURATION

The total duration to run the tests

**default**: 5*time.Second

### USERS

Number of concurrent users

**default**: 5

### SHOW_PROGRESS

Show graphical progress output, set to false when running on a scheduler

**default**: true
