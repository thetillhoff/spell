package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {

	// #####
	// START flag processing
	// #####

	NATOmode := false
	GERMANmode := false
	mode := "NATO"

	flag.BoolVar(&NATOmode, "nato", false, "[optional] Enable NATO spelling mode")
	flag.BoolVar(&GERMANmode, "german", false, "[optional] Enable GERMAN spelling mode")
	flag.Parse()

	if NATOmode && !GERMANmode {
		mode = "NATO"
	} else if GERMANmode && !NATOmode {
		mode = "GERMAN"
	} else {
		log.Fatal("Bad use of flags.")
	}

	// #####
	// END flag processing
	// START args processing
	// #####

	argsWithoutProg := strings.Join(flag.Args(), " ")

	// #####
	// END args processing
	// START spelling
	// #####

	if len(argsWithoutProg) == 0 {
		help()
	} else {
		for _, letter := range argsWithoutProg {
			if letter == ' ' {
				fmt.Printf("\n")
			} else {
				spellLetter(letter, mode)
			}
		}
	}

	// #####
	// END spelling
	// #####
}

// takes: letter, mode
// returns: nothing
// prints: spelled letter according to mode
func spellLetter(letter rune, mode string) {

	NATO := make(map[rune]string)
	NATO['a'] = "alfa"
	NATO['b'] = "bravo"
	NATO['c'] = "charlie"
	NATO['d'] = "delta"
	NATO['e'] = "echo"
	NATO['f'] = "foxtrott"
	NATO['g'] = "golf"
	NATO['h'] = "hotel"
	NATO['i'] = "india"
	NATO['j'] = "juliett"
	NATO['k'] = "kilo"
	NATO['l'] = "lima"
	NATO['m'] = "mike"
	NATO['n'] = "november"
	NATO['o'] = "oscar"
	NATO['p'] = "papa"
	NATO['q'] = "quebec"
	NATO['r'] = "romeo"
	NATO['s'] = "sierra"
	NATO['t'] = "tango"
	NATO['u'] = "uniform"
	NATO['v'] = "victor"
	NATO['w'] = "whiskey"
	NATO['x'] = "x-ray"
	NATO['y'] = "yankee"
	NATO['z'] = "zulu"

	GERMAN := make(map[rune]string)
	GERMAN['a'] = "anton"
	GERMAN['b'] = "berta"
	GERMAN['c'] = "cäsar"
	GERMAN['d'] = "dora"
	GERMAN['e'] = "emil"
	GERMAN['f'] = "friedrich"
	GERMAN['g'] = "gustav"
	GERMAN['h'] = "heinrich"
	GERMAN['i'] = "ida"
	GERMAN['j'] = "julius"
	GERMAN['k'] = "kaufmann"
	GERMAN['l'] = "ludwig"
	GERMAN['m'] = "martha"
	GERMAN['n'] = "nordpol"
	GERMAN['o'] = "otto"
	GERMAN['p'] = "paula"
	GERMAN['q'] = "quelle"
	GERMAN['r'] = "richard"
	GERMAN['s'] = "siegfried"
	GERMAN['t'] = "theodor"
	GERMAN['u'] = "ulrich"
	GERMAN['v'] = "viktor"
	GERMAN['w'] = "wilhelm"
	GERMAN['x'] = "xanthippe"
	GERMAN['y'] = "ypsilon"
	GERMAN['z'] = "zeppelin"

	switch mode {
	case "NATO":
		fmt.Printf("%s\n", NATO[letter])
	case "GERMAN":
		fmt.Printf("%s\n", GERMAN[letter])
	default:
		log.Fatal("Received unkown mode. (Blame the dev.)")
	}
}

// takes: nothing
// returns: nothing
// prints: help
func help() {
	fmt.Printf("  usage:\n  $> spell [-german|-nato<default>] <what to spell>")
}
