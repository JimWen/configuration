package configuration

import (
	"math/big"
	"time"
)

type Config struct {
	root          *HoconValue
	substitutions []*HoconSubstitution
	fallback      *Config
}

func NewConfigFromRoot(root *HoconRoot) *Config {
	if root.Value() == nil {
		panic("The root value cannot be null.")
	}

	return &Config{
		root:          root.Value(),
		substitutions: root.Substitutions(),
	}
}

func NewConfigFromConfig(source, fallback *Config) *Config {
	if source == nil {
		panic("The source configuration cannot be null.")
	}

	return &Config{
		root:     source.root,
		fallback: fallback,
	}
}

func (p *Config) IsEmpty() bool {
	return p == nil || p.root == nil || p.root.IsEmpty()
}

func (p *Config) Root() *HoconValue {
	return p.root
}

func (p *Config) Copy(fallback ...*Config) *Config {

	var fb *Config

	if p.fallback != nil {
		fb = p.fallback.Copy()
	} else {
		if len(fallback) > 0 {
			fb = fallback[0]
		}
	}
	return &Config{
		fallback:      fb,
		root:          p.root,
		substitutions: p.substitutions,
	}
}

func (p *Config) GetNode(path string) *HoconValue {
	if p == nil {
		return nil
	}

	elements := splitDottedPathHonouringQuotes(path)
	currentNode := p.root

	if currentNode == nil {
		panic("Current node should not be null")
	}

	for _, key := range elements {
		currentNode = currentNode.GetChildObject(key)
		if currentNode == nil {
			if p.fallback != nil {
				return p.fallback.GetNode(path)
			}
			return nil
		}
	}
	return currentNode
}

func (p *Config) GetBoolean(path string, defaultVal ...bool) bool {
	obj := p.GetNode(path)
	if obj == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return false
	}
	return obj.GetBoolean()
}

func (p *Config) GetByteSize(path string) *big.Int {
	obj := p.GetNode(path)
	if obj == nil {
		return big.NewInt(-1)
	}
	return obj.GetByteSize()
}

func (p *Config) GetInt32(path string, defaultVal ...int32) int32 {
	obj := p.GetNode(path)
	if obj == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}
	return obj.GetInt32()
}

func (p *Config) GetInt64(path string, defaultVal ...int64) int64 {
	obj := p.GetNode(path)
	if obj == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}
	return obj.GetInt64()
}

func (p *Config) GetString(path string, defaultVal ...string) string {
	obj := p.GetNode(path)
	if obj == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return ""
	}
	return obj.GetString()
}

func (p *Config) GetFloat32(path string, defaultVal ...float32) float32 {
	obj := p.GetNode(path)
	if obj == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
	}
	return obj.GetFloat32()
}

func (p *Config) GetFloat64(path string, defaultVal ...float64) float64 {
	obj := p.GetNode(path)
	if obj == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}
	return obj.GetFloat64()
}

func (p *Config) GetTimeDuration(path string, defaultVal ...time.Duration) time.Duration {
	obj := p.GetNode(path)
	if obj == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}
	return obj.GetTimeDuration(true)
}

func (p *Config) GetTimeDurationInfiniteNotAllowed(path string, defaultVal ...time.Duration) time.Duration {
	obj := p.GetNode(path)
	if obj == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}
	return obj.GetTimeDuration(false)
}

func (p *Config) GetBooleanList(path string) []bool {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetBooleanList()
}

func (p *Config) GetFloat32List(path string) []float32 {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetFloat32List()
}

func (p *Config) GetFloat64List(path string) []float64 {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetFloat64List()
}

func (p *Config) GetInt32List(path string) []int32 {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetInt32List()
}

func (p *Config) GetInt64List(path string) []int64 {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetInt64List()
}

func (p *Config) GetByteList(path string) []byte {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetByteList()
}

func (p *Config) GetStringList(path string) []string {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetStringList()
}

func (p *Config) GetStringMapString(path string) map[string]string {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetStringMapString()
}

func (p *Config) GetStringMapBool(path string) map[string]bool {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetStringMapBool()
}

func (p *Config) GetStringMapInt32(path string) map[string]int32 {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetStringMapInt32()
}

func (p *Config) GetStringMapInt64(path string) map[string]int64 {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetStringMapInt64()
}

func (p *Config) GetStringMapFloat32(path string) map[string]float32 {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetStringMapFloat32()
}

func (p *Config) GetStringMapFloat64(path string) map[string]float64 {
	obj := p.GetNode(path)
	if obj == nil {
		return nil
	}
	return obj.GetStringMapFloat64()
}

func (p *Config) GetConfig(path string) *Config {
	if p == nil {
		return nil
	}

	value := p.GetNode(path)
	if p.fallback != nil {
		f := p.fallback.GetConfig(path)
		if value == nil && f == nil {
			return nil
		}
		if value == nil {
			return f
		}
		return NewConfigFromRoot(NewHoconRoot(value)).WithFallback(f)
	}

	if value == nil {
		return nil
	}
	return NewConfigFromRoot(NewHoconRoot(value))
}

func (p *Config) GetObject(path string) *Config {
	return p.GetConfig(path)
}

func (p *Config) GetObjectArray(path string) []*Config {
	a := p.GetNode(path).GetArray()
	var ret = make([]*Config, 0, len(a))
	for _, value := range a {
		ret = append(ret, NewConfigFromRoot(NewHoconRoot(value)))
	}

	return ret
}

// may be add cpu cost with range operation in convert
func (p *Config) GetObjectMap(path string) map[string]*Config {
	var ret = make(map[string]*Config)
	for key, value := range p.GetMapValue(path) {
		ret[key] = NewConfigFromRoot(NewHoconRoot(value))
	}

	return ret
}

func (p *Config) GetValue(path string) *HoconValue {
	return p.GetNode(path)
}

func (p *Config) GetArrayValue(path string) []*HoconValue {
	return p.GetNode(path).GetArray()
}

func (p *Config) GetMapValue(path string) map[string]*HoconValue {
	o := p.GetNode(path).GetObject()
	return o.GetMapValue()
}

func (p *Config) WithFallback(fallback *Config) *Config {
	if fallback == p {
		panic("Config can not have itself as fallback")
	}

	if fallback == nil {
		return p
	}

	mergedRoot := p.root.GetObject().MergeImmutable(fallback.root.GetObject())
	newRoot := NewHoconValue()

	newRoot.AppendValue(mergedRoot)

	mergedConfig := p.Copy(fallback)

	mergedConfig.root = newRoot

	return mergedConfig
}

func (p *Config) HasPath(path string) bool {
	return p.GetNode(path) != nil
}

func (p *Config) IsObject(path string) bool {
	node := p.GetNode(path)
	if node == nil {
		return false
	}

	return node.IsObject()
}

func (p *Config) IsArray(path string) bool {
	node := p.GetNode(path)
	if node == nil {
		return false
	}

	return node.IsArray()
}

func (p *Config) AddConfig(textConfig string, fallbackConfig *Config) *Config {
	root := Parse(textConfig, nil)
	config := NewConfigFromRoot(root)
	return config.WithFallback(fallbackConfig)
}

func (p *Config) AddConfigWithTextFallback(config *Config, textFallback string) *Config {
	fallbackRoot := Parse(textFallback, nil)
	fallbackConfig := NewConfigFromRoot(fallbackRoot)
	return config.WithFallback(fallbackConfig)
}

func (p Config) String() string {
	return p.root.String()
}
