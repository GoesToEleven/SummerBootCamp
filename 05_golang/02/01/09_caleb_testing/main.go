package example

type DB struct{}

func (db *DB) Get(key string) string {
	return map[string]string{
		"first_name": "John",
		"last_name":  "Doe",
	}[key]
}

type Database interface {
	Get(string) string
}

func GetName(db Database) string {
	return db.Get("first_name") + " " + db.Get("last_name")
}




package example

import (
"io/ioutil"
"os"
"strings"
"testing"
)

type FakeDatabase struct {
	GetFunc func(string) string
}
func (fake *FakeDatabase) Get(key string) string {
	return fake.GetFunc(key)
}

func TestDatabase(t *testing.T) {
	fake := &FakeDatabase{
		GetFunc: func(key string) {
			return key
		},
	}
	result := GetName(fake)
	if result != "first_name last_name" {

	}
}

func TestLengthOfFile(t *testing.T) {
	n := LengthOfFile(strings.NewReader("hello"))
	if n != 5 {
		t.Errorf("invalid length")
	}

	ioutil.WriteFile("/tmp/fake-file", []byte("hello"), 0777)
	f, _ := os.Open("/tmp/fake-file")
	defer f.Close()
	defer os.Remove("/tmp/fake-file")
	n := LengthOfFile(f)
	if n != 5 {
		t.Errorf("invalid length")
	}
}


// this code was grabbed from lecture
// and does not actually work