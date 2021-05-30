package circle

type Circle struct {
	Space string `json:"space" csv:"space"`
	Name  string `json:"name" csv:"name"`
	URL   string `json:"url" csv:"url"`
}

type Circles []Circle

func (cl *Circles) Add(c *Circle) {
	*cl = append(*cl, *c)
}

func (cl *Circles) String() [][]string {
	var result [][]string
	for _, c := range *cl {
		result = append(result, []string{c.Space, c.Name, c.URL})
	}
	return result
}
