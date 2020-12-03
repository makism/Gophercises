
# v2

* Use channels properly for timer and input (and labels!).
* Use the flag module for parsing command line arguments.
* Rename variables to comfort go-style naming.

After this nice revamping, I stumbled upon the following issue:
> What happens if the user answers all the questions _before_ the timer elapses.

Well, the answer is simple: wait.

So, I had to refactor the following:
- move the for-loop outside the goroutine to force the iteration to stop if needed. 

# v1
Bare-basic implementation
