package gen

import (
	"fmt"
	"strings"

	"gorm.io/gen/field"
	"gorm.io/gen/internal/generate"
	"gorm.io/gen/internal/model"
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
		// 这里因为WithTableNameStrategy增加了前缀
		tableName := strings.Split(data.TableName, ".")[1]
		if tableName == name {
			return data
		}
	}
	return nil
}

// GetModelByTableName
func (g *Generator) GetModels() map[string]*generate.QueryStructMeta {
	return g.models
}

// BelongsTo
func (g *Generator) BelongsTo(key string, tableName string) model.CreateFieldOpt {
	return FieldRelate(field.BelongsTo, key, g.BuildModel(tableName), &field.RelateConfig{
		RelatePointer: true,
		GORMTag:       fmt.Sprintf("foreignKey:%s_id", tableName),
	})
}

// HasOne
func (g *Generator) HasOne(key string, tableName string) model.CreateFieldOpt {
	return FieldRelate(field.HasOne, key, g.BuildModel(tableName), &field.RelateConfig{
		RelatePointer: true,
		GORMTag:       "foreignKey:id",
	})
}

// HasMany
func (g *Generator) HasMany(key string, tableName string) model.CreateFieldOpt {
	return FieldRelate(field.HasMany, key, g.BuildModel(tableName), &field.RelateConfig{
		RelateSlicePointer: true,
		GORMTag:            "foreignKey:id",
	})
}

// HasTree 用于一对多的关联配置，以pid为外键进行关联，构建成 {items:[{items:...,nodes:[]}],nodes:[]}
func (g *Generator) HasTree(key string, tableName string) model.CreateFieldOpt {
	return FieldRelate(field.HasMany, key, g.BuildModel(tableName), &field.RelateConfig{
		RelateSlicePointer: true,
		GORMTag:            "foreignKey:pid;references:id",
	})
}

// HasTreeByParentID
func (g *Generator) HasTreeByParentID(key string, tableName string) model.CreateFieldOpt {
	return FieldRelate(field.HasMany, key, g.BuildModel(tableName), &field.RelateConfig{
		RelateSlicePointer: true,
		GORMTag:            "foreignKey:parent_id;references:id",
	})
}
