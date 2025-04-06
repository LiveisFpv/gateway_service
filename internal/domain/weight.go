package domain

type WeightRecord struct {
	User_id int
	Weight  float32
	Date    string
}

type WeightHistory struct {
	Data *[]WeightRecord
}
