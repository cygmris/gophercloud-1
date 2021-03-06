package testing

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/openstack/workflow/v2/executions"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreateExecution(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/executions", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusCreated)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `
			{
				"created_at": "2018-09-12 14:48:49",
				"description": "description",
				"id": "50bb59f1-eb77-4017-a77f-6d575b002667",
				"input": "{\"msg\": \"Hello\"}",
				"output": "{}",
				"params": "{\"namespace\": \"\", \"env\": {}}",
				"project_id": "778c0f25df0d492a9a868ee9e2fbb513",
				"root_execution_id": null,
				"state": "SUCCESS",
				"state_info": null,
				"task_execution_id": null,
				"updated_at": "2018-09-12 14:48:49",
				"workflow_id": "6656c143-a009-4bcb-9814-cc100a20bbfa",
				"workflow_name": "echo",
				"workflow_namespace": ""
			}
		`)
	})

	opts := &executions.CreateOpts{
		WorkflowID: "6656c143-a009-4bcb-9814-cc100a20bbfa",
		Input: map[string]interface{}{
			"msg": "Hello",
		},
		Description: "description",
	}

	actual, err := executions.Create(fake.ServiceClient(), opts).Extract()
	if err != nil {
		t.Fatalf("Unable to create execution: %v", err)
	}

	expected := &executions.Execution{
		ID:          "50bb59f1-eb77-4017-a77f-6d575b002667",
		Description: "description",
		Input: map[string]interface{}{
			"msg": "Hello",
		},
		Params: map[string]interface{}{
			"namespace": "",
			"env":       map[string]interface{}{},
		},
		Output:       map[string]interface{}{},
		ProjectID:    "778c0f25df0d492a9a868ee9e2fbb513",
		State:        "SUCCESS",
		WorkflowID:   "6656c143-a009-4bcb-9814-cc100a20bbfa",
		WorkflowName: "echo",
		CreatedAt:    time.Date(2018, time.September, 12, 14, 48, 49, 0, time.UTC),
		UpdatedAt:    time.Date(2018, time.September, 12, 14, 48, 49, 0, time.UTC),
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v, but was %#v", expected, actual)
	}
}

func TestGetExecution(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/executions/50bb59f1-eb77-4017-a77f-6d575b002667", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `
			{
				"created_at": "2018-09-12 14:48:49",
				"description": "description",
				"id": "50bb59f1-eb77-4017-a77f-6d575b002667",
				"input": "{\"msg\": \"Hello\"}",
				"output": "{}",
				"params": "{\"namespace\": \"\", \"env\": {}}",
				"project_id": "778c0f25df0d492a9a868ee9e2fbb513",
				"root_execution_id": null,
				"state": "SUCCESS",
				"state_info": null,
				"task_execution_id": null,
				"updated_at": "2018-09-12 14:48:49",
				"workflow_id": "6656c143-a009-4bcb-9814-cc100a20bbfa",
				"workflow_name": "echo",
				"workflow_namespace": ""
			}
		`)
	})

	actual, err := executions.Get(fake.ServiceClient(), "50bb59f1-eb77-4017-a77f-6d575b002667").Extract()
	if err != nil {
		t.Fatalf("Unable to get execution: %v", err)
	}

	expected := &executions.Execution{
		ID:          "50bb59f1-eb77-4017-a77f-6d575b002667",
		Description: "description",
		Input: map[string]interface{}{
			"msg": "Hello",
		},
		Params: map[string]interface{}{
			"namespace": "",
			"env":       map[string]interface{}{},
		},
		Output:       map[string]interface{}{},
		ProjectID:    "778c0f25df0d492a9a868ee9e2fbb513",
		State:        "SUCCESS",
		WorkflowID:   "6656c143-a009-4bcb-9814-cc100a20bbfa",
		WorkflowName: "echo",
		CreatedAt:    time.Date(2018, time.September, 12, 14, 48, 49, 0, time.UTC),
		UpdatedAt:    time.Date(2018, time.September, 12, 14, 48, 49, 0, time.UTC),
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v, but was %#v", expected, actual)
	}
}

func TestDeleteExecution(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/executions/1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
	res := executions.Delete(fake.ServiceClient(), "1")
	th.AssertNoErr(t, res.Err)
}
