# Elevator Project.

## Introduction
An elevators primary purpose is to transport people and items between floors. A sensible metric of performance could be something indicating how well and fast this primary task is done. This is not something we care about in this project. The elevators must avoid starvation and "silly actions", other than that, not much attention is given to their performance as elevators. This project is really about faults, how they are bound to happen, and how we can keep them from ruining our day. 

On the good days, your elevator system is bound to be under-appreciated. The corporate procurement officers is going to favor cheaper or faster alternatives. But when the day comes that an especially large earthquake simultaneously causes the largest tsunami the world has ever seen, cracks the concrete in a biological warfare lab and cause a nuclear reactor meltdown, your elevator system is going to be all worth it. Most of your server rooms will be flooded. Radioactive zombies will be chewing on your network cables. And most of your disks will be wiped out by the earthquake itself. In this chaotic post apocalyptic world, the only constant will be that your elevator system will remain in working order, and every order that is taken, will also be served.

## Specification
#### Technical Specification
The "technical specification" of the elevator system is simple but most likely ambiguous and incomplete. It should only be regarded as a subset of the "complete specification". Most of you will still be able to create a system adhering to specification just by reading the "technical specification". The technical specification can be found [here](SPECIFICATION.md)

The "complete specification" does not exist in written form. It is a metaphysical entity that can only be channeled into through the distorted mind of a teaching assistant. There are two ways to make sure you are interpreting the technical specification in a way that is compatible with the complete specification. The practical yet informal is to ask a student assistant. The student assistants are great! They have experience with interpreting the technical specification and have possibly answered the question you want an answer for a few time already.

#### FAQ
Before revealing the formal way to resolve specification ambiguities there is one more part of the specification that needs to be discussed. Namely the FAQ (TODO: insert link). The FAQ contains clarifications in regards to how the specification should be interpreted. There should never be conflicting answers between the FAQ and the technical specification, or between different FAQs.

If you're not satisified with the answer of a student assistant or believes that the issue at hand requires a formal explanation you should start reading the "old" FAQs. If no answer can be found you may fill out an issue following the [guidelines](CONTRIBUTING.md). If found warrented your issue will result in a new FAQ or change in the techincal specification, and you will at least get a response from the teaching staff.

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

#### Useful resources
We encourage submissions to this list! Tutorials, libraries, articles, blog posts, talks, videos...
 - [Python](http://python.org/)
   - [Official tutorial (2.7)](http://docs.python.org/2.7/tutorial/)
   - [Python for Programmers](https://wiki.python.org/moin/BeginnersGuide/Programmers) (Several websites/books/tutorials)
   - [Advanced Python Programming](http://www.slideshare.net/vishnukraj/advanced-python-programming)
   - [Socket Programming HOWTO](http://docs.python.org/2/howto/sockets.html)
 - C
   - [Amended C99 standard](http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf) (pdf)
   - [GNU C library](http://www.gnu.org/software/libc/manual/html_node/)
   - [POSIX '97 standard headers index](http://pubs.opengroup.org/onlinepubs/7990989775/headix.html)
   - [POSIX.1-2008 standard](http://pubs.opengroup.org/onlinepubs/9699919799/) (Unnavigable terrible website)
   - [Beej's network tutorial](http://beej.us/guide/bgnet/)
   - [Deep C](http://www.slideshare.net/olvemaudal/deep-c)
 - [Go](http://golang.org/)
   - [Official tour](http://tour.golang.org/)
   - [Go by Example](https://gobyexample.com/)
   - [Learning Go](http://www.miek.nl/projects/learninggo/)
   - From [the wiki](http://code.google.com/p/go-wiki/): [Articles](https://code.google.com/p/go-wiki/wiki/Articles), [Talks](https://code.google.com/p/go-wiki/wiki/GoTalks)
   - [Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw) (video): transforming problems into the for-select-loop form
 - [D](http://dlang.org/)
   - [The book](http://www.amazon.com/exec/obidos/ASIN/0321635361/) by Andrei Alexandrescu ([Chapter 1](http://www.informit.com/articles/article.aspx?p=1381876), [Chapter 13](http://www.informit.com/articles/article.aspx?p=1609144))
   - [Programming in D](http://ddili.org/ders/d.en/)
   - [Pragmatic D Tutorial](http://qznc.github.io/d-tut/)
   - [DConf talks](http://www.youtube.com/channel/UCzYzlIaxNosNLAueoQaQYXw/videos)
   - [Vibe.d](http://vibed.org/)
 - [Erlang](http://www.erlang.org/)
   - [Learn you some Erlang for great good!](http://learnyousomeerlang.com/content)
   - [Erlang: The Movie](http://www.youtube.com/watch?v=uKfKtXYLG78), [Erlang: The Movie II: The sequel](http://www.youtube.com/watch?v=rRbY3TMUcgQ)
 - [Rust](http://www.rust-lang.org/)
 - Java
   - [The Java Tutorials](http://docs.oracle.com/javase/tutorial/index.html)
   - [Java 8 API spec](http://docs.oracle.com/javase/8/docs/api/)
 - [Scala](http://scala-lang.org/)
   - [Learn](http://scala-lang.org/documentation/)
 - [C#](https://msdn.microsoft.com/en-us/library/kx37x362.aspx?f=255&MSPPError=-2147217396)
   - [C# 6.0 and the .NET 4.6 Framework by Andrew Troelsen (free pdf-version for NTNU students)](http://link.springer.com/book/10.1007/978-1-4842-1332-2)
   - [Mono (.NET on Linux)](http://www.mono-project.com/docs/)
   - [Introduction to Socket Programming with C#](http://www.codeproject.com/Articles/10649/An-Introduction-to-Socket-Programming-in-NET-using)
   - Importing native libraries: [general](http://www.codeproject.com/Articles/403285/P-Invoke-Tutorial-Basics-Part) and [for Linux](http://www.mono-project.com/docs/advanced/pinvoke/)

<!-- -->
 
 - Design and code quality
   - [The State of Sock Tubes](http://james-iry.blogspot.no/2009/04/state-of-sock-tubes.html): How "state" is pervasive even in message-passing- and functional languages
   - [Defactoring](http://raganwald.com/2013/10/08/defactoring.html): Removing flexibility to better express intent
   - [The Future of Programming](http://vimeo.com/71278954) (video): A presentation on what programming may look like 40 years from now... as if it was presented 40 years ago.
   - [Railway Oriented Programming](http://www.slideshare.net/ScottWlaschin/railway-oriented-programming): A functional approach to error handling
   - [Practical Unit Testing](https://www.youtube.com/watch?v=i_oA5ZWLhQc) (video): "Readable, Maintainable, and Trustworthy"
   - [Core Principles and Practices for Creating Lightweight Design](https://www.youtube.com/watch?v=3G-LO9T3D1M&t=4h31m25s) (video)
    - [Origins and pitfalls of the recursive mutex](http://zaval.org/resources/library/butenhof1.html). (TL;DR: Recursive mutexes are usually bad, because if you need one you're holding a lock for too long)

## Evaluation
The project give a maximum score of 25 points. Follow [this](https://github.com/TTK4145/Project/blob/master/EVALUATION.md#evaluation) link to see a detailed description of the evalation.

## Tips & Tricks
- For a process to access the io card (elevator hw) on the real time lab the user running the process must be in the iocard group. To add user student to the iocard group run `sudo usermod -a -G iocard student`.

## Contact
- If you find an error or something missing please post an issue or a pull request.
- Updated contact information can be found [here](https://www.ntnu.no/studier/emner/TTK4145).
