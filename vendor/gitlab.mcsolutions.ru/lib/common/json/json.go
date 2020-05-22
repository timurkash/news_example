package json

import (
	"github.com/ghodss/yaml"
	"github.com/json-iterator/go"
	"io"
	"io/ioutil"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary

	Decode = func(r io.Reader, obj interface{}) error {
		return json.NewDecoder(r).Decode(obj)
	}

	Encode = func(obj interface{}, w io.Writer) error {
		return json.NewEncoder(w).Encode(obj)
	}

	EncodeByte = func(obj interface{}) (*[]byte, error) {
		bytes, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		return &bytes, nil
	}

	DecodeYamlEncodeByte = func(r io.Reader) (*[]byte, error) {
		body, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}
		json, err := yaml.YAMLToJSON(body)
		if err != nil {
			return nil, err
		}
		return &json, nil
	}
)
