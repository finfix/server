package testUtils

import (
	"testing"

	"github.com/tidwall/sjson"
)

func NewJSONUpdater(t *testing.T, json string) JSONUpdater {
	return JSONUpdater{
		t:    t,
		json: json,
	}
}

type JSONUpdater struct {
	t    *testing.T
	json string
}

func (u JSONUpdater) Set(path string, value any) JSONUpdater {
	var err error
	u.json, err = sjson.Set(u.json, path, value)
	if err != nil {
		u.t.Fatalf("\n\nОшибка при изменении json: %v\n\n", err)
	}
	return u
}

func (u JSONUpdater) Delete(path string) JSONUpdater {
	var err error
	u.json, err = sjson.Delete(u.json, path)
	if err != nil {
		u.t.Fatalf("\n\nОшибка при удалении поля json: %v\n\n", err)
	}
	return u
}

func (u JSONUpdater) Get() string {
	return u.json
}
