
func Sqrt(x float64) float64 {
	
	check := float64(0)
	z := 1.0
	for i:= 0; i<5; i++{
		z -= (z*z - x ) /(2*x)
		if z - check < .01 {
			return z
		}
		check := z
		fmt.Println(check)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
