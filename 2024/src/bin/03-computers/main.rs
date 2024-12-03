use clap::Parser;
use regex::Regex;
use std::fs::read_to_string;

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
    print!("{result}");
    Ok(())
}

fn solve1(input: &str) -> Result<i32, std::io::Error>{
    let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();
    let sum = re.captures_iter(input).map(|cap| {
        let a = cap[1].parse::<i32>().unwrap();
        let b = cap[2].parse::<i32>().unwrap();
        a * b
    }).sum();
    Ok(sum)
}

fn solve2(input: &str) -> Result<i32, std::io::Error>{
    let re = Regex::new(r"mul\((\d+),(\d+)\)|do\(\)|don't\(\)").unwrap();
    let mut enabled = true;
    let mut sum = 0;
    re.captures_iter(input).for_each(|cap| {
        match &cap[0] {
            "do()" => enabled = true,
            "don't()" => enabled = false,
            _ => {
                if enabled {
                    let a = cap[1].parse::<i32>().unwrap();
                    let b = cap[2].parse::<i32>().unwrap();
                    sum += a * b
                }
            }
        }
    });
    Ok(sum)
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_part1() {
        let input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))";
        assert_eq!(solve1(input).unwrap(), 161);
    }

    #[test]
    fn test_part2() {
        let input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))";
        assert_eq!(solve2(input).unwrap(), 48);
    }
}
