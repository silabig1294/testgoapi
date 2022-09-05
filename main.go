package main

import (
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	// "github.com/gofiber/swagger"
	"github.com/jmoiron/sqlx"
	// _ "backend/docs/testapi"
	// swagger "github.com/arsmn/fiber-swagger/v2"
	// swagger "github.com/arsmn/fiber-swagger/v2"
)

type YearAccount struct {
	Year    string `db:"year" json:"year"`
	Account int    `db:"account_count" json:"account_count"`
}

type YearMonthAccount2 struct {
	Year    string `db:"year" json:"year"`
	Month   string `db:"month" json:"month"`
	Account int    `db:"account_count" json:"account_count"`
}

type AllAccount struct {
	AccountAll int `db:"account_all" json:"account_all"`
}
type YearMonthAccount struct {
	Year     string `db:"year" json:"year"`
	Month    string `db:"month" json:"month"`
	Yearback int    `db:"account_count" json:"account_count"`
}

type Date struct {
	Date_from string `json:"date_from"`
	Date_to   string `json:"date_to"`
}

type Date1 struct {
	Last_year string `json:"last_year"`
}

type Date3 struct {
	Date1 string `json:"date_1"`
	Date2 string `json:"date_2"`
	Date3 string `json:"date_3"`
	Date4 string `json:"date_4"`
	Date5 string `json:"date_5"`
}

type MonthAccount struct {
	Month   string `db:"month" json:"month"`
	Account int    `db:"account_count" json:"account_count"`
}

type AccountByYears struct {
	Year  string `json:"year"`
	Total int    `json:"total"`
	Data  []MonthAccount
}

type YearAgo struct {
	Year_total string `json:"year_total"`
	Account    int    `json:"counto"`
}

type Date2 struct {
	Year_back string `json:"year_back"`
}

type PercentOfUserYearly struct {
	Year                  string `db:"year" json:"year"`
	Percent_User_Per_Year string `db:"percent" json:"percent"`
}

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/one_id_account_dr_20220809") //one_id_account_dr_20220809//openid_db
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	// middle ware
	app.Use(recover.New())
	app.Use(cors.New())
	// app.Get("/swagger/*", swagger.HandlerDefault) // default
	// graph account per year and
	v1 := app.Group("/api/v1")
	v1.Get("/separation/yearly", YearlyAccountArray)
	v1.Get("/separation/yearly/query", YearlyAccountQueryArray)
	v1.Get("/gather/:y?/", YearlyAccountGather)
	v1.Get("/total/account", TotalAccount)
	v1.Get("/all/account", TotalAccountArray)
	// v1.Get("/percent/user",PercentUsersYearly)
	v1.Get("/percent/user", PercentUsersYearlyArray)

	// app.Get("/api/v1/percent_status_transac_api", PercentStatusTransactOfApi)
	v1.Get("service/citizen/:y?/:m?", YearlyMonthlyAccountGather)

	// table
	v2 := app.Group("/api/v2")
	v2.Get("/cumulative/total", YearlyTotalAccount)
	v2.Get("/separation/:y?", YearlyAccountQueryMap)
	// v2.Get("service/citizen/:y1?/:y2?", YearlyAllAccountGather)
	v2.Get("service/citizen/:y1?/:y2?", YearlyAllAccountGather2)

	// app.Get("/swagger/*", swagger.HandlerDefault) // default
	// app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
	// 	URL: "http://example.com/doc.json",
	// 	DeepLinking: false,
	// 	// Expand ("list") or Collapse ("none") tag groups by default
	// 	DocExpansion: "none",
	// 	// Prefill OAuth ClientId on Authorize popup
	// 	OAuth: &swagger.OAuthConfig{
	// 		AppName:  "OAuth Provider",
	// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
	// 	},
	// 	// Ability to change OAuth2 redirect uri location
	// 	OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	// }))

	app.Listen(":8000")

	// err := app.Listen(":1323")
	// if err != nil {
	// 	log.Fatalf("fiber.Listen failed %s",err)
	// }

}

