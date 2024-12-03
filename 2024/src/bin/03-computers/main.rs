use clap::Parser;
use regex::Regex;
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
    let mut s = String::new();
    input.read_to_string(&mut s)?;
    let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();
    let sum = re.captures_iter(s.as_str()).map(|cap| {
        let a = cap[1].parse::<i32>().unwrap();
        let b = cap[2].parse::<i32>().unwrap();
        a * b
    }).sum();
    Ok(sum)
}

fn solve2<R: BufRead>(_: &mut R) -> Result<i32, std::io::Error>{
    unimplemented!()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_part1() {
        let input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))";
        assert_eq!(solve1(&mut BufReader::new(input.as_bytes())).unwrap(), 161);
    }
}
