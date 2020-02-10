package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/Jerome1337/sac/shapes"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

var (
	shape, calculator string
	sizes             []float64
)

func prompt(ctx *cli.Context) error {
	if err := shapePrompt(); err != nil {
		return err
	}

	if err := calcPrompt(); err != nil {
		return err
	}

	if err := sizesPrompt("Size (a, b, c, ...)"); err != nil {
		return err
	}

	shapes.DynamicFuncCall(shape, calculator, sizes)

	return nil
}

func shapePrompt() error {
	prompt := promptui.Select{
		Label: "Select a shape",
		Items: []string{
			"triangle",
			"rectangle",
		},
	}

	_, s, err := prompt.Run()
	if err != nil {
		return err
	}

	shape = s

	return nil
}

func calcPrompt() error {
	prompt := promptui.Select{
		Label: "Select a calculator",
		Items: []string{"area"},
	}

	_, c, err := prompt.Run()
	if err != nil {
		return err
	}

	calculator = c

	return nil
}

func sizesPrompt(label string) error {
	sizesPrompt := promptui.Prompt{
		Label:    label,
		Validate: validateSizesInput,
	}

	givenSizes, err := sizesPrompt.Run()
	if err != nil {
		return err
	}

	formattedSizes := strings.Split(
		strings.Replace(givenSizes, " ", "", -1),
		",",
	)

	for _, size := range formattedSizes {
		num, err := strconv.ParseFloat(size, 64)
		if err != nil {
			return err
		}

		sizes = append(sizes, num)
	}

	return nil
}

func validateSizesInput(input string) error {
	regex := regexp.MustCompile(`(\d+,\s)+\d+$`)

	if ok := regex.MatchString(input); !ok {
		return errors.New("sizes are not properly formatted")
	}

	return nil
}
