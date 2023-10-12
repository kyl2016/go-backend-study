package verify

import (
	"bytepower_gold/base"
	"bytepower_gold/utility"
)

type DeviceCuidUserVerify struct {
	userVerify

	tableName  struct{} `pg:"device_cuid_user_verify,alias:t"`
	DeviceCuid string   `pg:"device_cuid"`

	identifyType  string `pg:"-"`
	identifyField string `pg:"-"`
	idPrefix      string `pg:"-"`
}

// func init() {
// 	registerVerify(func(ids base_service.IDsParam) verifyI {
// 		if formatID(ids.DeviceCuid) == "" {
// 			return nil
// 		}
// 		return DeviceCuidUserVerify{
// 			DeviceCuid:    ids.DeviceCuid,
// 			identifyType:  _CuidIdentifyType,
// 			identifyField: _DeviceCuidField,
// 			idPrefix:      _DeviceCuidUserIdentifyIDPrefix,
// 		}
// 	})
// }

func (identify DeviceCuidUserVerify) getIdentifyInfo() identifyInfo {
	return identifyInfo{identify.identifyType, identify.DeviceCuid}
}

func (identify DeviceCuidUserVerify) loadUserVerify(db base.DBService) (userVerify, bool, utility.Error) {
	if err := db.Model(&identify).Where(identify.identifyField+"= ?", identify.getIdentifyInfo().ID).Select(); err != nil {
		return userVerify{}, false, base.WrapDBErrorOnLoad(err)
	}
	return identify.userVerify, true, nil
}

func (identify DeviceCuidUserVerify) updateUserIdentityStatus(db base.DBService, uv userVerify) utility.Error {
	identify.userVerify = uv
	identify.ID = identify.idPrefix + utility.GenerateUUID(4)
	_, err := db.Model(&identify).
		OnConflict("("+identify.identifyField+") DO UPDATE").
		Set(_IdentityStatusField+"=?", uv.IdentityVerified).
		Set(_UpdateTimeField+"=?", uv.UpdateTime).
		Insert()
	return base.WrapDBErrorOnExec(err)
}

func (identify DeviceCuidUserVerify) increaseWithdrawNum(db base.DBService, uv userVerify) utility.Error {
	identify.userVerify = uv
	identify.ID = identify.idPrefix + utility.GenerateUUID(4)
	_, err := db.Model(&identify).
		OnConflict("("+identify.identifyField+") DO UPDATE").
		Set(_WithdrawTotalField+"=t."+_WithdrawTotalField+"+?", uv.WithdrawTotal).
		Set(_UpdateTimeField+"=?", uv.UpdateTime).
		Insert()
	return base.WrapDBErrorOnExec(err)
}
