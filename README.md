# goring
A simple consistent hashing ring implementation in Golang


#Usage

```go
  import(
    "fmt"
    "github.com/alash3al/goring"
  )
  
  func main(){
    // init
    ring := new ring.NewRing()
    
    // add node with its weight in the ring
    ring.Add("127.0.0.1:8000", 2)

    // add more nodes
    ring.Add("127.0.0.1:8001", 1).Add("127.0.0.1:8002", 5) // ...

    // find a node by anonoymous key
    node := ring.Get("usr:alash3al:name")

    // print it
    fmt.Println(node)

    // remove a node
    node.Remove("127.0.0.1:8002")
  }
```
