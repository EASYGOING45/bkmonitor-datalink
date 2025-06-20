// Code generated by go-queryset. DO NOT EDIT.
package resulttable

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set ESFieldQueryAliasOptionQuerySet

// ESFieldQueryAliasOptionQuerySet is an queryset type for ESFieldQueryAliasOption
type ESFieldQueryAliasOptionQuerySet struct {
	db *gorm.DB
}

// NewESFieldQueryAliasOptionQuerySet constructs new ESFieldQueryAliasOptionQuerySet
func NewESFieldQueryAliasOptionQuerySet(db *gorm.DB) ESFieldQueryAliasOptionQuerySet {
	return ESFieldQueryAliasOptionQuerySet{
		db: db.Model(&ESFieldQueryAliasOption{}),
	}
}

func (qs ESFieldQueryAliasOptionQuerySet) w(db *gorm.DB) ESFieldQueryAliasOptionQuerySet {
	return NewESFieldQueryAliasOptionQuerySet(db)
}

func (qs ESFieldQueryAliasOptionQuerySet) Select(fields ...ESFieldQueryAliasOptionDBSchemaField) ESFieldQueryAliasOptionQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *ESFieldQueryAliasOption) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *ESFieldQueryAliasOption) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) All(ret *[]ESFieldQueryAliasOption) error {
	return qs.db.Find(ret).Error
}

// BkTenantIDEq is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) BkTenantIDEq(bkTenantID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("bk_tenant_id = ?", bkTenantID))
}

// BkTenantIDGt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) BkTenantIDGt(bkTenantID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("bk_tenant_id > ?", bkTenantID))
}

// BkTenantIDGte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) BkTenantIDGte(bkTenantID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("bk_tenant_id >= ?", bkTenantID))
}

// BkTenantIDIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) BkTenantIDIn(bkTenantID ...string) ESFieldQueryAliasOptionQuerySet {
	if len(bkTenantID) == 0 {
		qs.db.AddError(errors.New("must at least pass one bkTenantID in BkTenantIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bk_tenant_id IN (?)", bkTenantID))
}

// BkTenantIDLike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) BkTenantIDLike(bkTenantID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("bk_tenant_id LIKE ?", bkTenantID))
}

// BkTenantIDLt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) BkTenantIDLt(bkTenantID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("bk_tenant_id < ?", bkTenantID))
}

// BkTenantIDLte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) BkTenantIDLte(bkTenantID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("bk_tenant_id <= ?", bkTenantID))
}

// BkTenantIDNe is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) BkTenantIDNe(bkTenantID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("bk_tenant_id != ?", bkTenantID))
}

// BkTenantIDNotIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) BkTenantIDNotIn(bkTenantID ...string) ESFieldQueryAliasOptionQuerySet {
	if len(bkTenantID) == 0 {
		qs.db.AddError(errors.New("must at least pass one bkTenantID in BkTenantIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bk_tenant_id NOT IN (?)", bkTenantID))
}

// BkTenantIDNotlike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) BkTenantIDNotlike(bkTenantID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("bk_tenant_id NOT LIKE ?", bkTenantID))
}

// Count is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// CreateTimeEq is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreateTimeEq(createTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("create_time = ?", createTime))
}

// CreateTimeGt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreateTimeGt(createTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("create_time > ?", createTime))
}

// CreateTimeGte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreateTimeGte(createTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("create_time >= ?", createTime))
}

// CreateTimeLt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreateTimeLt(createTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("create_time < ?", createTime))
}

// CreateTimeLte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreateTimeLte(createTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("create_time <= ?", createTime))
}

// CreateTimeNe is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreateTimeNe(createTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("create_time != ?", createTime))
}

// CreatorEq is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreatorEq(creator string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("creator = ?", creator))
}

// CreatorGt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreatorGt(creator string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("creator > ?", creator))
}

// CreatorGte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreatorGte(creator string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("creator >= ?", creator))
}

// CreatorIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreatorIn(creator ...string) ESFieldQueryAliasOptionQuerySet {
	if len(creator) == 0 {
		qs.db.AddError(errors.New("must at least pass one creator in CreatorIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("creator IN (?)", creator))
}

// CreatorLike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreatorLike(creator string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("creator LIKE ?", creator))
}

// CreatorLt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreatorLt(creator string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("creator < ?", creator))
}

