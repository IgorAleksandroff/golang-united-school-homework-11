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
	semaphore := make(chan struct{}, pool)
	wg := &sync.WaitGroup{}

	wg.Add(int(n))
	for i := int64(0); i < n; i++ {
		semaphore <- struct{}{}
		go func(index int64) {
			defer wg.Done()
			users[index] = getOne(index)
			<-semaphore
		}(i)
	}
	wg.Wait()
	return users
}
