use clap::Parser;
use regex::Regex;
use std::collections::HashMap;
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
    let re = Regex::new(r"(\d+)\s+(\d+)").unwrap();
    let mut a: Vec<i32> = Vec::new();
    let mut b: Vec<i32> = Vec::new();
    for line in input.lines() {
        let line = line?;
        let caps = re.captures(&line).unwrap();
        a.push(caps[1].parse().unwrap());
        b.push(caps[2].parse().unwrap());
    }
    a.sort();
    b.sort();
    let mut sum = 0;
    for i in 0..a.len() {
        let diff = (a[i] - b[i]).abs();
        sum += diff;
    }
    Ok(sum)
}

fn solve2<R: BufRead>(input: &mut R) -> Result<i32, std::io::Error>{
    let re = Regex::new(r"(\d+)\s+(\d+)").unwrap();
    let mut a: Vec<i32> = Vec::new();
    let mut b: HashMap<i32, i32> = HashMap::new();
    for line in input.lines() {
        let line = line?;
        let caps = re.captures(&line).unwrap();
        a.push(caps[1].parse().unwrap());
        *b.entry(caps[2].parse().unwrap()).or_default() += 1;
    }
    let mut sum = 0;
    for value in a {
        if let Some(&count) = b.get(&value) {
            sum += value * count;
        }
    }
    Ok(sum)
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_part1() {
        let input = "3   4
4   3
2   5
1   3
3   9
3   3";
        assert_eq!(solve1(&mut BufReader::new(input.as_bytes())).unwrap(), 11);
    }

    #[test]
    fn test_part2() {
        let input = "3   4
4   3
2   5
1   3
3   9
3   3";
        assert_eq!(solve2(&mut BufReader::new(input.as_bytes())).unwrap(), 31);
    }
}
