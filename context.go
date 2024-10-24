package clover

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
)

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Params  map[string]string
	mu      sync.RWMutex
	Keys    map[string]any
}

// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes  c.Keys if it was not used previously.
func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}

	c.Keys[key] = value
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exist it returns (nil, false)
func (c *Context) Get(key string) (value any, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists = c.Keys[key]
	return
}

func (c *Context) validate(target *interface{}, validatorOpts ...validator.Option) error {
	var err error = nil
	if len(validatorOpts) > 0 {
		validate := validator.New(validatorOpts...)
		err = validate.Struct(target)
		if err != nil {
			return err
		}
	}
	return err
}

func (c *Context) BodyForm(target *interface{}) error {
	err := c.Request.ParseForm()
	if err != nil {
		return err
	}
	for k, v := range c.Request.Form {
		err := json.Unmarshal([]byte(v[0]), target)
		if err != nil {
			return err
		}
		c.Params[k] = v[0]
	}
	return c.validate(target)
}
func (c *Context) BodyJSON(target *interface{}, validatorOpts ...validator.Option) error {
	err := json.NewDecoder(c.Request.Body).Decode(target)
	if err != nil {
		return err
	}
	return c.validate(target, validatorOpts...)
}

func (c *Context) ResponseJSON(code int, body interface{}) {
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(body)
}

func NewContext(w http.ResponseWriter, r *http.Request, p map[string]string) *Context {
	return &Context{
		Request: r,
		Writer:  w,
		Params:  p,
	}
}

type HandlerFunc func(*Context)

func (f HandlerFunc) ServeHTTP(ctx *Context) {
	f(ctx)
}

type H map[string]any
