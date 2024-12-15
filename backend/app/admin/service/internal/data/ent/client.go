// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"kratos-admin/app/admin/service/internal/data/ent/migrate"

	"kratos-admin/app/admin/service/internal/data/ent/menu"
	"kratos-admin/app/admin/service/internal/data/ent/organization"
	"kratos-admin/app/admin/service/internal/data/ent/position"
	"kratos-admin/app/admin/service/internal/data/ent/role"
	"kratos-admin/app/admin/service/internal/data/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Menu is the client for interacting with the Menu builders.
	Menu *MenuClient
	// Organization is the client for interacting with the Organization builders.
	Organization *OrganizationClient
	// Position is the client for interacting with the Position builders.
	Position *PositionClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Menu = NewMenuClient(c.config)
	c.Organization = NewOrganizationClient(c.config)
	c.Position = NewPositionClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:          ctx,
		config:       cfg,
		Menu:         NewMenuClient(cfg),
		Organization: NewOrganizationClient(cfg),
		Position:     NewPositionClient(cfg),
		Role:         NewRoleClient(cfg),
		User:         NewUserClient(cfg),
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
		ctx:          ctx,
		config:       cfg,
		Menu:         NewMenuClient(cfg),
		Organization: NewOrganizationClient(cfg),
		Position:     NewPositionClient(cfg),
		Role:         NewRoleClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Menu.
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
	c.Menu.Use(hooks...)
	c.Organization.Use(hooks...)
	c.Position.Use(hooks...)
	c.Role.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Menu.Intercept(interceptors...)
	c.Organization.Intercept(interceptors...)
	c.Position.Intercept(interceptors...)
	c.Role.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *MenuMutation:
		return c.Menu.mutate(ctx, m)
	case *OrganizationMutation:
		return c.Organization.mutate(ctx, m)
	case *PositionMutation:
		return c.Position.mutate(ctx, m)
	case *RoleMutation:
		return c.Role.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// MenuClient is a client for the Menu schema.
type MenuClient struct {
	config
}

