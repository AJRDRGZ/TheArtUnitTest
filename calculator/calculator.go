package calculator

type MemCalculator struct {
	sum int
}

func (this *MemCalculator) Add(number int) {
	this.sum += number
}

func (this *MemCalculator) Sum() int {
	temp := this.sum
	this.sum = 0
	return temp
}
