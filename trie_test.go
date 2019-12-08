package trie

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

var (
    mar        = []rune("Mar")
    marlon     = []rune("Marlon")
    marcelitte = []rune("Marcelitte")

    tuesday = []rune("Tuesday")
)

func TestInsertAndSearch(t *testing.T) {
    tr := New()

    tr.Insert(marlon)
    tr.Insert(marcelitte)

    assert.True(t, tr.Search(marlon))
    assert.True(t, tr.Search(marcelitte))

    assert.False(t, tr.Search(tuesday))
    assert.False(t, tr.Search(mar))
}

func TestDelete(t *testing.T) {
    tr := New()

    tr.Insert(marlon)
    tr.Insert(marcelitte)

    assert.True(t, tr.Delete(marlon))
    assert.False(t, tr.Delete(mar))
    assert.False(t, tr.Delete(tuesday))
    assert.False(t, tr.Search(marlon))
    assert.True(t, tr.Search(marcelitte))
}
