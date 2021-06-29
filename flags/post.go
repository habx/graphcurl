package flags

type Post struct {
	Requests
	Variables           map[string]string
	VariablesFromFile   map[string]string
	FilePath            string
	ExitFailIfFail      bool
	ExitFailIfCondition map[string]string
}
