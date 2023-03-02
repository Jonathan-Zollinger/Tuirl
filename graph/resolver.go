package graph

import "github.com/Jonathan-Zollinger/Tuirl/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	notecards []*model.Notecard
	sections  []*model.Section
}
