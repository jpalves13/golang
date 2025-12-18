package main

import "log"

func logApp(msg string) {
	log.Printf("[app] %s\n", msg)
}

func logPod(pod, comp, msg string) {
	log.Printf("[%s][%s] %s\n", pod, comp, msg)
}
