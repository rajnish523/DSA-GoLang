package dsaquestions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	array := []int{0,0,1,1,1,2,2,3,3,3,4}
	expected := 5
	output := RemoveDuplicatesFromSortedArray(array)
	require.Equal(t, expected, output);
}
