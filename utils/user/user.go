package user

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadUser() (float64, float64) {
	fmt.Print(`
	Para selecionarmos as melhores acomodações para você, precisamos que você selecione o número de estrelas de 1 a 10: 
	`)

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	star := strings.TrimRight(strings.Replace(text, "\n", "", -1), "\r")

	Stars, err := strconv.ParseFloat(star, 32)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(`
	Informe a quantidade de dias que planeja ficar de viagem: 
	`)

	text_, _ := reader.ReadString('\n')
	day := strings.TrimRight(strings.Replace(text_, "\n", "", -1), "\r")

	Days, err := strconv.ParseFloat(day, 32)

	if err != nil {
		fmt.Println(err)
	}

	return Stars, Days
}

func ConvertDaysToMinute(days float64) float64 {
	return days * 1440.0
}

func ReadPathAnswer() bool {
	fmt.Println(`
	Gostaria de saber o melhor trajeto para frequentar os hoteis ?
	`)

	fmt.Println(`
	1 - Sim
	2 - Não 
	`)

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	ans := strings.TrimRight(strings.Replace(text, "\n", "", -1), "\r")

	answer, err := strconv.ParseInt(ans, 32, 64)

	if err != nil {
		fmt.Println(err)
	}

	if answer == 2 {
		return false
	}

	return true
}
