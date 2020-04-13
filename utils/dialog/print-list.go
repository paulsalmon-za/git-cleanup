package dialog

// PrintList ...
func PrintList(collection []string) {
	if len(collection) == 0 {
		Info("No items to display")
	}
	for _, c := range collection {
		Info(c)
	}
}
