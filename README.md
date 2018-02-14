# Mad aliens invasion simulation
## Usage
You need to have [go installed](https://golang.org/doc/install), and GOPATH set
```
$ go build
$ ./invasion -alien 10
```
The alien flag set the number of alien you want for the simulation (For this data set it has to be between 2 and 15).
## Assumptions
- Three or four aliens can go to the same city at the same time. In this case they all die.
- Only 2 to 15 aliens can be created for this dataset.
