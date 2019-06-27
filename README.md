# HackEye-Hackathon


## To run server :
 
* `cd Server`
* `go build -o main -v ./cmd/... `
* `./main` 

## Scripts

### Cleaning Script
- Download [dataset](https://www.kaggle.com/datafiniti/food-ingredient-lists/downloads/ingredients%20v1.csv/1)
- `python3 <path to csv file> <directory to store resultant files>`
-  example <br> 
  `python3 data/ingredients.csv data/`
- <b>NEED pandas library to run</b>