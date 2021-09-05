package main

import (
	"fmt"
	"flag"
	"os"
	"strconv"
	df "github.com/go-gota/gota/dataframe"
	series "github.com/go-gota/gota/series"
)

func main() {

	// flag parse and checks ------------------------------------------------------------
	wordPtr := flag.String("i", "", "2-column tab-separated input file. column 1 is x, and column 2 is y.")
	verbPtr := flag.Bool("v", false, "verbose mode will track AUC along the coordinates.")
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("ERROR: Please supply at least 1 flag\n\n")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Input file:", *wordPtr)
	fmt.Println("Verbose mode:", *verbPtr)

	// import results -------------------------------------------------------------------
	ioContent, _ := os.Open(*wordPtr)
	df := df.ReadCSV(ioContent, df.WithDelimiter('\t'), df.HasHeader(true))

	// riemann sums ---------------------------------------------------------------------
	auc := riemann_sums(df, *verbPtr)
	fmt.Println("total auc:", auc)

	os.Exit(0)
}

func make_range(min, max int) []int {

	// Create a range of numbers from min to max.

	a := make([]int, max-min)
	for i := range a {
		a[i] = min + i
	}
	return(a)
}

func riemann_sums(input_df df.DataFrame, v bool) float64 {

	// input: df with x and y as headers
	// method: loop through coordinates and define current and next coordinates.
	// method: take trapezoidal area of the cross section and add it to auc.
	// output: area under the curve as float64

	// arrange df by increasing order and filter for coordinates >= 1.
	input_df = input_df.Arrange(df.Sort("x")).Filter(
		df.F{
			Colname:	"x",
			Comparator:	series.LessEq,
			Comparando:	1,
		},
	)

	// init important vars
	var auc float64 = 0
	var nrows int = input_df.Nrow()
	df_index := make_range(0, nrows - 1)
	x_col := []string{"x"}
	y_col := []string{"y"}

	for i := range df_index {

		// define x and f(x), then x_i and f(x_i). Then define coordinates as float64.
		x, _ := strconv.ParseFloat(  input_df.Subset(i).Select(x_col).Records()[1][0],  64  )
		y, _ := strconv.ParseFloat(  input_df.Subset(i).Select(y_col).Records()[1][0],  64  )
		x_i, _ := strconv.ParseFloat(  input_df.Subset(i+ 1).Select(x_col).Records()[1][0],  64  )
		y_i, _ := strconv.ParseFloat(  input_df.Subset(i+ 1).Select(y_col).Records()[1][0],  64  )

		// trapezoidal cross section + increment to auc
		tmp_area := trapezoid_area(y, y_i, x_i - x)
		auc += tmp_area

		// logging info
		if v == true {
			fmt.Println(  "x:", x, "y:", y, "x_i:", x_i, "y_i:", y_i, "tmp_area:", tmp_area, "auc:", auc  )
		}

	}

	return(auc)

}

func trapezoid_area(height1, height2, width float64) float64 {
	return(1.0/2.0 * (height1 + height2) * width)
}