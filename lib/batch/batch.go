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
	users := make([]user, n)
	jobs := make(chan int64, pool)
	wg := &sync.WaitGroup{}

	wg.Add(int(n))

	for i := int64(0); i < pool; i++ {
		go worker(jobs, users, wg)
	}

	for i := int64(0); i < n; i++ {
		jobs <- i
	}
	wg.Wait()
	return users
}

func worker(jobs <-chan int64, result []user, wg *sync.WaitGroup) {
	for j := range jobs {
		result[j] = getOne(j)
		wg.Done()
	}
}
