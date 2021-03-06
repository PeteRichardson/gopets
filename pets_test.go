package pets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPet(t *testing.T) {
	p := NewPet("Bella", "Cat", "American Shorthair", 5)
	assert.Equal(t, "Cat", p.Species)
	assert.Equal(t, "Bella", p.Name)
	assert.Equal(t, "American Shorthair", p.Breed)
	assert.Equal(t, 5, p.Age)
}

func TestRandomPet(t *testing.T) {
	p1 := RandomPet()
	p2 := RandomPet()
	// log.Println(p1)
	// log.Println(p2)
	// Assert that the pets are not identical
	assert.NotEqual(t, p1, p2, "Two random pets seem to be identical.")
}
