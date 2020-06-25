package main

import (
  "fmt"
  "os/exec"
  "time"
  "net/http"
)

func main() {
  resp, err := http.Get("http://127.0.0.1:5000/stop/gp_App2")
  fmt.Printf("being busy...")
  time.Sleep(5 * time.Second)
  fmt.Printf("..bye")
//   for {
//     fmt.Println("Hello from firecracker (run by unik from solo.io)")
//     out, _ := exec.Command("uname", "-a").CombinedOutput()
//     fmt.Printf("OS Version: %s\n", string(out))
//     time.Sleep(10 * time.Second)
//   }
}