// CreatorLte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreatorLte(creator string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("creator <= ?", creator))
}

// CreatorNe is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreatorNe(creator string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("creator != ?", creator))
}

// CreatorNotIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreatorNotIn(creator ...string) ESFieldQueryAliasOptionQuerySet {
	if len(creator) == 0 {
		qs.db.AddError(errors.New("must at least pass one creator in CreatorNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("creator NOT IN (?)", creator))
}

// CreatorNotlike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) CreatorNotlike(creator string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("creator NOT LIKE ?", creator))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) Delete() error {
	return qs.db.Delete(ESFieldQueryAliasOption{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(ESFieldQueryAliasOption{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(ESFieldQueryAliasOption{})
	return db.RowsAffected, db.Error
}

// FieldPathEq is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) FieldPathEq(fieldPath string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("field_path = ?", fieldPath))
}

// FieldPathGt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) FieldPathGt(fieldPath string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("field_path > ?", fieldPath))
}

// FieldPathGte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) FieldPathGte(fieldPath string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("field_path >= ?", fieldPath))
}

// FieldPathIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) FieldPathIn(fieldPath ...string) ESFieldQueryAliasOptionQuerySet {
	if len(fieldPath) == 0 {
		qs.db.AddError(errors.New("must at least pass one fieldPath in FieldPathIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("field_path IN (?)", fieldPath))
}

// FieldPathLike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) FieldPathLike(fieldPath string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("field_path LIKE ?", fieldPath))
}

// FieldPathLt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) FieldPathLt(fieldPath string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("field_path < ?", fieldPath))
}

// FieldPathLte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) FieldPathLte(fieldPath string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("field_path <= ?", fieldPath))
}

// FieldPathNe is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) FieldPathNe(fieldPath string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("field_path != ?", fieldPath))
}

// FieldPathNotIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) FieldPathNotIn(fieldPath ...string) ESFieldQueryAliasOptionQuerySet {
	if len(fieldPath) == 0 {
		qs.db.AddError(errors.New("must at least pass one fieldPath in FieldPathNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("field_path NOT IN (?)", fieldPath))
}

// FieldPathNotlike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) FieldPathNotlike(fieldPath string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("field_path NOT LIKE ?", fieldPath))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) GetUpdater() ESFieldQueryAliasOptionUpdater {
	return NewESFieldQueryAliasOptionUpdater(qs.db)
}

// IsDeletedEq is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) IsDeletedEq(isDeleted bool) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("is_deleted = ?", isDeleted))
}

// IsDeletedIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) IsDeletedIn(isDeleted ...bool) ESFieldQueryAliasOptionQuerySet {
	if len(isDeleted) == 0 {
		qs.db.AddError(errors.New("must at least pass one isDeleted in IsDeletedIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("is_deleted IN (?)", isDeleted))
}

// IsDeletedNe is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) IsDeletedNe(isDeleted bool) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("is_deleted != ?", isDeleted))
}

// IsDeletedNotIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) IsDeletedNotIn(isDeleted ...bool) ESFieldQueryAliasOptionQuerySet {
	if len(isDeleted) == 0 {
		qs.db.AddError(errors.New("must at least pass one isDeleted in IsDeletedNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("is_deleted NOT IN (?)", isDeleted))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) Limit(limit int) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) Offset(offset int) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs ESFieldQueryAliasOptionQuerySet) One(ret *ESFieldQueryAliasOption) error {
	return qs.db.First(ret).Error
}

// OrderAscByBkTenantID is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderAscByBkTenantID() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("bk_tenant_id ASC"))
}

// OrderAscByCreateTime is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderAscByCreateTime() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("create_time ASC"))
}

// OrderAscByCreator is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderAscByCreator() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("creator ASC"))
}

// OrderAscByFieldPath is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderAscByFieldPath() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("field_path ASC"))
}

// OrderAscByIsDeleted is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderAscByIsDeleted() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("is_deleted ASC"))
}

// OrderAscByPathType is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderAscByPathType() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("path_type ASC"))
}

// OrderAscByQueryAlias is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderAscByQueryAlias() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("query_alias ASC"))
}

// OrderAscByTableID is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderAscByTableID() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("table_id ASC"))
}

// OrderAscByUpdateTime is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderAscByUpdateTime() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("update_time ASC"))
}

