package main

import (
  "fmt"
  "os/exec"
  "time"
)

func main() {
  for {
    fmt.Println("Hello from firecracker (run by unik from solo.io)")
    out, _ := exec.Command("uname", "-a").CombinedOutput()
    fmt.Printf("OS Version: %s\n", string(out))
    time.Sleep(10 * time.Second)
  }
}
