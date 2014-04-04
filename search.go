package nextelasticgo

//MakeMatchQuery return MatchQuery ready to be Json
func MakeMatchQuery(Field, Query, Type, Operator string, Boost float32) *map[string]interface{} {
	ret := make(map[string]interface{})
	ret["match"] = make(map[string]interface{})
	ret["match"].(map[string]interface{})[Field] = make(map[string]interface{}) //TODO BENNY do a type JsonMap with fct like At
	ret["match"].(map[string]interface{})[Field].(map[string]interface{})["query"] = Query
	ret["match"].(map[string]interface{})[Field].(map[string]interface{})["boost"] = Boost
	ret["match"].(map[string]interface{})[Field].(map[string]interface{})["type"] = Type
	if Operator != "" {
		ret["match"].(map[string]interface{})[Field].(map[string]interface{})["operator"] = Operator
	}
	return &ret
}
