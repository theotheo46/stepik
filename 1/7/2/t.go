package main

// не меняйте импорты, они нужны для проверки
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errNegativeOverdraft error = errors.New("expect overdraft >= 0")
var errBalanceMoreThenOverdraft error = errors.New("balance cannot exceed overdraft")

// account представляет счет
type account struct {
	balance   int
	overdraft int
}

func main() {
	var acc account
	var trans []int
	var err error
	acc, trans, err = parseInput()
	if err == nil {
		fmt.Print("-> ")
		fmt.Println(acc, trans)
	} else {
		fmt.Println(err.Error())
	}
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int, error) {
	accSrc, transSrc := readInput()
	acc, err := parseAccount(accSrc)
	if err == nil {
		trans, err := parseTransactions(transSrc)
		return acc, trans, err
	} else {
		return account{}, nil, err
	}
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount(src string) (account, error) {
	parts := strings.Split(src, "/")
	balance, err := strconv.Atoi(parts[0])
	if err != nil {
		return account{}, err
	}
	overdraft, err := strconv.Atoi(parts[1])
	if err != nil {
		return account{}, err
	} else {
		if overdraft < 0 {
			return account{}, errNegativeOverdraft
		}
		if balance < -overdraft {
			return account{}, errBalanceMoreThenOverdraft
		}
	}
	return account{balance, overdraft}, nil
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) ([]int, error) {
	trans := make([]int, len(src))
	for idx, s := range src {
		t, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		} else {
			trans[idx] = t
		}
	}
	return trans, nil
}
