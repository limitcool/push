package config

type Config struct {
	Enable string
	Mirai
	PushPlus
}
type Mirai struct {
	Host   string
	Target int64
}

type PushPlus struct {
	Token string
	Title string
}
