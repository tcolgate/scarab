Opinions on load testing.
=========================

Terminology:
- Componenets under test. This is the actual code you have written
  that you wish to test. It does not initiallyinclude depdencies (DBs, external
  caches)
- System under test. This is the component under test, and all the depdendent
  systems, and hosting systems, networking infrastructure, load balancing,
  and anything else.

What is the goal?
- load testing is not about firing massive load at a platform to
  establish some magic maximum requests per second. This is the
  "car crash" method.
- The goal is to establish the rate of requests our application can
  handle whilst maintaining satisfactory response times.
- It is important to establish the scope of what is being tested.
  - if we have backend depdencies, we should load test them separetly
    and arrange for sufficient capacity to be available for us to
    test our componenet, with the performance of any backend systems
    impacting our results.
  - During testing we may discover that our response times are entirely
    constrained by our depedencies. In this situation we may need to
    scale the depdency to a sufficent size that it reflects the capacity
    available in the final sytem. We may decide our applications own
    overhead is negligible, and load test the depdendency directly.
  - If we are confident that our application can scale linearly (
    assuming capacity in depdencies), then we sould only load test
    an individual instance of our applciation. Multiple instances,
    and especially autoscaling, make it harder to interpret the results.
  - We should arrange that the system under test is stable during the
    load test. This include minimizing changes on the underlying
    infrastructure.

Preparation:
- Collect together any dashbaords that can help determine how the sytem
  under test is performing. These should include indications for load, and
  errors, for as much of the system as possible:
  - performance dashboards for the appplication.
  - performance dashboarsd for the depedencies.
  - performance dashbaords for relevant infrastructure.
- Clear the floor!
  - Ensure that there are no other load tests running against
    dependent systems.
  - Try and minimize maintenance to the system under test and
    related infrastructure during the testing.
  - Warn others that you are starting the test.
- Note the time and purpose of this test.

Running the test:

Collecting results:

Interpreting results:

