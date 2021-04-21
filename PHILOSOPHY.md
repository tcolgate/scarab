Opinions on load testing.
=========================

Terminology:
- Components under test. This is the actual code you have written that you wish
  to test. It does not initially include dependencies (DBs, external caches)
- 
- System under test. This is the component under test, and all the dependent
  systems, and hosting systems, networking infrastructure, load balancing,
  and anything else.

What is the goal?
- load testing is not about firing massive load at a platform to
  establish some magic maximum requests per second.
- The goal is to establish the rate of requests our application can
  handle whilst maintaining satisfactory response times.
- If we are confident that our application can scale linearly ( assuming
  capacity in dependencies), then we should only load test
  an individual instance of our application. Multiple instances,
  and especially auto scaling, make it harder to interpret the results.
- All systems have limits. Our goal is not to establish where the breaking
  point of our system is (the car crash method), but to determine the
  base line performance, and the load that causes that performance to drift
  beyond our desired performance.
- It is important to establish the scope of what is being tested.
  - Design a test that is representative of real world traffic. Sometimes
    this can be a single identical query. It is more common that different
    requests exert different loads, so the test should represent a reasonable
    range of load that might be seen under "normal circumstances".
  - if we have back end dependencies, we should load test them separately
    and arrange for sufficient capacity to be available for us to
    test our component, with the performance of any back end systems
    impacting our results.
  - During testing we may discover that our response times are entirely
    constrained by our dependencies. In this situation we may need to
    scale the dependency to a sufficient size that it reflects the capacity
    available in the final system. We may decide our applications own
    overhead is negligible, and load test the dependency directly.
  - We should arrange that the system under test is stable during the
    load test. This include minimizing changes on the underlying
    infrastructure.

Preparation:
- Collect together any dashboards that can help determine how the system
  under test is performing. These should include indications for load, and
  errors, for as much of the system as possible:
  - Performance dashboards for the application.
  - Performance dashboards for the dependencies. (we aren't testing them, but
    we should be able to tell if they are reaching capacity).
  - Performance dashboards for relevant infrastructure.
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

