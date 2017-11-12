package trie

type MemberInfo struct {
	Value string
	Count int64
}

func (m *MemberInfo) String() string {
	return fmt.Sprintf("%s(%d)", m.Value, m.Count)
}

type Branch struct {
	sync.RWMutex
	Branches  map[byte]*Branch
	LeafValue []byte
	End       bool
	Count     int64
}

func (b *Branch) NewBranch() *Branch {
	return &Branch{
		Branches: make(map[byte]*Branch),
		Count:    0,
}
