package main

func Chunk(slice []int, size int) {
	if size <= 0 {
		print("\n")
		return
	}
	var chunks [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	printChunks(chunks)
}

func printChunks(chunks [][]int) {
	print("[")
	for i, chunk := range chunks {
		print("[")
		for j, val := range chunk {
			print(val)
			if j < len(chunk)-1 {
				print(" ")
			}
		}
		print("]")
		if i < len(chunks)-1 {
			print(" ")
		}
	}
	print("]\n")
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