// NewMenuClient returns a client for the Menu from the given config.
func NewMenuClient(c config) *MenuClient {
	return &MenuClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `menu.Hooks(f(g(h())))`.
func (c *MenuClient) Use(hooks ...Hook) {
	c.hooks.Menu = append(c.hooks.Menu, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `menu.Intercept(f(g(h())))`.
func (c *MenuClient) Intercept(interceptors ...Interceptor) {
	c.inters.Menu = append(c.inters.Menu, interceptors...)
}

// Create returns a builder for creating a Menu entity.
func (c *MenuClient) Create() *MenuCreate {
	mutation := newMenuMutation(c.config, OpCreate)
	return &MenuCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Menu entities.
func (c *MenuClient) CreateBulk(builders ...*MenuCreate) *MenuCreateBulk {
	return &MenuCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MenuClient) MapCreateBulk(slice any, setFunc func(*MenuCreate, int)) *MenuCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MenuCreateBulk{err: fmt.Errorf("calling to MenuClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MenuCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MenuCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Menu.
func (c *MenuClient) Update() *MenuUpdate {
	mutation := newMenuMutation(c.config, OpUpdate)
	return &MenuUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MenuClient) UpdateOne(m *Menu) *MenuUpdateOne {
	mutation := newMenuMutation(c.config, OpUpdateOne, withMenu(m))
	return &MenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MenuClient) UpdateOneID(id int32) *MenuUpdateOne {
	mutation := newMenuMutation(c.config, OpUpdateOne, withMenuID(id))
	return &MenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Menu.
func (c *MenuClient) Delete() *MenuDelete {
	mutation := newMenuMutation(c.config, OpDelete)
	return &MenuDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MenuClient) DeleteOne(m *Menu) *MenuDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MenuClient) DeleteOneID(id int32) *MenuDeleteOne {
	builder := c.Delete().Where(menu.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MenuDeleteOne{builder}
}

// Query returns a query builder for Menu.
func (c *MenuClient) Query() *MenuQuery {
	return &MenuQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMenu},
		inters: c.Interceptors(),
	}
}

// Get returns a Menu entity by its id.
func (c *MenuClient) Get(ctx context.Context, id int32) (*Menu, error) {
	return c.Query().Where(menu.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MenuClient) GetX(ctx context.Context, id int32) *Menu {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a Menu.
func (c *MenuClient) QueryParent(m *Menu) *MenuQuery {
	query := (&MenuClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(menu.Table, menu.FieldID, id),
			sqlgraph.To(menu.Table, menu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, menu.ParentTable, menu.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Menu.
func (c *MenuClient) QueryChildren(m *Menu) *MenuQuery {
	query := (&MenuClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(menu.Table, menu.FieldID, id),
			sqlgraph.To(menu.Table, menu.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, menu.ChildrenTable, menu.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MenuClient) Hooks() []Hook {
	return c.hooks.Menu
}

// Interceptors returns the client interceptors.
func (c *MenuClient) Interceptors() []Interceptor {
	return c.inters.Menu
}

func (c *MenuClient) mutate(ctx context.Context, m *MenuMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MenuCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MenuUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MenuDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Menu mutation op: %q", m.Op())
	}
}

// OrganizationClient is a client for the Organization schema.
type OrganizationClient struct {
	config
}

// NewOrganizationClient returns a client for the Organization from the given config.
func NewOrganizationClient(c config) *OrganizationClient {
	return &OrganizationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `organization.Hooks(f(g(h())))`.
func (c *OrganizationClient) Use(hooks ...Hook) {
	c.hooks.Organization = append(c.hooks.Organization, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `organization.Intercept(f(g(h())))`.
func (c *OrganizationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Organization = append(c.inters.Organization, interceptors...)
}

// Create returns a builder for creating a Organization entity.
func (c *OrganizationClient) Create() *OrganizationCreate {
	mutation := newOrganizationMutation(c.config, OpCreate)
	return &OrganizationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Organization entities.
func (c *OrganizationClient) CreateBulk(builders ...*OrganizationCreate) *OrganizationCreateBulk {
	return &OrganizationCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OrganizationClient) MapCreateBulk(slice any, setFunc func(*OrganizationCreate, int)) *OrganizationCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OrganizationCreateBulk{err: fmt.Errorf("calling to OrganizationClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OrganizationCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OrganizationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Organization.
func (c *OrganizationClient) Update() *OrganizationUpdate {
	mutation := newOrganizationMutation(c.config, OpUpdate)
	return &OrganizationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrganizationClient) UpdateOne(o *Organization) *OrganizationUpdateOne {
	mutation := newOrganizationMutation(c.config, OpUpdateOne, withOrganization(o))
	return &OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrganizationClient) UpdateOneID(id uint32) *OrganizationUpdateOne {
	mutation := newOrganizationMutation(c.config, OpUpdateOne, withOrganizationID(id))
	return &OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Organization.
func (c *OrganizationClient) Delete() *OrganizationDelete {
	mutation := newOrganizationMutation(c.config, OpDelete)
	return &OrganizationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrganizationClient) DeleteOne(o *Organization) *OrganizationDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrganizationClient) DeleteOneID(id uint32) *OrganizationDeleteOne {
	builder := c.Delete().Where(organization.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrganizationDeleteOne{builder}
}

// Query returns a query builder for Organization.
func (c *OrganizationClient) Query() *OrganizationQuery {
	return &OrganizationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrganization},
		inters: c.Interceptors(),
	}
}

// Get returns a Organization entity by its id.
func (c *OrganizationClient) Get(ctx context.Context, id uint32) (*Organization, error) {
	return c.Query().Where(organization.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrganizationClient) GetX(ctx context.Context, id uint32) *Organization {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a Organization.
func (c *OrganizationClient) QueryParent(o *Organization) *OrganizationQuery {
	query := (&OrganizationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, id),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, organization.ParentTable, organization.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Organization.
func (c *OrganizationClient) QueryChildren(o *Organization) *OrganizationQuery {
	query := (&OrganizationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, id),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organization.ChildrenTable, organization.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrganizationClient) Hooks() []Hook {
	return c.hooks.Organization
}

// Interceptors returns the client interceptors.
func (c *OrganizationClient) Interceptors() []Interceptor {
	return c.inters.Organization
}

func (c *OrganizationClient) mutate(ctx context.Context, m *OrganizationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrganizationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrganizationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrganizationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Organization mutation op: %q", m.Op())
	}
}

// PositionClient is a client for the Position schema.
type PositionClient struct {
	config
}

// NewPositionClient returns a client for the Position from the given config.
func NewPositionClient(c config) *PositionClient {
	return &PositionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `position.Hooks(f(g(h())))`.
func (c *PositionClient) Use(hooks ...Hook) {
	c.hooks.Position = append(c.hooks.Position, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `position.Intercept(f(g(h())))`.
func (c *PositionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Position = append(c.inters.Position, interceptors...)
}

// Create returns a builder for creating a Position entity.
func (c *PositionClient) Create() *PositionCreate {
	mutation := newPositionMutation(c.config, OpCreate)
	return &PositionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Position entities.
func (c *PositionClient) CreateBulk(builders ...*PositionCreate) *PositionCreateBulk {
	return &PositionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PositionClient) MapCreateBulk(slice any, setFunc func(*PositionCreate, int)) *PositionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PositionCreateBulk{err: fmt.Errorf("calling to PositionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PositionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PositionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Position.
func (c *PositionClient) Update() *PositionUpdate {
	mutation := newPositionMutation(c.config, OpUpdate)
	return &PositionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PositionClient) UpdateOne(po *Position) *PositionUpdateOne {
	mutation := newPositionMutation(c.config, OpUpdateOne, withPosition(po))
	return &PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PositionClient) UpdateOneID(id uint32) *PositionUpdateOne {
	mutation := newPositionMutation(c.config, OpUpdateOne, withPositionID(id))
	return &PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Position.
func (c *PositionClient) Delete() *PositionDelete {
	mutation := newPositionMutation(c.config, OpDelete)
	return &PositionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PositionClient) DeleteOne(po *Position) *PositionDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PositionClient) DeleteOneID(id uint32) *PositionDeleteOne {
	builder := c.Delete().Where(position.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PositionDeleteOne{builder}
}

// Query returns a query builder for Position.
func (c *PositionClient) Query() *PositionQuery {
	return &PositionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePosition},
		inters: c.Interceptors(),
	}
}

// Get returns a Position entity by its id.
func (c *PositionClient) Get(ctx context.Context, id uint32) (*Position, error) {
	return c.Query().Where(position.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PositionClient) GetX(ctx context.Context, id uint32) *Position {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a Position.
func (c *PositionClient) QueryParent(po *Position) *PositionQuery {
	query := (&PositionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(position.Table, position.FieldID, id),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, position.ParentTable, position.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Position.
func (c *PositionClient) QueryChildren(po *Position) *PositionQuery {
	query := (&PositionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(position.Table, position.FieldID, id),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, position.ChildrenTable, position.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PositionClient) Hooks() []Hook {
	return c.hooks.Position
}

// Interceptors returns the client interceptors.
func (c *PositionClient) Interceptors() []Interceptor {
	return c.inters.Position
}

func (c *PositionClient) mutate(ctx context.Context, m *PositionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PositionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PositionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PositionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Position mutation op: %q", m.Op())
	}
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `role.Intercept(f(g(h())))`.
func (c *RoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Role = append(c.inters.Role, interceptors...)
}

// Create returns a builder for creating a Role entity.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RoleClient) MapCreateBulk(slice any, setFunc func(*RoleCreate, int)) *RoleCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RoleCreateBulk{err: fmt.Errorf("calling to RoleClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RoleCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id uint32) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoleClient) DeleteOneID(id uint32) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRole},
		inters: c.Interceptors(),
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id uint32) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id uint32) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// Interceptors returns the client interceptors.
func (c *RoleClient) Interceptors() []Interceptor {
	return c.inters.Role
}

func (c *RoleClient) mutate(ctx context.Context, m *RoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Role mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uint32) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uint32) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uint32) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uint32) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Menu, Organization, Position, Role, User []ent.Hook
	}
	inters struct {
		Menu, Organization, Position, Role, User []ent.Interceptor
	}
)
