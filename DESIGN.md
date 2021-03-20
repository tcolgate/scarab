# Scarab

Scarab is inspired by Locust.

The intension to overcome some short falls in load testing tools.  Generating
load is not typically hard, but managing the reults of tests can be awkward.

For a minimum viable implementation:

Manager:
- Long lived
- Capable of orchestrating different test workloads ("profiles"), and collating
  results.
- Archives results to storage.
- Show in progress running load tests for each profile.
- Report on utilization of workers.

Worker:
- Self contained process, registers the profile it support with the Manager.
- Waits for request to run one or more instances of the workload.
- Reports metrics back to the manager.

Metrics should be based on prometheus metrics, with the possible exception of
histograms. So far as possible we should reuse the prometheus client library
whereever possible.

#Potential enhancements:

Whilst metrics of test reults are of primary interest, additional information
can be useful when interpreting results.

- arrange snapshots of grafana dashboards that are requested in the profile.
- arrange for collection of logs (from loki), from workers, and potentially
  from related services, as request in the profile.
- arrange for collection of traces, from jaeger.

We may also take on the responsibility of anouncing load test runs to slack.

#Infrastructure

- grpc between componenents
  - manager and workers chart using grpc
  - ui -> manager using grpc
  - ui serves web conent and grpc-web
  - ui web interface -> ui server using grpc-web
