# Golang-SeqLock
A basice implementation of sequence lock in golang.
# License
This project is licensed under the terms of the simplified BSD license.
# How to use
```
go get -u github.com/MUSQQQ/Golang-SeqLock
```
I'm sorry for the very long example. I plan to prepare a better one inside this repo or make a new repo for showing the usage of this seqlock.
```go
var r = rand.New(rand.NewSource(1))

// func that shows the process of reading data
func ReadingData(seq *seqlock.SeqLock, wg *sync.WaitGroup) {
	tmp := int32(0)
	for {
		time.Sleep(time.Duration(r.Int31()) * 1)
		tmp = seq.RdRead()
		/*

			process of reading data

		*/
		if !seq.RdAgain(tmp) {
			fmt.Printf("counter after succesfully reading data: %d\n", tmp)
			break
		}
	}
	defer wg.Done()
}

// func that shows the process of writing data
func WritingData(seq *seqlock.SeqLock, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(r.Int31()) * 1)
	seq.WrLock()
	/*

		process of writing data

	*/
	fmt.Printf("writing in progress, counter is odd: %d\n", seq.Counter)
	seq.WrUnlock()
	defer wg.Done()
}
```

# Why
It was a mini project for my Multithreaded Programming classes