package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	ch := make(chan FreqMap)
	result := FreqMap{}
	for _, ama := range texts {
		go func(word string) {
			ch <- Frequency(word)
		}(ama)
	}
	for range texts {
		for letter, count := range <-ch {
			result[letter] += count
		}
	}
	return result
}
