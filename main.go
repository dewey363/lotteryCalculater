package main

import (
	"bufio"
	"flag"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jin925/lotteryCalculater/lottery"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	iFileName, oFileName *string
)

func init() {
	iFileName = flag.String("i", "input.txt", "path of input file")
	oFileName = flag.String("o", "output.xls", "path of output file")
	flag.Parse()
}

func main() {
	// 从输入文件中读取内容
	f, err := os.Open(*iFileName)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	var input []lottery.Input
	for s.Scan() {
		line := s.Text()
		reg := regexp.MustCompile(`^((?:\d+,)+\d+)取(\d+)$`)
		res := reg.FindStringSubmatch(line)
		if len(res) == 0 {
			log.Fatal("文件输入格式不正确")
		}
		str, pick := res[1], res[2]
		pickInt, err := strconv.Atoi(pick)
		if err != nil {
			log.Fatal("文件输入格式不正确")
		}
		input = append(input, lottery.Input{String: str, Pick: pickInt})
	}
	zuhe := lottery.ZuHe(input)
	names := lottery.GetResultName()
	names = append([]string{"号码"}, names...)
	// 写文档
	oFile := excelize.NewFile()
	oFile.SetSheetRow("sheet1", "A1", &names) // 首行
	for k, v := range zuhe {
		var r []string
		for _, v := range lottery.Compute(v) {
			r = append(r, strconv.Itoa(v.Count))
		}
		data := append([]string{v}, r...)
		oFile.SetSheetRow("Sheet1", "A"+strconv.Itoa(k+2), &data)
	}
	if err := oFile.SaveAs(*oFileName); err != nil {
		log.Fatal(err)
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
}
