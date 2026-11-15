import (
	"slices"
	"maps"
)
func groupAnagrams(strs []string) [][]string {
	ana := make(map[[26]int][]string)
	for _, word := range strs {
		count := [26]int{}
		for _, char := range word {
			count[char - 'a']++
		}
		ana[count] = append(ana[count], word)
	}
	return slices.Collect(maps.Values(ana))
}
