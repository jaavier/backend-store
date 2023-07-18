package storage

type Memory struct {
	String        string
	IntNumber     int
	Float64       float64
	Boolean       bool
	StringList    []string
	IntNumberList []int
	Float64List   []float64
	BoolList      []bool
}

type MemoryStorage struct {
	StringMap    map[string]Memory
	IntMap       map[int]Memory
	Float64Map   map[float64]Memory
	InterfaceMap map[interface{}]Memory
}

func (s *MemoryStorage) GetString(key string) (string, bool) {
	memory, exists := s.StringMap[key]
	return memory.String, exists
}

func (s *MemoryStorage) GetInt(key int) (int, bool) {
	memory, exists := s.IntMap[key]
	return memory.IntNumber, exists
}

func (s *MemoryStorage) GetFloat64(key float64) (float64, bool) {
	memory, exists := s.Float64Map[key]
	return memory.Float64, exists
}

func (s *MemoryStorage) GetBool(key interface{}) (bool, bool) {
	memory, exists := s.InterfaceMap[key]
	return memory.Boolean, exists
}

func (s *MemoryStorage) SetString(key string, value Memory) {
	s.StringMap[key] = value
}

func (s *MemoryStorage) SetInt(key int, value Memory) {
	s.IntMap[key] = value
}

func (s *MemoryStorage) SetFloat64(key float64, value Memory) {
	s.Float64Map[key] = value
}

func (s *MemoryStorage) SetBool(key interface{}, value Memory) {
	s.InterfaceMap[key] = value
}

func (s *MemoryStorage) SetStringList(key string, value []string) {
	memory := s.StringMap[key]
	memory.StringList = append(memory.StringList, value...)
	s.StringMap[key] = memory
}

func (s *MemoryStorage) SetIntList(key int, value []int) {
	memory := s.IntMap[key]
	memory.IntNumberList = append(memory.IntNumberList, value...)
	s.IntMap[key] = memory
}

func (s *MemoryStorage) SetFloat64List(key float64, value []float64) {
	memory := s.Float64Map[key]
	memory.Float64List = append(memory.Float64List, value...)
	s.Float64Map[key] = memory
}

func (s *MemoryStorage) SetBoolList(key interface{}, value []bool) {
	memory := s.InterfaceMap[key]
	memory.BoolList = append(memory.BoolList, value...)
	s.InterfaceMap[key] = memory
}

func (s *MemoryStorage) CreateMemory(kindOfMemory string) {
	switch kindOfMemory {
	case "string":
		s.StringMap = make(map[string]Memory)
	case "int":
		s.IntMap = make(map[int]Memory)
	case "float64":
		s.Float64Map = make(map[float64]Memory)
	default:
		s.InterfaceMap = make(map[interface{}]Memory)
	}
}
