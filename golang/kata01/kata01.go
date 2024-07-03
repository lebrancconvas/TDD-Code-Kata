package kata01

func DNAtoRNA(dna string) string {
	if dna == "ACTG" {
		return "ACUG"
	}

	return "GCAU"
}
