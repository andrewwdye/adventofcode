use clap::Parser;
use std::{fs::read_to_string, thread::sleep};

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

struct Guard {
    position: (i32, i32),
    velocity: (i32, i32),
    width: i32,
    height: i32,
}

impl Guard {
    fn new(input: &str, width: i32, height: i32) -> Self {
        let re = regex::Regex::new(r"-?\d+").unwrap();
        let results: Vec<i32> = re.find_iter(input).map(|m| m.as_str().parse::<i32>().unwrap()).collect();
        Guard { position: (results[0], results[1]), velocity: (results[2], results[3]), width, height }
    }

    fn simulate_steps(&self, steps: i32) -> (i32, i32) {
        let mut x = self.position.0 + self.velocity.0 * steps;
        if x < 0 {
            x = self.width - x.abs() % self.width;
        }
        x = x % self.width;
        let mut y = self.position.1 + self.velocity.1 * steps;
        if y < 0 {
            y = self.height - y.abs() % self.height;
        }
        y = y % self.height;
        (x, y)
    }
}

fn solve1(input: &str) -> Result<i32, std::io::Error>{
    solve1_internal(input, 101, 103)
}

fn solve1_internal(input: &str, width: i32, height: i32) -> Result<i32, std::io::Error>{
    let mut guards: Vec<Guard> = Vec::new();
    for line in input.lines() {
        guards.push(Guard::new(line, width, height));
    }
    let mut quads = vec![0; 4];
    let mut grid = vec![vec![0; width as usize]; height as usize];
    for g in &guards {
        let (x, y) = g.simulate_steps(100);
        grid[y as usize][x as usize] = 1;
        if x < g.width / 2 && y < g.height / 2 {
            quads[0] += 1;
        } else if x > g.width / 2 && y < g.height / 2 {
            quads[1] += 1;
        } else if x < g.width / 2 && y > g.height / 2 {
            quads[2] += 1;
        } else if x > g.width / 2 && y > g.height / 2 {
            quads[3] += 1;
        }
    }
    // for row in grid {
    //     for cell in row {
    //         print!("{:3}", cell);
    //     }
    //     println!();
    // }
    // println!("{:?}", quads);
    let mut mult = 1;
    for q in quads {
        mult *= q;
    }
    Ok(mult)
}

fn solve2(input: &str) -> Result<i32, std::io::Error>{
    let width = 101;
    let height = 103;
    let mut guards: Vec<Guard> = Vec::new();
    for line in input.lines() {
        guards.push(Guard::new(line, width, height));
    }
    'search:
    for i in 1.. {
        let mut grid = vec![vec![0; width as usize]; height as usize];
        for g in &guards {
            let (x, y) = g.simulate_steps(i);
            grid[y as usize][x as usize] = 1;
        }
        for row in &grid {
            if row.iter().map(|&c| if c > 0 { 1 } else { 0 }).sum::<i32>() > 30 {
                println!("{} seconds:", i);
                for row in grid {
                    for cell in row {
                        let c = if cell > 0 { '*' } else { ' ' };
                        print!("{}", c);
                    }
                    println!();
                }
                sleep(std::time::Duration::from_secs(1));
                continue 'search;
            }
        }
    }
    Ok(0)
}

#[cfg(test)]
mod tests {
    use super::*;

    const SAMPLE_INPUT: &str = "p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3";

    #[test]
    fn test_part1() {
        assert_eq!(solve1_internal(SAMPLE_INPUT, 11, 7).unwrap(), 12);
    }
}
