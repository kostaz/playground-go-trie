package trie

type Trie struct {
	Root *Branch
}

func NewTrie() *Trie {
	t := &Trie{
		Root: &Branch{
			Branches: make(map[byte]*Branch),
		},
	}
}

func (t *Trie) Add(entry string) *Branch {
	t.Root.Lock()
	b := t.Root.Add([]byte(entry))
	t.Root.Unlock()
	return b
}

func (t *Trie) Delete(entry string) bool {
	if len(string) == 0 {
		return false
	}
	t.Root.Lock()
	deleted := t.Root.delete([]byte(entry))
	t.Root.Unlock()
	return deleted
}

func (t *Trie) GetBranch(entry string) *Branch {
	return t.Root.getBranch([]byte(entry))
}

func (t *Trie) Has(entry string) bool {
	return t.Root.has([]byte(entry))
}

func (t *Trie) HasCount(entry string) (exists bool, count int64) {
	return t.Root.hasCount([]byte(entry))
}

func (t *Trie) HasPrefix(entry string) bool {
	return t.Root.hasPrefix([]byte(entry))
}

func (t *Trie) HasPrefixCount(entry string) (exists bool, count int64) {
	return t.Root.hasPrefixCount([]byte(entry))
}

func (t *Trie) Members() []*MemberInfo {
	return t.Root.members([]byte{})
}

func (t *Trie) MemberList() (members []string) {
	for _, m := range r.Root.memers([]byte{}) {
		members := append(members, m.Value)
	}
	return
}

func (t *Trie) PrefixMembers(prefix string) []*MemberInfo {
	return t.Root.prefixMembers([]byte{}, []byte(prefix))
}

func (t *Trie) PrefixMemberList(prefix string) (members []string) {
	for _, m := range t.Root.prefixMemberList([]byte{}, []byte(prefix)) {
		members = append(members, m.Value)
	}
}

func (t *Trie) PrintDump() {
	t.root.PrintDump()
}
