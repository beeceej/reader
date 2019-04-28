package reader

// EnvKey is some identifier, used to extract a value out of the environment
type EnvKey string

// EnvVal represents the value at a given key
type EnvVal interface{}

// Env is the Env that our Reader works with
type Env map[EnvKey]EnvVal

// MonadReader is a reader, with access to it's environment
type MonadReader struct {
	Reader
	Env Env
}

// Ask returns the current environment from the reader
func (m MonadReader) Ask() Env {
	return m.Env
}

// Bind is defined as MonadReader r a -> (r -> MonadReader r a) -> MonadReader r a
// Bind takes in a function which accepts an env,
// and returns a new MonadReader
func (m MonadReader) Bind(fn func(env Env) MonadReader) MonadReader {
	return fn(m.Ask())
}

// With takes a function that takes an env,
// and the MonadReader supplies the env the function
func (m MonadReader) With(fn func(env Env)) {
	fn(m.Ask())
}

// Reader is a function which given an environment, will return some value
type Reader interface {
	Run(Env) EnvVal
}

// AReader is type representing an instance of a the interface Reader
type AReader func(Env) EnvVal

// Run allows AReader to satisfy the Reader interface
// When you some reader says r.Run(someEnv) it is the same as saying r(someEnv)
func (runReader AReader) Run(env Env) EnvVal {
	return runReader(env)
}

// KVReader will return a MonadReader,
// and set the env[k] = v
func KVReader(env Env, k EnvKey, v EnvVal) MonadReader {
	env[k] = v
	return MonadReader{
		AReader(func(env Env) EnvVal {
			return env[k]
		}),
		env,
	}
}

// NewReader takes in an environment and gives you an empty Reader
func NewReader(env Env) MonadReader {
	return MonadReader{
		AReader(func(env Env) EnvVal {
			return ""
		}),
		env,
	}
}

func NewEnv() map[EnvKey]EnvVal {
	return map[EnvKey]EnvVal{}
}
