package main


import (
  "fmt"
  "errors"
)



const winePrice = 100

func main() {
    change, err := buyWine(28, 118)
    if err != nil {
      fmt.Println("Покупка не успешна:", err.Error())
    }else {
      fmt.Println("Ваша сдача - %d. Хорошего дня!\n", change)
    }

    change, err = buyWine(17, 118)
    if err != nil {
      fmt.Println("Покупка не успешна:", err.Error())
    }else {
      fmt.Println("Ваша сдача - %d. Хорошего дня!", change)
    }

    change, err = buyWine(33, 98)
    if err != nil {
      fmt.Println("Покупка не успешна:", err.Error())
    }else {
      fmt.Println("Ваша сдача - %d. Хорошего дня!", change)
    }

}

func buyWine(age, moneyAmount int) (int, error) {
    if age >= 18 && moneyAmount >= winePrice {
        return moneyAmount - winePrice, nil
    } else if age < 18 {
        return moneyAmount, errors.New("Пей вишневый сок, малыш")
    } else {
        return moneyAmount, errors.New("У Вас недостаточно средств")
    }
}
