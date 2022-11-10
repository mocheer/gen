package gen

import (
	"fmt"

	"gorm.io/gen/internal/generate"
)

// GetTables
func (g *Generator) GetTables(opts ...ModelOpt) (tableModels []string) {
	tableList, err := g.db.Migrator().GetTables()
	if err != nil {
		panic(fmt.Errorf("get all tables fail: %w", err))
	}

	g.info(fmt.Sprintf("find %d table from db: %s", len(tableList), tableList))

	return tableList
}

// BuildModel
func (g *Generator) BuildModel(tableName string, opts ...ModelOpt) *generate.QueryStructMeta {
	meta := g.GetModel(tableName)
	if meta == nil {
		meta = g.GenerateModel(tableName, opts...)
	}
	return meta
}

// GetModel
func (g *Generator) GetModel(name string) *generate.QueryStructMeta {
	for _, data := range g.models {
		if data == nil {
			continue
		}
		if data.TableName == name {
			return data
		}
	}
	return nil
}

// GetModelByTableName
func (g *Generator) GetModels() map[string]*generate.QueryStructMeta {
	return g.models
}
