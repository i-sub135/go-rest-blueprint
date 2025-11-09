// Package common provides shared resources that can be reused across multiple features.
//
// This package contains:
//   - model: Shared GORM models and entities that are used by multiple features
//   - repository: Shared repository implementations for database operations
//   - utils: Common utility functions and helpers used throughout the application
//
// The common package follows the principle that if a component (model, repository, or utility)
// is used by more than one feature, it should be moved here to promote code reuse and
// maintain a single source of truth.
//
// Example usage:
//
//	import "github.com/i-sub135/go-rest-blueprint/source/common/model"
//	import "github.com/i-sub135/go-rest-blueprint/source/common/repository"
//	import "github.com/i-sub135/go-rest-blueprint/source/common/utils/http_resp_utils"
//
// Structure:
//
//	common/
//	├── model/           # Shared GORM models and database entities
//	├── repository/      # Shared repository implementations
//	└── glob_utils/           # Common utility functions and helpers
package common
