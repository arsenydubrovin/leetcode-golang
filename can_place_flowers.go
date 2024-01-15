package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	var cnt int

	for i := 0; i < len(flowerbed); i++ {
		sum := flowerbed[i]
		if i > 0 {
			sum += flowerbed[i-1]
		}
		if i < len(flowerbed)-1 {
			sum += flowerbed[i+1]
		}
		if sum == 0 {
			cnt++
			i++
		}
	}

	return n <= cnt
}
