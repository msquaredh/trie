package trie

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

var (
    mar        = []rune("Mar")
    marl       = []rune("Marl")
    marlon     = []rune("Marlon")
    marcelitte = []rune("Marcelitte")

    tuesday = []rune("Tuesday")
)

func TestInsertAndSearch(t *testing.T) {
    tr := New()

    tr.Insert(marlon)
    tr.Insert(marcelitte)
    // Make sure we see the inserts
    assert.True(t, tr.Search(marlon))
    assert.True(t, tr.Search(marcelitte))
    // verify things that aren't inserted aren't there
    assert.False(t, tr.Search(tuesday))
    assert.False(t, tr.Search(mar))
    // insert an existing prefix
    tr.Insert(mar)
    assert.True(t, tr.Search(mar))

}

func TestDelete(t *testing.T) {
    tr := New()

    tr.Insert(marlon)
    tr.Insert(marcelitte)
    tr.Insert(mar)

    assert.True(t, tr.Delete(marlon))
    assert.False(t, tr.Search(marlon))
    assert.True(t, tr.Search(marcelitte))
    assert.True(t, tr.Search(mar))

    assert.True(t, tr.Delete(mar))
    assert.False(t, tr.Search(mar))
    assert.True(t, tr.Search(marcelitte))

    assert.False(t, tr.Delete(tuesday))
    assert.False(t, tr.Delete(marl))
}
