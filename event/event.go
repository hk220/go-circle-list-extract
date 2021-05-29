package event

type Event struct {
	CircleListURL string `mapstructure:"circle_list_url"`
	Parser        string `mapstructure:"parser"`
}
