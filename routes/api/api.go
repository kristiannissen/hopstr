package api

type Hop struct {
  Name string
}

func LoadHops() (map, int) {
  hops := make(map[string]string)
  hops["citra"] = "Citra"

  return hops, 200
}
