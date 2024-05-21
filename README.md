[![Build](https://github.com/skribisto/goker/actions/workflows/go.yml/badge.svg)](https://github.com/skribisto/goker/actions/workflows/go.yml)
[![License](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0.fr.html)

# Goker - A Poker Game in Go

## Project Overview

**Goker** is a poker game implemented in Go. This project was created so I can learn the Go programming language. Goker is a fully playable poker game featuring rudimentary CPU intelligence and multiplayer capabilities. The game runs in a Docker container.

## Features

- **Playable Poker Game**: Enjoy a Texas Hold'em game of poker with friends or against CPU opponents. (hard-coded 1 player vs 2 CPUs for now)
- **Rudimentary CPU Intelligence**: Test your skills against basic AI opponents.
- **Multiplayer Support**: Play with multiple human players in the same game with minor tweak.
- **Dockerized**: Run the game effortlessly using Docker.
- **CI/CD**: Very basic CI/CD that automatically builds a docker image and pushes it on Docker Hub for each release.

## Installation Instructions

To start playing Goker, you need Docker installed on your machine. Once Docker is set up, you can run the game with the following command:

```sh
docker run -it skribisto/goker:latest
```

This command pulls the latest Goker image from Docker Hub and starts the game in an interactive terminal session.

## Usage

Goker is primarily intended for my own learning experience on Go language. It showcases (very) basic game logic, CPU decision-making, and multiplayer interaction within a Dockerized environment.

### Future Improvements

Planned future improvements include:

- **Web Application**: Adding a web interface for an enhanced user experience.
- **Performance Enhancements**: Optimizing the game for better performance.
- **Code Refactoring**: Refactoring the codebase to comply with Go standards and best practices.

No commitments are made to code those improvements under any timeline.

## License

Goker is released under the GNU General Public License v3.0. You are free to use, modify, and distribute this software under the terms of the GPLv3. For more details, refer to the [LICENSE](LICENSE) file.

## Contributing

If you have ideas for new features, improvements, or bug fixes, please feel free to open an issue.
Comments on my code, style, choices are very welcome.

## Contact

For any inquiries, suggestions, or feedback, you can reach out via GitHub issues or directly contact me.

---
*This project is a work in progress*

