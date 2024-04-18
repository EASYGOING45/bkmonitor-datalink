// Code generated by go-queryset. DO NOT EDIT.
package space

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set SpaceResourceQuerySet

// SpaceResourceQuerySet is an queryset type for SpaceResource
type SpaceResourceQuerySet struct {
	db *gorm.DB
}

// NewSpaceResourceQuerySet constructs new SpaceResourceQuerySet
func NewSpaceResourceQuerySet(db *gorm.DB) SpaceResourceQuerySet {
	return SpaceResourceQuerySet{
		db: db.Model(&SpaceResource{}),
	}
}

func (qs SpaceResourceQuerySet) w(db *gorm.DB) SpaceResourceQuerySet {
	return NewSpaceResourceQuerySet(db)
}

func (qs SpaceResourceQuerySet) Select(fields ...SpaceResourceDBSchemaField) SpaceResourceQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *SpaceResource) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *SpaceResource) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) All(ret *[]SpaceResource) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// CreateTimeEq is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreateTimeEq(createTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("create_time = ?", createTime))
}

// CreateTimeGt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreateTimeGt(createTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("create_time > ?", createTime))
}

// CreateTimeGte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreateTimeGte(createTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("create_time >= ?", createTime))
}

// CreateTimeLt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreateTimeLt(createTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("create_time < ?", createTime))
}

// CreateTimeLte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreateTimeLte(createTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("create_time <= ?", createTime))
}

// CreateTimeNe is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreateTimeNe(createTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("create_time != ?", createTime))
}

// CreatorEq is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreatorEq(creator string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("creator = ?", creator))
}

// CreatorGt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreatorGt(creator string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("creator > ?", creator))
}

// CreatorGte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreatorGte(creator string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("creator >= ?", creator))
}

// CreatorIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreatorIn(creator ...string) SpaceResourceQuerySet {
	if len(creator) == 0 {
		qs.db.AddError(errors.New("must at least pass one creator in CreatorIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("creator IN (?)", creator))
}

// CreatorLike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreatorLike(creator string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("creator LIKE ?", creator))
}

// CreatorLt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreatorLt(creator string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("creator < ?", creator))
}

// CreatorLte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreatorLte(creator string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("creator <= ?", creator))
}

// CreatorNe is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreatorNe(creator string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("creator != ?", creator))
}

// CreatorNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreatorNotIn(creator ...string) SpaceResourceQuerySet {
	if len(creator) == 0 {
		qs.db.AddError(errors.New("must at least pass one creator in CreatorNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("creator NOT IN (?)", creator))
}

// CreatorNotlike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) CreatorNotlike(creator string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("creator NOT LIKE ?", creator))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) Delete() error {
	return qs.db.Delete(SpaceResource{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(SpaceResource{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(SpaceResource{})
	return db.RowsAffected, db.Error
}

// DimensionValuesEq is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DimensionValuesEq(dimensionValues string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("dimension_values = ?", dimensionValues))
}

// DimensionValuesGt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DimensionValuesGt(dimensionValues string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("dimension_values > ?", dimensionValues))
}

// DimensionValuesGte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DimensionValuesGte(dimensionValues string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("dimension_values >= ?", dimensionValues))
}

// DimensionValuesIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DimensionValuesIn(dimensionValues ...string) SpaceResourceQuerySet {
	if len(dimensionValues) == 0 {
		qs.db.AddError(errors.New("must at least pass one dimensionValues in DimensionValuesIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("dimension_values IN (?)", dimensionValues))
}

// DimensionValuesLike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DimensionValuesLike(dimensionValues string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("dimension_values LIKE ?", dimensionValues))
}

// DimensionValuesLt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DimensionValuesLt(dimensionValues string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("dimension_values < ?", dimensionValues))
}

// DimensionValuesLte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DimensionValuesLte(dimensionValues string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("dimension_values <= ?", dimensionValues))
}

// DimensionValuesNe is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DimensionValuesNe(dimensionValues string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("dimension_values != ?", dimensionValues))
}

// DimensionValuesNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DimensionValuesNotIn(dimensionValues ...string) SpaceResourceQuerySet {
	if len(dimensionValues) == 0 {
		qs.db.AddError(errors.New("must at least pass one dimensionValues in DimensionValuesNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("dimension_values NOT IN (?)", dimensionValues))
}

// DimensionValuesNotlike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) DimensionValuesNotlike(dimensionValues string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("dimension_values NOT LIKE ?", dimensionValues))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) GetUpdater() SpaceResourceUpdater {
	return NewSpaceResourceUpdater(qs.db)
}

// IdEq is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) IdEq(id int) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("id = ?", id))
}

// IdGt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) IdGt(id int) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("id > ?", id))
}

// IdGte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) IdGte(id int) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("id >= ?", id))
}

// IdIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) IdIn(id ...int) SpaceResourceQuerySet {
	if len(id) == 0 {
		qs.db.AddError(errors.New("must at least pass one id in IdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", id))
}

// IdLt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) IdLt(id int) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("id < ?", id))
}

// IdLte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) IdLte(id int) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("id <= ?", id))
}

// IdNe is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) IdNe(id int) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("id != ?", id))
}

// IdNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) IdNotIn(id ...int) SpaceResourceQuerySet {
	if len(id) == 0 {
		qs.db.AddError(errors.New("must at least pass one id in IdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", id))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) Limit(limit int) SpaceResourceQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) Offset(offset int) SpaceResourceQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs SpaceResourceQuerySet) One(ret *SpaceResource) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreateTime is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderAscByCreateTime() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("create_time ASC"))
}

// OrderAscByCreator is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderAscByCreator() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("creator ASC"))
}

// OrderAscByDimensionValues is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderAscByDimensionValues() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("dimension_values ASC"))
}

// OrderAscById is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderAscById() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByResourceId is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderAscByResourceId() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("resource_id ASC"))
}

// OrderAscByResourceType is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderAscByResourceType() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("resource_type ASC"))
}

// OrderAscBySpaceId is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderAscBySpaceId() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("space_id ASC"))
}

// OrderAscBySpaceTypeId is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderAscBySpaceTypeId() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("space_type_id ASC"))
}

// OrderAscByUpdateTime is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderAscByUpdateTime() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("update_time ASC"))
}

// OrderAscByUpdater is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderAscByUpdater() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("updater ASC"))
}

// OrderDescByCreateTime is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderDescByCreateTime() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("create_time DESC"))
}

// OrderDescByCreator is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderDescByCreator() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("creator DESC"))
}

// OrderDescByDimensionValues is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderDescByDimensionValues() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("dimension_values DESC"))
}

// OrderDescById is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderDescById() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByResourceId is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderDescByResourceId() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("resource_id DESC"))
}

// OrderDescByResourceType is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderDescByResourceType() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("resource_type DESC"))
}

// OrderDescBySpaceId is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderDescBySpaceId() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("space_id DESC"))
}

// OrderDescBySpaceTypeId is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderDescBySpaceTypeId() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("space_type_id DESC"))
}

// OrderDescByUpdateTime is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderDescByUpdateTime() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("update_time DESC"))
}

// OrderDescByUpdater is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) OrderDescByUpdater() SpaceResourceQuerySet {
	return qs.w(qs.db.Order("updater DESC"))
}

// ResourceIdEq is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdEq(resourceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_id = ?", resourceId))
}

// ResourceIdGt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdGt(resourceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_id > ?", resourceId))
}

// ResourceIdGte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdGte(resourceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_id >= ?", resourceId))
}

// ResourceIdIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdIn(resourceId ...string) SpaceResourceQuerySet {
	if len(resourceId) == 0 {
		qs.db.AddError(errors.New("must at least pass one resourceId in ResourceIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("resource_id IN (?)", resourceId))
}

// ResourceIdIsNotNull is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdIsNotNull() SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_id IS NOT NULL"))
}

// ResourceIdIsNull is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdIsNull() SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_id IS NULL"))
}

