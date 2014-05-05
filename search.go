package nego

import (
	"github.com/bennyscetbun/jsongo"
)

//MakeMatchQuery return MatchQuery ready to be Json
func MakeMatchQuery(Field, Query, Type, Operator string, Boost float32) *jsongo.JSONNode {
	ret := jsongo.JSONNode{}
	ret.At("match", Field, "query").Val(Query)
	ret.At("match", Field, "boost").Val(Boost)
	ret.At("match", Field, "type").Val(Type)
	if Operator != "" {
		ret.At("match", Field, "operator").Val(Operator)
	}
	return &ret
}
