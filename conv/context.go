package conv

type Context map[string]any

func (c Context) Set(key string, value any) {
	c[key] = value
}

func (c Context) Get(key string) any {
	return c[key]
}
