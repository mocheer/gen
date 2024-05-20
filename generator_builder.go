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
func (g *Generator) BelongsTo(fieldName string, tableName string) model.CreateFieldOpt {
	tag := field.GormTag{}
	tag.Set("foreignKey", fmt.Sprintf("%s_id", tableName))
	return FieldRelate(field.BelongsTo, fieldName, g.BuildModel(tableName), &field.RelateConfig{
		RelatePointer: true,
		GORMTag:       tag,
	})
}

// HasOne
func (g *Generator) HasOne(fieldName string, tableName string) model.CreateFieldOpt {
	tag := field.GormTag{}
	tag.Set("foreignKey", "id")
	return FieldRelate(field.HasOne, fieldName, g.BuildModel(tableName), &field.RelateConfig{
		RelatePointer: true,
		GORMTag:       tag,
	})
}

// HasMany
func (g *Generator) HasMany(fieldName string, tableName string) model.CreateFieldOpt {
	tag := field.GormTag{}
	tag.Set("foreignKey", "id")
	return FieldRelate(field.HasMany, fieldName, g.BuildModel(tableName), &field.RelateConfig{
		RelateSlicePointer: true,
		GORMTag:            tag,
	})
}

// HasTree 用于一对多的关联配置，以pid为外键进行关联，构建成 {items:[{items:...,nodes:[]}],nodes:[]}
func (g *Generator) HasTree(fieldName string, tableName string) model.CreateFieldOpt {
	tag := field.GormTag{}
	tag.Set("foreignKey", "pid")
	tag.Set("references", "id")
	return FieldRelate(field.HasMany, fieldName, g.BuildModel(tableName), &field.RelateConfig{
		RelateSlicePointer: true,
		GORMTag:            tag,
	})
}

// HasTreeByParentID
func (g *Generator) HasTreeByParentID(fieldName string, tableName string) model.CreateFieldOpt {
	tag := field.GormTag{}
	tag.Set("foreignKey", "parent_id")
	tag.Set("references", "id")
	return FieldRelate(field.HasMany, fieldName, g.BuildModel(tableName), &field.RelateConfig{
		RelateSlicePointer: true,
		GORMTag:            tag,
	})
}

// HasColumn
func (g *Generator) HasColumn(fieldName string, typeName string) model.CreateFieldOpt {
	tag := field.Tag{}
	tag.Set("json", ns.ColumnName("", fieldName))

	tag2 := field.GormTag{}
	tag2.Set("-", "")
	return func(*model.Field) *model.Field {
		return &model.Field{
			Name:    fieldName,
			Type:    typeName,
			Tag:     tag,
			GORMTag: tag2,
			// NewTag:       config.NewTag,
			// OverwriteTag: config.OverwriteTag,
		}
	}
}