// func PercentUsersYearly(c *fiber.Ctx) error {
// 	var err error
// 	percent := []PercentOfUserYearly{}
// 	query := "select year(created_at) year, (count(id) * 100.0)/(select count(id) from account) percent from account group by year(created_at) having percent > 0 order by percent desc"
// 	err = db.Select(&percent, query)
// 	if err != nil {
// 		return err
// 	}
// 	return c.JSON(percent)
// }
func YearlyAllAccountGather2(c *fiber.Ctx) error {
	y1 := c.Params("y1")
	y2 := c.Params("y2")
	if y1 == "" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 94165) as total where date between date(?) and date(now());" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		o5 := AccountByYears{Year: year, Total: years[3].Account, Data: years[0:3]}
		result := []AccountByYears{o5}
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
		return c.JSON(result)
		// years := []YearAccount{}
		// query := "select year, @total := @total + account as account_count from counteryearlyaccount, (Select @total := 0) as total"
		// err := db.Select(&years, query)
		// if err != nil {
		// 	return err
		// }
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":  "success",
		// })
	} else if y1 == "2018" && y2 == "2022" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		time1 := today.Year() - 3
		time2 := today.Year() - 2
		time3 := today.Year() - 1
		time4 := today.Year()
		year := strconv.Itoa(time)
		year1 := strconv.Itoa(time1)
		year2 := strconv.Itoa(time2)
		year3 := strconv.Itoa(time3)
		year4 := strconv.Itoa(time4)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 0) as total where date between date(?) and date(now());" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		// year_2018 := []MonthAccount{}
		// for i:=0;i<len(years)-40;i++{

		// 	// year_2018 := []MonthAccount{}
		// 	// s := append(year_2018,years[i].Month,years[i].Account)
		// 	year_2018[i] = years[i]
		// 	// year_2018[i].Account = years[i].Account
		// }
		o1 := AccountByYears{Year: year, Total: years[9].Account, Data: years[0:10]}
		o2 := AccountByYears{Year: year1, Total: years[21].Account, Data: years[10:22]}
		o3 := AccountByYears{Year: year2, Total: years[33].Account, Data: years[22:34]}
		o4 := AccountByYears{Year: year3, Total: years[45].Account, Data: years[34:46]}
		o5 := AccountByYears{Year: year4, Total: years[49].Account, Data: years[46:50]}
		result := []AccountByYears{o1, o2, o3, o4, o5}

		// fmt.Println(year1)
		// return c.JSON(fiber.Map{
		// 	"results": result,
		// 	"status":    "success",
		// })
		return c.JSON(result)
	} else if y1 == "2019" && y2 == "2022" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() - 3 //Date(-5,0,0).Format("2006-01-01")
		time1 := today.Year() - 2
		time2 := today.Year() - 1
		time3 := today.Year()
		year := strconv.Itoa(time)
		year1 := strconv.Itoa(time1)
		year2 := strconv.Itoa(time2)
		year3 := strconv.Itoa(time3)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 9136) as total where date between date(?) and date(now());" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		// o1 := AccountByYears{Year: year4,Total: years[9].Account,Data: years[0:10]}
		o2 := AccountByYears{Year: year, Total: years[11].Account, Data: years[0:12]}
		o3 := AccountByYears{Year: year1, Total: years[23].Account, Data: years[12:24]}
		o4 := AccountByYears{Year: year2, Total: years[35].Account, Data: years[24:36]}
		o5 := AccountByYears{Year: year3, Total: years[39].Account, Data: years[36:40]}
		// fmt.Println(o1)
		result := []AccountByYears{o2, o3, o4, o5}
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
		return c.JSON(result)
	} else if y1 == "2020" && y2 == "2022" {
		var err error
		request := Date1{}
		today := time.Now()
		// time4 := today.Year()-4
		// time3 := today.Year()-3
		time := today.Year() - 2 //Date(-5,0,0).Format("2006-01-01")
		time1 := today.Year() - 1
		time2 := today.Year()
		year := strconv.Itoa(time)
		year1 := strconv.Itoa(time1)
		year2 := strconv.Itoa(time2)
		// year3 := strconv.Itoa(time3)
		// year4 := strconv.Itoa(time4)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 20465) as total where date between date(?) and date(now());" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		o3 := AccountByYears{Year: year, Total: years[11].Account, Data: years[0:12]}
		o4 := AccountByYears{Year: year1, Total: years[23].Account, Data: years[12:24]}
		o5 := AccountByYears{Year: year2, Total: years[27].Account, Data: years[24:28]}
		// fmt.Println(o1,o2)
		result := []AccountByYears{o3, o4, o5}
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
		return c.JSON(result)
	} else if y1 == "2021" && y2 == "2022" {
		var err error
		request := Date1{}
		today := time.Now()
		// time4 := today.Year()-4
		// time3 := today.Year()-3 //Date(-5,0,0).Format("2006-01-01")
		// time2 := today.Year()-2
		time := today.Year() - 1 //Date(-5,0,0).Format("2006-01-01")
		time1 := today.Year()
		year := strconv.Itoa(time)
		year1 := strconv.Itoa(time1)
		// year2 := strconv.Itoa(time2)
		// year3 := strconv.Itoa(time3)
		// year4 := strconv.Itoa(time4)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 39610) as total where date between date(?) and date(now());" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		o4 := AccountByYears{Year: year, Total: years[11].Account, Data: years[0:12]}
		o5 := AccountByYears{Year: year1, Total: years[15].Account, Data: years[12:16]}
		// fmt.Println(o1,o2,o3)
		result := []AccountByYears{o4, o5}
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
		return c.JSON(result)
	} else if y1 == "2022" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 94165) as total where date between date(?) and date(now());" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		o5 := AccountByYears{Year: year, Total: years[3].Account, Data: years[0:3]}
		result := []AccountByYears{o5}
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
		return c.JSON(result)
	} else if y1 == "2018" && y2 == "2021" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		time1 := today.Year() - 3
		time2 := today.Year() - 2
		time3 := today.Year() - 1

		year := strconv.Itoa(time)
		year1 := strconv.Itoa(time1)
		year2 := strconv.Itoa(time2)
		year3 := strconv.Itoa(time3)

		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time_2 := today.Year() - 1
		year_2 := strconv.Itoa(time_2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year_2 + month2 + day2
		request.Date_to = time_back2
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 0) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		o1 := AccountByYears{Year: year, Total: years[9].Account, Data: years[0:10]}
		o2 := AccountByYears{Year: year1, Total: years[21].Account, Data: years[10:22]}
		o3 := AccountByYears{Year: year2, Total: years[33].Account, Data: years[22:34]}
		o4 := AccountByYears{Year: year3, Total: years[45].Account, Data: years[34:46]}
		// o5 := AccountByYears{Year: year4,Total: years[49].Account,Data: years[46:50]}
		result := []AccountByYears{o1, o2, o3, o4}
		return c.JSON(result)
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
	} else if y1 == "2018" && y2 == "2020" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		time1 := today.Year() - 3
		time2 := today.Year() - 2
		year := strconv.Itoa(time)
		year1 := strconv.Itoa(time1)
		year2 := strconv.Itoa(time2)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time_2 := today.Year() - 2
		year_2 := strconv.Itoa(time_2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year_2 + month2 + day2
		request.Date_to = time_back2
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 0) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		o1 := AccountByYears{Year: year, Total: years[9].Account, Data: years[0:10]}
		o2 := AccountByYears{Year: year1, Total: years[21].Account, Data: years[10:22]}
		o3 := AccountByYears{Year: year2, Total: years[33].Account, Data: years[22:34]}
		result := []AccountByYears{o1, o2, o3}
		return c.JSON(result)
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
	} else if y1 == "2018" && y2 == "2019" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		time1 := today.Year() - 3
		year1 := strconv.Itoa(time1)
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 3
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 0) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		o1 := AccountByYears{Year: year, Total: years[9].Account, Data: years[0:10]}
		o2 := AccountByYears{Year: year1, Total: years[21].Account, Data: years[10:22]}
		result := []AccountByYears{o1, o2}
		return c.JSON(result)
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
	} else if y1 == "2018" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 4
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 0) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		o1 := AccountByYears{Year: year, Total: years[9].Account, Data: years[0:10]}
		result := []AccountByYears{o1}
		return c.JSON(result)
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
	} else if y1 == "2019" && y2 == "2021" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 3 //Date(-5,0,0).Format("2006-01-01")
		time1 := today.Year() - 2
		time2 := today.Year() - 1
		// time3 := today.Year()
		year := strconv.Itoa(time)
		year1 := strconv.Itoa(time1)
		year2 := strconv.Itoa(time2)
		// year3 := strconv.Itoa(time3)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time_2 := today.Year() - 1
		year_2 := strconv.Itoa(time_2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year_2 + month2 + day2
		request.Date_to = time_back2
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 9136) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		o2 := AccountByYears{Year: year, Total: years[11].Account, Data: years[0:12]}
		o3 := AccountByYears{Year: year1, Total: years[23].Account, Data: years[12:24]}
		o4 := AccountByYears{Year: year2, Total: years[35].Account, Data: years[24:36]}

		result := []AccountByYears{o2, o3, o4}
		return c.JSON(result)
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
	} else if y1 == "2019" && y2 == "2020" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 3 //Date(-5,0,0).Format("2006-01-01")
		time1 := today.Year() - 2
		year := strconv.Itoa(time)
		year1 := strconv.Itoa(time1)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 2
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 9136) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		o2 := AccountByYears{Year: year, Total: years[11].Account, Data: years[0:12]}
		o3 := AccountByYears{Year: year1, Total: years[23].Account, Data: years[12:24]}
		result := []AccountByYears{o2, o3}
		return c.JSON(result)
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
	} else if y1 == "2019" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 3 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 3
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 9136) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		o2 := AccountByYears{Year: year, Total: years[11].Account, Data: years[0:12]}
		return c.JSON(o2)
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
	} else if y1 == "2020" && y2 == "2021" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 2 //Date(-5,0,0).Format("2006-01-01")
		time1 := today.Year() - 1
		year := strconv.Itoa(time)
		year1 := strconv.Itoa(time1)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 1
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 20465) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		o3 := AccountByYears{Year: year, Total: years[11].Account, Data: years[0:12]}
		o4 := AccountByYears{Year: year1, Total: years[23].Account, Data: years[12:24]}
		result := []AccountByYears{o3, o4}
		return c.JSON(result)
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
	} else if y1 == "2020" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 2 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 2
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 20465) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		o3 := AccountByYears{Year: year, Total: years[11].Account, Data: years[0:12]}
		result := []AccountByYears{o3}
		return c.JSON(result)
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
	} else if y1 == "2021" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 1 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 1
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []MonthAccount{}
		query := "select monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 39610) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		o4 := AccountByYears{Year: year, Total: years[11].Account, Data: years[0:12]}

		return c.JSON(o4)
		// return c.JSON(fiber.Map{
		// 	"results": years,
		// 	"status":    "success",
		// })
	}
	return fiber.ErrBadGateway
}

