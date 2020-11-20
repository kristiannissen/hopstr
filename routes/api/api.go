package api

type Hop struct {
  Name string
}

func LoadHops() (map, int) {
  hops := make(map[string]Hop)
  hops["citra"] = Hop{Name: "Citra"}

  return hops, 200
}
