# Thoughts

Thoughts is a command line tool to store your thoughts that randomly pop up as you write code.
This prevents you from having to context switch away from the command line. It keeps you focused, so you can context switch only after you're finished with your deep work session.


Here's what happens in a world without Thoughts:
1. You're writing some code and troubleshooting a bug.
2. You run some command on the terminal to test the bug.
3. You have a random thought about that funny cat video you saw the other day.
4. You go on YouTube to search for that cat video.
5. You then come back to your code, losing context, and needing to ramp back up on the code. :(


Here's what happens with Thoughts:
1. Same as above.
2. Same as above.
3. Same as above.
4. `$ thoughts look up that funny cat video`
5. Continue solving the bug until you squash it.
6. Take a break and congratulate yourself for fixing it.
7. `$ thoughts l`
8. "Oh yeah, I need to watch that funny cat video!"
9. Win the day for maintaining your deep work session. :D


## Installation

This will create a binary and store it in /usr/local/bin

`$ ./install`


## Usage

`$ thoughts read blog post on caching`

`$ thoughts reach out to Jon about that video he sent`

`$ thoughts join the Go Discord channel`

`$ thoughts write a tool to prevent context switching`

```
$ thoughts l

read blog post on caching
reach out to Jon about that video he sent
join the Go Discord channel
write a tool to prevent context switching
```

You can also list all thoughts from a specific date

```
$ thoughts l 2024-01-29
```

Or list all thoughts from all dates

```
$ thoughts l all
```