func YearlyTotalAccount(c *fiber.Ctx) error {
	years := []YearAccount{}
	query := "select year, @total := @total + account as account_count from counteryearlyaccount, (Select @total := 0) as total"
	err := db.Select(&years, query)

	if err != nil {
		return err
	}

	return fiber.ErrBadRequest
}

func YearlyAllAccountGather(c *fiber.Ctx) error {
	y1 := c.Params("y1")
	y2 := c.Params("y2")
	if y1 == "" {
		years := []YearAccount{}
		query := "select year, @total := @total + account as account_count from counteryearlyaccount, (Select @total := 0) as total"
		err := db.Select(&years, query)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2018" && y2 == "2022" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 0) as total where date between date(?) and date(now());" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2019" && y2 == "2022" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() - 3 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 9136) as total where date between date(?) and date(now());" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2020" && y2 == "2022" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() - 2 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 20465) as total where date between date(?) and date(now());" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2021" && y2 == "2022" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() - 1 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 39610) as total where date between date(?) and date(now());" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2022" && y2 == "2022" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 94165) as total where date between date(?) and date(now());" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2018" && y2 == "2021" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 1
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 0) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2018" && y2 == "2020" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 2
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 0) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2018" && y2 == "2019" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 3
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 0) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2018" && y2 == "2018" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 4
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 0) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2019" && y2 == "2021" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 3 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 1
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 9136) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2019" && y2 == "2020" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 3 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 2
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 9136) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2019" && y2 == "2019" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 3 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 3
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 9136) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2020" && y2 == "2021" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 2 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 1
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 20465) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2020" && y2 == "2020" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 2 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 2
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 20465) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	} else if y1 == "2021" && y2 == "2021" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 1 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 1
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		years := []YearMonthAccount2{}
		query := "select year,monthname(str_to_date(month,'%m')) as month, @total := @total + account as account_count from counteraccount, (Select @total := 39610) as total where date between date(?) and date(?);" //,monthname(str_to_date(month,'%s'))
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"results": years,
			"status":  "success",
		})
	}
	return fiber.ErrBadGateway
}

