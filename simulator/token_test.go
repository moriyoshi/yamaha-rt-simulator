package simulator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenTrieAdd(t *testing.T) {
	trie := TokenTrie{}
	assert.Equal(t, true, trie.Add([]Token{TShow, TConfig}, 2, "aaa"))
	assert.Equal(t, false, trie.Add([]Token{TShow, TConfig}, 2, "bbb"))
	assert.Equal(t, true, trie.Add([]Token{TShow, TCopyright}, 2, "ccc"))
	{
		tis, meta, nexts := trie.LookupByString(nil, []string{"show"})
		assert.Equal(t, 1, len(tis))
		assert.Equal(t, nil, meta)
		assert.ElementsMatch(t, []Token{TConfig, TCopyright}, nexts)
	}
	{
		tis, meta, nexts := trie.LookupByString(nil, []string{"show", "config"})
		assert.Equal(t, 2, len(tis))
		assert.Equal(t, "aaa", meta)
		assert.Equal(t, []Token{}, nexts)
	}
	{
		tis, meta, nexts := trie.LookupByString(nil, []string{"show", "copyright"})
		assert.Equal(t, 2, len(tis))
		assert.Equal(t, "ccc", meta)
		assert.Equal(t, []Token{}, nexts)
	}
}

func TestTokenTrieLookupStringC(t *testing.T) {
	trie := TokenTrie{}
	assert.Equal(t, true, trie.Add([]Token{TPp, TSelect}, 2, "aaa"))
	assert.Equal(t, true, trie.Add([]Token{TPp, TServer}, 2, "bbb"))
	assert.Equal(t, true, trie.Add([]Token{TPppoe, TUse}, 2, "ccc"))
	assert.Equal(t, true, trie.Add([]Token{TPpp, TCcp}, 2, "ddd"))
	{
		tis, cands, candsForDisp := trie.LookupByStringC(nil, []string{"pp"})
		assert.Equal(t, 0, len(tis))
		assert.ElementsMatch(t, []string{"pp", "ppp", "pppoe"}, cands)
		assert.ElementsMatch(t, []string{"pp", "ppp", "pppoe"}, candsForDisp)
	}
	{
		tis, cands, candsForDisp := trie.LookupByStringC(nil, []string{"pp", "se"})
		assert.Equal(t, []TokenInstance{{TPp, "pp"}}, tis)
		assert.ElementsMatch(t, []string{"select", "server"}, cands)
		assert.ElementsMatch(t, []string{"select", "server"}, candsForDisp)
	}
}
