package utilities

import "github.com/mitchellh/mapstructure"

func AutoMapper(source interface{}, target interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           target,
		TagName:          "json", // Use the "json" struct tag for field mapping
		WeaklyTypedInput: true,   // Allow weakly typed input values
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(source)
	if err != nil {
		return err
	}

	return nil
}
