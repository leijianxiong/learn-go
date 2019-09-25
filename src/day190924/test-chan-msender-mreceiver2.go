package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
	"strconv"
)

//关闭通道原则: 1 不能在接受者方关闭数据通道, 2 不能在多个发送者中关闭数据通道
//m个发送者 m个接收者: 使用一个第三方

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100
	const NumReceivers = 10
	const NumSenders = 10

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown below.
	// Its reveivers are all senders and receivers of dataCh.
	toStop := make(chan string, 1)
	// the channel toStop is used to notify the moderator
	// to close the additional signal channel (stopCh).
	// Its senders are any senders and receivers of dataCh.
	// Its reveiver is the moderator goroutine shown below.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <- toStop // part of the trick used to notify the moderator
		// to close the additional signal channel.
		//stopCh <- struct{}{}
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				log.Println("sender#" + id + " rand v=", value)
				if value == 0 {
					// here, a trick is used to notify the moderator
					// to close the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// the first select here is to try to exit the
				// goroutine as early as possible.
				select {
				case v, ok := <- stopCh:
					log.Println("sender#"+id+" return", v, ok)
					return
				default:
				}

				select {
				case dataCh <- value:
				default:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// same as senders, the first select here is to
				// try to exit the goroutine as early as possible.
				select {
				case v, ok := <- stopCh:
					log.Println("receiver#"+id+" return", v, ok)
					return
				default:
				}

				select {
				case value := <-dataCh:
					log.Println("receiver#" + id + " v=", value)
					if value == MaxRandomNumber-1 {
						// the same trick is used to notify the moderator
						// to close the additional signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				default:
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}