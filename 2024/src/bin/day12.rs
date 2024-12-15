use clap::Parser;
use std::{collections::HashSet, fmt::Debug, fs::read_to_string};

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

struct Garden {
    plots: Vec<Vec<char>>
}

impl Garden {
    fn new(input: &str) -> Self {
        let mut plots = Vec::new();
        for line in input.lines() {
            plots.push(line.chars().collect());
        }
        Garden { plots }
    }

    fn height(&self) -> usize {
        self.plots.len()
    }

    fn width(&self) -> usize {
        self.plots[0].len()
    }
}

impl Debug for Garden {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for row in &self.plots {
            for plot in row {
                write!(f, "{}", plot)?;
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

fn solve1(input: &str) -> Result<usize, std::io::Error>{
    let garden = Garden::new(input);
    let mut visited: HashSet<(usize, usize)> = HashSet::new();
    let mut total = 0;
    for y in 0..garden.height() {
        for x in 0..garden.width() {
            let (area, perim) = expand(&garden, &mut visited, x, y, garden.plots[y][x]);
            total += area * perim;
        }
    }
    Ok(total)
}

fn expand(garden: &Garden, visited: &mut HashSet<(usize, usize)>, x: usize, y: usize, t: char) -> (usize, usize) {
    if garden.plots[y][x] != t {
        return (0, 1);
    }
    if visited.contains(&(x, y)) {
        return (0, 0);
    }
    visited.insert((x, y));
    let mut total_area = 1;
    let mut total_perimeter  = 0;
    for (dx, dy) in &[(0, -1), (1, 0), (0, 1), (-1, 0)] {
        let nx = x as i32 + dx;
        let ny = y as i32 + dy;
        if nx >= 0 && nx < garden.width() as i32 && ny >= 0 && ny < garden.height() as i32 {
            let (area, perim) = expand(garden, visited, nx as usize, ny as usize, t);
            total_area += area;
            total_perimeter += perim;
        } else {
            total_perimeter += 1;
        }
    }
    (total_area, total_perimeter)
}

fn solve2(input: &str) -> Result<usize, std::io::Error>{
    let garden = Garden::new(input);
    let mut visited: HashSet<(usize, usize)> = HashSet::new();
    let mut total = 0;
    for y in 0..garden.height() {
        for x in 0..garden.width() {
            let (area, corners) = expand2(&garden, &mut visited, x, y, garden.plots[y][x]);
            let price = area * corners;
            // if area > 0 {
            //     println!("{}: {} * {} = {}", garden.plots[y][x], area, corners, price);
            // }
            total += price;
        }
    }
    Ok(total)
}

/// Returns the area and corners (sides...) of newly discovered plots
fn expand2(garden: &Garden, visited: &mut HashSet<(usize, usize)>, x: usize, y: usize, t: char) -> (usize, usize) {
    if garden.plots[y][x] != t || visited.contains(&(x, y)) {
        return (0, 0);
    }
    visited.insert((x, y));
    let mut total_area = 1;
    let mut total_corners = 0;
    let dirs = [(0, -1), (1, 0), (0, 1), (-1, 0)];
    for i in 0..dirs.len() {
        let dleft = dirs[i];
        let dright = dirs[(i + 1) % dirs.len()];
        let ddiag = (dleft.0 + dright.0, dleft.1 + dright.1);
        let vleft = try_at(garden, dleft.0, dleft.1, x, y);
        let vright = try_at(garden, dright.0, dright.1, x, y);
        let vdiag = try_at(garden, ddiag.0, ddiag.1, x, y);
        if vleft != Some(t) && vright != Some(t) {
            total_corners += 1;
        } else if vleft == Some(t) && vright == Some(t) && vdiag != Some(t) {
            total_corners += 1;
        }
    }
    for (dx, dy) in &dirs {
        let nx = x as i32 + dx;
        let ny = y as i32 + dy;
        if nx >= 0 && nx < garden.width() as i32 && ny >= 0 && ny < garden.height() as i32 {
            let (area, corners) = expand2(garden, visited, nx as usize, ny as usize, t);
            total_area += area;
            total_corners += corners;
        }
    }
    (total_area, total_corners)
}

fn try_at(garden: &Garden, dx: i32, dy: i32, x: usize, y: usize) -> Option<char> {
    let nx = x as i32 + dx;
    let ny = y as i32 + dy;
    if nx >= 0 && nx < garden.width() as i32 && ny >= 0 && ny < garden.height() as i32 {
        Some(garden.plots[ny as usize][nx as usize])
    } else {
        None
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const SAMPLE_INPUT: &str = "AAAA
BBCD
BBCC
EEEC";

    const SAMPLE_INPUT2: &str = "EEEEE
EXXXX
EEEEE
EXXXX
EEEEE";

    const SAMPLE_INPUT3: &str = "AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA";

    const SAMPLE_INPUT4: &str = "RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT).unwrap(), 140);
    }

    #[test]
    fn test_part2() {
        assert_eq!(solve2(SAMPLE_INPUT).unwrap(), 80);
        assert_eq!(solve2(SAMPLE_INPUT2).unwrap(), 236);
        assert_eq!(solve2(SAMPLE_INPUT3).unwrap(), 368);
        assert_eq!(solve2(SAMPLE_INPUT4).unwrap(), 1206);
    }
}
