package post

import (
	"encoding/json"
	"fmt"
	"github.com/imdario/mergo"
	"github.com/tidwall/sjson"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/habx/graphcurl/flags"
	"github.com/habx/graphcurl/graphrequest"
	"github.com/habx/graphcurl/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	// Command performs the "repos webhook" function
	Command = &cobra.Command{
		Use:   "post",
		Short: "Post your graphQL request",
		Long:  `You can use env variable for flags`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			v := viper.New()
			flags.BindFlags(cmd, v)
			return nil
		},
		PreRun: validation,
		Run:    run,
	}
	l           *zap.SugaredLogger
	requests    flags.Post
	QueryString string
)

func init() {

	// requests
	Command.Flags().StringVarP(&requests.URL, "url", "u", "", "URL API graphrequest")
	MarkFlagRequiredErr(Command.MarkFlagRequired("url"))
	Command.Flags().StringVarP(&requests.UserAgent, "user-agent", "A", "graphcurl/"+flags.Version, " Send User-Agent <name> to server")
	Command.Flags().StringToStringVarP(&requests.Headers, "headers", "H", map[string]string{}, "Pass custom header(s) to server")
	Command.Flags().IntVarP(&requests.Retry, "retry", "", 0, "Retry request if transient problems occur")
	Command.Flags().IntVarP(&requests.RetryDelay, "retry-delay", "", 0, "Wait time between retrie")
	Command.Flags().StringToStringVarP(&requests.Variables, "variables", "V", map[string]string{}, "Graphql variables (work with json data ; example info.aa -> info={'aa':'value'}; info.bb -> {'bb':'value'} )")
	Command.Flags().StringToStringVarP(&requests.VariablesFromFile, "variables-from-file", "", map[string]string{}, "Graphql variables form file (example: files.list=/root/list.json)")

	Command.Flags().StringVarP(&requests.FilePath, "file-path", "f", "", "File for your graphrequest Query")
	MarkFlagRequiredErr(Command.MarkFlagRequired("file-path"))
	Command.Flags().BoolVarP(&requests.ExitFailIfFail, "fail", "", false, "Exit 1 if transient problems")
	Command.Flags().StringToStringVarP(&requests.ExitFailIfCondition, "exit-fail-if-condition", "", map[string]string{}, "Exit 1 if first map params is true (value = value)")
}
func validation(cmd *cobra.Command, cmdLineArgs []string) {
	if flags.Slient {
		flags.LogLevel = "-99"
	}
	l = logger.GetLogger(flags.LogLevel).Sugar().With("commands", "post")
	if _, err := os.Stat(requests.FilePath); os.IsNotExist(err) {
		l.Fatalw("Check your query file path", "err", err)
	}
	content, err := ioutil.ReadFile(requests.FilePath)

	if err != nil {
		l.Fatalw("cannot read file", "err", err)
	}
	QueryString = string(content)
}

func run(cmd *cobra.Command, cmdLineArgs []string) {
	var err error
	var data interface{}
	variablesMerged, err := mergeVariables(requests.Variables, requests.VariablesFromFile)
	if err != nil {
		l.Fatalw("cannot merge variables, file and cli modes", "err", err)
	}
	l.Debugw("Merged variables", "variables", BuildVariablesStruct(variablesMerged))
	newPostRequest := graphrequest.PostConfig{
		URL:       requests.URL,
		Headers:   requests.Headers,
		Query:     QueryString,
		Variables: BuildVariablesStruct(variablesMerged),
		Logger:    l,
	}
	newPostRequest.SetupUserAgent(requests.UserAgent)
	l.Infow("Exec http request", "URL", requests.URL)
	if requests.Retry > 0 {
		data, err = newPostRequest.PostRetry(requests.Retry, requests.RetryDelay)
	} else {
		data, err = newPostRequest.Post()
	}
	if err != nil {
		if requests.ExitFailIfFail {
			l.Fatalw("request error (exit if fail)", "err", err)
		} else {
			l.Errorw("request error", "err", err)
			return
		}
	}
	if len(requests.ExitFailIfCondition) == 1 {
		for k, v := range requests.ExitFailIfCondition {
			if k == v {
				l.Fatalw("Exit (exit with condition)", "key", k, "value", v)
			}
		}
	}
	JSONString, _ := json.Marshal(data)
	// Cannot use logger for slient mode
	fmt.Println(string(JSONString))
}
func MarkFlagRequiredErr(err error) {
	if err != nil {
		panic(err)
	}
}

func BuildVariablesStruct(variables map[string]interface{}) map[string]interface{} {
	variablesStruct := make(map[string]interface{})
	variablesJSON := []byte("{}")
	for k, v := range variables {
		variablesJSON, _ = sjson.SetBytesOptions(variablesJSON, k, v, &sjson.Options{
			Optimistic:     false,
			ReplaceInPlace: false,
		})
		variablesStructTmp := make(map[string]interface{})
		_ = json.Unmarshal(variablesJSON, &variablesStructTmp)
		_ = mergo.Merge(&variablesStruct, variablesStructTmp)
	}
	return variablesStruct
}

func mergeVariables(variablesFromCli map[string]string, variablesFromFile map[string]string) (map[string]interface{}, error) {
	variables := make(map[string]interface{})
	variablesFromFileLoaded, err := loadVariablesFromFile(variablesFromFile)
	if err != nil {
		return nil, err
	}
	// from files
	for k, v := range variablesFromFileLoaded {
		variables[k] = v
	}
	// from cli
	for k, v := range variablesFromCli {
		if v == "true" || v == "false" {
			vToBool, _ := strconv.ParseBool(v)
			if err == nil {
				variables[k] = vToBool
			}
		} else {
			variables[k] = v
		}

	}
	return variables, nil
}
func loadVariablesFromFile(variablesFromFile map[string]string) (map[string]interface{}, error) {
	variablesFromFileLoaded := make(map[string]interface{})
	for k, file := range variablesFromFile {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return nil, err
		}
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		var variablesFromFile interface{}
		err = json.Unmarshal(content, &variablesFromFile)
		if err != nil {
			return nil, err
		}
		variablesFromFileLoaded[k] = variablesFromFile
	}
	return variablesFromFileLoaded, nil

}
