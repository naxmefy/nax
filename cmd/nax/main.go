package main

import (
	"context"
	"fmt"
	"github.com/naxmefy/nax/pkg/server"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	fmt.Println("hello nax!")
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = s.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("time to close the server")
	err = s.Shutdown(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
