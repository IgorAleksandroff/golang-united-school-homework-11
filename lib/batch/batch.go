package batch

import (
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
	jobs := make(chan int64, n)
	results := make(chan user, n)
	users := make([]user, 0, n)

	for w := int64(1); w <= pool; w++ {
		go worker(jobs, results)
	}

	for j := int64(0); j < n; j++ {
		jobs <- j
	}
	close(jobs)
	for a := int64(1); a <= n; a++ {
		detestedUser := <-results
		users = append(users, detestedUser)
	}
	return users
}

func worker(jobs <-chan int64, results chan<- user) {
	for j := range jobs {
		results <- getOne(j)
	}
}
