package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(">>%s>> ", strings.Replace(dir, home, "~", 1))
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		comands := strings.Split(scanner.Text(), "|")
		for _, value := range comands {
			args := strings.Fields(value)
			switch args[0] {
			case "cd": //смена директории
				if len(args) != 2 {
					fmt.Println("не корректно указанна директория")
					continue
				}

				switch {
				case strings.HasPrefix(args[1], "."): //Из текущего каталога
					dir = ParseArg(strings.TrimPrefix(args[1], "."), dir)
				case strings.HasPrefix(args[1], "/"): //Из домашнего каталога
					dir = ParseArg(strings.TrimPrefix(args[1], "/"), home)
				default:
					continue
				}

			case "pwd":
				fmt.Println(dir)
			case "echo":
				fmt.Println(strings.TrimSpace(strings.TrimPrefix(value, "echo")))
			case "kill":
				pid, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Println("не корректно указан PID")
					continue
				}
				proc, err := os.FindProcess(pid)
				if err != nil {
					fmt.Println(err)
				}

				err = proc.Kill()
				if err != nil {
					fmt.Println(err)
				}
			case "ps":
				dirNow, err := os.Open("/proc")
				if err != nil {
					fmt.Println(err)
				}

				dirs, err := dirNow.ReadDir(0)
				if err != nil {
					fmt.Println(err)
				}

				result := make([][]string, 0, 0)
				lenMax := make([]int, 3, 3)
				result = append(result, []string{"NAME", "STATE", "PID"})
				for _, val := range dirs {
					if val.IsDir() {
						_, err := strconv.Atoi(val.Name())
						if err == nil {
							fileStats, err := os.Open(fmt.Sprintf("/proc/%s/status", val.Name()))
							if err != nil {
								fmt.Println(err)
							}

							scanner := bufio.NewScanner(fileStats)
							line := make([]string, 3, 3)
							for scanner.Scan() {
								switch {
								case strings.HasPrefix(scanner.Text(), "Name:"):
									line[0] = strings.TrimPrefix(scanner.Text(), "Name:\t")
									if len(line[0]) > lenMax[0] {
										lenMax[0] = len(line[0])
									}
								case strings.HasPrefix(scanner.Text(), "State:"):
									line[1] = strings.TrimPrefix(scanner.Text(), "State:\t")
									if len(line[1]) > lenMax[1] {
										lenMax[1] = len(line[1])
									}
								case strings.HasPrefix(scanner.Text(), "Pid:"):
									line[2] = strings.TrimPrefix(scanner.Text(), "Pid:\t")
									if len(line[2]) > lenMax[2] {
										lenMax[2] = len(line[2])
									}
								default:
									continue
								}
							}

							if !strings.Contains(line[1], "sleep") {
								continue
							}

							result = append(result, line)
						}
					}
				}

				for _, line := range result {
					space := make([]string, 3, 3)
					for idx := range space {
						for i := 0; i < lenMax[idx]-len(line[idx]); i++ {
							space[idx] += " "
						}
					}

					fmt.Printf("%s%s %s%s %s%s\n", line[2], space[2], line[1], space[1], line[0], space[0])

				}
			case "\\quit":
				return
			default:
				fmt.Printf("Unknown command: %s\n", args[0])
			}
		}
		fmt.Printf(">>%s>> ", strings.Replace(dir, home, "~", 1))
	}
}

func ParseArg(s, dir string) string {
	dirNew := dir
	nameDirs := strings.Split(s, "/")
	for _, val := range nameDirs {
		switch {
		case strings.HasPrefix(val, "."):
			for _, str := range strings.Split(val, "") {
				if str != "." || dirNew == "/" {
					return dir
				}

				dirNew = path.Dir(dirNew)

			}

		case val == "":
			continue
		default:
			dirNow, err := os.Open(dirNew)
			if err != nil {
				return dir
			}
			defer dirNow.Close()

			dirs, err := dirNow.ReadDir(0)
			if err != nil {
				return dir
			}

			var ok bool
			for _, v := range dirs {
				if v.IsDir() && v.Name() == val {
					dirNew = path.Join(dirNew, v.Name())
					ok = true
					break
				}
			}

			if !ok {
				return dir
			}
		}
	}
	return dirNew
}
