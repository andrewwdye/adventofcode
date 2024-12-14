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

struct Map {
    values: Vec<Vec<u8>>,
}

impl Map {
    fn new(input: &str) -> Self {
        let mut values: Vec<Vec<u8>> = Vec::new();
        for line in input.lines() {
            values.push(line.as_bytes().to_vec().iter().map(|&c| c - b'0').collect());
        }
        Map { values }
    }

    fn height(&self) -> usize {
        self.values.len()
    }

    fn width(&self) -> usize {
        self.values[0].len()
    }

    fn at(&self, x: usize, y: usize) -> u8 {
        self.values[y][x]
    }
}

impl std::fmt::Debug for Map {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for row in &self.values {
            for v in row {
                write!(f, "{}", v)?;
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

fn solve1(input: &str) -> Result<i32, std::io::Error>{
    let m = Map::new(input);
    let mut total = 0;
    for y in 0..m.height() {
        for x in 0..m.width() {
            let score = explore(&m, x, y, 0).len();
            total += score as i32;
        }
    }
    Ok(total)
}

fn explore(map: &Map, x: usize, y: usize, seek: u8) -> HashSet<(usize, usize)> {
    let elevation = map.at(x, y);
    if elevation != seek {
        return HashSet::new();
    }
    if elevation == 9 {
        let mut h = HashSet::new();
        h.insert((x, y));
        return h;
    }
    let mut all: HashSet<(usize, usize)> = HashSet::new();
    for (dx, dy) in &[(0, -1), (1, 0), (0, 1), (-1, 0)] {
        let nx = x as i32 + dx;
        let ny = y as i32 + dy;
        if nx < 0 || nx >= map.width() as i32 || ny < 0 || ny >= map.height() as i32 {
            continue;
        }
        let nx = nx as usize;
        let ny = ny as usize;
        let found = explore(map, nx, ny, elevation + 1);
        for (fx, fy) in found {
            all.insert((fx, fy));
        }
    }
    all
}

fn solve2(_: &str) -> Result<i32, std::io::Error>{
    unimplemented!()
}

#[cfg(test)]
mod tests {
    use super::*;

    const SAMPLE_INPUT: &str = "89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT).unwrap(), 36);
    }
}
