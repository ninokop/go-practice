package util

import (
	"log"
	"testing"
)

func TestInterSection(t *testing.T) {
	d := [][]int{
		{1, 3, 5, 7, 9, 20, 30, 50, 99}, //9
		{5, 8, 9, 20},                   //4
		{2, 7, 9, 30, 55},               //5
		{6, 8, 9, 10, 20, 30, 50},       //7
	}

	log.Printf("%v", MultiInterSection(d))
}
