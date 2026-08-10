package hero

type Hero struct {
	name   string // lowercase ie not exported
	health int    // lowercase ie not exported
}

// Setters
// we pass pointers to modify the original instance
func (h *Hero) Set_name(name string) {
	h.name = name
}

func (h *Hero) Set_health(health int) {
	h.health = health
}

// Setters
func (h Hero) Get_name() string {
	return h.name
}

func (h Hero) Get_health() int {
	return h.health
}
