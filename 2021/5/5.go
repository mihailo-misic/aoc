package main

import "fmt"

func main() {
	area := make([][]int, 1000)
	for i := range area {
		area[i] = make([]int, 1000)
	}

	for _, line := range lines {
		line.Draw(&area)
	}

	overlaps := 0
	for _, row := range area {
		for _, num := range row {
			if num > 1 {
				overlaps++
			}
		}
	}

	fmt.Println(overlaps)
}

func (l *Line) Draw(area *[][]int) {
	c := l.Start

	for {
		(*area)[c.Y][c.X] += 1

		xD := 0
		if c.X < l.End.X {
			xD = 1
		} else if c.X > l.End.X {
			xD = -1
		}
		c.X += xD

		yD := 0
		if c.Y < l.End.Y {
			yD = 1
		} else if c.Y > l.End.Y {
			yD = -1
		}
		c.Y += yD

		if xD == 0 && yD == 0 {
			break
		}
	}
}

type Coord struct {
	X int
	Y int
}

type Line struct {
	Start Coord
	End   Coord
}

var lines = []Line{
	Line{Start: Coord{X: 217, Y: 490}, End: Coord{X: 217, Y: 764}},
	Line{Start: Coord{X: 44, Y: 270}, End: Coord{X: 373, Y: 599}},
	Line{Start: Coord{X: 440, Y: 139}, End: Coord{X: 440, Y: 303}},
	Line{Start: Coord{X: 161, Y: 663}, End: Coord{X: 345, Y: 663}},
	Line{Start: Coord{X: 848, Y: 963}, End: Coord{X: 908, Y: 963}},
	Line{Start: Coord{X: 299, Y: 207}, End: Coord{X: 162, Y: 70}},
	Line{Start: Coord{X: 77, Y: 346}, End: Coord{X: 77, Y: 686}},
	Line{Start: Coord{X: 693, Y: 743}, End: Coord{X: 693, Y: 127}},
	Line{Start: Coord{X: 96, Y: 459}, End: Coord{X: 96, Y: 779}},
	Line{Start: Coord{X: 864, Y: 39}, End: Coord{X: 233, Y: 670}},
	Line{Start: Coord{X: 58, Y: 79}, End: Coord{X: 203, Y: 79}},
	Line{Start: Coord{X: 158, Y: 596}, End: Coord{X: 463, Y: 291}},
	Line{Start: Coord{X: 633, Y: 293}, End: Coord{X: 136, Y: 293}},
	Line{Start: Coord{X: 656, Y: 474}, End: Coord{X: 656, Y: 72}},
	Line{Start: Coord{X: 148, Y: 754}, End: Coord{X: 947, Y: 754}},
	Line{Start: Coord{X: 535, Y: 780}, End: Coord{X: 535, Y: 460}},
	Line{Start: Coord{X: 821, Y: 701}, End: Coord{X: 821, Y: 796}},
	Line{Start: Coord{X: 592, Y: 200}, End: Coord{X: 592, Y: 610}},
	Line{Start: Coord{X: 620, Y: 786}, End: Coord{X: 722, Y: 786}},
	Line{Start: Coord{X: 632, Y: 731}, End: Coord{X: 536, Y: 731}},
	Line{Start: Coord{X: 825, Y: 640}, End: Coord{X: 195, Y: 10}},
	Line{Start: Coord{X: 956, Y: 547}, End: Coord{X: 956, Y: 387}},
	Line{Start: Coord{X: 25, Y: 32}, End: Coord{X: 981, Y: 988}},
	Line{Start: Coord{X: 870, Y: 613}, End: Coord{X: 870, Y: 16}},
	Line{Start: Coord{X: 369, Y: 780}, End: Coord{X: 369, Y: 362}},
	Line{Start: Coord{X: 348, Y: 924}, End: Coord{X: 243, Y: 924}},
	Line{Start: Coord{X: 28, Y: 114}, End: Coord{X: 540, Y: 114}},
	Line{Start: Coord{X: 702, Y: 690}, End: Coord{X: 702, Y: 335}},
	Line{Start: Coord{X: 836, Y: 442}, End: Coord{X: 184, Y: 442}},
	Line{Start: Coord{X: 602, Y: 11}, End: Coord{X: 602, Y: 651}},
	Line{Start: Coord{X: 76, Y: 988}, End: Coord{X: 608, Y: 988}},
	Line{Start: Coord{X: 15, Y: 922}, End: Coord{X: 951, Y: 922}},
	Line{Start: Coord{X: 363, Y: 18}, End: Coord{X: 296, Y: 18}},
	Line{Start: Coord{X: 130, Y: 580}, End: Coord{X: 516, Y: 580}},
	Line{Start: Coord{X: 799, Y: 335}, End: Coord{X: 858, Y: 335}},
	Line{Start: Coord{X: 571, Y: 842}, End: Coord{X: 571, Y: 800}},
	Line{Start: Coord{X: 684, Y: 654}, End: Coord{X: 684, Y: 971}},
	Line{Start: Coord{X: 815, Y: 674}, End: Coord{X: 66, Y: 674}},
	Line{Start: Coord{X: 575, Y: 612}, End: Coord{X: 575, Y: 919}},
	Line{Start: Coord{X: 652, Y: 126}, End: Coord{X: 822, Y: 296}},
	Line{Start: Coord{X: 391, Y: 493}, End: Coord{X: 730, Y: 493}},
	Line{Start: Coord{X: 810, Y: 479}, End: Coord{X: 810, Y: 807}},
	Line{Start: Coord{X: 397, Y: 420}, End: Coord{X: 780, Y: 37}},
	Line{Start: Coord{X: 187, Y: 740}, End: Coord{X: 869, Y: 740}},
	Line{Start: Coord{X: 175, Y: 626}, End: Coord{X: 175, Y: 169}},
	Line{Start: Coord{X: 773, Y: 901}, End: Coord{X: 773, Y: 44}},
	Line{Start: Coord{X: 45, Y: 130}, End: Coord{X: 45, Y: 17}},
	Line{Start: Coord{X: 226, Y: 253}, End: Coord{X: 252, Y: 279}},
	Line{Start: Coord{X: 481, Y: 928}, End: Coord{X: 481, Y: 521}},
	Line{Start: Coord{X: 121, Y: 506}, End: Coord{X: 121, Y: 50}},
	Line{Start: Coord{X: 306, Y: 386}, End: Coord{X: 653, Y: 733}},
	Line{Start: Coord{X: 115, Y: 635}, End: Coord{X: 208, Y: 542}},
	Line{Start: Coord{X: 619, Y: 67}, End: Coord{X: 212, Y: 67}},
	Line{Start: Coord{X: 82, Y: 79}, End: Coord{X: 972, Y: 969}},
	Line{Start: Coord{X: 15, Y: 20}, End: Coord{X: 15, Y: 933}},
	Line{Start: Coord{X: 606, Y: 136}, End: Coord{X: 500, Y: 136}},
	Line{Start: Coord{X: 791, Y: 250}, End: Coord{X: 791, Y: 316}},
	Line{Start: Coord{X: 128, Y: 931}, End: Coord{X: 781, Y: 278}},
	Line{Start: Coord{X: 11, Y: 365}, End: Coord{X: 11, Y: 226}},
	Line{Start: Coord{X: 705, Y: 326}, End: Coord{X: 57, Y: 326}},
	Line{Start: Coord{X: 778, Y: 632}, End: Coord{X: 173, Y: 27}},
	Line{Start: Coord{X: 121, Y: 624}, End: Coord{X: 121, Y: 737}},
	Line{Start: Coord{X: 30, Y: 815}, End: Coord{X: 909, Y: 815}},
	Line{Start: Coord{X: 18, Y: 114}, End: Coord{X: 869, Y: 965}},
	Line{Start: Coord{X: 554, Y: 741}, End: Coord{X: 554, Y: 771}},
	Line{Start: Coord{X: 284, Y: 826}, End: Coord{X: 945, Y: 826}},
	Line{Start: Coord{X: 386, Y: 654}, End: Coord{X: 295, Y: 654}},
	Line{Start: Coord{X: 235, Y: 848}, End: Coord{X: 418, Y: 848}},
	Line{Start: Coord{X: 536, Y: 59}, End: Coord{X: 497, Y: 59}},
	Line{Start: Coord{X: 156, Y: 922}, End: Coord{X: 29, Y: 922}},
	Line{Start: Coord{X: 57, Y: 718}, End: Coord{X: 174, Y: 718}},
	Line{Start: Coord{X: 964, Y: 774}, End: Coord{X: 964, Y: 426}},
	Line{Start: Coord{X: 729, Y: 950}, End: Coord{X: 729, Y: 254}},
	Line{Start: Coord{X: 896, Y: 117}, End: Coord{X: 152, Y: 861}},
	Line{Start: Coord{X: 603, Y: 919}, End: Coord{X: 603, Y: 776}},
	Line{Start: Coord{X: 176, Y: 472}, End: Coord{X: 573, Y: 472}},
	Line{Start: Coord{X: 25, Y: 970}, End: Coord{X: 939, Y: 56}},
	Line{Start: Coord{X: 478, Y: 482}, End: Coord{X: 38, Y: 482}},
	Line{Start: Coord{X: 155, Y: 936}, End: Coord{X: 956, Y: 135}},
	Line{Start: Coord{X: 351, Y: 621}, End: Coord{X: 133, Y: 403}},
	Line{Start: Coord{X: 513, Y: 323}, End: Coord{X: 103, Y: 323}},
	Line{Start: Coord{X: 679, Y: 167}, End: Coord{X: 679, Y: 983}},
	Line{Start: Coord{X: 910, Y: 456}, End: Coord{X: 241, Y: 456}},
	Line{Start: Coord{X: 16, Y: 266}, End: Coord{X: 16, Y: 829}},
	Line{Start: Coord{X: 338, Y: 791}, End: Coord{X: 973, Y: 156}},
	Line{Start: Coord{X: 564, Y: 73}, End: Coord{X: 564, Y: 676}},
	Line{Start: Coord{X: 196, Y: 800}, End: Coord{X: 339, Y: 800}},
	Line{Start: Coord{X: 15, Y: 776}, End: Coord{X: 973, Y: 776}},
	Line{Start: Coord{X: 719, Y: 134}, End: Coord{X: 719, Y: 775}},
	Line{Start: Coord{X: 730, Y: 692}, End: Coord{X: 272, Y: 692}},
	Line{Start: Coord{X: 247, Y: 770}, End: Coord{X: 244, Y: 770}},
	Line{Start: Coord{X: 853, Y: 720}, End: Coord{X: 940, Y: 720}},
	Line{Start: Coord{X: 685, Y: 379}, End: Coord{X: 873, Y: 379}},
	Line{Start: Coord{X: 944, Y: 647}, End: Coord{X: 944, Y: 206}},
	Line{Start: Coord{X: 67, Y: 974}, End: Coord{X: 967, Y: 74}},
	Line{Start: Coord{X: 828, Y: 194}, End: Coord{X: 355, Y: 194}},
	Line{Start: Coord{X: 596, Y: 522}, End: Coord{X: 596, Y: 169}},
	Line{Start: Coord{X: 677, Y: 970}, End: Coord{X: 638, Y: 970}},
	Line{Start: Coord{X: 587, Y: 427}, End: Coord{X: 587, Y: 354}},
	Line{Start: Coord{X: 804, Y: 488}, End: Coord{X: 469, Y: 153}},
	Line{Start: Coord{X: 355, Y: 653}, End: Coord{X: 787, Y: 221}},
	Line{Start: Coord{X: 798, Y: 873}, End: Coord{X: 133, Y: 873}},
	Line{Start: Coord{X: 565, Y: 798}, End: Coord{X: 534, Y: 829}},
	Line{Start: Coord{X: 239, Y: 273}, End: Coord{X: 20, Y: 273}},
	Line{Start: Coord{X: 942, Y: 138}, End: Coord{X: 398, Y: 138}},
	Line{Start: Coord{X: 499, Y: 743}, End: Coord{X: 958, Y: 284}},
	Line{Start: Coord{X: 913, Y: 466}, End: Coord{X: 514, Y: 466}},
	Line{Start: Coord{X: 504, Y: 705}, End: Coord{X: 504, Y: 983}},
	Line{Start: Coord{X: 455, Y: 863}, End: Coord{X: 451, Y: 863}},
	Line{Start: Coord{X: 638, Y: 255}, End: Coord{X: 425, Y: 255}},
	Line{Start: Coord{X: 338, Y: 724}, End: Coord{X: 338, Y: 457}},
	Line{Start: Coord{X: 147, Y: 880}, End: Coord{X: 928, Y: 99}},
	Line{Start: Coord{X: 11, Y: 955}, End: Coord{X: 806, Y: 160}},
	Line{Start: Coord{X: 566, Y: 961}, End: Coord{X: 231, Y: 961}},
	Line{Start: Coord{X: 870, Y: 560}, End: Coord{X: 611, Y: 560}},
	Line{Start: Coord{X: 714, Y: 925}, End: Coord{X: 859, Y: 925}},
	Line{Start: Coord{X: 484, Y: 946}, End: Coord{X: 905, Y: 946}},
	Line{Start: Coord{X: 112, Y: 394}, End: Coord{X: 266, Y: 394}},
	Line{Start: Coord{X: 191, Y: 728}, End: Coord{X: 191, Y: 635}},
	Line{Start: Coord{X: 983, Y: 806}, End: Coord{X: 217, Y: 40}},
	Line{Start: Coord{X: 575, Y: 286}, End: Coord{X: 730, Y: 286}},
	Line{Start: Coord{X: 366, Y: 323}, End: Coord{X: 366, Y: 211}},
	Line{Start: Coord{X: 383, Y: 990}, End: Coord{X: 834, Y: 990}},
	Line{Start: Coord{X: 834, Y: 976}, End: Coord{X: 26, Y: 168}},
	Line{Start: Coord{X: 819, Y: 492}, End: Coord{X: 819, Y: 648}},
	Line{Start: Coord{X: 257, Y: 522}, End: Coord{X: 257, Y: 199}},
	Line{Start: Coord{X: 756, Y: 176}, End: Coord{X: 244, Y: 176}},
	Line{Start: Coord{X: 165, Y: 199}, End: Coord{X: 569, Y: 199}},
	Line{Start: Coord{X: 896, Y: 943}, End: Coord{X: 18, Y: 65}},
	Line{Start: Coord{X: 986, Y: 642}, End: Coord{X: 354, Y: 10}},
	Line{Start: Coord{X: 864, Y: 381}, End: Coord{X: 349, Y: 381}},
	Line{Start: Coord{X: 177, Y: 982}, End: Coord{X: 977, Y: 182}},
	Line{Start: Coord{X: 458, Y: 254}, End: Coord{X: 458, Y: 920}},
	Line{Start: Coord{X: 550, Y: 322}, End: Coord{X: 550, Y: 297}},
	Line{Start: Coord{X: 956, Y: 748}, End: Coord{X: 270, Y: 62}},
	Line{Start: Coord{X: 412, Y: 305}, End: Coord{X: 292, Y: 305}},
	Line{Start: Coord{X: 201, Y: 571}, End: Coord{X: 375, Y: 571}},
	Line{Start: Coord{X: 608, Y: 139}, End: Coord{X: 608, Y: 330}},
	Line{Start: Coord{X: 646, Y: 718}, End: Coord{X: 432, Y: 504}},
	Line{Start: Coord{X: 449, Y: 325}, End: Coord{X: 449, Y: 115}},
	Line{Start: Coord{X: 315, Y: 971}, End: Coord{X: 955, Y: 331}},
	Line{Start: Coord{X: 248, Y: 143}, End: Coord{X: 477, Y: 143}},
	Line{Start: Coord{X: 956, Y: 858}, End: Coord{X: 111, Y: 13}},
	Line{Start: Coord{X: 776, Y: 608}, End: Coord{X: 739, Y: 608}},
	Line{Start: Coord{X: 44, Y: 842}, End: Coord{X: 548, Y: 842}},
	Line{Start: Coord{X: 590, Y: 487}, End: Coord{X: 590, Y: 792}},
	Line{Start: Coord{X: 978, Y: 127}, End: Coord{X: 978, Y: 748}},
	Line{Start: Coord{X: 620, Y: 948}, End: Coord{X: 852, Y: 948}},
	Line{Start: Coord{X: 67, Y: 403}, End: Coord{X: 67, Y: 122}},
	Line{Start: Coord{X: 340, Y: 256}, End: Coord{X: 346, Y: 256}},
	Line{Start: Coord{X: 803, Y: 58}, End: Coord{X: 474, Y: 387}},
	Line{Start: Coord{X: 876, Y: 448}, End: Coord{X: 876, Y: 55}},
	Line{Start: Coord{X: 78, Y: 288}, End: Coord{X: 565, Y: 288}},
	Line{Start: Coord{X: 235, Y: 80}, End: Coord{X: 480, Y: 80}},
	Line{Start: Coord{X: 949, Y: 880}, End: Coord{X: 949, Y: 666}},
	Line{Start: Coord{X: 529, Y: 734}, End: Coord{X: 529, Y: 332}},
	Line{Start: Coord{X: 780, Y: 973}, End: Coord{X: 780, Y: 824}},
	Line{Start: Coord{X: 900, Y: 279}, End: Coord{X: 698, Y: 279}},
	Line{Start: Coord{X: 290, Y: 438}, End: Coord{X: 34, Y: 694}},
	Line{Start: Coord{X: 766, Y: 569}, End: Coord{X: 766, Y: 443}},
	Line{Start: Coord{X: 729, Y: 690}, End: Coord{X: 729, Y: 137}},
	Line{Start: Coord{X: 72, Y: 938}, End: Coord{X: 72, Y: 893}},
	Line{Start: Coord{X: 960, Y: 563}, End: Coord{X: 960, Y: 322}},
	Line{Start: Coord{X: 669, Y: 293}, End: Coord{X: 578, Y: 293}},
	Line{Start: Coord{X: 396, Y: 388}, End: Coord{X: 984, Y: 388}},
	Line{Start: Coord{X: 675, Y: 694}, End: Coord{X: 211, Y: 230}},
	Line{Start: Coord{X: 152, Y: 743}, End: Coord{X: 63, Y: 743}},
	Line{Start: Coord{X: 203, Y: 660}, End: Coord{X: 391, Y: 660}},
	Line{Start: Coord{X: 582, Y: 806}, End: Coord{X: 906, Y: 806}},
	Line{Start: Coord{X: 698, Y: 837}, End: Coord{X: 698, Y: 483}},
	Line{Start: Coord{X: 869, Y: 320}, End: Coord{X: 595, Y: 594}},
	Line{Start: Coord{X: 283, Y: 817}, End: Coord{X: 283, Y: 861}},
	Line{Start: Coord{X: 919, Y: 926}, End: Coord{X: 919, Y: 235}},
	Line{Start: Coord{X: 16, Y: 64}, End: Coord{X: 930, Y: 978}},
	Line{Start: Coord{X: 980, Y: 25}, End: Coord{X: 16, Y: 989}},
	Line{Start: Coord{X: 181, Y: 890}, End: Coord{X: 952, Y: 119}},
	Line{Start: Coord{X: 877, Y: 731}, End: Coord{X: 877, Y: 364}},
	Line{Start: Coord{X: 130, Y: 55}, End: Coord{X: 130, Y: 111}},
	Line{Start: Coord{X: 30, Y: 298}, End: Coord{X: 590, Y: 858}},
	Line{Start: Coord{X: 134, Y: 933}, End: Coord{X: 134, Y: 41}},
	Line{Start: Coord{X: 711, Y: 853}, End: Coord{X: 711, Y: 196}},
	Line{Start: Coord{X: 123, Y: 206}, End: Coord{X: 841, Y: 924}},
	Line{Start: Coord{X: 130, Y: 585}, End: Coord{X: 130, Y: 394}},
	Line{Start: Coord{X: 161, Y: 952}, End: Coord{X: 531, Y: 952}},
	Line{Start: Coord{X: 455, Y: 830}, End: Coord{X: 455, Y: 919}},
	Line{Start: Coord{X: 612, Y: 817}, End: Coord{X: 30, Y: 817}},
	Line{Start: Coord{X: 461, Y: 474}, End: Coord{X: 106, Y: 119}},
	Line{Start: Coord{X: 511, Y: 100}, End: Coord{X: 581, Y: 30}},
	Line{Start: Coord{X: 263, Y: 550}, End: Coord{X: 263, Y: 814}},
	Line{Start: Coord{X: 976, Y: 973}, End: Coord{X: 14, Y: 11}},
	Line{Start: Coord{X: 749, Y: 876}, End: Coord{X: 380, Y: 876}},
	Line{Start: Coord{X: 731, Y: 226}, End: Coord{X: 731, Y: 659}},
	Line{Start: Coord{X: 630, Y: 682}, End: Coord{X: 570, Y: 622}},
	Line{Start: Coord{X: 914, Y: 780}, End: Coord{X: 311, Y: 780}},
	Line{Start: Coord{X: 975, Y: 274}, End: Coord{X: 87, Y: 274}},
	Line{Start: Coord{X: 328, Y: 957}, End: Coord{X: 724, Y: 957}},
	Line{Start: Coord{X: 357, Y: 950}, End: Coord{X: 357, Y: 659}},
	Line{Start: Coord{X: 466, Y: 580}, End: Coord{X: 466, Y: 726}},
	Line{Start: Coord{X: 854, Y: 425}, End: Coord{X: 854, Y: 559}},
	Line{Start: Coord{X: 39, Y: 106}, End: Coord{X: 39, Y: 82}},
	Line{Start: Coord{X: 675, Y: 711}, End: Coord{X: 956, Y: 711}},
	Line{Start: Coord{X: 204, Y: 117}, End: Coord{X: 672, Y: 585}},
	Line{Start: Coord{X: 867, Y: 101}, End: Coord{X: 49, Y: 919}},
	Line{Start: Coord{X: 849, Y: 88}, End: Coord{X: 784, Y: 88}},
	Line{Start: Coord{X: 394, Y: 249}, End: Coord{X: 394, Y: 730}},
	Line{Start: Coord{X: 865, Y: 188}, End: Coord{X: 125, Y: 928}},
	Line{Start: Coord{X: 316, Y: 918}, End: Coord{X: 722, Y: 918}},
	Line{Start: Coord{X: 781, Y: 336}, End: Coord{X: 781, Y: 551}},
	Line{Start: Coord{X: 821, Y: 826}, End: Coord{X: 258, Y: 826}},
	Line{Start: Coord{X: 597, Y: 273}, End: Coord{X: 597, Y: 653}},
	Line{Start: Coord{X: 726, Y: 266}, End: Coord{X: 90, Y: 902}},
	Line{Start: Coord{X: 701, Y: 701}, End: Coord{X: 941, Y: 701}},
	Line{Start: Coord{X: 105, Y: 401}, End: Coord{X: 949, Y: 401}},
	Line{Start: Coord{X: 890, Y: 486}, End: Coord{X: 890, Y: 205}},
	Line{Start: Coord{X: 651, Y: 409}, End: Coord{X: 651, Y: 408}},
	Line{Start: Coord{X: 450, Y: 88}, End: Coord{X: 51, Y: 88}},
	Line{Start: Coord{X: 29, Y: 478}, End: Coord{X: 29, Y: 667}},
	Line{Start: Coord{X: 676, Y: 293}, End: Coord{X: 676, Y: 77}},
	Line{Start: Coord{X: 380, Y: 773}, End: Coord{X: 962, Y: 773}},
	Line{Start: Coord{X: 253, Y: 836}, End: Coord{X: 429, Y: 836}},
	Line{Start: Coord{X: 833, Y: 706}, End: Coord{X: 123, Y: 706}},
	Line{Start: Coord{X: 689, Y: 167}, End: Coord{X: 665, Y: 143}},
	Line{Start: Coord{X: 375, Y: 540}, End: Coord{X: 375, Y: 346}},
	Line{Start: Coord{X: 867, Y: 222}, End: Coord{X: 746, Y: 343}},
	Line{Start: Coord{X: 99, Y: 832}, End: Coord{X: 370, Y: 561}},
	Line{Start: Coord{X: 133, Y: 349}, End: Coord{X: 133, Y: 815}},
	Line{Start: Coord{X: 950, Y: 981}, End: Coord{X: 12, Y: 43}},
	Line{Start: Coord{X: 195, Y: 466}, End: Coord{X: 644, Y: 466}},
	Line{Start: Coord{X: 84, Y: 876}, End: Coord{X: 84, Y: 720}},
	Line{Start: Coord{X: 128, Y: 237}, End: Coord{X: 34, Y: 331}},
	Line{Start: Coord{X: 872, Y: 947}, End: Coord{X: 960, Y: 947}},
	Line{Start: Coord{X: 641, Y: 220}, End: Coord{X: 641, Y: 472}},
	Line{Start: Coord{X: 489, Y: 950}, End: Coord{X: 489, Y: 441}},
	Line{Start: Coord{X: 508, Y: 513}, End: Coord{X: 721, Y: 300}},
	Line{Start: Coord{X: 394, Y: 137}, End: Coord{X: 332, Y: 137}},
	Line{Start: Coord{X: 456, Y: 672}, End: Coord{X: 625, Y: 503}},
	Line{Start: Coord{X: 65, Y: 463}, End: Coord{X: 540, Y: 463}},
	Line{Start: Coord{X: 207, Y: 745}, End: Coord{X: 529, Y: 423}},
	Line{Start: Coord{X: 948, Y: 888}, End: Coord{X: 891, Y: 831}},
	Line{Start: Coord{X: 39, Y: 642}, End: Coord{X: 165, Y: 642}},
	Line{Start: Coord{X: 20, Y: 228}, End: Coord{X: 20, Y: 386}},
	Line{Start: Coord{X: 706, Y: 50}, End: Coord{X: 57, Y: 699}},
	Line{Start: Coord{X: 66, Y: 736}, End: Coord{X: 66, Y: 840}},
	Line{Start: Coord{X: 944, Y: 450}, End: Coord{X: 915, Y: 479}},
	Line{Start: Coord{X: 697, Y: 988}, End: Coord{X: 697, Y: 862}},
	Line{Start: Coord{X: 987, Y: 969}, End: Coord{X: 57, Y: 39}},
	Line{Start: Coord{X: 64, Y: 813}, End: Coord{X: 64, Y: 468}},
	Line{Start: Coord{X: 814, Y: 85}, End: Coord{X: 469, Y: 85}},
	Line{Start: Coord{X: 667, Y: 749}, End: Coord{X: 154, Y: 236}},
	Line{Start: Coord{X: 755, Y: 337}, End: Coord{X: 755, Y: 50}},
	Line{Start: Coord{X: 536, Y: 185}, End: Coord{X: 536, Y: 217}},
	Line{Start: Coord{X: 732, Y: 48}, End: Coord{X: 529, Y: 48}},
	Line{Start: Coord{X: 33, Y: 578}, End: Coord{X: 430, Y: 578}},
	Line{Start: Coord{X: 511, Y: 658}, End: Coord{X: 669, Y: 658}},
	Line{Start: Coord{X: 294, Y: 352}, End: Coord{X: 353, Y: 352}},
	Line{Start: Coord{X: 109, Y: 937}, End: Coord{X: 820, Y: 226}},
	Line{Start: Coord{X: 465, Y: 346}, End: Coord{X: 465, Y: 114}},
	Line{Start: Coord{X: 878, Y: 965}, End: Coord{X: 34, Y: 121}},
	Line{Start: Coord{X: 259, Y: 933}, End: Coord{X: 576, Y: 933}},
	Line{Start: Coord{X: 240, Y: 750}, End: Coord{X: 240, Y: 296}},
	Line{Start: Coord{X: 567, Y: 633}, End: Coord{X: 899, Y: 965}},
	Line{Start: Coord{X: 29, Y: 609}, End: Coord{X: 169, Y: 469}},
	Line{Start: Coord{X: 962, Y: 532}, End: Coord{X: 962, Y: 921}},
	Line{Start: Coord{X: 443, Y: 875}, End: Coord{X: 395, Y: 875}},
	Line{Start: Coord{X: 831, Y: 584}, End: Coord{X: 510, Y: 263}},
	Line{Start: Coord{X: 859, Y: 35}, End: Coord{X: 84, Y: 810}},
	Line{Start: Coord{X: 829, Y: 285}, End: Coord{X: 829, Y: 463}},
	Line{Start: Coord{X: 486, Y: 661}, End: Coord{X: 883, Y: 661}},
	Line{Start: Coord{X: 371, Y: 672}, End: Coord{X: 959, Y: 84}},
	Line{Start: Coord{X: 722, Y: 532}, End: Coord{X: 722, Y: 241}},
	Line{Start: Coord{X: 49, Y: 216}, End: Coord{X: 468, Y: 216}},
	Line{Start: Coord{X: 308, Y: 343}, End: Coord{X: 308, Y: 277}},
	Line{Start: Coord{X: 183, Y: 626}, End: Coord{X: 264, Y: 545}},
	Line{Start: Coord{X: 748, Y: 611}, End: Coord{X: 356, Y: 611}},
	Line{Start: Coord{X: 67, Y: 85}, End: Coord{X: 925, Y: 943}},
	Line{Start: Coord{X: 726, Y: 972}, End: Coord{X: 726, Y: 272}},
	Line{Start: Coord{X: 841, Y: 222}, End: Coord{X: 841, Y: 867}},
	Line{Start: Coord{X: 597, Y: 250}, End: Coord{X: 813, Y: 250}},
	Line{Start: Coord{X: 20, Y: 631}, End: Coord{X: 555, Y: 631}},
	Line{Start: Coord{X: 803, Y: 846}, End: Coord{X: 589, Y: 632}},
	Line{Start: Coord{X: 276, Y: 654}, End: Coord{X: 222, Y: 708}},
	Line{Start: Coord{X: 400, Y: 952}, End: Coord{X: 672, Y: 952}},
	Line{Start: Coord{X: 939, Y: 173}, End: Coord{X: 534, Y: 173}},
	Line{Start: Coord{X: 638, Y: 316}, End: Coord{X: 638, Y: 935}},
	Line{Start: Coord{X: 578, Y: 120}, End: Coord{X: 578, Y: 101}},
	Line{Start: Coord{X: 54, Y: 457}, End: Coord{X: 723, Y: 457}},
	Line{Start: Coord{X: 904, Y: 713}, End: Coord{X: 904, Y: 721}},
	Line{Start: Coord{X: 902, Y: 180}, End: Coord{X: 99, Y: 983}},
	Line{Start: Coord{X: 590, Y: 426}, End: Coord{X: 174, Y: 10}},
	Line{Start: Coord{X: 740, Y: 975}, End: Coord{X: 309, Y: 975}},
	Line{Start: Coord{X: 84, Y: 242}, End: Coord{X: 803, Y: 961}},
	Line{Start: Coord{X: 28, Y: 667}, End: Coord{X: 362, Y: 333}},
	Line{Start: Coord{X: 73, Y: 703}, End: Coord{X: 73, Y: 354}},
	Line{Start: Coord{X: 902, Y: 26}, End: Coord{X: 902, Y: 365}},
	Line{Start: Coord{X: 602, Y: 455}, End: Coord{X: 578, Y: 431}},
	Line{Start: Coord{X: 339, Y: 686}, End: Coord{X: 339, Y: 846}},
	Line{Start: Coord{X: 764, Y: 444}, End: Coord{X: 311, Y: 444}},
	Line{Start: Coord{X: 780, Y: 535}, End: Coord{X: 862, Y: 453}},
	Line{Start: Coord{X: 333, Y: 127}, End: Coord{X: 911, Y: 127}},
	Line{Start: Coord{X: 451, Y: 296}, End: Coord{X: 451, Y: 832}},
	Line{Start: Coord{X: 849, Y: 681}, End: Coord{X: 849, Y: 580}},
	Line{Start: Coord{X: 309, Y: 672}, End: Coord{X: 309, Y: 913}},
	Line{Start: Coord{X: 339, Y: 411}, End: Coord{X: 147, Y: 411}},
	Line{Start: Coord{X: 907, Y: 478}, End: Coord{X: 974, Y: 545}},
	Line{Start: Coord{X: 444, Y: 753}, End: Coord{X: 855, Y: 342}},
	Line{Start: Coord{X: 642, Y: 285}, End: Coord{X: 683, Y: 244}},
	Line{Start: Coord{X: 307, Y: 633}, End: Coord{X: 751, Y: 633}},
	Line{Start: Coord{X: 292, Y: 128}, End: Coord{X: 767, Y: 603}},
	Line{Start: Coord{X: 969, Y: 92}, End: Coord{X: 647, Y: 414}},
	Line{Start: Coord{X: 80, Y: 120}, End: Coord{X: 942, Y: 982}},
	Line{Start: Coord{X: 886, Y: 810}, End: Coord{X: 740, Y: 810}},
	Line{Start: Coord{X: 205, Y: 846}, End: Coord{X: 168, Y: 846}},
	Line{Start: Coord{X: 878, Y: 230}, End: Coord{X: 72, Y: 230}},
	Line{Start: Coord{X: 186, Y: 822}, End: Coord{X: 628, Y: 822}},
	Line{Start: Coord{X: 472, Y: 66}, End: Coord{X: 472, Y: 609}},
	Line{Start: Coord{X: 251, Y: 753}, End: Coord{X: 129, Y: 753}},
	Line{Start: Coord{X: 575, Y: 959}, End: Coord{X: 102, Y: 959}},
	Line{Start: Coord{X: 582, Y: 194}, End: Coord{X: 858, Y: 194}},
	Line{Start: Coord{X: 43, Y: 986}, End: Coord{X: 43, Y: 589}},
	Line{Start: Coord{X: 355, Y: 402}, End: Coord{X: 751, Y: 402}},
	Line{Start: Coord{X: 982, Y: 292}, End: Coord{X: 86, Y: 292}},
	Line{Start: Coord{X: 329, Y: 966}, End: Coord{X: 329, Y: 379}},
	Line{Start: Coord{X: 475, Y: 291}, End: Coord{X: 475, Y: 924}},
	Line{Start: Coord{X: 625, Y: 70}, End: Coord{X: 625, Y: 350}},
	Line{Start: Coord{X: 358, Y: 467}, End: Coord{X: 981, Y: 467}},
	Line{Start: Coord{X: 319, Y: 700}, End: Coord{X: 736, Y: 283}},
	Line{Start: Coord{X: 657, Y: 247}, End: Coord{X: 654, Y: 247}},
	Line{Start: Coord{X: 450, Y: 803}, End: Coord{X: 450, Y: 497}},
	Line{Start: Coord{X: 812, Y: 15}, End: Coord{X: 812, Y: 425}},
	Line{Start: Coord{X: 649, Y: 160}, End: Coord{X: 377, Y: 160}},
	Line{Start: Coord{X: 684, Y: 491}, End: Coord{X: 690, Y: 491}},
	Line{Start: Coord{X: 925, Y: 429}, End: Coord{X: 772, Y: 429}},
	Line{Start: Coord{X: 138, Y: 91}, End: Coord{X: 883, Y: 91}},
	Line{Start: Coord{X: 602, Y: 121}, End: Coord{X: 774, Y: 293}},
	Line{Start: Coord{X: 700, Y: 531}, End: Coord{X: 451, Y: 531}},
	Line{Start: Coord{X: 250, Y: 216}, End: Coord{X: 800, Y: 766}},
	Line{Start: Coord{X: 550, Y: 784}, End: Coord{X: 289, Y: 784}},
	Line{Start: Coord{X: 53, Y: 759}, End: Coord{X: 228, Y: 759}},
	Line{Start: Coord{X: 678, Y: 310}, End: Coord{X: 645, Y: 343}},
	Line{Start: Coord{X: 147, Y: 70}, End: Coord{X: 171, Y: 46}},
	Line{Start: Coord{X: 130, Y: 653}, End: Coord{X: 130, Y: 103}},
	Line{Start: Coord{X: 292, Y: 640}, End: Coord{X: 731, Y: 640}},
	Line{Start: Coord{X: 797, Y: 762}, End: Coord{X: 618, Y: 762}},
	Line{Start: Coord{X: 154, Y: 75}, End: Coord{X: 964, Y: 885}},
	Line{Start: Coord{X: 222, Y: 523}, End: Coord{X: 557, Y: 523}},
	Line{Start: Coord{X: 989, Y: 103}, End: Coord{X: 989, Y: 964}},
	Line{Start: Coord{X: 335, Y: 61}, End: Coord{X: 422, Y: 61}},
	Line{Start: Coord{X: 782, Y: 954}, End: Coord{X: 160, Y: 332}},
	Line{Start: Coord{X: 82, Y: 929}, End: Coord{X: 82, Y: 528}},
	Line{Start: Coord{X: 732, Y: 540}, End: Coord{X: 635, Y: 637}},
	Line{Start: Coord{X: 950, Y: 362}, End: Coord{X: 798, Y: 362}},
	Line{Start: Coord{X: 415, Y: 566}, End: Coord{X: 916, Y: 566}},
	Line{Start: Coord{X: 588, Y: 446}, End: Coord{X: 743, Y: 291}},
	Line{Start: Coord{X: 495, Y: 46}, End: Coord{X: 495, Y: 435}},
	Line{Start: Coord{X: 913, Y: 561}, End: Coord{X: 303, Y: 561}},
	Line{Start: Coord{X: 788, Y: 902}, End: Coord{X: 788, Y: 698}},
	Line{Start: Coord{X: 81, Y: 783}, End: Coord{X: 715, Y: 149}},
	Line{Start: Coord{X: 867, Y: 990}, End: Coord{X: 867, Y: 558}},
	Line{Start: Coord{X: 145, Y: 919}, End: Coord{X: 145, Y: 725}},
	Line{Start: Coord{X: 850, Y: 861}, End: Coord{X: 727, Y: 861}},
	Line{Start: Coord{X: 535, Y: 129}, End: Coord{X: 535, Y: 496}},
	Line{Start: Coord{X: 922, Y: 772}, End: Coord{X: 922, Y: 917}},
	Line{Start: Coord{X: 882, Y: 559}, End: Coord{X: 672, Y: 349}},
	Line{Start: Coord{X: 496, Y: 80}, End: Coord{X: 496, Y: 948}},
	Line{Start: Coord{X: 915, Y: 244}, End: Coord{X: 516, Y: 643}},
	Line{Start: Coord{X: 633, Y: 461}, End: Coord{X: 748, Y: 461}},
	Line{Start: Coord{X: 899, Y: 341}, End: Coord{X: 677, Y: 341}},
	Line{Start: Coord{X: 66, Y: 981}, End: Coord{X: 878, Y: 169}},
	Line{Start: Coord{X: 68, Y: 24}, End: Coord{X: 984, Y: 940}},
	Line{Start: Coord{X: 12, Y: 880}, End: Coord{X: 23, Y: 869}},
	Line{Start: Coord{X: 779, Y: 514}, End: Coord{X: 779, Y: 752}},
	Line{Start: Coord{X: 878, Y: 641}, End: Coord{X: 949, Y: 641}},
	Line{Start: Coord{X: 264, Y: 919}, End: Coord{X: 229, Y: 919}},
	Line{Start: Coord{X: 213, Y: 281}, End: Coord{X: 213, Y: 196}},
	Line{Start: Coord{X: 538, Y: 149}, End: Coord{X: 538, Y: 278}},
	Line{Start: Coord{X: 184, Y: 478}, End: Coord{X: 364, Y: 298}},
	Line{Start: Coord{X: 301, Y: 136}, End: Coord{X: 923, Y: 758}},
	Line{Start: Coord{X: 559, Y: 266}, End: Coord{X: 559, Y: 986}},
	Line{Start: Coord{X: 384, Y: 37}, End: Coord{X: 384, Y: 558}},
	Line{Start: Coord{X: 815, Y: 529}, End: Coord{X: 800, Y: 514}},
	Line{Start: Coord{X: 33, Y: 80}, End: Coord{X: 624, Y: 80}},
	Line{Start: Coord{X: 561, Y: 261}, End: Coord{X: 215, Y: 607}},
	Line{Start: Coord{X: 169, Y: 944}, End: Coord{X: 169, Y: 921}},
	Line{Start: Coord{X: 673, Y: 42}, End: Coord{X: 164, Y: 42}},
	Line{Start: Coord{X: 820, Y: 977}, End: Coord{X: 424, Y: 581}},
	Line{Start: Coord{X: 816, Y: 29}, End: Coord{X: 802, Y: 29}},
	Line{Start: Coord{X: 374, Y: 924}, End: Coord{X: 121, Y: 671}},
	Line{Start: Coord{X: 962, Y: 555}, End: Coord{X: 426, Y: 19}},
	Line{Start: Coord{X: 982, Y: 199}, End: Coord{X: 860, Y: 77}},
	Line{Start: Coord{X: 334, Y: 62}, End: Coord{X: 359, Y: 62}},
	Line{Start: Coord{X: 960, Y: 785}, End: Coord{X: 260, Y: 85}},
	Line{Start: Coord{X: 681, Y: 280}, End: Coord{X: 860, Y: 280}},
	Line{Start: Coord{X: 184, Y: 925}, End: Coord{X: 184, Y: 30}},
	Line{Start: Coord{X: 332, Y: 398}, End: Coord{X: 858, Y: 924}},
	Line{Start: Coord{X: 405, Y: 270}, End: Coord{X: 218, Y: 270}},
	Line{Start: Coord{X: 261, Y: 846}, End: Coord{X: 29, Y: 614}},
	Line{Start: Coord{X: 591, Y: 941}, End: Coord{X: 591, Y: 716}},
	Line{Start: Coord{X: 313, Y: 502}, End: Coord{X: 313, Y: 637}},
	Line{Start: Coord{X: 930, Y: 259}, End: Coord{X: 779, Y: 259}},
	Line{Start: Coord{X: 432, Y: 15}, End: Coord{X: 566, Y: 149}},
	Line{Start: Coord{X: 51, Y: 182}, End: Coord{X: 223, Y: 182}},
	Line{Start: Coord{X: 603, Y: 536}, End: Coord{X: 603, Y: 281}},
	Line{Start: Coord{X: 139, Y: 703}, End: Coord{X: 825, Y: 17}},
	Line{Start: Coord{X: 965, Y: 22}, End: Coord{X: 55, Y: 932}},
	Line{Start: Coord{X: 389, Y: 608}, End: Coord{X: 771, Y: 608}},
	Line{Start: Coord{X: 209, Y: 617}, End: Coord{X: 923, Y: 617}},
	Line{Start: Coord{X: 769, Y: 672}, End: Coord{X: 769, Y: 236}},
	Line{Start: Coord{X: 163, Y: 717}, End: Coord{X: 638, Y: 717}},
	Line{Start: Coord{X: 801, Y: 604}, End: Coord{X: 136, Y: 604}},
	Line{Start: Coord{X: 974, Y: 881}, End: Coord{X: 110, Y: 17}},
	Line{Start: Coord{X: 187, Y: 226}, End: Coord{X: 929, Y: 968}},
	Line{Start: Coord{X: 430, Y: 949}, End: Coord{X: 473, Y: 949}},
	Line{Start: Coord{X: 899, Y: 279}, End: Coord{X: 899, Y: 224}},
	Line{Start: Coord{X: 964, Y: 806}, End: Coord{X: 964, Y: 876}},
	Line{Start: Coord{X: 635, Y: 190}, End: Coord{X: 349, Y: 190}},
	Line{Start: Coord{X: 142, Y: 656}, End: Coord{X: 142, Y: 216}},
	Line{Start: Coord{X: 740, Y: 814}, End: Coord{X: 35, Y: 109}},
	Line{Start: Coord{X: 588, Y: 956}, End: Coord{X: 534, Y: 956}},
	Line{Start: Coord{X: 107, Y: 968}, End: Coord{X: 707, Y: 968}},
	Line{Start: Coord{X: 787, Y: 639}, End: Coord{X: 787, Y: 50}},
	Line{Start: Coord{X: 964, Y: 491}, End: Coord{X: 964, Y: 148}},
	Line{Start: Coord{X: 30, Y: 70}, End: Coord{X: 30, Y: 323}},
	Line{Start: Coord{X: 30, Y: 905}, End: Coord{X: 806, Y: 129}},
	Line{Start: Coord{X: 592, Y: 419}, End: Coord{X: 91, Y: 419}},
	Line{Start: Coord{X: 73, Y: 87}, End: Coord{X: 973, Y: 987}},
	Line{Start: Coord{X: 540, Y: 683}, End: Coord{X: 540, Y: 139}},
	Line{Start: Coord{X: 422, Y: 107}, End: Coord{X: 422, Y: 90}},
	Line{Start: Coord{X: 935, Y: 74}, End: Coord{X: 935, Y: 590}},
	Line{Start: Coord{X: 728, Y: 566}, End: Coord{X: 188, Y: 26}},
	Line{Start: Coord{X: 839, Y: 313}, End: Coord{X: 839, Y: 620}},
	Line{Start: Coord{X: 723, Y: 898}, End: Coord{X: 723, Y: 719}},
	Line{Start: Coord{X: 679, Y: 814}, End: Coord{X: 679, Y: 617}},
	Line{Start: Coord{X: 203, Y: 633}, End: Coord{X: 417, Y: 633}},
	Line{Start: Coord{X: 36, Y: 812}, End: Coord{X: 546, Y: 302}},
	Line{Start: Coord{X: 112, Y: 316}, End: Coord{X: 598, Y: 802}},
	Line{Start: Coord{X: 798, Y: 773}, End: Coord{X: 989, Y: 964}},
	Line{Start: Coord{X: 914, Y: 69}, End: Coord{X: 520, Y: 69}},
	Line{Start: Coord{X: 213, Y: 556}, End: Coord{X: 213, Y: 19}},
	Line{Start: Coord{X: 795, Y: 516}, End: Coord{X: 795, Y: 220}},
	Line{Start: Coord{X: 348, Y: 803}, End: Coord{X: 664, Y: 803}},
	Line{Start: Coord{X: 910, Y: 861}, End: Coord{X: 238, Y: 189}},
	Line{Start: Coord{X: 633, Y: 691}, End: Coord{X: 594, Y: 691}},
	Line{Start: Coord{X: 96, Y: 166}, End: Coord{X: 96, Y: 60}},
	Line{Start: Coord{X: 278, Y: 848}, End: Coord{X: 854, Y: 272}},
	Line{Start: Coord{X: 64, Y: 370}, End: Coord{X: 64, Y: 815}},
	Line{Start: Coord{X: 386, Y: 196}, End: Coord{X: 386, Y: 222}},
	Line{Start: Coord{X: 888, Y: 330}, End: Coord{X: 888, Y: 834}},
	Line{Start: Coord{X: 166, Y: 482}, End: Coord{X: 37, Y: 482}},
	Line{Start: Coord{X: 594, Y: 283}, End: Coord{X: 594, Y: 865}},
	Line{Start: Coord{X: 515, Y: 267}, End: Coord{X: 515, Y: 448}},
	Line{Start: Coord{X: 707, Y: 279}, End: Coord{X: 239, Y: 747}},
	Line{Start: Coord{X: 302, Y: 745}, End: Coord{X: 302, Y: 268}},
	Line{Start: Coord{X: 210, Y: 830}, End: Coord{X: 885, Y: 155}},
	Line{Start: Coord{X: 592, Y: 180}, End: Coord{X: 592, Y: 324}},
	Line{Start: Coord{X: 245, Y: 154}, End: Coord{X: 245, Y: 613}},
	Line{Start: Coord{X: 607, Y: 954}, End: Coord{X: 545, Y: 954}},
	Line{Start: Coord{X: 854, Y: 951}, End: Coord{X: 19, Y: 116}},
	Line{Start: Coord{X: 77, Y: 878}, End: Coord{X: 963, Y: 878}},
	Line{Start: Coord{X: 759, Y: 585}, End: Coord{X: 759, Y: 892}},
	Line{Start: Coord{X: 750, Y: 918}, End: Coord{X: 750, Y: 130}},
	Line{Start: Coord{X: 62, Y: 716}, End: Coord{X: 329, Y: 983}},
	Line{Start: Coord{X: 785, Y: 880}, End: Coord{X: 785, Y: 590}},
	Line{Start: Coord{X: 318, Y: 794}, End: Coord{X: 318, Y: 599}},
	Line{Start: Coord{X: 403, Y: 547}, End: Coord{X: 719, Y: 863}},
	Line{Start: Coord{X: 742, Y: 803}, End: Coord{X: 742, Y: 937}},
	Line{Start: Coord{X: 680, Y: 579}, End: Coord{X: 680, Y: 425}},
	Line{Start: Coord{X: 268, Y: 404}, End: Coord{X: 826, Y: 962}},
	Line{Start: Coord{X: 425, Y: 959}, End: Coord{X: 710, Y: 959}},
	Line{Start: Coord{X: 406, Y: 823}, End: Coord{X: 976, Y: 253}},
	Line{Start: Coord{X: 359, Y: 361}, End: Coord{X: 165, Y: 361}},
	Line{Start: Coord{X: 276, Y: 861}, End: Coord{X: 657, Y: 480}},
	Line{Start: Coord{X: 74, Y: 260}, End: Coord{X: 743, Y: 929}},
	Line{Start: Coord{X: 194, Y: 129}, End: Coord{X: 194, Y: 651}},
	Line{Start: Coord{X: 879, Y: 835}, End: Coord{X: 65, Y: 21}},
	Line{Start: Coord{X: 16, Y: 977}, End: Coord{X: 980, Y: 13}},
	Line{Start: Coord{X: 538, Y: 525}, End: Coord{X: 624, Y: 439}},
	Line{Start: Coord{X: 985, Y: 789}, End: Coord{X: 985, Y: 510}},
	Line{Start: Coord{X: 699, Y: 850}, End: Coord{X: 560, Y: 711}},
	Line{Start: Coord{X: 301, Y: 48}, End: Coord{X: 477, Y: 224}},
	Line{Start: Coord{X: 28, Y: 938}, End: Coord{X: 905, Y: 61}},
	Line{Start: Coord{X: 844, Y: 530}, End: Coord{X: 793, Y: 530}},
	Line{Start: Coord{X: 286, Y: 325}, End: Coord{X: 936, Y: 975}},
	Line{Start: Coord{X: 368, Y: 122}, End: Coord{X: 677, Y: 431}},
	Line{Start: Coord{X: 924, Y: 153}, End: Coord{X: 924, Y: 774}},
	Line{Start: Coord{X: 783, Y: 498}, End: Coord{X: 783, Y: 148}},
	Line{Start: Coord{X: 250, Y: 392}, End: Coord{X: 578, Y: 392}},
	Line{Start: Coord{X: 465, Y: 345}, End: Coord{X: 573, Y: 345}},
	Line{Start: Coord{X: 860, Y: 763}, End: Coord{X: 860, Y: 40}},
	Line{Start: Coord{X: 373, Y: 226}, End: Coord{X: 599, Y: 226}},
	Line{Start: Coord{X: 169, Y: 562}, End: Coord{X: 169, Y: 292}},
	Line{Start: Coord{X: 408, Y: 123}, End: Coord{X: 569, Y: 123}},
	Line{Start: Coord{X: 510, Y: 396}, End: Coord{X: 733, Y: 396}},
	Line{Start: Coord{X: 199, Y: 20}, End: Coord{X: 199, Y: 770}},
	Line{Start: Coord{X: 892, Y: 631}, End: Coord{X: 237, Y: 631}},
	Line{Start: Coord{X: 671, Y: 863}, End: Coord{X: 705, Y: 863}},
	Line{Start: Coord{X: 141, Y: 530}, End: Coord{X: 141, Y: 630}},
	Line{Start: Coord{X: 467, Y: 159}, End: Coord{X: 367, Y: 159}},
	Line{Start: Coord{X: 729, Y: 501}, End: Coord{X: 255, Y: 975}},
	Line{Start: Coord{X: 578, Y: 871}, End: Coord{X: 578, Y: 225}},
	Line{Start: Coord{X: 821, Y: 363}, End: Coord{X: 821, Y: 820}},
}
