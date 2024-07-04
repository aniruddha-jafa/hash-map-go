```sh
$ go build
# linear probing 
$ ./hash-map-go -n=8 -impl=linear < tale.txt > out/tale_linear.csv
# chaining
$ ./hash-map-go -n=8 -impl=chain < tale.txt > out/tale_chain.csv
```