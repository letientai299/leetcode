package main

type ParkingSystem struct {
	spaces [3]int
}

func Constructor_1603(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		spaces: [3]int{big, medium, small},
	}
}

func (p *ParkingSystem) AddCar(carType int) bool {
	if p.spaces[carType-1] == 0 {
		return false
	}

	p.spaces[carType-1]--
	return true
}
