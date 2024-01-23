package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	"honnef.co/go/gotraceui/color"
)

var (
	errExitAfterParsing = errors.New("we were instructed to exit after parsing")
	errExitAfterLoading = errors.New("we were instructed to exit after loading")
)

type debugGraph struct {
	title           string
	width           time.Duration
	background      color.Oklch
	fixedZero       bool
	stickyLastValue bool

	mu     sync.Mutex
	values []struct {
		when time.Time
		val  float64
	}
}

type DebugWindow struct {
	cvStart           debugGraph
	cvEnd             debugGraph
	cvY               debugGraph
	cvPxPerNs         debugGraph
	animationProgress debugGraph
	animationRatio    debugGraph
	frametimes        debugGraph
}

func writeMemprofile(s string) {
	f, err := os.Create(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, "couldn't write memory profile:", err)
		return
	}
	defer f.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Fprintln(os.Stderr, "couldn't write memory profile:", err)
	}
}

func assert(b bool, msg string) {
	if !b {
		panic(fmt.Sprintf("failed assertion: %s", msg))
	}
}
