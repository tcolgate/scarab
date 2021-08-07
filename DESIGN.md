# Scarab

Scarab is inspired by Locust.

The intension to overcome some short falls in load testing tools.  Generating
load is not typically hard, but managing the results of tests can be awkward.

For a minimum viable implementation:

Manager:
- Long lived
- Capable of orchestrating different test workloads ("profiles"), and collating
  results.
- Test runs should be
  - Named (with sensible defaults.
  - Have all actions (start, stop, updates) annotated, and allow arbitrary
  - A description/reason for running the test should be given.
  - (something about learnings, possibly via annotations as mentioned above)
- Archives results to storage.
- Show in progress running load tests for each profile.
- Report on utilization of workers.

Worker:
- Self contained process, registers the profile it supports with the Manager.
- Waits for request to run one or more instances of the profile. Reports
  metrics back to the manager.

Metrics should be based on prometheus metrics, with the possible exception of
histograms. So far as possible we should reuse the prometheus client library
wherever possible.

#Potential enhancements:

Whilst metrics of test results are of primary interest, additional information
can be useful when interpreting results.

- arrange snapshots of Grafana dashboards that are requested in the profile.
- arrange for collection of logs (from Loki), from workers, and potentially
  from related services, as request in the profile.
- arrange for collection of traces, from jaeger.
- We may also take on the responsibility of announcing load test runs to slack.

#Infrastructure

- gRPC between components
  - manager and workers chat using grpc
  - ui -> manager using grpc
  - ui serves web content and grpc-web
  - ui web interface -> ui server using grpc-web
  - cli -> manager via grpc
