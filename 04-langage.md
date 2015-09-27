# Le langage

## Chaines de caracteres

~~~ {.go}
    myString := "I'm a string"
    myString = "Hello " + myString
    fmt.Println(myString) // Hello I'm a string
~~~

## Types numeriques

~~~ {.go}
    myInt := 2
    myFloat := 1.2
    otherFloat := float64(myInt) * myFloat
    fmt.Println(otherFloat) // 2.4
~~~

## Boucles

~~~ {.go}
    langages := []string{"Go", "Java", "Ruby", "Python"}

    for i, langage := range langages {
        fmt.Println(langage, "a le rang", i)
    }
~~~

affiche :

~~~ {.sh}
    Go a le rang 0
    Java a le rang 1
    Ruby a le rang 2
    Python a le rang 3
~~~

## Tableaux associatifs

~~~ {.go}
    starWarsYears := map[string]int{
        "A New Hope": 1977,
        "The Empire Strikes Back": 1980,
        "Revenge of the Sith": 2005,
    }
~~~

ajout / suppression :

~~~ {.go}
    starWarsYears["The Force Awakens"] = 2015
    delete(starWarsYears, "Revenge of the Sith")
~~~

boucle :

~~~ {.go}
    for title, year := range starWarsYears {
        fmt.Println(title, "was released in", year)
    }
~~~

## Fonctions

visibilité :

~~~ {.go}
func giveMePear(fruit string) // non visible en dehors du package
func GiveMePear(fruit string) // visible en dehors du package
~~~

paramètres :

~~~ {.go}
func giveMePear(fruit string) {
	fruit = "pear"
}

func main() {
	fruit := "banana"
	giveMePear(fruit)
	fmt.Println(fruit) // banana
}
~~~

## Pointeurs

~~~ {.go}
func giveMePear(fruit *string) {
	*fruit = "pear"
}

func main() {
	fruit := "banana"
	giveMePear(&fruit)
	fmt.Println(fruit) // pear
}
~~~

## Structures

~~~ {.go}
	type Movie struct {
		Actors      []string
		Rating      float32
		ReleaseYear int
		Title       string
	}
	episodeIV := Movie{
		Title:       "Star Wars: A New Hope",
		Rating:      5.0,
		ReleaseYear: 1977,
	}
~~~

type methods :

~~~ {.go}
func (movie Movie) DisplayTitle() string {
    return fmt.Sprintf("%s (%d)", movie.Title, movie.ReleaseYear)
}
~~~

