package engine

type (
	Task struct {
		Path string
		Envs map[string]string
		Args []string
		Name string
	}
)