// OrderAscByUpdater is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderAscByUpdater() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("updater ASC"))
}

// OrderDescByBkTenantID is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderDescByBkTenantID() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("bk_tenant_id DESC"))
}

// OrderDescByCreateTime is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderDescByCreateTime() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("create_time DESC"))
}

// OrderDescByCreator is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderDescByCreator() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("creator DESC"))
}

// OrderDescByFieldPath is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderDescByFieldPath() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("field_path DESC"))
}

// OrderDescByIsDeleted is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderDescByIsDeleted() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("is_deleted DESC"))
}

// OrderDescByPathType is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderDescByPathType() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("path_type DESC"))
}

// OrderDescByQueryAlias is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderDescByQueryAlias() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("query_alias DESC"))
}

// OrderDescByTableID is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderDescByTableID() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("table_id DESC"))
}

// OrderDescByUpdateTime is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderDescByUpdateTime() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("update_time DESC"))
}

// OrderDescByUpdater is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) OrderDescByUpdater() ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Order("updater DESC"))
}

// PathTypeEq is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) PathTypeEq(pathType string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("path_type = ?", pathType))
}

// PathTypeGt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) PathTypeGt(pathType string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("path_type > ?", pathType))
}

// PathTypeGte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) PathTypeGte(pathType string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("path_type >= ?", pathType))
}

// PathTypeIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) PathTypeIn(pathType ...string) ESFieldQueryAliasOptionQuerySet {
	if len(pathType) == 0 {
		qs.db.AddError(errors.New("must at least pass one pathType in PathTypeIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("path_type IN (?)", pathType))
}

// PathTypeLike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) PathTypeLike(pathType string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("path_type LIKE ?", pathType))
}

// PathTypeLt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) PathTypeLt(pathType string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("path_type < ?", pathType))
}

// PathTypeLte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) PathTypeLte(pathType string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("path_type <= ?", pathType))
}

// PathTypeNe is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) PathTypeNe(pathType string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("path_type != ?", pathType))
}

// PathTypeNotIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) PathTypeNotIn(pathType ...string) ESFieldQueryAliasOptionQuerySet {
	if len(pathType) == 0 {
		qs.db.AddError(errors.New("must at least pass one pathType in PathTypeNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("path_type NOT IN (?)", pathType))
}

// PathTypeNotlike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) PathTypeNotlike(pathType string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("path_type NOT LIKE ?", pathType))
}

// QueryAliasEq is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) QueryAliasEq(queryAlias string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("query_alias = ?", queryAlias))
}

// QueryAliasGt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) QueryAliasGt(queryAlias string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("query_alias > ?", queryAlias))
}

// QueryAliasGte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) QueryAliasGte(queryAlias string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("query_alias >= ?", queryAlias))
}

// QueryAliasIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) QueryAliasIn(queryAlias ...string) ESFieldQueryAliasOptionQuerySet {
	if len(queryAlias) == 0 {
		qs.db.AddError(errors.New("must at least pass one queryAlias in QueryAliasIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("query_alias IN (?)", queryAlias))
}

// QueryAliasLike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) QueryAliasLike(queryAlias string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("query_alias LIKE ?", queryAlias))
}

// QueryAliasLt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) QueryAliasLt(queryAlias string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("query_alias < ?", queryAlias))
}

// QueryAliasLte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) QueryAliasLte(queryAlias string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("query_alias <= ?", queryAlias))
}

// QueryAliasNe is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) QueryAliasNe(queryAlias string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("query_alias != ?", queryAlias))
}

// QueryAliasNotIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) QueryAliasNotIn(queryAlias ...string) ESFieldQueryAliasOptionQuerySet {
	if len(queryAlias) == 0 {
		qs.db.AddError(errors.New("must at least pass one queryAlias in QueryAliasNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("query_alias NOT IN (?)", queryAlias))
}

// QueryAliasNotlike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) QueryAliasNotlike(queryAlias string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("query_alias NOT LIKE ?", queryAlias))
}

// TableIDEq is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) TableIDEq(tableID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("table_id = ?", tableID))
}

// TableIDGt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) TableIDGt(tableID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("table_id > ?", tableID))
}

// TableIDGte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) TableIDGte(tableID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("table_id >= ?", tableID))
}

// TableIDIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) TableIDIn(tableID ...string) ESFieldQueryAliasOptionQuerySet {
	if len(tableID) == 0 {
		qs.db.AddError(errors.New("must at least pass one tableID in TableIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("table_id IN (?)", tableID))
}

// TableIDLike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) TableIDLike(tableID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("table_id LIKE ?", tableID))
}

// TableIDLt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) TableIDLt(tableID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("table_id < ?", tableID))
}

// TableIDLte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) TableIDLte(tableID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("table_id <= ?", tableID))
}

// TableIDNe is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) TableIDNe(tableID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("table_id != ?", tableID))
}

// TableIDNotIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) TableIDNotIn(tableID ...string) ESFieldQueryAliasOptionQuerySet {
	if len(tableID) == 0 {
		qs.db.AddError(errors.New("must at least pass one tableID in TableIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("table_id NOT IN (?)", tableID))
}

// TableIDNotlike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) TableIDNotlike(tableID string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("table_id NOT LIKE ?", tableID))
}

// UpdateTimeEq is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdateTimeEq(updateTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("update_time = ?", updateTime))
}

// UpdateTimeGt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdateTimeGt(updateTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("update_time > ?", updateTime))
}

// UpdateTimeGte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdateTimeGte(updateTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("update_time >= ?", updateTime))
}

// UpdateTimeLt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdateTimeLt(updateTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("update_time < ?", updateTime))
}

// UpdateTimeLte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdateTimeLte(updateTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("update_time <= ?", updateTime))
}

// UpdateTimeNe is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdateTimeNe(updateTime time.Time) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("update_time != ?", updateTime))
}

// UpdaterEq is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdaterEq(updater string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("updater = ?", updater))
}

// UpdaterGt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdaterGt(updater string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("updater > ?", updater))
}

// UpdaterGte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdaterGte(updater string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("updater >= ?", updater))
}

// UpdaterIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdaterIn(updater ...string) ESFieldQueryAliasOptionQuerySet {
	if len(updater) == 0 {
		qs.db.AddError(errors.New("must at least pass one updater in UpdaterIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updater IN (?)", updater))
}

// UpdaterLike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdaterLike(updater string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("updater LIKE ?", updater))
}

// UpdaterLt is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdaterLt(updater string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("updater < ?", updater))
}

// UpdaterLte is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdaterLte(updater string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("updater <= ?", updater))
}

// UpdaterNe is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdaterNe(updater string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("updater != ?", updater))
}

// UpdaterNotIn is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdaterNotIn(updater ...string) ESFieldQueryAliasOptionQuerySet {
	if len(updater) == 0 {
		qs.db.AddError(errors.New("must at least pass one updater in UpdaterNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updater NOT IN (?)", updater))
}

// UpdaterNotlike is an autogenerated method
// nolint: dupl
func (qs ESFieldQueryAliasOptionQuerySet) UpdaterNotlike(updater string) ESFieldQueryAliasOptionQuerySet {
	return qs.w(qs.db.Where("updater NOT LIKE ?", updater))
}

// SetBkTenantID is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) SetBkTenantID(bkTenantID string) ESFieldQueryAliasOptionUpdater {
	u.fields[string(ESFieldQueryAliasOptionDBSchema.BkTenantID)] = bkTenantID
	return u
}

// SetCreateTime is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) SetCreateTime(createTime time.Time) ESFieldQueryAliasOptionUpdater {
	u.fields[string(ESFieldQueryAliasOptionDBSchema.CreateTime)] = createTime
	return u
}

// SetCreator is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) SetCreator(creator string) ESFieldQueryAliasOptionUpdater {
	u.fields[string(ESFieldQueryAliasOptionDBSchema.Creator)] = creator
	return u
}

// SetFieldPath is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) SetFieldPath(fieldPath string) ESFieldQueryAliasOptionUpdater {
	u.fields[string(ESFieldQueryAliasOptionDBSchema.FieldPath)] = fieldPath
	return u
}

// SetIsDeleted is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) SetIsDeleted(isDeleted bool) ESFieldQueryAliasOptionUpdater {
	u.fields[string(ESFieldQueryAliasOptionDBSchema.IsDeleted)] = isDeleted
	return u
}

// SetPathType is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) SetPathType(pathType string) ESFieldQueryAliasOptionUpdater {
	u.fields[string(ESFieldQueryAliasOptionDBSchema.PathType)] = pathType
	return u
}

