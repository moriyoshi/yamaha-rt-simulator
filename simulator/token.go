package simulator

import (
	"strings"
)

type Token interface {
	Candidates(interface{}, string, []string, []string) ([]string, []string)
	Matches(interface{}, string) (string, bool)
}

type AnyToken struct{}

func (*AnyToken) Candidates(_ interface{}, _ string, candidates []string, candidatesForDisplay []string) ([]string, []string) {
	return candidates, candidatesForDisplay
}

func (*AnyToken) Matches(_ interface{}, im string) (string, bool) {
	return im, true
}

type EitherToken []Token

func (et *EitherToken) Candidates(ctx interface{}, im string, candidates []string, candidatesForDisplay []string) ([]string, []string) {
	for _, t := range *et {
		candidates, candidatesForDisplay = t.Candidates(ctx, im, candidates, candidatesForDisplay)
	}
	return candidates, candidatesForDisplay
}

func (et *EitherToken) Matches(ctx interface{}, im string) (canon string, ok bool) {
	for _, t := range *et {
		canon, ok = t.Matches(ctx, im)
		if ok {
			break
		}
	}
	return
}

type SymbolToken struct {
	Image string
}

func (sym *SymbolToken) Candidates(_ interface{}, im string, candidates []string, candidatesForDisplay []string) ([]string, []string) {
	im = strings.ToLower(im)
	if strings.HasPrefix(sym.Image, im) {
		candidates = append(candidates, sym.Image)
		candidatesForDisplay = append(candidatesForDisplay, sym.Image)
	}
	return candidates, candidatesForDisplay
}

func (sym *SymbolToken) Matches(_ interface{}, im string) (string, bool) {
	im = strings.ToLower(im)
	if im == sym.Image {
		return im, true
	} else {
		return "", false
	}
}

type EqualArgToken struct {
	Arg string
}

func (sym *EqualArgToken) Candidates(_ interface{}, im string, candidates []string, candidatesForDisplay []string) ([]string, []string) {
	im = strings.ToLower(im)
	if strings.HasPrefix(sym.Arg, im) {
		candidates = append(candidates, sym.Arg)
		candidatesForDisplay = append(candidatesForDisplay, sym.Arg)
	}
	return candidates, candidatesForDisplay
}

func (sym *EqualArgToken) Matches(_ interface{}, im string) (string, bool) {
	im = strings.ToLower(im)
	if strings.HasPrefix(im, sym.Arg) && len(im) > len(sym.Arg) {
		return im, true
	} else {
		return "", false
	}
}

type TokenInstance struct {
	Token Token
	Image string
}

type TokenTrieNode struct {
	Value []Token
	Subs  map[Token]*TokenTrieNode
	Meta  interface{}
}

type TokenTrie struct {
	Root TokenTrieNode
}

type TokenTrieOp int

const (
	TokenTrieOpNone TokenTrieOp = iota
	TokenTrieOpAdd
	TokenTrieOpUpdate
	TokenTrieOpAddOrUpdate
)

func (ttn *TokenTrieNode) Lookup(val []Token, i int) (int, *TokenTrieNode) {
	if i < len(val) && ttn.Subs != nil {
		if sttn := ttn.Subs[val[i]]; sttn != nil {
			return sttn.Lookup(val, i+1)
		}
	}
	return i, ttn
}

func (ttn *TokenTrieNode) LookupByString(ctx interface{}, val []string, canons []string) ([]string, *TokenTrieNode) {
	if len(val) > 0 && ttn.Subs != nil {
		for t, sttn := range ttn.Subs {
			if canon, ok := t.Matches(ctx, val[0]); ok {
				return sttn.LookupByString(ctx, val[1:], append(canons, canon))
			}
		}
	}
	return canons, ttn
}

func (ttn *TokenTrieNode) LookupByStringC(ctx interface{}, val []string, canons []string) ([]string, *TokenTrieNode, []string, []string) {
	var candidates []string
	var candidatesForDisplay []string
	var sttn *TokenTrieNode
	if len(val) > 0 && ttn.Subs != nil {
		for t, _sttn := range ttn.Subs {
			var newCandidates []string
			var canon string
			var matched bool
			if len(val) > 1 {
				canon, matched = t.Matches(ctx, val[0])
			}
			if matched {
				candidates = []string{canon}
				sttn = _sttn
				break
			} else {
				newCandidates, candidatesForDisplay = t.Candidates(ctx, val[0], candidates, candidatesForDisplay)
			}
			if len(candidates) == 0 && len(newCandidates) == 1 {
				sttn = _sttn
			}
			candidates = newCandidates
		}
	}
	if len(candidates) == 1 {
		return sttn.LookupByStringC(ctx, val[1:], append(canons, candidates[0]))
	}
	return canons, ttn, candidates, candidatesForDisplay
}

func (tt *TokenTrie) Add(val []Token, n int, meta interface{}) bool {
	i, ttn := tt.Root.Lookup(val, 0)
	if i == len(val) {
		return false
	}
	if i > n {
		return false
	}
	for {
		if i >= n {
			ttn.Meta = meta
		}
		if i >= len(val) {
			break
		}
		nttn := &TokenTrieNode{
			Value: val[:i+1],
			Subs:  nil,
			Meta:  nil,
		}
		if ttn.Subs == nil {
			ttn.Subs = map[Token]*TokenTrieNode{}
		}
		ttn.Subs[val[i]] = nttn
		ttn = nttn
		i++
	}
	return true
}

func (tt *TokenTrie) LookupByString(ctx interface{}, val []string) ([]TokenInstance, interface{}) {
	canons, ttn, _, _ := tt.Root.LookupByStringC(ctx, val, make([]string, 0, len(val)))
	tis := make([]TokenInstance, len(ttn.Value))
	for i := 0; i < len(ttn.Value); i++ {
		tis[i] = TokenInstance{ttn.Value[i], canons[i]}
	}
	if ttn == nil {
		return tis, nil
	} else {
		return tis, ttn.Meta
	}
}

func (tt *TokenTrie) LookupByStringC(ctx interface{}, val []string) ([]TokenInstance, []string, []string) {
	canons, ttn, candidates, candidatesForDisplay := tt.Root.LookupByStringC(ctx, val, make([]string, 0, len(val)))
	tis := make([]TokenInstance, len(ttn.Value))
	for i := 0; i < len(ttn.Value); i++ {
		tis[i] = TokenInstance{ttn.Value[i], canons[i]}
	}
	return tis, candidates, candidatesForDisplay
}
