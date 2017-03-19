package dummy

import (
	"io"

	"github.com/golang/protobuf/jsonpb"

	"github.com/octavore/ketchup/proto/ketchup/models"
)

type DummyTemplateStore struct {
	Themes map[string]*models.Theme
}

func New() *DummyTemplateStore {
	return &DummyTemplateStore{
		Themes: map[string]*models.Theme{},
	}
}

func (d *DummyTemplateStore) List() ([]*models.Theme, error) {
	themes := []*models.Theme{}
	for _, t := range d.Themes {
		themes = append(themes, t)
	}
	return themes, nil
}

func (d *DummyTemplateStore) Add(r io.Reader) (*models.Theme, error) {
	theme := &models.Theme{}
	err := jsonpb.Unmarshal(r, theme)
	if err != nil {
		return nil, err
	}
	d.Themes[theme.GetName()] = theme
	return theme, nil
}

func (d *DummyTemplateStore) Get(themeName string) (*models.Theme, error) {
	return d.Themes[themeName], nil
}

func (d *DummyTemplateStore) GetTemplate(t *models.Theme, template string) (*models.ThemeTemplate, error) {
	return t.GetTemplates()[template], nil
}

func (d *DummyTemplateStore) GetAsset(t *models.Theme, asset string) (*models.ThemeAsset, error) {
	return t.GetAssets()[asset], nil
}