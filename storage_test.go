package main

import (
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"io/ioutil"
	"testing"
	// "os"
	"strconv"
)

func BenchmarkJsonMarshal(b *testing.B) {

	server := Server{}
	server.m.Lock()

	// add tokens
	for i := 0; i < 10; i++ {
		_ = server.AddToken()
	}

	// add projects
	for i := 0; i < 100; i++ {
		p := server.AddProject(randomdata.SillyName() + " " + randomdata.Street())

		// add groupings
		for j := 0; j < 10; j++ {
			g := p.AddGrouping(randomdata.State(randomdata.Small) + " " + randomdata.SillyName())

			// add subjects
			for k := 0; k < 10; k++ {
				s := g.AddSubject(randomdata.LastName() + " " + strconv.Itoa(randomdata.Number(13, 985)) + " " + randomdata.SillyName())

				// add people to 10% of the subject
				if randomdata.Number(0, 1000) < 100 {
					s.Locked = true
					s.People = fmt.Sprintf("%s, %s, %s", randomdata.SillyName(), randomdata.SillyName(), randomdata.SillyName())
				}
			}
		}
	}

	b.ResetTimer()

	j, err := json.Marshal(server)

	if err != nil {
		b.Error(err)
	}

	file, err := ioutil.TempFile("", "talkapply_benchmark")
	// defer os.Remove(file.Name())

	if _, err := file.Write(j); err != nil {
		b.Error(err)
	}

	server.m.Unlock()
	file.Close()

}
