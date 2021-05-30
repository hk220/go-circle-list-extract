/*
Copyright © 2021 Kazuki Hara

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
