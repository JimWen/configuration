package configuration

import (
	"encoding/json"
	"io/ioutil"
)

func ParseString(text string, includeCallback ...IncludeCallback) *Config {
	var callback IncludeCallback
	if len(includeCallback) > 0 {
		callback = includeCallback[0]
	} else {
		callback = defaultIncludeCallback
	}
	root := Parse(text, callback)
	return NewConfigFromRoot(root)
}

func LoadConfig(filename string) *Config {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return ParseString(string(data), defaultIncludeCallback)
}

func FromObject(obj interface{}) *Config {
	data, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return ParseString(string(data), defaultIncludeCallback)
}

func defaultIncludeCallback(filename string) *HoconRoot {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return Parse(string(data), defaultIncludeCallback)
}
