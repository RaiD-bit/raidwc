
# raidwc 

this a tool similar to wc in go.


I am trying to learn go by implementing the challenges mentioned at [coding challenges](https://codingchallenges.fyi/challenges).

This the first challenge in that list: [challenge-wc](https://codingchallenges.fyi/challenges/challenge-wc) 


## Installation

To run this locally.

Step 1: clone this project.

Step 2: navigate to the root folder

```bash
    go build .
```
you should see an executalbe by the name of raidwc. 

You can now use this raidwc executalbe to run usual wc cmds.


## Usage/Examples
The challenge has 5 tasks. 
As part of this challenge we support the following flags 

```
Usage of ./raidwc:
  -c    get byte count
  -l    get line count
  -m    get char count
  -w    get word count
```

Also the final task is being able to read from standard input if no filename is specified. Something like this: 


```bash
❯ cat test2.txt | ./raidwc -w -l
    10   2
```

