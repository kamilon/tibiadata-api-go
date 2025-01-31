package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverviewAll(t *testing.T) {
	data, err := ioutil.ReadFile("../testdata/spells/overviewall.html")
	if err != nil {
		t.Errorf("File reading error: %s", err)
		return
	}

	spellsOverviewJson := TibiaSpellsOverviewV3Impl("", string(data))
	assert := assert.New(t)

	assert.Equal(142, len(spellsOverviewJson.Spells.Spells))

	firstSpell := spellsOverviewJson.Spells.Spells[0]
	assert.Equal("Animate Dead Rune (adana mort)", firstSpell.Name)
	assert.Equal("Animate Dead Rune", firstSpell.Spell)
	assert.Equal("adana mort", firstSpell.Formula)
	assert.Equal(27, firstSpell.Level)
	assert.Equal(600, firstSpell.Mana)
	assert.Equal(1200, firstSpell.Price)
	assert.False(firstSpell.GroupAttack)
	assert.False(firstSpell.GroupHealing)
	assert.True(firstSpell.GroupSupport)
	assert.False(firstSpell.TypeInstant)
	assert.True(firstSpell.TypeRune)
	assert.True(firstSpell.PremiumOnly)

	findPersonSpell := spellsOverviewJson.Spells.Spells[53]
	assert.Equal("Find Person (exiva \"name\")", findPersonSpell.Name)
	assert.Equal("Find Person", findPersonSpell.Spell)
	assert.Equal("exiva \"name\"", findPersonSpell.Formula)
	assert.Equal(8, findPersonSpell.Level)
	assert.Equal(20, findPersonSpell.Mana)
	assert.Equal(80, findPersonSpell.Price)
	assert.False(findPersonSpell.GroupAttack)
	assert.False(findPersonSpell.GroupHealing)
	assert.True(findPersonSpell.GroupSupport)
	assert.True(findPersonSpell.TypeInstant)
	assert.False(findPersonSpell.TypeRune)
	assert.False(findPersonSpell.PremiumOnly)
}

func TestOverviewDruid(t *testing.T) {
	data, err := ioutil.ReadFile("../testdata/spells/overviewdruid.html")
	if err != nil {
		t.Errorf("File reading error: %s", err)
		return
	}

	spellsOverviewJson := TibiaSpellsOverviewV3Impl("druid", string(data))
	assert := assert.New(t)

	assert.Equal("druid", spellsOverviewJson.Spells.SpellsVocationFilter)
	assert.Equal(73, len(spellsOverviewJson.Spells.Spells))
}
