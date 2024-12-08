use clap::Parser;
use std::{collections::{HashMap, HashSet}, fs::read_to_string};

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

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
struct Location {
    x: i32,
    y: i32
}

struct FrequencyMap {
    height: i32,
    width: i32,
    locations: HashMap<char, Vec<Location>>,
}

impl FrequencyMap {
    fn new(input: &str) -> Self {
        let mut height = 0;
        let mut width = 0;
        let mut locations: HashMap<char, Vec<Location>> = HashMap::new();
        for (y, line) in input.lines().enumerate() {
            height += 1;
            width = line.len() as i32;
            for (x, c) in line.chars().enumerate() {
                if c == '.' {
                    continue;
                }
                locations.entry(c).or_default().push(Location { x: x as i32, y: y as i32 });
            }
        }
        FrequencyMap { locations, height, width }
    }
}

fn solve1(input: &str) -> Result<i32, std::io::Error>{
    let map = FrequencyMap::new(input);
    let mut antinodes: HashSet<Location> = HashSet::new();
    for (_, locations) in map.locations.into_iter() {
        for (i, l1) in locations.iter().enumerate() {
            for l2 in locations[i+1..].iter() {
                let dx = l1.x - l2.x;
                let dy = l1.y - l2.y;
                let a1 = Location{ x: l1.x + dx, y: l1.y + dy };
                let a2 = Location{ x: l2.x - dx, y: l2.y - dy };
                for a in &[a1, a2] {
                    if a.x >= 0 && a.x < map.width && a.y >= 0 && a.y < map.height {
                        antinodes.insert(*a);
                    }
                }
            }
        }
    }
    Ok(antinodes.len() as i32)
}

fn solve2(_: &str) -> Result<i32, std::io::Error>{
    unimplemented!()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    const SAMPLE_INPUT: &str = "............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............";

    #[test]
    fn test_sample_part1() {
        assert_eq!(solve1(SAMPLE_INPUT).unwrap(), 14);
    }
}
