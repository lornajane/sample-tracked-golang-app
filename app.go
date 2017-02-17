package main

import "fmt"
import "github.com/IBM-Bluemix/cf_deployment_tracker_client_go"

func main() {
	cf_deployment_tracker.Track()
	fmt.Printf("hello, world\n")
}
