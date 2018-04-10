# Evaluation

In TTK4145 25% of the final grade is based on the elevator project. The project also give a maximum score of 25 points. This does not mean that the project score necesarily translates directly into the the 25% of your final grade. Some of the evaluations might not be practical to scale after [NTNUs grading scale](https://innsida.ntnu.no/wiki/-/wiki/English/Grading+scale+using+percentage+points), and at others we might fail to scale the difficulty level properly. The only thing that can be thought of as given is how much value each of the evaluations have relatively to each other. This is given in the table below.

| Design review | Code review | Completion | Total |
| ------------- | ----------- | ---------- | ----- |
| 8             | 9           | 8          | 25    |

## Design Review
### Motivation
 - If you do good on the design review you might get a better grade in TTK4145.
 - This is a good opportunity to get direct feedback from the teaching staff on a large part of your project.
 - For you to be able to present your design, you will need to think about it. The process of thinking about your design might reveal weaknesses.

### Expectations
 - You are expected to deliver a handout on blackboard and bring printed copies to the presentation. The handout should ideally be two to three pages and no more than four pages. You are not expected to include text, and you are expected to include figures.
 - You are expected to give a short talk (5-10min) about the design if asked to.
 - You are expected to answer questions about your design.

### Specification
You will need to consult the [specification](https://github.com/TTK4145/Project#technical-specification) for details, but the key part is to make sure that no accepted orders will be lost.

### Evaluation Criteria
#### Fault Tolerance
 - Your system should handle any single failure at any time and still operate after specifications.
 - You should know what errors that are likely to occur, and how they are handled in your system. 

#### Code Quality
 - You should partition the system into modules.
 - You should minimize the couplings/dependencies between the modules, and make these couplings explicit.
 - You should give the modules names that summarize their responsibilities.
 - You should ensure that the module interfaces are consistent with the module responsibilities. 

#### Design of RT systems
 - You should be able to explain what strengths/weaknesses that is relevant for your chosen tools (programming language etc). 
 - You should be able to explain how the abstractions you have chosen are represent-able with your chosen tools.

#### Presentation
 - You should be able to convey your solution clearly.

### Disclaimers
You should not think of the design document as a complete specification but rather a draft. You should use feedback from the design review and your gained insight into the task at hand to improve/change your design as you go along.

## Code review
### Evaluation Criteria
#### Personal skills
 - You should be able to evaluate code quality based on the [code evaluation criterias](https://github.com/TTK4145/Project/blob/master/EVALUATION.md#Code-evaluation)
 
#### Code evaluation
##### Top level
 - The entry point (or similar) should document what components/modules the system consists of.
 - It should be clear how different components/modules communicate and depend on each other, and how they depend on external components.
 - Naming should be consistent over the whole program.
 - Language features should be used in a good way.
 - Comments should give information that is difficult to express in code.
     - Comments should not be repetition of code.
     - Comments should generally express intent on a different abstraction level than code.
     
##### Modules
 - Modules should minimize accessibility to their internals.
     - Programmatically enforce this where possible (Erlang, Rust, etc)
     - Enforce this through conventions if there are no language features to enable this (Python, etc)
 - Modules should hide their implementation details.
 - Module interfaces should represent a consistent abstraction.
 - The module names should fit well with the module interface.
 - Module interfaces should be simple yet complete.
     - Avoid interfaces that are hard to use correctly.
     
##### Routines
 - A routine should have a concrete purpose (do one thing well).
     - The word "and" in a routine name usually means that something is "wrong".
 - The routine name should indicate its purpose.
 - Routines should be simple enough to veirfy from inspection that it fulfills its purpose without doing anything more.
 
## Completion Test
The completion test aims to test wether your system adheres to specification in a time efficient way. Some things to take note of in the different parts of the completion test follows.

#### Normal operation
When either one, two or three elevators are initialized they should adher to specification.

Note: *You will be allowed to reinitialized your elevator system between testing with one, two and three elevators*

#### Fault tolerance
The elevator system should adher to specification when faults are introduced.

#### Packet loss testing
Your elevator system must adher to specification even with simulated packet loss on your network adapter (not on localhost).

Note: *To test packet loss time-efficiently an unrealisticly high packet loss probability is used. This will cause loss of several packets in a row. If you rely on heartbeat timeouts or similar, you're adviced to require 10 lost packets before timeout occurs. (timeout time >= 10x heartbeat period)*


