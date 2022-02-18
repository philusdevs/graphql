package models 

type People struct {
  getAllPeople {
    Name string: `json:name`
    Height int: `json:height`
    Mass string: `json:mass`
    Gender string: `json:gender`
    Homeworld string: `json:homeworld`
  }
}


type Search struct ($personName: String!) {
    searchPerson(personName: $personName) {
		Name string: `json:name`
		Height int: `json:height`
		Mass string: `json:mass`
		Gender string: `json:gender`
		Homeworld string: `json:homeworld`
    }
  }



  type NextPage struct ($pageNumber: String!) {
    next(pageNumber: $pageNumber) {
		Name string: `json:name`
		Height int: `json:height`
		Mass string: `json:mass`
		Gender string: `json:gender`
		Homeworld string: `json:homeworld`
    }
  }