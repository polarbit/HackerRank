package main

// https://www.hackerrank.com/challenges/zurikela/problem

import (
	"fmt"
	"os"
  "bufio"
  "strconv"
  "log"
  "flag"
  "io/ioutil"
)

var isLocal *bool = flag.Bool("local", false, "Is local development environment")

func main() {
  flag.Parse()
  if *isLocal == false {
    log.SetOutput(ioutil.Discard)
  }

  var scanner *bufio.Scanner

  if *isLocal {
    file, err := os.Open("sampleinput")
    if err != nil {
      fmt.Println(err)
      return
    }
    defer file.Close()
    scanner = bufio.NewScanner(file)
  } else {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
    scanner = bufio.NewScanner(reader)
  }
  
  scanner.Split(bufio.ScanWords)
  
  q := scanInt(scanner)

  for i :=0; i<q; i++ {
    readOp(scanner)
  }

  fmt.Println(calculateIndependents())
}

type set struct {
  k int
  nodes []*node
}

type node struct {
  set *set
  to []*node
}

var nextK int = 1
var sets = make(map[int]*set, 0)

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
  log.Println("A", x)

  createSet(x)

  dump()
}

func opB(x1, x2 int) {
  log.Println("B", x1, x2)

  s1 := sets[x1]
  s2 := sets[x2]

  for _, n1 := range s1.nodes {
    for _, n2 := range s2.nodes {
      n1.to = append(n1.to, n2)
      n2.to = append(n2.to, n1)
    }
  }

  dump()
}

func opC(x int) {
  log.Println("C", x)

  oldSet := sets[x]
  newSet := createSet(0)
 
  for _, n := range oldSet.nodes {
    moveNode(n, newSet)
  }

  dump()
}

func moveNode(node *node, newSet *set){
    if node.set == newSet {
        return
    }

    oldSet := node.set
    node.set = newSet
    newSet.nodes = append(newSet.nodes, node)

    for _, to := range node.to {
     if to.set == oldSet || to.set == newSet {
       continue
     }
      
      moveNode(to, newSet)
    }

    if sets[oldSet.k] != nil {
      delete(sets, oldSet.k)
    }
}

func createSet(nodes int) *set {
  newSet := set{k:nextK, nodes: make([]*node, nodes) }

  for i :=0; i<nodes; i++ {
    newSet.nodes[i] = &node{set: &newSet, to: make([]*node, 0)}
  }

  sets[nextK] = &newSet
  nextK++
  return &newSet
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

func calculateIndependents() int{
  
  result := 0
  for _, set := range sets {

    if len(set.nodes) < 2 {
      continue
    }
    
    for _, nodex := range set.nodes {
      
      edgeCount := 0
      for _, to := range nodex.to {
        
        if nodex.set == to.set {
          edgeCount += 1
        }

        if edgeCount >= 2 {
          break
        }
      }

      if edgeCount < 2 {
        result += 1
      }
    }
  }

  return result
}

func dump() {

  log.Println("====================")

  for _, s := range sets { 
    log.Println("Set:", s.k, "Count:", len(s.nodes), s.nodes)

    for _, n := range s.nodes {
      log.Println("Node To:", n.to)
    }
  }

  log.Println("Independents:", calculateIndependents())
}
