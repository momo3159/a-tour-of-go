package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	v:= make([][]uint8, dy)
	
	for y := range v {
	  row := make([]uint8, dx)
	  for x := range row {
	    row[x] = uint8(x^y)
	  }
	  v[y] = row
	}
	
	return v
} 

func main() {
	pic.Show(Pic)
}

