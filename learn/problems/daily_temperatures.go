package problems

type monotonicStack []int

func (s *monotonicStack) Push(val int) {
	*s = append(*s, val)
}

func (s *monotonicStack) Pop() int {
	if len(*s) == 0 {
		return 0
	}

	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return element
}

func (s monotonicStack) Peek() int {
	if len(s) == 0 {
		return 0
	}

	return s[len(s)-1]
}

// Given an array temperatures where temperatures[i] is the temperature on day i, return an array answer such that:
// answer[i] = number of days you have to wait after day i to get a warmer temperature
// if there is no future warmer day, answer[i] = 0
// Example
// temperatures = [73,74,75,71,69,72,76,73]
// output = [1,1,4,2,1,1,0,0]
func dailyTemperatures(temperatures []int) []int {
	stack := monotonicStack{}
	output := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack.Peek()] {
			output[stack.Peek()] = i - stack.Pop()
		}
		stack.Push(i)
	}

	return output
}
