package circle

type Circle struct {
	Space string `json:"space" csv:"space"`
	Name  string `json:"name" csv:"name"`
	URL   string `json:"url" csv:"url"`
}

type CircleList struct {
	Circles []Circle
}

func (cl *CircleList) Add(c *Circle) {
	cl.Circles = append(cl.Circles, *c)
}

func (cl *CircleList) String() [][]string {
	var result [][]string
	for _, c := range cl.Circles {
		result = append(result, []string{c.Space, c.Name, c.URL})
	}
	return result
}
