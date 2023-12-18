package main

import (
	"fmt"
	"sync"

	"github.com/iverly/go-mcping/mcping"
)

var wg sync.WaitGroup

func routine(ips chan string) {
	for {
		ip := <-ips
		p := mcping.NewPinger()
		r, err := p.Ping(ip, 25565)
		if err != nil {
			continue
		} else {
			fmt.Printf("%d/%d\n%s\n%s\n%s\n", r.PlayerCount.Online, r.PlayerCount.Max, ip, r.Version, r.Motd)
		}
		wg.Done()
	}
}

func main() {
	routines := 500
	ips := make(chan string, routines)

	fmt.Println("Starting...")

	for i := 0; i <= routines; i++ {
		go routine(ips)
	}

	for i0 := 0; i0 <= 255; i0++ {
		if i0 == 10 {
			continue
		}
		for i1 := 0; i1 <= 255; i1++ {
			if (i0 == 172 && i1 >= 16 && i1 <= 31) || (i0 == 192 && i1 == 168) {
				continue
			}
			for i2 := 0; i2 <= 255; i2++ {
				for i3 := 0; i3 <= 255; i3++ {
					ips <- fmt.Sprintf("%d.%d.%d.%d", i0, i1, i2, i3)
					wg.Add(1)
				}
			}
		}
	}

	wg.Wait()
}
