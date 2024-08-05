// cake.go

// Cake provides a simulation of a concurrent cake shop with numerous parameters.

// Use this command to run the benchmarks:
// $ go test -bench=. go-cake

package cake

import (
	"math/rand"
	"time"
)

type Shop struct {
	Cakes          int           // Number of cakes to bake.
	BakeTime       time.Duration // Time to bake one cake.
	BakeStdDev     time.Duration // Standard deviation of baking time.
	BakeBuf        int           // Buffer slots between baking and icing.
	NumIcers       int           // Number of cooks doing icing.
	IceTime        time.Duration // Time to ice one cake.
	IceStdDev      time.Duration // Standard deviation of icing time.
	IceBuf         int           // Buffer slots between icing and inscribing.
	InscribeTime   time.Duration // Time to inscribe one cake.
	InscribeStdDev time.Duration // Standard deviation of inscribing time.
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscriber(chan cake) {
	for i := 0; i < s.Cakes; i++ {
		work(s.InscribeTime, s.InscribeStdDev)
	}
}

// Work runs the simulation 'runs' times.
func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}

// work blocks the calling goroutine for a period of time that is normally
// distributed around d with a standard deviation of stddev.
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}
