package reader

type EnvVal interface{}
type EnvKey string

// Env is the Env that our Reader works with
type Env map[EnvKey]EnvVal

// MonadReader is reader, with access to it's environment
type MonadReader struct {
	Reader
	Env Env
}

// Ask returns the environment from the reader
func (m MonadReader) Ask() Env {
	return m.Env
}

func (m MonadReader) Bind(fn func(env Env) MonadReader) MonadReader {
	return fn(m.Ask())
}

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

// KVReader will return a MonadReader, with k, v set.
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
