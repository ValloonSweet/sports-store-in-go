package templates

import (
	"errors"
	"html/template"
	"platform/config"
	"sync"
)

var once = sync.Once{}

func LoadTemplates(c config.Configuration) (err error) {
	path, ok := c.GetString("templates:path")
	if !ok {
		return errors.New("cannot load template config")
	}

	reload := c.GetBoolDefault("templates:reload", true)
	once.Do(func() {
		doLoad := func() (t *template.Template) {
			t = template.New("htmlTemplates")
			t.Funcs(map[string]interface{}{
				"body":    func() string { return "" },
				"layout":  func() string { return "" },
				"handler": func() interface{} { return "" },
			})
			t, err = t.ParseGlob(path)
			return
		}
		if reload {
			getTemplates = doLoad
		} else {
			templates := doLoad()
			getTemplates = func() *template.Template {
				t, _ := templates.Clone()
				return t
			}
		}
	})
	return
}
