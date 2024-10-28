package tomlx

// Field 表示消息中的字段
type Field struct {
	Type   string `toml:"type"`
	Number int    `toml:"number"`
	Key    bool   `toml:"key"`
}

// Message 表示一个消息
type Message struct {
	Fields map[string]Field `toml:"-"`
}

// Enum 表示一个枚举
type Enum struct {
	Values map[string]int `toml:",inline"`
}

// Config 表示整个配置
type Config struct {
	GoPackage string             `toml:"go_package"`
	PbOut     string             `toml:"pb_out"`
	Messages  map[string]Message `toml:"msg"`
	Enums     map[string]Enum    `toml:"enum"`
}

// UnmarshalTOML 自定义反序列化方法，用于处理嵌套的字段
func (m *Message) UnmarshalTOML(data interface{}) error {
	m.Fields = make(map[string]Field)
	if rawMap, ok := data.(map[string]interface{}); ok {
		for key, value := range rawMap {
			if fieldMap, ok := value.(map[string]interface{}); ok {
				field := Field{}
				if t, ok := fieldMap["type"].(string); ok {
					field.Type = t
				}
				if n, ok := fieldMap["number"].(int64); ok {
					field.Number = int(n)
				}
				if k, ok := fieldMap["key"].(bool); ok {
					field.Key = k
				}
				m.Fields[key] = field
			}
		}
	}
	return nil
}