// ResourceIdLike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdLike(resourceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_id LIKE ?", resourceId))
}

// ResourceIdLt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdLt(resourceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_id < ?", resourceId))
}

// ResourceIdLte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdLte(resourceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_id <= ?", resourceId))
}

// ResourceIdNe is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdNe(resourceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_id != ?", resourceId))
}

// ResourceIdNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdNotIn(resourceId ...string) SpaceResourceQuerySet {
	if len(resourceId) == 0 {
		qs.db.AddError(errors.New("must at least pass one resourceId in ResourceIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("resource_id NOT IN (?)", resourceId))
}

// ResourceIdNotlike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceIdNotlike(resourceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_id NOT LIKE ?", resourceId))
}

// ResourceTypeEq is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceTypeEq(resourceType string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_type = ?", resourceType))
}

// ResourceTypeGt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceTypeGt(resourceType string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_type > ?", resourceType))
}

// ResourceTypeGte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceTypeGte(resourceType string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_type >= ?", resourceType))
}

// ResourceTypeIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceTypeIn(resourceType ...string) SpaceResourceQuerySet {
	if len(resourceType) == 0 {
		qs.db.AddError(errors.New("must at least pass one resourceType in ResourceTypeIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("resource_type IN (?)", resourceType))
}

// ResourceTypeLike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceTypeLike(resourceType string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_type LIKE ?", resourceType))
}

// ResourceTypeLt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceTypeLt(resourceType string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_type < ?", resourceType))
}

// ResourceTypeLte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceTypeLte(resourceType string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_type <= ?", resourceType))
}

// ResourceTypeNe is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceTypeNe(resourceType string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_type != ?", resourceType))
}

// ResourceTypeNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceTypeNotIn(resourceType ...string) SpaceResourceQuerySet {
	if len(resourceType) == 0 {
		qs.db.AddError(errors.New("must at least pass one resourceType in ResourceTypeNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("resource_type NOT IN (?)", resourceType))
}

// ResourceTypeNotlike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) ResourceTypeNotlike(resourceType string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("resource_type NOT LIKE ?", resourceType))
}

// SpaceIdEq is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceIdEq(spaceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_id = ?", spaceId))
}

// SpaceIdGt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceIdGt(spaceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_id > ?", spaceId))
}

// SpaceIdGte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceIdGte(spaceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_id >= ?", spaceId))
}

// SpaceIdIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceIdIn(spaceId ...string) SpaceResourceQuerySet {
	if len(spaceId) == 0 {
		qs.db.AddError(errors.New("must at least pass one spaceId in SpaceIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("space_id IN (?)", spaceId))
}

// SpaceIdLike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceIdLike(spaceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_id LIKE ?", spaceId))
}

// SpaceIdLt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceIdLt(spaceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_id < ?", spaceId))
}

// SpaceIdLte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceIdLte(spaceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_id <= ?", spaceId))
}

// SpaceIdNe is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceIdNe(spaceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_id != ?", spaceId))
}

// SpaceIdNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceIdNotIn(spaceId ...string) SpaceResourceQuerySet {
	if len(spaceId) == 0 {
		qs.db.AddError(errors.New("must at least pass one spaceId in SpaceIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("space_id NOT IN (?)", spaceId))
}

// SpaceIdNotlike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceIdNotlike(spaceId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_id NOT LIKE ?", spaceId))
}

// SpaceTypeIdEq is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceTypeIdEq(spaceTypeId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_type_id = ?", spaceTypeId))
}

// SpaceTypeIdGt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceTypeIdGt(spaceTypeId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_type_id > ?", spaceTypeId))
}

// SpaceTypeIdGte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceTypeIdGte(spaceTypeId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_type_id >= ?", spaceTypeId))
}

// SpaceTypeIdIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceTypeIdIn(spaceTypeId ...string) SpaceResourceQuerySet {
	if len(spaceTypeId) == 0 {
		qs.db.AddError(errors.New("must at least pass one spaceTypeId in SpaceTypeIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("space_type_id IN (?)", spaceTypeId))
}

// SpaceTypeIdLike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceTypeIdLike(spaceTypeId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_type_id LIKE ?", spaceTypeId))
}

// SpaceTypeIdLt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceTypeIdLt(spaceTypeId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_type_id < ?", spaceTypeId))
}

// SpaceTypeIdLte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceTypeIdLte(spaceTypeId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_type_id <= ?", spaceTypeId))
}

// SpaceTypeIdNe is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceTypeIdNe(spaceTypeId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_type_id != ?", spaceTypeId))
}

// SpaceTypeIdNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceTypeIdNotIn(spaceTypeId ...string) SpaceResourceQuerySet {
	if len(spaceTypeId) == 0 {
		qs.db.AddError(errors.New("must at least pass one spaceTypeId in SpaceTypeIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("space_type_id NOT IN (?)", spaceTypeId))
}

// SpaceTypeIdNotlike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) SpaceTypeIdNotlike(spaceTypeId string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("space_type_id NOT LIKE ?", spaceTypeId))
}

// UpdateTimeEq is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdateTimeEq(updateTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("update_time = ?", updateTime))
}

// UpdateTimeGt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdateTimeGt(updateTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("update_time > ?", updateTime))
}

// UpdateTimeGte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdateTimeGte(updateTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("update_time >= ?", updateTime))
}

// UpdateTimeLt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdateTimeLt(updateTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("update_time < ?", updateTime))
}

// UpdateTimeLte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdateTimeLte(updateTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("update_time <= ?", updateTime))
}

// UpdateTimeNe is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdateTimeNe(updateTime time.Time) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("update_time != ?", updateTime))
}

// UpdaterEq is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdaterEq(updater string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("updater = ?", updater))
}

// UpdaterGt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdaterGt(updater string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("updater > ?", updater))
}

// UpdaterGte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdaterGte(updater string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("updater >= ?", updater))
}

// UpdaterIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdaterIn(updater ...string) SpaceResourceQuerySet {
	if len(updater) == 0 {
		qs.db.AddError(errors.New("must at least pass one updater in UpdaterIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updater IN (?)", updater))
}

// UpdaterLike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdaterLike(updater string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("updater LIKE ?", updater))
}

// UpdaterLt is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdaterLt(updater string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("updater < ?", updater))
}

// UpdaterLte is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdaterLte(updater string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("updater <= ?", updater))
}

// UpdaterNe is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdaterNe(updater string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("updater != ?", updater))
}

// UpdaterNotIn is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdaterNotIn(updater ...string) SpaceResourceQuerySet {
	if len(updater) == 0 {
		qs.db.AddError(errors.New("must at least pass one updater in UpdaterNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updater NOT IN (?)", updater))
}

// UpdaterNotlike is an autogenerated method
// nolint: dupl
func (qs SpaceResourceQuerySet) UpdaterNotlike(updater string) SpaceResourceQuerySet {
	return qs.w(qs.db.Where("updater NOT LIKE ?", updater))
}

// SetCreateTime is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) SetCreateTime(createTime time.Time) SpaceResourceUpdater {
	u.fields[string(SpaceResourceDBSchema.CreateTime)] = createTime
	return u
}

// SetCreator is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) SetCreator(creator string) SpaceResourceUpdater {
	u.fields[string(SpaceResourceDBSchema.Creator)] = creator
	return u
}

// SetDimensionValues is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) SetDimensionValues(dimensionValues string) SpaceResourceUpdater {
	u.fields[string(SpaceResourceDBSchema.DimensionValues)] = dimensionValues
	return u
}

// SetId is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) SetId(id int) SpaceResourceUpdater {
	u.fields[string(SpaceResourceDBSchema.Id)] = id
	return u
}

