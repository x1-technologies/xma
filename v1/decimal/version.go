// Package: xma/v1/decimal
// File: version.go

// Acknowledgement:
//   This implementation is a reference
//   Referenced Project Name: shopspring/decimal
//   Referenced Project Owner(s): Jason Biegel (from New York City, US), Mateusz Wos (from Cracow, Poland)
//   Referenced Project URL: https://github.com/shopspring/decimal
//   Referencer: x1-tech.com
//   Reference Project: xma
//   Feel grateful to the contributors

package decimal

const (
	Version        = "1.0.10000"
	LastUpdateDate = "2025-10-04Z"
)

const (
	SyncModeFull        = "FULL"
	SyncModeBranchMerge = "BRANCH_MERGE"
)

type CodeSyncRecord struct {
	XmaDecimalVersion        string
	ShopSpringDecimalVersion string
	SyncMode                 string
	SyncDate                 string
}

var CodeSyncRecords = []CodeSyncRecord{
	{
		XmaDecimalVersion:        "1.0.10000",
		ShopSpringDecimalVersion: "1.4.0",
		SyncMode:                 SyncModeFull,
		SyncDate:                 "2025-10-04Z",
	},
}
