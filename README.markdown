tmachine is a Go implementation of the Google's Turing machine doodle.

It is programmable, but the demo program implements the cool Fibonacci-sequence generator that HN user cduan so nicely reverse-engineered [here](http://sbf5.com/~cduan/technical/turing/).

There was an amusing opportunity to write literal TM instructions here by bitwise-ORing 'opcodes', so Branch|Blank|Down means move down in the instruction board if the symbol under head is blank.

To run tmachine: `go run tmachine.go`. 

Sample compressed output (unfortunately, Chrome doesn't do a gread job, visually, with the combining underscore).

```
go run tmachine.go 
 0  2 [ 1͟                              ] 82  1
 0  2 [ 1͟                              ] 82  1
17  3 [ 1 ͟                             ] 68 ?_↓
32  4 [ 1 1͟                            ] 81  0
17  5 [ 11 0͟                           ] 68 ?_↓
32  6 [ 11 01͟                          ] 81  0
 9  7 [ 111 10͟                         ] 68 ?_↓
17  8 [ 1110 01͟                        ] 68 ?_↓
32  9 [ 1110 011͟                       ] 81  0
 9 10 [ 11101 110͟                      ] 68 ?_↓
17 11 [ 111010 101͟                     ] 68 ?_↓
32 12 [ 111010 1011͟                    ] 81  0
17 13 [ 1110101 0110͟                   ] 68 ?_↓
32 14 [ 1110101 01101͟                  ] 81  0
 9 15 [ 11101011 11010͟                 ] 68 ?_↓
17 16 [ 111010110 10101͟                ] 68 ?_↓
32 17 [ 111010110 101011͟               ] 81  0
17 18 [ 1110101101 010110͟              ] 68 ?_↓
32 19 [ 1110101101 0101101͟             ] 81  0
 9 20 [ 11101011011 1011010͟            ] 68 ?_↓
17 21 [ 111010110110 0110101͟           ] 68 ?_↓
32 22 [ 111010110110 01101011͟          ] 81  0
 9 23 [ 1110101101101 11010110͟         ] 68 ?_↓
17 24 [ 11101011011010 10101101͟        ] 68 ?_↓
32 25 [ 11101011011010 101011011͟       ] 81  0
```
