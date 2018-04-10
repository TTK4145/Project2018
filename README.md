# Elevator Project.

## Introduction
An elevators primary purpose is to transport people and items between floors. A sensible metric of performance could be something indicating how well and fast this primary task is done. This is not something we care about in this project. The elevators must avoid starvation and "silly actions", other than that, not much attention is given to their performance as elevators. This project is really about faults, how they are bound to happen, and how we can keep them from ruining our day. 

On the good days, your elevator system is bound to be under-appreciated. The corporate procurement officers is going to favor cheaper or faster alternatives. But when the day comes that an especially large earthquake simultaneously causes the largest tsunami the world has ever seen, cracks the concrete in a biological warfare lab and cause a nuclear reactor meltdown, your elevator system is going to be all worth it. Most of your server rooms will be flooded. Radioactive zombies will be chewing on your network cables. And most of your disks will be wiped out by the earthquake itself. In this chaotic post apocalyptic world, the only constant will be that your elevator system will remain in working order, and every order that is taken, will also be served.

## Specification
#### Technical Specification
The "technical specification" of the elevator system is simple but most likely ambiguous and incomplete. It should only be regarded as a subset of the "complete specification". Most of you will still be able to create a system adhering to specification just by reading the "technical specification". The technical specification can be found [here](SPECIFICATION.md)

The "complete specification" does not exist in written form. It is a metaphysical entity that can only be channeled into through the distorted mind of a teaching assistant. There are two ways to make sure you are interpreting the technical specification in a way that is compatible with the complete specification. The practical yet informal is to ask a student assistant. The student assistants are great! They have experience with interpreting the technical specification and have possibly answered the question you want an answer for a few time already.

#### FAQ
Before revealing the formal way to resolve specification ambiguities there is one more part of the specification that needs to be discussed. Namely the [FAQ](FAQ/README.md). The FAQ contains clarifications in regards to how the specification should be interpreted. There should never be conflicting answers between the FAQ and the technical specification, or between different FAQs.

If you are not satisfied with the answer of a student assistant, or believe that the issue at hand requires a formal explanation you should start with reading the existing FAQ answers. If no answer can be found you may fill out an issue following the guidelines [guidelines](CONTRIBUTING.md). Your issue will result in a new FAQ entry or change in the technical specification, and you will at least get a response from the teaching staff.

#### A word of comfort
If this seem daunting to you, fear not. You can choose to not care about any of the formalities surrounding the technical specification and FAQ process. If you instead use your best intuition on how elevators should work and ask as many questions as you can come up with to the student assistants and your fellow students, you will come up with something that is identical to the complete specification, and you too will be able to channel it through your very own (possibly distorted) mind.

## Programming Language
We do not impose any constraints on which programming languages you can use. As long as it is possible to use it for the task at hand, you are allowed to do so. You are responsible of keeping track of any tools you might need for making your project work (compilers, dependencies, build systems, etc) and the versioning of these, but we know these things can be hard and you should not hesitate to ask for help. 

#### State of support for some popular languages

| Languages | Student assistant support | Driver support                  | Network support                       |
|-----------|---------------------------|---------------------------------|---------------------------------------|
| Go        | Good+                     | ?                               | UDP based network module              |
| C         | Good                      | Git submodule                   | UDP/TCP based network module          |
| C++       | Good                      | Git submodule                   | The C module                          |
| Python    | Good-                     | ?                               | Only native support                   |
| Erlang    | Decent+                   | ?                               | Erlang is distributed by nature :tada:|
| Rust      | Decent+                   | ?                               | UDP based network module              |
| D         | Decent? (uncertain)       | ?                               | UDP based network module              |
| Ada       | Some                      | ?                               | Only native support                   |
| Java      | Not much                  | ?                               | Only native support                   |
| C#        | Unknown                   | ?                               | Only native support                   |

## Submission
Follow [this](SUBMISSION.md) link to find out how you submit the project.

## Evaluation
The project give a maximum score of 25 points. Follow [this](https://github.com/TTK4145/Project/blob/master/EVALUATION.md#evaluation) link to see a detailed description of the evalation.

## Tips & Tricks
- For a process to access the io card (elevator hw) on the real time lab the user running the process must be in the iocard group. To add user student to the iocard group run `sudo usermod -a -G iocard student`.

## Contact
- If you find an error or something missing please post an issue or a pull request.
- Updated contact information can be found [here](https://www.ntnu.no/studier/emner/TTK4145).
