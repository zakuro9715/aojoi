package main

var mapping = [][]string{
	{}, // Level 0
	{"0510", "0521", "0532", "0543", "0554", "0565", "0576"},
	{"0511", "0512", "0522", "0533", "0577"},
	{"0513", "0523", "0544", "0555", "0566"},
	{"0516", "0545", "0556", "0578"},
	{"0517", "0515", "0528", "0524", "0526", "0538", "0534", "0535", "0549", "0546", "0557", "0558", "0571", "0567", "0568", "0579"},
	{"0518", "0514", "0527", "0525", "0529", "0539", "0551", "0547", "0560", "0572", "0569", "0573", "0519"},
	{"0530", "0537", "0562", "0580"},
	{"0520", "0531", "0540", "0541", "0542", "0536", "0550", "0548", "0563", "0574", "0570"},
	{"0564", "0559", "0581"},
	{"0552", "0575"},
	{"0553"},
	{}, // 12
}

func getLevel(problemId string) int {
	for level, v := range mapping {
		for _, id := range v {
			if id == problemId {
				return level
			}
		}
	}
	return -1
}
