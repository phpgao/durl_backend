// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/phpgao/durl_backend/internal/data/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/phpgao/durl_backend/internal/data/ent/tleaf"
	"github.com/phpgao/durl_backend/internal/data/ent/tshorturl"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// TLeaf is the client for interacting with the TLeaf builders.
	TLeaf *TLeafClient
	// TShortUrl is the client for interacting with the TShortUrl builders.
	TShortUrl *TShortUrlClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.TLeaf = NewTLeafClient(c.config)
	c.TShortUrl = NewTShortUrlClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		TLeaf:     NewTLeafClient(cfg),
		TShortUrl: NewTShortUrlClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		TLeaf:     NewTLeafClient(cfg),
		TShortUrl: NewTShortUrlClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		TLeaf.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.TLeaf.Use(hooks...)
	c.TShortUrl.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.TLeaf.Intercept(interceptors...)
	c.TShortUrl.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *TLeafMutation:
		return c.TLeaf.mutate(ctx, m)
	case *TShortUrlMutation:
		return c.TShortUrl.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// TLeafClient is a client for the TLeaf schema.
type TLeafClient struct {
	config
}

// NewTLeafClient returns a client for the TLeaf from the given config.
func NewTLeafClient(c config) *TLeafClient {
	return &TLeafClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tleaf.Hooks(f(g(h())))`.
func (c *TLeafClient) Use(hooks ...Hook) {
	c.hooks.TLeaf = append(c.hooks.TLeaf, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `tleaf.Intercept(f(g(h())))`.
func (c *TLeafClient) Intercept(interceptors ...Interceptor) {
	c.inters.TLeaf = append(c.inters.TLeaf, interceptors...)
}

// Create returns a builder for creating a TLeaf entity.
func (c *TLeafClient) Create() *TLeafCreate {
	mutation := newTLeafMutation(c.config, OpCreate)
	return &TLeafCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TLeaf entities.
func (c *TLeafClient) CreateBulk(builders ...*TLeafCreate) *TLeafCreateBulk {
	return &TLeafCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TLeafClient) MapCreateBulk(slice any, setFunc func(*TLeafCreate, int)) *TLeafCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TLeafCreateBulk{err: fmt.Errorf("calling to TLeafClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TLeafCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TLeafCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TLeaf.
func (c *TLeafClient) Update() *TLeafUpdate {
	mutation := newTLeafMutation(c.config, OpUpdate)
	return &TLeafUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TLeafClient) UpdateOne(t *TLeaf) *TLeafUpdateOne {
	mutation := newTLeafMutation(c.config, OpUpdateOne, withTLeaf(t))
	return &TLeafUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TLeafClient) UpdateOneID(id int64) *TLeafUpdateOne {
	mutation := newTLeafMutation(c.config, OpUpdateOne, withTLeafID(id))
	return &TLeafUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TLeaf.
func (c *TLeafClient) Delete() *TLeafDelete {
	mutation := newTLeafMutation(c.config, OpDelete)
	return &TLeafDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TLeafClient) DeleteOne(t *TLeaf) *TLeafDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TLeafClient) DeleteOneID(id int64) *TLeafDeleteOne {
	builder := c.Delete().Where(tleaf.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TLeafDeleteOne{builder}
}

// Query returns a query builder for TLeaf.
func (c *TLeafClient) Query() *TLeafQuery {
	return &TLeafQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTLeaf},
		inters: c.Interceptors(),
	}
}

// Get returns a TLeaf entity by its id.
func (c *TLeafClient) Get(ctx context.Context, id int64) (*TLeaf, error) {
	return c.Query().Where(tleaf.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TLeafClient) GetX(ctx context.Context, id int64) *TLeaf {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TLeafClient) Hooks() []Hook {
	return c.hooks.TLeaf
}

// Interceptors returns the client interceptors.
func (c *TLeafClient) Interceptors() []Interceptor {
	return c.inters.TLeaf
}

func (c *TLeafClient) mutate(ctx context.Context, m *TLeafMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TLeafCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TLeafUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TLeafUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TLeafDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TLeaf mutation op: %q", m.Op())
	}
}

// TShortUrlClient is a client for the TShortUrl schema.
type TShortUrlClient struct {
	config
}

// NewTShortUrlClient returns a client for the TShortUrl from the given config.
func NewTShortUrlClient(c config) *TShortUrlClient {
	return &TShortUrlClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tshorturl.Hooks(f(g(h())))`.
func (c *TShortUrlClient) Use(hooks ...Hook) {
	c.hooks.TShortUrl = append(c.hooks.TShortUrl, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `tshorturl.Intercept(f(g(h())))`.
func (c *TShortUrlClient) Intercept(interceptors ...Interceptor) {
	c.inters.TShortUrl = append(c.inters.TShortUrl, interceptors...)
}

// Create returns a builder for creating a TShortUrl entity.
func (c *TShortUrlClient) Create() *TShortUrlCreate {
	mutation := newTShortUrlMutation(c.config, OpCreate)
	return &TShortUrlCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TShortUrl entities.
func (c *TShortUrlClient) CreateBulk(builders ...*TShortUrlCreate) *TShortUrlCreateBulk {
	return &TShortUrlCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TShortUrlClient) MapCreateBulk(slice any, setFunc func(*TShortUrlCreate, int)) *TShortUrlCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TShortUrlCreateBulk{err: fmt.Errorf("calling to TShortUrlClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TShortUrlCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TShortUrlCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TShortUrl.
func (c *TShortUrlClient) Update() *TShortUrlUpdate {
	mutation := newTShortUrlMutation(c.config, OpUpdate)
	return &TShortUrlUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TShortUrlClient) UpdateOne(tu *TShortUrl) *TShortUrlUpdateOne {
	mutation := newTShortUrlMutation(c.config, OpUpdateOne, withTShortUrl(tu))
	return &TShortUrlUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TShortUrlClient) UpdateOneID(id int64) *TShortUrlUpdateOne {
	mutation := newTShortUrlMutation(c.config, OpUpdateOne, withTShortUrlID(id))
	return &TShortUrlUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TShortUrl.
func (c *TShortUrlClient) Delete() *TShortUrlDelete {
	mutation := newTShortUrlMutation(c.config, OpDelete)
	return &TShortUrlDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TShortUrlClient) DeleteOne(tu *TShortUrl) *TShortUrlDeleteOne {
	return c.DeleteOneID(tu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TShortUrlClient) DeleteOneID(id int64) *TShortUrlDeleteOne {
	builder := c.Delete().Where(tshorturl.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TShortUrlDeleteOne{builder}
}

// Query returns a query builder for TShortUrl.
func (c *TShortUrlClient) Query() *TShortUrlQuery {
	return &TShortUrlQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTShortUrl},
		inters: c.Interceptors(),
	}
}

// Get returns a TShortUrl entity by its id.
func (c *TShortUrlClient) Get(ctx context.Context, id int64) (*TShortUrl, error) {
	return c.Query().Where(tshorturl.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TShortUrlClient) GetX(ctx context.Context, id int64) *TShortUrl {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TShortUrlClient) Hooks() []Hook {
	return c.hooks.TShortUrl
}

// Interceptors returns the client interceptors.
func (c *TShortUrlClient) Interceptors() []Interceptor {
	return c.inters.TShortUrl
}

func (c *TShortUrlClient) mutate(ctx context.Context, m *TShortUrlMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TShortUrlCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TShortUrlUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TShortUrlUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TShortUrlDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TShortUrl mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		TLeaf, TShortUrl []ent.Hook
	}
	inters struct {
		TLeaf, TShortUrl []ent.Interceptor
	}
)
