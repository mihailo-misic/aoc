package main

import (
	"fmt"
	"strconv"

	. "../utils"
)

type Packet struct {
	Start  int
	Buffer string
	Ver    int
	Id     int
	LenId  int
	Len    int
	Count  int
	NumBin string
	Num    int
	End    int
}

func NewPacket(start, length, count int) *Packet {
	return &Packet{
		Start:  start,
		Buffer: "",
		Ver:    -1,
		Id:     -1,
		LenId:  -1,
		Len:    length,
		Count:  count,
		NumBin: "",
		Num:    -1,
		End:    -1,
	}
}

func main() {
	input := ReadFile("./sinput.txt")

	toRun := []string{}
	for _, line := range input {
		allBits := ""
		for _, r := range line {
			allBits += hexToBin[string(r)]
		}
		toRun = append(toRun, allBits)
	}

	for _, bits := range toRun {
		fmt.Println("bits", bits)
		run(bits)
	}

}

func run(allBits string) {
	packets := []*Packet{}
	packets = append(packets, NewPacket(0, -1, -1))

	for idx := 0; idx < len(allBits); idx++ {
		cp := packets[len(packets)-1]
		cp.Buffer += string(allBits[idx])
		cpCur := idx - cp.Start

		if cp.Count != -1 { // Count handling (LenId = 1)
			if cp.Count == 0 {
				cp.End = idx
			}
		}

		if cp.Len != -1 { // Len handling (LenId = 0)
			if cp.Len == 0 {
				cp.End = idx
			}

			cp.Len--
		}

		if cp.End != -1 { // End handling
			packets = append(packets, NewPacket(idx, -1, -1))
			cp = packets[len(packets)-1]
			cp.Buffer += string(allBits[idx])
		}

		if cpCur < 3 { // Version handling
			if len(cp.Buffer) == 3 {
				cp.Ver = binToDec(cp.Buffer)
				cp.Buffer = ""
			}
			continue
		}
		if cpCur < 6 { // Id handling
			if len(cp.Buffer) == 3 {
				cp.Id = binToDec(cp.Buffer)
				cp.Buffer = ""
			}
			continue
		}

		if cp.Id == 4 { // Litteral packet handling
			if len(cp.Buffer) == 1 {
				continue
			}
			if cp.Buffer[0] == '1' && len(cp.Buffer) == 5 {
				cp.NumBin += cp.Buffer[1:]
				cp.Buffer = ""
				continue
			}
			if cp.Buffer[0] == '0' && len(cp.Buffer) == 5 {
				cp.NumBin += cp.Buffer[1:]
				cp.Buffer = ""
				cp.Num = binToDec(cp.NumBin)
				cp.End = idx
				if cp.Count != -1 {
					cp.Count--
					packets = append(packets, NewPacket(idx+1, -1, cp.Count))
				}
				continue
			}
		}

		if cpCur == 6 { // LenId handling
			cp.LenId = binToDec(cp.Buffer)
			cp.Buffer = ""
			continue
		}

		if cp.LenId == 0 { // Handling Lengther
			if cpCur < 7+15 {
				if cpCur == 7+14 {
					length := binToDec(cp.Buffer)
					packets = append(packets, NewPacket(idx+1, length, -1))
					cp.Buffer = ""
				}
				continue
			}
		}

		if cp.LenId == 1 { // Handling Counter
			if cpCur < 7+11 {
				if cpCur == 7+10 {
					count := binToDec(cp.Buffer)
					packets = append(packets, NewPacket(idx+1, -1, count))
					cp.Buffer = ""
				}
				continue
			}
		}
	}

	ans := 0
	nums := []int{}
	calc(p, packets)

	for i := len(packets) - 1; i >= 0; i-- {
		p := packets[i]

		if p.Id == -1 {
			nums = []int{}
			continue
		}
		fmt.Printf("ID:%v Num:%v Start:%v\n", p.Id, p.Num, p.Start)
		if p.Id == 4 {
			nums = append(nums, p.Num)
			continue
		}

	}

	// 14533618 -- too low
	// 2060457365095 -- too high
	fmt.Println("Answer", ans)
}

func calc(p *Packet, packets []*Packet) int {
	if p.Id == -1 {
		calc(packets[0], packets[1:])
	}
	if p.Id == 4 {
		return p.Num
	}

	fun := getFun(p.Id)
}

func getFun(Id int) func(...int) int {
	switch Id {
	case 0:
		return sum
	case 1:
		return product
	case 2:
		return minimum
	case 3:
		return maximum
	case 5:
		return greater
	case 6:
		return less
	default:
		return equal
	}
}

func sum(nums ...int) (res int) {
	for _, n := range nums {
		res += n
	}
	return
}

func product(nums ...int) (res int) {
	res = nums[0]
	for i, n := range nums {
		if i == 0 {
			continue
		}
		res *= n
	}
	return
}

func minimum(nums ...int) (min int) {
	min = 999999999999
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return
}

func maximum(nums ...int) (max int) {
	max = -999999999999
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return
}

func greater(nums ...int) (res int) {
	if nums[1] > nums[0] {
		return 1
	}
	return 0
}

func less(nums ...int) (res int) {
	if nums[1] < nums[0] {
		return 1
	}
	return 0
}

func equal(nums ...int) (res int) {
	if nums[1] == nums[0] {
		return 1
	}
	return 0
}

func binToDec(bin string) int {
	i, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(i)
}

var hexToBin = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}
