package main

import "log"

func logApp(msg string) {
	log.Printf("[app] %s\n", msg)
}

func logPod(podID, component, msg string) {
	log.Printf("[%s][%s] %s\n", podID, component, msg)
}
