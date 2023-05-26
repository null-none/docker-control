package main

import (
	Controller "./controllers"
)

func main() {
	Controller.RunContainerBackground("hello-world")
	Controller.ListContainers()
	Controller.StopRunningContainers()
	Controller.ListImages()
	Controller.PullImage("hello-world")
}
