package freight

import (
	"context"
	"fmt"
	"testing"
)

func TestProjectResolve(t *testing.T) {

	p := &Project{
		Name:   "test",
		Remote: "github.com/ChrisMcKenzie/accord",
		Path:   "tools/accord",
	}

	pro, err := p.Resolve(context.Background(), State{
		CWD: "/tmp/test-freight",
	})

	go func() {
		for msg := range <-pro {
			fmt.Println(msg)
		}
	}()

	if err != nil {
		t.Fatal(err)
	}

}
