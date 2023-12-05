package testingFunc

import (
	"testing"

	"github.com/tidwall/sjson"
)

func NewJSONUpdater(t *testing.T, json string) jsonUpdater {
	return jsonUpdater{
		t:    t,
		json: json,
	}
}

type jsonUpdater struct {
	t    *testing.T
	json string
}

func (u jsonUpdater) Set(path string, value any) jsonUpdater {
	var err error
	u.json, err = sjson.Set(u.json, path, value)
	if err != nil {
		u.t.Fatalf("\n\nОшибка при изменении json: %v\n\n", err)
	}
	return u
}

func (u jsonUpdater) Delete(path string) jsonUpdater {
	var err error
	u.json, err = sjson.Delete(u.json, path)
	if err != nil {
		u.t.Fatalf("\n\nОшибка при удалении поля json: %v\n\n", err)
	}
	return u
}

func (u jsonUpdater) Get() string {
	return u.json
}
