package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

var testSet []BDONG
var trainSet []BDONG
var k int

func main3() {
	k = 36
	fmt.Println("total ")
	fmt.Println(len(bdongs))
	for i := range bdongs {
		trainSet = append(trainSet, bdongs[i])
	}

	var predictions []string
	fmt.Println("test lenght")
	fmt.Println(len(testSet))

	for x := 0; x < len(testSet); x++ {

	}
	for x := range testSet {
		result := testCase(trainSet, testSet[x], k)
		predictions = append(predictions, result[0].key)
		fmt.Printf("Predicted: %s, Actual: %s\n", result[0].key, testSet[x].Institucion)
	}

	accuracy := getAccuracy(testSet, predictions)
	fmt.Printf("Accuracy: %f%s\n", accuracy, "%")
}

func testCase(trainSetA []BDONG, testSetObject BDONG, k int) sortedClassVotes {
	fmt.Println(testSetObject)
	neighbors := getNeighbors(trainSetA, testSetObject, k)
	result := getResponse(neighbors)
	return result
}

func getAccuracy(testSet []BDONG, predictions []string) float64 {
	correct := 0

	for x := range testSet {
		if testSet[x].Institucion == predictions[x] {
			correct += 1
		}
	}

	return (float64(correct) / float64(len(testSet))) * 100.00
}

type classVote struct {
	key   string
	value int
}

type sortedClassVotes []classVote

func (scv sortedClassVotes) Len() int           { return len(scv) }
func (scv sortedClassVotes) Less(i, j int) bool { return scv[i].value < scv[j].value }
func (scv sortedClassVotes) Swap(i, j int)      { scv[i], scv[j] = scv[j], scv[i] }

func getResponse(neighbors []BDONG) sortedClassVotes {
	classVotes := make(map[string]int)

	for x := range neighbors {
		response := neighbors[x].Institucion
		if contains(classVotes, response) {
			classVotes[response] += 1
		} else {
			classVotes[response] = 1
		}
	}

	scv := make(sortedClassVotes, len(classVotes))
	i := 0
	for k, v := range classVotes {
		scv[i] = classVote{k, v}
		i++
	}

	sort.Sort(sort.Reverse(scv))
	return scv
}

type distancePair struct {
	record   BDONG
	distance float64
}

type distancePairs []distancePair

func (slice distancePairs) Len() int           { return len(slice) }
func (slice distancePairs) Less(i, j int) bool { return slice[i].distance < slice[j].distance }
func (slice distancePairs) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

func getNeighbors(trainingSet []BDONG, testRecord BDONG, k int) []BDONG {
	var distances distancePairs
	for i := range trainingSet {

		dist := euclidianDistance(testRecord, trainingSet[i])
		distances = append(distances, distancePair{trainingSet[i], dist})
	}

	sort.Sort(distances)

	var neighbors []BDONG

	for x := 0; x < k; x++ {
		neighbors = append(neighbors, distances[x].record)
	}

	return neighbors
}
func concurrent() {

}

func euclidianDistance(instanceOne BDONG, instanceTwo BDONG) float64 {
	var distance float64

	distance += math.Pow(float64((instanceOne.Numero - instanceTwo.Numero)), 2)

	return math.Sqrt(distance)
}

func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func contains(votesMap map[string]int, name string) bool {
	for s := range votesMap {
		if s == name {
			return true
		}
	}
	return false
}
