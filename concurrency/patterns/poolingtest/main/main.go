package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lovexiaoe/golangpros/concurrency/patterns/poolingtest"
)

const (
	numberOfGoroutines = 25
	pooledResources    = 2
)

//对象模仿一个可以共享的资源：数据库连接
type dbConnect struct {
	ID int32
}

//实现io.Closer接口，使得dbConnect能够被pool管理，Close方法释放占用的资源。
func (dbConn *dbConnect) Close() error {
	log.Println("db连接关闭：", dbConn.ID)
	return nil
}

var idCounter int32

func main() {
	var wg sync.WaitGroup
	wg.Add(numberOfGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < numberOfGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("程序关闭")
	p.Close()
}

//定义pool的工厂方法，新建pool时作为参数传入。该工厂方法创建一个先的db链接。
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("创建了新的链接：", id)
	return &dbConnect{id}, nil
}

//从Pool中取出资源，进行查询处理。
func performQueries(query int, p *pool.Pool) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)
	log.Printf("查询: 查询ID[%d] 链接ID[%d]\n", query, conn.(*dbConnect).ID)
}
