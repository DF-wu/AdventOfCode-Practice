package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func filereader() string {
	inputFile, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("file err")
		return ""
	}
	inputString := string(inputFile)
	return inputString
}

func main() {
	input := filereader()
	// p1(input)
	t := time.Now()
	p2(input)
	fmt.Println("final time use: ", time.Since(t))
}

type Reflection struct {
	title       string
	dsr_maxtrix [][3]int // desti source range
}

func p2(input string) {
	intputTokens := strings.Split(input, "\n")
	// consume \r for Windows
	for i, v := range intputTokens {
		intputTokens[i] = strings.Trim(v, "\r")
	}

	reflections := []Reflection{}
	matrixCollection := []string{}
	seeds := []int{}
	// save pool starting pivot
	// only 7 pools
	poolPivots := [7]int{}
	pivotCtr := 0
	for tokenidx, v := range intputTokens {
		// only invoke at the first iteration
		if strings.Contains(v, "seeds") {
			for _, seed := range strings.Fields(v) {
				// the first token is "seeds:", ignore it
				if seed == "seeds:" {
					continue
				} else {
					// convert string to int
					seedInt, _ := strconv.Atoi(seed)
					seeds = append(seeds, seedInt)
				}
			}
			// pass the first iteration
			continue
		}

		// consume first ""
		if tokenidx == 1 {
			continue
		}

		// consume title
		if strings.Contains(v, "map") {
			title := ""
			// fmt.Sscanln(v, "%s map:", &title)
			fmt.Sscanf(v, "%s map", &title)
			refs := Reflection{title: title}
			reflections = append(reflections, refs)

			// save pool starting pivot
			poolPivots[pivotCtr] = tokenidx
			pivotCtr++
		} else if v == "" || tokenidx == len(intputTokens)-1 {
			// end of matrix reading and when it's the last line
			// convert && flush matrixCollection
			// retrieve last reflection
			r := reflections[len(reflections)-1]
			for _, num := range matrixCollection {
				d, s, t := 0, 0, 0
				fmt.Sscanf(num, "%d %d %d", &d, &s, &t)
				tmpArr := [3]int{d, s, t}
				r.dsr_maxtrix = append(r.dsr_maxtrix, tmpArr)

				// reflections[len(reflections)-1].dsr_maxtrix = append(reflections[len(reflections)-1].dsr_maxtrix, tmpArr)
			}
			// update object
			reflections[len(reflections)-1] = r
			fmt.Println(reflections[len(reflections)-1])

			// flush matrixCollection for next matrix
			matrixCollection = []string{}

		} else {
			//consume matrix
			matrixCollection = append(matrixCollection, v)
		}

	}
	// convert seeds to its full range
	// 會爆炸!!
	// for i := 0; i < len(seeds); i = i + 2 {
	// 	s := seeds[i]
	// 	r := seeds[i+1]
	// 	for j := s; j < s+r; j++ {
	// 		fullSeeds = append(fullSeeds, j)
	// 	}

	// }
	fmt.Println("data done.")
	// let's start the reflecting
	// The Reflections are in order. so we can iterate it.
	// when the Reflection.titile change. we can start a new reflection
	pivotCtr = 0
	// set a large number ceiling

	ch := make(chan int, len(seeds)/2)
	wg := new(sync.WaitGroup)

	for i := 0; i < len(seeds); i = i + 2 {
		s := seeds[i]
		r := seeds[i+1]

		wg.Add(1)
		go func() {

			p2_worker(s, r, ch, reflections, wg)
		}()

	}

	wg.Wait()
	close(ch)
	localmin := 9999999999
	resArr := []int{}
	for i := 0; i < len(seeds)/2; i++ {
		val := <-ch
		resArr = append(resArr, val)
		localmin = min(localmin, val)
	}
	fmt.Println(resArr)
	// for v := range ch {
	// 	fmt.Println(v)
	// 	localmin = min(localmin, v)
	// }
	fmt.Println("!!!!!!!!!!!! ans", localmin)

	// 	for _, seed := range fullSeeds {
	// 		currCorrespond := seed
	// 		// always 7 pool to use
	// 		// by title

	// 		for i := 0; i < len(reflections); i++ {
	// 			//iter matrix
	// 			// flag to check if no mapping found -> direct copy

	// 			for _, dsr := range reflections[i].dsr_maxtrix {
	// 				if currCorrespond >= dsr[1] && currCorrespond <= dsr[1]+dsr[2] {
	// 					// reflect
	// 					offset := currCorrespond - dsr[1]
	// 					currCorrespond = dsr[0] + offset

	// 					break
	// 				}
	// 			}

	// 			// fmt.Println(currCorrespond)
	// 		}
	// 		// fmt.Println("final location:", currCorrespond)
	// 		ans = int(math.Min(float64(ans), float64(currCorrespond)))

	// 	}
	// 	anslist = append(anslist, ans)

	// }

}

