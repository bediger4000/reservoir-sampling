# Reservoir Sampling Filter

This code makes a Unix-style filter program that chooses some number of lines
from stdin at random, then prints those lines to stdout.

## Building and Running

```
$ go build reservoir.go
$ ./reservoir 7 < /usr/share/dict/words
escutcheons
insinuator's
bone
cc
identification's
segues
comer
$
```

Input is from stdin, output on stdout.
The mandatory command line argument is the number of lines to select from stdin.

This code based on [Algorithm R](https://en.wikipedia.org/wiki/Reservoir_sampling#Simple_algorithm)
from Wikipedia.

## Daily Coding Problem: Problem #911 [Medium] 

This problem was asked by Facebook.

Given a stream of elements too large to store in memory,
pick a random element from the stream with uniform probability.

### Analysis

Isn't this just reservoir sampling, with a reservoir size of 1?
