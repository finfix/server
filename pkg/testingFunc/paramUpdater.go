package testingFunc

import "net/url"

func NewParamUpdater(values map[string]string) paramUpdater {
	return paramUpdater{
		values: values,
	}
}

type paramUpdater struct {
	values map[string]string
}

func (u paramUpdater) Set(key, value string) paramUpdater {
	newValues := copyMap(u.values)
	newValues[key] = value
	u.values = newValues
	return u
}

func (u paramUpdater) Delete(key string) paramUpdater {
	newValues := copyMap(u.values)
	delete(newValues, key)
	u.values = newValues
	return u
}

func (u paramUpdater) Get() url.Values {
	urlValues := url.Values{}
	for key, value := range u.values {
		urlValues.Set(key, value)
	}
	return urlValues
}
