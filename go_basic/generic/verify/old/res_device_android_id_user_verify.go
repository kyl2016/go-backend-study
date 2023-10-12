package verify

import (
	"bytepower_gold/base"
	"bytepower_gold/utility"
)

type DeviceAndroidIDUserVerify struct {
	userVerify

	tableName       struct{} `pg:"device_android_id_user_verify,alias:t"`
	DeviceAndroidID string   `pg:"device_android_id"`

	identifyType  string `pg:"-"`
	identifyField string `pg:"-"`
	idPrefix      string `pg:"-"`
}

// func init() {
// 	registerVerify(func(ids base_service.IDsParam) verifyI {
// 		if formatID(ids.DeviceAndroidID) == "" {
// 			return nil
// 		}

// 		return DeviceAndroidIDUserVerify{
// 			DeviceAndroidID: ids.DeviceAndroidID,
// 			identifyType:    _AndroidIDIdentifyType,
// 			identifyField:   _DeviceAndroidIDField,
// 			idPrefix:        _DeviceAndroidIDUserIdentifyPrefix,
// 		}
// 	})
// }

func (identify DeviceAndroidIDUserVerify) getIdentifyInfo() identifyInfo {
	return identifyInfo{identify.identifyType, identify.DeviceAndroidID}
}

func (identify DeviceAndroidIDUserVerify) loadUserVerify(db base.DBService) (userVerify, bool, utility.Error) {
	if err := db.Model(&identify).Where(identify.identifyField+"= ?", identify.getIdentifyInfo().ID).Select(); err != nil {
		return userVerify{}, false, base.WrapDBErrorOnLoad(err)
	}
	return identify.userVerify, true, nil
}

func (identify DeviceAndroidIDUserVerify) updateUserIdentityStatus(db base.DBService, uv userVerify) utility.Error {
	identify.userVerify = uv
	identify.ID = identify.idPrefix + utility.GenerateUUID(4)
	_, err := db.Model(&identify).
		OnConflict("("+identify.identifyField+") DO UPDATE").
		Set(_IdentityStatusField+"=?", uv.IdentityVerified).
		Set(_UpdateTimeField+"=?", uv.UpdateTime).
		Insert()
	return base.WrapDBErrorOnExec(err)
}

func (identify DeviceAndroidIDUserVerify) increaseWithdrawNum(db base.DBService, uv userVerify) utility.Error {
	identify.userVerify = uv
	identify.ID = identify.idPrefix + utility.GenerateUUID(4)
	_, err := db.Model(&identify).
		OnConflict("("+identify.identifyField+") DO UPDATE").
		Set(_WithdrawTotalField+"=t."+_WithdrawTotalField+"+?", uv.WithdrawTotal).
		Set(_UpdateTimeField+"=?", uv.UpdateTime).
		Insert()
	return base.WrapDBErrorOnExec(err)
}
