import (
	"slices"
	"maps"
)
func groupAnagrams(strs []string) [][]string {
	ana := make(map[[26]int][]string)
	for _, char := range strs {
		count := [26]int{}
		for _, a := range char {
			count[a - 'a']++
		}
		ana[count] = append(ana[count], char)
	}
	return slices.Collect(maps.Values(ana))
}
