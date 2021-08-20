# Golang-SeqLock
A basice implementation of sequence lock in golang.
# License
This project is licensed under the terms of the simplified BSD license.
# How to use
```
go get -u github.com/MUSQQQ/Golang-SeqLock
```
I'm sorry for not the greatest example. I plan to prepare a better one inside this repo or make a new repo for showing the usage of this seqlock.
```go

// func that shows the process of reading data
func ReadingData(seq *seqlock.SeqLock, wg *sync.WaitGroup) {
	tmp := uint32(0)
	for {
		tmp = seq.RdRead()
		/*

			reading data

		*/
		if !seq.RdAgain(tmp) {
			break
		}
	}
	defer wg.Done()
}

// func that shows the process of writing data
func WritingData(seq *seqlock.SeqLock, wg *sync.WaitGroup) {
	seq.WrLock()
	/*

		writing data

	*/
	seq.WrUnlock()
	defer wg.Done()
}
```

# Why
It was a mini project for my Multithreaded Programming classes