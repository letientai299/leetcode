package main

func minOperations1598(logs []string) int {
	dep := 0
	for _, s := range logs {
		switch s {
		case "../":
			if dep > 0 {
				dep--
			}
		case "./":
			continue
		default:
			dep++
		}
	}
	return dep
}
