package testingFunc

import "net/url"

func NewParamUpdater(values map[string]string) ParamUpdater {
	return ParamUpdater{
		values: values,
	}
}

type ParamUpdater struct {
	values map[string]string
}

func (u ParamUpdater) Set(key, value string) ParamUpdater {
	newValues := copyMap(u.values)
	newValues[key] = value
	u.values = newValues
	return u
}

func (u ParamUpdater) Delete(key string) ParamUpdater {
	newValues := copyMap(u.values)
	delete(newValues, key)
	u.values = newValues
	return u
}

func (u ParamUpdater) Get() url.Values {
	urlValues := url.Values{}
	for key, value := range u.values {
		urlValues.Set(key, value)
	}
	return urlValues
}
