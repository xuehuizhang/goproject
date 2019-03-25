package model

type Config struct {
	LogPath string
	LogLevel string

	ChanSize int
	KafkaAddr string
	KafkaPort int

	CollectConf []CollectConf
}
