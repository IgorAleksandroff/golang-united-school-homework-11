package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	users := make([]user, 0, n)
	jobs := make(chan int64, pool)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	wg.Add(int(n))
	for i := int64(0); i < pool; i++ {
		go func() {
			for j := range jobs {
				detectedUser := getOne(j)
				mu.Lock()
				users = append(users, detectedUser)
				mu.Unlock()
				wg.Done()
			}
		}()
	}

	for i := int64(0); i < n; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return users
}
