// Code generated by github.com/jim-minter/go-cosmosdb, DO NOT EDIT.

package cosmosdb

import (
	"net/http"

	pkg "github.com/jim-minter/rp/pkg/api"
)

type leaseDocumentClient struct {
	*databaseClient
	path string
}

// LeaseDocumentClient is a leaseDocument client
type LeaseDocumentClient interface {
	Create(string, *pkg.LeaseDocument, *Options) (*pkg.LeaseDocument, error)
	List() LeaseDocumentIterator
	ListAll() (*pkg.LeaseDocuments, error)
	Get(string, string) (*pkg.LeaseDocument, error)
	Replace(string, *pkg.LeaseDocument, *Options) (*pkg.LeaseDocument, error)
	Delete(string, *pkg.LeaseDocument, *Options) error
	Query(string, *Query) LeaseDocumentIterator
	QueryAll(string, *Query) (*pkg.LeaseDocuments, error)
}

type leaseDocumentListIterator struct {
	*leaseDocumentClient
	continuation string
	done         bool
}

type leaseDocumentQueryIterator struct {
	*leaseDocumentClient
	partitionkey string
	query        *Query
	continuation string
	done         bool
}

// LeaseDocumentIterator is a leaseDocument iterator
type LeaseDocumentIterator interface {
	Next() (*pkg.LeaseDocuments, error)
}

// NewLeaseDocumentClient returns a new leaseDocument client
func NewLeaseDocumentClient(collc CollectionClient, collid string) LeaseDocumentClient {
	return &leaseDocumentClient{
		databaseClient: collc.(*collectionClient).databaseClient,
		path:           collc.(*collectionClient).path + "/colls/" + collid,
	}
}

func (c *leaseDocumentClient) all(i LeaseDocumentIterator) (*pkg.LeaseDocuments, error) {
	allleaseDocuments := &pkg.LeaseDocuments{}

	for {
		leaseDocuments, err := i.Next()
		if err != nil {
			return nil, err
		}
		if leaseDocuments == nil {
			break
		}

		allleaseDocuments.Count += leaseDocuments.Count
		allleaseDocuments.ResourceID = leaseDocuments.ResourceID
		allleaseDocuments.LeaseDocuments = append(allleaseDocuments.LeaseDocuments, leaseDocuments.LeaseDocuments...)
	}

	return allleaseDocuments, nil
}

func (c *leaseDocumentClient) Create(partitionkey string, newleaseDocument *pkg.LeaseDocument, options *Options) (leaseDocument *pkg.LeaseDocument, err error) {
	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)
	if options != nil {
		setOptions(options, headers)
	}
	err = c.do(http.MethodPost, c.path+"/docs", "docs", c.path, http.StatusCreated, &newleaseDocument, &leaseDocument, headers)
	return
}

func (c *leaseDocumentClient) List() LeaseDocumentIterator {
	return &leaseDocumentListIterator{leaseDocumentClient: c}
}

func (c *leaseDocumentClient) ListAll() (*pkg.LeaseDocuments, error) {
	return c.all(c.List())
}

func (c *leaseDocumentClient) Get(partitionkey, leaseDocumentid string) (leaseDocument *pkg.LeaseDocument, err error) {
	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)
	err = c.do(http.MethodGet, c.path+"/docs/"+leaseDocumentid, "docs", c.path+"/docs/"+leaseDocumentid, http.StatusOK, nil, &leaseDocument, headers)
	return
}

func (c *leaseDocumentClient) Replace(partitionkey string, newleaseDocument *pkg.LeaseDocument, options *Options) (leaseDocument *pkg.LeaseDocument, err error) {
	if newleaseDocument.ETag == "" {
		return nil, ErrETagRequired
	}
	headers := http.Header{}
	headers.Set("If-Match", newleaseDocument.ETag)
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)
	if options != nil {
		setOptions(options, headers)
	}
	err = c.do(http.MethodPut, c.path+"/docs/"+newleaseDocument.ID, "docs", c.path+"/docs/"+newleaseDocument.ID, http.StatusOK, &newleaseDocument, &leaseDocument, headers)
	return
}

func (c *leaseDocumentClient) Delete(partitionkey string, leaseDocument *pkg.LeaseDocument, options *Options) error {
	if leaseDocument.ETag == "" {
		return ErrETagRequired
	}
	headers := http.Header{}
	headers.Set("If-Match", leaseDocument.ETag)
	headers.Set("X-Ms-Documentdb-Partitionkey", `["`+partitionkey+`"]`)
	if options != nil {
		setOptions(options, headers)
	}
	return c.do(http.MethodDelete, c.path+"/docs/"+leaseDocument.ID, "docs", c.path+"/docs/"+leaseDocument.ID, http.StatusNoContent, nil, nil, headers)
}

func (c *leaseDocumentClient) Query(partitionkey string, query *Query) LeaseDocumentIterator {
	return &leaseDocumentQueryIterator{leaseDocumentClient: c, partitionkey: partitionkey, query: query}
}

func (c *leaseDocumentClient) QueryAll(partitionkey string, query *Query) (*pkg.LeaseDocuments, error) {
	return c.all(c.Query(partitionkey, query))
}

func (i *leaseDocumentListIterator) Next() (leaseDocuments *pkg.LeaseDocuments, err error) {
	if i.done {
		return
	}

	headers := http.Header{}
	if i.continuation != "" {
		headers.Set("X-Ms-Continuation", i.continuation)
	}

	err = i.do(http.MethodGet, i.path+"/docs", "docs", i.path, http.StatusOK, nil, &leaseDocuments, headers)
	if err != nil {
		return
	}

	i.continuation = headers.Get("X-Ms-Continuation")
	i.done = i.continuation == ""

	return
}

func (i *leaseDocumentQueryIterator) Next() (leaseDocuments *pkg.LeaseDocuments, err error) {
	if i.done {
		return
	}

	headers := http.Header{}
	headers.Set("X-Ms-Documentdb-Isquery", "True")
	headers.Set("Content-Type", "application/query+json")
	if i.partitionkey != "" {
		headers.Set("X-Ms-Documentdb-Partitionkey", `["`+i.partitionkey+`"]`)
	} else {
		headers.Set("X-Ms-Documentdb-Query-Enablecrosspartition", "True")
	}
	if i.continuation != "" {
		headers.Set("X-Ms-Continuation", i.continuation)
	}

	err = i.do(http.MethodPost, i.path+"/docs", "docs", i.path, http.StatusOK, &i.query, &leaseDocuments, headers)
	if err != nil {
		return
	}

	i.continuation = headers.Get("X-Ms-Continuation")
	i.done = i.continuation == ""

	return
}
