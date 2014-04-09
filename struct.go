package nego

import "encoding/json"

//ElasticMappingType Structure to use for mapping
type ElasticMappingType struct {
	Type                       string   `json:"type,omitempty"`
	IndexAnalyser              string   `json:"index_analyzer,omitempty"`
	SearchAnalyser             string   `json:"search_analyzer,omitempty"`
	Payloads                   *bool    `json:"payloads,omitempty"`
	PreserveSeparators         *bool    `json:"preserve_separators,omitempty"`
	PreservePositionIncrements *bool    `json:"preserve_position_increments,omitempty"`
	MaxInputLength             *uint    `json:"max_input_length,omitempty"`
	Index                      string   `json:"index,omitempty"`
	Boost                      *float64 `json:"boost,omitempty"`
	IncludeInAll               *bool    `json:"include_in_all,omitempty"`
	OmitNorms                  *bool    `json:"omit_norms,omitempty"`
}

//ElasticSettingToken Structure to use for settings
type ElasticSettingToken struct {
	Type      string   `json:"type"`
	Tokenizer string   `json:"tokenizer"`
	Filter    []string `json:"filter"`
}

//GetFromElasticStruct Structure when asking one index in ElasticSearch
type GetFromElasticStruct struct {
	Index   string          `json:"_index"`
	Type    string          `json:"_type"`
	ID      string          `json:"_id"`
	Version int64           `json:"_version"`
	Found   bool            `json:"_found"`
	Source  json.RawMessage `json:"_source"`
}

// SearchHitsStruct Got the info of an Hit from a Search
type SearchHitsStruct struct {
	Total    int     `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []struct {
		Index  string          `json:"_index"`
		Type   string          `json:"_type"`
		ID     string          `json:"_id"`
		Score  float64         `json:"_score"`
		Source json.RawMessage `json:"_source"`
	} `json:"hits"`
}

// SearchFromElasticStruct Struct to help to recover the results
type SearchFromElasticStruct struct {
	Took     float64 `json:"took"`
	TimedOut bool    `json:"timed_out"`
	Shards   struct {
		Total      float64 `json:"total"`
		Successful float64 `json:"successful"`
		Failed     float64 `json:"failed"`
	} `json:"_shards"`
	Hits SearchHitsStruct `json:"hits"`
}
