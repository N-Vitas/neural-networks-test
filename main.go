package main

import (
	"encoding/json"
	"github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"github.com/NOX73/go-neural/persist"
	"io/ioutil"
	"log"
)

const (
	mozgFile   = "json/mozgFile.json"
	speed      = 0.1
	tranerFile = "json/tranerFile.json"
)

type Sample struct {
	In  []float64
	Out []float64
}

func main() {
	// createNetwork()
	n := loadNetwork()

	// testNetwork(n)
	// learnNetwork(n)
	testNetwork(n)

	saveNetwork(n)
}
func getParams(in []float64) string {
	var (a string;b string;c string)
	if(in[0] == 1){a = "Есть подарки"}else{a = "Нет подарков"}
	if(in[1] == 1){b = "Далеко ехать"}else{b = "Рядом ехать"}
	if(in[2] == 1){c = "Скучно"}else{c = "Весело"}
	return a + " " + b + " " + c
}
func total(in float64) string {
	if in > 0.5 {
		return "Да"
	}else{
		return "Нет"
	}
}
func testNetwork(n *neural.Network) {
	// i := 5
	log.Println("--------------------------------------")	
	samples := loadSample()		
	for _, s := range samples {
		log.Printf(getParams(s.In))
		// log.Println("samples",s)
		res := n.Calculate(s.In)
		// e := learn.Evaluation(n, s.In, s.Out)
		log.Printf("Пойдем бухать => %s, ожидаем ответ %s",total(res[0]),total(s.Out[0]))	
		// log.Printf("Пойдем бухать => %s, а ожидаем %s \t %.3f \t\t %v",total(res[0]),total(s.Out[0]),e, res)		
	}
}
func learnNetwork(n *neural.Network) {
	samples := loadSample()
	// есть водка,идет дождь,друг
	for i := 0; i < 10000; i++ {
		for _, s := range samples {
			// log.Println("samples",s)
			learn.Learn(n, s.In, s.Out, speed)
		}

	}

}

func loadSample() []Sample {

	s := []Sample{}
	b, _ := ioutil.ReadFile(tranerFile)
	json.Unmarshal([]byte(b), &s)
	// log.Println(s)
	return s
}

func loadNetwork() *neural.Network {
	return persist.FromFile(mozgFile)
}

func saveNetwork(n *neural.Network) {
	persist.ToFile(mozgFile, n)
}

func createNetwork() {

	n := neural.NewNetwork(3, []int{3, 2, 1})
	n.RandomizeSynapses()

	persist.ToFile(mozgFile, n)
}
