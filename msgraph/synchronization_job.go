package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SynchronizationJobClient performs operations on Azure AD identity synchronization jobs.
type SynchronizationJobClient struct {
	BaseClient Client
}

// NewSynchronizationJobClient returns a new SynchronizationJobClient.
func NewSynchronizationJobClient(tenantId string) *SynchronizationJobClient {
	return &SynchronizationJobClient{
		BaseClient: NewClient(VersionBeta, tenantId),
	}
}

// List returns a list of synchronization jobs for the given service principal.
func (c *SynchronizationJobClient) List(ctx context.Context, svcpID string) (*[]SynchronizationJob, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		ValidStatusCodes:       []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s/synchronization/jobs", svcpID),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("SynchronizationJobClient.BaseClient.Get(): %w", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %w", err)
	}

	var data struct {
		Value []SynchronizationJob
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("Unmarshal(): %w", err)
	}

	return &data.Value, status, nil
}

// Get retrieves a synchronization job.
func (c *SynchronizationJobClient) Get(ctx context.Context, svcpID string, jobID string) (*SynchronizationJob, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		ValidStatusCodes:       []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s/synchronization/jobs/%s", svcpID, jobID),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("SynchronizationJobClient.BaseClient.Get(): %w", err)
	}

	if err != nil {
		return nil, status, fmt.Errorf("SynchronizationJobClient.BaseClient.Get(): %w", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %w", err)
	}

	var data SynchronizationJob
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("Unmarshal(): %w", err)
	}

	return &data, status, nil
}
