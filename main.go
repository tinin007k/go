package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"sync"
	"time"
)

type person struct {
	name string
	age  int
}

type vertex struct {
	x, y float64
}

type shape interface {
	area() float64
	perimeter() float64
}

type cirle struct {
	radius float64
}

type rectangle struct {
	height float64
	width  float64
}

type base struct {
	num int
}

type container struct {
	base
	str string
}

type describer interface {
	describe() string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type argError struct {
	arg  int
	prob string
}

func main() {
	// var i int
	// msg := "test"
	// fmt.Println(msg)
	// msg1 := first.Hello("Toy")
	// fmt.Println(msg1)

	// i = 100
	// fmt.Println(i)

	// for i := 9; i < 12; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	/* 	// whatami := func(i interface{}) {
	   	// 	switch i.(type) {
	   	// 	case int:
	   	// 		fmt.Println("I am an int")
	   	// 	case bool:
	   	// 		fmt.Println("I am a bool")
	   	// 	default:
	   	// 		fmt.Println("Unknown type")
	   	// 	}
	   	// } */

	//whatami("hey")
	//whatami(123)
	//whatami(true)
	// structSample()
	// per := newPerson("Maple")
	// fmt.Println(per.name)
	// fmt.Println(per.age)
	//arraySample()

	//slicessample()
	//mapsSample()
	//rangeSample()
	//num := []int{5, 6, 7, 8, 9}
	//variadicSample(num...)
	// nextInt := closureSample()
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// nextInt1 := closureSample()
	// fmt.Println(nextInt1())

	// j := 1
	// fmt.Println(j)
	// pointerSample(&j)
	// fmt.Println(j)

	// vertex := vertex{10.0, 5.0}
	// fmt.Println(vertex.area())
	// vertex.scale(10)
	// fmt.Println(vertex.area())

	// r := rectangle{2.0, 6.0}
	// c := cirle{5.0}

	// measure(r)
	// measure(c)

	// co := container{
	// 	base: base{
	// 		num: 3,
	// 	},
	// 	str: "some str",
	// }

	// fmt.Println(co.describe())
	// fmt.Println(co.base.num)

	// var d describer
	// d = co
	// fmt.Println("from describe: ", d.describe())
	//m := map[string]int{"k1": 1, "k2": 2, "k3": 4}
	//fmt.Println(genricsSample(m))

	// if k, ae := f2(42); ae != nil { //common go idiom for if
	// 	fmt.Println("f2 failed: ", ae)
	// } else {
	// 	fmt.Println("f1 worked:", k)
	// }

	// _, e := f2(42)
	// if aae, ok := e.(*argError); ok {
	// 	fmt.Println(aae.arg)
	// 	fmt.Println(aae.prob)
	// }

	// f("once")
	// go f("goroutine")

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// time.Sleep(time.Second)
	// fmt.Println("done")

	//channelSample()
	// ping := make(chan string, 1)
	// pong := make(chan string, 1)
	// ping <- "ping"
	// biChannelSample(ping, pong)
	// fmt.Println(<-pong)

	// selectSample()
	// timeoutSample()
	// nonBlockingChanSample()
	// closeChannel()
	// timerSample()
	// tickerSample()
	// workerPoolsSample()
	waitGroupSample()
}

func workerN1(id int, task int) {
	fmt.Println("worker: ", id, "started")
	time.Sleep(2 * time.Second)
	fmt.Println("worker: ", id, "finished task: ", task)
}

func waitGroupSample() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			workerN1(i, i)
		}()
	}
	wg.Wait()
}

func worker(in <-chan int, op chan<- int, id int) {
	for j := range in {
		fmt.Println("worker: ", id, "started job: ", j)
		time.Sleep(1 * time.Second)
		fmt.Println("worker: ", id, "finished procesing job: ", j)
		op <- j * 2
	}
}

func workerPoolsSample() {
	const numJobs = 5
	job := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for i := 1; i < 6; i++ {
		go worker(job, result, i)
	}

	for j := 1; j < 6; j++ {
		job <- j
	}
	close(job)

	for r := 1; r < 6; r++ {
		<-result
	}

	fmt.Println("All jobs one")
}

func tickerSample() {
	done := make(chan bool)
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Done, ticker stopped!!!")
				return
			case t := <-ticker.C:
				fmt.Println("ticker ticked at: ", t.Local())
			}
		}

	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker has ticker 3 times and stopped")
}

func timerSample() {
	done := make(chan bool)
	start := time.Now()
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer1 fired.")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
		done <- true
	}()

	// stop2 := timer2.Stop()
	// if stop2 {
	// 	fmt.Println("Timer 2 stopped")
	// }

	<-done
	fmt.Println("Elapsed time: ", time.Duration(time.Since(start)))

}