func p2_worker(s int, r int, ch chan int, reflections []Reflection, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.Now()
	// let's start the reflecting
	localmin := 9999999999
	for i := s; i < s+r; i++ {
		currCorrespond := i
		// always 7 pool to use
		// by title
		for i := 0; i < len(reflections); i++ {
			//iter matrix
			// flag to check if no mapping found -> direct copy
			for _, dsr := range reflections[i].dsr_maxtrix {
				//
				if currCorrespond >= dsr[1] && currCorrespond < dsr[1]+dsr[2] {
					// reflect
					currCorrespond = currCorrespond - dsr[1] + dsr[0]

					break

				}
			}
			// fmt.Print(" This round:", currCorrespond, ",")
		}
		// fmt.Println(" end :", currCorrespond)
		localmin = min(localmin, currCorrespond)

		// fmt.Println("final location:", currCorrespond)
		if i%100000000 == 0 {
			// fmt.Println(currCorrespond)
			fmt.Println(time.Since(t))
		}
	}
	ch <- localmin

}

func p1(input string) {

	intputTokens := strings.Split(input, "\n")
	// consume \r for Windows
	for i, v := range intputTokens {
		intputTokens[i] = strings.Trim(v, "\r")
	}

	reflections := []Reflection{}
	matrixCollection := []string{}
	seeds := []int{}

	// save pool starting pivot
	// only 7 pools
	poolPivots := [7]int{}
	pivotCtr := 0
	for tokenidx, v := range intputTokens {
		// only invoke at the first iteration
		if strings.Contains(v, "seeds") {
			for _, seed := range strings.Fields(v) {
				// the first token is "seeds:", ignore it
				if seed == "seeds:" {
					continue
				} else {
					// convert string to int
					seedInt, _ := strconv.Atoi(seed)
					seeds = append(seeds, seedInt)
				}
			}
			// pass the first iteration
			continue
		}

		// consume first ""
		if tokenidx == 1 {
			continue
		}

		// consume title
		if strings.Contains(v, "map") {
			title := ""
			// fmt.Sscanln(v, "%s map:", &title)
			fmt.Sscanf(v, "%s map", &title)
			refs := Reflection{title: title}
			reflections = append(reflections, refs)

			// save pool starting pivot
			poolPivots[pivotCtr] = tokenidx
			pivotCtr++
		} else if v == "" || tokenidx == len(intputTokens)-1 {
			// end of matrix reading and when it's the last line
			// convert && flush matrixCollection
			// retrieve last reflection
			r := reflections[len(reflections)-1]
			for _, num := range matrixCollection {
				d, s, t := 0, 0, 0
				fmt.Sscanf(num, "%d %d %d", &d, &s, &t)
				tmpArr := [3]int{d, s, t}
				r.dsr_maxtrix = append(r.dsr_maxtrix, tmpArr)

				// reflections[len(reflections)-1].dsr_maxtrix = append(reflections[len(reflections)-1].dsr_maxtrix, tmpArr)
			}
			// update object
			reflections[len(reflections)-1] = r
			fmt.Println(reflections[len(reflections)-1], tokenidx)

			// flush matrixCollection for next matrix
			matrixCollection = []string{}

		} else {
			//consume matrix
			matrixCollection = append(matrixCollection, v)
		}

	}

	// let's start the reflecting
	// The Reflections are in order. so we can iterate it.
	// when the Reflection.titile change. we can start a new reflection
	pivotCtr = 0
	// set a large number ceiling
	ans := 4294967295
	for _, seed := range seeds {
		currCorrespond := seed
		// always 7 pool to use
		// by title

		for i := 0; i < len(reflections); i++ {
			//iter matrix
			// flag to check if no mapping found -> direct copy

			for _, dsr := range reflections[i].dsr_maxtrix {
				if currCorrespond >= dsr[1] && currCorrespond <= dsr[1]+dsr[2] {
					// reflect
					offset := currCorrespond - dsr[1]
					currCorrespond = dsr[0] + offset

					break
				}
			}

			// fmt.Println(currCorrespond)
		}
		fmt.Println("final location:", currCorrespond)
		ans = int(math.Min(float64(ans), float64(currCorrespond)))
	}
	fmt.Println(ans)

}
