package main

import (
	"fmt"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
)

type Packet struct {
	Start  int
	Buffer string
	Ver    int
	Id     int
	LenId  int
	LenT   int
	Len    int
	CountT int
	Count  int
	NumBin string
	Num    int
	End    int
	Parent *Packet
}

func NewPacket(start, length, count int, parent *Packet) *Packet {
	return &Packet{
		Start:  start,
		Buffer: "",
		Ver:    -1,
		Id:     -1,
		LenId:  -1,
		LenT:   -1,
		Len:    length,
		CountT: -1,
		Count:  count,
		NumBin: "",
		Num:    -1,
		End:    -1,
		Parent: parent,
	}
}

func DecOpen(idx int, openPackets []*Packet) []*Packet {
	oneUp := openPackets[len(openPackets)-1]
	oneUp.Count--
	if oneUp.CountT != -1 && oneUp.Count == 0 {
		//oneUp.End = idx
		openPackets = openPackets[:len(openPackets)-1]

		if len(openPackets) >= 1 {
			return DecOpen(idx, openPackets)
		}

	}

	return openPackets
}

func main() {
	input := ReadFile("./input.txt")

	toRun := []string{}
	for _, line := range input {
		allBits := ""
		for _, r := range line {
			allBits += hexToBin[string(r)]
		}
		toRun = append(toRun, allBits)
	}

	for _, bits := range toRun {
		fmt.Println("bits", len(bits))
		run(bits)
	}
}

func run(allBits string) {
	packets := []*Packet{}
	packets = append(packets, NewPacket(0, -1, -1, nil))
	openPackets := []*Packet{}

	for idx := 0; idx < len(allBits); idx++ {
		cp := packets[len(packets)-1]
		cp.Buffer += string(allBits[idx])
		cpCur := idx - cp.Start

		if cp.Count != -1 { // Count handling (LenId = 1)
			if cp.Count == 0 {
				cp.End = idx
				openPackets = openPackets[:len(openPackets)-1]
				cp.Count--
			}
		}

		if cp.Len != -1 { // Len handling (LenId = 0)
			if cp.Len == 0 {
				cp.End = idx
				openPackets = openPackets[:len(openPackets)-1]
			}

			cp.Len--
		}

		if cp.End == idx { // End handling for Lengther and Counter
			cp.End = idx - 1
			if cp.CountT == -1 {
				//		openPackets = DecOpen(cp.End, openPackets)
			}

			parent := &Packet{}
			if len(openPackets) > 0 {
				parent = openPackets[len(openPackets)-1]
			}

			packets = append(packets, NewPacket(idx, cp.Len, cp.Count, parent))
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
				openPackets = DecOpen(idx, openPackets)

				if cp.Count != -1 {
					cp.Count--
				}
				if cp.Count != 0 && cp.Len != 0 {
					packets = append(packets, NewPacket(idx+1, cp.Len, cp.Count, cp.Parent))
				}
			}

			continue
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
					cp.Len = length
					cp.LenT = length
					cp.End = idx + length
					openPackets = append(openPackets, cp)

					packets = append(packets, NewPacket(idx+1, length, -1, cp))
					cp.Buffer = ""
				}
			}
		}

		if cp.LenId == 1 { // Handling Counter
			if cpCur < 7+11 {
				if cpCur == 7+10 {
					count := binToDec(cp.Buffer)
					cp.Count = count
					cp.CountT = count
					openPackets = append(openPackets, cp)

					packets = append(packets, NewPacket(idx+1, -1, count, cp))
					cp.Buffer = ""
				}
			}
		}
	}

	ver := 0

	for i := len(packets) - 2; i >= 0; i-- {
		p := packets[i]

		if p.Ver != -1 {
			ver += p.Ver
		}

		if p.End == -1 {
			np := packets[i+1]
			for k := 1; k < p.CountT; k++ {
				for _, cp := range packets {
					if cp.Start == np.End+1 {
						np = cp
						break
					}
				}
			}
			p.End = np.End
		}

		par := p.Parent
		spacer := ""
		for par != nil {
			spacer += " "
			par = par.Parent
		}
		if p.Id == 4 {
			fmt.Printf("  %sID:%v   Start:%v   End:%v   Num:%v \n", spacer, p.Id, p.Start, p.End, p.Num)
		} else {
			if p.LenT != -1 {
				if p.Start+p.LenT+21 != p.End {
					fmt.Println("PANIC")
				}
				fmt.Printf("> %sID:%v   Start:%v   End:%v   Len:%v   LID:%v \n", spacer, p.Id, p.Start, p.End, p.LenT, p.LenId)
			} else {
				if p.End == -1 {
					fmt.Println("PANIC")
				}
				fmt.Printf("> %sID:%v   Start:%v   End:%v   Cnt:%v   LID:%v \n", spacer, p.Id, p.Start, p.End, p.CountT, p.LenId)
			}
		}
	}

	fmt.Println(calc(packets[:len(packets)-1])[0].Num)

	// 14533618 -- too low
	// 1114600142730 -- Good
	// 1127730943879 -- too high
	// 2060457365095 -- too high
	fmt.Println("VR", ver)
}

func calc(ps []*Packet) []*Packet {
	if len(ps) == 1 {
		return ps
	}

	for _, p := range ps {
		if p.Id != 4 {
			children, simple := getChildren(p, ps)

			if simple {
				nums := []int{}
				for _, pk := range children {
					nums = append(nums, pk.Num)
				}
				p.Num = getFun(p.Id)(nums...)
				fmt.Println(p.Id, p.Start, p.End, nums, p.Num)
				p.Id = 4

				ps = removeChildren(ps, children)

				break
			}
		}
	}

	return calc(ps)
}

func removeChildren(ps, toRemove []*Packet) []*Packet {
	newPs := []*Packet{}

	hMap := map[int]bool{}
	for _, pr := range toRemove {
		hMap[pr.Start] = true
	}

	for _, p := range ps {
		if hMap[p.Start] != true {
			newPs = append(newPs, p)
		}
	}

	return newPs
}

func getChildren(p *Packet, ps []*Packet) (children []*Packet, simple bool) {
	simple = true

	for _, cp := range ps {
		if cp.Start > p.Start {
			if cp.End > p.End {
				return children, simple
			}

			if cp.Id != 4 {
				simple = false
			}
			children = append(children, cp)
		}
	}

	return children, simple
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
	for _, n := range nums[1:] {
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
	if nums[0] > nums[1] {
		return 1
	}
	return 0
}

func less(nums ...int) (res int) {
	if nums[0] < nums[1] {
		return 1
	}
	return 0
}

func equal(nums ...int) (res int) {
	if nums[0] == nums[1] {
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
