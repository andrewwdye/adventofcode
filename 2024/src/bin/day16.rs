use clap::Parser;
use core::cmp::Reverse;
use std::{
    collections::{BinaryHeap, HashMap, HashSet},
    fmt::Debug,
    fs::read_to_string,
};

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
    let (cost, _) = m.solve();
    Ok(cost)
}

fn solve2(input: &str) -> Result<i32, std::io::Error> {
    let m = Maze::new(input);
    let (_, best_path_tiles) = m.solve();
    Ok(best_path_tiles as i32)
}

type Location = (i32, i32);

#[derive(Debug, Clone, Copy, Eq, PartialEq, Hash)]
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

#[derive(Eq, PartialEq)]
struct Entry {
    location: Location,
    direction: Direction,
    cost: i32,
    path: HashSet<Location>,
}

impl Debug for Entry {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("Entry")
            .field("location", &self.location)
            .field("direction", &self.direction)
            .field("cost", &self.cost)
            .field("path", &self.path.len())
            .finish()
    }
}

impl Ord for Entry {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        self.cost.cmp(&other.cost)
    }
}

impl PartialOrd for Entry {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        self.cost.partial_cmp(&other.cost)
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

    fn solve(&self) -> (i32, usize) {
        let mut visited = HashMap::new();
        let mut h: BinaryHeap<Reverse<Entry>> = BinaryHeap::new();
        let mut p = HashSet::new();
        p.insert(self.start);
        h.push(Reverse(Entry {
            location: self.start,
            direction: Direction::East,
            cost: 0,
            path: p,
        }));
        let mut min_cost = i32::MAX;
        let mut best_path_tiles = HashSet::new();
        loop {
            let entry = h.pop().unwrap().0;
            if entry.cost > min_cost {
                return (min_cost, best_path_tiles.len());
            }
            let element = self.grid[entry.location.1 as usize][entry.location.0 as usize];
            if element == 'E' {
                min_cost = entry.cost;
                for loc in entry.path.iter() {
                    best_path_tiles.insert(loc.clone());
                }
                continue;
            }
            if element == '#' {
                continue;
            }
            if let Some(&prev_cost) = visited.get(&(entry.location, entry.direction)) {
                if entry.cost > prev_cost {
                    continue;
                }
            }
            visited.insert((entry.location, entry.direction), entry.cost);
            for (dir, turn_cost) in [
                (entry.direction, 0),
                (entry.direction.left(), 1000),
                (entry.direction.right(), 1000),
            ] {
                let (dx, dy) = dir.steps();
                let next = (entry.location.0 + dx, entry.location.1 + dy);
                if entry.path.contains(&next) {
                    continue;
                }
                let mut path = entry.path.clone();
                path.insert(next);
                let e = Entry {
                    location: next,
                    direction: dir,
                    cost: entry.cost + turn_cost + 1,
                    path: path,
                };
                h.push(Reverse(e));
            }
        }
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

    #[test]
    fn test_part2() {
        assert_eq!(solve2(SAMPLE_INPUT1).unwrap(), 45);
        assert_eq!(solve2(SAMPLE_INPUT2).unwrap(), 64);
    }
}
