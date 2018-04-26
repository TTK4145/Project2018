# Technical Specification

### 1. Basic Definitions
##### 1.1. An order is *placed* when an order button is pressed on the elevator panel inside an elevator, or the elevator panel outside the elevator.
##### 1.2. A *cab call* (or *cab order*) is any order that is placed from a panel inside an elevator.
##### 1.3. A *hall call* (or *hall order*) is any order that is placed from a panel outside an elevator.
##### 1.4. The *participating elevators* is a set of elevators that remain constant over a test case. This means that even though an elevator is temporarily disabled or non functioning it is still a participating elevator. Furthermore, any elevator that may influence the outcome of the test is a participating elevator.
##### 1.5. The number of participating elevators is refered to as `n`.
##### 1.6. Everyone of the `n` participating elevators have the exact same floors.
##### 1.7. The number of floors any of elevator see is refered to as `m`.
##### 1.8. An order is *accepted* when a participating elevator changes the light corresponding to the order from off to on. If an order is already accepted and this happens, it changes nothing.
##### 1.9. An order is *completed* when a participating elevator open it's doors at the corresponding floor and the corresponding light is in (or is turned to) off state for this particular elevator. We say that this particular elevator has served the order.

### 2. Elevator Behaviour
##### 2.1. This specification describes software for controlling `n` elevators working in parallel across `m` floors.
##### 2.2. When a hall order is accepted, it must be served within reasonable time. Such order can be served by any of the `n` participating elevators.
##### 2.3 When a cab order is accepted, it must be served within reasonable time. The only elevator able to serve such order is the elevator corresponding to the panel where the order was placed.
##### 2.4. It is not reasonable to expect a cab order to be completed as long as the elevator able to serve it is non-functional. However, it must be served after function is restored.
##### 2.5 Multiple elevators must be more efficient than one, in cases where this is reasonable to expect.
##### 2.6 The elevator system must avoid doing unnecessary actions. 
##### 2.7 The elevators should have a sense of direction, more specifically, the elevators must avoid serving hall-up and hall-down orders on the same floor at the same time.
##### 2.8. A placed hall order should be ignored if there is no way to serve the order (as described in [2.2](SPECIFICATION.md#2-elevator-behaviour)) in the pressence of faults (as described in [3. Fault Tolerance](SPECIFICATION.md#3-fault-tolerance))
##### 2.9 A placed hall order can be intermittently disregarded as long as multiple placement attempts have a low probability of failing (it is allowed to disregard an order due to 3 udp packets being dropped in a row, or due to one of the elevators is busy initializing, etc).
##### 2.10. The door must never be open while moving.
##### 2.11. The door must only be open when the elevator is at a floor.
##### 2.12. When the door is opened, it should remain open for at least 2 seconds.
##### 2.13. When no input is given, the door should close within 5 seconds.
##### 2.14. The panel lights must be synchronized between elevators whenever it is reasonable to do so.

### 3. Fault Tolerance
##### 3.1. Assume that errors will happen (both expected and unexpected).
##### 3.2. Assume that errors will be resolved, but not necessarily within a definite amount of time.
##### 3.3. Assume that multiple errors will not happen simultaneously. More specifically, assume that you have sufficient time to do necessary preparation before the next error will occur as long as you detect the first error and recover from it within reasonable time.
##### 3.4. Even though errors will not happen simultaneously they can be "active" simultaneously. For instance when the second error happens before the first one is resolved.
##### 3.5. Loss of packets in UDP is not regarded as an error. UDP is in nature unreliable and can rightfully drop arbitrary packets. As a consequence of this, you may encounter an error at the same time as UDP packets are dropped.
##### 3.6. A computer that loses power (and therefore loses network) is regarded as a single failure.
##### 3.7. The (fictional) computer system contains redundant disks with replacements readily available. You may write to file, disregarding the chances of losing anything.

### 4. Unspecified Behaviour
##### 4.1. What happens after the stop button is pressed is intentionally unspecified.
##### 4.2. What happens after the obstruction button switch is turned on is intentionally unspecified.

### 5. Practical Implementation
##### 5.1 You must be able to create a single executable that runs your elevator software. (This rule should not limit your choice of tools. If you're writing in a language that requires an interpreter and/or have problems generating executables, talk to the student assistants or the person administrating this project to get help).
##### 5.2 Your program must not depend on any configuration files.
##### 5.3 Your program must not depend on any command line arguments except those listed below. You are free to ignore any or all of these, but your program must work in the presence of the listed ones.
 - `--id <id>` where all participating elevators will be started with an unique `<id>` in the range `[0, 255]`.