func PercentUsersYearlyArray(c *fiber.Ctx) error {
	var err error
	percents := []PercentOfUserYearly{}
	query := "select year(created_at) year, (count(id) * 100.0)/(select count(id) from account) percent from account group by year(created_at) having percent > 0 order by percent desc"
	err = db.Select(&percents, query)
	if err != nil {
		return err
	}
	return c.JSON(percents)
}

func TotalAccount(c *fiber.Ctx) error {
	var err error
	total := AllAccount{}
	//select count(id) as account_count from account where date(created_at);
	query := "select count(id) as account_all from account where date(created_at)"
	err = db.Get(&total, query)
	if err != nil {
		return err
	}
	return c.JSON(total)
}

func TotalAccountArray(c *fiber.Ctx) error {
	var err error
	totals := []AllAccount{}
	//select count(id) as account_count from account where date(created_at);
	query := "select count(id) as account_all from account where date(created_at)"
	err = db.Select(&totals, query)
	if err != nil {
		return err
	}
	return c.JSON(totals)
}

func YearlyAccountArray(c *fiber.Ctx) error {
	var err error
	year := []YearAccount{}
	query := "select year(created_at) year, count(id) account_count from account group by year(created_at) order by year(created_at) asc"
	err = db.Select(&year, query)
	if err != nil {
		return err
	}
	return c.JSON(year)
}

func YearlyAccountGather(c *fiber.Ctx) error {
	y := c.Params("y")
	if y == "" {
		return fiber.ErrBadRequest
	} else if y == "5" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		time_cur := today.Year()
		back_year := strconv.Itoa(time_cur)
		// back_year1 :=  c.Query("back_year1");
		// back_year1 = back_year
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s-%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(now())", back_year, year)
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "4" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() - 3 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		time_cur := today.Year()
		back_year := strconv.Itoa(time_cur)
		// back_year1 :=  c.Query("back_year1");
		// back_year1 = back_year
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s-%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(now())", back_year, year)
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "3" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() - 2 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		time_cur := today.Year()
		back_year := strconv.Itoa(time_cur)
		// back_year1 :=  c.Query("back_year1");
		// back_year1 = back_year
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s-%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(now())", back_year, year)
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() - 1 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		time_cur := today.Year()
		back_year := strconv.Itoa(time_cur)
		// back_year1 :=  c.Query("back_year1");
		// back_year1 = back_year
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s-%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(now())", back_year, year)
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "1" {
		var err error
		request := Date1{}
		today := time.Now()
		time := today.Year() //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		// back_year1 :=  c.Query("back_year1");
		// back_year1 = back_year
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(now())", year)
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(years)
	}
	return fiber.ErrBadRequest
}

func YearlyAccountQueryArray(c *fiber.Ctx) error {
	var err error
	year := []YearAccount{}
	query := "select year(created_at) year, count(id) account_count from account"

	if y := c.Query("y"); y == "" {
		query = fmt.Sprintf("%s group by year(created_at) order by year(created_at) asc", query)
	}
	if y := c.Query("y"); y != "" {
		query = fmt.Sprintf("%s where year(created_at) like '%s' group by year(created_at) order by year(created_at) asc", query, y)
	}
	err = db.Select(&year, query)
	if err != nil {
		return err
	}
	return c.JSON(year)
}

func YearlyAccountQueryMap(c *fiber.Ctx) error {
	// request1 := Date2{}
	// y := request1.Year_back
	y := c.Params("y")
	// if err != nil { //ParamsInt
	// 	return fiber.ErrBadRequest
	// }
	if y == "" {
		request := Date1{}
		var err error
		// err = c.BodyParser(&request)
		// if err != nil {
		// 	return err
		// }
		today := time.Now()
		time := today.Year() //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		years := []YearMonthAccount{}
		query := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc"
		// query := "select year(created_at) year,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by year(created_at) order by year(created_at) asc"
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"last_year": years,
			"status":    "success",
		})
		// return c.JSON(years)
		// return fiber.ErrUnprocessableEntity
	}
	if y == "5" {
		today := time.Now()
		time5 := today.Year() - 5 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time5)
		month := "-01"
		day := "-01"
		time_back_five := year + month + day
		y = time_back_five
		years := []YearAccount{}
		query := "select year(created_at) year,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by year(created_at) order by year(created_at) asc"
		err := db.Select(&years, query, y)
		if err != nil {
			return err
		}
		request := Date1{}
		// err = c.BodyParser(&request)
		// if err != nil {
		// 	return err
		// }
		today_cur := time.Now()
		time_cur := today_cur.Year() //Date(-5,0,0).Format("2006-01-01")
		year_cur := strconv.Itoa(time_cur)
		month_cur := "-01"
		day_cur := "-01"
		time_back := year_cur + month_cur + day_cur
		request.Last_year = time_back
		years2 := []YearMonthAccount{}
		query2 := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc;"
		err = db.Select(&years2, query2, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"year_back": years,
			"last_year": years2,
			"status":    "success",
		})

	} else if y == "4" {
		today := time.Now()
		time4 := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time4)
		month := "-01"
		day := "-01"
		time_back_four := year + month + day
		y = time_back_four
		years := []YearAccount{}
		query := "select year(created_at) year,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by year(created_at) order by year(created_at) asc"
		err := db.Select(&years, query, y)
		if err != nil {
			return err
		}
		request := Date1{}
		// err = c.BodyParser(&request)
		// if err != nil {
		// 	return err
		// }
		today_cur := time.Now()
		time_cur := today_cur.Year() //Date(-5,0,0).Format("2006-01-01")
		year_cur := strconv.Itoa(time_cur)
		month_cur := "-01"
		day_cur := "-01"
		time_back := year_cur + month_cur + day_cur
		request.Last_year = time_back
		years2 := []YearMonthAccount{}
		query2 := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc;"
		err = db.Select(&years2, query2, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"year_back": years,
			"last_year": years2,
			"status":    "success",
		})

	} else if y == "3" {
		today := time.Now()
		time3 := today.Year() - 3 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time3)
		month := "-01"
		day := "-01"
		time_back_three := year + month + day
		y = time_back_three
		years := []YearAccount{}
		query := "select year(created_at) year,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by year(created_at) order by year(created_at) asc"
		err := db.Select(&years, query, y)
		if err != nil {
			return err
		}
		request := Date1{}
		// err = c.BodyParser(&request)
		// if err != nil {
		// 	return err
		// }
		today_cur := time.Now()
		time_cur := today_cur.Year() //Date(-5,0,0).Format("2006-01-01")
		year_cur := strconv.Itoa(time_cur)
		month_cur := "-01"
		day_cur := "-01"
		time_back := year_cur + month_cur + day_cur
		request.Last_year = time_back
		years2 := []YearMonthAccount{}
		query2 := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc;"
		err = db.Select(&years2, query2, request.Last_year)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"year_back": years,
			"last_year": years2,
			"status":    "success",
		})

	} else if y == "2" {
		today := time.Now()
		time2 := today.Year() - 2 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time2)
		month := "-01"
		day := "-01"
		time_back_two := year + month + day
		y = time_back_two
		years := []YearAccount{}
		query := "select year(created_at) year,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by year(created_at) order by year(created_at) asc"
		err := db.Select(&years, query, y)
		if err != nil {
			return err
		}
		request := Date1{}
		// err = c.BodyParser(&request)
		// if err != nil {
		// 	return err
		// }
		today_cur := time.Now()
		time_cur := today_cur.Year() //Date(-5,0,0).Format("2006-01-01")
		year_cur := strconv.Itoa(time_cur)
		month_cur := "-01"
		day_cur := "-01"
		time_back := year_cur + month_cur + day_cur
		request.Last_year = time_back
		years2 := []YearMonthAccount{}
		query2 := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc;"
		err = db.Select(&years2, query2, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"year_back": years,
			"last_year": years2,
			"status":    "success",
		})
	} else if y == "1" {
		today := time.Now()
		time1 := today.Year() - 1 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time1)
		month := "-01"
		day := "-01"
		time_back_one := year + month + day
		y = time_back_one
		years := []YearAccount{}
		query := "select year(created_at) year,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by year(created_at) order by year(created_at) asc"
		err := db.Select(&years, query, y)
		if err != nil {
			return err
		}
		request := Date1{}
		// err = c.BodyParser(&request)
		// if err != nil {
		// 	return err
		// }
		today_cur := time.Now()
		time_cur := today_cur.Year() //Date(-5,0,0).Format("2006-01-01")
		year_cur := strconv.Itoa(time_cur)
		month_cur := "-01"
		day_cur := "-01"
		time_back := year_cur + month_cur + day_cur
		request.Last_year = time_back
		years2 := []YearMonthAccount{}
		query2 := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc;"
		err = db.Select(&years2, query2, request.Last_year)
		if err != nil {
			return err
		}
		// var code int
		// // Retrieve the custom status code if it's an fiber.*Error
		// if e, ok := err.(*fiber.Error); ok {
		//     code = e.Code
		// }

		// // Send custom error page
		// err = c.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
		// if err != nil {
		//     // In case the SendFile fails
		//     return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		// }
		return c.JSON(fiber.Map{
			"year_back": years,
			"last_year": years2,
			"status":    "success",
		})
	} else if y == "0" {
		request := Date1{}
		var err error
		// err = c.BodyParser(&request)
		// if err != nil {
		// 	return err
		// }
		today := time.Now()
		time := today.Year()
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Last_year = time_back
		years := []YearMonthAccount{}
		query := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc;"
		// query := "select year(created_at) year,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by year(created_at) order by year(created_at) asc"
		err = db.Select(&years, query, request.Last_year)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"last_year": years,
			"status":    "success",
		})
	} else {
		years := []YearAccount{}
		query := "select year(created_at) year,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by year(created_at) order by year(created_at) asc"
		err := db.Select(&years, query, y)
		if err != nil {
			return err
		}
		return fiber.ErrBadRequest
	}
}

func YearlyMonthlyAccountGather(c *fiber.Ctx) error {
	y := c.Params("y")
	m := c.Params("m")
	if y == "" {
		return fiber.ErrBadRequest
	} else if y == "2018" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 4
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2019" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 3
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2020" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 2
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2021" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year() - 1
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "12" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-12"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "11" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-11"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "10" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-10"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "09" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-09"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "08" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-08"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "07" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-07"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "06" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-06"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "05" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-05"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "04" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-04"
		day2 := "-30"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "03" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-03"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "02" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-02"
		day2 := "-28"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	} else if y == "2022" && m == "01" {
		var err error
		request := Date{}
		today := time.Now()
		time := today.Year() - 4 //Date(-5,0,0).Format("2006-01-01")
		year := strconv.Itoa(time)
		month := "-01"
		day := "-01"
		time_back := year + month + day
		request.Date_from = time_back
		time2 := today.Year()
		year2 := strconv.Itoa(time2)
		month2 := "-01"
		day2 := "-31"
		time_back2 := year2 + month2 + day2
		request.Date_to = time_back2
		// time_cur  := today.Year()
		// back_year := strconv.Itoa(time_cur)
		years := []YearAccount{}
		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)", year2)
		err = db.Select(&years, query, request.Date_from, request.Date_to)
		if err != nil {
			return err
		}
		return c.JSON(years)
	}
	return fiber.ErrBadRequest
}

