package ex02

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func Test_solutionByWaitGroup(t *testing.T) {
	// for fun - taking value from output
	var bf bytes.Buffer
	log.SetOutput(&bf)
	log.SetFlags(0)
	t.Cleanup(func() {
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags)
	})

	tests := []struct {
		name   string
		args   []int
		wanted []int
	}{
		{
			name:   "From task",
			args:   []int{2, 4, 6, 8, 10},
			wanted: []int{4, 16, 36, 64, 100},
		},
		{
			name:   "Empty slice",
			args:   []int{},
			wanted: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solutionByWaitGroup(tt.args)
			result := preparingOutput(bf)
			bf.Reset()
			if !reflect.DeepEqual(result, tt.wanted) {
				t.Errorf("Test Failed, expect %v, have %v", result, tt.wanted)
			}
		})
	}
}

func preparingOutput(bf bytes.Buffer) (outputAsIntSlice []int) {
	outputAsStr := bf.String()
	outputAsStrSlice := strings.Fields(outputAsStr)
	outputAsIntSlice = make([]int, len(outputAsStrSlice))
	for i, s := range outputAsStrSlice {
		outputAsIntSlice[i], _ = strconv.Atoi(s)
	}
	slices.Sort(outputAsIntSlice)
	return
}
