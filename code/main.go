package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
)

/*
Для добавления echo в проект:
terminal -> 	go mod init
				go get github.com/labstack/echo/v4
*/

func main() {
	fmt.Println("We're ready!")

	s := echo.New()

	s.GET("/days", Handler)

	err := s.Start(":8081")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx echo.Context) error {
	t := time.Now()
	n := fmt.Sprintf("Сегодня %02d.%02d.%02d. Текущее время %02d часов %02d минут.\n\n", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())

	err_n := ctx.String(http.StatusOK, n)
	if err_n != nil {
		return err_n
	}

	// Мой ДР
	d_my := time.Date(2024, time.September, 5, 0, 00, 0, 0, time.UTC)
	my_born := time.Date(1997, time.September, 5, 0, 0, 0, 0, time.UTC)
	durat_my := time.Until(d_my)
	my_b := fmt.Sprintf("Я родился %02d.%02d.%02d. В 2024 году до него осталось %d дней. Это %d часов. Это %d минут. Это %d секунд.\nМне будет %d лет. \n\n",
		my_born.Day(), my_born.Month(), my_born.Year(), int(durat_my.Hours())/24, int(durat_my.Hours()), int(durat_my.Minutes()), int(durat_my.Seconds()),
		d_my.Year()-my_born.Year())

	err_my := ctx.String(http.StatusOK, my_b)
	if err_my != nil {
		return err_my
	}

	// Сестры ДР
	d_sis := time.Date(2024, time.May, 1, 0, 0, 0, 0, time.UTC)
	sis_born := time.Date(2006, time.May, 1, 0, 0, 0, 0, time.UTC)
	durat_sis := time.Until(d_sis)
	sis_b := fmt.Sprintf("Моя Сестра родилась %02d.%02d.%02d. В 2024 году до него осталось %d дней. Это %d часов. Это %d минут. Это %d секунд.\nЕй будет %d лет. \n\n",
		sis_born.Day(), sis_born.Month(), sis_born.Year(), int(durat_sis.Hours())/24, int(durat_sis.Hours()), int(durat_sis.Minutes()), int(durat_sis.Seconds()),
		d_sis.Year()-sis_born.Year())

	err_sis := ctx.String(http.StatusOK, sis_b)
	if err_sis != nil {
		return err_sis
	}

	// Мамин ДР
	d_mom := time.Date(2024, time.March, 7, 0, 0, 0, 0, time.UTC)
	mom_born := time.Date(1977, time.March, 7, 0, 0, 0, 0, time.UTC)
	durat_mom := time.Until(d_mom)
	mom_b := fmt.Sprintf("Моя Мама родилась %02d.%02d.%02d. В 2024 году до него осталось %d дней. Это %d часов. Это %d минут. Это %d секунд.\nЕй будет %d лет. \n\n",
		mom_born.Day(), mom_born.Month(), mom_born.Year(), int(durat_mom.Hours())/24, int(durat_mom.Hours()), int(durat_mom.Minutes()), int(durat_mom.Seconds()),
		d_mom.Year()-mom_born.Year())

	err_mom := ctx.String(http.StatusOK, mom_b)
	if err_mom != nil {
		return err_mom
	}

	// Папин ДР
	d_dad := time.Date(2024, time.May, 1, 0, 0, 0, 0, time.UTC)
	dad_born := time.Date(1972, time.May, 1, 0, 0, 0, 0, time.UTC)
	durat_dad := time.Until(d_dad)
	dad_b := fmt.Sprintf("Мой Папа родился %02d.%02d.%02d. В 2024 году до него осталось %d дней. Это %d часов. Это %d минут. Это %d секунд.\nЕму будет %d лет. \n\n",
		dad_born.Day(), dad_born.Month(), dad_born.Year(), int(durat_dad.Hours())/24, int(durat_dad.Hours()), int(durat_dad.Minutes()), int(durat_dad.Seconds()),
		d_dad.Year()-dad_born.Year())

	err_dad := ctx.String(http.StatusOK, dad_b)
	if err_dad != nil {
		return err_dad
	}

	// Бабушки Риты ДР
	d_rita := time.Date(2024, time.December, 21, 0, 0, 0, 0, time.UTC)
	rita_born := time.Date(1949, time.December, 21, 0, 0, 0, 0, time.UTC)
	rita_age := d_rita.Year() - rita_born.Year()
	durat_rita := time.Until(d_rita)
	rita_b := fmt.Sprintf("Бабушка Рита родилась %02d.%02d.%02d. В 2024 году до него осталось %d дней. Это %d часов. Это %d минут. Это %d секунд.\nЕй будет %d лет. \n\n",
		rita_born.Day(), rita_born.Month(), rita_born.Year(), int(durat_rita.Hours())/24, int(durat_rita.Hours()), int(durat_rita.Minutes()), int(durat_rita.Seconds()), rita_age)

	err_rita := ctx.String(http.StatusOK, rita_b)
	if err_rita != nil {
		return err_rita
	}

	// Бабушки Тани ДР
	d_tanya := time.Date(2024, time.February, 7, 0, 0, 0, 0, time.UTC)
	tanya_born := time.Date(1947, time.February, 7, 0, 0, 0, 0, time.UTC)
	tanya_age := d_tanya.Year() - tanya_born.Year()
	durat_tanya := time.Until(d_rita)
	tanya_b := fmt.Sprintf("Бабушка Таня родилась %02d.%02d.%02d. В 2024 году до него осталось %d дней. Это %d часов. Это %d минут. Это %d секунд.\nЕй будет %d лет. \n\n",
		tanya_born.Day(), tanya_born.Month(), tanya_born.Year(), int(durat_tanya.Hours())/24, int(durat_tanya.Hours()), int(durat_tanya.Minutes()), int(durat_tanya.Seconds()), tanya_age)

	err_tanya := ctx.String(http.StatusOK, tanya_b)
	if err_tanya != nil {
		return err_tanya
	}

	// Данилин ДР
	d_pupan := time.Date(2024, time.November, 25, 0, 0, 0, 0, time.UTC)
	pupan_born := time.Date(1997, time.November, 25, 0, 0, 0, 0, time.UTC)
	durat_pupan := time.Until(d_pupan)
	pupan_b := fmt.Sprintf("Данил родился %02d.%02d.%02d. В 2024 году до него осталось %d дней. Это %d часов. Это %d минут. Это %d секунд.\nЕму будет %d лет. \n\n",
		pupan_born.Day(), pupan_born.Month(), pupan_born.Year(), int(durat_pupan.Hours())/24, int(durat_pupan.Hours()), int(durat_pupan.Minutes()), int(durat_pupan.Seconds()), d_pupan.Year()-pupan_born.Year())

	err_pupan := ctx.String(http.StatusOK, pupan_b)
	if err_pupan != nil {
		return err_pupan
	}

	// Пашин ДР
	d_shara := time.Date(2024, time.September, 7, 0, 0, 0, 0, time.UTC)
	shara_born := time.Date(1999, time.September, 7, 0, 0, 0, 0, time.UTC)
	durat_shara := time.Until(d_shara)
	shara_b := fmt.Sprintf("Паша родился %02d.%02d.%02d. В 2024 году до него осталось %d дней. Это %d часов. Это %d минут. Это %d секунд.\nЕму будет %d лет. \n\n",
		shara_born.Day(), shara_born.Month(), shara_born.Year(), int(durat_shara.Hours())/24, int(durat_shara.Hours()), int(durat_shara.Minutes()), int(durat_shara.Seconds()), d_shara.Year()-shara_born.Year())

	err_shara := ctx.String(http.StatusOK, shara_b)
	if err_shara != nil {
		return err_shara
	}

	return nil
}

