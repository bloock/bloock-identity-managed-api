package repository

import "github.com/iden3/iden3comm"

type VerificationZkpRepository interface {
	PackageManager() iden3comm.PackageManager
}
