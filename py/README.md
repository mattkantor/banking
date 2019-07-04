## FIZZBUZZ for Banking

### Requirements

- Python 3.5 +
- pip3
- MacOS or Unix. Sorry.

### Installation

`$ make init`

`$ source env/bin/activate`

`$ make test`

`$ make run`

`$ cat ../results_py.txt`


## Some notes on style

- Why not have a controller? The answer is to hide the guts of things and keep it simple.  We have a singleton which has one command - deposit.  
You really dont need anything else.  It's all hidden.
- as opposed to doing anything persistent, I wanted a way to quickly and easily match the logic to the rules of the loading. Given this is an exercise, in reality we mah not want to group by mondays and date offsets, especially if its the only optiion.  I'm sort of thinking we have multiple ways of storing our data, this is good way for this exercise.
- In go, it would be a bit more elegant I think.
- Obviously this isnt optimized for much, there's no persistence or anything.

