package graphrequest

import (
	"context"
	"github.com/machinebox/graphql"
	"go.uber.org/zap"
	"time"
)

type PostConfig struct {
	URL       string
	Headers   map[string]string
	Query     string
	Variables map[string]interface{}
	Logger    *zap.SugaredLogger
}

func (p *PostConfig) PostRetry(count int, delay int) (interface{}, error) {
	var err error
	var data interface{}
	for i := 0; i < count; i++ {
		if count > 1 {
			p.Logger.Infow("Retry post status", "URL", p.URL, "countLoop", count, "step", i+1)
		}
		data, err = p.Post()
		if err != nil {
			p.Logger.Errorw("Retry post, there was an error", "URL", p.URL, "countLoop", count, "step", i+1)
		} else {
			return data, nil
		}
		// exit for last loop
		if (i + 1) == count {
			return data, err
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}
	return data, err
}

func (p *PostConfig) SetupUserAgent(userAgent string) {
	p.Headers["User-agent"] = userAgent
}
func (p *PostConfig) Post() (interface{}, error) {
	client := graphql.NewClient(p.URL)
	client.Log = func(s string) { p.Logger.Debug("Request: " + s) }
	req := graphql.NewRequest(p.Query)
	for k, v := range p.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range p.Variables {
		req.Var(k, v)
	}
	var respData interface{}
	if err := client.Run(context.Background(), req, &respData); err != nil {
		return respData, err
	}
	return respData, nil
}
