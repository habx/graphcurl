// nolint: typecheck
package post

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_BuildVariablesStruct(t *testing.T) {
	variables, _ := mergeVariables(map[string]string{
		"one":           "a",
		"two":           "b",
		"cliNested.one": "a",
		"cliNested.two": "a",
	},
		map[string]string{
			"test": "../../tests/datafile/variables.json",
		})
	Convey("Standard processing", t, func() {
		build := BuildVariablesStruct(variables)
		So(build, ShouldNotBeNil)
		So(build, ShouldResemble, map[string]interface{}{
			"cliNested": map[string]interface{}{
				"one": "a",
				"two": "a",
			},
			"one": "a",
			"test": map[string]interface{}{
				"taskErrorMessage": "",
				"taskStatus":       "succeeded",
			},
			"two": "b",
		})
	})
}
func Test_mergeVariables(t *testing.T) {
	Convey("Bad file don't exists", t, func() {
		variables, err := mergeVariables(map[string]string{
			"one": "a",
			"two": "b",
		},
			map[string]string{
				"test": "../../tests/dataf",
			})
		So(variables, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("Merge is ok", t, func() {
		variables, err := mergeVariables(map[string]string{
			"one":     "a",
			"two":     "b",
			"enabled": "true",
		},
			map[string]string{
				"test": "../../tests/datafile/variables.json",
			})
		So(variables, ShouldResemble, map[string]interface{}{
			"one": "a",
			"test": map[string]interface{}{
				"taskErrorMessage": "",
				"taskStatus":       "succeeded"},
			"two":     "b",
			"enabled": true,
		})
		So(err, ShouldBeNil)
	})
}

func Test_loadVariablesFromFile(t *testing.T) {
	Convey("Bad file", t, func() {
		Convey("don't exists", func() {
			variables, err := loadVariablesFromFile(map[string]string{
				"bad": "bad_file.json",
			})
			So(variables, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
		Convey("other error (is directory)", func() {
			variables, err := loadVariablesFromFile(map[string]string{
				"bad": "../../tests/datafile",
			})
			So(variables, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})

	Convey("bad file format", t, func() {
		variables, err := loadVariablesFromFile(map[string]string{
			"bad": "../../tests/datafile/variables_bad.json",
		})
		So(variables, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
	Convey("File exists", t, func() {
		variables, err := loadVariablesFromFile(map[string]string{
			"test": "../../tests/datafile/variables.json",
		})
		So(variables, ShouldResemble, map[string]interface{}{
			"test": map[string]interface{}{
				"taskErrorMessage": "",
				"taskStatus":       "succeeded",
			},
		})
		So(err, ShouldBeNil)
	})
}

func Test_Command(t *testing.T) {
	Convey("Standard command", t, func() {
		Command.SetArgs([]string{
			"-u=https://api.spacex.land/graphql/",
			"-f=../../tests/query/users_aggregate.graphql",
		})
		err := Command.Execute()
		So(err, ShouldBeNil)
	})

	Convey("Standard command", t, func() {
		Command.SetArgs([]string{
			"-u=https://api.spacex.land/graphql/",
			"-f=../../tests/query/users_aggregate.graphql",
		})
		err := Command.Execute()
		So(err, ShouldBeNil)
	})

	Convey("Standard command with variable", t, func() {
		Command.SetArgs([]string{
			"-u=https://api.spacex.land/graphql/",
			"-f=../../tests/query/users_aggregate.graphql",
			"-V='test=ok'",
			"-V='test.nested=ok'",
		})
		err := Command.Execute()
		So(err, ShouldBeNil)
	})

	Convey("Standard command with retry", t, func() {
		Command.SetArgs([]string{
			"-u=https://api.spacex.land/graphql/",
			"-f=../../tests/query/users_aggregate.graphql",
			"--retry=1",
		})
		err := Command.Execute()
		So(err, ShouldBeNil)
	})

	Convey("Standard command with retry with fail", t, func() {
		Command.SetArgs([]string{
			"-u=https://example.com",
			"-f=../../tests/query/users_aggregate.graphql",
			"--retry=1",
			"--retry-delay=4",
		})
		err := Command.Execute()
		So(err, ShouldBeNil)
	})
}
