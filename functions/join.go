package functions

// Join returns a string from joining each of the elements.
func (ss StringSliceType) Join(glue string) (s string) {
	for i, element := range ss {
		if i > 0 {
			s += glue
		}

		s += string(element)
	}

	return s
}
