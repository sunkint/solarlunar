package solarlunar

import (
	"fmt"
	"testing"
)

func TestSolarToChineseLuanr(t *testing.T) {
	solarDate := "1990-05-06"
	fmt.Println(SolarToChineseLunar(solarDate))
}

func TestSolarToSimpleLunar(t *testing.T) {
	solarDate := "1990-05-06"
	fmt.Println(SolarToSimpleLunar(solarDate))
}

func TestLunarToSolar(t *testing.T) {
	lunarDate := "1990-04-12"
	fmt.Println(LunarToSolar(lunarDate, false))
}

func TestSolarToSplitLunar(t *testing.T) {
	lunarDate := "1996-09-06"
	fmt.Println(SolarToSplitLunar(lunarDate))
}