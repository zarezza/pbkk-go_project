package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Recipe struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
}

func loadRecipes(filename string) ([]Recipe, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var recipes []Recipe
	err = json.Unmarshal(data, &recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func getUserIngredients() []string {
	fmt.Print("Enter ingredients (comma-separated): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return strings.Split(input, ",")
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if strings.TrimSpace(a) == strings.TrimSpace(item) {
			return true
		}
	}
	return false
}

func findRecipes(recipes []Recipe, ingredients []string) []Recipe {
	var matches []Recipe

	for _, recipe := range recipes {
		match := true
		for _, reqIngredient := range recipe.Ingredients {
			if !contains(ingredients, reqIngredient) {
				match = false
				break
			}
		}

		if match {
			matches = append(matches, recipe)
		}
	}
	return matches
}

func main() {
	recipes, err := loadRecipes("recipes.json")
	if err != nil {
		log.Fatal("Error loading recipes:", err)
	}

	fmt.Println("Recipes loaded successfully!")

	userIngredients := getUserIngredients()

	matchingRecipes := findRecipes(recipes, userIngredients)
	fmt.Println("\nRecipes you can make:")
	if len(matchingRecipes) == 0 {
		fmt.Println("No recipes found with the given ingredients.")
	} else {
		for _, recipe := range matchingRecipes {
			fmt.Println("-", recipe.Name)
		}
	}
}
