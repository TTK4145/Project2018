# Submission

## How to submit

1. Push all your commits into your designated github repo (the one that was created for you by using github classroom).
2. Add a README.md (or add to your README.md) where you describe what kind of libraries/code you've used that you're not the author of. Feel free to add other relevant info to the README.md as well.
3. Anonymize your code/readme to the extent that there are no obvious way to identify you from your code (remove group number and names).
4. Look through the code and make sure you're happy using it for the [code review](EVALUATION.md#code-review)
5. Create an executable binary and test it until you're happy using it for the [acceptance test](EVALUATION.md#completion-test).
6. Create a new [release](https://help.github.com/articles/creating-releases/) for your project repo.
7. Upload the binary to your newly created release.
8. Download the .zip in your release and upload it to blackboard **together with a link to the release**.

## Binary file

Before adding a binary file was mandatory people would sometime lose points due to undefined behaviour compiling differently on different compilers. A lot of people have administrative access to the computers on the real time lab and you cannot be sure what has been done to them. If you're having trouble creating executable binaries you should talk get a student assistant to help you. If you can take other measures to assure that you're running the code you tested (e.g. containers) this is also fine.

## Readme file.

Remember that other people will read your code. A readme file can be a good way to guide readers (including future self) to the relevant parts. You can still assume that the readers will know conventions for the programming language you're using.

## Tips and Tricks
 - To test packet loss time-efficiently an unrealisticly high packet loss probability is used. This will cause loss of several packets in a row. If you rely on heartbeat timeouts or similar, you're adviced to require 10 lost packets before timeout occurs. (timeout time >= 10x heartbeat period
