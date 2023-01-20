package sample

import (
	proto_gen "github.com/negarciacamilo/go_grpc/proto-gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewKeyboard() *proto_gen.Keyboard {
	keyboard := &proto_gen.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

// NewCPU returns a new sample CPU
func NewCPU() *proto_gen.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &proto_gen.CPU{
		Brand:   brand,
		Name:    name,
		Cores:   uint32(numberCores),
		Threads: uint32(numberThreads),
		MinGhz:  minGhz,
		MaxGhz:  maxGhz,
	}

	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *proto_gen.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	gpu := &proto_gen.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &proto_gen.Memory{
			Value: uint64(memGB),
			Unit:  proto_gen.Memory_GIGABYTE,
		},
	}

	return gpu
}

// NewRAM returns a new sample RAM
func NewRAM() *proto_gen.Memory {
	memGB := randomInt(4, 64)

	ram := &proto_gen.Memory{
		Value: uint64(memGB),
		Unit:  proto_gen.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a new sample SSD
func NewSSD() *proto_gen.Storage {
	memGB := randomInt(128, 1024)

	ssd := &proto_gen.Storage{
		Driver: proto_gen.Storage_SDD,
		Memory: &proto_gen.Memory{
			Value: uint64(memGB),
			Unit:  proto_gen.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD returns a new sample HDD
func NewHDD() *proto_gen.Storage {
	memTB := randomInt(1, 6)

	hdd := &proto_gen.Storage{
		Driver: proto_gen.Storage_HDD,
		Memory: &proto_gen.Memory{
			Value: uint64(memTB),
			Unit:  proto_gen.Memory_TERABYTE,
		},
	}

	return hdd
}

// NewScreen returns a new sample Screen
func NewScreen() *proto_gen.Screen {
	screen := &proto_gen.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// NewLaptop returns a new sample Laptop
func NewLaptop() *proto_gen.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &proto_gen.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*proto_gen.GPU{NewGPU()},
		Storages: []*proto_gen.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &proto_gen.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   timestamppb.Now(),
	}

	return laptop
}

// RandomLaptopScore returns a random laptop score
func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}
