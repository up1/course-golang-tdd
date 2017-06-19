package math

func Average(datas []float64) float64 {
	total := float64(0)
	for _, v := range datas {
		total += v
	}
	return total / float64(len(datas))
}
