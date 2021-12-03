package anagramms_test

import (
	"github.com/asavt7/go-little-tasks/tasks/anagramms"
	"testing"
)

func TestIsAnagrams(t *testing.T) {

	t.Run("is", func(t *testing.T) {
		if !anagramms.IsAnagrams("asdfgh", "adghsf") {
			t.Errorf("its anagram!")
		}
	})
	t.Run("not", func(t *testing.T) {
		if anagramms.IsAnagrams("121asdfgh", "989adghsf") {
			t.Errorf("its not anagram!")
		}
	})

}
