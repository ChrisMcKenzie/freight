package manifest

import (
	"errors"
	"reflect"

	"github.com/ChrisMcKenzie/freight"
	"github.com/ChrisMcKenzie/freight/tasks"
	"github.com/mitchellh/mapstructure"
)

type Project struct {
	Name   string `mapstructure:"name"`
	Remote string `mapstructure:"remote"`
	Path   string `mapstructure:"path"`

	AfterTasks []map[string]interface{} `mapstructure:"after" hcl:"after"`
}

// ProjectParseHook handles the parsing of config files for freight
func ProjectParseHook() mapstructure.DecodeHookFunc {
	return func(s reflect.Type, d reflect.Type, t interface{}) (interface{}, error) {

		switch d {
		case reflect.TypeOf([]*freight.Project{}):
			return parseProjects(t)
		case reflect.TypeOf(freight.BaseConfig{}):
			return parseBaseConfig(t)
		}

		return t, nil
	}
}

func parseBaseConfig(t interface{}) (interface{}, error) {
	var result freight.BaseConfig

	bc := t.([]map[string]interface{})[0]

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &result,
	})
	if err != nil {
		return t, err
	}
	if err := decoder.Decode(bc); err != nil {
		return nil, err
	}

	return result, nil
}

func parseProjects(t interface{}) (interface{}, error) {
	var result []*freight.Project

	for _, val := range t.([]map[string]interface{}) {
		for k, v := range val {
			p := v.([]map[string]interface{})[0]
			var pproj Project
			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				Result: &pproj,
			})
			if err != nil {
				return t, err
			}
			if err := decoder.Decode(p); err != nil {
				return nil, err
			}

			proj := &freight.Project{
				Name:   k,
				Remote: pproj.Remote,
				Path:   pproj.Path,

				AfterTasks: make([]tasks.Task, len(pproj.AfterTasks)),
			}

			for pos, cfg := range pproj.AfterTasks {
				for n, val := range cfg {
					p := val.([]map[string]interface{})[0]
					task := tasks.GetTask(n)
					if task == nil {
						return t, errors.New("task not found")
					}

					decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
						Result: &task,
					})
					if err != nil {
						return t, err
					}
					if err := decoder.Decode(p); err != nil {
						return nil, err
					}

					proj.AfterTasks[pos] = task
				}
			}

			result = append(result, proj)
		}
	}
	return result, nil
}
