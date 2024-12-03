use clap::Parser;
use std::io::{BufRead, BufReader};
use std::fs::File;

#[derive(Parser)]
struct Cli {
    #[arg(short, long, value_parser = clap::value_parser!(u8).range(1..=2), default_value_t = 1)]
    part: u8,

    input: String
}

fn main() -> Result<(), std::io::Error>{
    let cli = Cli::parse();
    let result = match cli.part {
        1 => solve1(&mut BufReader::new(File::open(cli.input)?))?,
        2 => solve2(&mut BufReader::new(File::open(cli.input)?))?,
        _ => unreachable!(),
    };
    print!("{result}");
    Ok(())
}

fn solve1<R: BufRead>(input: &mut R) -> Result<i32, std::io::Error>{
    let mut safe = 0; 
    for line in input.lines() {
        let levels = line?
            .split(' ')
            .map(|s| s.parse::<i32>().unwrap())
            .collect::<Vec<i32>>();
        let mut diffs = levels.iter().zip(levels.iter().skip(1)).map(|(&x, &y)| y - x);
        let increasing = diffs.clone().next().unwrap() > 0;
        if diffs.all(|x| {
            x.abs() >= 1 && x.abs() <= 3 &&
            if increasing {
                x > 0
            } else {
                x < 0
            }
        }) {
            safe += 1;
        }
    }
    Ok(safe)
}

fn solve2<R: BufRead>(input: &mut R) -> Result<i32, std::io::Error>{
    let mut safe = 0; 
'outer:
    for line in input.lines() {
        let levels = line?
            .split(' ')
            .map(|s| s.parse::<i32>().unwrap())
            .collect::<Vec<i32>>();
        for skip in 0..levels.len() {
            let filtered = levels.iter().enumerate().filter(|&(i, _)| i != skip).map(|(_, x)| *x);
            let diffs = filtered.clone().zip(filtered.clone().skip(1)).map(|(x, y)| y - x);
            let increasing = diffs.clone().next().unwrap() > 0;
            if diffs.clone().all(|x| {
                x.abs() >= 1 && x.abs() <= 3 &&
                if increasing {
                    x > 0
                } else {
                    x < 0
                }
            }) {
                safe += 1;
                continue 'outer;
            }
        }
    }
    Ok(safe)
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_part1() {
        let input = "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9";
        assert_eq!(solve1(&mut BufReader::new(input.as_bytes())).unwrap(), 2);
    }

    #[test]
    fn test_part2() {
        let input = "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9";
        assert_eq!(solve2(&mut BufReader::new(input.as_bytes())).unwrap(), 4);
    }
}