// func YearlyAccount2(c *fiber.Ctx) error {
// 	request := Date2{}
// 	var err error
// 	err = c.BodyParser(&request)
// 	if err != nil {
// 		return err
// 	}
// 	if request.Year_back == "" {
// 		request := Date1{}
// 		var err error
// 		err = c.BodyParser(&request)
// 		if err != nil {
// 			return err
// 		}
// 		// if request.Last_year == "" {
// 		// 	return fiber.ErrUnprocessableEntity
// 		// }
// 		today := time.Now()
// 		time := today.Year()  //Date(-5,0,0).Format("2006-01-01")
// 		year := strconv.Itoa(time)
// 		month := "-01"
// 		day := "-01"
// 		time_back := year+month+day
// 		request.Last_year = time_back
// 		years := []YearMonthAccount{}
// 		query := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc"
// 		// query := "select year(created_at) year,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by year(created_at) order by year(created_at) asc"
// 		err = db.Select(&years, query, request.Last_year)
// 		if err != nil {
// 			return err
// 		}
// 		for _,y := range years {
// 			fmt.Printf("%s\t %s\t %d\t \n",y.Year,y.Month,y.Yearback)
// 		}
// 		x := years[0].Yearback
// 		fmt.Printf("%d\n",x)
// 		return c.JSON(fiber.Map{
// 			"last_year": years,
// 			// "status":    "success",
// 		})
// 		// return c.JSON(years)
// 		// return fiber.ErrUnprocessableEntity
// 	}else if request.Year_back == "4" {
// 		var err error
// 		request := Date{}
// 		today := time.Now()
// 		time := today.Year()-4 //Date(-5,0,0).Format("2006-01-01")
// 		year := strconv.Itoa(time)
// 		month := "-01"
// 		day := "-01"
// 		time_back := year+month+day
// 		request.Date_from = time_back
// 		time2 := today.Year()
// 		year2 := strconv.Itoa(time2)
// 		month2 := "-12"
// 		day2 := "-31"
// 		time_back2 := year2+month2+day2
// 		request.Date_to = time_back2
// 		// time_cur  := today.Year()
// 		// back_year := strconv.Itoa(time_cur)
// 		years := []YearAccount{}
// 		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)",year2)
// 		err = db.Select(&years, query ,request.Date_from ,request.Date_to)
// 		if err != nil {
// 			return err
// 		}
// 		// request := Date1{}
// 		// var err error
// 		// err = c.BodyParser(&request)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		// today_cur := time.Now()
// 		// time_cur := today_cur.Year()  //Date(-5,0,0).Format("2006-01-01")
// 		// year_cur := strconv.Itoa(time_cur)
// 		// month_cur := "-01"
// 		// day_cur := "-01"
// 		// time_back := year_cur+month_cur+day_cur
// 		// request.Last_year = time_back
// 		// years2 := []YearMonthAccount{}
// 		// query2 := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc;"
// 		// err = db.Select(&years2, query2, request.Last_year)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		// var code int
//         // // Retrieve the custom status code if it's an fiber.*Error
//         // if e, ok := err.(*fiber.Error); ok {
//         //     code = e.Code
//         // }

//         // // Send custom error page
//         // err = c.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
//         // if err != nil {
//         //     // In case the SendFile fails
//         //     return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
//         // }
// 		return c.JSON(fiber.Map{
// 			"year_back": years,
// 			// "last_year": years2,
// 			"status":    "success",
// 		})

// 	} else if request.Year_back == "3" {
// 		var err error
// 		request := Date{}
// 		today := time.Now()
// 		time := today.Year()-4 //Date(-5,0,0).Format("2006-01-01")
// 		year := strconv.Itoa(time)
// 		month := "-01"
// 		day := "-01"
// 		time_back := year+month+day
// 		request.Date_from = time_back
// 		time2 := today.Year()-1
// 		year2 := strconv.Itoa(time2)
// 		month2 := "-12"
// 		day2 := "-31"
// 		time_back2 := year2+month2+day2
// 		request.Date_to = time_back2
// 		// time_cur  := today.Year()
// 		// back_year := strconv.Itoa(time_cur)
// 		years := []YearAccount{}
// 		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)",year2)
// 		err = db.Select(&years, query ,request.Date_from ,request.Date_to)
// 		if err != nil {
// 			return err
// 		}
// 		// request := Date1{}
// 		// var err error
// 		// err = c.BodyParser(&request)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		// today_cur := time.Now()
// 		// time_cur := today_cur.Year()  //Date(-5,0,0).Format("2006-01-01")
// 		// year_cur := strconv.Itoa(time_cur)
// 		// month_cur := "-01"
// 		// day_cur := "-01"
// 		// time_back := year_cur+month_cur+day_cur
// 		// request.Last_year = time_back
// 		// years2 := []YearMonthAccount{}
// 		// query2 := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc;"
// 		// err = db.Select(&years2, query2, request.Last_year)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		return c.JSON(fiber.Map{
// 			"year_back": years,
// 			// "last_year": years2,
// 			"status":    "success",
// 		})