// SetQueryAlias is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) SetQueryAlias(queryAlias string) ESFieldQueryAliasOptionUpdater {
	u.fields[string(ESFieldQueryAliasOptionDBSchema.QueryAlias)] = queryAlias
	return u
}

// SetTableID is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) SetTableID(tableID string) ESFieldQueryAliasOptionUpdater {
	u.fields[string(ESFieldQueryAliasOptionDBSchema.TableID)] = tableID
	return u
}

// SetUpdateTime is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) SetUpdateTime(updateTime time.Time) ESFieldQueryAliasOptionUpdater {
	u.fields[string(ESFieldQueryAliasOptionDBSchema.UpdateTime)] = updateTime
	return u
}

// SetUpdater is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) SetUpdater(updater string) ESFieldQueryAliasOptionUpdater {
	u.fields[string(ESFieldQueryAliasOptionDBSchema.Updater)] = updater
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u ESFieldQueryAliasOptionUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set ESFieldQueryAliasOptionQuerySet

// ===== BEGIN of ESFieldQueryAliasOption modifiers

// ESFieldQueryAliasOptionDBSchemaField describes database schema field. It requires for method 'Update'
type ESFieldQueryAliasOptionDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f ESFieldQueryAliasOptionDBSchemaField) String() string {
	return string(f)
}

// ESFieldQueryAliasOptionDBSchema stores db field names of ESFieldQueryAliasOption
var ESFieldQueryAliasOptionDBSchema = struct {
	TableID    ESFieldQueryAliasOptionDBSchemaField
	BkTenantID ESFieldQueryAliasOptionDBSchemaField
	FieldPath  ESFieldQueryAliasOptionDBSchemaField
	PathType   ESFieldQueryAliasOptionDBSchemaField
	QueryAlias ESFieldQueryAliasOptionDBSchemaField
	IsDeleted  ESFieldQueryAliasOptionDBSchemaField
	Creator    ESFieldQueryAliasOptionDBSchemaField
	CreateTime ESFieldQueryAliasOptionDBSchemaField
	Updater    ESFieldQueryAliasOptionDBSchemaField
	UpdateTime ESFieldQueryAliasOptionDBSchemaField
}{

	TableID:    ESFieldQueryAliasOptionDBSchemaField("table_id"),
	BkTenantID: ESFieldQueryAliasOptionDBSchemaField("bk_tenant_id"),
	FieldPath:  ESFieldQueryAliasOptionDBSchemaField("field_path"),
	PathType:   ESFieldQueryAliasOptionDBSchemaField("path_type"),
	QueryAlias: ESFieldQueryAliasOptionDBSchemaField("query_alias"),
	IsDeleted:  ESFieldQueryAliasOptionDBSchemaField("is_deleted"),
	Creator:    ESFieldQueryAliasOptionDBSchemaField("creator"),
	CreateTime: ESFieldQueryAliasOptionDBSchemaField("create_time"),
	Updater:    ESFieldQueryAliasOptionDBSchemaField("updater"),
	UpdateTime: ESFieldQueryAliasOptionDBSchemaField("update_time"),
}

// Update updates ESFieldQueryAliasOption fields by primary key
// nolint: dupl
func (o *ESFieldQueryAliasOption) Update(db *gorm.DB, fields ...ESFieldQueryAliasOptionDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"table_id":     o.TableID,
		"bk_tenant_id": o.BkTenantID,
		"field_path":   o.FieldPath,
		"path_type":    o.PathType,
		"query_alias":  o.QueryAlias,
		"is_deleted":   o.IsDeleted,
		"creator":      o.Creator,
		"create_time":  o.CreateTime,
		"updater":      o.Updater,
		"update_time":  o.UpdateTime,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update ESFieldQueryAliasOption %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// ESFieldQueryAliasOptionUpdater is an ESFieldQueryAliasOption updates manager
type ESFieldQueryAliasOptionUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewESFieldQueryAliasOptionUpdater creates new ESFieldQueryAliasOption updater
// nolint: dupl
func NewESFieldQueryAliasOptionUpdater(db *gorm.DB) ESFieldQueryAliasOptionUpdater {
	return ESFieldQueryAliasOptionUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&ESFieldQueryAliasOption{}),
	}
}

// ===== END of ESFieldQueryAliasOption modifiers

// ===== END of all query sets
