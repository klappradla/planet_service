package data

func Planets() []byte {
  var jsonData = []byte(`
    [
      {
        "name": "Tatooine",
        "rotation_period": 23,
        "orbital_period": 304,
        "diameter": 10465,
        "climate": "arid",
        "gravity": "1 standard",
        "terrain": "desert",
        "surface_water": 1,
        "population": 200000
      },
      {
        "name": "Alderaan",
        "rotation_period": 24,
        "orbital_period": 364,
        "diameter": 12500,
        "climate": "temperate",
        "gravity": "1 standard",
        "terrain": "grasslands, mountains",
        "surface_water": 40,
        "population": 2000000000
      },
      {
        "name": "Yavin IV",
        "rotation_period": 24,
        "orbital_period": 4818,
        "diameter": 10200,
        "climate": "temperate, tropical",
        "gravity": "1 standard",
        "terrain": "jungle, rainforests",
        "surface_water": 8,
        "population": 1000
      },
      {
        "name": "Hoth",
        "rotation_period": 23,
        "orbital_period": 549,
        "diameter": 7200,
        "climate": "frozen",
        "gravity": "1.1 standard",
        "terrain": "tundra, ice caves, mountain ranges",
        "surface_water": 100,
        "population": null
      },
      {
        "name": "Dagobah",
        "rotation_period": 23,
        "orbital_period": 341,
        "diameter": 8900,
        "climate": "murky",
        "gravity": "N/A",
        "terrain": "swamp, jungles",
        "surface_water": 8,
        "population": 0
      },
      {
        "name": "Bespin",
        "rotation_period": 12,
        "orbital_period": 5110,
        "diameter": 118000,
        "climate": "temperate",
        "gravity": "1.5 (surface), 1 standard (Cloud City)",
        "terrain": "gas giant",
        "surface_water": 0,
        "population": 6000000
      },
      {
        "name": "Endor",
        "rotation_period": 18,
        "orbital_period": 402,
        "diameter": 4900,
        "climate": "temperate",
        "gravity": "0.85 standard",
        "terrain": "forests, mountains, lakes",
        "surface_water": 8,
        "population": 30000000
      },
      {
        "name": "Naboo",
        "rotation_period": 26,
        "orbital_period": 312,
        "diameter": 12120,
        "climate": "temperate",
        "gravity": "1 standard",
        "terrain": "grassy hills, swamps, forests, mountains",
        "surface_water": 12,
        "population": 4500000000
      },
      {
        "name": "Coruscant",
        "rotation_period": 24,
        "orbital_period": 368,
        "diameter": 12240,
        "climate": "temperate",
        "gravity": "1 standard",
        "terrain": "cityscape, mountains",
        "surface_water": null,
        "population": 1000000000000
      },
      {
        "name": "Kamino",
        "rotation_period": 27,
        "orbital_period": 463,
        "diameter": 19720,
        "climate": "temperate",
        "gravity": "1 standard",
        "terrain": "ocean",
        "surface_water": 100,
        "population": 1000000000
      },
      {
        "name": "Geonosis",
        "rotation_period": 30,
        "orbital_period": 256,
        "diameter": 11370,
        "climate": "temperate, arid",
        "gravity": "0.9 standard",
        "terrain": "rock, desert, mountain, barren",
        "surface_water": 5,
        "population": 100000000000
      }
    ]
  `)
  return jsonData
}