// SetResourceId is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) SetResourceId(resourceId *string) SpaceResourceUpdater {
	u.fields[string(SpaceResourceDBSchema.ResourceId)] = resourceId
	return u
}

// SetResourceType is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) SetResourceType(resourceType string) SpaceResourceUpdater {
	u.fields[string(SpaceResourceDBSchema.ResourceType)] = resourceType
	return u
}

// SetSpaceId is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) SetSpaceId(spaceId string) SpaceResourceUpdater {
	u.fields[string(SpaceResourceDBSchema.SpaceId)] = spaceId
	return u
}

// SetSpaceTypeId is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) SetSpaceTypeId(spaceTypeId string) SpaceResourceUpdater {
	u.fields[string(SpaceResourceDBSchema.SpaceTypeId)] = spaceTypeId
	return u
}

// SetUpdateTime is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) SetUpdateTime(updateTime time.Time) SpaceResourceUpdater {
	u.fields[string(SpaceResourceDBSchema.UpdateTime)] = updateTime
	return u
}

// SetUpdater is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) SetUpdater(updater string) SpaceResourceUpdater {
	u.fields[string(SpaceResourceDBSchema.Updater)] = updater
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u SpaceResourceUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set SpaceResourceQuerySet

// ===== BEGIN of SpaceResource modifiers

// SpaceResourceDBSchemaField describes database schema field. It requires for method 'Update'
type SpaceResourceDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f SpaceResourceDBSchemaField) String() string {
	return string(f)
}

// SpaceResourceDBSchema stores db field names of SpaceResource
var SpaceResourceDBSchema = struct {
	Id              SpaceResourceDBSchemaField
	SpaceTypeId     SpaceResourceDBSchemaField
	SpaceId         SpaceResourceDBSchemaField
	ResourceType    SpaceResourceDBSchemaField
	ResourceId      SpaceResourceDBSchemaField
	DimensionValues SpaceResourceDBSchemaField
	Creator         SpaceResourceDBSchemaField
	CreateTime      SpaceResourceDBSchemaField
	Updater         SpaceResourceDBSchemaField
	UpdateTime      SpaceResourceDBSchemaField
}{

	Id:              SpaceResourceDBSchemaField("id"),
	SpaceTypeId:     SpaceResourceDBSchemaField("space_type_id"),
	SpaceId:         SpaceResourceDBSchemaField("space_id"),
	ResourceType:    SpaceResourceDBSchemaField("resource_type"),
	ResourceId:      SpaceResourceDBSchemaField("resource_id"),
	DimensionValues: SpaceResourceDBSchemaField("dimension_values"),
	Creator:         SpaceResourceDBSchemaField("creator"),
	CreateTime:      SpaceResourceDBSchemaField("create_time"),
	Updater:         SpaceResourceDBSchemaField("updater"),
	UpdateTime:      SpaceResourceDBSchemaField("update_time"),
}

// Update updates SpaceResource fields by primary key
// nolint: dupl
func (o *SpaceResource) Update(db *gorm.DB, fields ...SpaceResourceDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":               o.Id,
		"space_type_id":    o.SpaceTypeId,
		"space_id":         o.SpaceId,
		"resource_type":    o.ResourceType,
		"resource_id":      o.ResourceId,
		"dimension_values": o.DimensionValues,
		"creator":          o.Creator,
		"create_time":      o.CreateTime,
		"updater":          o.Updater,
		"update_time":      o.UpdateTime,
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

		return fmt.Errorf("can't update SpaceResource %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// SpaceResourceUpdater is an SpaceResource updates manager
type SpaceResourceUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewSpaceResourceUpdater creates new SpaceResource updater
// nolint: dupl
func NewSpaceResourceUpdater(db *gorm.DB) SpaceResourceUpdater {
	return SpaceResourceUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&SpaceResource{}),
	}
}

// ===== END of SpaceResource modifiers

// ===== END of all query sets