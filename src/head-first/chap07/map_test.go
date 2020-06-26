package chap07

import "testing"

func TestInit(t *testing.T) {
	var rankMap map[string]int
	rankMap = make(map[string]int)

	rankMap["gold"] = 1
	rankMap["silver"] = 2
	rankMap["bronze"] = 3

	t.Log(rankMap)
	t.Logf("%#v", rankMap)
}

func TestLiterals(t *testing.T) {
	literalMap := map[string]float64{"a": 1.2, "b": 5.6}
	t.Logf("%#v", literalMap)

	t.Log(literalMap["a"])
}

func TestMapCount(t *testing.T) {
	strings := []string{"a", "a", "b", "b", "c", ""}
	countMap := make(map[string]int)
	for _, str := range strings {
		countMap[str]++
	}
	t.Log(countMap)

	for key, value := range countMap {
		t.Log(key, value)
	}
}
