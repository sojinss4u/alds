package main

import (
	"bytes"
	"testing"
)

func TestPrintDfs(t *testing.T) {
	g := Graph{}
	v_list := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for _, item := range v_list {
		g.AddVertex(item)
	}
	e_list := map[string]int{"AB": 1, "AE": 1, "BF": 1, "FI": 1, "FG": 1, "GC": 1, "GF": 1, "GJ": 1, "IH": 1, "IF": 1, "HE": 1, "HD": 1, "DE": 1, "DH": 1, "ED": 1, "EH": 1, "EA": 1}
	for k, v := range e_list {
		g.AddEdge(string(k[0]), string(k[1]), v)
	}
	t.Run("TestPrintDfsSuccess", func(t *testing.T) {
		expect := "ABFGCJIHDE"
		b := bytes.Buffer{}
		g.dfs(g.vertices["A"], 1)
		if g.PrintDfs(&b); expect != b.String() {
			t.Errorf("Expect %s, Got %s", expect, b.String())
		}
	})
}

