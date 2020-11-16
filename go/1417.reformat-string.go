package main

func reformat(s string) string {
	buf := make([]byte, 0, len(s))
	var cs []byte
	var ds []byte
	for _, x := range s {
		if x >= 'a' && x <= 'z' {
			cs = append(cs, byte(x))
		} else {
			ds = append(ds, byte(x))
		}
	}
	if len(ds) > len(cs) {
		cs, ds = ds, cs
	}

	if len(cs)-len(ds) >= 2 {
		return ""
	}

	buf = []byte(s[:0])
	for i := range cs {
		buf = append(buf, cs[i])
		buf = append(buf, ds[i])
	}

	if len(cs) > len(ds) {
		buf = append(buf, cs[len(cs)-1])
	}

	return string(buf)
}
