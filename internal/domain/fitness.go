package domain

type Fitness struct {
	User_id       int
	Fitness_id    int
	Fitness_date  string
	Fitness_train string
}

type Plan_train struct {
	Data *[]Train
}

type Train struct {
	Train_id    int
	Title       string
	Description string
	User_level  int
	Date        string
}
