use clap::Parser;
use std::{collections::HashSet, fs::read_to_string};

#[derive(Parser)]
struct Cli {
    #[arg(short, long, value_parser = clap::value_parser!(u8).range(1..=2), default_value_t = 1)]
    part: u8,

    input: String
}

fn main() -> Result<(), std::io::Error>{
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

fn solve1(input: &str) -> Result<i32, std::io::Error>{
    let room: Room = input.lines()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();
    let mut guard = find_guard(&room);
    let mut visited: HashSet<Location> = HashSet::new();
    while guard.in_room(&room) {
        visited.insert(guard.location);
        guard.step(&room);
    }
    Ok(visited.len() as i32)
}

fn solve2(input: &str) -> Result<i32, std::io::Error>{
    let room: Room = input.lines()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();
    let guard = find_guard(&room);
    let mut stuck = 0;
    for y in 0..room.len() {
        for x in 0..room[y].len() {
            if room[y][x] != '.' {
                continue;
            }
            // TODO: could also prerun and check for places to put an obstruction based on what we visited
            let mut guard = guard.clone();
            let mut room = room.clone();
            room[y][x] = '#';
            let mut visited: HashSet<Guard> = HashSet::new();
            while guard.in_room(&room) {
                if visited.contains(&guard) {
                    stuck += 1;
                    break;
                }
                visited.insert(guard);
                guard.step(&room);
            }
        }
    }
    Ok(stuck)
}

type Room = Vec<Vec<char>>;

#[derive(Debug, PartialEq, Eq, Hash, Clone, Copy)]
enum Direction {
    Up,
    Down,
    Left,
    Right
}

#[derive(Debug, PartialEq, Eq, Hash, Clone, Copy)]
struct Location {
    x: i32,
    y: i32
}

impl Location {
    fn in_room(&self, room: &Room) -> bool {
        self.y >= 0 && self.y < room.len() as i32 && 
        self.x >= 0 && self.x < room[self.y as usize].len() as i32
    }

}

#[derive(Debug, PartialEq, Eq, Hash, Clone, Copy)]
struct Guard {
    location: Location,
    direction: Direction
}

impl Guard {
    fn in_room(&self, room: &Room) -> bool {
        self.location.in_room(room)
    }

    fn step(&mut self, room: &Room) {
        let new_location = match self.direction {
            Direction::Up => Location{x: self.location.x, y: self.location.y - 1},
            Direction::Down => Location{x: self.location.x, y: self.location.y + 1},
            Direction::Left => Location{x: self.location.x - 1, y: self.location.y},
            Direction::Right => Location{x: self.location.x + 1, y: self.location.y}
        };
        if new_location.in_room(room) && room[new_location.y as usize][new_location.x as usize] == '#' {
            self.direction = match self.direction {
                Direction::Up => Direction::Right,
                Direction::Down => Direction::Left,
                Direction::Left => Direction::Up,
                Direction::Right => Direction::Down
            };
        } else {
            self.location = new_location;
        }
    }
}

fn find_guard(room: &Room) -> Guard {
    for y in 0..room.len() {
        for x in 0..room[y].len() {
            match room[y][x] {
                '^' => return Guard { location: Location{x: x as i32, y: y as i32}, direction: Direction::Up },
                'v' => return Guard { location: Location{x: x as i32, y: y as i32}, direction: Direction::Down },
                '<' => return Guard { location: Location{x: x as i32, y: y as i32}, direction: Direction::Left },
                '>' => return Guard { location: Location{x: x as i32, y: y as i32}, direction: Direction::Right },
                _ => {}
            }
        }
    }
    unreachable!()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    const SAMPLE_INPUT: &str = "....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT).unwrap(), 41);
    }

    #[test]
    fn test_part2() {
        assert_eq!(solve2(SAMPLE_INPUT).unwrap(), 6);
    }
}
