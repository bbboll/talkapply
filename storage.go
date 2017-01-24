package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/sha3"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Token struct {
	Value string
}

func (t *Token) Verify(s string) error {
	if t.Value == hashStringForToken(s) {
		return nil
	}
	return errors.New("Hash verification failed.")
}

// Generates a new random admin token. Returns the token object and a plain un-hased version of the token
func NewToken() (*Token, string) {

	bytes := make([]byte, 10)
	for i := 0; i < 10; i++ {
		bytes[i] = byte(rand.Intn(57-48) + 48)
	}

	t := &Token{}
	t.Value = hashStringForToken(string(bytes))
	return t, string(bytes)
}

func hashStringForToken(s string) string {
	h := sha3.New512()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

type Server struct {
	adminTokens []*Token
	Projects    []*Project `json:"projects"`
	m           sync.Mutex
}

// Adds a new admin token to the server and returns the un-hashed version
func (s *Server) AddToken() string {
	t, p := NewToken()
	s.adminTokens = append(s.adminTokens, t)
	return p
}

// Adds a new project to the server
func (s *Server) AddProject(title string) *Project {
	p := NewProject(title)
	s.Projects = append(s.Projects, p)
	return p
}

type Project struct {
	Title        string
	Groupings    []*Grouping `json:"groupings"`
	AccessNumber uint64
}

func NewProject(title string) *Project {
	return &Project{
		Title: title,
	}
}

func (p *Project) AddGrouping(title string) *Grouping {
	g := NewGrouping(title)
	p.Groupings = append(p.Groupings, g)
	return g
}

type Grouping struct {
	Title     string
	Subjects  []*Subject `json:"subjects"`
	SortIndex uint8
}

func NewGrouping(title string) *Grouping {
	return &Grouping{
		Title: title,
	}
}

func (g *Grouping) AddSubject(title string) *Subject {
	s := NewSubject(title)
	g.Subjects = append(g.Subjects, s)
	return s
}

type Subject struct {
	Title            string `json:"title"`
	People           string `json:"people"`
	Locked           bool   `json:"locked"`
	lastModification uint64 `json:"lastModification"`
}

func NewSubject(title string) *Subject {
	return &Subject{
		Title:  title,
		Locked: false,
	}
}

type Storage struct {
	Server   *Server
	filename string
	seconds  int
}

func initStorage(seconds int, filename string) *Storage {

	rand.Seed(time.Now().UTC().UnixNano())
	strg := &Storage{
		Server:   &Server{},
		filename: filename,
		seconds:  seconds,
	}

	if strg.crashed() {
		fmt.Printf("talkapply crashed and will now start restoring the data")

		j, err := ioutil.ReadFile(filename)

		// if the file is broken someone fucked up
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(100)
		}

		err = json.Unmarshal(j, strg.Server)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(101)
		}
	}

	go strg.loop()
	return nil
}

// checks if the talkapply crashed and needs to restore the data to the ram
func (strg *Storage) crashed() bool {
	if _, err := os.Stat(strg.filename); err == nil {
		return true
	}
	return false
}

// save the in memory values to json file
func (strg *Storage) save() {

	strg.Server.m.Lock()
	j, err := json.Marshal(strg.Server)

	if err != nil {
		fmt.Println("Failed to save storage to disk. \n%s", err)
	}

	err = ioutil.WriteFile(strg.filename, j, 0644)

	if err != nil {
		fmt.Println("Failed to save storage to disk. \n%s", err)
	}

	strg.Server.m.Unlock()
	fmt.Println("Saved memory to disk.")
}

func (strg *Storage) loop() {
	for {
		strg.save()
		time.Sleep(time.Duration(strg.seconds) * time.Second)
	}
}
