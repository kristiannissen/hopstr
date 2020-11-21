package api

type Hop struct {
  Name string
}

func LoadHops() (map[string]string, int) {
    m := make(map[string]string)
    m["1"] = "Citra"
    m["2"] = "Amarillo"

    return m, 200
}
