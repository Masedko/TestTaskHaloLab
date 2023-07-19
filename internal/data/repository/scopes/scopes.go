package scopes

import (
	"TestTaskHaloLab/internal/core/dtos"
	"gorm.io/gorm"
)

// TopN Scope for finding top N species
func TopN(n int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("quantity DESC").Limit(n)
	}
}

// InCoordinates Scope for finding sensors in given coordinates
func InCoordinates(coordinates dtos.Coordinates) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Where("coordinate_x BETWEEN ? AND ?", coordinates.XMin, coordinates.XMax).
			Where("coordinate_y BETWEEN ? AND ?", coordinates.YMin, coordinates.YMax).
			Where("coordinate_z BETWEEN ? AND ?", coordinates.ZMin, coordinates.ZMax)
	}
}

// GetWithGroupName Scope for finding sensors with given groupName
func GetWithGroupName(groupName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Where("sensor_name = ?", groupName)
	}
}

// GetWithGroup Scope for finding sensors with given group
func GetWithGroup(group *dtos.Group) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Where("sensor_name = ? AND sensor_index = ?", group.Name, group.Index)
	}
}

// GetWithGroups Scope for finding sensors with given group
func GetWithGroups(groups []dtos.Group) func(db *gorm.DB) *gorm.DB {
	names, indexes := getGroupNamesAndIndexes(groups)
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Where("sensor_name IN ? AND sensor_index IN ?", names, indexes)
	}
}

func GetWithTimeInterval(timeInterval *dtos.TimeInterval) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Where("created_at BETWEEN ? AND ?", timeInterval.StarTime, timeInterval.EndTime)
	}
}

// Util function for scope
func getGroupNamesAndIndexes(groups []dtos.Group) ([]string, []int) {
	names := make([]string, len(groups))
	indexes := make([]int, len(groups))
	for i, group := range groups {
		names[i] = group.Name
	}
	return names, indexes
}
