package merge

import "testing"

func TestMergeChannelsGoroutines(t *testing.T) {
	if testMergeFunc(5, 5, MergeChannelsGoroutines) != 25 {
		t.Errorf("received not all messages!")
	}
}

func TestMergeChannelsReflect(t *testing.T) {
	if testMergeFunc(5, 5, MergeChannelsReflect) != 25 {
		t.Errorf("received not all messages!")
	}

}

func BenchmarkMergeChannelsGoroutines_little(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testMergeFunc_little(MergeChannelsGoroutines)
	}
}
func BenchmarkMergeChannelsGoroutines_middle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testMergeFunc_middle(MergeChannelsGoroutines)
	}
}
func BenchmarkMergeChannelsGoroutines_big(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testMergeFunc_big(MergeChannelsGoroutines)
	}
}

func BenchmarkMergeChannelsReflect_little(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testMergeFunc_little(MergeChannelsReflect)
	}
}
func BenchmarkMergeChannelsReflect_middle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testMergeFunc_middle(MergeChannelsReflect)
	}
}
func BenchmarkMergeChannelsReflect_big(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testMergeFunc_big(MergeChannelsReflect)
	}
}

func testMergeFunc_little(mergeFunc func(...<-chan int) <-chan int) int {
	return testMergeFunc(10, 100, mergeFunc)
}
func testMergeFunc_middle(mergeFunc func(...<-chan int) <-chan int) int {
	return testMergeFunc(100, 1000, mergeFunc)
}
func testMergeFunc_big(mergeFunc func(...<-chan int) <-chan int) int {
	return testMergeFunc(100, 10000, mergeFunc)
}

func testMergeFunc(chanNum, messNum int, mergeFunc func(...<-chan int) <-chan int) int {
	chs := initChannels(chanNum, messNum)

	resultCh := mergeFunc(chs...)

	i := 0
	for range resultCh {
		i++
	}
	return i
}

func initChannels(chanNum, numOfMessages int) []<-chan int {
	res := make([]<-chan int, 0, chanNum)
	for i := 0; i < chanNum; i++ {
		res = append(res, createChan(numOfMessages))
	}
	return res
}
