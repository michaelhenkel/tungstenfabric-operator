// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gapic-generator. DO NOT EDIT.

package talent

import (
	"context"
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> v0.0.4
	"math"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// JobCallOptions contains the retry settings for each method of JobClient.
type JobCallOptions struct {
	CreateJob          []gax.CallOption
	GetJob             []gax.CallOption
	UpdateJob          []gax.CallOption
	DeleteJob          []gax.CallOption
	ListJobs           []gax.CallOption
	BatchDeleteJobs    []gax.CallOption
	SearchJobs         []gax.CallOption
	SearchJobsForAlert []gax.CallOption
}

func defaultJobClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("jobs.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultJobCallOptions() *JobCallOptions {
	retry := map[[2]string][]gax.CallOption{
		{"default", "idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
	}
	return &JobCallOptions{
		CreateJob:          retry[[2]string{"default", "non_idempotent"}],
		GetJob:             retry[[2]string{"default", "idempotent"}],
		UpdateJob:          retry[[2]string{"default", "non_idempotent"}],
		DeleteJob:          retry[[2]string{"default", "idempotent"}],
		ListJobs:           retry[[2]string{"default", "idempotent"}],
		BatchDeleteJobs:    retry[[2]string{"default", "non_idempotent"}],
		SearchJobs:         retry[[2]string{"default", "non_idempotent"}],
		SearchJobsForAlert: retry[[2]string{"default", "non_idempotent"}],
	}
}

// JobClient is a client for interacting with Cloud Talent Solution API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type JobClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	jobClient talentpb.JobServiceClient

	// The call options for this service.
	CallOptions *JobCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewJobClient creates a new job service client.
//
// A service handles job management, including job CRUD, enumeration and search.
func NewJobClient(ctx context.Context, opts ...option.ClientOption) (*JobClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultJobClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &JobClient{
		conn:        conn,
		CallOptions: defaultJobCallOptions(),

		jobClient: talentpb.NewJobServiceClient(conn),
	}
	c.setGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *JobClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *JobClient) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *JobClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateJob creates a new job.
//
// Typically, the job becomes searchable within 10 seconds, but it may take
// up to 5 minutes.
func (c *JobClient) CreateJob(ctx context.Context, req *talentpb.CreateJobRequest, opts ...gax.CallOption) (*talentpb.Job, error) {
<<<<<<< HEAD
	ctx = insertMetadata(ctx, c.xGoogMetadata)
=======
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
>>>>>>> v0.0.4
	opts = append(c.CallOptions.CreateJob[0:len(c.CallOptions.CreateJob):len(c.CallOptions.CreateJob)], opts...)
	var resp *talentpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobClient.CreateJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetJob retrieves the specified job, whose status is OPEN or recently EXPIRED
// within the last 90 days.
func (c *JobClient) GetJob(ctx context.Context, req *talentpb.GetJobRequest, opts ...gax.CallOption) (*talentpb.Job, error) {
<<<<<<< HEAD
	ctx = insertMetadata(ctx, c.xGoogMetadata)
=======
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
>>>>>>> v0.0.4
	opts = append(c.CallOptions.GetJob[0:len(c.CallOptions.GetJob):len(c.CallOptions.GetJob)], opts...)
	var resp *talentpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobClient.GetJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateJob updates specified job.
//
// Typically, updated contents become visible in search results within 10
// seconds, but it may take up to 5 minutes.
func (c *JobClient) UpdateJob(ctx context.Context, req *talentpb.UpdateJobRequest, opts ...gax.CallOption) (*talentpb.Job, error) {
<<<<<<< HEAD
	ctx = insertMetadata(ctx, c.xGoogMetadata)
=======
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "job.name", req.GetJob().GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
>>>>>>> v0.0.4
	opts = append(c.CallOptions.UpdateJob[0:len(c.CallOptions.UpdateJob):len(c.CallOptions.UpdateJob)], opts...)
	var resp *talentpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobClient.UpdateJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteJob deletes the specified job.
//
// Typically, the job becomes unsearchable within 10 seconds, but it may take
// up to 5 minutes.
func (c *JobClient) DeleteJob(ctx context.Context, req *talentpb.DeleteJobRequest, opts ...gax.CallOption) error {
<<<<<<< HEAD
	ctx = insertMetadata(ctx, c.xGoogMetadata)
=======
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
>>>>>>> v0.0.4
	opts = append(c.CallOptions.DeleteJob[0:len(c.CallOptions.DeleteJob):len(c.CallOptions.DeleteJob)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.jobClient.DeleteJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ListJobs lists jobs by filter.
func (c *JobClient) ListJobs(ctx context.Context, req *talentpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
<<<<<<< HEAD
	ctx = insertMetadata(ctx, c.xGoogMetadata)
=======
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
>>>>>>> v0.0.4
	opts = append(c.CallOptions.ListJobs[0:len(c.CallOptions.ListJobs):len(c.CallOptions.ListJobs)], opts...)
	it := &JobIterator{}
	req = proto.Clone(req).(*talentpb.ListJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*talentpb.Job, string, error) {
		var resp *talentpb.ListJobsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.jobClient.ListJobs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.Jobs, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// BatchDeleteJobs deletes a list of [Job][google.cloud.talent.v4beta1.Job]s by filter.
func (c *JobClient) BatchDeleteJobs(ctx context.Context, req *talentpb.BatchDeleteJobsRequest, opts ...gax.CallOption) error {
<<<<<<< HEAD
	ctx = insertMetadata(ctx, c.xGoogMetadata)
=======
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
>>>>>>> v0.0.4
	opts = append(c.CallOptions.BatchDeleteJobs[0:len(c.CallOptions.BatchDeleteJobs):len(c.CallOptions.BatchDeleteJobs)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.jobClient.BatchDeleteJobs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// SearchJobs searches for jobs using the provided
// [SearchJobsRequest][google.cloud.talent.v4beta1.SearchJobsRequest].
//
// This call constrains the
// [visibility][google.cloud.talent.v4beta1.Job.visibility] of jobs present in
// the database, and only returns jobs that the caller has permission to
// search against.
func (c *JobClient) SearchJobs(ctx context.Context, req *talentpb.SearchJobsRequest, opts ...gax.CallOption) *SearchJobsResponse_MatchingJobIterator {
<<<<<<< HEAD
	ctx = insertMetadata(ctx, c.xGoogMetadata)
=======
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
>>>>>>> v0.0.4
	opts = append(c.CallOptions.SearchJobs[0:len(c.CallOptions.SearchJobs):len(c.CallOptions.SearchJobs)], opts...)
	it := &SearchJobsResponse_MatchingJobIterator{}
	req = proto.Clone(req).(*talentpb.SearchJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*talentpb.SearchJobsResponse_MatchingJob, string, error) {
		var resp *talentpb.SearchJobsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.jobClient.SearchJobs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.MatchingJobs, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// SearchJobsForAlert searches for jobs using the provided
// [SearchJobsRequest][google.cloud.talent.v4beta1.SearchJobsRequest].
//
// This API call is intended for the use case of targeting passive job
// seekers (for example, job seekers who have signed up to receive email
// alerts about potential job opportunities), and has different algorithmic
// adjustments that are targeted to passive job seekers.
//
// This call constrains the
// [visibility][google.cloud.talent.v4beta1.Job.visibility] of jobs present in
// the database, and only returns jobs the caller has permission to search
// against.
func (c *JobClient) SearchJobsForAlert(ctx context.Context, req *talentpb.SearchJobsRequest, opts ...gax.CallOption) *SearchJobsResponse_MatchingJobIterator {
<<<<<<< HEAD
	ctx = insertMetadata(ctx, c.xGoogMetadata)
=======
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
>>>>>>> v0.0.4
	opts = append(c.CallOptions.SearchJobsForAlert[0:len(c.CallOptions.SearchJobsForAlert):len(c.CallOptions.SearchJobsForAlert)], opts...)
	it := &SearchJobsResponse_MatchingJobIterator{}
	req = proto.Clone(req).(*talentpb.SearchJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*talentpb.SearchJobsResponse_MatchingJob, string, error) {
		var resp *talentpb.SearchJobsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.jobClient.SearchJobsForAlert(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.MatchingJobs, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// JobIterator manages a stream of *talentpb.Job.
type JobIterator struct {
	items    []*talentpb.Job
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*talentpb.Job, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *JobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *JobIterator) Next() (*talentpb.Job, error) {
	var item *talentpb.Job
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *JobIterator) bufLen() int {
	return len(it.items)
}

func (it *JobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// SearchJobsResponse_MatchingJobIterator manages a stream of *talentpb.SearchJobsResponse_MatchingJob.
type SearchJobsResponse_MatchingJobIterator struct {
	items    []*talentpb.SearchJobsResponse_MatchingJob
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*talentpb.SearchJobsResponse_MatchingJob, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SearchJobsResponse_MatchingJobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SearchJobsResponse_MatchingJobIterator) Next() (*talentpb.SearchJobsResponse_MatchingJob, error) {
	var item *talentpb.SearchJobsResponse_MatchingJob
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SearchJobsResponse_MatchingJobIterator) bufLen() int {
	return len(it.items)
}

func (it *SearchJobsResponse_MatchingJobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