// 	} else if request.Year_back == "2" {
// 		var err error
// 		request := Date{}
// 		today := time.Now()
// 		time := today.Year()-4 //Date(-5,0,0).Format("2006-01-01")
// 		year := strconv.Itoa(time)
// 		month := "-01"
// 		day := "-01"
// 		time_back := year+month+day
// 		request.Date_from = time_back
// 		time2 := today.Year()-2
// 		year2 := strconv.Itoa(time2)
// 		month2 := "-12"
// 		day2 := "-31"
// 		time_back2 := year2+month2+day2
// 		request.Date_to = time_back2
// 		// time_cur  := today.Year()
// 		// back_year := strconv.Itoa(time_cur)
// 		years := []YearAccount{}
// 		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)",year2)
// 		err = db.Select(&years, query ,request.Date_from ,request.Date_to)
// 		if err != nil {
// 			return err
// 		}
// 		// request := Date1{}
// 		// var err error
// 		// err = c.BodyParser(&request)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		// today_cur := time.Now()
// 		// time_cur := today_cur.Year()  //Date(-5,0,0).Format("2006-01-01")
// 		// year_cur := strconv.Itoa(time_cur)
// 		// month_cur := "-01"
// 		// day_cur := "-01"
// 		// time_back := year_cur+month_cur+day_cur
// 		// request.Last_year = time_back
// 		// years2 := []YearMonthAccount{}
// 		// query2 := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc;"
// 		// err = db.Select(&years2, query2, request.Last_year)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		return c.JSON(fiber.Map{
// 			"year_back": years,
// 			// "last_year": years2,
// 			"status":    "success",
// 		})
// 	} else if request.Year_back == "1" {
// 		var err error
// 		request := Date{}
// 		today := time.Now()
// 		time := today.Year()-4 //Date(-5,0,0).Format("2006-01-01")
// 		year := strconv.Itoa(time)
// 		month := "-01"
// 		day := "-01"
// 		time_back := year+month+day
// 		request.Date_from = time_back
// 		time2 := today.Year()-3
// 		year2 := strconv.Itoa(time2)
// 		month2 := "-12"
// 		day2 := "-31"
// 		time_back2 := year2+month2+day2
// 		request.Date_to = time_back2
// 		// time_cur  := today.Year()
// 		// back_year := strconv.Itoa(time_cur)
// 		years := []YearAccount{}
// 		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)",year2)
// 		err = db.Select(&years, query ,request.Date_from ,request.Date_to)
// 		if err != nil {
// 			return err
// 		}
// 		// request := Date1{}
// 		// var err error
// 		// err = c.BodyParser(&request)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		// today_cur := time.Now()
// 		// time_cur := today_cur.Year()  //Date(-5,0,0).Format("2006-01-01")
// 		// year_cur := strconv.Itoa(time_cur)
// 		// month_cur := "-01"
// 		// day_cur := "-01"
// 		// time_back := year_cur+month_cur+day_cur
// 		// request.Last_year = time_back
// 		// years2 := []YearMonthAccount{}
// 		// query2 := "select year(created_at) year,month(created_at) month,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by month(created_at) order by month(created_at) asc;"
// 		// err = db.Select(&years2, query2, request.Last_year)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		return c.JSON(fiber.Map{
// 			"year_back": years,
// 			// "last_year": years2,
// 			"status":    "success",
// 		})
// 	}else if request.Year_back == "0" {
// 		var err error
// 		request := Date{}
// 		today := time.Now()
// 		time := today.Year()-4 //Date(-5,0,0).Format("2006-01-01")
// 		year := strconv.Itoa(time)
// 		month := "-01"
// 		day := "-01"
// 		time_back := year+month+day
// 		request.Date_from = time_back
// 		time2 := today.Year()-4
// 		year2 := strconv.Itoa(time2)
// 		month2 := "-12"
// 		day2 := "-31"
// 		time_back2 := year2+month2+day2
// 		request.Date_to = time_back2
// 		// time_cur  := today.Year()
// 		// back_year := strconv.Itoa(time_cur)
// 		years := []YearAccount{}
// 		query := fmt.Sprintf("select '%s' as year, count(id) as account_count from account where date(created_at) between date(?) and date(?)",year2)
// 		err = db.Select(&years, query ,request.Date_from ,request.Date_to)
// 		if err != nil {
// 			return err
// 		}
// 		// var code int
//         // // Retrieve the custom status code if it's an fiber.*Error
//         // if e, ok := err.(*fiber.Error); ok {
//         //     code = e.Code
//         // }

//         // // Send custom error page
//         // err = c.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
//         // if err != nil {
//         //     // In case the SendFile fails
//         //     return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
//         // }
// 		return c.JSON(fiber.Map{
// 			"last_year": years,
// 			"status":    "success",
// 		})
// 	}else{
// 		years := []YearAccount{}
// 		query := "select year(created_at) year,count(id) as account_count from account where date(created_at) between date(?) and date(now()) group by year(created_at) order by year(created_at) asc"
// 		err = db.Select(&years, query, request.Year_back)
// 		if err != nil {
// 			return err
// 		}
// 		// var code int
//         // // Retrieve the custom status code if it's an fiber.*Error
//         // if e, ok := err.(*fiber.Error); ok {
//         //     code = e.Code
//         // }

//         // // Send custom error page
//         // err = c.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
//         // if err != nil {
//         //     // In case the SendFile fails
//         //     return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
//         // }
// 		// return c.JSON(fiber.Map{
// 		// 	"year_back": years,
// 		// 	"status":    "success",
// 		// })
// 		return fiber.ErrUnprocessableEntity
//     }
// }
