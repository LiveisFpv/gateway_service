package domain

type Dishes struct {
	Dishes_id   int
	User_id     int
	Dishes_name string
	Dishes_num  int
	Dishes_date string
}

type Plan_diet struct {
	Data *[]Diet
}

type Diet struct {
	Dishes_id   int
	Time        string
	Title       string
	Callory     float64
	Fat         float64
	Protein     float64
	Carbs       float64
	Description string
	Weight      float64
	Date        string
}
