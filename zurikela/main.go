package main

// https://www.hackerrank.com/challenges/zurikela/problem

import (
	"fmt"
	"os"
  "bufio"
  "strconv"
)

func main() {
	file, err := os.Open("sampleinput")

	if err != nil {
		fmt.Println(err)
		return
	}

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanWords)
  
  q := scanInt(scanner)

  for i :=0; i<q; i++ {
    readOp(scanner)
  }
  
	defer file.Close()
}

func readOp(scanner *bufio.Scanner){
  scanner.Scan()
  op := scanner.Text()

  switch op {
    case "A":
      opA(scanInt(scanner))
    case "B":
      opB(scanInt(scanner), scanInt(scanner))
    case "C":
      opC(scanInt(scanner))
    default:
      panic("Invalid op: " + op)
  }
}

func opA(x int) {
  fmt.Printf("A %v \n", x)
}

func opB(x1, x2 int) {
  fmt.Printf("B %v %v \n", x1, x2)
}

func opC(x int) {
  fmt.Printf("C %v \n", x)
}

func scanInt(scanner *bufio.Scanner) int {
  scanner.Scan()
  txt := scanner.Text()

  val, err := strconv.Atoi(txt)
  if err != nil {
    panic("Int scan failed on Atoi !")
  }

  return val
}
