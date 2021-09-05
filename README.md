# GO-AUC

Using GO to implement Riemann Sums with trapezoids to approximate areas under a curve.

## How to use

Assuming you have a local installation of GO, please follow the instructions:

```bash
$ git clone https://github.com/gartician/GO-AUC.git

$ cd GO-AUC

# print usage
$ ./auc -h

Usage of ./auc:
  -i string
        2-column tab-separated input file. column 1 is x, and column 2 is y.
  -v    verbose mode will track AUC along the coordinates.
  
# use the provided input files
$ ./auc -i real_example.txt
Input file: real_example.txt
Verbose mode: false
total auc: 0.762655
```

The input file is a 2-column tsv file with an x and y column that looks like this:

```
x       y
0.00    0.00
0.00    0.01
0.00    0.02
0.00    0.03
0.00    0.04
0.00    0.05
0.00    0.06
0.00    0.07
0.00    0.08
```

