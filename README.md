# HackEye-Hackathon


## Pre-requisites :
* A PostgreSQL database running on preferably docker.
* DB containing a table named `groceries` .

## To build and run server :
 
* Enter the server's directory : ` cd Server`
* Build and run the executable : `./build.sh`

## To build server on docker :

* Run `docker build -t wmgroceries .` 
* Run `docker run -d -p 8008:8008 wmgroceries:latest`

## Scripts

### Cleaning Script
- Download [dataset](https://www.kaggle.com/datafiniti/food-ingredient-lists/downloads/ingredients%20v1.csv/1)
- `python3 <path to csv file> <directory to store resultant files>`
-  example <br> 
  `python3 data/ingredients.csv data/`
- <b>NEED pandas library to run</b>
