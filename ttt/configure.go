package ttt

import (
	"bufio"
	"strings"
)

func Configure(reader *bufio.Reader, options *Options) {
	ClearScreen()
	print(YELLOW + "\n\n\n\t\tSETTINGS\n\n\n" + ORIGINAL)

	var choice string
	var err error
	for {
		print(RED + "\tPlay Computer or Human:\n\n" + ORIGINAL)
		print("\t\t1) computer\n")
		print("\t\t2) human\n")
		print(BLUE + "\n\t\t>_ " + ORIGINAL)
		choice, err = reader.ReadString(byte('\n'))
		if err == nil {
			choice = strings.TrimSpace(choice)
			if choice == "1" {
				options.Computer = true
				options.Starter = options.SymbolTop
				break
			} else if choice == "2" {
				options.Computer = false
				break
			}
		}
	}

	for {
		print(RED + "\n\n\n\tAlter settings:\n\n" + ORIGINAL)
		print("\t\ty) yes\n")
		print("\t\tn) no\n")
		print("\t\tq) quit\n")
		print(BLUE + "\n\t\t>_ " + ORIGINAL)
		choice, err = reader.ReadString(byte('\n'))
		if err == nil {
			choice = strings.TrimSpace(choice)
			if choice == "y" || choice == "yes" || choice == "1" {
				if options.Computer {
					options.SymbolTop = Computer(reader)
					options.SymbolComp = options.SymbolTop
				} else {
					options.SymbolTop = Top(reader)
				}
				options.SymbolBtm = Bottom(reader)

				if options.Computer {
					options.Starter = options.SymbolComp
				} else {
					options.Starter = Starter(reader, options.SymbolTop, options.SymbolBtm)
				}

				// symbols must be different
				if options.SymbolTop == options.SymbolBtm {
					continue
				}
				break
			} else if choice == "n" || choice == "no" || choice == "2" {
				break
			} else if choice == "q" || choice == "quit" || choice == "3" {
				return
			} else {
				continue
			}
		}
	}
}

func Opponent(reader *bufio.Reader) (computer bool) {
	for {
		print("\n\tPlay against:\n")
		print("\t\t1) computer\n")
		print("\t\t2) human\n")
		print(BLUE + "\n\t\t>_ " + ORIGINAL)

		choice, err := reader.ReadString(byte('\n'))
		if err == nil {
			choice = strings.TrimSpace(choice)
			if choice == "1" {
				computer = true
				break
			} else if choice == "2" {
				computer = false
				break
			} else {
				continue
			}
		}
	}
	return
}

func Top(reader *bufio.Reader) (token string) {
	for {
		print("\n\tFirst Symbol: (any printable ascii character)\n")
		print(BLUE + "\n\t\t>_ " + ORIGINAL)
		choice, err := reader.ReadString(byte('\n'))
		if err == nil {
			choice = strings.TrimSpace(choice)
			if len(choice) > 0 {
				xter := choice[0]
				if xter > 32 && xter < 127 {
					token = string(xter)
					return
				} else {
					continue
				}
			} else {
				continue
			}
		}
	}
}

func Bottom(reader *bufio.Reader) (token string) {
	for {
		print("\n\tSecond Symbol: (any printable ascii character)\n")
		print(BLUE + "\n\t\t>_ " + ORIGINAL)
		choice, err := reader.ReadString(byte('\n'))
		if err == nil {
			choice = strings.TrimSpace(choice)
			if len(choice) > 0 {
				xter := choice[0]
				if xter > 32 && xter < 127 {
					token = string(xter)
					return
				} else {
					continue
				}
			} else {
				continue
			}
		}
	}
}

func Computer(reader *bufio.Reader) (token string) {
	for {
		println("\n\tComputer Symbol: (any printable ascii character)\n")
		print(BLUE + "\n\t\t>_ " + ORIGINAL)

		choice, err := reader.ReadString(byte('\n'))
		if err == nil {
			choice = strings.TrimSpace(choice)
			if len(choice) > 0 {
				xter := choice[0]
				if xter > 32 && xter < 127 {
					token = string(xter)
					return
				} else {
					continue
				}
			} else {
				ClearScreen()
				continue
			}
		}
	}
}

func Starter(reader *bufio.Reader, symbolOne, symbolTwo string) (token string) {
	for {
		println("\n\tStarter Symbol: (any of the two aforementioned symbols)\n")
		print(BLUE + "\n\t\t>_ " + ORIGINAL)

		choice, err := reader.ReadString(byte('\n'))
		if err == nil {
			choice = strings.TrimSpace(choice)
			if len(choice) > 0 {
				xter := choice[0]
				if xter > 32 && xter < 127 {
					token = string(xter)
					if token == symbolOne || token == symbolTwo {
						return
					} else {
						continue
					}
				} else {
					continue
				}
			} else {
				continue
			}
		}
	}
}
