package msgraph_test

import (
	"fmt"
	"testing"

	"github.com/manicminer/hamilton/internal/test"
	"github.com/manicminer/hamilton/msgraph"
)

func TestSynchronizationJobTest(t *testing.T) {
	c := test.NewTest(t)
	defer c.CancelFunc()

	jobs := testSynchronizationJob_List(t, c, "e87dd102-a303-4df9-bc17-8206ea54b44c")
	fmt.Println(jobs)
}

func testSynchronizationJob_List(t *testing.T, c *test.Test, id string) (jobs *[]msgraph.SynchronizationJob) {
	jobs, _, err := c.SynchronizationJobClient.List(c.Context, id)
	if err != nil {
		t.Fatalf("SynchronizationJobClient.List(): %v", err)
	}

	if jobs == nil {
		t.Fatal("SynchronizationJobClient.List(): jobs is nil")
	}

	return
}