func closeChannel() {
	job := make(chan int, 3)
	// done := make(chan bool)

	fmt.Println("Sending 3 jobs to the queue")
	for j := 0; j < 3; j++ {
		job <- j
		fmt.Println("Sent job task: ", j)
	}
	close(job)
	fmt.Println("All job tasks sent, job closed")

	for task := range job {
		fmt.Println("Got task: ", task)
	}
	fmt.Println("All tasks done")
}

func nonBlockingChanSample() {
	message := make(chan string)

	select {
	case res := <-message:
		fmt.Println("received from mmessage: ", res)
	default:
		fmt.Println("no recieve as no message on channel")
	}

	signal := make(chan string)
	msg := "hi"
	select {
	case signal <- msg:
		fmt.Println("sent message: ", msg)
	default:
		fmt.Println("mesg not sent as no receiver")
	}
	//double recv and sent
	select {
	case res := <-message:
		fmt.Println("got message: ", res)
	case signal <- msg:
		fmt.Println("sent msg: ", msg)
	default:
		fmt.Println("no activity")
	}
}

func timeoutSample() {

	pipe := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		pipe <- "no response"
	}()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("timeout detected")
	case res := <-pipe:
		fmt.Println("got response: ", res)
	}

	go func() {
		time.Sleep(2 * time.Second)
		pipe <- "after 2 seconds"
	}()

	select {
	case res := <-pipe:
		fmt.Println("got after 3 seconds: ", res)
	case <-time.After(2 * time.Second):
		fmt.Println("No timeout a wait is more than sleep time.")
	}
}

func selectSample() {
	start := time.Now()
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("Received: ", msg1)
		case msg2 := <-chan2:
			fmt.Println("Received: ", msg2)
		}
		fmt.Println(time.Since(start))
	}
}

func biChannelSample(ping <-chan string, pong chan<- string) {
	in := <-ping
	pong <- in
}

func channelSample() {
	message := make(chan string, 2)
	go func() {
		message <- "ping"
		message <- "ping1"
	}()

	fmt.Println(<-message)
	msg := <-message
	fmt.Println("msg from channel: ", msg)
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, " from: ", i)
	}
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(i int) (int, error) {
	if i == 42 {
		return -1, &argError{42, "can't work with it"}
	} else {
		return i, nil
	}
}

func genricsSample[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for k, v := range m {
		r = append(r, v)
		fmt.Println("key iterated: ", k)
	}
	return r
}

func measure(s shape) {
	fmt.Print("Area: ")
	fmt.Println(s.area())
	fmt.Print("Perimeter: ")
	fmt.Println(s.perimeter())
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c cirle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c cirle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * r.height * r.width
}

func (v *vertex) scale(factor int) {
	v.x = v.x * float64(factor)
	v.y = v.y * float64(factor)
}

func (v *vertex) area() float64 {
	return v.x*v.x + v.y*v.y
}

func pointerSample(i *int) {
	*i = 2
}

func closureSample() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func variadicSample(i ...int) {
	var sum int
	for n := range i {
		sum += n
		fmt.Printf("%d -> %d \n", n, sum)
	}
	fmt.Printf("sum is %d \n", sum)
}

func rangeSample() {
	arr := [3]int{1, 2, 3}
	for num, i := range arr {
		if num == 1 {
			fmt.Println("num is equal to 1")
		} else {
			fmt.Println(i)
			fmt.Println(num)
		}
	}

	m1 := map[string]int{"k1": 1, "k2": 2}
	for k, v := range m1 {
		fmt.Printf("%s -> %d \n", k, v)
	}

	slice1 := []int{1, 2, 3, 4, 5}
	for n, i := range slice1 {
		fmt.Println(n, i)
	}

}

func mapsSample() {
	m1 := make(map[string]int)

	m1["k1"] = 1
	m1["k2"] = 2

	fmt.Println(m1)

	delete(m1, "k1")

	_, prs := m1["k1"]
	prs0, prs1 := m1["k2"]

	n0 := map[string]int{"k1": 1, "k2": 3}
	n1 := map[string]int{"k1": 1, "k2": 4}

	if maps.Equal(n0, n1) {
		fmt.Println("maps are equal")
	} else {
		fmt.Println("maps are unequal")
	}

	fmt.Println(n0)
	fmt.Println(prs)
	fmt.Println(prs0)
	fmt.Println(prs1)
}

func slicessample() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	slice1 := make([]string, 3)
	slice1 = append(slice1, "a", "b")
	fmt.Println(slice1)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	fmt.Println(t2[:3])
	fmt.Println(t2[1:2])
	fmt.Println(t2[0])

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)

}

func arraySample() {
	var a [5]int
	a[1] = 10
	a[2] = 3
	fmt.Println(a)

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	var twoD [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
	fmt.Println(len(twoD[0]))

}

func structSample() {
	p := person{"tom", 34}
	fmt.Println(p.name)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog.name)
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 38
	return &p
}
