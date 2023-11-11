# Pokedexcli

This is a for fun Pokemon project that allows its user to explore the vast regions in all of pokemon using the Pokeapi! You can explore, catch, view the pokemon and store them within your pokedex. Gotta catch em all!

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)


## Installation
 

```bash
  git clone https://github.com/nmowens95/pokedexcli
```
    
## Deployment

To use this project, from the command line run: 

```bash
  go build && ./pokedexcli
```
## Lessons Learned


- How to make use an api and use it for different endpoint
- Making use of a go routine, mutexes and pointers
- How to make a unit text in go
- Managed different packages and how they interact specific to go
- Working with a Repl style project from the command line
## Roadmap
- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Refactor your code to organize it better and make it more testable
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon