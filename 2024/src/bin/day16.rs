use clap::Parser;
use std::{collections::HashSet, fs::read_to_string};

#[derive(Parser)]
struct Cli {
    #[arg(short, long, value_parser = clap::value_parser!(u8).range(1..=2), default_value_t = 1)]
    part: u8,

    input: String,
}

fn main() -> Result<(), std::io::Error> {
    let cli = Cli::parse();
    let input = read_to_string(cli.input)?;
    let result = match cli.part {
        1 => solve1(input.as_str())?,
        2 => solve2(input.as_str())?,
        _ => unreachable!(),
    };
    println!("{result}");
    Ok(())
}

fn solve1(input: &str) -> Result<i32, std::io::Error> {
    let m = Maze::new(input);
    Ok(m.solve())
}

fn solve2(_: &str) -> Result<i32, std::io::Error> {
    unimplemented!()
}

type Location = (i32, i32);

#[derive(Debug, Clone, Copy)]
enum Direction {
    North = 0,
    East,
    South,
    West,
}

impl Direction {
    fn left(&self) -> Self {
        match self {
            Self::North => Self::West,
            Self::East => Self::North,
            Self::South => Self::East,
            Self::West => Self::South,
        }
    }

    fn right(&self) -> Self {
        match self {
            Self::North => Self::East,
            Self::East => Self::South,
            Self::South => Self::West,
            Self::West => Self::North,
        }
    }

    fn steps(&self) -> (i32, i32) {
        match self {
            Self::North => (0, -1),
            Self::East => (1, 0),
            Self::South => (0, 1),
            Self::West => (-1, 0),
        }
    }
}

struct Maze {
    grid: Vec<Vec<char>>,
    start: Location,
}

impl Maze {
    fn new(input: &str) -> Self {
        let grid: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();
        let start = (1 as i32, grid.len() as i32 - 2);
        Self { grid, start }
    }

    fn solve(&self) -> i32 {
        return self.recurse(self.start, Direction::East, &mut HashSet::new());
    }

    fn recurse(
        &self,
        location: Location,
        direction: Direction,
        visited: &mut HashSet<Location>,
    ) -> i32 {
        let element = self.grid[location.1 as usize][location.0 as usize];
        if element == 'E' {
            return 0;
        }
        if element == '#' || visited.contains(&location) {
            return i32::MAX;
        }
        visited.insert(location);
        // Loop over each dir and find cheapest path
        let mut min_cost = i32::MAX;
        for (dir, turn_cost) in [
            (direction, 0),
            (direction.left(), 1000),
            (direction.right(), 1000),
        ] {
            let (dx, dy) = dir.steps();
            let next = (location.0 + dx, location.1 + dy);
            let recurse_cost = self.recurse(next, dir, visited);
            if recurse_cost == i32::MAX {
                continue;
            }
            let cost = recurse_cost + 1 + turn_cost;
            if cost < min_cost {
                min_cost = cost;
            }
        }
        visited.remove(&location);
        min_cost
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const SAMPLE_INPUT1: &str = "###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############";

    const SAMPLE_INPUT2: &str = "#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT1).unwrap(), 7036);
        assert_eq!(solve1(SAMPLE_INPUT2).unwrap(), 11048);
    }
}
