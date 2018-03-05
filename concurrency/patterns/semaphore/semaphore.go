//信号量是一种普通的同步机制，可用来实现互斥锁，限制访问多个资源。如解决reader/writer问题。
//本例采用多个通道实现一个信号量，用信号量解决多个读取和一个写入的问题。
package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type (
	//定义semaphore类型为chan struct{}
	semaphore chan struct{}
)

type (
	readerWriter struct {
		name string
		//当写入时，会强制停止读取的goroutine
		write sync.WaitGroup
		//信号量，允许同时有多个读取goroutine
		readerControl semaphore
		//通知正在运行的的goroutine关闭
		shutdown chan struct{}
		//用于goroutines报告它们自己的关闭
		reportShutdown sync.WaitGroup
		//定义同时读取的goroutine总数。
		maxReads int
		//定义启动的goroutine个数。
		numReaders int
		//记录目前的read数目
		currentReads int32
	}
)

func main() {
	log.Println("程序启动")

	//first := start("readerWriter1", 3, 6)

	second := start("readerWriter2", 2, 2)

	time.Sleep(2 * time.Second)

	shutdown(second)
	//shutdown(first, second)

	log.Println("程序结束")

}

//创建一个readerWriter,启动(reader/writer)goroutines
func start(name string, maxReads int, numReaders int) *readerWriter {
	rw := readerWriter{
		name:          name,
		shutdown:      make(chan struct{}),
		readerControl: make(semaphore, maxReads),
		maxReads:      maxReads,
		numReaders:    numReaders,
	}
	//启动一定数量的读取goroutine.
	rw.reportShutdown.Add(numReaders)
	for goroutine := 0; goroutine < numReaders; goroutine++ {
		log.Println("-----read goroutine:", goroutine)
		go rw.reader(goroutine)
	}
	//启动一个写入goroutine
	rw.reportShutdown.Add(1)
	go rw.writer()

	return &rw

}

func shutdown(readerWriters ...*readerWriter) {
	var waitShutdown sync.WaitGroup
	waitShutdown.Add(len(readerWriters))

	for _, readerWriter := range readerWriters {
		go readerWriter.stop(&waitShutdown)
	}

	waitShutdown.Wait()
}

func (rw *readerWriter) stop(waitShutdown *sync.WaitGroup) {
	defer waitShutdown.Done()
	log.Printf("%s\t: #####> 停止", rw.name)
	close(rw.shutdown)
	rw.reportShutdown.Wait()
	log.Printf("%s\t: #####> 已经停止", rw.name)
}

// 读取操作开启，
func (rw *readerWriter) reader(reader int) {
	defer rw.reportShutdown.Done()
	//一直执行读取操作，直到接收到shutdown
	for {
		log.Println("-----read goroutine for ", reader)
		select {
		case <-rw.shutdown:
			log.Printf("%s\t: #> reader关闭", rw.name)
			return
		default:
			rw.performRead(reader)
		}
	}
}

// 执行具体读取操作
func (rw *readerWriter) performRead(reader int) {
	//获取一个读取锁。
	rw.ReadLock(reader)
	count := atomic.AddInt32(&rw.currentReads, 1)
	// 模仿一些读取操作
	log.Printf("%s\t: 读取goroutine %d 开始\t- 共[%d]个读取操作\n", rw.name, reader, count)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	count = atomic.AddInt32(&rw.currentReads, -1)
	log.Printf("%s\t: 读取goroutine %d 结束\t- 共[%d]个读取操作\n", rw.name, reader, count)

	rw.ReadUnlock(reader)
}

// 写入操作开启
func (rw *readerWriter) writer() {
	defer rw.reportShutdown.Done()
	//一直执行写入操作，直到接收到shutdown
	for {
		select {
		case <-rw.shutdown:
			log.Printf("%s\t: #> Writer关闭", rw.name)
			return
		default:
			rw.performWrite()
		}
	}
}

// 执行具体的写入操作
func (rw *readerWriter) performWrite() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	log.Printf("%s\t: *****> Writing Pending\n", rw.name)

	// Get a write lock for this critical section.
	rw.WriteLock()

	// 模仿一些写入操作
	log.Printf("%s\t: *****> 写入开始", rw.name)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("%s\t: *****> 写入结束", rw.name)

	rw.WriteUnlock()
}

// ReadLock 向信号量channel中写入一个单位。
func (rw *readerWriter) ReadLock(reader int) {
	// 如果有读取在进行，等待读取完成。
	rw.write.Wait()
	rw.readerControl.Acquire(1)
}

// ReadUnlock 释放信号量中的一个单位.
func (rw *readerWriter) ReadUnlock(reader int) {
	rw.readerControl.Release(1)
}

// WriteLock 阻止所有的读取操作，让写入操作能安全操作。
func (rw *readerWriter) WriteLock() {
	rw.write.Add(1)
	rw.readerControl.Acquire(rw.maxReads)
}

// WriteUnlock 释放写锁，允许读取操作发生.
func (rw *readerWriter) WriteUnlock() {
	rw.readerControl.Release(rw.maxReads)
	rw.write.Done()
}

// 尝试从信号量中获取指定数量的buffer, readerWriter有一个信号量变量，实现读取控制。
func (s semaphore) Acquire(buffers int) {
	var e struct{}

	for buffer := 0; buffer < buffers; buffer++ {
		s <- e
	}
}

// 释放指定数量buffer，返还给信号量。
func (s semaphore) Release(buffers int) {
	for buffer := 0; buffer < buffers; buffer++ {
		<-s
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}
