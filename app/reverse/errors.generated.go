package reverse

import "github.com/Mortaza-Karimi/Xray-core/blob/main/common/errors"

type errPathObjHolder struct{}

func newError(values ...interface{}) *errors.Error {
	return errors.New(values...).WithPathObj(errPathObjHolder{})
}
